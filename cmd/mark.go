package cmd

import (
	"fmt"
	"strconv"

	"todo.go/model"
	"todo.go/service"
	"todo.go/utils"
)

func HandleMark(filename string, args []string) {
	if len(args) > 3 || len(args) < 2 {
		fmt.Println("Usage todo mark <ID> <status>")
		fmt.Println()
		fmt.Println("Status commands: ")
		fmt.Println(" todo")
		fmt.Println(" in-progress")
		fmt.Println(" done")
	}

	markFilters := map[string]model.Status{
		"todo":        model.StatusNotDone,
		"in-progress": model.StatusInProgress,
		"done":        model.StatusDone,
	}

	id, err:= strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Invalid ID format")
		return
	}
	
	if  status, ok := markFilters[args[2]]; ok {
		err := service.MarkTaskStatusById(filename, status, id)
		utils.PrintErr(err)

		fmt.Println("Status updated successfully"); 
		return
	}

	fmt.Println("No such command exist")
}
