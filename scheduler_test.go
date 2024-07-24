package gotask

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestOnce(t *testing.T) {
	t1, err := NewOnceTask()
	if err != nil {
		t.Fatal(err)
	}
	sch := NewScheduler()
	ctx := context.Background()

	testFunc := func(ctx context.Context) error {
		log.Println("testFunc")
		return nil
	}
	t1.TaskFunc = testFunc
	sch.AddTask(t1)
	// sch.AddFunc(testFunc, t1)
	err = sch.Start(ctx)
	if err != nil {
		t.Fatal(err)
	}
	//defer sch.Stop()
}

func TestScheduler(t *testing.T) {
	scheduler := NewScheduler()
	t1, err := NewOnceTask()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Run OnceTask", func(t *testing.T) {
		// Channel for orchestrating when the task ran
		doneCh := make(chan struct{})
		testFunc := func(ctx context.Context) error {
			log.Println("testFunc")
			return nil
		}
		// Setup A task
		scheduler.AddFunc(testFunc, t1)
		// Start the scheduler
		scheduler.Start(context.Background())
		// Make sure it runs especially when we want it too
		for i := 0; i < 6; i++ {
			select {
			case <-doneCh:
				continue
			case <-time.After(2 * time.Second):
				t.Errorf("Scheduler failed to execute the scheduled tasks %d run within 2 seconds", i)
			}
		}
	})
	t.Run("Run OnceTask With Delay", func(t *testing.T) {
		t2, err := NewOnceTask(WithDelay(time.Second * 4))
		if err != nil {
			t.Fatal(err)
		}
		doneCh := make(chan struct{})
		testFunc := func(ctx context.Context) error {
			log.Println("testFunc")
			return nil
		}
		// Setup A task
		scheduler.AddFunc(testFunc, t2)
		// Start the scheduler
		scheduler.Start(context.Background())
		// Make sure it runs especially when we want it too
		for i := 0; i < 6; i++ {
			select {
			case <-doneCh:
				continue
			case <-time.After(2 * time.Second):
				t.Errorf("Scheduler failed to execute the scheduled tasks %d run within 2 seconds", i)
			}
		}
	})
}
