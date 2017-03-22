package kademlia

type PingArgs struct {
	RPCHeader
}

type PingReply struct {
	RPCHeader
}

func (k *Kademlia) NewPingArgs() *PingArgs {
	return &PingArgs{k.NewRPCHeaderWithID()}
}

func (k *Kademlia) NewPingReply() *PingReply {
	return &PingReply{}
}

// Makes Ping RPC Call
func (k *Kademlia) Ping(recipient Contact) error {
	client, err := k.Dial(recipient)
	if err != nil {
		return err
	}
	return client.Call("KademliaRPC.PingRPC", k.NewPingArgs(), k.NewPingReply())
}

// Handles Ping RPC Call
func (k *Kademlia) PingRPC(req *PingArgs, resp *PingReply) {
	k.HandleRPC(&req.RPCHeader, &resp.RPCHeader)
}
