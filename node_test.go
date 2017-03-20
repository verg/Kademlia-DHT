package kademlia

import "testing"

func TestNewNode(t *testing.T) {
	id := "0123456789abcdef0123456789abcdef01234567"
	node, err := NewNodeID(id)
	if err != nil || node.String() != id {
		t.Errorf("Expected %s, got: %s", id, node.String())
	}
}

func TestXor(t *testing.T) {
	a := NodeID{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	b := NodeID{1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 17}
	expected := NodeID{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2}
	got := a.Xor(b)
	if got != expected {
		t.Errorf("got: %s, expected: %s", got.String(), expected.String())
	}
}

func TestLess(t *testing.T) {
	smallest := NodeID{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	middle := NodeID{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 20}
	largest := NodeID{1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 17}
	if smallest.Less(smallest) {
		t.Error("Same NodeID should not be less than self")
	}
	if !smallest.Less(middle) {
		t.Error("smallest should be Less than middle")
	}
	if middle.Less(smallest) {
		t.Error("Middle should not be Less than smallest")
	}

	if !smallest.Less(largest) {
		t.Error("smallest should be Less than largest")
	}
	if largest.Less(smallest) {
		t.Error("Largest should not be Less than smallest")
	}
}

func TestBucketIndex(t *testing.T) {
	a := NodeID{}
	maxBucketID := IDLen*8 - 1
	if a.BucketIndex() != maxBucketID {
		t.Errorf("Got: %d Expected %d", a.BucketIndex(), maxBucketID)
	}
	b := NodeID{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	if b.BucketIndex() != 7 {
		t.Errorf("Got: %d Expected %d", b.BucketIndex(), 7)
	}
}
