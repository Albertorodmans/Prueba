package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/banwire/api-exam/models"
	base "github.com/banwire/api-exam/service"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

var coluComercio = base.Connectiondbase().DB("Test").C("comercios")

func UpdateComerce(c *gin.Context) {

	//var com models.Comercio
	comID := c.Param("id")

	//obtener la estructura que nos envía el cliente
	var com models.Comercio
	err := c.BindJSON(&com)
	//si ocurrió un error, notificarlo al cliente
	if err != nil && !strings.Contains(err.Error(), "parsing time") {

		c.JSON(http.StatusOK,
			gin.H{
				"Accion": err.Error(),
				"json":   &com,
			})
		return
	}

	//si no contiene el nombre y la comisión, regresarlo como error

	if !validateComerce(com) || comID == "" {
		c.JSON(http.StatusOK,
			gin.H{
				"Accion": "El nombre está vació o la comisión no cumple con lo requerido",
				"json":   com,
			})
		return
	}

	var recordOri models.Comercio
	comObjectID := bson.ObjectIdHex(comID)
	err = coluComercio.Find(bson.M{"merchant_id": comObjectID}).One(&recordOri)
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"Accion": "Error al buscar el comercio " + err.Error(),
				"json":   recordOri,
			})
		return
	}

	recordOri.UpdatedAt = time.Now()
	recordOri.MerchantName = com.MerchantName
	recordOri.Commission = com.Commission

	coluComercio.Update(bson.M{"merchant_id": comObjectID}, recordOri)
	c.JSON(http.StatusOK,
		gin.H{
			"Accion": "Ejecutada",
			"json":   &recordOri,
		})
	return
}

/* func buscartodo(c *gin.Context) {
	var todo models.Comercio
	//todoID := c.Param("id")
	if todo.MerchantId == "" {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	_todo := models.Comercio{MerchantId: todo.MerchantId, MerchantName: todo.MerchantName, Commission: todo.Commission}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todo})
} */

/* func BuscaTodo(c *gin.Context) {
	var todos []models.Comercio
	var _todos []models.Comercio
	coluComercio.Find(&todos)
	if len(todos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	//transforms the todos for building a good response
	for _, item := range todos {
		//completed := false
		if item.Commission == 10 {
			//  completed = true
		} else {
			//completed = false
		}
		_todos = append(_todos, models.Comercio{MerchantId: item.MerchantId, MerchantName: item.MerchantName, Commission: item.Commission})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})
} */
