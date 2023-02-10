package main

import (
	"fmt"
    "log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/ivanes285/GoPostgresAPI/db"
	"github.com/ivanes285/GoPostgresAPI/models"
	"github.com/ivanes285/GoPostgresAPI/routes/v1"
	"github.com/joho/godotenv"
)

func main() {
	err := run()    // ejecutamos la función run y la asignamos a la variable err
	if err != nil { // si err es diferente de nil entonces mostramos el error en la consola
		panic(err)
	}

}

func run() error {

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

	// MIDDLEWARES
	app.Use(logger.New())         // logger permite mostrar en la consola las peticiones que se hacen a la API
	app.Use(recover.New())        // recover permite mostrar en la consola los errores y no se caiga el servidor en el caso que se ejecute un panic
	app.Use(cors.New(cors.Config{ // Configuración de CORS para permitir el acceso a la API desde cualquier origen
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection,Access-Control-Allow-Origin",
		AllowCredentials: true,
	}))

	//Routes
	routes.AddUsersGroup(app)
	routes.AddTasksGroup(app)
	

	// PORT
	PORT := os.Getenv("PORT") // Obtenemos el puerto de la variable de entorno PORT
	if PORT == "" {
		PORT = "4000"
	}

	app.Listen(":" + PORT) // Iniciamos el servidor y si hay un error lo mostramos en la consola
	fmt.Println("Server is running on port", PORT)

	return nil // retornamos nil porque no hay error

}
