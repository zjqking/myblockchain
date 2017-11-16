package web

//简单演示，不使用第三方WEB框架包

import (
	"fmt"
	"net/http"

	"zjq.com/myblockchain/core"
)

func sayHelloService(rsp http.ResponseWriter, req *http.Request) {
	rsp.Write([]byte("Welcome to my block chain!"))
}

func initService(rsp http.ResponseWriter, req *http.Request) {
	rsp.Write([]byte("This is init!\n"))

	core.NewBlockChain()

	rsp.Write([]byte("Blockchain init finish!\n"))
}

func addService(rsp http.ResponseWriter, req *http.Request) {
	rsp.Write([]byte("This is add!\n"))

	bc := core.GetInstance()
	if nil == bc {
		rsp.Write([]byte("Init blockchain first"))
		return
	}

	n := core.GetBlockNum()
	message := fmt.Sprintf("This is %d block.", n)
	bc.AddBlock(message)
	n++
	core.SetBlockNum(n)

	rsp.Write([]byte("Block add finish!\n"))
}

func listService(rsp http.ResponseWriter, req *http.Request) {
	rsp.Write([]byte("Blockchain data begin:"))

	bc := core.GetInstance()
	if nil == bc {
		rsp.Write([]byte("Init blockchain first"))
		return
	}

	var message string
	for _, block := range bc.Blocks {
		message = fmt.Sprintf("Prev hash: %x\n", block.Head.PrevBlockHash)
		message += fmt.Sprintf("Data: %s\n", block.Data)
		message += fmt.Sprintf("Hash: %x\n", block.Head.Hash)
		rsp.Write([]byte(message))
	}

	rsp.Write([]byte("Blockchain data end."))
}

//简单网络服务
func WebRun() {
	http.HandleFunc("/", sayHelloService)
	http.HandleFunc("/init", initService)
	http.HandleFunc("/add", addService)
	http.HandleFunc("/list", listService)

	err := http.ListenAndServe(":9080", nil)
	if nil != err {
		panic(err)
	}
}
