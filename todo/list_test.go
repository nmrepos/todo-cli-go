package todo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddAndTasks(t *testing.T) {
	l := NewList()
	// initially empty
	if got := l.Tasks(); len(got) != 0 {
		t.Fatalf("expected empty list, got %v", got)
	}

	// add tasks
	id1 := l.Add("task1")
	if id1 != 1 {
		t.Errorf("expected id 1, got %d", id1)
	}
	id2 := l.Add("task2")
	if id2 != 2 {
		t.Errorf("expected id 2, got %d", id2)
	}

	tasks := l.Tasks()
	want := []Task{
		{ID: 1, Text: "task1", Done: false},
		{ID: 2, Text: "task2", Done: false},
	}
	if !reflect.DeepEqual(tasks, want) {
		t.Errorf("Tasks() = %v; want %v", tasks, want)
	}
}

func TestComplete(t *testing.T) {
	l := NewList()
	l.Add("task")

	// complete existing task
	ok := l.Complete(1)
	if !ok {
		t.Errorf("Complete(1) = false; want true")
	}
	tasks := l.Tasks()
	if !tasks[0].Done {
		t.Errorf("after Complete, Done = %v; want true", tasks[0].Done)
	}

	// complete non-existent
	if ok := l.Complete(99); ok {
		t.Errorf("Complete(99) = true; want false")
	}
}

func TestRemove(t *testing.T) {
	l := NewList()
	// add three tasks
	for i := 1; i <= 3; i++ {
		l.Add(fmt.Sprintf("task%d", i))
	}

	// remove middle task
	ok := l.Remove(2)
	if !ok {
		t.Errorf("Remove(2) = false; want true")
	}
	tasks := l.Tasks()
	want := []Task{
		{ID: 1, Text: "task1", Done: false},
		{ID: 3, Text: "task3", Done: false},
	}
	if !reflect.DeepEqual(tasks, want) {
		t.Errorf("Tasks after Remove = %v; want %v", tasks, want)
	}

	// remove non-existent
	if ok := l.Remove(99); ok {
		t.Errorf("Remove(99) = true; want false")
	}
}

func TestTasksImmutability(t *testing.T) {
	l := NewList()
	l.Add("immutable")

	tasks := l.Tasks()
	// modify returned slice
	tasks[0].Done = true

	tasks2 := l.Tasks()
	if tasks2[0].Done {
		t.Errorf("Tasks returned slice is not a copy; internal state mutated")
	}
}
