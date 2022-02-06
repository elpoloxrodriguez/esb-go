package api

import (
	"net/http"

	models "github.com/elpoloxrodriguez/esb-inea/models/mysql"
)

type Aplicacion struct{}

//Listar Aplicacion mysql
func (u *Aplicacion) MostrarDatos(w http.ResponseWriter, r *http.Request) {
	var aplicacion models.Aplicacion
	Cabecera(w, r)
	j, _ := aplicacion.MostrarAplicacion()
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
