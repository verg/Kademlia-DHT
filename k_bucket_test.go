package kademlia

import (
	"reflect"
	"testing"
)

func TestFullKBucket(t *testing.T) {
	b := NewKBucket()
	for i := 0; i < BucketSize; i++ {
		if b.Full() {
			t.Errorf("Bucket should not be full before adding %d Contacts", BucketSize)
		}
		c := BuildContact()
		b.Update(c)
	}

	if !b.Full() {
		t.Errorf("Bucket should be full before adding %d Contacts", BucketSize)
	}
}

func TestUpdateExistingContact(t *testing.T) {
	b := NewKBucket()
	firstContact := BuildContact()
	secondContact := BuildContact()
	thirdContact := BuildContact()
	b.Update(firstContact)
	b.Update(secondContact)
	b.Update(thirdContact)

	front := b.Front().Value.(Contact)
	if !reflect.DeepEqual(front, firstContact) {
		t.Error("Least recently added Contact should be at Front of the list")
	}
	b.Update(firstContact)
	if b.Len() != 3 {
		t.Error("Updating existing contact should not change size")
	}

	back := b.Back().Value.(Contact)
	if !reflect.DeepEqual(back, firstContact) {
		t.Error("Recently updated Contact should be at Front of the list")
	}

}

func BuildContact() Contact {
	return Contact{ID: NewRandomNodeID()}
}
