package kademlia

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type RPCIdentifier [IDLen]byte

type Kademlia struct {
	table *RoutingTable
}

type KademliaRPC struct {
	kademlia *Kademlia
}

type RPCHeader struct {
	Sender Contact
	RPCID  RPCIdentifier
}

func NewKademlia(self Contact) *Kademlia {
	return &Kademlia{table: NewRoutingTable(self)}
}

func (k *Kademlia) NewRPCHeaderWithID() RPCHeader {
	return RPCHeader{
		Sender: k.table.self,
		RPCID:  RPCIdentifier(IDLenRandBytes()),
	}
}

func (k *Kademlia) Serve() error {
	rpc.Register(&KademliaRPC{k})
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", k.table.self.Addr)
	if err != nil {
		return err
	}
	go http.Serve(listener, nil)
	return nil
}

func (k *Kademlia) Dial(contact Contact) (*rpc.Client, error) {
	client, err := rpc.DialHTTP("tcp", contact.Addr)
	return client, err
}

func (k *Kademlia) HandleRPC(request *RPCHeader, resp *RPCHeader) error {
	if request.RPCID != resp.RPCID {
		msg := fmt.Sprintf("RPCID Mismatch. Expected: %s, Got: %s",
			request.RPCID, resp.RPCID)
		return errors.New(msg)
	}

	k.table.Update(request.Sender)
	return nil
}
