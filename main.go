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

func loadViewLogin(c *gin.Context) {

	if len(userLogged) != 0 {
		c.HTML(http.StatusOK, "users.html", gin.H{
			"user":  userLogged,
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
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirmPassword")
	firstname := c.PostForm("firstname")
	lastname := c.PostForm("lastname")
	birthdate := c.PostForm("birthdate")

	if len(username) > 0 && len(password) > 0 && len(confirmPassword) > 0 && len(firstname) > 0 && len(lastname) > 0 && len(birthdate) > 0 {
		if password == confirmPassword {
			newUser := user{Username: username, Password: password, ConfirmPassword: confirmPassword, FirstName: firstname, LastName: lastname, Birthdate: birthdate}
			users = append(users, newUser)
			c.HTML(http.StatusOK, "login.html", gin.H{
				"message": "your user was create successfully",
			})
		} else {
			c.HTML(http.StatusOK, "create.html", gin.H{
				"answer": "The passwords are not equals",
			})
		}
	} else {
		c.HTML(http.StatusOK, "create.html", gin.H{
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
				userLogged := a
				c.HTML(http.StatusOK, "users.html", gin.H{
					"username": userLogged.Username,
					"users":    users,
				})
				return
			} else {
				c.HTML(http.StatusOK, "login.html", gin.H{
					"message": "Incorrect Password",
				})
				return
			}
		}

	}
	c.HTML(http.StatusOK, "login.html", gin.H{
		"message": "This user doesn't exist",
	})

}

func logout(c *gin.Context) {
	userLogged = []user{}
	c.Redirect(http.StatusMovedPermanently, "/users")
}
