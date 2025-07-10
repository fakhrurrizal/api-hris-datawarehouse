package repository

import (
	"hris-datawarehouse/config"
)

type TerminationRatio struct {
	Label string `json:"label"`
	Total int    `json:"total"`
}

type EmpRatio struct {
	Name  string `json:"name"`
	Total int    `json:"total"`
}

func GetEmployeeTerminationByReason(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state, gender string,
) (result []PerformanceScoreCount, err error) {

	query := `
		SELECT 
			CASE 
				WHEN f.TermReason IS NULL OR TRIM(f.TermReason) = '' THEN 'Other'
				ELSE f.TermReason
			END AS name,
			COUNT(*) AS total
		FROM fact_employment f
	`

	if state != "" || gender != "" || managerID != 0 {
		query += " INNER JOIN dim_employee e ON f.EmpID = e.EmpID"
	}

	query += " WHERE f.Is_Terminated = 'Yes'"

	var args []interface{}

	if startDate != "" && endDate != "" {
		query += " AND f.DateofTermination BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}

	if deptID != 0 {
		query += " AND f.DeptID = ?"
		args = append(args, deptID)
	}

	if empStatusID != 0 {
		query += " AND f.EmpStatusID = ?"
		args = append(args, empStatusID)
	}

	if positionID != 0 {
		query += " AND f.PositionID = ?"
		args = append(args, positionID)
	}

	if managerID != 0 {
		query += " AND f.ManagerID = ?"
		args = append(args, managerID)
	}

	if state != "" {
		query += " AND e.State = ?"
		args = append(args, state)
	}

	if gender != "" {
		query += " AND e.Gender = ?"
		args = append(args, gender)
	}

	query += `
		GROUP BY name
		ORDER BY total DESC
	`

	err = config.DB.Raw(query, args...).Scan(&result).Error
	return
}

func GetEmployeeTerminationByDepartment(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state, gender string,
) (result []PerformanceScoreCount, err error) {
	query := `
		SELECT d.Department AS name, COUNT(*) AS total
		FROM fact_employment f
		INNER JOIN dim_employee e ON f.EmpID = e.EmpID
		INNER JOIN dim_department d ON f.DeptID = d.DeptID
		WHERE f.Is_Terminated = 'Yes'
	`
	var args []interface{}

	if startDate != "" && endDate != "" {
		query += " AND f.DateofTermination BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}

	query += `
		AND (? = 0 OR f.DeptID = ?)
		AND (? = 0 OR f.EmpStatusID = ?)
		AND (? = 0 OR f.PositionID = ?)
		AND (? = '' OR e.State = ?)
		AND (? = '' OR e.Gender = ?)
	`

	query += " GROUP BY d.Department, f.DeptID ORDER BY total DESC"

	args = append(args,
		deptID, deptID,
		empStatusID, empStatusID,
		positionID, positionID,
		state, state,
		gender, gender,
	)

	err = config.DB.Raw(query, args...).Scan(&result).Error
	return
}

func GetEmployeeTerminationRatio(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state, gender string,
) (result []EmpRatio, err error) {
	query := `
		SELECT 
			sub.termination_label AS name, 
			COUNT(*) AS total
		FROM (
			SELECT 
				CASE 
					WHEN LOWER(f.TermReason) LIKE '%career change%'
						OR LOWER(f.TermReason) LIKE '%relocation%'
						OR LOWER(f.TermReason) LIKE '%return to school%'
						OR LOWER(f.TermReason) LIKE '%more money%'
						OR LOWER(f.TermReason) LIKE '%unhappy%'
						OR LOWER(f.TermReason) LIKE '%maternity%'
						OR LOWER(f.TermReason) LIKE '%retiring%'
					THEN 'Voluntarily Terminated'

					WHEN LOWER(f.TermReason) LIKE '%gross misconduct%'
						OR LOWER(f.TermReason) LIKE '%no-call%'
						OR LOWER(f.TermReason) LIKE '%performance%'
						OR LOWER(f.TermReason) LIKE '%attendance%'
					THEN 'Terminated for Cause'

					ELSE 'Other'
				END AS termination_label
			FROM fact_employment f
			INNER JOIN dim_employee e ON f.EmpID = e.EmpID
			WHERE f.Is_Terminated = 'Yes'
			AND f.TermReason IS NOT NULL
			AND TRIM(f.TermReason) != ''
	`
	var args []interface{}

	if startDate != "" && endDate != "" {
		query += " AND f.DateofTermination BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}

	query += `
			AND (? = 0 OR f.DeptID = ?)
			AND (? = 0 OR f.EmpStatusID = ?)
			AND (? = 0 OR f.PositionID = ?)
			AND (? = '' OR e.State = ?)
			AND (? = '' OR e.Gender = ?)
		) AS sub
		GROUP BY sub.termination_label
		ORDER BY total DESC
	`

	args = append(args,
		deptID, deptID,
		empStatusID, empStatusID,
		positionID, positionID,
		state, state,
		gender, gender,
	)

	err = config.DB.Raw(query, args...).Scan(&result).Error
	return
}
