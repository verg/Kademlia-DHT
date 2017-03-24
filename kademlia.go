package kademlia

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"time"

	database "github.com/syndtr/goleveldb/leveldb"
)

const DB_Path = "db/"

type RPCIdentifier [IDLen]byte

type Kademlia struct {
	table *RoutingTable
	db    *database.DB
}

type KademliaRPC struct {
	kademlia *Kademlia
}

type RPCHeader struct {
	Sender Contact
	RPCID  RPCIdentifier
}

func NewKademlia(self Contact) *Kademlia {
	conn, err := database.OpenFile(DB_Path+self.String(), nil)
	if err != nil {
		panic("Failed to open db conn")
	}
	return &Kademlia{table: NewRoutingTable(self), db: conn}
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
	conn, err := net.DialTimeout("tcp", contact.Addr, 4*time.Second)
	if err != nil {
		return nil, err
	}
	return rpc.NewClient(conn), nil
}

func (k *Kademlia) HandleRPC(request *RPCHeader, resp *RPCHeader) error {
	if request.RPCID != resp.RPCID {
		msg := fmt.Sprintf("RPCID Mismatch. Expected: %s, Got: %s",
			request.RPCID, resp.RPCID)
		return errors.New(msg)
	}

	k.table.Update(request.Sender)
	resp.Sender = k.table.self
	return nil
}
