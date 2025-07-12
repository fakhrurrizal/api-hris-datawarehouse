package repository

import (
	"fmt"
	"hris-datawarehouse/app/models"
	"hris-datawarehouse/app/reqres"
	"hris-datawarehouse/app/utils"
	"hris-datawarehouse/config"
	"time"

	"github.com/guregu/null"
)

type TotalResult struct {
	Total       int64
	LastUpdated time.Time
}

func GetDepartments(param reqres.ReqPaging) (data reqres.ResPaging, err error) {
	var responses []models.DimDepartment
	where := "1=1"
	var args []interface{}

	if param.Search != "" {
		where += " AND Department ILIKE ?"
		args = append(args, "%"+param.Search+"%")
	}

	var total int64
	countQuery := "SELECT COUNT(*) FROM dim_department WHERE " + where
	err = config.DB.Raw(countQuery, args...).Scan(&total).Error
	if err != nil {
		return
	}

	validSortFields := map[string]bool{
		"DeptID":     true,
		"Department": true,
	}

	if !validSortFields[param.Sort] {
		param.Sort = "DeptID"
	}

	if param.Order != "ASC" && param.Order != "DESC" {
		param.Order = "ASC"
	}

	param.Offset = (param.Page - 1) * param.Limit

	dataQuery := fmt.Sprintf("SELECT * FROM dim_department WHERE %s ORDER BY %s %s LIMIT ? OFFSET ?",
		where, param.Sort, param.Order)
	argsWithLimit := append(args, param.Limit, param.Offset)

	err = config.DB.Raw(dataQuery, argsWithLimit...).Scan(&responses).Error
	if err != nil {
		return
	}

	data = utils.PopulateResPaging(&param, responses, total, total, null.Time{})
	return
}

func GetEmployees(param reqres.ReqPaging) (data reqres.ResPaging, err error) {
	var responses []models.DimEmployee
	where := "1=1"
	var args []interface{}

	if param.Search != "" {
		where += " AND Employee_Name ILIKE ?"
		args = append(args, "%"+param.Search+"%")
	}

	var total int64
	err = config.DB.Raw("SELECT COUNT(*) FROM dim_employee WHERE "+where, args...).Scan(&total).Error
	if err != nil {
		return
	}

	utils.ValidateSort(&param, map[string]bool{
		"EmpID":         true,
		"Employee_Name": true,
	}, "EmpID")

	dataQuery := fmt.Sprintf("SELECT * FROM dim_employee WHERE %s ORDER BY %s %s LIMIT ? OFFSET ?", where, param.Sort, param.Order)
	args = append(args, param.Limit, param.Offset)

	err = config.DB.Raw(dataQuery, args...).Scan(&responses).Error
	if err != nil {
		return
	}

	data = utils.PopulateResPaging(&param, responses, total, total, null.Time{})
	return
}

func GetEmploymentStatuses(param reqres.ReqPaging) (data reqres.ResPaging, err error) {
	var responses []models.DimEmploymentStatus
	where := "1=1"
	var args []interface{}

	if param.Search != "" {
		where += " AND EmploymentStatus ILIKE ?"
		args = append(args, "%"+param.Search+"%")
	}

	var total int64
	err = config.DB.Raw("SELECT COUNT(*) FROM dim_employment_status WHERE "+where, args...).Scan(&total).Error
	if err != nil {
		return
	}

	utils.ValidateSort(&param, map[string]bool{
		"EmpStatusID":      true,
		"EmploymentStatus": true,
	}, "EmpStatusID")

	dataQuery := fmt.Sprintf("SELECT * FROM dim_employment_status WHERE %s ORDER BY %s %s LIMIT ? OFFSET ?", where, param.Sort, param.Order)
	args = append(args, param.Limit, param.Offset)

	err = config.DB.Raw(dataQuery, args...).Scan(&responses).Error
	if err != nil {
		return
	}

	data = utils.PopulateResPaging(&param, responses, total, total, null.Time{})
	return
}

func GetManagers(param reqres.ReqPaging) (data reqres.ResPaging, err error) {
	var responses []models.DimManager
	where := "1=1"
	var args []interface{}

	if param.Search != "" {
		where += " AND ManagerName ILIKE ?"
		args = append(args, "%"+param.Search+"%")
	}

	var total int64
	err = config.DB.Raw("SELECT COUNT(*) FROM dim_manager WHERE "+where, args...).Scan(&total).Error
	if err != nil {
		return
	}

	utils.ValidateSort(&param, map[string]bool{
		"ManagerID":   true,
		"ManagerName": true,
	}, "ManagerID")

	dataQuery := fmt.Sprintf("SELECT * FROM dim_manager WHERE %s ORDER BY %s %s LIMIT ? OFFSET ?", where, param.Sort, param.Order)
	args = append(args, param.Limit, param.Offset)

	err = config.DB.Raw(dataQuery, args...).Scan(&responses).Error
	if err != nil {
		return
	}

	data = utils.PopulateResPaging(&param, responses, total, total, null.Time{})
	return
}

