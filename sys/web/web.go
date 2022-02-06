package web

//Copyright Carlos Peña
//Controlador del MiddleWare
import (
	"fmt"
	"net/http"

	"github.com/elpoloxrodriguez/esb-inea/sys/web/api"
	"github.com/gorilla/mux"
)

//Variables de Control
var (
	Enrutador   = mux.NewRouter()
	wUsuario    api.WUsuario
	lstCiudades api.ListaCiudades
)

//Cargar los diferentes modulos del sistema
func Cargar() {
	CargarModulosLogin()
	CargarModulosDevel()
	Principal()
}

//CargarModulosWebSite Cargador de modulos web
func CargarModulosLogin() {
	fmt.Println("[Rutas de Produccion] ✅")
	Enrutador.HandleFunc("/inea/api/web/loginW", wUsuario.LoginW).Methods("POST")
	Enrutador.HandleFunc("/inea/api/web/cambiarclave", wUsuario.CambiarClave).Methods("POST")
	Enrutador.HandleFunc("/inea/recuperarclave", wUsuario.RecuperarW).Methods("POST")
}

func CargarModulosDevel() {
	fmt.Println("[Rutas de Desarrollo] ✅")
	Enrutador.HandleFunc("/devel/api/wusuario/listar", wUsuario.Listar).Methods("GET")
	Enrutador.HandleFunc("/devel/api/ciudades/listado", lstCiudades.Listar).Methods("GET")
}

//Principal Página inicial del sistema o bienvenida
func Principal() {
	prefix := http.StripPrefix("/", http.FileServer(http.Dir("public_web/dist")))
	Enrutador.PathPrefix("/").Handler(prefix)
}
