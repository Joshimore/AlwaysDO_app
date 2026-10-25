package todo

import (
	"fmt"
	"time"
)

func NewTask(id int, title, category string) Task{

	task := Task{
		ID: id, 
		Title: title, 
		Category: category,
		CreatedAt: time.Now(),
	}

	return task

}

func NextID(task []Task) int{

	maxID := 0
	
	for i := 0; i < len(task); i++{
		fmt.Println(task[i].ID)

		if task[i].ID > maxID{
			maxID = task[i].ID;
		}	
	}
	
	return maxID + 1

}
