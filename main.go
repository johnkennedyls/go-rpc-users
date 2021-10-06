package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a rpc user.
type user struct {
	Username          string `json:"username"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	Birthday          string `json:"birthday"`
	Password          string `json:"password"`
	PasswordToConfirm string `json:"passwordToConfirm"`
}

var loggedUser = []user{}

// users slice to seed record album data.
var users = []user{
	{Username: "jkls1998", FirstName: "John Kennedy", LastName: "Landazuri Sandoval", Birthday: "17/10/1998", Password: "johnkennedy", PasswordToConfirm: "johnkennedy"},
	{Username: "nimiasaca", FirstName: "Nimia", LastName: "Sandoval Carabalí", Birthday: "6/08/1967", Password: "nimia", PasswordToConfirm: "nimia"},
	{Username: "kenken", FirstName: "Kennedy", LastName: "Landazuri Cortez", Birthday: "11/11/1972", Password: "kenkenken", PasswordToConfirm: "kenkenken"},
	{Username: "stefa2005", FirstName: "Rut Stefany", LastName: "Landazuri Sandoval", Birthday: "11/08/2005", Password: "stefany", PasswordToConfirm: "stefany"},
}

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("login.html", "create.html", "users.html")
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

func addUser(c *gin.Context) {
	username := c.PostForm("username")
	firstname := c.PostForm("firstname")
	lastname := c.PostForm("lastname")
	birthday := c.PostForm("birthday")
	password := c.PostForm("password")
	passwordToConfirm := c.PostForm("passwordToConfirm")

	if len(username) > 0 && len(password) > 0 && len(passwordToConfirm) > 0 && len(firstname) > 0 && len(lastname) > 0 && len(birthday) > 0 {
		if password == passwordToConfirm {
			newUser := user{Username: username, FirstName: firstname, LastName: lastname, Birthday: birthday, Password: password, PasswordToConfirm: passwordToConfirm}
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

func login(c *gin.Context) {
	username := c.PostForm("Username")
	password := c.PostForm("Password")

	for _, a := range users {
		if username == a.Username {
			if password == a.Password {
				loggedUser := a
				c.HTML(http.StatusOK, "usersPage.html", gin.H{
					"username": loggedUser.Username,
					"users":    users,
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
