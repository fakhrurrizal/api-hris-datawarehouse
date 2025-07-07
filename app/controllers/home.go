package controllers

import (
	"fmt"
	"hris-datawarehouse/app/utils"
	"hris-datawarehouse/config"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	return c.JSON(200, "Welcome "+config.LoadConfig().AppName)
}

type Commit struct {
	ID string `json:"id"`
}

// Version godoc
// @Summary Get Version Build
// @Description Get Version Build
// @Tags Home
// @Accept json
// @Produce json
// @Success 200
// @Router /version [get]
// @Security ApiKeyAuth
func Version(c echo.Context) error {
	commitID, timestamp, err := getLastCommitInfo()
	if err != nil {
		fmt.Println("Error:", err)
		return c.JSON(400, map[string]interface{}{
			"message": "failed get version build",
		})
	}
	formattedTimestamp, err := convertTimestamp(timestamp)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return c.JSON(200, map[string]interface{}{
		"data": map[string]interface{}{
			"build_id":   commitID,
			"last_build": timestamp,
			"version":    formattedTimestamp,
		},
		"message": "",
	})

}

func convertTimestamp(timestamp string) (string, error) {
	t, err := time.ParseInLocation("Mon Jan 2 15:04:05 2006", timestamp, utils.GetTimeLocation())
	if err != nil {
		return "", err
	}
	return t.Format("02-01-2006 15:04:05"), nil
}

func getLastCommitInfo() (string, string, error) {
	cmd := exec.Command("git", "log", "-1", "--pretty=format:%h %cd", "--date=local")
	output, err := cmd.Output()
	if err != nil {
		return "", "", err
	}
	outputStr := strings.TrimSpace(string(output))
	fields := strings.Fields(outputStr)
	if len(fields) <= 2 {
		return "", "", fmt.Errorf("unexpected output format: %q", outputStr)
	}
	return fields[0], strings.Join(fields[1:], " "), nil
}

// Test godoc
func Test(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	data := ""
	return c.JSON(404, map[string]interface{}{
		"data": map[string]interface{}{
			"detail": data,
			"id":     id,
		},
		"message": "",
	})
}
