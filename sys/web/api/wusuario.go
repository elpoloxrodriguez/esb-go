package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/elpoloxrodriguez/esb-inea/sys/seguridad"
	"github.com/elpoloxrodriguez/esb-inea/util"
)

type WUsuario struct{}

//Crear Usuario del sistema
func (u *WUsuario) Crear(w http.ResponseWriter, r *http.Request) {
	var usuario seguridad.Usuario
	var m util.Mensajes
	Cabecera(w, r)
	ip := strings.Split(r.RemoteAddr, ":")
	usuario.FirmaDigital.DireccionIP = ip[0]
	usuario.FirmaDigital.Tiempo = time.Now()
	usuario.FechaCreacion = time.Now()

	e := json.NewDecoder(r.Body).Decode(&usuario)
	util.Error(e)
	// if ip[0] == "192.168.6.45" {

	e = usuario.Salvar()
	if e != nil {
		w.WriteHeader(http.StatusForbidden)
		m.Msj = e.Error()
		m.Tipo = 1
		fmt.Println("Err: ", e.Error())
	} else {
		w.WriteHeader(http.StatusOK)
		m.Msj = "Usuario creado"
		m.Tipo = 0
	}
	// } else {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	m.Msj = "El equipo donde no esta autorizado"
	// 	m.Tipo = 2
	// }

	m.Fecha = time.Now()
	j, _ := json.Marshal(m)
	w.Write(j)
}

//Clave de tipos
type Clave struct {
	ID        string `json:"id"`
	Cedula    string `json:"cedula"`
	Login     string `json:"login"`
	Clave     string `json:"clave"`
	Nueva     string `json:"nueva"`
	Repetir   string `json:"repetir"`
	Correo    string `json:"correo"`
	Coleccion string `json:"coleccion"`
}

//LoginW conexion para solicitud de token
func (u *WUsuario) LoginW(w http.ResponseWriter, r *http.Request) {
	var usuario seguridad.WUsuario
	// var traza fanb.Traza
	CabeceraW(w, r)
	e := json.NewDecoder(r.Body).Decode(&usuario)
	fmt.Println(usuario)
	util.Error(e)
	//fmt.Println("login... ", usuario.Cedula, usuario.Nombre, util.GenerarHash256([]byte(usuario.Clave)))
	err := usuario.Validar(usuario.Cedula, util.GenerarHash256([]byte(usuario.Clave)))
	if err != nil {

		w.Header().Set("Content-Type", "application/text")
		fmt.Println("Error en la conexion del usuario")
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Usuario y clave no validas")

	} else {

		if usuario.Cedula != "" && usuario.Componente != "" {
			usuario.Clave = ""
			token := seguridad.WGenerarJWT(usuario)
			result := seguridad.RespuestaToken{Token: token}
			j, e := json.Marshal(result)
			util.Error(e)
			w.WriteHeader(http.StatusOK)
			w.Write(j)
		}
	}

}
