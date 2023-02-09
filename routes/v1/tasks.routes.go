package routes

import (
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ivanes285/GoPostgresAPI/db"
	"github.com/ivanes285/GoPostgresAPI/models"
)

func AddTasksGroup(app *fiber.App) {
	users := app.Group("/api/v1/tasks")
	users.Get("/", getTasks)
	users.Get("/:id", getTaskById)
	users.Post("/", createTask)
	users.Put("/:id", updateTask)
	users.Delete("/:id", deleteTask)
}

func createTask(c *fiber.Ctx) error {

   var	task models.Task // creamos una variable de tipo User
	if err := c.BodyParser(&task); //? &task es un puntero a la variable task en este caso se le pasa el valor del body y no hay necesidad de usar task = c.BodyParser() y mas bien verificamos si hay un error en el body y lo validamos con el if

	err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "An error has occurred with data,Invalid body"},
		)
	}
	
	// VALIDATE TASK
	if task.Title == "" || task.Title == " " {
		return c.Status(400).JSON(fiber.Map{
			"message": "title is required"},
		)
	}
	if task.Description== "" || task.Description== " " {
		return c.Status(400).JSON(fiber.Map{
			"message": "description is required"},
		)
	}
	
	result:= db.DB.Create(&task) // aqui se crea el usuario en la BDD postgreSQL 
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
		"message": task,
	
	})
}

func getTasks(c *fiber.Ctx) error {
	var tasks []models.Task // creamos una variable de tipo User 
	db.DB.Find(&tasks) // aqui se obtienen todos los usuarios de la BDD postgreSQL y se almacenan en la variable users
   return c.Status(200).JSON(fiber.Map{
	   "tasks": tasks,
   })
}


// GET USER BY ID
func getTaskById(c *fiber.Ctx) error {
	var taskById models.Task
	 id := c.Params("id")
	 
	 db.DB.Find(&taskById, id)
	 return c.Status(200).JSON(fiber.Map{
		 "user": taskById,
	 })
 }
 
 func updateTask(c *fiber.Ctx) error {
 
	 var taskUpdate models.Task
	 var task models.Task
 
	 id := c.Params("id")
 
	 if err := c.BodyParser(&task);	err != nil {
		 return c.Status(400).JSON(fiber.Map{
			 "message": "An error has occurred with data,Invalid body"},
		 )
	 }
	// VALIDATE TASK
	if task.Title == "" || task.Title == " " {
		return c.Status(400).JSON(fiber.Map{
			"message": "title is required"},
		)
	}
	if task.Description== "" || task.Description== " " {
		return c.Status(400).JSON(fiber.Map{
			"message": "description is required"},
		)
	}
 
	 db.DB.First(&taskUpdate, id)
	  
	  taskUpdate.Title= task.Title
	  taskUpdate.Description= task.Description

	 result:=db.DB.Save(&taskUpdate)
	 err := result.Error
	 var messageErr string
	 if err != nil {
		 //?: reflect permite acceder a los campos de una estructura y obtener su valor
		 code:= reflect.ValueOf(err).Elem().FieldByName("Code").String()  // aqui se obtiene el codigo de error de la BDD
		 detail:= reflect.ValueOf(err).Elem().FieldByName("Detail").String() // aqui se obtiene el detalle del error de la BDD
		 td:= strings.Split(detail,"=")
		 campo:=strings.Split(td[1], " ")
	 
		 if code == "23505" {
			 messageErr= "Ya existe una tarea con este título "+campo[0]+" intente con otro"
		 }
 
		 return c.Status(400).JSON(fiber.Map{
			 "messageError":messageErr, 
		 },
		 )
	 }
 
	 rowsAffected := result.RowsAffected
	 var message string
	 if rowsAffected == 0 {
		 message = "The task does not exist "
	 } else {
		 message = "Task updated"
	 }
	 return c.Status(200).JSON(fiber.Map{
		 "message": message,
		 "data": taskUpdate,
	 })
 }
 
 func deleteTask(c *fiber.Ctx) error {
	 var taskDelete models.Task
	 id := c.Params("id")
	 // result:=db.DB.Delete(&taskDelete, id)  //?Delete es un metodo que elimina un registro de la BDD pero solo visualmente(temporal), no se elimina en la BDD
	 result:= db.DB.Unscoped().Delete(&taskDelete, id) //? Unscoped es un metodo que elimina un registro de la BDD de forma permanente
	 
	 err := result.Error
	 if err != nil {
		 return c.Status(400).JSON(fiber.Map{
			 "message": " An error has occurred "},
		 )
	 }
 
	 rowsAffected := result.RowsAffected
	 var message string
	 if rowsAffected == 0 {
		 message = "The task does not exist or has already been deleted"
	 } else {
		 message = "Task deleted"
	 }
	 return c.Status(200).JSON(fiber.Map{
		 "message": message,
		 
	 })
 }
 