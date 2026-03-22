package cmd

import (
	"fmt"

	"todo.go/model"
	"todo.go/service"
	"todo.go/utils"
)

func HandleList(filename string, args []string) {

	if len(args) < 2 {
		tasks, err := service.ListTask(filename)
		utils.PrintErr(err)

		utils.PrintTaskList(tasks)
		return
	}

	var listFilters = map[string]model.Status{
		"todo":        model.StatusNotDone,
		"in-progress": model.StatusInProgress,
		"done":        model.StatusDone,
	}

	if status, ok := listFilters[args[1]]; ok {
		tasks, err := service.ListByStatus(filename, status)
		utils.PrintErr(err)

		utils.PrintTaskList(tasks)
		return
	}

	fmt.Println("Invalid command")

}
