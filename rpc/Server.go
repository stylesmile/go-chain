package main

import (
	"../core"
	"encoding/json"
	"io"
	"net/http"
)

var blockchain *core.Blockchain

// http server
func run() {
	// 对外暴露端口
	http.HandleFunc("/blockchain/get", blockchainGetHandler)
	http.HandleFunc("/blockchain/write", blockchainWriteHandler)
	// 监听端口8080
	http.ListenAndServe("localhost:8080", nil)
}

// 读数据
func blockchainGetHandler(w http.ResponseWriter, r *http.Request) {
	bytes, error := json.Marshal(blockchain)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}
	// 返回json 字节
	io.WriteString(w, string(bytes))
}

// 写数据
func blockchainWriteHandler(w http.ResponseWriter, r *http.Request) {
	blockData := r.URL.Query().Get("data")
	blockchain.SendData(blockData)
	blockchainGetHandler(w, r)
}

func main() {
	blockchain = core.NewBlockchain()
	run()
}
