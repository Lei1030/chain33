package blockchain

import (
	"testing"
	"time"
)
import "math/rand"

func TestTask(t *testing.T) {
	task := newTask(time.Second)

	if task.InProgress() {
		t.Error("task not start")
		return
	}
	task.Start(1, 10)
	perm := rand.Perm(10)
	for i := 0; i < len(perm); i++ {
		time.Sleep(time.Millisecond * 500)
		task.Done(int64(perm[i]) + 1)
		if i < len(perm)-1 && task.InProgress() == false {
			t.Error("task not done, but InProgress is false")
			return
		}
		if i == len(perm)-1 && task.InProgress() == true {
			t.Error("task is done, but InProgress is true")
		}
	}
}