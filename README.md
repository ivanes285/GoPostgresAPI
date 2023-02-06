# GoPostgresAPI
Esta aplicacion es una backend realizada con el framework fiber de GO.

#SERVER (backend)


### Paquetes
- Instalamos air que nos permitira aplicar y reiniciar de forma automatica al servidor de go
```
go install github.com/cosmtrek/air@latest
```
-Una vez instalado ejecutamos los comandos que nos permitiran establecer un archivo de configuracion y el otro ejecutar nuestro servidor respectivamente
```
ait init
```
```
air
```

### DATABASE
- Utilizamos el paquete gorm y tambien el driver de la respectiva bdd que vamos a usar en este caso postgreSQL para GO
```
go get -u gorm.io/gorm
```
```
 go get -u gorm.io/driver/postgres 
```

### DOKER
- Iniciamos docker y posterior a eso ejecutamos en el terminal el comando para crear un container de postgres con docker
- Instalamos primero la imagen de postgres
 ```
 docker pull postgres
 ```
- Creamos un contenedor en este caso con el nombre mypostgress , configuramos el puerto en el que quiero que se ejecute, las variables de entorno para la password y el user y por ulltimo le decismo -d(ditach) para que el contenedor se quede ejecutando en segundo plano
```
docker run --name mypostgres -p 5432:5432 -e POSTGRES_PASSWORD=sistemas -e POSTGRES_USER=ivan -d postgres
```
- Con el siguiente comando podremos visualizar nuestros contenedores 
```
docker ps 
```
