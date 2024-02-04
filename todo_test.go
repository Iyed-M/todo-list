package todo_test

import (
	"testing"

	"github.com/Iyed-M/todo-list"
)

func TestAdd(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)

	if taskName != l[0].Task {
		t.Errorf("Expected %q got %q", taskName, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}

	taskName := "NewTask"
	l.Add(taskName)

	if taskName != l[0].Task {
		t.Errorf("Expected %q got %q", taskName, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("New Task should not be completed")
	

	l.Complete(0)
	if !l[0].Done {
		t.Errorf("New Task should be completed")
	}
}

func TestDeletet *Testing.T) {
	l  := todo.List{}
	taskNames :=  []string{"Task1", "Task2", "Task3"

	for _,  v := range taskNames {
		l.Add(v)
	}

	if l[0].Task != taskNames[0] {
		t.Errorf("Expected %q got %q instead", taskNames[0], l[0].Task)
	}
}
