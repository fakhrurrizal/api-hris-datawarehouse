package repository

import (
	"hris-datawarehouse/config"
)

type SalaryData struct {
	Name  string  `json:"name"`
	Total float64 `json:"total"`
}

type SalaryMinMaxResponse struct {
	Highest SalaryMinMax `json:"highest"`
	Lowest  SalaryMinMax `json:"lowest"`
}

type SalaryMinMax struct {
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	EmpID  int     `json:"emp_id,omitempty"`
}



func GetAverageSalaryPerDepartment(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state, gender string,
) (result []SalaryData, err error) {
	query := `
		SELECT d.Department AS name, AVG(e.Salary) AS total
		FROM dim_employee e
		INNER JOIN dim_department d ON e.DeptID = d.DeptID
		WHERE e.DateofTermination IS NULL
		AND e.Salary > 0
	`
	var args []interface{}
	if startDate != "" && endDate != "" {
		query += " AND e.DateofHire BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}
	query += `
		AND (? = 0 OR e.DeptID = ?)
		AND (? = 0 OR e.EmpStatusID = ?)
		AND (? = 0 OR d.ManagerID = ?)
		AND (? = 0 OR e.PositionID = ?)
		AND (? = '' OR e.State = ?)
		AND (? = '' OR e.Gender = ?)
	`
	query += " GROUP BY d.Department, d.DeptID ORDER BY total DESC"
	args = append(args,
		deptID, deptID,
		empStatusID, empStatusID,
		managerID, managerID,
		positionID, positionID,
		state, state,
		gender, gender,
	)
	err = config.DB.Raw(query, args...).Scan(&result).Error
	return
}

func GetAverageSalaryPerPositionWithCount(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state, gender string,
) (result []SalaryData, err error) {
	query := `
		SELECT 
			p.Position AS name, 
			ROUND(AVG(e.Salary), 2) AS total,
			COUNT(e.EmpID) AS count
		FROM dim_employee e
		INNER JOIN dim_position p ON e.PositionID = p.PositionID
		INNER JOIN dim_department d ON e.DeptID = d.DeptID
		WHERE e.DateofTermination IS NULL
		AND e.Salary > 0
	`
	var args []interface{}
	if startDate != "" && endDate != "" {
		query += " AND e.DateofHire BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}
	query += `
		AND (? = 0 OR e.DeptID = ?)
		AND (? = 0 OR e.EmpStatusID = ?)
		AND (? = 0 OR d.ManagerID = ?)
		AND (? = 0 OR e.PositionID = ?)
		AND (? = '' OR e.State = ?)
		AND (? = '' OR e.Gender = ?)
	`
	query += " GROUP BY p.Position, p.PositionID ORDER BY total DESC"
	args = append(args,
		deptID, deptID,
		empStatusID, empStatusID,
		managerID, managerID,
		positionID, positionID,
		state, state,
		gender, gender,
	)
	err = config.DB.Raw(query, args...).Scan(&result).Error
	return
}

func GetHighestSalary(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state, gender string,
) (result SalaryMinMax, err error) {
	query := `
		SELECT 
				e.Employee_Name AS name,
			e.Salary AS salary,
			e.EmpID AS emp_id
		FROM dim_employee e
		INNER JOIN dim_department d ON e.DeptID = d.DeptID
		WHERE e.DateofTermination IS NULL
		AND e.Salary > 0
	`
	var args []interface{}
	if startDate != "" && endDate != "" {
		query += " AND e.DateofHire BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}
	query += `
		AND (? = 0 OR e.DeptID = ?)
		AND (? = 0 OR e.EmpStatusID = ?)
		AND (? = 0 OR d.ManagerID = ?)
		AND (? = 0 OR e.PositionID = ?)
		AND (? = '' OR e.State = ?)
		AND (? = '' OR e.Gender = ?)
	`
	query += " ORDER BY e.Salary DESC LIMIT 1"
	args = append(args,
		deptID, deptID,
		empStatusID, empStatusID,
		managerID, managerID,
		positionID, positionID,
		state, state,
		gender, gender,
	)
	err = config.DB.Raw(query, args...).Scan(&result).Error
	return
}

// Repository Function - Get Lowest Salary
func GetLowestSalary(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state, gender string,
) (result SalaryMinMax, err error) {
	query := `
		SELECT 
			e.Employee_Name AS name,
			e.Salary AS salary,
			e.EmpID AS emp_id
		FROM dim_employee e
		INNER JOIN dim_department d ON e.DeptID = d.DeptID
		WHERE e.DateofTermination IS NULL
		AND e.Salary > 0
	`
	var args []interface{}
	if startDate != "" && endDate != "" {
		query += " AND e.DateofHire BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}
	query += `
		AND (? = 0 OR e.DeptID = ?)
		AND (? = 0 OR e.EmpStatusID = ?)
		AND (? = 0 OR d.ManagerID = ?)
		AND (? = 0 OR e.PositionID = ?)
		AND (? = '' OR e.State = ?)
		AND (? = '' OR e.Gender = ?)
	`
	query += " ORDER BY e.Salary ASC LIMIT 1"
	args = append(args,
		deptID, deptID,
		empStatusID, empStatusID,
		managerID, managerID,
		positionID, positionID,
		state, state,
		gender, gender,
	)
	err = config.DB.Raw(query, args...).Scan(&result).Error
	return
}

func GetHighestAndLowestSalary(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state, gender string,
) (result SalaryMinMaxResponse, err error) {
	highest, err := GetHighestSalary(startDate, endDate, empStatusID, managerID, positionID, deptID, state, gender)
	if err != nil {
		return result, err
	}

	lowest, err := GetLowestSalary(startDate, endDate, empStatusID, managerID, positionID, deptID, state, gender)
	if err != nil {
		return result, err
	}

	result.Highest = highest
	result.Lowest = lowest
	return result, nil
}
