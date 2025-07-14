package models

import "time"

type DimDepartment struct {
	DeptID     int    `json:"dept_id" db:"DeptID"`
	Department string `json:"department" db:"Department"`
	ManagerID  int    `json:"manager_id" db:"ManagerID"`
}

type DimEmployee struct {
	EmpID                     int        `json:"emp_id" db:"EmpID"`
	EmployeeName              string     `json:"employee_name" db:"Employee_Name"`
	DateofBirth               time.Time  `json:"date_of_birth" db:"DateofBirth"`
	Gender                    string     `json:"gender" db:"Gender"`
	CitizenDesc               string     `json:"citizen_desc" db:"CitizenDesc"`
	HispanicLatino            string     `json:"hispanic_latino" db:"HispanicLatino"`
	RaceDesc                  string     `json:"race_desc" db:"RaceDesc"`
	MaritalStatusID           int        `json:"marital_status_id" db:"MaritalStatusID"`
	EmpStatusID               int        `json:"emp_status_id" db:"EmpStatusID"`
	DeptID                    int        `json:"dept_id" db:"DeptID"`
	PositionID                int        `json:"position_id" db:"PositionID"`
	Salary                    int        `json:"salary" db:"Salary"`
	DateofTermination         *time.Time `json:"date_of_termination" db:"DateofTermination"`
	TermReason                string     `json:"term_reason" db:"TermReason"`
	RecruitmentSource         string     `json:"recruitment_source" db:"RecruitmentSource"`
	PerfScoreID               int        `json:"perf_score_id" db:"PerfScoreID"`
	DateofHire                time.Time  `json:"date_of_hire" db:"DateofHire"`
	FromDiversityJobFairID    int        `json:"from_diversity_job_fair_id" db:"FromDiversityJobFairID"`
	DaysLateLast30            int        `json:"days_late_last_30" db:"DaysLateLast30"`
	State                     string     `json:"state" db:"State"`
	Zip                       string     `json:"zip" db:"Zip"`
	EmpSatisfaction           int        `json:"emp_satisfaction" db:"EmpSatisfaction"`
	SpecialProjectsCount      int        `json:"special_projects_count" db:"SpecialProjectsCount"`
	LastPerformanceReviewDate *time.Time `json:"last_performance_review_date" db:"LastPerformanceReview_Date"`
	Absences                  int        `json:"absences" db:"Absences"`
}

type DimEmploymentStatus struct {
	EmpStatusID      int    `json:"emp_status_id" db:"EmpStatusID"`
	EmploymentStatus string `json:"employment_status" db:"EmploymentStatus"`
}

type DimManager struct {
	ManagerID   int    `json:"manager_id" db:"ManagerID"`
	ManagerName string `json:"manager_name" db:"ManagerName"`
}

type DimMaritalStatus struct {
	MaritalStatusID int    `json:"marital_status_id" db:"MaritalStatusID"`
	MarriedID       int    `json:"married_id" db:"MarriedID"`
	MaritalDesc     string `json:"marital_desc" db:"MaritalDesc"`
}

type DimPerformance struct {
	PerfScoreID      int    `json:"perf_score_id" db:"PerfScoreID"`
	PerformanceScore string `json:"performance_score" db:"PerformanceScore"`
}

type DimPosition struct {
	PositionID int    `json:"position_id" db:"PositionID"`
	Position   string `json:"position" db:"Position"`
}

type EmploymentResponse struct {
	EmpID              int     `json:"emp_id"`
	EmployeeName       string  `json:"employee_name"`
	Position           string  `json:"position"`
	Department         string  `json:"department"`
	ManagerName        string  `json:"manager_name"`
	DateOfHire         string  `json:"date_of_hire"`
	DateOfTermination  *string `json:"date_of_termination,omitempty"`
	TermReason         *string `json:"term_reason,omitempty"`
	Salary             float64 `json:"salary"`
	Gender             string  `json:"gender"`
	State              string  `json:"state"`
	Zip                string  `json:"zip"`
	CitizenDesc        string  `json:"citizen_desc"`
	HispanicLatino     string  `json:"hispanic_latino"`
	RaceDesc           string  `json:"race_desc"`
	TenureDays         int     `json:"tenure_days"`
	DaysLateLast30     int     `json:"days_late_last_30"`
	Absences           int     `json:"absences"`
	RecruitmentSource  string  `json:"recruitment_source"`
	MaritalDesc       string  `json:"marital_desc"`
}
