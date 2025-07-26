package data

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
)

type todo struct {
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

func parseJson(slicePtr *[]todo) error {
	fileData, err := os.ReadFile("data.json")
	if err != nil {
		return fmt.Errorf("Error reading json: %v", err)
	}

	jsonerr := json.Unmarshal(fileData, slicePtr)
	if jsonerr != nil {
		return fmt.Errorf("Error parsing json: %v", jsonerr)
	}
	return nil
}

func reWriteJson(slice []todo) error {
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

func searchTodos(todos []todo, searchId string) int {
	for i, todo := range todos {
		if todo.Id == searchId {
			return i
		}
	}
	return -1
}

func AddTodo(description string) {
	//
	if checkFileNotExists() {
		createFile()
	}

	var todos []todo
	if err := parseJson(&todos); err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}

	newTodo := todo{
		Id:          hashString(description),
		Description: description,
		Status:      false,
	}
	todos = append(todos, newTodo)
	if err := reWriteJson(todos); err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
}

func EditTodo(id string, description string, status bool) {
	if checkFileNotExists() {
		fmt.Fprintf(os.Stderr, "Error editing item: data.json does not exist")
		return
	}

	var todos []todo
	parseJson(&todos)

	indexToEdit := searchTodos(todos, id)
	if indexToEdit == -1 {
		fmt.Fprintf(os.Stderr, "Error searching id in data: no entry of %s found in data.json", id)
		return
	}
	todos[indexToEdit] = todo{
		Id:          id,
		Description: description,
		Status:      status,
	}

	if err := reWriteJson(todos); err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
}

func DeleteTodo(id string) {
	if checkFileNotExists() {
		fmt.Fprintf(os.Stderr, "Error deleting item: data.json does not exist")
		return
	}

	var todos []todo
	parseJson(&todos)

	indexToRemove := searchTodos(todos, id)
	if indexToRemove == -1 {
		fmt.Fprintf(os.Stderr, "Error searching id in data: no entry of %s found in data.json", id)
		return
	}
	todos = append(todos[:indexToRemove], todos[indexToRemove+1:]...)

	if err := reWriteJson(todos); err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
}
