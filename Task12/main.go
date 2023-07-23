package main

import (
	"html/template"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Id         int
	Title      string
	Content    string
	Duration   string
	StartDate  string
	EndDate    string
	NodeJs     bool
	ReactJs    bool
	VueJs      bool
	JavaScript bool
}

var dataProjects = []Project{
	{
		Title:      "Wolf",
		Content:    "The wolf has upright ears and it also resembles a dog with perky ears in other respects. Telling a wolf and a dog apart is often difficult, but the wolf's tail, which hangs straight down, is a tell-tale sign.",
		Duration:   "20 day",
		StartDate:  "01 Jan 2023",
		EndDate:    "21 Jan 2023",
		NodeJs:     true,
		ReactJs:    true,
		VueJs:      true,
		JavaScript: true,
	},
	{
		Title:      "Fox",
		Content:    "The fox has reddish-brown fur, a white chest and a bushy, white-tipped tail, called a brush. Its nose and ears are pointed.",
		Duration:   "1 month",
		StartDate:  "01 Feb 2023",
		EndDate:    "04 Mar 2023",
		NodeJs:     true,
		ReactJs:    true,
		VueJs:      false,
		JavaScript: true,
	},
	{
		Title:      "Stork",
		Content:    "Storks are large birds with long legs, necks, and bills. They are wading birds, which means they typically walk or stand in shallow water while feeding. There are 17 species, or types, of stork.",
		Duration:   "10 day",
		StartDate:  "05 Mar 2023",
		EndDate:    "15 Mar 2023",
		NodeJs:     true,
		ReactJs:    false,
		VueJs:      false,
		JavaScript: true,
	},
	{
		Title:      "Deer",
		Content:    "Deer, (family Cervidae), any of 43 species of hoofed ruminants in the order Artiodactyla, notable for having two large and two small hooves on each foot and also for having antlers in the males of most species and in the females of one species.",
		Duration:   "1 year",
		StartDate:  "01 Apr 2023",
		EndDate:    "02 Apr 2024",
		NodeJs:     false,
		ReactJs:    false,
		VueJs:      false,
		JavaScript: true,
	},
}

func main() {
	e := echo.New()

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

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	tmpl, err := template.ParseFiles("index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	data := map[string]interface{}{
		"Projects": dataProjects,
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

func projectDetail(c echo.Context) error {
	id := c.Param("id") // misal : 1

	tmpl, err := template.ParseFiles("html/project-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	idToInt, _ := strconv.Atoi(id)

	projectDetail := Project{}

	for index, data := range dataProjects {
		// index += 1
		if index == idToInt {
			projectDetail = Project{
				Title:      data.Title,
				Content:    data.Content,
				Duration:   data.Duration,
				StartDate:  data.StartDate,
				EndDate:    data.EndDate,
				NodeJs:     data.NodeJs,
				ReactJs:    data.ReactJs,
				VueJs:      data.VueJs,
				JavaScript: data.JavaScript,
			}
		}
	}

	data := map[string]interface{}{
		"id":      id,
		"Project": projectDetail,
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

	fmtStartDate := (startDate.Format("02 January 2006"))
	fmtEndtDate := (endDate.Format("02 January 2006"))

	newProject := Project{
		Title:      title,
		Content:    description,
		Duration:   countDuration(endDate, startDate),
		StartDate:  fmtStartDate,
		EndDate:    fmtEndtDate,
		NodeJs:     (chkBox1 == "NodeJs"),
		ReactJs:    (chkBox2 == "ReactJs"),
		VueJs:      (chkBox3 == "VueJs"),
		JavaScript: (chkBox4 == "JavaScript"),
	}

	dataProjects = append(dataProjects, newProject)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func deleteProject(c echo.Context) error {
	id := c.Param("id")

	idToInt, _ := strconv.Atoi(id)

	dataProjects = append(dataProjects[:idToInt], dataProjects[idToInt+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func formEditProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	editProject := Project{}

	for index, data := range dataProjects {
		// index += 1
		if id == index {
			editProject = Project{
				Id:         index,
				Title:      data.Title,
				Content:    data.Content,
				Duration:   data.Duration,
				StartDate:  data.StartDate,
				EndDate:    data.EndDate,
				NodeJs:     data.NodeJs,
				ReactJs:    data.ReactJs,
				VueJs:      data.VueJs,
				JavaScript: data.JavaScript,
			}
		}
	}

	data := map[string]interface{}{
		"Project": editProject,
	}
	var tmpl, err = template.ParseFiles("html/edit-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), data)
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

	fmtStartDate := (startDate.Format("02 January 2006"))
	fmtEndtDate := (endDate.Format("02 January 2006"))

	dataProjects[id].Title = title
	dataProjects[id].Content = description
	dataProjects[id].Duration = countDuration(endDate, startDate)
	dataProjects[id].StartDate = fmtStartDate
	dataProjects[id].EndDate = fmtEndtDate
	dataProjects[id].NodeJs = (chkBox1 == "NodeJs")
	dataProjects[id].ReactJs = (chkBox2 == "ReactJs")
	dataProjects[id].VueJs = (chkBox3 == "VueJs")
	dataProjects[id].JavaScript = (chkBox4 == "JavaScript")

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func countDuration(endDate time.Time, startDate time.Time) string {

	duration := endDate.Sub(startDate)
	days := (int(math.Floor(duration.Hours() / 24)))
	month := days / 30

	var result string

	if days > 30 {
		result = strconv.Itoa(month) + " month"
		if month >= 12 {
			result = strconv.Itoa(month/12) + " year"
		}
		return result
	}

	return strconv.Itoa(days) + " day"
}
