package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Respuesta struct {
	Estado        int         `json:"estado"`
	Dato          interface{} `json:"dato"`
	Mensaje       string      `json:"mensaje"`
	contenType    string
	responseWrite http.ResponseWriter
}

func CrearDefectoRespuesta(rw http.ResponseWriter) Respuesta {
	return Respuesta{
		Estado:        http.StatusOK,
		responseWrite: rw,
		contenType:    "application/json",
	}
}

func (resp *Respuesta) NoEncontrado() {
	resp.Estado = http.StatusNotFound
	resp.Mensaje = "NO se ha encontrado la fuente"
}

func (resp *Respuesta) Enviar() {
	resp.responseWrite.Header().Set("Content-Type", resp.contenType)
	resp.responseWrite.WriteHeader(resp.Estado)

	salida, _ := json.Marshal(&resp)
	fmt.Fprintln(resp.responseWrite, string(salida))
}

func EnviarDato(rw http.ResponseWriter, dato interface{}, mensaje string) {
	Respuesta := CrearDefectoRespuesta(rw)
	Respuesta.Dato = dato
	Respuesta.Mensaje = mensaje
	Respuesta.Enviar()
}

func EnviarNOEncontrado(rw http.ResponseWriter) {
	respuesta := CrearDefectoRespuesta(rw)
	respuesta.NoEncontrado()
	respuesta.Enviar()
}

func (resp *Respuesta) NoProcesaEntidad() {
	resp.Estado = http.StatusUnprocessableEntity
	resp.Mensaje = "NO se han podido procesar las entidades"
}

func EnviarNoProcesaEntidad(rw http.ResponseWriter) {
	respuesta := CrearDefectoRespuesta(rw)
	respuesta.NoProcesaEntidad()
	respuesta.Enviar()
}
