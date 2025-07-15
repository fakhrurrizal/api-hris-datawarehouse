package repository

import "hris-datawarehouse/config"

type GlobalCount struct {
	Name  string `json:"name"`
	Total int    `json:"total"`
}

func buildFilterQueryForGroup(
	empStatusID, managerID, positionID, deptID int,
	state, gender string,
) (string, []interface{}) {
	var filter string
	var args []interface{}

	if deptID != 0 {
		filter += " AND f.DeptID = ?"
		args = append(args, deptID)
	}
	if empStatusID != 0 {
		filter += " AND f.EmpStatusID = ?"
		args = append(args, empStatusID)
	}
	if managerID != 0 {
		filter += " AND f.ManagerID = ?"
		args = append(args, managerID)
	}
	if positionID != 0 {
		filter += " AND f.PositionID = ?"
		args = append(args, positionID)
	}
	if state != "" {
		filter += " AND e.State = ?"
		args = append(args, state)
	}
	if gender != "" {
		filter += " AND e.Gender = ?"
		args = append(args, gender)
	}
	return filter, args
}

func GetEmployeeCountPerDepartment(
	startDate, endDate string,
	empStatusID, managerID, positionID int,
	state, gender string,
) (result []GlobalCount, err error) {
	query := `
		SELECT d.Department AS name, COUNT(DISTINCT f.EmpID) AS total
		FROM fact_employment f
		JOIN dim_employee e USING(EmpID)
		JOIN dim_department d USING(DeptID)
		WHERE f.DateofTermination IS NULL
	`

	var args []interface{}
	if startDate != "" && endDate != "" {
		query += " AND f.DateofHire BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}

	filter, filterArgs := buildFilterQueryForGroup(empStatusID, managerID, positionID, 0, state, gender)
	query += filter

	query += " GROUP BY f.DeptID, d.Department ORDER BY total DESC"
	args = append(args, filterArgs...)

	err = config.DB.Raw(query, args...).Scan(&result).Error
	return
}

func GetEmployeeCountPerGender(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state string,
) (result []GlobalCount, err error) {
	query := `
		SELECT e.Gender AS name, COUNT(DISTINCT f.EmpID) AS total
		FROM fact_employment f
		JOIN dim_employee e USING(EmpID)
		JOIN dim_department d USING(DeptID)
		WHERE f.DateofTermination IS NULL
	`

	var args []interface{}
	if startDate != "" && endDate != "" {
		query += " AND f.DateofHire BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}

	filter, filterArgs := buildFilterQueryForGroup(empStatusID, managerID, positionID, deptID, state, "")
	query += filter

	query += " GROUP BY e.Gender ORDER BY total DESC"
	args = append(args, filterArgs...)

	err = config.DB.Raw(query, args...).Scan(&result).Error
	return
}

func GetEmployeeCountPerRecruitmentSource(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state, gender string,
) (result []GlobalCount, err error) {
	query := `
		SELECT f.RecruitmentSource AS name, COUNT(DISTINCT f.EmpID) AS total
		FROM fact_employment f
		JOIN dim_employee e USING(EmpID)
		JOIN dim_department d USING(DeptID)
		WHERE f.Is_Terminated = 0
	`

	var args []interface{}
	if startDate != "" && endDate != "" {
		query += " AND f.DateofHire BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}

	filter, filterArgs := buildFilterQueryForGroup(empStatusID, managerID, positionID, deptID, state, gender)
	query += filter

	query += " GROUP BY f.RecruitmentSource ORDER BY total DESC"
	args = append(args, filterArgs...)

	err = config.DB.Raw(query, args...).Scan(&result).Error
	return
}

func GetEmployeeCountPerCitizenDesc(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state, gender string,
) (result []GlobalCount, err error) {
	query := `
		SELECT e.CitizenDesc AS name, COUNT(DISTINCT f.EmpID) AS total
		FROM fact_employment f
		JOIN dim_employee e USING(EmpID)
		JOIN dim_department d USING(DeptID)
		WHERE f.DateofTermination IS NULL
	`

	var args []interface{}
	if startDate != "" && endDate != "" {
		query += " AND f.DateofHire BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}

	filter, filterArgs := buildFilterQueryForGroup(empStatusID, managerID, positionID, deptID, state, gender)
	query += filter

	query += " GROUP BY e.CitizenDesc ORDER BY total DESC"
	args = append(args, filterArgs...)

	err = config.DB.Raw(query, args...).Scan(&result).Error
	return
}
func GetEmployeeCountPerRaceDesc(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state, gender string,
) (result []GlobalCount, err error) {
	query := `
		SELECT e.RaceDesc AS name, COUNT(DISTINCT f.EmpID) AS total
		FROM fact_employment f
		JOIN dim_employee e USING(EmpID)
		JOIN dim_department d USING(DeptID)
		WHERE f.DateofTermination IS NULL
	`

	var args []interface{}
	if startDate != "" && endDate != "" {
		query += " AND f.DateofHire BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}

	filter, filterArgs := buildFilterQueryForGroup(empStatusID, managerID, positionID, deptID, state, gender)
	query += filter

	query += " GROUP BY e.RaceDesc ORDER BY total DESC"
	args = append(args, filterArgs...)

	err = config.DB.Raw(query, args...).Scan(&result).Error
	return
}

func GetEmployeeMaritalRatio(
	deptID, empStatusID, managerID, positionID int,
	state, gender, startDate, endDate string,
) (result []EmpRatio, err error) {
	query := `
		SELECT ms.MaritalDesc AS name, COUNT(*) AS total
		FROM fact_employment f
		JOIN dim_employee e ON f.EmpID = e.EmpID
		JOIN dim_marital_status ms ON e.MaritalStatusID = ms.MaritalStatusID
		WHERE f.DateofTermination IS NULL
	`

	var args []interface{}

	if startDate != "" && endDate != "" {
		query += ` AND f.DateofHire BETWEEN ? AND ?`
		args = append(args, startDate, endDate)
	}

	query += `
		AND (? = 0 OR f.DeptID = ?)
		AND (? = 0 OR f.EmpStatusID = ?)
		AND (? = 0 OR f.ManagerID = ?)
		AND (? = 0 OR f.PositionID = ?)
		AND (? = '' OR e.State = ?)
		AND (? = '' OR e.Gender = ?)
	`

	args = append(args,
		deptID, deptID,
		empStatusID, empStatusID,
		managerID, managerID,
		positionID, positionID,
		state, state,
		gender, gender,
	)

	query += `
		GROUP BY ms.MaritalDesc
		ORDER BY total DESC
	`

	err = config.DB.Raw(query, args...).Scan(&result).Error
	return
}

func GetEmployeeAgeRatio(
	deptID, empStatusID, managerID, positionID int,
	state, gender, startDate, endDate string,
) (result []EmpRatio, err error) {
	query := `
		SELECT 
			CASE
				WHEN age < 25 THEN '< 25'
				WHEN age BETWEEN 25 AND 34 THEN '25-34'
				WHEN age BETWEEN 35 AND 44 THEN '35-44'
				WHEN age BETWEEN 45 AND 54 THEN '45-54'
				ELSE '55+'
			END AS name,
			COUNT(*) AS total
		FROM (
			SELECT 
				TIMESTAMPDIFF(YEAR, e.DateofBirth, CURDATE()) AS age
			FROM fact_employment f
			JOIN dim_employee e ON f.EmpID = e.EmpID
			WHERE f.DateofTermination IS NULL
	`

	var args []interface{}

	if startDate != "" && endDate != "" {
		query += ` AND f.DateofHire BETWEEN ? AND ?`
		args = append(args, startDate, endDate)
	}

	query += `
		AND (? = 0 OR f.DeptID = ?)
		AND (? = 0 OR f.EmpStatusID = ?)
		AND (? = 0 OR f.ManagerID = ?)
		AND (? = 0 OR f.PositionID = ?)
		AND (? = '' OR e.State = ?)
		AND (? = '' OR e.Gender = ?)
		) sub
		GROUP BY name
		ORDER BY name
	`

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