func GetMaritalStatuses(param reqres.ReqPaging) (data reqres.ResPaging, err error) {
	var responses []models.DimMaritalStatus
	where := "1=1"
	var args []interface{}

	if param.Search != "" {
		where += " AND MaritalDesc ILIKE ?"
		args = append(args, "%"+param.Search+"%")
	}

	var total int64
	err = config.DB.Raw("SELECT COUNT(*) FROM dim_marital_status WHERE "+where, args...).Scan(&total).Error
	if err != nil {
		return
	}

	utils.ValidateSort(&param, map[string]bool{
		"MaritalStatusID": true,
		"MaritalDesc":     true,
	}, "MaritalStatusID")

	dataQuery := fmt.Sprintf("SELECT * FROM dim_marital_status WHERE %s ORDER BY %s %s LIMIT ? OFFSET ?", where, param.Sort, param.Order)
	args = append(args, param.Limit, param.Offset)

	err = config.DB.Raw(dataQuery, args...).Scan(&responses).Error
	if err != nil {
		return
	}

	data = utils.PopulateResPaging(&param, responses, total, total, null.Time{})
	return
}

func GetPerformances(param reqres.ReqPaging) (data reqres.ResPaging, err error) {
	var responses []models.DimPerformance
	where := "1=1"
	var args []interface{}

	if param.Search != "" {
		where += " AND PerformanceScore ILIKE ?"
		args = append(args, "%"+param.Search+"%")
	}

	var total int64
	err = config.DB.Raw("SELECT COUNT(*) FROM dim_performance WHERE "+where, args...).Scan(&total).Error
	if err != nil {
		return
	}

	utils.ValidateSort(&param, map[string]bool{
		"PerfScoreID":      true,
		"PerformanceScore": true,
	}, "PerfScoreID")

	dataQuery := fmt.Sprintf("SELECT * FROM dim_performance WHERE %s ORDER BY %s %s LIMIT ? OFFSET ?", where, param.Sort, param.Order)
	args = append(args, param.Limit, param.Offset)

	err = config.DB.Raw(dataQuery, args...).Scan(&responses).Error
	if err != nil {
		return
	}

	data = utils.PopulateResPaging(&param, responses, total, total, null.Time{})
	return
}

func GetPositions(param reqres.ReqPaging) (data reqres.ResPaging, err error) {
	var responses []models.DimPosition
	where := "1=1"
	var args []interface{}

	if param.Search != "" {
		where += " AND Position ILIKE ?"
		args = append(args, "%"+param.Search+"%")
	}

	var total int64
	err = config.DB.Raw("SELECT COUNT(*) FROM dim_position WHERE "+where, args...).Scan(&total).Error
	if err != nil {
		return
	}

	utils.ValidateSort(&param, map[string]bool{
		"PositionID": true,
		"Position":   true,
	}, "PositionID")

	dataQuery := fmt.Sprintf("SELECT * FROM dim_position WHERE %s ORDER BY %s %s LIMIT ? OFFSET ?", where, param.Sort, param.Order)
	args = append(args, param.Limit, param.Offset)

	err = config.DB.Raw(dataQuery, args...).Scan(&responses).Error
	if err != nil {
		return
	}

	data = utils.PopulateResPaging(&param, responses, total, total, null.Time{})
	return
}

func GetEmploymentWithFilters(param reqres.ReqPaging) (data reqres.ResPaging, err error) {
	var responses []models.EmploymentResponse
	where := "f.Is_Terminated = 'No'"
	var args []interface{}

	if param.Search != "" {
		where += " AND e.Employee_Name ILIKE ?"
		args = append(args, "%"+param.Search+"%")
	}

	var total int64
	countQuery := fmt.Sprintf(`
		SELECT COUNT(*)
		FROM fact_employment f
		JOIN dim_employee e ON f.EmpID = e.EmpID
		JOIN dim_department d ON f.DeptID = d.DeptID
		JOIN dim_position p ON f.PositionID = p.PositionID
		LEFT JOIN dim_manager m ON f.ManagerID = m.ManagerID
		WHERE %s
	`, where)

	err = config.DB.Raw(countQuery, args...).Scan(&total).Error
	if err != nil {
		return
	}

	utils.ValidateSort(&param, map[string]bool{
		"e.Employee_Name": true,
		"p.Position":      true,
		"d.Department":    true,
		"f.DateofHire":    true,
		"f.Salary":        true,
	}, "e.Employee_Name")

	dataQuery := fmt.Sprintf(`
		SELECT 
			f.EmpID,
			e.Employee_Name AS employee_name,
			p.Position,
			d.Department,
			COALESCE(m.ManagerName, '-') AS ManagerName,
			DATE_FORMAT(f.DateofHire, '%%Y-%%m-%%d') AS DateOfHire,
			DATE_FORMAT(f.DateofTermination, '%%Y-%%m-%%d') AS DateOfTermination,
			f.TermReason,
			f.Salary,
			e.Gender,
			e.State,
			e.Zip,
			e.CitizenDesc,
			e.HispanicLatino,
			e.RaceDesc,
			f.Tenure_Days,
			f.DaysLateLast30,
			f.Absences,
			f.RecruitmentSource
		FROM fact_employment f
		JOIN dim_employee e ON f.EmpID = e.EmpID
		JOIN dim_department d ON f.DeptID = d.DeptID
		JOIN dim_position p ON f.PositionID = p.PositionID
		LEFT JOIN dim_manager m ON f.ManagerID = m.ManagerID
		WHERE %s
		ORDER BY %s %s
		LIMIT ? OFFSET ?
	`, where, param.Sort, param.Order)

	args = append(args, param.Limit, param.Offset)

	err = config.DB.Raw(dataQuery, args...).Scan(&responses).Error
	if err != nil {
		return
	}

	data = utils.PopulateResPaging(&param, responses, total, total, null.Time{})
	return
}
