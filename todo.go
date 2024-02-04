package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, t)
}

func (l *List) Complete(i int) error {
	index := i
	ls := *l

	if index < 0 || index > len(ls) {
		return fmt.Errorf("unvalid item %d", i)
	}

	ls[index].CompletedAt = time.Now()
	ls[index].Done = true

	return nil
}

func (l *List) Delete(i int) error {
	if i < 1 || i > len(*l) {
		return fmt.Errorf("unvalid item %d", i)
	}

	index := i - 1
	ls := *l
	*l = append(ls[:index], ls[index+1:]...)
	return nil
}

func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, js, 0644)
}

func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}
