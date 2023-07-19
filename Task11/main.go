package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/assets", "assets")

	e.GET("/intro", intro) //(route, function)
	e.GET("/home", home)
	e.GET("/addProject", addProject)
	e.GET("/contact", contact)
	e.GET("/testimonial", testimonial)
	e.GET("/projectDetail/:id", projectDetail)

	e.POST("/addDataProject", addDataProject)
	e.POST("/addDataContact", addDataContact)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

// handler / controller (di php)

func intro(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"name": "Hello World",
	})
}

func home(c echo.Context) error {
	tmpl, err := template.ParseFiles("index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)
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

	projectDetail := map[string]interface{}{
		"id":      id,
		"title":   "Dumbways Web App",
		"content": "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Ratione, odio dolore tempora voluptatem soluta placeat illum ea quia iste vitae neque quis nostrum animi nesciunt ipsum quidem, ab natus sapiente error! Asperiores voluptate eos eveniet pariatur consequatur assumenda laboriosam corporis, repellendus quasi quos eius expedita nam nemo! Architecto, reiciendis vitae recusandae suscipit ipsum natus quos quo rem fuga, perferendis nesciunt. Sit deleniti ducimus, aperiam eaque hic nobis doloribus explicabo perspiciatis magni delectus quos. Autem soluta est distinctio laboriosam culpa enim, excepturi temporibus numquam consequuntur nemo voluptas laudantium aliquam pariatur ratione non consectetur impedit, commodi harum dolore minima. Porro, ab illum?",
	}

	return tmpl.Execute(c.Response(), projectDetail)
}

func addDataProject(c echo.Context) error {

	title := c.FormValue("title")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	description := c.FormValue("description")
	nodejs := c.FormValue("checkbox1")
	reactjs := c.FormValue("checkbox2")
	vuejs := c.FormValue("checkbox3")
	javacscript := c.FormValue("checkbox4")

	fmt.Println("______________________________________________________________")
	fmt.Println("Projec tName: ", title)
	fmt.Println("Duration: ", startDate, "to", endDate)
	// fmt.Println("startDate: ", startDate)
	// fmt.Println("endDate: ", endDate)
	fmt.Println("Description: ", description)
	fmt.Println("Technologies: ", nodejs, "   ", reactjs, "   ", vuejs, "   ", javacscript)
	// fmt.Println("nodeJs: ", nodejs)
	// fmt.Println("reactJs: ", reactjs)
	// fmt.Println("vueJs: ", vuejs)
	// fmt.Println("JavaScript: ", javacscript)

	return c.Redirect(http.StatusMovedPermanently, "/addProject")
}

func addDataContact(c echo.Context) error {

	name := c.FormValue("name")
	email := c.FormValue("email")
	phoneNumber := c.FormValue("phoneNumber")
	subject := c.FormValue("subject")
	message := c.FormValue("message")

	fmt.Println("______________________________________________________________")
	fmt.Println("Name: ", name)
	fmt.Println("Email: ", email)
	fmt.Println("Phone Number: ", phoneNumber)
	fmt.Println("Subject: ", subject)
	fmt.Println("Message: ", message)

	return c.Redirect(http.StatusMovedPermanently, "/contact")
}
