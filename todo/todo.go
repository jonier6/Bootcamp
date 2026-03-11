package todo

import (
	"fmt"
	"time"
	"encoding/json"
	"os"
)

// type item private
type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// type list public
type List []item


func (l *List) Add(task string) {
	newItem := item{
		Task:      task,
		Done:      false,
		CreatedAt: time.Now(),
	}

	*l = append(*l, newItem)
}


func (l *List) Complete(index int) error {
	if index < 0 || index >= len(*l) {
		return fmt.Errorf("invalid index")
	}

	(*l)[index].Done = true
	(*l)[index].CompletedAt = time.Now()

	return nil
}


func (l *List) Delete(index int) error {
	if index < 0 || index >= len(*l) {
		return fmt.Errorf("invalid index")
	}

	*l = append((*l)[:index], (*l)[index+1:]...)
	return nil
}


func (l *List) Save(filename string) error {
	data, err := json.Marshal(l)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}


func (l *List) Get(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// Do not unmarshall if it is empty
	if len(data) == 0 {
		return nil
	}

	
	err = json.Unmarshal(data, l)
	if err != nil {
		return err
	}

	return nil
}

func (l List) String() string {

	var result string

	for i, task := range l {

		if task.Done {
			result += fmt.Sprintf("- [X] %d: %s\n", i, task.Task)
		} else {
			result += fmt.Sprintf("- [ ] %d: %s\n", i, task.Task)
		}

	}

	return result
}