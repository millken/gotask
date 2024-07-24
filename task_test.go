package gotask

import (
	"testing"
	"time"
)

func TestOnceTask(t *testing.T) {
	t1, err := NewOnceTask()
	if err != nil {
		t.Fatal(err)
	}
	n1 := t1.Next()
	if n1 != RunNow {
		t.Fatalf("expect %v, but got %v", RunNow, n1)
	}
	n1 = t1.Next()
	if n1 != Zero {
		t.Fatalf("expect %v, but got %v", Zero, n1)
	}

	t2, err := NewOnceTask(WithDelay(10 * time.Second))
	if err != nil {
		t.Fatal(err)
	}
	n2 := t2.Next()
	if n2 != time.Now().Unix()+10 {
		t.Fatalf("expect %v, but got %v", time.Now().Unix()+10, n2)
	}
	n2 = t2.Next()
	if n2 != Zero {
		t.Fatalf("expect %v, but got %v", Zero, n2)
	}
}
