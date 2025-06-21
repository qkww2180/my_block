package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64
)

const targetBits = 8 // 目标难度，表示哈希值的前导零数量;值越大，难度越高

type ProofOfWork struct {
	block  *Block
	target *big.Int // 任何有效的哈希值，当被看作一个大整数时，都必须小于这个 target。
}

// NewProofOfWork builds and returns a ProofOfWork
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	// 将 target 左移 (256 - targetBits) 位
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

// prepareData prepares the data for hashing
func (pow *ProofOfWork) prepareDataForHash(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			Int64ToBytes(pow.block.Timestamp),
			Int64ToBytes(int64(targetBits)),
			Int64ToBytes(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

// Run performs a proof-of-work
// and returns the nonce that satisfies the target and the hash of the block
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)

	for nonce < maxNonce {
		data := pow.prepareDataForHash(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 { // 如果哈希值小于目标值，则成功
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	data := pow.prepareDataForHash(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(pow.target) == -1 // 如果哈希值小于目标值，则验证通过
}
