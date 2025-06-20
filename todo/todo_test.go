package todo

import "testing"

func TestAddTask(t *testing.T) {
	l := NewList()
	id := l.Add("task one")
	if id != 1 {
		t.Fatalf("expected id 1 got %d", id)
	}
	if len(l.tasks) != 1 {
		t.Fatalf("expected 1 task got %d", len(l.tasks))
	}
}

func TestCompleteTask(t *testing.T) {
	l := NewList()
	id := l.Add("task")
	ok := l.Complete(id)
	if !ok {
		t.Fatal("complete returned false")
	}
	if !l.tasks[0].Done {
		t.Fatal("task not marked done")
	}
}

func TestRemoveTask(t *testing.T) {
	l := NewList()
	id := l.Add("task")
	ok := l.Remove(id)
	if !ok {
		t.Fatal("remove returned false")
	}
	if len(l.tasks) != 0 {
		t.Fatalf("expected 0 tasks got %d", len(l.tasks))
	}
}

func TestTasksReturnsCopy(t *testing.T) {
	l := NewList()
	l.Add("a")
	list1 := l.Tasks()
	list1[0].Text = "b"
	if l.tasks[0].Text != "a" {
		t.Fatal("Tasks did not return copy")
	}
}
