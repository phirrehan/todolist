package db

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

func hashString(s string) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])
}

func checkFileNotExists() bool {
	_, err := os.Stat("./data.json")
	if err != nil {
		return true
	}
	return false
}

func createFile() {
	file, err := os.Create("data.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
	}
	defer file.Close()

	bytes, err := file.WriteString("[]")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
	}

	fmt.Printf("Wrote %d Number of Bytes\n", bytes)
}

func parseJson(slicePtr *[]Todo) error {
	fileData, err := os.ReadFile("data.json")
	if err != nil {
		return fmt.Errorf("Error reading json: %v", err)
	}

	if err := json.Unmarshal(fileData, slicePtr); err != nil {
		return fmt.Errorf("Error parsing json: %v", err)
	}
	return nil
}

func encodeJson(slice []Todo) error {
	newJson, err := json.Marshal(slice)
	if err != nil {
		return fmt.Errorf("Error encoding json: %v", err)
	}

	newFile, err := os.Create("data.json")
	if err != nil {
		return fmt.Errorf("Error creating new file: %v", err)
	}
	defer newFile.Close()

	newFile.Write(newJson)
	return nil
}

func searchTodos(todos []Todo, searchId string) int {
	for i, todo := range todos {
		if todo.Id == searchId {
			return i
		}
	}
	return -1
}

func GetTodos() []Todo {
	var todos []Todo
	err := parseJson(&todos)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting Todos: %v", err)
	}
	return todos
}

func AddTodo(description string) {
	//
	if checkFileNotExists() {
		createFile()
	}

	var todos []Todo
	if err := parseJson(&todos); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return
	}

	newTodo := Todo{
		Id:          hashString(description),
		Description: description,
		Status:      false,
	}
	todos = append(todos, newTodo)
	if err := encodeJson(todos); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
}

func EditTodo(id string, description string, status bool) {
	if checkFileNotExists() {
		fmt.Fprintf(os.Stderr, "Error editing item: data.json does not exist\n")
		return
	}

	var todos []Todo
	parseJson(&todos)

	indexToEdit := searchTodos(todos, id)
	if indexToEdit == -1 {
		fmt.Fprintf(os.Stderr, "Specified ID not found\n")
		return
	}
	todos[indexToEdit] = Todo{
		Id:          id,
		Description: description,
		Status:      status,
	}

	if err := encodeJson(todos); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return
	}
}

func DeleteTodo(id string) {
	if checkFileNotExists() {
		fmt.Fprintf(os.Stderr, "Error deleting item: data.json does not exist\n")
		return
	}

	var todos []Todo
	parseJson(&todos)

	indexToRemove := searchTodos(todos, id)
	if indexToRemove == -1 {
		fmt.Fprintf(os.Stderr, "Specified ID does not exist\n")
		return
	}
	todos = append(todos[:indexToRemove], todos[indexToRemove+1:]...)

	if err := encodeJson(todos); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return
	}
}
