package store

import (
	"errors"

	"github.com/taaag51/go_todo_app/entitiy"
)

var (
	Tasks = &TaskStore{Tasks: map[int]*entitiy.Task{}}

	ErrNotFound = errors.New("not found")
)

type TaskStore struct {
	// 動作確認用の仮実装なのであえてexportしてる
	LastID entitiy.TaskID
	Tasks  map[entitiy.TaskID]*entitiy.Task
}

func (ts *TaskStore) Add(t *entitiy.Task) (int, error) {
	ts.LastID++
	t.ID = ts.LastID
	ts.Tasks[t.ID] = t
	return t.ID, nil

}

// All はソート済みタスク一覧を返す
func (ts *TaskStore) All() entitiy.Tasks {
	tasks := make([]*entitiy.Task, len(ts.Tasks))
	for i, t := range ts.Tasks {
		tasks[i-1] = t
	}
	return tasks
}
