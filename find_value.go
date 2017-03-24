package kademlia

type FindValueArgs struct {
	RPCHeader
	Target NodeID
}

type FindValueReply struct {
	RPCHeader
	Contacts []Contact
	Value    string
}

func (k *Kademlia) NewFindValueArgs(target NodeID) *FindValueArgs {
	return &FindValueArgs{
		RPCHeader: k.NewRPCHeaderWithID(),
		Target:    target,
	}
}

func NewFindValueReply() *FindValueReply {
	return &FindValueReply{}
}

// Asks a node for a value.
// Returns the value if the node has it. Otherwise returns Contacts as in FindNode
func (k *Kademlia) FindValue(contact Contact, target NodeID) (string, []Contact, error) {
	client, err := k.Dial(contact)
	if err != nil {
		return "", nil, err
	}
	args := k.NewFindValueArgs(target)
	reply := NewFindValueReply()
	if err := client.Call("KademliaRPC.FindValueRPC", &args, &reply); err != nil {
		return "", nil, err
	}
	return reply.Value, reply.Contacts, nil
}

func (k *Kademlia) FindValueRPC(req *FindValueArgs, resp *FindValueReply) error {
	err := k.HandleRPC(&req.RPCHeader, &resp.RPCHeader)
	if err != nil {
		return err
	}
	value, err := k.db.Get(req.Target[:], nil)

	if err != nil {
		panic("Read from db failed")
	}
	if value == nil {
		resp.Contacts = k.table.FindClosest(req.Target, BucketSize)
	} else {
		resp.Value = string(value)
	}

	return nil
}
