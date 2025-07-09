package repository

import (
	"hris-datawarehouse/config"
)

type TerminationRatio struct {
	Label string `json:"label"`
	Total int    `json:"total"`
}

func GetEmployeeTerminationByReason(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state, gender string,
) (result []PerformanceScoreCount, err error) {
	query := `
		SELECT e.TermReason AS name, COUNT(*) AS total
		FROM dim_employee e
		WHERE e.DateofTermination IS NOT NULL
		AND e.TermReason IS NOT NULL
		AND e.TermReason != ''
	`
	var args []interface{}

	if startDate != "" && endDate != "" {
		query += " AND e.DateofTermination BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}

	query += `
		AND (? = 0 OR e.DeptID = ?)
		AND (? = 0 OR e.EmpStatusID = ?)
		AND (? = 0 OR e.PositionID = ?)
		AND (? = '' OR e.State = ?)
		AND (? = '' OR e.Gender = ?)
	`
	query += " GROUP BY e.TermReason ORDER BY total DESC"

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

// 2. Jumlah karyawan keluar per Department
func GetEmployeeTerminationByDepartment(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state, gender string,
) (result []PerformanceScoreCount, err error) {
	query := `
		SELECT dept.Department AS name, COUNT(*) AS total
		FROM dim_employee e
		INNER JOIN dim_department dept ON e.DeptID = dept.DeptID
		WHERE e.DateofTermination IS NOT NULL
	`
	var args []interface{}

	if startDate != "" && endDate != "" {
		query += " AND e.DateofTermination BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}

	query += `
		AND (? = 0 OR e.DeptID = ?)
		AND (? = 0 OR e.EmpStatusID = ?)
		AND (? = 0 OR e.PositionID = ?)
		AND (? = '' OR e.State = ?)
		AND (? = '' OR e.Gender = ?)
	`
	query += " GROUP BY dept.Department, dept.DeptID ORDER BY total DESC"

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

type EmpRatio struct {
	Name  string `json:"name"`
	Total int    `json:"total"`
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
					WHEN LOWER(e.TermReason) LIKE '%career change%'
						OR LOWER(e.TermReason) LIKE '%relocation%'
						OR LOWER(e.TermReason) LIKE '%return to school%'
						OR LOWER(e.TermReason) LIKE '%more money%'
						OR LOWER(e.TermReason) LIKE '%unhappy%'
						OR LOWER(e.TermReason) LIKE '%maternity%'
						OR LOWER(e.TermReason) LIKE '%retiring%'
					THEN 'Voluntarily Terminated'

					WHEN LOWER(e.TermReason) LIKE '%gross misconduct%'
						OR LOWER(e.TermReason) LIKE '%no-call%'
						OR LOWER(e.TermReason) LIKE '%performance%'
						OR LOWER(e.TermReason) LIKE '%attendance%'
					THEN 'Terminated for Cause'

					ELSE 'Other'
				END AS termination_label
			FROM dim_employee e
			WHERE e.DateofTermination IS NOT NULL
				AND e.TermReason IS NOT NULL
				AND TRIM(e.TermReason) != ''
	`

	var args []interface{}

	if startDate != "" && endDate != "" {
		query += " AND e.DateofTermination BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}

	query += `
			AND (? = 0 OR e.DeptID = ?)
			AND (? = 0 OR e.EmpStatusID = ?)
			AND (? = 0 OR e.PositionID = ?)
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
