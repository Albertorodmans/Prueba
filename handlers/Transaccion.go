package handlers

import (
	"fmt"
	"net/http"

	"github.com/banwire/api-exam/models"
	base "github.com/banwire/api-exam/service"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

var colTransaccion = base.Connectiondbase().DB("Test").C("transaccion")

func AddTransaccion(c *gin.Context) {

	merchant_id := c.Param("id")

	//obtener la estructura que nos envía el cliente
	var tran models.Transaccion
	err := c.BindJSON(&tran)
	//si ocurrió un error, notificarlo al cliente
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"Accion": err.Error(),
				"json":   &tran,
			})
		return
	}
	//si no contiene el nombre y la comisión, regresarlo como error
	if !validateTransaction(tran) || merchant_id == "" {
		c.JSON(http.StatusOK,
			gin.H{
				"Accion": "Datos incompletos",
				"json":   &tran,
			})
		return
	}

	//buscar el comercio correspondiente al parámetro
	var comercio models.Comercio
	comObjectID := bson.ObjectIdHex(merchant_id)
	err = colComercio.Find(bson.M{"merchant_id": comObjectID}).One(&comercio)
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"Accion": "Error al buscar el comercio " + err.Error(),
				"json":   comercio,
			})
		return
	}

	tran.TransaccionId = bson.NewObjectId()
	tran.Commission = comercio.Commission
	tran.MerchantId = comObjectID

	colTransaccion.Insert(tran)
	c.JSON(http.StatusOK,
		gin.H{
			"Accion": "Ejecutada",
			"json":   &tran,
		})
	return
}

func validateTransaction(tr models.Transaccion) bool {
	if tr.Amount == 0 {
		return false
	}
	return true
}

func GetProfits(c *gin.Context) {

	var comercios models.Comercios
	err := colComercio.Find(bson.D{}).All(&comercios)
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"Accion": "Error al buscar el comercio " + err.Error(),
				"json":   comercios,
			})
		return
	}
	ganancia := 0.0
	for _, v := range comercios {
		//var transacciones models.Transaccions

		pipeline := []bson.M{
			{
				"$match": bson.M{
					"merchant_id": v.MerchantId,
				},
			},
			{
				"$group": bson.M{
					"_id":   "",
					"total": bson.M{"$sum": "$amount"},
				},
			},
		}
		result := []bson.M{}
		err = colTransaccion.Pipe(pipeline).All(&result)

		if len(result) > 0 {
			totalT := result[0]["total"].(float64)
			totalT = (totalT * float64(v.Commission)) / 100
			ganancia += totalT
		}

	}

	c.JSON(http.StatusOK,
		gin.H{
			"Accion": "Ejecutada",
			"json":   ganancia,
		})

	return
}

func GetEspecificProfit(c *gin.Context) {

	comID := c.Param("id")
	if comID == "" {
		c.JSON(http.StatusOK,
			gin.H{
				"Accion": "Proporcione un identificador válido",
				"json":   "",
			})
		return
	}

	comObjectID := bson.ObjectIdHex(comID)

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"merchant_id": comObjectID,
			},
		},
		{
			"$group": bson.M{
				"_id":      "",
				"comision": bson.M{"$first": "$commission"},
				"total":    bson.M{"$sum": "$amount"},
			},
		},
	}
	result := []bson.M{}
	err := colTransaccion.Pipe(pipeline).All(&result)
	fmt.Println("result and err", result, err)
	ganancia := 0.0
	if len(result) > 0 {
		totalT := result[0]["total"].(float64)
		comision := float64(result[0]["comision"].(int))
		totalT = (totalT * float64(comision)) / 100
		ganancia += totalT
	}

	c.JSON(http.StatusOK,
		gin.H{
			"Accion": "Ejecutada",
			"json":   ganancia,
		})

	return

}
