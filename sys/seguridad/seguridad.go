// se refiere a la confianza en algo.
package seguridad

import (
	"crypto/rsa"
	"io/ioutil"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/elpoloxrodriguez/esb-inea/util"
)

//Constantes Generales
const (
	ENCRIPTAMIENTO             = "md5"
	ACTIVARLIMITECONEXIONES    = true
	DESACTIVARLIMITECONEXIONES = false
)

//Variables de Seguridad
var (
	LlavePrivada *rsa.PrivateKey
	LlavePublica *rsa.PublicKey
	LlaveJWT     string
)

//init Función inicial del sistema
func init() {
	bytePrivados, err := ioutil.ReadFile("./certificados/jwt/private.rsa")
	util.Fatal(err)
	LlavePrivada, err = jwt.ParseRSAPrivateKeyFromPEM(bytePrivados)
	bytePublicos, err := ioutil.ReadFile("./certificados/jwt/public.rsa.pub")
	util.Fatal(err)
	LlavePublica, err = jwt.ParseRSAPublicKeyFromPEM(bytePublicos)
}

//GenerarJWT Json Web Token
func GenerarJWT(u Usuario) string {
	peticion := Reclamaciones{
		Usuario: u,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 259200).Unix(),
			Issuer:    "Conexion Bus Empresarial INEA",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, peticion)
	rs, e := token.SignedString(LlavePrivada)
	util.Fatal(e)
	return rs
}

//WGenerarJWT Json Web Token
func WGenerarJWT(u WUsuario) string {
	peticion := WReclamaciones{
		WUsuario: u,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 360).Unix(),
			Issuer:    "Conexión Bus de Servicio Empresarial ESB-INEA",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, peticion)
	rs, e := token.SignedString(LlavePrivada)
	util.Fatal(e)
	return rs
}
