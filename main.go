package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// album represents data about a rpc user.
type user struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	//Birthday          string `json:"birthday"`
	Password          string `json:"password"`
	PasswordToConfirm string `json:"passwordToConfirm"`
	Email             string `json:"email"`
	Country           string `json:"country"`
}

var loggedUser = []user{}

// users slice to seed record album data.
var users = []user{
	{Username: "jkls1998", FirstName: "John Kennedy", LastName: "Landazuri Sandoval", Password: "johnkennedy", PasswordToConfirm: "johnkennedy", Email: "email1@gmail.com", Country: "Colombia"},
	{Username: "nimiasaca", FirstName: "Nimia", LastName: "Sandoval Carabalí", Password: "nimia", PasswordToConfirm: "nimia", Email: "email2@gmail.com", Country: "Colombia"},
	{Username: "kenken", FirstName: "Kennedy", LastName: "Landazuri Cortez", Password: "kenkenken", PasswordToConfirm: "kenkenken", Email: "email3@gmail.com", Country: "Colombia"},
	{Username: "stefa2005", FirstName: "Rut Stefany", LastName: "Landazuri Sandoval", Password: "stefany", PasswordToConfirm: "stefany", Email: "email4@gmail.com", Country: "Colombia"},
}

func main() {
	router := gin.Default()
	//add html files
	router.LoadHTMLFiles("login.html", "register.html", "usersPage.html")
	//add routes css
	router.Static("/css", "./css")

	router.GET("/", defaultCharge)
	router.GET("/users", loadUserLoginView)
	router.POST("/users", login)
	router.GET("/create", loadRegisterView)
	router.POST("/create", addUser)
	router.GET("/list", loadTableView)
	router.GET("/logout", logout)
	router.Run("localhost:8080")
}

func defaultCharge(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/users")
}

func loadTableView(c *gin.Context) {
	c.HTML(http.StatusOK, "usersPage.html", gin.H{
		"users": users,
	})
}

func loadRegisterView(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"answer": " ",
	})
}

func loadUserLoginView(c *gin.Context) {

	if len(loggedUser) != 0 {
		c.HTML(http.StatusOK, "usersPage.html", gin.H{
			"user":  loggedUser,
			"users": users,
		})
		return
	} else {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"message": " ",
		})
	}
}

func hsAndSalt(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)

	return string(hash), err
}

func addUser(c *gin.Context) {
	username := c.PostForm("username")
	firstname := c.PostForm("firstname")
	lastname := c.PostForm("lastname")
	//birthday := c.PostForm("birthday")
	password := c.PostForm("password")
	passwordToConfirm := c.PostForm("passwordToConfirm")
	email := c.PostForm("email")
	country := c.PostForm("country")

	if len(username) > 0 && len(password) > 0 && len(passwordToConfirm) > 0 && len(firstname) > 0 && len(lastname) > 0 && len(email) > 0 && len(country) > 0 {
		if password == passwordToConfirm {
			newUser := user{Username: username, FirstName: firstname, LastName: lastname, Password: password, PasswordToConfirm: passwordToConfirm, Email: email, Country: country}
			users = append(users, newUser)
			c.HTML(http.StatusOK, "login.html", gin.H{
				"message": "Usuario creado exitosamente",
			})
		} else {
			c.HTML(http.StatusOK, "create.html", gin.H{
				"answer": "Contraseñas no coinciden",
			})
		}
	} else {
		c.HTML(http.StatusOK, "usersPage.html", gin.H{
			"answer": "There can be no empty fields",
		})
	}
}

func comparePasswords(hsdPwd string, pwd []byte) error {
	byteHash := []byte(hsdPwd)
	log.Println("hsdp: " + hsdPwd + "   pwd:" + string(pwd))
	e := bcrypt.CompareHashAndPassword(byteHash, pwd)

	return e
}
func login(c *gin.Context) {
	//username := c.PostForm("Username")
	password := c.PostForm("Password")
	email := c.PostForm("Email")

	for _, a := range users {
		if email == a.Email {
			if password == a.Password {
				loggedUser := a
				c.HTML(http.StatusOK, "usersPage.html", gin.H{
					"email": loggedUser.Email,
					"users": users,
				})
				return
			} else {
				c.HTML(http.StatusOK, "login.html", gin.H{
					"message": "Contraseña Incorrecta",
				})
				return
			}
		}

	}
	c.HTML(http.StatusOK, "login.html", gin.H{
		"message": "Usuario no encontrado, se recomienda registrarse de nuevo",
	})

}

func logout(c *gin.Context) {
	loggedUser = []user{}
	c.Redirect(http.StatusMovedPermanently, "/users")
}
