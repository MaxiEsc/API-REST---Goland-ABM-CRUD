package models

import (
	"goapirest/db"
)

type Usuario struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Usuarios []Usuario

//Scheems

const UsuarioScheme string = `CREATE TABLE Usuario (
	id INT(10) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(50) NOT NULL,
	password VARCHAR(150) NOT NULL,
	email VARCHAR(150) NOT NULL,
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)`

// Constructor de usuario
func NuevoUsuario(username, password, email string) *Usuario {
	user := &Usuario{
		Username: username,
		Password: password,
		Email:    email,
	}

	return user
}

//Crear e insertar a la base de datos

/*
Ahora es cuando nos sirve el polimorfismo ya que gracias a esa reescritua en el archivo database.go
se puede permitir el uso del metodo en cuestion
*/

func CrearUsuario(username, password, email string) *Usuario {
	elUsuario := NuevoUsuario(username, password, email)
	elUsuario.agregar()
	return elUsuario
}

// Insertar un registro
func (user *Usuario) agregar() {
	sql := "INSERT usuario SET username=?, password=?, email=?"
	result, _ := db.Exec(sql, user.Username, user.Password, user.Email)
	user.Id, _ = result.LastInsertId()
}

// Listar todos los registros
func ListarUsuarios() (Usuarios, error) {

	sql := "SELECT id, username, password, email FROM usuario"
	usuarios := Usuarios{}
	filas, error := db.Query(sql)
	//Hacemos una especie de while que recorra la lista obtenida en la variable filas
	for filas.Next() {
		//Creamos una variable de tipo Usuario
		usuario := Usuario{}
		//Con Scan se hace una lectura de las direcciones contenidas en las filas.
		filas.Scan(&usuario.Id, &usuario.Username, &usuario.Password, &usuario.Email)
		//Mas tarde vemos al slicen de usuarios y le asignamos los datos dentro de su recorrido
		//De manera que añadimos al slicen con append lo que se encuentra en bucle de filas mediante la direccion
		//con ya los datos de usuarios iterados asi mismo entonces se le asignan al slicen.
		usuarios = append(usuarios, usuario)
	}

	return usuarios, error
}

// Obtener un registro
func ObtenerUsuario(id int) (*Usuario, error) {
	usuario := NuevoUsuario("", "", "")

	sql := "SELECT id, username,password, email FROM usuario WHERE id=?"

	if filas, error := db.Query(sql, id); error != nil {
		return nil, error
	} else {
		//Aqui tenemos que usar la variable usuarios y cuando se escanee la variable filas se le asigna
		//la direccion de la misma para copiar lo que tenga filas en la direccion de usuario
		for filas.Next() {
			filas.Scan(&usuario.Id, &usuario.Username, &usuario.Password, &usuario.Email)
		}

		return usuario, nil
	}
}

// Actualizar registro
func (usuario *Usuario) ActualizarUsuario() {

	sql := "UPDATE usuario SET username=?, password=?, email=? WHERE id=?"
	db.Exec(sql, usuario.Username, usuario.Password, usuario.Email, usuario.Id)
}

// Guardar o editar registro
func (usuario *Usuario) Guardar() {
	if usuario.Id == 0 {
		usuario.agregar()
	} else {
		usuario.ActualizarUsuario()
	}
}

// Eliminar un registro
func (Usuario *Usuario) EliminarUsuario() {
	sql := "DELETE FROM usuario WHERE id=?"
	db.Exec(sql, Usuario.Id)
}
