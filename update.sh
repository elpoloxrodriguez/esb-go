#!/bin/sh

echo "Bajando el servicio esb.inea"
echo "Por favor espere..."
pkill esb.inea

echo "Eliminando la versión actual del servicio esb.inea"
echo "Por favor	espere..."
rm -rf esb.inea

echo "Compilando la nueva versión del servicio esb.inea"
echo "Por favor	espere..."
go build -o esb.inea servicio.go

echo "Lanzando nueva versión del servicio esb.inea"
echo "Por favor	espere..."
./esb.inea &
