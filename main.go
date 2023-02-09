package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ivanes285/GoPostgresAPI/db"
	"github.com/ivanes285/GoPostgresAPI/models"
	"github.com/ivanes285/GoPostgresAPI/routes/v1"
	"github.com/joho/godotenv"
)

func main() {
	err := run()    // ejecutamos la función run y la asignamos a la variable err
	if err != nil { // si err es diferente de nil entonces mostramos el error en la consola
		log.Fatal(err)
	}

	
}

func run () error{

	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}


	//?Database connection
	db.DBConnection()

	//!! AutoMigrate crea la tabla en la base de datos si no existe y toma como referencia la estructura del modelo Task y                       crea las columnas de la tabla de acuerdo a los campos de la estructura Task , {} significa que no hay ninguna configuración adicional
    db.DB.AutoMigrate(models.Task{}) 
    db.DB.AutoMigrate(models.User{})
	

    //? Inicializamos la aplicación
	app := fiber.New()
	

	//Routes
	app.Get("/", routes.Home)
    routes.AddUsersGroup(app)


	app.Listen(":3000")


	return nil // retornamos nil porque no hay error

}