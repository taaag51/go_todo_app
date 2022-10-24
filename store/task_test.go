package store

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/taaag51/go_todo_app/testutil"
)

func TestRepository_ListTasks(t *testing.T) {
	ctx := context.Background()
	// entity.Taskを作成する他のテストケースと混ぜるとテストがフェイルする
	// そのため、トランザクションを張ることでこのテストケースの中だけのテーブル状態にする
	tx, err := testutil.OpenDBForTest(t).BeginTxx(ctx, nil)

	// このテストケースが完了したら元に戻す
	t.Cleanup(func() { _ = tx.Rollback() })
	if err != nil {
		t.Fatal(err)
	}
	wants := prepareTasks(ctx, t, tx)

	sut := &Repositry{}

	gots, err := sut.ListTasks(ctx, tx)
	if err != nil {
		t.Fatal("unexpected err: %v", err)
	}
	if d := cmp.Diff(gots, wants); len(d) != 0 {
		t.Errorf("differs: (-got +want)\n%s", d)
	}

}
