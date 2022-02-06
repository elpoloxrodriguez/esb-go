/*
Copyright 2022 Ing. Andrés Ricardo Rodríguez Durán.Todos los derechos reservados.
En informática un Bus de Servicio Empresarial (ESB por sus siglas en inglés)
es un modelo de arquitectura de software que gestiona la comunicación entre
servicios web. Es un componente fundamental de la Arquitectura Orientada a
Servicios.
Un ESB generalmente proporciona una capa de abstracción construida sobre
una implementación de un sistema de mensajes de empresa que permita a los
expertos en integración explotar el valor del envío de mensajes sin tener que
escribir código. Al contrario que sucede con la clásica integración de
aplicaciones de empresa (IAE) que se basa en una pila monolítica sobre una
implantación distribuida cuando se hace necesario, de modo que trabajen
arquitectura hub and spoke, un bus de servicio de empresa se construye sobre
unas funciones base que se dividen en sus partes constituyentes, con una
armoniosamente según la demanda.
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/elpoloxrodriguez/esb-inea/sys"
	"github.com/elpoloxrodriguez/esb-inea/sys/web"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

//Variables de Control
var (
	Enrutador = mux.NewRouter()
)

func init() {
	fmt.Println("")
	fmt.Println("..........................................................")
	fmt.Println(".... Versión del Bus de Servicio Empresarial", sys.Version, " ....")
	fmt.Println("..........................................................")
	fmt.Println("")
	if sys.MongoDB {
		fmt.Println("")
		fmt.Println("..........................................................")
		fmt.Println("... Iniciando Carga de Elementos Para el servidor WEB   ...")
		fmt.Println("..........................................................")
		fmt.Println("")
	}
}

func PuertosProduccion() {
	fmt.Println("... Puertos de Desarrollo ✅ ...")
	server := &http.Server{
		Handler:      context.ClearHandler(web.Enrutador),
		Addr:         ":" + sys.PUERTO_SSL,
		WriteTimeout: 280 * time.Second,
		ReadTimeout:  280 * time.Second,
	}
	fmt.Println("Servidor Escuchando en el puerto:", sys.PUERTO_SSL)
	go server.ListenAndServeTLS("certificados/https/cert.pem", "certificados/https/key.pem")

	srv := &http.Server{
		Handler:      context.ClearHandler(web.Enrutador),
		Addr:         ":" + sys.PUERTO_STANDAR,
		WriteTimeout: 280 * time.Second,
		ReadTimeout:  280 * time.Second,
	}
	fmt.Println("Servidor Escuchando en el puerto:", sys.PUERTO_STANDAR)
	fmt.Println("")
	go srv.ListenAndServe()
}

func PuertosDesarrollo() {
	fmt.Println("... Puertos de Produccion ✅ ...")
	srvs := &http.Server{
		Handler:      context.ClearHandler(web.Enrutador),
		Addr:         ":" + sys.PUERTO,
		WriteTimeout: 280 * time.Second,
		ReadTimeout:  280 * time.Second,
	}
	fmt.Println("Servidor Escuchando en el puerto:", sys.PUERTO)
	go srvs.ListenAndServe()

	serverx := &http.Server{
		Handler:      context.ClearHandler(web.Enrutador),
		Addr:         ":" + sys.PUERTO_SSL_STANDAR,
		WriteTimeout: 280 * time.Second,
		ReadTimeout:  280 * time.Second,
	}
	fmt.Println("Servidor Escuchando en el puerto:", sys.PUERTO_SSL_STANDAR)
	fmt.Println("..........................................................")
	log.Fatal(serverx.ListenAndServeTLS("certificados/https/cert.pem", "certificados/https/key.pem"))
}

func main() {
	fmt.Println("")
	fmt.Println("..........................................................")
	fmt.Println("..... Inciando la carga de las rutas API del sistema .....")
	fmt.Println("..........................................................")
	web.Cargar()
	fmt.Println("")

	fmt.Println("")
	fmt.Println("..........................................................")
	fmt.Println("..... Abriendo los puertos de conexion del servidor  .....")
	fmt.Println("..........................................................")
	fmt.Println("")
	PuertosProduccion()
	PuertosDesarrollo()
}
