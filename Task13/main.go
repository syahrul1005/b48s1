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

var dataProjects = []Project{}

func main() {
	e := echo.New()

	connection.DatabaseConnection()

	e.Static("/assets", "assets")

	e.GET("/", home)
	e.GET("/addProject", addProject)
	e.GET("/contact", contact)
	e.GET("/testimonial", testimonial)
	e.GET("/projectDetail/:id", projectDetail)
	// e.GET("/formEditProject/:id", formEditProject)

	// e.POST("/addDataProject", addDataProject)
	// e.POST("/deleteProject/:id", deleteProject)
	// e.POST("/editProject", editProject)
	// e.POST("/addDataContact", addDataContact)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	tmpl, err := template.ParseFiles("index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	dataBlogs, errBlogs := connection.Conn.Query(context.Background(), "SELECT id, name, start_date, end_date, description, technologies, image FROM tb_projects")

	if errBlogs != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var resultBlogs []Project
	for dataBlogs.Next() {
		var each = Project{}

		err := dataBlogs.Scan(&each.Id, &each.Title, &each.StartDate, &each.EndDate, &each.Content, &each.Technologies, &each.Image)

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

		resultBlogs = append(resultBlogs, each)
	}

	data := map[string]interface{}{
		"Projects": resultBlogs,
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
	id := c.Param("Id")

	idToInt, _ := strconv.Atoi(id)

	projectDetail := Project{}

	for index, data := range dataProjects {

		if index == idToInt {
			projectDetail = Project{
				Id:           data.Id,
				Title:        data.Title,
				Content:      data.Content,
				Duration:     data.Duration,
				StartDate:    data.StartDate,
				EndDate:      data.EndDate,
				Technologies: data.Technologies,
			}
		}
	}

	data := map[string]interface{}{
		"id":      id,
		"Project": projectDetail,
	}

	tmpl, err := template.ParseFiles("html/project-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
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

// func addDataProject(c echo.Context) error {

// 	title := c.FormValue("title")
// 	startdate := c.FormValue("startDate")
// 	enddate := c.FormValue("endDate")
// 	description := c.FormValue("description")
// 	chkBox1 := c.FormValue("checkbox1")
// 	chkBox2 := c.FormValue("checkbox2")
// 	chkBox3 := c.FormValue("checkbox3")
// 	chkBox4 := c.FormValue("checkbox4")

// 	startDate, _ := time.Parse("2006-01-02", startdate)
// 	endDate, _ := time.Parse("2006-01-02", enddate)

// 	// fmtStartDate := (startDate.Format("02 January 2006"))
// 	// fmtEndtDate := (endDate.Format("02 January 2006"))

// 	newProject := Project{
// 		Title:      title,
// 		Content:    description,
// 		Duration:   countDuration(endDate, startDate),
// 		StartDate:  startDate,
// 		EndDate:    endDate,
// 		NodeJs:     (chkBox1 == "NodeJs"),
// 		ReactJs:    (chkBox2 == "ReactJs"),
// 		VueJs:      (chkBox3 == "VueJs"),
// 		JavaScript: (chkBox4 == "JavaScript"),
// 	}

// 	dataProjects = append(dataProjects, newProject)

// 	return c.Redirect(http.StatusMovedPermanently, "/")
// }

// func deleteProject(c echo.Context) error {
// 	id := c.Param("id")

// 	idToInt, _ := strconv.Atoi(id)

// 	dataProjects = append(dataProjects[:idToInt], dataProjects[idToInt+1:]...)

// 	return c.Redirect(http.StatusMovedPermanently, "/")
// }

// func formEditProject(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	editProject := Project{}

// 	for index, data := range dataProjects {
// 		// index += 1
// 		if id == index {
// 			editProject = Project{
// 				Id:         index,
// 				Title:      data.Title,
// 				Content:    data.Content,
// 				Duration:   data.Duration,
// 				StartDate:  data.StartDate,
// 				EndDate:    data.EndDate,
// 				NodeJs:     data.NodeJs,
// 				ReactJs:    data.ReactJs,
// 				VueJs:      data.VueJs,
// 				JavaScript: data.JavaScript,
// 			}
// 		}
// 	}

// 	data := map[string]interface{}{
// 		"Project": editProject,
// 	}
// 	var tmpl, err = template.ParseFiles("html/edit-project.html")

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}

// 	return tmpl.Execute(c.Response(), data)
// }

// func editProject(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.FormValue("id"))

// 	title := c.FormValue("title")
// 	startdate := c.FormValue("startDate")
// 	enddate := c.FormValue("endDate")
// 	description := c.FormValue("description")
// 	chkBox1 := c.FormValue("checkbox1")
// 	chkBox2 := c.FormValue("checkbox2")
// 	chkBox3 := c.FormValue("checkbox3")
// 	chkBox4 := c.FormValue("checkbox4")

// 	startDate, _ := time.Parse("2006-01-02", startdate)
// 	endDate, _ := time.Parse("2006-01-02", enddate)

// 	// fmtStartDate := (startDate.Format("02 January 2006"))
// 	// fmtEndtDate := (endDate.Format("02 January 2006"))

// 	dataProjects[id].Title = title
// 	dataProjects[id].Content = description
// 	dataProjects[id].Duration = countDuration(endDate, startDate)
// 	dataProjects[id].StartDate = startDate
// 	dataProjects[id].EndDate = endDate
// 	dataProjects[id].NodeJs = (chkBox1 == "NodeJs")
// 	dataProjects[id].ReactJs = (chkBox2 == "ReactJs")
// 	dataProjects[id].VueJs = (chkBox3 == "VueJs")
// 	dataProjects[id].JavaScript = (chkBox4 == "JavaScript")

// 	return c.Redirect(http.StatusMovedPermanently, "/")
// }
