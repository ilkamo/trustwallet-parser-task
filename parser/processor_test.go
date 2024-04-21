package parser

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/ilkamo/ethparser-go/tests/mock"
	"github.com/ilkamo/ethparser-go/types"
)

func TestParser_processBlock(t *testing.T) {
	log := &mock.Logger{}
	ethMock := mock.EthereumClient{
		MostRecentBlock: 2,
		BlockByNumber:   types.Block{},
	}

	p := NewParser(
		ethMock,
		log,
		time.Millisecond,
		mock.TransactionsRepository{},
		mock.ObserverRepository{
			WantError: errors.New("observer error"),
		},
	)
	require.NotNil(t, p)

	tx := types.Transaction{
		Hash:  "0x005295d8C90Fe127932C6fE78daE6D5a4B975098",
		From:  "0x995295d8C90Fe127932C6fE78daE6D5a4B975098",
		To:    "0x225295d8C90Fe127932C6fE78daE6D5a4B975098",
		Value: "0x0123",
	}

	err := p.processBlock(context.Background(), types.Block{Transactions: []types.Transaction{tx}})
	require.Error(t, err)
}