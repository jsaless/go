package main

import (
	"fmt"
	"net/http"
)

var taskItems = []string{"Learn Go"}

func main() {
	http.HandleFunc("/", helloFunc)
	http.HandleFunc("/tasks", showTasks)

	http.ListenAndServe(":8080", nil)
}

func showTasks(writer http.ResponseWriter, _ *http.Request) {
	for index, taskItem := range taskItems {
		fmt.Fprintf(writer, "%d. %s\n", index+1, taskItem)
	}
}

func helloFunc(writer http.ResponseWriter, _ *http.Request) {
	greeting := "Hello everybody! Welcome to ToDo App"
	fmt.Fprintln(writer, greeting)
}
