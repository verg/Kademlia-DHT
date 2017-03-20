package kademlia

import "net"

type Contact struct {
	ID   NodeID
	Addr net.IP
}

func NewContact(id NodeID, addr net.IP) Contact {
	return Contact{
		ID:   id,
		Addr: addr,
	}
}
