package storage

import (
	"alwaysdo/internal/todo"
	"encoding/json"
	"os"
)

func SaveTasks(path string, tasks []todo.Task) error {

	jsondata, err := json.MarshalIndent(tasks, "", "  ")

	if err != nil {
		return err
	}

	err = os.WriteFile(path, jsondata, 0600)

	if err != nil {
		return err
	}

	return nil

}
