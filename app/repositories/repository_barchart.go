package repository

import "hris-datawarehouse/config"

type GlobalCount struct {
	Name  string `json:"name"`
	Total int    `json:"total"`
}

func GetEmployeeCountPerDepartment(
	startDate, endDate string,
	empStatusID, managerID, positionID int,
	state, gender string,
) (result []GlobalCount, err error) {
	query := `
		SELECT d.Department AS name, COUNT(e.EmpID) AS total
		FROM dim_employee e
		INNER JOIN dim_department d ON e.DeptID = d.DeptID
		WHERE e.DateofTermination IS NULL
	`

	var args []interface{}

	if startDate != "" && endDate != "" {
		query += " AND e.DateofHire BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}

	// Add filter conditions directly
	query += `
		AND (? = 0 OR e.DeptID = ?)
		AND (? = '' OR e.Gender = ?)
		AND (? = 0 OR e.EmpStatusID = ?)
		AND (? = 0 OR d.ManagerID = ?)
		AND (? = 0 OR e.PositionID = ?)
		AND (? = '' OR e.State = ?)
	`

	query += " GROUP BY e.DeptID, d.Department ORDER BY total DESC"

	args = append(args,
		0, 0,
		gender, gender,
		empStatusID, empStatusID,
		managerID, managerID,
		positionID, positionID,
		state, state,
	)

	err = config.DB.Raw(query, args...).Scan(&result).Error
	return
}

func GetEmployeeCountPerGender(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state string,
) (result []GlobalCount, err error) {
	query := `
		SELECT e.Gender AS name, COUNT(e.EmpID) AS total
		FROM dim_employee e
		INNER JOIN dim_department d ON e.DeptID = d.DeptID
		WHERE e.DateofTermination IS NULL
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
	`

	query += " GROUP BY e.Gender ORDER BY total DESC"

	args = append(args,
		deptID, deptID,
		empStatusID, empStatusID,
		managerID, managerID,
		positionID, positionID,
		state, state,
	)

	err = config.DB.Raw(query, args...).Scan(&result).Error
	return
}


func GetEmployeeCountPerRecruitmentSource(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state, gender string,
) (result []GlobalCount, err error) {
	query := `
		SELECT e.RecruitmentSource AS name, COUNT(e.EmpID) AS total
		FROM dim_employee e
		INNER JOIN dim_department d ON e.DeptID = d.DeptID
		WHERE e.DateofTermination IS NULL
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

	query += " GROUP BY e.RecruitmentSource ORDER BY total DESC"

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
func GetEmployeeCountPerCitizenDesc(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state, gender string,
) (result []GlobalCount, err error) {
	query := `
		SELECT e.CitizenDesc AS name, COUNT(e.EmpID) AS total
		FROM dim_employee e
		INNER JOIN dim_department d ON e.DeptID = d.DeptID
		WHERE e.DateofTermination IS NULL
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

	query += " GROUP BY e.CitizenDesc ORDER BY total DESC"

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

func GetEmployeeCountPerRaceDesc(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state, gender string,
) (result []GlobalCount, err error) {
	query := `
		SELECT e.RaceDesc AS name, COUNT(e.EmpID) AS total
		FROM dim_employee e
		INNER JOIN dim_department d ON e.DeptID = d.DeptID
		WHERE e.DateofTermination IS NULL
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
	query += " GROUP BY e.RaceDesc ORDER BY total DESC"
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