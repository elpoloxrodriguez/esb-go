package models

import (
	"encoding/json"

	"github.com/elpoloxrodriguez/esb-inea/sys"

	"gopkg.in/mgo.v2/bson"
)

type Ciudad struct {
	Capital string `json:"capital"`
	Nombre  string `json:"nombre"`
}

type Municipio struct {
	Nombre    string   `json:"nombre"`
	Parroquia []string `json:"parroquia"`
}

type Estados struct {
	Nombre    string      `json:"nombre"`
	Codigo    string      `json:"codigo"`
	Ciudad    []Ciudad    `json:"ciudad"`
	Municipio []Municipio `json:"municipio"`
}

//Listar el sistema de Ciudades
func (u *Ciudad) ListaCiudades() (j []byte, err error) {
	var lstCiudades []Estados
	c := sys.MGOSession.DB(sys.CBASE).C("estados")
	err = c.Find(bson.M{}).Select(bson.M{}).All(&lstCiudades)
	j, _ = json.Marshal(lstCiudades)
	return
}
