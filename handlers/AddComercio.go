package handlers

import (
	"net/http"
	"time"

	"github.com/banwire/api-exam/models"
	base "github.com/banwire/api-exam/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/mgo.v2/bson"
)

var colComercio = base.Connectiondbase().DB("Test").C("comercios")

func AddnewComerce(c *gin.Context) {

	//obtener la estructura que nos envía el cliente
	var Comerce models.Comercio

	//si ocurrió un error, notificarlo al cliente
	if err := c.ShouldBindBodyWith(&Comerce, binding.JSON); err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"Accion": err.Error(),
				"json":   &Comerce,
			})
		return
	}
	//si no contiene el nombre y la comisión, regresarlo como error
	if !validateComerce(Comerce) {
		c.JSON(http.StatusOK,
			gin.H{
				"Accion": "El nombre está vació o la comisión no cumple con lo requerido",
				"json":   &Comerce,
			})
		return
	}

	Comercio := new(models.Comercio)
	Comercio.MerchantId = Comerce.MerchantId
	Comercio.CreatedAt = time.Now()
	Comercio.UpdatedAt = Comercio.CreatedAt
	//Asignacion de comisión al comercio
	Comercio.Commission = Comerce.Commission

	if err := c.ShouldBindBodyWith(&Comercio, binding.JSON); err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"Accion": "Error al parsear la información: " + err.Error(),
				"json":   &Comercio,
			})

	} else {
		colComercio.Insert(Comercio)
		c.JSON(http.StatusOK,
			gin.H{
				"Accion": "Ejecutada",
				"json":   &Comercio,
			})
		return
	}

}

func AddnewComerces(c *gin.Context) {

	//obtener la estructura que nos envía el cliente
	var comerces models.Comercios
	err := c.BindJSON(&comerces)
	//si ocurrió un error, notificarlo al cliente
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"Accion": err.Error(),
				"json":   &comerces,
			})
		return
	}

	//Recorrer todos los comercios
	for _, v := range comerces {
		if !validateComerce(*v) {
			c.JSON(http.StatusOK,
				gin.H{
					"Accion": "El nombre está vació o la comisión no cumple con lo requerido",
					"json":   &v,
				})
			return
		}

		v.MerchantId = bson.NewObjectId()
		v.CreatedAt = time.Now()
		v.UpdatedAt = v.CreatedAt

		err = colComercio.Insert(v)
		if err != nil {
			c.JSON(http.StatusOK,
				gin.H{
					"Accion": "Error al insertar: " + err.Error(),
					"json":   &comerces,
				})
			return
		}
	}

	c.JSON(http.StatusOK,
		gin.H{
			"Accion": "Ejecutada",
			"json":   &comerces,
		})
	return
}

func validateComerce(Comerce models.Comercio) bool {
	//si no contiene el nombre y la comisión, regresarlo como error
	if Comerce.Commission <= 0 || Comerce.Commission > 100 || Comerce.MerchantName == "" {
		return false
	}
	return true
}
