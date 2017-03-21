package kademlia

import (
	"reflect"
	"testing"
)

func TestFindeClosest(t *testing.T) {
	ourself := ContactFromHexString("1000000000000000000000000000000000000000")
	rt := NewRoutingTable(ourself)
	target := ContactFromHexString("2333333300000000000000000000000000000000")
	closest := ContactFromHexString("2222222200000000000000000000000000000000")

	rt.Update(ContactFromHexString("1111111100000000000000000000000000000000"))
	rt.Update(closest)
	rt.Update(ContactFromHexString("3333333300000000000000000000000000000000"))
	rt.Update(ContactFromHexString("4444444400000000000000000000000000000000"))

	numberToFind := 1
	got := rt.FindClosest(target.ID, numberToFind)
	if len(got) != 1 {
		t.Error("FindClosest should return the specified amount of contacts")
	}

	if !reflect.DeepEqual(got[0], closest) {
		t.Error("Closest should return the closestNode to the target in the 0th index")
	}

	numberToFind = 2
	got = rt.FindClosest(target.ID, numberToFind)
	if len(got) != 2 {
		t.Error("FindClosest should return the specified amount of contacts")
	}

	if !reflect.DeepEqual(got[0], closest) {
		t.Error("Closest should return the closestNode to the target in the 0th index")
	}
}

func ContactFromHexString(id string) Contact {
	node, _ := NewNodeID(id)
	return Contact{ID: node}
}
