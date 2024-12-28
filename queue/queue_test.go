package queue

import (
	"fmt"
	"testing"

	"github.com/GirigiriG/DST/adt"
)

func TestEnqueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(&adt.Node{Value: 3})

	if v, _ := q.Dequeue(); v.Value != 3 {
		t.Errorf("expected %v got %v", 3, v)
	}

	fmt.Println(q.IsEmpty())
}
func TestDequeue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(&adt.Node{Value: 3})
	q.Dequeue()
	fmt.Println(q.IsEmpty())
}
