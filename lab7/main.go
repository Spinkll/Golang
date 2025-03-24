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

// Структура для співробітника
type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
}

// Масив співробітників і м'ютекс для потокобезпечного доступу
var employees []Employee
var mu sync.Mutex

// Отримати всіх співробітників
func getEmployees(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

// Додати нового співробітника
func addEmployee(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var newEmployee Employee
	// Розпаковуємо JSON з тіла запиту в структуру
	if err := json.NewDecoder(r.Body).Decode(&newEmployee); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Визначаємо унікальний ID
	newEmployee.ID = len(employees) + 1
	employees = append(employees, newEmployee)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newEmployee)
}

// Оновити інформацію про співробітника
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
			employees[i].Name = updatedEmployee.Name
			employees[i].Position = updatedEmployee.Position
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(employees[i])
			return
		}
	}

	http.Error(w, "Employee not found", http.StatusNotFound)
}

// Видалити співробітника
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
	// Ініціалізуємо початкових співробітників
	employees = append(employees, Employee{ID: 1, Name: "John Doe", Position: "Developer"})
	employees = append(employees, Employee{ID: 2, Name: "Jane Smith", Position: "Manager"})

	// Визначаємо маршрути для маршрутизатора mux
	r.HandleFunc("/employees", getEmployees).Methods("GET")
	r.HandleFunc("/employees/add", addEmployee).Methods("POST")
	r.HandleFunc("/employees/update", updateEmployee).Methods("PUT")
	r.HandleFunc("/employees/delete", deleteEmployee).Methods("DELETE")

	// Додаємо підтримку CORS для всіх запитів
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	}).Handler(r)

	// Запускаємо сервер на порту 8080
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", handler)
}
