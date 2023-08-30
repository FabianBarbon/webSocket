package db

import (
	"database/sql"
	"fmt"

	// Import the Azure AD driver module (also imports the regular driver package)
	_ "github.com/go-sql-driver/mysql"
)

var dbHost = "telemetria.cchfdbjtfv0t.us-east-2.rds.amazonaws.com"
var dbPort = "3306"
var dbUser = "telemetria_dev"
var dbPass = "Qd!897BLOMtZdTrAmpu*!sZl"
var dbName = "xtamtelemetria_Dev"
var UrlDB = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

// guarda la conexion
var db *sql.DB

// conexion
func Connect() {
	connection, err := sql.Open("mysql", UrlDB)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexion a la base de datos exitosa!")
	db = connection
}

// cerrar conexion
func Close() {
	db.Close()
}

// revisar conexion
func Ping() {

	if err := db.Ping(); err != nil {
		panic(err)
	}
}
