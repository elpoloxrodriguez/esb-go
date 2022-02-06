package util

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Mensajes struct {
	Msj         string    `json:"msj,onmitempty"`
	Tipo        int       `json:"tipo,onmitempty"`
	Fecha       time.Time `json:"fecha,onmitempty"`
	Responsable int       `json:"responsable,onmitempty"`
}

//NullTime Tiempo nulo
type NullTime struct {
	Time  time.Time
	Valid bool
}

const layout string = "2006-01-02"

//ValidarNullFloat64 los campos nulos de la base de datos y retornar su valor original
func ValidarNullFloat64(b sql.NullFloat64) (f float64) {
	if b.Valid {
		str := strconv.FormatFloat(b.Float64, 'f', 2, 64)
		f, _ = strconv.ParseFloat(str, 64)
	} else {
		f = 0
	}
	return
}

//ValidarNullTime los campos nulos de la base de datos y retorna fecha
func ValidarNullTime(b interface{}) (t time.Time) {
	t, e := b.(time.Time)
	if !e {
		return time.Now()
	}
	return
}

//ValidarNullString Validar los campos nulos de la base de datos y retornar su valor original
func ValidarNullString(b sql.NullString) (s string) {
	if b.Valid {
		s = b.String
	} else {
		s = "null"
	}
	return
}

func GetFechaConvert(f sql.NullString) (dateStamp time.Time) {
	fecha := ValidarNullString(f)
	if fecha != "null" {
		dateString := strings.Replace(fecha, "/", "-", -1)
		dateStamp, _ = time.Parse(layout, dateString)
	}
	return
}

//Fatal Error
func Fatal(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

//GenerarHash256 Generar Claves 256 para usuarios
func GenerarHash256(password []byte) (encry string) {
	h := sha256.New()
	h.Write(password)
	encry = hex.EncodeToString(h.Sum(nil))
	return

}

//Error Procesa errores del sistema
func Error(e error) {
	if e != nil {
		fmt.Println("\n Utilidad Error: ", e.Error())
	}
}

//EjecutarScript ejecucion de comandos
func EjecutarScript() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	if err := exec.CommandContext(ctx, "/bin/sh", "update.sh").Run(); err != nil {
		// This will fail after 100 milliseconds. The 5 second sleep
		// will be interrupted.
	}
	return
}

//GitAll Actualizando proyectos
func GitAll(paquete string, cmd string, origen string) (out []byte, err error) {
	if paquete == "bus" {
		origen = "."
	} else {
		origen = "public_web/SSSIFANB/" + paquete + "/"
	}
	argstr := []string{"gitall.sh", origen, cmd}
	out, err = exec.Command("/bin/sh", argstr...).Output()
	Error(err)
	return
}
