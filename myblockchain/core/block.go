package core

//引用必须的标准库
import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

//最简单的区块定义
//区块头
type BlockHead struct {
	Timestamp     int64  //时间戳，也就是该区块的创建时间
	PrevBlockHash []byte //上一个区块的哈希值
	Hash          []byte //本区块的哈希值
	Nonce         int    //难度值
}

//区块
type Block struct {
	Head BlockHead //区块头
	Data []byte    //本区块实际记录的信息
}

//生成新块时，只需 Data 与 PrevBlockHash
//当前块的哈希值会基于 Data 和 PrevBlockHash 计算得到
func NewBlock(data string, prevBlockHash []byte) *Block {
	head := BlockHead{
		Timestamp:     time.Now().Unix(),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		Nonce:         0}

	block := &Block{
		Head: head,
		Data: []byte(data)}

	block.SetHash()

	return block
}

// 设置当前块哈希值，采用公式：sha256(PrevBlockHash + Data + Timestamp)
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Head.Timestamp, 10))
	headers := bytes.Join([][]byte{b.Head.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Head.Hash = hash[:]
}

//创世块，空数据
func NewGenesisBlock() *Block {
	genesisHash := "201711151421"
	block := NewBlock("This is my block chain.", []byte(genesisHash))

	return block
}
