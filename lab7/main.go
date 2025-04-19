package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Employee struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
	Position  string `json:"position"`
}

var employees []Employee
var mu sync.Mutex
var nextId int = 0

func getEmployees(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var newEmployee Employee
	if err := json.NewDecoder(r.Body).Decode(&newEmployee); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newEmployee.ID = nextId
	nextId++
	employees = append(employees, newEmployee)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newEmployee)
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedEmployee Employee
	if err := json.NewDecoder(r.Body).Decode(&updatedEmployee); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, emp := range employees {
		if emp.ID == id {
			employees[i].FirstName = updatedEmployee.FirstName
			employees[i].LastName = updatedEmployee.LastName
			employees[i].Age = updatedEmployee.Age
			employees[i].Position = updatedEmployee.Position
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(employees[i])
			return
		}
	}

	http.Error(w, "Employee not found", http.StatusNotFound)
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, emp := range employees {
		if emp.ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Employee not found", http.StatusNotFound)
}

func main() {
	r := mux.NewRouter()

	employees = append(employees, Employee{ID: 1, FirstName: "John", LastName: "Doe", Age: 30, Position: "Developer"})
	employees = append(employees, Employee{ID: 2, FirstName: "Jane", LastName: "Smith", Age: 35, Position: "Manager"})

	nextId = len(employees) + 1

	r.HandleFunc("/employees", getEmployees).Methods("GET")
	r.HandleFunc("/employee", addEmployee).Methods("POST")
	r.HandleFunc("/employees/update", updateEmployee).Methods("PATCH")
	r.HandleFunc("/employees/delete", deleteEmployee).Methods("DELETE")

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	}).Handler(r)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", handler)
}
