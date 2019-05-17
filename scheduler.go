package scheduler

import (
	"fmt"
	"sync"
	"time"
)

// Scheduler will receive your func and run at the other time the you want
type Scheduler struct {
	jobID uint
	jobs map[uint]*time.Timer
	jobsRWMutex sync.RWMutex
}

// NewScheduler will return newly created scheduler
func NewScheduler() *Scheduler {
	return &Scheduler{
		jobID: 0,
		jobs: make(map[uint]*time.Timer),
	}
}

// StartJob will add job to the map and return jobID
func (scheduler *Scheduler) StartJob(delay time.Duration, f func()) uint {
	scheduler.jobID++

	scheduler.jobsRWMutex.Lock()
	scheduler.jobs[scheduler.jobID] = time.AfterFunc(delay, f)
	scheduler.jobsRWMutex.Unlock()

	return scheduler.jobID
}

// StopJob will receive jobID and return bool for showing that it success or not
func (scheduler *Scheduler) StopJob(jobID uint) error {
	scheduler.jobsRWMutex.RLock()
	defer scheduler.jobsRWMutex.RUnlock()
	timer, ok := scheduler.jobs[jobID]
	if !ok {
		return fmt.Errorf("not found this job id: %d", jobID)
	}

	timer.Stop()
	return nil
}
