package todo

import (
	"fmt"
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

func (l *List) Delete(i int) error {
	if i < 1 || i > len(*l) {
		return fmt.Errorf("unvalid item %d", i)
	}

	index := i - 1
	ls := *l
	*l = append(ls[:index], ls[index+1:]...)
	return nil
}
