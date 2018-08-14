package handlers

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/likecoin/likechain/abci/context"
	"github.com/likecoin/likechain/abci/types"
)

func TestCheckTransfer(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	ctx := context.NewMockContext(mockCtrl)
	// TODO: mock ctx calls

	rawTx := &types.Transaction{}
	res := checkTransfer(ctx, rawTx)
	t.Log(res)
	// TODO
}

func TestDeliverTransfer(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	ctx := context.NewMockContext(mockCtrl)
	// TODO: mock ctx calls

	rawTx := &types.Transaction{}
	res := deliverTransfer(ctx, rawTx)
	t.Log(res)
	// TODO
}

func TestValidateTransferSignature(t *testing.T) {
	tx := &types.TransferTransaction{}
	if !validateTransferSignature(tx.Sig) {
		t.Error("Validate TransferSignature failed")
	}
}

func TestValidateTransferTransaction(t *testing.T) {
	tx := &types.TransferTransaction{}
	if !validateTransferTransaction(tx) {
		t.Error("Validate TransferTransaction failed")
	}
}

func TestTransfer(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	ctx := context.NewMockContext(mockCtrl)
	// TODO: mock ctx calls

	tx := &types.TransferTransaction{}
	transfer(ctx, tx)
	// TODO
}