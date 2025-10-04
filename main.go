package main

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gin-gonic/gin"
)

const (
	// 强烈建议从环境变量读取，这里为了演示方便直接写入
	// e.g., "https://eth-mainnet.g.alchemy.com/v2/YOUR_ALCHEMY_KEY"
	nodeUrl = "YOUR_ETHEREUM_NODE_RPC_URL"
)

func main() {
	// rpc.Dial会建立原始的RPC连接，我们需要它来调用debug模块
	rpcClient, err := rpc.Dial(nodeUrl)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// ethclient.NewClient是基于RPC连接的封装，方便调用标准接口
	// 虽然我们主要用rpcClient，但保留它也很方便
	ethClient := ethclient.NewClient(rpcClient)
	_ = ethClient // 避免 "unused variable" 错误

	log.Println("Successfully connected to Ethereum node.")

	r := gin.Default()
	
	// 设置API路由
	r.GET("/trace", func(c *gin.Context) {
		txHash := c.Query("tx_hash")
		if txHash == "" {
			c.JSON(400, gin.H{"error": "tx_hash is required"})
			return
		}
		
		// 在这里调用我们的追踪逻辑
		// trace, err := getTransactionTrace(rpcClient, txHash)
		// ...
		
		c.JSON(200, gin.H{
			"message": "trace feature coming soon",
			"tx_hash": txHash,
		})
	})

	r.Run(":8080") // 启动服务在 8080 端口
}
