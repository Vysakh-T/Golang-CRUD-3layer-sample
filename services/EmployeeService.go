package services

import (
	"log"

	"src/sampleCRUD/models"
	"src/sampleCRUD/repositories"
)

func InsertEmployee(employee models.Employee) int64 {

	insertID := repositories.InsertEmployee(employee)
	return insertID

}

func GetEmployee(id int64) models.Employee {

	employee, err := repositories.Getemployee(int64(id))

	if err != nil {
		log.Fatalf("Unable to get employee. %v", err)
	}

	return employee

}

func GetAllEmployees() []models.Employee {

	employees, err := repositories.GetAllemployees()

	if err != nil {
		log.Fatalf("Unable to get all employee. %v", err)
	}

	return employees
}

func UpdateEmployee(id int64, employee models.Employee) int64 {

	updatedRows := repositories.Updateemployee(int64(id), employee)

	// fmt.Println(employee)

	return updatedRows
}

func DeleteEmployee(id int64) int64 {

	deletedRows := repositories.Deleteemployee(int64(id))
	return deletedRows

}
