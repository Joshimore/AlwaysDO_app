package storage

import (
	"errors"
	"alwaysdo/internal/todo"
	"encoding/json"
	"os"
)

func SaveTasks(path string, tasks []todo.Task) error{
	
	jsondata, err := json.MarshalIndent(tasks, "", "  ")

	if err != nil {
		return err
	}

	err = os.WriteFile(path, jsondata, 0600)

	if err != nil{
		return err
	}

	return nil

}

func LoadTasks(path string) ([]todo.Task, error){

	data, err := os.ReadFile(path)
	if err != nil{
		if errors.Is(err, os.ErrNotExist){
			return []todo.Task{}, nil
		}

		return nil, err
	}


	var tasks []todo.Task
	err = json.Unmarshal(data, &tasks)
	if err != nil{
		return nil, err
	}

	return tasks, nil	
}
