package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/*
Con el comando :
 "go get -u github.com/go-sql-driver/mysql"
 Instalamos el driver de para la conexion de MySql con Goland
*/

// Se construye con esta estructura el link de la conexion
// Username:password@tcp(Localhost3306)/database
const url = "root:12345678@tcp(localhost:3306)/goweb_db"

// Variable Global de un tipo especifico de dato
// Nos permite Guardar la conexion cuando de conecta valga la redundancia
var db *sql.DB

// Funcion para realizar la conexion
func ConexionDB() {
	conexion, error := sql.Open("mysql", url)

	if error != nil {
		panic(error)
	}

	fmt.Println("Conexion Exitosa")
	db = conexion
}

// Funcion para cerrar la conexión
func CerrarConexion() {
	db.Close()
}

// Verifica la conexion
func VerificaConexion() {

	//Recordar que en Goland podemos crear variables dentro del IF
	if err := db.Ping(); err != nil {
		panic(err)
	}

}

// Verificar si una tabla existe
func ExisteTabla(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	filas, error := db.Query(sql)
	if error != nil {
		fmt.Println("Error: ", error)
	}

	return filas.Next()
}

// Crea una tabla Usuario
func CreaTabla(schema string, name string) {

	if !ExisteTabla(name) {
		_, error := db.Exec(schema)

		if error != nil {
			fmt.Printf("ERROR: %s", error)
		}
	}
}

//Polimorfismo del metodo
/*
 elminamos "(db *DB)" para usuarlo como una funcion asi lo modificamos a placer
*/
func Exec(query string, args ...any) (sql.Result, error) {
	ConexionDB()
	result, error := db.Exec(query, args...)
	CerrarConexion()
	if error != nil {
		fmt.Println(error)
	}

	return result, error
}

/*
Repetimos el polimorfismo con el metodo Query

Polimorfismo de Query
*/
func Query(query string, args ...any) (*sql.Rows, error) {
	ConexionDB()
	filas, error := db.Query(query, args...)
	CerrarConexion()
	if error != nil {
		fmt.Println(error)
	}

	return filas, error
}

//Reiniciar registro de una tabla

func CortarTabla(tablename string) {
	sql := fmt.Sprintf("TRUNCATE %s", tablename)
	Exec(sql)
}
