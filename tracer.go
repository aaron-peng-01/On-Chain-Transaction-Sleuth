package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

type CallFrame struct {
	Type    string      `json:"type"`
	From    string      `json:"from"`
	To      string      `json:"to"`
	Value   string      `json:"value,omitempty"`
	Gas     string      `json:"gas"`
	GasUsed string      `json:"gasUsed"`
	Input   string      `json:"input"`
	Output  string      `json:"output,omitempty"`
	Calls   []CallFrame `json:"calls,omitempty"` // 递归结构，用于表示子调用
}


// getTransactionTrace 调用 debug_traceTransaction
func getTransactionTrace(client *rpc.Client, txHash string) (*CallFrame, error) {
	var result CallFrame
	txHashCommon := common.HexToHash(txHash)

	// Geth的 `debug_traceTransaction` 允许传入一个JS脚本作为Tracer。
	// `callTracer` 是内建tracer。
	err := client.CallContext(context.Background(), &result, "debug_traceTransaction", txHashCommon, map[string]string{"tracer": "callTracer"})
	if err != nil {
		return nil, fmt.Errorf("failed to call debug_traceTransaction: %w", err)
	}

	return &result, nil
}
