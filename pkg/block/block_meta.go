package block

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/tendermint/tendermint/pkg/meta"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

// BlockMeta contains meta information.
type BlockMeta struct {
	BlockID   meta.BlockID `json:"block_id"`
	BlockSize int          `json:"block_size"`
	Header    meta.Header  `json:"header"`
	NumTxs    int          `json:"num_txs"`
}

// NewBlockMeta returns a new BlockMeta.
func NewBlockMeta(block *Block, blockParts *meta.PartSet) *BlockMeta {
	return &BlockMeta{
		BlockID:   meta.BlockID{block.Hash(), blockParts.Header()},
		BlockSize: block.Size(),
		Header:    block.Header,
		NumTxs:    len(block.Data.Txs),
	}
}

func (bm *BlockMeta) ToProto() *tmproto.BlockMeta {
	if bm == nil {
		return nil
	}

	pb := &tmproto.BlockMeta{
		BlockID:   bm.BlockID.ToProto(),
		BlockSize: int64(bm.BlockSize),
		Header:    *bm.Header.ToProto(),
		NumTxs:    int64(bm.NumTxs),
	}
	return pb
}

func BlockMetaFromProto(pb *tmproto.BlockMeta) (*BlockMeta, error) {
	if pb == nil {
		return nil, errors.New("blockmeta is empty")
	}

	bm := new(BlockMeta)

	bi, err := meta.BlockIDFromProto(&pb.BlockID)
	if err != nil {
		return nil, err
	}

	h, err := meta.HeaderFromProto(&pb.Header)
	if err != nil {
		return nil, err
	}

	bm.BlockID = *bi
	bm.BlockSize = int(pb.BlockSize)
	bm.Header = h
	bm.NumTxs = int(pb.NumTxs)

	return bm, bm.ValidateBasic()
}

// ValidateBasic performs basic validation.
func (bm *BlockMeta) ValidateBasic() error {
	if err := bm.BlockID.ValidateBasic(); err != nil {
		return err
	}
	if !bytes.Equal(bm.BlockID.Hash, bm.Header.Hash()) {
		return fmt.Errorf("expected BlockID#Hash and Header#Hash to be the same, got %X != %X",
			bm.BlockID.Hash, bm.Header.Hash())
	}
	return nil
}
