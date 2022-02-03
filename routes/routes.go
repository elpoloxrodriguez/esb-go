package routes

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
	Principal()
	Login()
}

func Login() {
	Enrutador.HandleFunc("/ipsfa/api/web/loginW", wUsuario.LoginW).Methods("POST")
}

//Principal Página inicial del sistema o bienvenida
func Principal() {
	prefix := http.StripPrefix("/", http.FileServer(http.Dir("public_web/dist")))
	Enrutador.PathPrefix("/").Handler(prefix)
}
