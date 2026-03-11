package todo

import (
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	var list List

	list.Add("Learn Go")

	if len(list) != 1 {
		t.Errorf("expected list length 1, got %d", len(list))
	}

	if list[0].Task != "Learn Go" {
		t.Errorf("expected task 'Learn Go', got %s", list[0].Task)
	}

	if list[0].Done != false {
		t.Errorf("expected Done to be false")
	}
}

func TestComplete(t *testing.T) {
	var list List

	list.Add("Learn Go")

	err := list.Complete(0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	
	if list[0].Task != "Learn Go" {
		t.Fatalf("wrong task completed")
	}

	if list[0].Done != true {
		t.Errorf("expected task to be completed")
	}
}

func TestDelete(t *testing.T) {
	var list List

	list.Add("Task 1")
	list.Add("Task 2")

	err := list.Delete(0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(list) != 1 {
		t.Errorf("expected list length 1, got %d", len(list))
	}

	
	if list[0].Task != "Task 2" {
		t.Errorf("wrong task deleted")
	}
}

func TestSaveAndGet(t *testing.T) {
	tf, err := os.CreateTemp("", "todo_test_*.json")
	if err != nil {
		t.Fatalf("error creating temp file: %v", err)
	}

	filename := tf.Name()
	tf.Close()

	
	defer os.Remove(filename)

	
	var list1 List
	list1.Add("Task A")
	list1.Add("Task B")
	list1.Complete(1)

	err = list1.Save(filename)
	if err != nil {
		t.Fatalf("error saving list: %v", err)
	}

	
	var list2 List
	err = list2.Get(filename)
	if err != nil {
		t.Fatalf("error getting list: %v", err)
	}

	if len(list1) != len(list2) {
		t.Fatalf("lists have different lengths")
	}

	for i := range list1 {

		
		if list1[i].Task != list2[i].Task {
			t.Fatalf("tasks are different")
		}

		
		if list1[i].Done != list2[i].Done {
			t.Errorf("done status mismatch for task %s", list1[i].Task)
		}
	}
}
func TestString(t *testing.T) {

	var list List

	list.Add("Task 1")
	list.Add("Task 2")

	list.Complete(0)

	expected :=
		"- [X] 0: Task 1\n" +
			"- [ ] 1: Task 2\n"

	result := list.String()

	if result != expected {
		t.Errorf("expected:\n%s\ngot:\n%s", expected, result)
	}
}