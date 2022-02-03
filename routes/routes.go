package routes

//Controlador del MiddleWare
import (
	"net/http"

	"github.com/gorilla/mux"
)

//Variables de Control
var (
	Enrutador = mux.NewRouter()
)

//Cargar los diferentes modulos del sistema
func Cargar() {
	Principal()
}

//Principal Página inicial del sistema o bienvenida
func Principal() {
	prefix := http.StripPrefix("/", http.FileServer(http.Dir("public_web/dist")))
	Enrutador.PathPrefix("/").Handler(prefix)
}
