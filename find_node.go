package kademlia

import (
	"container/heap"
	"sort"
)

// Know as alpha in the spec.
const MaxRequests = 3

type FindNodeArgs struct {
	RPCHeader
	Target NodeID
}

type FindNodeReply struct {
	RPCHeader
	Contacts []Contact
}

func (k *Kademlia) NewFindNodeArgs(target NodeID) *FindNodeArgs {
	return &FindNodeArgs{
		RPCHeader: k.NewRPCHeaderWithID(),
		Target:    target,
	}
}

func (k *Kademlia) NewFindNodeReply() *FindNodeReply {
	return &FindNodeReply{Contacts: []Contact{}}
}

// Makes RPC Call to request list of nodes closest to target from individual contact
func (k *Kademlia) FindNode(recipient Contact, target NodeID, found chan *FindNodeReply) error {
	client, err := k.Dial(recipient)
	if err != nil {
		return err
	}
	args, reply := k.NewFindNodeArgs(target), k.NewFindNodeReply()
	err = client.Call("Kademlia.FindNodeRPC", args, reply)

	if err == nil {
		reply.Sender = recipient
	} else {
		k.table.Update(reply.Sender)
	}
	found <- reply
	return nil
}

// Handler for FindNode RPC
func (k *Kademlia) FindNodeRPC(req *FindNodeArgs, resp *FindNodeReply) error {
	err := k.HandleRPC(&req.RPCHeader, &resp.RPCHeader)
	if err != nil {
		return err
	}
	resp.Contacts = k.table.FindClosest(req.Target, BucketSize)
	return nil
}

// Iteratively searches Node Tree for target node
func (k *Kademlia) SearchForNode(target NodeID, n int) (shortList []Contact) {
	seen := make(map[string]bool) // key: nodeID.String()
	found := make(chan *FindNodeReply)

	contactsPQ := Contacts(k.table.FindClosest(target, MaxRequests))
	closest := contactsPQ[0]

	// Send async FindRPCs to each initial contact
	pending, replyCount := 0, 0
	for contactsPQ.Len() > 0 {
		contact := heap.Pop(&contactsPQ).(Contact)
		seen[contact.String()] = true
		pending++
		go k.FindNode(contact, target, found)
	}

	for pending > 0 {
		reply := <-found
		pending--
		if len(reply.Contacts) > 0 { // discard unresponsive recipients
			shortList = append(shortList, reply.Sender)
			replyCount++
		}

		didImprove := false
		for _, contact := range reply.Contacts {
			if _, haveSeen := seen[contact.String()]; !haveSeen {
				seen[contact.String()] = true
				heap.Push(&contactsPQ, contact)
				if contact.ID.Less(closest.ID) {
					closest = contact
					didImprove = true
				}
			}
		}

		finishedRequesting := replyCount >= n && !didImprove
		for pending < MaxRequests && contactsPQ.Len() > 0 && !finishedRequesting {
			pending++
			contact := heap.Pop(&contactsPQ).(Contact)
			go k.FindNode(contact, target, found)
		}

		if finishedRequesting && pending > 0 {
			go func() { // Drain chanel for any oustanding requests
				for pending > 0 {
					<-found
					pending--
				}
			}()
			break
		}
	}

	sort.Sort(Contacts(shortList))
	shortList = shortList[:n]
	return shortList
}
