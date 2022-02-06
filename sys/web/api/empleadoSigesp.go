package api

import (
	"net/http"

	models "github.com/elpoloxrodriguez/esb-inea/models/posgresql"
)

type EmpleadoSIGESP struct{}

//Listar Empleado SIGESP PoatgreSql
func (u *EmpleadoSIGESP) MostrarDatoEmpleadoSIGESP(w http.ResponseWriter, r *http.Request) {
	var empleadoSIGESP models.BuscarEmpleadoSigesp
	Cabecera(w, r)
	j, _ := empleadoSIGESP.MostrarEmpleadoSIGESP()
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
