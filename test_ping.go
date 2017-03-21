package kademlia

import "testing"

const LocalHost = "127.0.0.1:8989"

func TestPing(t *testing.T) {
	ourself := BuildContactWithAddr(LocalHost)
	k := NewKademlia(ourself)
	k.Serve()

	other := BuildContactWithAddr(LocalHost)
	err := k.Ping(other)
	if err != nil {
		t.Error(err)
	}
}

func BuildContactWithAddr(addr string) Contact {
	c := BuildContact()
	c.Addr = addr
	return c
}
