package worker

import "errors"

var ErrQueueFull = errors.New("worker queue full")
