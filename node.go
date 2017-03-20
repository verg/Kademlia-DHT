package kademlia

import (
	"encoding/hex"
	"errors"
	"math/rand"
)

const IDLen = 20

type NodeID [IDLen]byte

// Create new Node with hex-encoded ID
func NewNodeID(h string) (res NodeID, e error) {
	id, err := hex.DecodeString(h)
	if err != nil {
		return res, errors.New("Error decoding hex-encoded ID")
	}
	for i, b := range id {
		res[i] = b
	}
	return res, nil
}

func NewRandomNodeID() (res NodeID) {
	for i := range res {
		res[i] = byte(rand.Intn(256))
	}
	return res
}

func (n NodeID) Xor(other NodeID) (res NodeID) {
	for i, b := range n {
		res[i] = b ^ other[i]
	}
	return res
}

// Determined by number of leading zeros
func (n NodeID) BucketIndex() (bucket int) {
	for i, b := range n {
		for j := 0; j < 8; j++ {
			if (b>>uint8(7-j))&0x1 != 0 {
				return i*8 + j
			}
		}
	}
	return IDLen*8 - 1
}

func (n NodeID) Less(other interface{}) bool {
	for i, b := range n {
		if b != other.(NodeID)[i] {
			return b < other.(NodeID)[i]
		}
	}
	return false
}

func (n NodeID) String() string {
	return hex.EncodeToString(n[:])
}
