package main

import (
	"context"
	"html/template"
	"math"
	"net/http"
	"strconv"
	"task13/connection"
	"task13/middleware"
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
	Author       string
	UserId       int
	Technologies []string
	NodeJs       bool
	ReactJs      bool
	VueJs        bool
	JavaScript   bool
	Image        string
	GetRole      string
	CekUserId    bool
}

type User struct {
	Id             int
	Name           string
	Email          string
	HashedPassword string
	Role           string
}

type UserLoginSession struct {
	IsLogin bool
	Name    string
	CekRole bool
}

var userLoginSession = UserLoginSession{}

func main() {
	e := echo.New()

	connection.DatabaseConnection()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secretkeycookie"))))

	e.Static("/assets", "assets")
	e.Static("/uploads", "uploads")

	e.GET("/", home)
	e.GET("/addProject", addProject)
	e.GET("/contact", contact)
	e.GET("/testimonial", testimonial)
	e.GET("/projectDetail/:id", projectDetail)
	e.GET("/formEditProject/:id", formEditProject)

	e.POST("/addDataProject", middleware.UploadFile(addDataProject))
	e.POST("/deleteProject/:id", deleteProject)
	e.POST("/editProject", middleware.UploadFile(editProject))

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

	dataProjects, errBlogs := connection.Conn.Query(context.Background(), "SELECT tb_user.id, tb_user.role, tb_user.name, tb_projects.id, tb_projects.title, tb_projects.start_date, tb_projects.end_date, tb_projects.description, tb_projects.technologies, tb_projects.image, tb_projects.user_id  FROM tb_projects LEFT JOIN tb_user ON tb_projects.user_id = tb_user.id")

	if errBlogs != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	var resultProjects []Project

	for dataProjects.Next() {
		var each = Project{}
		var eachUser = User{}

		err := dataProjects.Scan(&eachUser.Id, &eachUser.Role, &eachUser.Name, &each.Id, &each.Title, &each.StartDate, &each.EndDate, &each.Content, &each.Technologies, &each.Image, &each.UserId)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		each.Duration = countDuration(each.EndDate, each.StartDate)
		each.Author = eachUser.Name
		each.GetRole = eachUser.Role

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
		if sess.Values["Id"] == each.UserId {
			each.CekUserId = true
		}

		resultProjects = append(resultProjects, each)
	}

	if sess.Values["isLogin"] != true {
		userLoginSession.IsLogin = false
	} else {
		userLoginSession.IsLogin = true
		userLoginSession.Name = sess.Values["Name"].(string)
	}

	if sess.Values["Role"] != "Admin" {
		userLoginSession.CekRole = false
	} else {
		userLoginSession.CekRole = true
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

	var each = Project{}
	var eachUser = User{}

	errQuery := connection.Conn.QueryRow(context.Background(), "SELECT tb_user.id, tb_user.name, tb_projects.id, tb_projects.title, tb_projects.start_date, tb_projects.end_date, tb_projects.description, tb_projects.technologies, tb_projects.image, tb_projects.user_id  FROM tb_projects LEFT JOIN tb_user ON tb_projects.user_id = tb_user.id WHERE tb_projects.id=$1", idToInt).Scan(&eachUser.Id, &eachUser.Name, &each.Id, &each.Title, &each.StartDate, &each.EndDate, &each.Content, &each.Technologies, &each.Image, &each.UserId)

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, errQuery.Error())
	}

	each.Duration = countDuration(each.EndDate, each.StartDate)
	each.Author = eachUser.Name

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
		"Project":          each,
		"startDate":        each.StartDate.Format("02 Jan 2006"),
		"endDate":          each.EndDate.Format("02 Jan 2006"),
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

	var each = Project{}
	var eachUser = User{}

	errQuery := connection.Conn.QueryRow(context.Background(), "SELECT * FROM public.tb_projects WHERE id=$1", idToInt).Scan(&each.Id, &each.Title, &each.StartDate, &each.EndDate, &each.Content, &each.Technologies, &each.Image, &each.UserId)

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, errQuery.Error())
	}

	each.Duration = countDuration(each.EndDate, each.StartDate)
	each.Author = eachUser.Name

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
		"Id":               id,
		"Project":          each,
		"StartDate":        each.StartDate.Format("2006-01-02"),
		"EndDate":          each.EndDate.Format("2006-01-02"),
		"Flash":            flash,
		"UserLoginSession": userLoginSession,
	}

	return tmpl.Execute(c.Response(), data)
}

func addDataProject(c echo.Context) error {
	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	title := c.FormValue("title")
	startdate := c.FormValue("startDate")
	enddate := c.FormValue("endDate")
	description := c.FormValue("description")
	chkBox1 := c.FormValue("nodejs")
	chkBox2 := c.FormValue("reactjs")
	chkBox3 := c.FormValue("vuejs")
	chkBox4 := c.FormValue("javascript")
	technologies := []string{chkBox1, chkBox2, chkBox3, chkBox4}
	userId := sess.Values["Id"].(int)
	image := c.Get("dataFile").(string)

	startDate, _ := time.Parse("2006-01-02", startdate)
	endDate, _ := time.Parse("2006-01-02", enddate)

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_projects (title, start_date, end_date, description, technologies, image, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7)", title, startDate, endDate, description, technologies, image, userId)

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
	Id, _ := strconv.Atoi(c.FormValue("id"))

	title := c.FormValue("title")
	startdate := c.FormValue("startDate")
	enddate := c.FormValue("endDate")
	description := c.FormValue("description")
	chkBox1 := c.FormValue("nodejs")
	chkBox2 := c.FormValue("reactjs")
	chkBox3 := c.FormValue("vuejs")
	chkBox4 := c.FormValue("javascript")
	technologies := []string{chkBox1, chkBox2, chkBox3, chkBox4}
	image := c.Get("dataFile").(string)

	startDate, _ := time.Parse("2006-01-02", startdate)
	endDate, _ := time.Parse("2006-01-02", enddate)

	_, err := connection.Conn.Exec(context.Background(), "UPDATE tb_projects SET title=$1, start_date=$2, end_date=$3, description=$4, technologies=$5, image=$6 WHERE id=$7", title, startDate, endDate, description, technologies, image, Id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
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
	role := c.FormValue("role")
	email := c.FormValue("email")
	password := c.FormValue("password")

	hashedPss, errPss := bcrypt.GenerateFromPassword([]byte(password), 10)

	if errPss != nil {
		return c.JSON(http.StatusInternalServerError, errPss.Error())
	}

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_user (name, email, password, role) VALUES ($1, $2, $3, $4)", name, email, hashedPss, role)

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

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, name, email, password, role FROM tb_user WHERE email=$1", email).Scan(&user.Id, &user.Name, &user.Email, &user.HashedPassword, &user.Role)

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

	sess.Options.MaxAge = 10800
	sess.Values["Message"] = "Login Success"
	sess.Values["Status"] = true
	sess.Values["Id"] = user.Id
	sess.Values["Name"] = user.Name
	sess.Values["Email"] = user.Email
	sess.Values["Role"] = user.Role
	sess.Values["isLogin"] = true

	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
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
