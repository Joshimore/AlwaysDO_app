package todo

import (
	"fmt"
	"time"
)

func NewTask(id int, title, category string) Task {

	task := Task{
		ID:        id,
		Title:     title,
		Category:  category,
		CreatedAt: time.Now(),
	}

	return task

}

func NextID(tasks []Task) int {

	maxID := 0

	for i := 0; i < len(tasks); i++ {
		fmt.Println(tasks[i].ID)

		if tasks[i].ID > maxID {
			maxID = tasks[i].ID
		}
	}

	return maxID + 1

}

//func FindTask(tasks []Task, id int) (int, bool){}
