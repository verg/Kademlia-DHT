package kademlia

import "container/list"

const BucketSize = 20 // refered to as "k" in the spec

type KBucket struct {
	*list.List
}

func NewKBucket() *KBucket {
	return &KBucket{list.New()}
}

func (k *KBucket) Update(contact Contact) {
	elem := k.find(contact)
	if elem != nil {
		k.MoveToBack(elem)
	} else if !k.Full() { // TODO Handle case when full
		k.PushBack(contact)
	}
}

func (k *KBucket) find(contact Contact) *list.Element {
	for elem := k.Front(); elem != nil; elem = elem.Next() {
		if contact.ID == elem.Value.(Contact).ID {
			return elem
		}
	}
	return nil
}

func (k *KBucket) Full() bool {
	return k.Len() >= BucketSize
}
