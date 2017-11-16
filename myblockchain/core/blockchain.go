package core

//简单的区块链
type BlockChain struct {
	Blocks []*Block
}

var bc *BlockChain
var blockNum int

//产生一个新的区块链
//默认产生一个新的创世块
func NewBlockChain() *BlockChain {
	blockNum = 0
	bc = &BlockChain{[]*Block{NewGenesisBlock()}}
	return bc
}

//增加一个区块
func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Head.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

//返回全局变量
func GetInstance() *BlockChain {
	return bc
}

//设置blockNum
func SetBlockNum(n int) {
	blockNum = n
}

//返回blockNum
func GetBlockNum() int {
	return blockNum
}
