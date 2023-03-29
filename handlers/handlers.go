package handlers

import (
	"encoding/json"
	"fmt"
	"goapirest/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const USER_ELIMINADO = "Se ha Eliminado el Usuario correctamente"
const USER_AGREGADO = "Se ha Creado el Usuario correctamente"
const USER_MODIFICADO = "Se ha Modificado el Usuario correctamente"
const USER_LISTADO = "Se han listado Los Usuarios correctamente"
const USER_BUSCADO = "Se ha encontrado el Usuario correctamente"

func GetUsuarios(rw http.ResponseWriter, r *http.Request) {
	if usuarios, error := models.ListarUsuarios(); error != nil {
		models.EnviarNOEncontrado(rw)
	} else {
		models.EnviarDato(rw, usuarios, USER_LISTADO)
	}
}
func GetUsuario(rw http.ResponseWriter, r *http.Request) {
	if usuario, error := obtenerUsuarioPeticion(r); error != nil {
		models.EnviarNOEncontrado(rw)
	} else {
		models.EnviarDato(rw, usuario, USER_BUSCADO)
	}
}

func CreateUsuario(rw http.ResponseWriter, r *http.Request) {

	usuario := models.Usuario{}
	decodificar := json.NewDecoder(r.Body)

	if error := decodificar.Decode(&usuario); error != nil {
		models.EnviarNoProcesaEntidad(rw)
		fmt.Fprintln(rw, "Error")
	} else {
		usuario.Guardar()
		models.EnviarDato(rw, usuario, USER_AGREGADO)
		fmt.Fprintln(rw, "Proceso realizado correctamente")
	}
}

func UpdateUsuario(rw http.ResponseWriter, r *http.Request) {

	//Obtener registro
	var usId int64

	if usuario, error := obtenerUsuarioPeticion(r); error != nil {
		fmt.Println("error en obtener usuario")
		models.EnviarNOEncontrado(rw)
		fmt.Fprintln(rw, "Error")
	} else {
		usId = usuario.Id
	}

	usuario := models.Usuario{}
	decodificar := json.NewDecoder(r.Body)

	if error := decodificar.Decode(&usuario); error != nil {
		fmt.Println("error en decodificaciones")
		models.EnviarNoProcesaEntidad(rw)
		fmt.Fprintln(rw, "Error")
	} else {
		usuario.Id = usId
		usuario.Guardar()
		models.EnviarDato(rw, usuario, USER_MODIFICADO)
		fmt.Fprintln(rw, "Proceso realizado correctamente")
	}
}

func DeleteUsuario(rw http.ResponseWriter, r *http.Request) {
	if usuario, error := obtenerUsuarioPeticion(r); error != nil {
		models.EnviarNOEncontrado(rw)
	} else {
		usuario.EliminarUsuario()
		models.EnviarDato(rw, usuario, USER_ELIMINADO)
	}
}

func obtenerUsuarioPeticion(r *http.Request) (models.Usuario, error) {
	//obtener el id
	v := mux.Vars(r)
	usuarioId, _ := strconv.Atoi(v["id"])
	if usuario, error := models.ObtenerUsuario(usuarioId); error != nil {
		return *usuario, error
	} else {
		return *usuario, nil
	}
}
