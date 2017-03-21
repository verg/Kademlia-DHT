package kademlia

import (
	"encoding/hex"
	"errors"
	"fmt"
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
	if len(id) > IDLen {
		return res, errors.New(fmt.Sprintf("ID longer than %d bytes", IDLen))
	}
	for i, b := range id {
		res[i] = b
	}
	return res, nil
}

func NewRandomNodeID() (res NodeID) {
	return NodeID(IDLenRandBytes())
}

func IDLenRandBytes() (b [IDLen]byte) {
	for i := range b {
		b[i] = byte(rand.Intn(256))
	}
	return b
}

func (n NodeID) Xor(other NodeID) (res NodeID) {
	for i, b := range n {
		res[i] = b ^ other[i]
	}
	return res
}

// Returns the index of the target node. Determined by the number of leading 0s.
// Can be commputed from either node since Xor is communicative.
func (n NodeID) BucketIndex(target NodeID) (bucket int) {
	distance := n.Xor(target)
	for i, b := range distance {
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
