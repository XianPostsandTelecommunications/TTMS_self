package work

import "sync"

type Worker struct {
	l          sync.RWMutex
	taskChan   chan func()
	workerChan chan func()
}
