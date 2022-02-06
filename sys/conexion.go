//configuraciones del sistema
package sys

import (
	"database/sql"
	"fmt"

	mgo "gopkg.in/mgo.v2"

	"github.com/elpoloxrodriguez/esb-inea/util"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

//MongoDBConexion Conexion a Mongo DB
func MongoDBConexion(mapa map[string]CadenaDeConexion) {
	c := mapa["mongodb"]
	MGOSession, Error = mgo.Dial(c.Host + ":27017")
	fmt.Println("[MongoDB: ", c.Host, " ✅]")
	util.Error(Error)
	// if MGOSession.Ping() != nil {
	// 	fmt.Println("[MongoDB: ", c.Host, " ❌] ", MGOSession.Ping())
	// } else {
	// 	fmt.Println("[MongoDB: ", c.Host, " ✅]")
	// }
}

//ConexionPACE
func ConexionPosgreSql(mapa map[string]CadenaDeConexion) {
	c := mapa["postgresql"]
	cadena := "user=" + c.Usuario + " dbname=" + c.Basedatos + " password=" + c.Clave + " host=" + c.Host + " sslmode=disable"
	PostgreSQL, _ = sql.Open("postgres", cadena)
	if PostgreSQL.Ping() != nil {
		fmt.Println("[PostgreSql: ", c.Host, " ❌] ", PostgreSQL.Ping())
	} else {
		fmt.Println("")
		fmt.Println("..........................................................")
		fmt.Println("........ Inciando la carga de las Bases de Datos  ........")
		fmt.Println("..........................................................")
		fmt.Println("")

		fmt.Println("[PostgreSql: ", c.Host, " ✅]")
	}
}

//ConexionMYSQL
func ConexionMYSQL(mapa map[string]CadenaDeConexion) {
	c := mapa["mysql"]
	cadena := c.Usuario + ":" + c.Clave + "@tcp(" + c.Host + ":3306)/" + c.Basedatos
	MysqlFullText, _ = sql.Open("mysql", cadena)
	if MysqlFullText.Ping() != nil {
		fmt.Println("[MySql: ❌] ", MysqlFullText.Ping())
	} else {
		fmt.Println("[MySql: ", c.Host, " ✅]")
	}
	return
}
