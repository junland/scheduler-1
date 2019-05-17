package scheduler_test

import (
	"github.com/qapquiz/scheduler"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewScheduler(t *testing.T) {
	actualScheduler := scheduler.NewScheduler()
	expectedScheduler := &scheduler.Scheduler{}

	assert.NotEqual(t, expectedScheduler, actualScheduler)
}

func TestScheduler_StartJob(t *testing.T) {
	isRun := false

	actualScheduler := scheduler.NewScheduler()
	jobID := actualScheduler.StartJob(time.Millisecond, func() { isRun = true })

	time.Sleep(2 * time.Millisecond)

	assert.True(t, isRun)
	assert.NotEqual(t, 0, jobID)
}

func TestScheduler_StopJob(t *testing.T) {
	isRun := false

	actualScheduler := scheduler.NewScheduler()
	jobID := actualScheduler.StartJob(time.Millisecond, func() { isRun = true })

	err := actualScheduler.StopJob(jobID)

	assert.Nil(t, err)
	assert.False(t, isRun)
}

func TestScheduler_StopJob_MustError(t *testing.T) {
	isRun := false

	actualScheduler := scheduler.NewScheduler()
	_ = actualScheduler.StartJob(time.Millisecond, func() { isRun = true })

	err := actualScheduler.StopJob(uint(10))

	assert.NotNil(t, err)
	assert.False(t, isRun)
}
