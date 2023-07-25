package main

import (
	"context"
	"html/template"
	"math"
	"net/http"
	"strconv"
	"task13/connection"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Id           int
	Title        string
	Content      string
	Duration     string
	StartDate    time.Time
	EndDate      time.Time
	Technologies []string
	NodeJs       bool
	ReactJs      bool
	VueJs        bool
	JavaScript   bool
	Image        string
}

// var dataProjects = []Project{}

func main() {
	e := echo.New()

	connection.DatabaseConnection()

	e.Static("/assets", "assets")

	e.GET("/", home)
	e.GET("/addProject", addProject)
	e.GET("/contact", contact)
	e.GET("/testimonial", testimonial)
	e.GET("/projectDetail/:id", projectDetail)
	e.GET("/formEditProject/:id", formEditProject)

	e.POST("/addDataProject", addDataProject)
	e.POST("/deleteProject/:id", deleteProject)
	e.POST("/editProject", editProject)

	// e.POST("/addDataContact", addDataContact)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	tmpl, err := template.ParseFiles("index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	dataProjects, errBlogs := connection.Conn.Query(context.Background(), "SELECT id, name, start_date, end_date, description, technologies, image FROM tb_projects")

	if errBlogs != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var resultProjects []Project
	for dataProjects.Next() {
		var each = Project{}

		err := dataProjects.Scan(&each.Id, &each.Title, &each.StartDate, &each.EndDate, &each.Content, &each.Technologies, &each.Image)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		each.Duration = countDuration(each.EndDate, each.StartDate)

		if checkValue(each.Technologies, "ReactJs") {
			each.ReactJs = true
		}
		if checkValue(each.Technologies, "JavaScript") {
			each.JavaScript = true
		}
		if checkValue(each.Technologies, "VueJs") {
			each.VueJs = true
		}
		if checkValue(each.Technologies, "NodeJs") {
			each.NodeJs = true
		}

		resultProjects = append(resultProjects, each)
	}

	data := map[string]interface{}{
		"Projects": resultProjects,
	}

	return tmpl.Execute(c.Response(), data)
}

func countDuration(endDate time.Time, startDate time.Time) string {

	duration := endDate.Sub(startDate)
	days := (int(math.Floor(duration.Hours() / 24)))
	month := days / 30

	var result string

	if days > 30 {
		result = strconv.Itoa(month) + " month"
		if month > 12 {
			result = strconv.Itoa(month/12) + " year"
		}
		return result
	}

	return strconv.Itoa(days) + " day"
}

func checkValue(tech []string, object string) bool {
	for _, dataTech := range tech {
		if dataTech == object {
			return true
		}
	}
	return false
}

func projectDetail(c echo.Context) error {
	id := c.Param("id")

	tmpl, err := template.ParseFiles("html/project-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	idToInt, _ := strconv.Atoi(id)

	projectDetail := Project{}

	errQuery := connection.Conn.QueryRow(context.Background(), "SELECT * FROM public.tb_projects WHERE id=$1", idToInt).Scan(&projectDetail.Id, &projectDetail.Title, &projectDetail.StartDate, &projectDetail.EndDate, &projectDetail.Content, &projectDetail.Technologies, &projectDetail.Image)

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, errQuery.Error())
	}

	projectDetail.Duration = countDuration(projectDetail.EndDate, projectDetail.StartDate)

	if checkValue(projectDetail.Technologies, "ReactJs") {
		projectDetail.ReactJs = true
	}
	if checkValue(projectDetail.Technologies, "JavaScript") {
		projectDetail.JavaScript = true
	}
	if checkValue(projectDetail.Technologies, "VueJs") {
		projectDetail.VueJs = true
	}
	if checkValue(projectDetail.Technologies, "NodeJs") {
		projectDetail.NodeJs = true
	}

	data := map[string]interface{}{
		"id":        id,
		"Project":   projectDetail,
		"startDate": projectDetail.StartDate.Format("02 Jan 2006"),
		"endDate":   projectDetail.EndDate.Format("02 Jan 2006"),
	}

	return tmpl.Execute(c.Response(), data)
}

func addProject(c echo.Context) error {
	tmpl, err := template.ParseFiles("html/add-my-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)
}

func contact(c echo.Context) error {
	tmpl, err := template.ParseFiles("html/contact-me.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)
}

func testimonial(c echo.Context) error {
	tmpl, err := template.ParseFiles("html/testimonial.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)
}

func formEditProject(c echo.Context) error {
	id := c.Param("id")

	data := map[string]interface{}{
		"Id": id,
	}
	tmpl, err := template.ParseFiles("html/edit-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), data)
}

func addDataProject(c echo.Context) error {

	title := c.FormValue("title")
	startdate := c.FormValue("startDate")
	enddate := c.FormValue("endDate")
	description := c.FormValue("description")
	chkBox1 := c.FormValue("checkbox1")
	chkBox2 := c.FormValue("checkbox2")
	chkBox3 := c.FormValue("checkbox3")
	chkBox4 := c.FormValue("checkbox4")

	startDate, _ := time.Parse("2006-01-02", startdate)
	endDate, _ := time.Parse("2006-01-02", enddate)

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_projects (name, start_date, end_date, description, technologies, image) VALUES ($1, $2, $3, $4, $5, $6)", title, startDate, endDate, description, []string{chkBox1, chkBox2, chkBox3, chkBox4}, "image.jpg")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func deleteProject(c echo.Context) error {
	id := c.Param("id")

	idToInt, _ := strconv.Atoi(id)

	connection.Conn.Exec(context.Background(), "DELETE FROM tb_projects WHERE id=$1", idToInt)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func editProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.FormValue("id"))

	title := c.FormValue("title")
	startdate := c.FormValue("startDate")
	enddate := c.FormValue("endDate")
	description := c.FormValue("description")
	chkBox1 := c.FormValue("checkbox1")
	chkBox2 := c.FormValue("checkbox2")
	chkBox3 := c.FormValue("checkbox3")
	chkBox4 := c.FormValue("checkbox4")

	startDate, _ := time.Parse("2006-01-02", startdate)
	endDate, _ := time.Parse("2006-01-02", enddate)

	_, err := connection.Conn.Exec(context.Background(), "UPDATE tb_projects SET name=$1, start_date=$2, end_date=$3, description=$4, technologies=$5, image=$6 WHERE id=$7", title, startDate, endDate, description, []string{chkBox1, chkBox2, chkBox3, chkBox4}, "image.jpg", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}
