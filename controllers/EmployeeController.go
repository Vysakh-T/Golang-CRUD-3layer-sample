package controllers

import (
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"net/http" // used to access the request and response object of the api
	"strconv"

	"src/sampleCRUD/middlewares" // models package where User schema is defined
	"src/sampleCRUD/models"
	"src/sampleCRUD/services"

	// package used to covert string into int type

	"github.com/gorilla/mux" // used to get the params from the route

	_ "github.com/lib/pq" // postgres golang driver
)

// Router is exported and used in main.go
func EmployeeController() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/employee/{id}", getEmployee).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/employee", getAllEmployee).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/employee", createEmployee).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/employee/{id}", updateEmployee).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/employee/{id}", deleteEmployee).Methods("DELETE", "OPTIONS")

	return router
}

func createEmployee(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var employee models.Employee

	err := json.NewDecoder(r.Body).Decode(&employee)

	if err != nil {
		panic("Unable to decode the request body")
	}

	insertID := services.InsertEmployee(employee)

	res := middlewares.Response{
		ID:      insertID,
		Message: "User created successfully",
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// GetUser will return a single employee by its id
func getEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get the userid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		panic("Unable to convert the string into int")
	}

	// call the getUser function with employee id to retrieve a single employee
	employee := services.GetEmployee(int64(id))

	// send the response
	json.NewEncoder(w).Encode(employee)
}

// GetAllUser will return all the users
func getAllEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get all the users in the db
	employees := services.GetAllEmployees()

	// send all the users as response
	json.NewEncoder(w).Encode(employees)
}

// UpdateUser update employee's detail in the postgres db
func updateEmployee(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the userid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		panic("Unable to convert the string into int")
	}

	// create an empty employee of type models.Employee
	var employee models.Employee

	// decode the json request to employee
	err = json.NewDecoder(r.Body).Decode(&employee)

	if err != nil {
		panic("Unable to decode the request body")
	}

	// call update employee to update the employee
	updatedRows := services.UpdateEmployee(int64(id), employee)

	// format the message string
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := middlewares.Response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// DeleteUser delete employee's detail in the postgres db
func deleteEmployee(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the userid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id in string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		panic("Unable to convert the string into int")
	}

	// call the deleteUser, convert the int to int64
	deletedRows := services.DeleteEmployee(int64(id))

	// format the message string
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := middlewares.Response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}
