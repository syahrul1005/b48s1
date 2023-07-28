package main

import (
	"context"
	"html/template"
	"math"
	"net/http"
	"strconv"
	"task13/connection"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Project struct {
	Id           int
	Title        string
	Content      string
	Duration     string
	StartDate    time.Time
	EndDate      time.Time
	CreatedAt    time.Time
	CreatedAtStr string
	UserId       int
	Technologies []string
	NodeJs       bool
	ReactJs      bool
	VueJs        bool
	JavaScript   bool
	Image        string
}

type User struct {
	Id             int
	Name           string
	Email          string
	HashedPassword string
}

type UserLoginSession struct {
	IsLogin bool
	Name    string
}

var userLoginSession = UserLoginSession{}

func main() {
	e := echo.New()

	connection.DatabaseConnection()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secretkeycookie"))))

	e.Static("/assets", "assets")

	e.GET("/", home)
	e.GET("/my-project", myProject)
	e.GET("/addProject", addProject)
	e.GET("/contact", contact)
	e.GET("/testimonial", testimonial)
	e.GET("/projectDetail/:id", projectDetail)
	e.GET("/formEditProject/:id", formEditProject)

	e.POST("/addDataProject", addDataProject)
	e.POST("/deleteProject/:id", deleteProject)
	e.POST("/editProject", editProject)

	// e.POST("/addDataContact", addDataContact)

	// auth
	e.GET("/form-register", formRegister)
	e.GET("/form-login", formLogin)

	e.POST("/register", register)
	e.POST("/login", login)

	e.POST("/logout", logout)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	tmpl, err := template.ParseFiles("index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	dataProjects, errBlogs := connection.Conn.Query(context.Background(), "SELECT id, name, start_date, end_date, description, technologies, image, created_at, user_id FROM public.tb_projects")

	if errBlogs != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var resultProjects []Project
	for dataProjects.Next() {
		var each = Project{}

		err := dataProjects.Scan(&each.Id, &each.Title, &each.StartDate, &each.EndDate, &each.Content, &each.Technologies, &each.Image, &each.CreatedAt, &each.UserId)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		each.Duration = countDuration(each.EndDate, each.StartDate)

		each.CreatedAtStr = PostAt(each.CreatedAt)

		// fmt.Println("Long date:", each.CreatedAt.Format("Mon, Jan 2 2006 15:04:05 "))

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

	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	if sess.Values["isLogin"] != true {
		userLoginSession.IsLogin = false
	} else {
		userLoginSession.IsLogin = true
		userLoginSession.Name = sess.Values["Name"].(string)
	}

	data := map[string]interface{}{
		"Projects":         resultProjects,
		"UserLoginSession": userLoginSession,
	}

	return tmpl.Execute(c.Response(), data)
}

func PostAt(datePost time.Time) string {
	postAt := datePost.Format("02 Jan 2006 15:04")

	return postAt
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

func myProject(c echo.Context) error {
	tmpl, err := template.ParseFiles("html/my-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	dataProjects, errBlogs := connection.Conn.Query(context.Background(), "SELECT id, name, start_date, end_date, description, technologies, image, created_at, user_id FROM public.tb_projects")

	if errBlogs != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var resultProjects []Project
	for dataProjects.Next() {
		var each = Project{}

		err := dataProjects.Scan(&each.Id, &each.Title, &each.StartDate, &each.EndDate, &each.Content, &each.Technologies, &each.Image, &each.CreatedAt, &each.UserId)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		each.Duration = countDuration(each.EndDate, each.StartDate)

		each.CreatedAtStr = PostAt(each.CreatedAt)

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

	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	if sess.Values["isLogin"] != true {
		userLoginSession.IsLogin = false
	} else {
		userLoginSession.IsLogin = true
		userLoginSession.Name = sess.Values["Name"].(string)
	}

	flash := map[string]interface{}{
		"FlashMessage": sess.Values["Message"],
		"FlashStatus":  sess.Values["Status"],
	}

	delete(sess.Values, "Message")
	delete(sess.Values, "Status")
	sess.Save(c.Request(), c.Response())

	data := map[string]interface{}{
		"Projects":         resultProjects,
		"Flash":            flash,
		"UserLoginSession": userLoginSession,
	}

	return tmpl.Execute(c.Response(), data)
}

func projectDetail(c echo.Context) error {
	id := c.Param("id")

	tmpl, err := template.ParseFiles("html/project-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	idToInt, _ := strconv.Atoi(id)

	projectDetail := Project{}

	errQuery := connection.Conn.QueryRow(context.Background(), "SELECT * FROM public.tb_projects WHERE id=$1", idToInt).Scan(&projectDetail.Id, &projectDetail.Title, &projectDetail.StartDate, &projectDetail.EndDate, &projectDetail.Content, &projectDetail.Technologies, &projectDetail.Image, &projectDetail.CreatedAt, &projectDetail.UserId)

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, errQuery.Error())
	}

	projectDetail.Duration = countDuration(projectDetail.EndDate, projectDetail.StartDate)
	projectDetail.CreatedAtStr = PostAt(projectDetail.CreatedAt)

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

	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	if sess.Values["isLogin"] != true {
		userLoginSession.IsLogin = false
	} else {
		userLoginSession.IsLogin = true
		userLoginSession.Name = sess.Values["Name"].(string)
	}

	data := map[string]interface{}{
		"id":               id,
		"Project":          projectDetail,
		"startDate":        projectDetail.StartDate.Format("02 Jan 2006"),
		"endDate":          projectDetail.EndDate.Format("02 Jan 2006"),
		"UserLoginSession": userLoginSession,
	}

	return tmpl.Execute(c.Response(), data)
}

func addProject(c echo.Context) error {
	tmpl, err := template.ParseFiles("html/add-my-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	if sess.Values["isLogin"] != true {
		userLoginSession.IsLogin = false
	} else {
		userLoginSession.IsLogin = true
		userLoginSession.Name = sess.Values["Name"].(string)
	}

	flash := map[string]interface{}{
		"UserLoginSession": userLoginSession,
	}

	return tmpl.Execute(c.Response(), flash)
}

func contact(c echo.Context) error {
	tmpl, err := template.ParseFiles("html/contact-me.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	if sess.Values["isLogin"] != true {
		userLoginSession.IsLogin = false
	} else {
		userLoginSession.IsLogin = true
		userLoginSession.Name = sess.Values["Name"].(string)
	}

	flash := map[string]interface{}{
		"UserLoginSession": userLoginSession,
	}

	return tmpl.Execute(c.Response(), flash)
}

func testimonial(c echo.Context) error {
	tmpl, err := template.ParseFiles("html/testimonial.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	if sess.Values["isLogin"] != true {
		userLoginSession.IsLogin = false
	} else {
		userLoginSession.IsLogin = true
		userLoginSession.Name = sess.Values["Name"].(string)
	}

	flash := map[string]interface{}{
		"UserLoginSession": userLoginSession,
	}

	return tmpl.Execute(c.Response(), flash)
}

func formEditProject(c echo.Context) error {
	id := c.Param("id")

	tmpl, err := template.ParseFiles("html/edit-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	idToInt, _ := strconv.Atoi(id)

	projectDetail := Project{}

	errQuery := connection.Conn.QueryRow(context.Background(), "SELECT * FROM public.tb_projects WHERE id=$1", idToInt).Scan(&projectDetail.Id, &projectDetail.Title, &projectDetail.StartDate, &projectDetail.EndDate, &projectDetail.Content, &projectDetail.Technologies, &projectDetail.Image, &projectDetail.CreatedAt, &projectDetail.UserId)

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, errQuery.Error())
	}

	projectDetail.Duration = countDuration(projectDetail.EndDate, projectDetail.StartDate)
	projectDetail.CreatedAtStr = PostAt(projectDetail.CreatedAt)

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

	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	if sess.Values["isLogin"] != true {
		userLoginSession.IsLogin = false
	} else {
		userLoginSession.IsLogin = true
		userLoginSession.Name = sess.Values["Name"].(string)
	}

	flash := map[string]interface{}{
		"FlashMessage": sess.Values["Message"],
		"FlashStatus":  sess.Values["Status"],
	}

	data := map[string]interface{}{
		"id":               id,
		"Project":          projectDetail,
		"StartDate":        projectDetail.StartDate.Format("2006-01-02"),
		"EndDate":          projectDetail.EndDate.Format("2006-01-02"),
		"Flash":            flash,
		"UserLoginSession": userLoginSession,
	}

	return tmpl.Execute(c.Response(), data)
}

func addDataProject(c echo.Context) error {

	title := c.FormValue("title")
	startdate := c.FormValue("startDate")
	enddate := c.FormValue("endDate")
	description := c.FormValue("description")
	chkBox1 := c.FormValue("nodejs")
	chkBox2 := c.FormValue("reactjs")
	chkBox3 := c.FormValue("vuejs")
	chkBox4 := c.FormValue("javascript")
	technologies := []string{chkBox1, chkBox2, chkBox3, chkBox4}

	startDate, _ := time.Parse("2006-01-02", startdate)
	endDate, _ := time.Parse("2006-01-02", enddate)

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_projects (name, start_date, end_date, description, technologies, image) VALUES ($1, $2, $3, $4, $5, $6)", title, startDate, endDate, description, technologies, "image.jpg")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(http.StatusMovedPermanently, "/my-project")
}

func deleteProject(c echo.Context) error {
	id := c.Param("id")

	idToInt, _ := strconv.Atoi(id)

	connection.Conn.Exec(context.Background(), "DELETE FROM tb_projects WHERE id=$1", idToInt)

	return c.Redirect(http.StatusMovedPermanently, "/my-project")
}

func editProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.FormValue("id"))

	title := c.FormValue("title")
	startdate := c.FormValue("startDate")
	enddate := c.FormValue("endDate")
	description := c.FormValue("description")
	chkBox1 := c.FormValue("nodejs")
	chkBox2 := c.FormValue("reactjs")
	chkBox3 := c.FormValue("vuejs")
	chkBox4 := c.FormValue("javascript")
	technologies := []string{chkBox1, chkBox2, chkBox3, chkBox4}

	startDate, _ := time.Parse("2006-01-02", startdate)
	endDate, _ := time.Parse("2006-01-02", enddate)

	_, err := connection.Conn.Exec(context.Background(), "UPDATE tb_projects SET name=$1, start_date=$2, end_date=$3, description=$4, technologies=$5, image=$6 WHERE id=$7", title, startDate, endDate, description, technologies, "image.jpg", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(http.StatusMovedPermanently, "/my-project")
}

func formRegister(c echo.Context) error {
	tmpl, err := template.ParseFiles("html/form-register.html")

	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	flash := map[string]interface{}{
		"FlashMessage":     sess.Values["Message"],
		"FlashStatus":      sess.Values["Status"],
		"UserLoginSession": userLoginSession,
	}

	delete(sess.Values, "Message")
	delete(sess.Values, "Status")
	sess.Save(c.Request(), c.Response())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return tmpl.Execute(c.Response(), flash)
}

func register(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	hashedPss, errPss := bcrypt.GenerateFromPassword([]byte(password), 10)

	if errPss != nil {
		return c.JSON(http.StatusInternalServerError, errPss.Error())
	}

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_user (name, email, password) VALUES ($1, $2, $3)", name, email, hashedPss)

	if err != nil {
		return redirectWithMessage(c, "Register Failure", false, "/form-register")
	}

	return redirectWithMessage(c, "Register Success", true, "/form-login")
}

func formLogin(c echo.Context) error {

	tmpl, err := template.ParseFiles("html/form-login.html")

	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	flash := map[string]interface{}{
		"FlashMessage":     sess.Values["Message"],
		"FlashStatus":      sess.Values["Status"],
		"UserLoginSession": userLoginSession,
	}

	delete(sess.Values, "Message")
	delete(sess.Values, "Status")
	sess.Save(c.Request(), c.Response())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return tmpl.Execute(c.Response(), flash)
}

func login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	user := User{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, name, email, password FROM tb_user WHERE email=$1", email).Scan(&user.Id, &user.Name, &user.Email, &user.HashedPassword)

	if err != nil {
		return redirectWithMessage(c, "Incorrect email or password", false, "/form-login")
	}

	errPss := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))

	if errPss != nil {
		return redirectWithMessage(c, "Incorrect email or password", false, "/form-login")
	}

	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	sess.Options.MaxAge = 10800 //3 hours
	sess.Values["Message"] = "Login Success"
	sess.Values["Status"] = true
	sess.Values["Id"] = user.Id
	sess.Values["Name"] = user.Name
	sess.Values["Email"] = user.Email
	sess.Values["isLogin"] = true

	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/my-project")
}

func redirectWithMessage(c echo.Context, Message string, Status bool, redirectPath string) error {
	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	sess.Values["Message"] = Message
	sess.Values["Status"] = Status
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, redirectPath)
}

func logout(c echo.Context) error {
	sess, _ := session.Get("session", c)

	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}
