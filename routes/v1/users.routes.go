package routes

import (
	"fmt"
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

	// coll := db.GetDBCollection("users") // obtenemos la coleccion de la base de datos goreactmongo que se encuentra en el paquete db

	// var users []models.User                          //creamos un slice(array de tamño dinamico) de tipo User
	// results, err := coll.Find(c.Context(), bson.M{}) //retorna primero en formato bson
	// if err != nil {
	// 	return c.Status(400).JSON(fiber.Map{
	// 		"message": err.Error()},
	// 	)
	// }

	// if err = results.All(c.Context(), &users); // aqui recien estamos asignando el resultado a la variable users que es un slice de tipo User
	// err != nil {
	// 	return c.Status(400).JSON(fiber.Map{
	// 		"message": "Could not find users"},
	// 	)
	// }

	// return c.Status(200).JSON(fiber.Map{
	// 	"users": users,
	// })
	return c.Status(200).JSON(fiber.Map{
		"users": "users",
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
		code:= reflect.ValueOf(err).Elem().FieldByName("Code").String()
		detail:= reflect.ValueOf(err).Elem().FieldByName("Detail").String()
	    td:= strings.Split(detail,"=")
		if code == "23505" {
			 message = "Campo Repetido"
		}
		fmt.Println("typo de DATO",reflect.TypeOf(err))
		return c.Status(400).JSON(fiber.Map{
			"messageError":message,
		     "error":td[1],
		},
		)
	}

	return c.Status(200).JSON(fiber.Map{
		"RowsAffected": result.RowsAffected ,
		"user": user,
	

	})
}

// GET USER BY ID
func getUserById(c *fiber.Ctx) error {
	// coll := db.GetDBCollection("users")

	// // find the user by id
	// id := c.Params("id") //obtenemos el id del parametro
	// if id == "" {
	// 	return c.Status(400).JSON(fiber.Map{
	// 		"error": "id is required",
	// 	})
	// }

	// objectId, err := primitive.ObjectIDFromHex(id) //convertimos el id a un objeto de tipo ObjectID
	// if err != nil {
	// 	return c.Status(400).JSON(fiber.Map{
	// 		"error": "invalid id",
	// 	})
	// }

	// user := models.User{}

	// err = coll.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&user)
	// if err != nil {
	// 	return c.Status(500).JSON(fiber.Map{
	// 		"error": "Could not find user",
	// 	})
	// }

	// return c.Status(200).JSON(fiber.Map{"user": user})
	return c.Status(200).JSON(fiber.Map{
		"users": "User by id",
	})
}

func updateUser(c *fiber.Ctx) error {
	// 	// validate the body
	// 	b := new(models.User)
	// 	if err := c.BodyParser(b); err != nil {
	// 		return c.Status(400).JSON(fiber.Map{
	// 			"error": "Invalid body",
	// 		})
	// 	}

	// 	// get the id
	// 	id := c.Params("id")
	// 	if id == "" {
	// 		return c.Status(400).JSON(fiber.Map{
	// 			"error": "id is required",
	// 		})
	// 	}
	// 	objectId, err := primitive.ObjectIDFromHex(id)
	// 	if err != nil {
	// 		return c.Status(400).JSON(fiber.Map{
	// 			"error": "invalid id",
	// 		})
	// 	}

	// 	// update the user
	// 	coll := db.GetDBCollection("users")
	// 	result, err := coll.UpdateOne(c.Context(), bson.M{"_id": objectId}, bson.M{"$set": b})
	// 	if err != nil {
	// 		return c.Status(500).JSON(fiber.Map{
	// 			"error":   "Failed to update book",
	// 			"message": err.Error(),
	// 		})
	// 	}
	// 	message :="User update successfully"
	// 	status := 200

	//    if result.ModifiedCount == 0 {
	// 	   message = "User not found or already updated"
	// 	   status = 404
	//    }
	// 	// return the user
	// 	return c.Status(status).JSON(fiber.Map{
	// 		"message": message,
	// 	})
	return c.Status(200).JSON(fiber.Map{
		"users": "User updated",
	})
}

func deleteUser(c *fiber.Ctx) error {
	// // get the id
	// id := c.Params("id")
	// if id == "" {
	// 	return c.Status(400).JSON(fiber.Map{
	// 		"error": "id is required",
	// 	})
	// }
	// objectId, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	return c.Status(400).JSON(fiber.Map{
	// 		"error": "invalid id",
	// 	})
	// }

	// // delete the user
	// coll := db.GetDBCollection("users")
	// result, err := coll.DeleteOne(c.Context(), bson.M{"_id": objectId})
	// if err != nil {
	// 	return c.Status(500).JSON(fiber.Map{
	// 		"error":   "Could not delete user.Failed to delete user",
	// 		"message": err.Error(),
	// 	})
	// }

	// message := "User deleted successfully"
	// status := 200

	// if result.DeletedCount == 0 {
	// 	message = "User not found or already deleted"
	// 	status = 404
	// }

	// return c.Status(status).JSON(fiber.Map{
	// 	"message": message,
	// })
	return c.Status(200).JSON(fiber.Map{
		"users": "User deleted",
	})
}
