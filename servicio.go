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

	"githib.com/elpoloxrodriguez/esb-inea/routes"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

//Variables de Control
var (
	Enrutador = mux.NewRouter()
)

func init() {
	fmt.Println("")
	fmt.Println("Versión del Panel V.1")
	fmt.Println("")
	if false {
		fmt.Println("")
		fmt.Println("..........................................................")
		fmt.Println("... Iniciando Carga de Elemento Para el servidor WEB   ...")
		fmt.Println("..........................................................")
		fmt.Println("")
	}
}

func main() {
	fmt.Println("Inciando la carga del sistema")
	routes.Cargar()

	srv := &http.Server{
		Handler:      context.ClearHandler(Enrutador),
		Addr:         ":81",
		WriteTimeout: 280 * time.Second,
		ReadTimeout:  280 * time.Second,
	}
	fmt.Println("Servidor Escuchando en el puerto:", 81)
	go srv.ListenAndServe()

	srvs := &http.Server{
		Handler:      context.ClearHandler(Enrutador),
		Addr:         ":8080",
		WriteTimeout: 280 * time.Second,
		ReadTimeout:  280 * time.Second,
	}
	fmt.Println("Servidor Escuchando en el puerto:", 8080)
	go srvs.ListenAndServe()
	//
	//https://dominio.com/* Protocolo de capa de seguridad
	server := &http.Server{
		Handler:      context.ClearHandler(Enrutador),
		Addr:         ":2608",
		WriteTimeout: 280 * time.Second,
		ReadTimeout:  280 * time.Second,
	}
	fmt.Println("Servidor Escuchando en el puerto:", 2608)
	go server.ListenAndServeTLS("certificados/https/cert.pem", "certificados/https/key.pem")

	serverx := &http.Server{
		Handler:      context.ClearHandler(Enrutador),
		Addr:         ":443",
		WriteTimeout: 280 * time.Second,
		ReadTimeout:  280 * time.Second,
	}
	fmt.Println("Servidor Escuchando en el puerto:", 443)
	log.Fatal(serverx.ListenAndServeTLS("certificados/https/cert.pem", "certificados/https/key.pem"))
}
