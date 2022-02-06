package models

import (
	"encoding/json"
	"strconv"

	"github.com/elpoloxrodriguez/esb-inea/sys"
)

type Aplicacion struct {
	ID        int    `json:"id,omitempty"`
	Nombre    string `json:"nombre,omitempty"`
	Direccion string `json:"direccion,omitempty"`
	Email     string `json:"email,omitempty"`
	Lema      string `json:"lema,omitempty"`
	Iniciales string `json:"iniciales,omitempty"`
	Slogan    string `json:"slogan,omitempty"`
	Creador   string `json:"creador,omitempty"`
	Fundada   string `json:"fundada,omitempty"`
	Sitioweb  string `json:"sitioweb,omitempty"`
	Logo      string `json:"logo,omitempty"`
	Status    int64  `json:"status,omitempty"`
}

//Listar el sistema de mysql
func (u *Aplicacion) MostrarAplicacion() (jSon []byte, err error) {
	rows, err := sys.MysqlFullText.Query("SELECT * FROM aplicacion")
	if err != nil {
		panic(err.Error())
	}
	var lst []interface{}
	for rows.Next() {
		var mysqldata Aplicacion
		var status int
		var id, nombre, direccion, email, lema, iniciales, slogan, creador, fundada, sitioweb, logo string
		err = rows.Scan(&id, &nombre, &direccion, &email, &lema, &iniciales, &slogan, &creador, &fundada, &sitioweb, &logo, &status)
		if err != nil {
			panic(err.Error())
		}
		mysqldata.ID, _ = strconv.Atoi(id)
		mysqldata.Nombre = nombre
		mysqldata.Direccion = direccion
		mysqldata.Email = email
		mysqldata.Lema = lema
		mysqldata.Iniciales = iniciales
		mysqldata.Slogan = slogan
		mysqldata.Creador = creador
		mysqldata.Fundada = fundada
		mysqldata.Sitioweb = sitioweb
		mysqldata.Logo = logo
		mysqldata.Status = int64(status)
		lst = append(lst, mysqldata)
	}
	jSon, err = json.Marshal(lst)
	return
}
