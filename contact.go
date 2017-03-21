package kademlia

type Contact struct {
	ID   NodeID
	Addr string
}

func NewContact(id NodeID, addr string) Contact {
	return Contact{
		ID:   id,
		Addr: addr,
	}
}

type Contacts []Contact

func (c Contacts) Len() int           { return len(c) }
func (c Contacts) Less(i, j int) bool { return c[i].ID.Less(c[j].ID) }
func (c Contacts) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
