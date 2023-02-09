package routes

import (

	"reflect"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ivanes285/GoPostgresAPI/db"
	"github.com/ivanes285/GoPostgresAPI/models"
)


func AddUsersGroup(app *fiber.App) {
	users := app.Group("/api/v1/users")
	users.Get("/", getUsers)
	users.Get("/:id", getUserById)
	users.Post("/", createUser)
	users.Put("/:id", updateUser)
	users.Delete("/:id", deleteUser)
}

// GET ALL USERS
func getUsers(c *fiber.Ctx) error {
     var users []models.User // creamos una variable de tipo User 
	 db.DB.Find(&users) // aqui se obtienen todos los usuarios de la BDD postgreSQL y se almacenan en la variable users
	return c.Status(200).JSON(fiber.Map{
		"users": users,
	})
}

// CREATE USER
func createUser(c *fiber.Ctx) error {

   var	user models.User // creamos una variable de tipo User

	if err := c.BodyParser(&user); //? &user es un puntero a la variable user en este caso se le pasa el valor del body y no hay necesidad de usar user = c.BodyParser() y mas bien verificamos si hay un error en el body y lo validamos con el if

	err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "An error has occurred with data,Invalid body"},
		)
	}
	
	// VALIDATE USER
	if user.FirstName == "" || user.FirstName == " " {
		return c.Status(400).JSON(fiber.Map{
			"message": "FirstName is required"},
		)
	}
	if user.LastName == "" || user.LastName == " " {
		return c.Status(400).JSON(fiber.Map{
			"message": "LastName is required"},
		)
	}
	if user.Email == "" || user.Email == " " {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email is required"},
		)
	}

	result:= db.DB.Create(&user) // aqui se crea el usuario en la BDD postgreSQL 
	
	err := result.Error
	var message string
	if err != nil {
		//?: reflect permite acceder a los campos de una estructura y obtener su valor
		code:= reflect.ValueOf(err).Elem().FieldByName("Code").String()  // aqui se obtiene el codigo de error de la BDD
		detail:= reflect.ValueOf(err).Elem().FieldByName("Detail").String() // aqui se obtiene el detalle del error de la BDD
	    td:= strings.Split(detail,"=")
		if code == "23505" {
			 message = "Campo Repetido"
		}
	
		return c.Status(400).JSON(fiber.Map{
			"messageError":message,
		     "error":td[1],
		},
		)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": user,
	
	})
}

// GET USER BY ID
func getUserById(c *fiber.Ctx) error {
   var userById models.User
	id := c.Params("id")
    
	db.DB.Find(&userById, id)
	return c.Status(200).JSON(fiber.Map{
		"user": userById,
	})
}

func updateUser(c *fiber.Ctx) error {

	var userUpdate models.User
	var	user models.User 

	id := c.Params("id")

	if err := c.BodyParser(&user);	err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "An error has occurred with data,Invalid body"},
		)
	}
	// VALIDATE fields USER
	if user.FirstName == "" || user.FirstName == " " {
		return c.Status(400).JSON(fiber.Map{
			"message": "FirstName is required"},
		)
	}
	if user.LastName == "" || user.LastName == " " {
		return c.Status(400).JSON(fiber.Map{
			"message": "LastName is required"},
		)
	}
	if user.Email == "" || user.Email == " " {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email is required"},
		)
	}

	db.DB.First(&userUpdate, id)
	 
     userUpdate.FirstName= user.FirstName
	 userUpdate.LastName= user.LastName
	 userUpdate.Email= user.Email

	result:=db.DB.Save(&userUpdate)


	err := result.Error
	var messageErr string
	if err != nil {
		//?: reflect permite acceder a los campos de una estructura y obtener su valor
		code:= reflect.ValueOf(err).Elem().FieldByName("Code").String()  // aqui se obtiene el codigo de error de la BDD
		detail:= reflect.ValueOf(err).Elem().FieldByName("Detail").String() // aqui se obtiene el detalle del error de la BDD
	    td:= strings.Split(detail,"=")
		campo:=strings.Split(td[1], " ")
	
		if code == "23505" {
			messageErr= "Ya existe un usuario con este dato "+campo[0]+" intente con otro"
			
		}

		return c.Status(400).JSON(fiber.Map{
			"messageError":messageErr, 
		},
		)
	}

	rowsAffected := result.RowsAffected
	var message string
	if rowsAffected == 0 {
		message = "The user does not exist "
	} else {
		message = "User updated"
	}
	return c.Status(200).JSON(fiber.Map{
		"users": message,
		"data": userUpdate,
	})
}

func deleteUser(c *fiber.Ctx) error {
	var userDelete models.User
	id := c.Params("id")
	// result:=db.DB.Delete(&userDelete, id)  //?Delete es un metodo que elimina un registro de la BDD pero solo visualmente(temporal), no se elimina en la BDD
    result:= db.DB.Unscoped().Delete(&userDelete, id) //? Unscoped es un metodo que elimina un registro de la BDD de forma permanente
	
	err := result.Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": " An error has occurred "},
		)
	}

	rowsAffected := result.RowsAffected
	var message string
	if rowsAffected == 0 {
		message = "The user does not exist or has already been deleted"
	} else {
		message = "User deleted"
	}
	return c.Status(200).JSON(fiber.Map{
		"message": message,
		
	})
}
