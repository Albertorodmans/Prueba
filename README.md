# api-exam
Examen BanWire.

## Ejercicio

Después de unas largas horas TI y Operaciones decidió, que es necesario realizar un api la cual permita gestionar comercios, comisiones y transacciones. Esto para tener mayor visibilidad en las transaccionalidad de los comercios.

## Definiciones

* Comercio: usuarios Banwire de los cuales procesamos sus pagos(transacciones).   
* Transacción: compra realizada por los usuarios finales de los comercios. 
* Comisión: monto en porcentaje cobrado a los comercios por transacción. 

## Requerimientos

Construir una API en el lenguaje de programacion golang utilizando el framework gin, que permita lo siguiente: 

* La API debe permitir la alta de los comercios y por cada comercio asignarle una comisión.
* Las comisiones deben ser medidas en porcentajes 1 - 100.
* La API debe permitir editar los comercios y sus comisiones.
* La API debe permitir agregar transacciones por comercio (alta de transacciones).
* Por cada transacción se debe aplicar la comisión asignada al comercio.
* La API debe permitir obtener las ganancias de todos los comercios Banwire (suma de todas las comisiones de todos los comercios).
* La API debe mostrar las ganancias por comercio (suma de comisiones de un solo comercio).

## Estructuras basicas

En seguida see muesta un ejemplo de las estructuras que deben tener las transacciones y los comercios BanWire.

#### Comercio
```js
{
  "merchant_id": 1,
  "merchant_name": "comercio1",
  "commission": 10,
  "created_at": 1602181741,
  "updated_at": 1602181741
} 
```
#### Transacción

```js
{
  "transaction_id": 20234,
  "merchant_id": 1,
  "amount": 100.50,
  "commission": 10,
  "fee": 10.50,
  //"items": [],
  "created_at": 1602181741
}
```
## Extras

* Patrones de diseño.
* Utilizar contenedores (Docker) para correr el API.

## Instalación

Para correr el proyecto con golang, por default corre en el puerto 3000.

```bash
go build -o app && ./app  
```

Para correr el proyecto con docer-compose.

```bash
docker-compose build && docker-compose up -d  
```