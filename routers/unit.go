package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	m "ymo.me/sbum-end/models"
	u "ymo.me/sbum-end/utils"
)

func createUnit(context *gin.Context) {
	context.Header("Content-Type", "application/json")

	var requestUnit m.Unit
	if err := context.ShouldBindJSON(&requestUnit); err != nil {
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}

	err := u.InsertUnit(requestUnit)
	if err != nil {
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "အတိုင်းအတာယူနစ်သစ်အား ထည့်သွင်းထားရှိပြီးပါပြီ။",
	})
}

func deleteUnit(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestUnit m.Unit
	if err := context.ShouldBindJSON(&requestUnit); err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}

	err := u.DeleteUnit(requestUnit)
	if err != nil {
		fmt.Println(" error ", err.Error())
		context.JSON(http.StatusNotAcceptable, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "ထိုအတိုင်းအတာယူနစ်အား ဖျက်သိမ်းပြီးပါပြီး။",
	})
}

func addUnitConverter(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestConverter m.Converter
	if err := context.ShouldBindJSON(&requestConverter); err != nil {
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	err := u.InsertUnitConverter(requestConverter)
	if err != nil {
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "အတိုင်းအတာယူနစ်နှုန်းထားသတ်မှတ် ပြီးပါပြီး။",
	})

}

func deleteUnitConverter(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestConverter m.Converter

	if err := context.ShouldBindJSON(&requestConverter); err != nil {
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	err := u.DeleteUnitConverter(requestConverter)
	if err != nil {
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	fmt.Print(" want to delete ", requestConverter)
	context.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "အတိုင်းအတာယူနစ်နှုန်းထားသတ်မှတ် ပြီးပါပြီး။",
	})
}

func fetchAllUnit(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	units, err := u.FindAllUnit()
	if err != nil {
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   units,
	})
}
