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

func (k *Kademlia) Ping(recipient Contact) error {
	client, err := k.Dial(recipient)
	if err != nil {
		return err
	}
	return client.Call("KademliaRPC.PingRPC", k.NewPingArgs(), k.NewPingReply())
}

func (k *Kademlia) PingRPC(req *PingArgs, res *PingReply) {
	k.HandleRPC(&req.RPCHeader, &req.RPCHeader)
}
