package gotask

import "time"

type option struct {
	Delay   time.Duration
	ErrFunc func(error)
	TaskFunc
}

// Option the tracer provider option
type Option func(opt *option) error

// WithDelay set the delay time
func WithDelay(delay time.Duration) Option {
	return func(opt *option) error {
		opt.Delay = delay
		return nil
	}
}

// WithTaskFunc set the task func
func WithTaskFunc(taskFunc TaskFunc) Option {
	return func(opt *option) error {
		opt.TaskFunc = taskFunc
		return nil
	}
}

// WithErrFunc set the error func
func WithErrFunc(errfunc func(error)) Option {
	return func(opt *option) error {
		opt.ErrFunc = errfunc
		return nil
	}
}
