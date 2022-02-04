package web

//Copyright Carlos Peña
//Controlador del MiddleWare
import (
	"net/http"

	"github.com/elpoloxrodriguez/esb-inea/sys/web/api"
	"github.com/gorilla/mux"
)

//Variables de Control
var (
	Enrutador = mux.NewRouter()
	wUsuario  api.WUsuario
)

//Cargar los diferentes modulos del sistema
func Cargar() {
	CargarModulosWebSite()
	Principal()
}

//CargarModulosWebSite Cargador de modulos web
func CargarModulosWebSite() {
	Enrutador.HandleFunc("/inea/api/web/loginW", wUsuario.LoginW).Methods("POST")
	Enrutador.HandleFunc("/inea/api/web/cambiarclave", wUsuario.CambiarClave).Methods("POST")
	Enrutador.HandleFunc("/inea/recuperarclave", wUsuario.RecuperarW).Methods("POST")

}

//Principal Página inicial del sistema o bienvenida
func Principal() {
	prefix := http.StripPrefix("/", http.FileServer(http.Dir("public_web/dist")))
	Enrutador.PathPrefix("/").Handler(prefix)
}
