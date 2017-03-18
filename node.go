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

func (n NodeID) String() string {
	return hex.EncodeToString(n[:])
}
