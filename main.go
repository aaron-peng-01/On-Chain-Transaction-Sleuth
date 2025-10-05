package main

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gin-gonic/gin"
)

const (
	// e.g., "https://eth-mainnet.g.alchemy.com/v2/YOUR_ALCHEMY_KEY"
	nodeUrl = "YOUR_ETHEREUM_NODE_RPC_URL"
)

func main() {
	// rpc.Dial建立原始的RPC连接，使其调用debug模块
	rpcClient, err := rpc.Dial(nodeUrl)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	ethClient := ethclient.NewClient(rpcClient)
	_ = ethClient // 避免 "unused variable" 错误

	log.Println("Successfully connected to Ethereum node.")

	r := gin.Default()
	
	r.GET("/trace", func(c *gin.Context) {
		txHash := c.Query("tx_hash")
		if txHash == "" {
			c.JSON(400, gin.H{"error": "tx_hash is required"})
			return
		}

		trace, err := getTransactionTrace(rpcClient, txHash)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		
		c.JSON(200, trace)
	})

	r.Run(":8080")
}
