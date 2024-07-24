package gotask

import (
	"context"
	"time"
)

type Scheduler struct {
	tasks []Task
	queue []Task
}

// NewScheduler creates a new scheduler instance. This is the main entry point for the package.
func NewScheduler() *Scheduler {
	s := &Scheduler{
		tasks: make([]Task, 0),
	}
	return s
}

// Add adds a new task to the scheduler. This function returns the task ID for the new task.
func (s *Scheduler) AddFunc(taskFunc TaskFunc, task Task) {
	task.SetTaskFunc(taskFunc)
	s.tasks = append(s.tasks, task)
}

func (s *Scheduler) AddTask(task Task) {
	s.tasks = append(s.tasks, task)
}

func (s *Scheduler) AddTasks(tasks []Task) {
	s.tasks = append(s.tasks, tasks...)
}

// Start starts the scheduler. This function blocks until the scheduler is stopped.
func (s *Scheduler) Start(ctx context.Context) error {
	s.queue = make([]Task, len(s.tasks))
	tick := time.Tick(time.Second)
	go func() {
		select {
		case <-ctx.Done():
			return
		case t := <-tick:
			s.tick(ctx, t)
		}
	}()
	return nil
}

// safeOps safely change task's data
func (s *Scheduler) safeOps(f func()) {
	// t.Lock()
	// defer t.Unlock()

	f()
}

func (s *Scheduler) tick(ctx context.Context, t time.Time) {
	var next int64
	tn := t.Unix()
	tasks := s.tasks
	for k, t := range tasks {
		if t == nil {
			continue
		}
		next = t.Next()
		if next == RunNow {
			go t.Func(ctx)
			continue
		}
		time.AfterFunc(time.Duration(next-tn)*time.Second, func() {
			s.tasks[k] = t
			s.queue[k] = nil
			go t.Func(ctx)
		})
		s.tasks[k] = nil
		s.queue[k] = t
	}
}

func (s *Scheduler) run(t Task) {
	// err := t.Func(context.Background())
	// if err != nil && t.ErrFunc != nil {
	// 	t.ErrFunc(err)
	// }
}
