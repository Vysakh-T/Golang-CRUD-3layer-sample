package repositories

import (
	"database/sql"
	"fmt"

	"src/sampleCRUD/middlewares"
	"src/sampleCRUD/models"
)

// insert one employee in the DB

func InsertEmployee(employee models.Employee) int64 {

	db := middlewares.CreateConnection()

	defer db.Close()

	sqlStatement := `INSERT INTO newEmployee (name, age) VALUES ($1, $2) RETURNING id`

	var id int64

	err := db.QueryRow(sqlStatement, employee.Name, employee.Age).Scan(&id)

	if err != nil {
		panic("Unable to execute the query")
	}

	fmt.Printf("Inserted a single record %v", id)

	return id
}

// get one employee from the DB by its employeeid

func Getemployee(id int64) (models.Employee, error) {

	db := middlewares.CreateConnection()

	defer db.Close()

	var employee models.Employee

	sqlStatement := `SELECT * FROM newEmployee WHERE id=$1`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&employee.ID, &employee.Name, &employee.Age)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return employee, nil
	case nil:
		return employee, nil
	default:
		panic("Unable to scan the row")
	}
}

// get one employee from the DB by its employeeid

func GetAllemployees() ([]models.Employee, error) {

	db := middlewares.CreateConnection()

	defer db.Close()

	var employees []models.Employee

	sqlStatement := `SELECT * FROM newEmployee`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic("Unable to execute the query")
	}

	defer rows.Close()

	for rows.Next() {
		var employee models.Employee

		err = rows.Scan(&employee.ID, &employee.Name, &employee.Age)

		if err != nil {
			panic("Unable to scan the row")
		}

		employees = append(employees, employee)

	}

	return employees, err
}

// update employee in the DB

func Updateemployee(id int64, employee models.Employee) int64 {

	db := middlewares.CreateConnection()

	defer db.Close()

	sqlStatement := `UPDATE newEmployee SET name=$2, age=$3 WHERE id=$1`

	res, err := db.Exec(sqlStatement, id, employee.Name, employee.Age)

	if err != nil {
		panic("Unable to execute the query.")
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		panic("Error while checking the affected rows")
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

// delete employee in the DB

func Deleteemployee(id int64) int64 {

	db := middlewares.CreateConnection()

	defer db.Close()

	sqlStatement := `DELETE FROM newEmployee WHERE id=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		panic("Unable to execute the query")
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		panic("Error while checking the affected rows")
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}
