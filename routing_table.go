package kademlia

import "sort"

type RoutingTable struct {
	self     Contact
	kBuckets [IDLen]*KBucket
}

func NewRoutingTable(self Contact) *RoutingTable {
	buckets := [IDLen]*KBucket{}
	for i := range buckets {
		buckets[i] = NewKBucket()
	}
	return &RoutingTable{
		self:     self,
		kBuckets: buckets,
	}
}

func (rt *RoutingTable) Update(contact Contact) {
	bucketIndex := rt.self.ID.BucketIndex(contact.ID)
	rt.kBuckets[bucketIndex].Update(contact)
}

// Finds n closest Contacts to the target node
func (rt *RoutingTable) FindClosest(target NodeID, n int) (contacts []Contact) {
	bucketIndex := rt.self.ID.BucketIndex(target)
	rt.appendContactsFromBucket(bucketIndex, &contacts)

	// if we've found fewer than n contacts, append others in any order
	for i := 0; i < len(rt.kBuckets) && len(contacts) < n; i++ {
		if i == bucketIndex {
			continue
		}
		rt.appendContactsFromBucket(i, &contacts)
	}

	contacts = contacts[:n]       // remove any extra contacts
	sort.Sort(Contacts(contacts)) // sort by distance so closestNode is a front
	return contacts
}

func (rt *RoutingTable) appendContactsFromBucket(bucket int, contacts *[]Contact) {
	elem := rt.kBuckets[bucket].Front()
	for ; elem != nil; elem = elem.Next() {
		*contacts = append(*contacts, elem.Value.(Contact))
	}
}
