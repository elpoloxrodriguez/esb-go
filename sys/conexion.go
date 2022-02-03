//configuraciones del sistema
package sys

import (
	"database/sql"
	"fmt"

	mgo "gopkg.in/mgo.v2"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

//MongoDBConexion Conexion a Mongo DB
func MongoDBConexion(mapa map[string]CadenaDeConexion) {
	c := mapa["mongodb"]
	MGOSession, Error = mgo.Dial(c.Host + ":27017")
	fmt.Println("Cargando Conexión Con MongoDB...")
	// util.Error(Error)
	// if e != nil {
	// 	fmt.Println("\n Utilidad Error: ", e.Error())
	// }
}

//ConexionPACE
func ConexionPosgreSql(mapa map[string]CadenaDeConexion) {
	c := mapa["postgresql"]
	cadena := "user=" + c.Usuario + " dbname=" + c.Basedatos + " password=" + c.Clave + " host=" + c.Host + " sslmode=disable"
	PostgreSQLPACE, _ = sql.Open("postgres", cadena)
	if PostgreSQLPACE.Ping() != nil {
		fmt.Println("[PostgreSql: ", c.Host, " Error...] ", PostgreSQLPACE.Ping())
	} else {
		fmt.Println("[PostgreSql: ", c.Host, " OK...]")
	}
}

//ConexionMYSQL
func ConexionMYSQL(mapa map[string]CadenaDeConexion) {
	c := mapa["mysql"]
	cadena := c.Usuario + ":" + c.Clave + "@tcp(" + c.Host + ":3306)/sssifanb"
	MysqlFullText, _ = sql.Open("mysql", cadena)
	if MysqlFullText.Ping() != nil {
		fmt.Println("[mysql FULLTEXT: Error...] ", MysqlFullText.Ping())
	} else {
		fmt.Println("[MySql: ", c.Host, " OK...]")
	}
	return
}
