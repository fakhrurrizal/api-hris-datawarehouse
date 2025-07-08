package repository

import (
	"hris-datawarehouse/config"
)

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

func GetEmployeeTerminationRatio(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state, gender string,
) (result []PerformanceScoreCount, err error) {
	query := `
		SELECT 
			CASE 
				WHEN e.TermReason = 'Voluntarily Terminated' THEN 'Voluntarily Terminated'
				WHEN e.TermReason = 'Terminated for Cause' THEN 'Terminated for Cause'
				ELSE 'Other'
			END AS name,
			COUNT(*) AS total
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
	query += `
		GROUP BY 
			CASE 
				WHEN e.TermReason = 'Voluntarily Terminated' THEN 'Voluntarily Terminated'
				WHEN e.TermReason = 'Terminated for Cause' THEN 'Terminated for Cause'
				ELSE 'Other'
			END
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
