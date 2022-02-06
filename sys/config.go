package sys

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/elpoloxrodriguez/esb-inea/util"
	mgo "gopkg.in/mgo.v2"
)

type config struct{}

//Variables del modelo
var (
	MySQL         bool = false
	MongoDB       bool = false
	SQLServer     bool = false
	Oracle        bool = false
	BaseDeDatos   BaseDatos
	MGOSession    *mgo.Session
	PostgreSQL    *sql.DB
	MysqlFullText *sql.DB
	Error         error
	Version       string = "V.2.2.2"
)

//Constantes del sistema
const (
	ACTIVAR_CONEXION_REMOTA       bool   = true
	DESACTIVAR_CONEXION_REMOTA    bool   = false
	ACTIVAR_LOG_REGISTRO          bool   = true
	DESACTIVAR_LOG_REGISTRO       bool   = false
	ACTIVAR_ROLES                 bool   = true
	DESACTIVAR_ROLES              bool   = false
	ACTIVAR_LIMITE_DE_CONSULTA    bool   = true
	DESACTIVAR_LIMITE_DE_CONSULTA bool   = false
	PUERTO                        string = "80"
	PUERTO_STANDAR                string = "8080"
	PUERTO_SSL                    string = "2608"
	PUERTO_SSL_STANDAR            string = "443"
	CODIFCACION_DE_ARCHIVOS       string = "UTF-8"
	MAXIMO_LIMITE_DE_USUARIO      int    = 100
	MAXIMO_LIMITE_DE_CONSULTAS    int    = 10
)

//BaseDatos Estructuras
type BaseDatos struct {
	CadenaDeConexion map[string]CadenaDeConexion
}

//CadenaDeConexion Conexion de datos
type CadenaDeConexion struct {
	Driver    string
	Usuario   string
	Basedatos string
	Clave     string
	Host      string
	Puerto    string
	StrUrl    string
}

//Conexiones 0:PostgreSQL, 1:MySQL, 2:MongoDB
var Conexiones []CadenaDeConexion

//init Inicio y control
func init() {
	var a util.Archivo
	a.NombreDelArchivo = "sys/config_dev.json"
	data, _ := a.LeerTodo()
	e := json.Unmarshal(data, &Conexiones)
	for _, valor := range Conexiones {
		switch valor.Driver {
		case "postgresql":
			cad := make(map[string]CadenaDeConexion)
			cad["postgresql"] = CadenaDeConexion{
				Driver:    valor.Driver,
				Usuario:   valor.Usuario,
				Basedatos: valor.Basedatos,
				Clave:     valor.Clave,
				Host:      valor.Host,
				Puerto:    valor.Puerto,
			}
			ConexionPosgreSql(cad)
		case "mysql":
			MySQL = true
			cad := make(map[string]CadenaDeConexion)
			cad["mysql"] = CadenaDeConexion{
				Driver:    valor.Driver,
				Usuario:   valor.Usuario,
				Basedatos: valor.Basedatos,
				Clave:     valor.Clave,
				Host:      valor.Host,
				Puerto:    valor.Puerto,
			}
			ConexionMYSQL(cad)
		case "mongodb":
			MongoDB = true
			cad := make(map[string]CadenaDeConexion)
			cad["mongodb"] = CadenaDeConexion{
				Driver:    valor.Driver,
				Usuario:   valor.Usuario,
				Basedatos: valor.Basedatos,
				Clave:     valor.Clave,
				Host:      valor.Host,
				Puerto:    valor.Puerto,
			}
			MongoDBConexion(cad)
		}
	}
	if e != nil {
		fmt.Println("\n Utilidad Error: ", e.Error())
	}
}
