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

func (c Contact) String() string {
	return c.ID.String()
}

type Contacts []Contact

func (c Contacts) Len() int           { return len(c) }
func (c Contacts) Less(i, j int) bool { return c[i].ID.Less(c[j].ID) }
func (c Contacts) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func (h *Contacts) Push(e interface{}) { *h = append(*h, e.(Contact)) }

func (h *Contacts) Pop() (e interface{}) {
	e = (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return e
}
