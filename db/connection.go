package db

import (
	"log"
	// "os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB     // variable global ya que sino no se puede acceder desde otros paquetes a los metodos de la BDD postgreSQL

func DBConnection()  {
	//data source name(dsn) referencia de la base de datos a la cual nos vamos a conectar
	dsn:= "host=dbpostgres.cygxjxa4hy63.us-east-1.rds.amazonaws.com user=ivan password=sistemas dbname=postgresgo port=5432" // DSN es una varriable de entorno y significa data source name
	var err error
	DB, err = gorm.Open(postgres.Open(dsn),  &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}else{
		log.Println("Database is connected")
	}

}