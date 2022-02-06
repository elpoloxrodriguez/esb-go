package api

import (
	"net/http"

	models "github.com/elpoloxrodriguez/esb-inea/models/mongoDB"
)

type ListaCiudades struct{}

//Listar Ciudades
func (u *ListaCiudades) Listar(w http.ResponseWriter, r *http.Request) {
	var ciudades models.Ciudad
	Cabecera(w, r)
	j, _ := ciudades.ListaCiudades()
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
