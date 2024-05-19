package productcontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/harryprasetyo4000/go-restapi-test.git/models"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var products []models.Product
	startTime := time.Now()
	// START CODE

	models.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{"products": products})

	// END CODE
	defer func() {
		endTime := time.Now()
		CreateAndInsertApiLog(c, startTime, endTime)
	}()
}

func Show(c *gin.Context) {
	var product []models.Product
	startTime := time.Now()
	// START CODE

	// MENGAMBIL DATA BERDASARKAN PARAM ID
	id := c.Param("id")

	// mengambil data berdasrkan primary key di table product
	if err := models.DB.First(&product, id).Error; err != nil {

		// CASE JIKA ERROR KARNA DATA TIDAK ADA DAN ERROR KARNA INTERNAL SERVER ERROR
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Not Found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error", "errorMessage": err.Error()})
			return
		}
	}

	// JIKA TIDAK ERROR MENAMPILKAN DATA
	c.JSON(http.StatusOK, gin.H{"data": product})

	// END CODE
	defer func() {
		endTime := time.Now()
		CreateAndInsertApiLog(c, startTime, endTime)
	}()
}

func Create(c *gin.Context) {
	var product models.Product
	startTime := time.Now()
	// START CODE

	// MENGAMBIL BODY JSON, JIKA TIDAK SAMA RETURN Bad Request
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// JIKA BODY SAMA, MENULIS KE DB
	models.DB.Create(&product)

	// JIKA TIDAK ERROR MENAMPILKAN DATA
	c.JSON(http.StatusOK, gin.H{"data": product})

	// END CODE
	defer func() {
		endTime := time.Now()
		CreateAndInsertApiLog(c, startTime, endTime)
	}()
}

func Update(c *gin.Context) {
	var product models.Product
	startTime := time.Now()
	// START CODE

	// MENGAMBIL DATA BERDASARKAN PARAM ID
	id := c.Param("id")

	// MENGAMBIL BODY JSON, JIKA TIDAK SAMA RETURN Bad Request
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// MENGUPDATE PRODUCT BERDASARKAN ID, JIKA ID TIDAK ADA MAKA RETURN ID PRODUCT NOT FOUND
	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "ID Product Not Found"})
		return
	}

	// JIKA TIDAK ERROR MENAMPILKAN DATA
	c.JSON(http.StatusOK, gin.H{"message": "Update Succesfully", "data": product})

	// END CODE
	defer func() {
		endTime := time.Now()
		CreateAndInsertApiLog(c, startTime, endTime)
	}()
}

func Delete(c *gin.Context) {
	var product models.Product
	startTime := time.Now()
	// START CODE

	var input struct {
		Id json.Number
	}

	// MENGAMBIL BODY JSON, JIKA TIDAK SAMA RETURN Bad Request
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// MENGAMBIL DATA ID SEBAGAI INT64
	id, _ := input.Id.Int64()

	// MENGHAPUS PRODUCT BERDASARKAN ID, JIKA ID PRODUCT TIDAK ADA RETURN FAILED DELETE PRODUCT
	if models.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed Delete Product"})
		return
	}

	// JIKA TIDAK ERROR MENAMPILKAN DATA
	c.JSON(http.StatusOK, gin.H{"message": "DELETE Succesfully"})

	// END CODE
	defer func() {
		endTime := time.Now()
		CreateAndInsertApiLog(c, startTime, endTime)
	}()
}

// MENULIS APILOG KE DB
func CreateAndInsertApiLog(c *gin.Context, startTime, endTime time.Time) {
	// Get request parameters
	getParams := c.Request.URL.Query()
	getParamsString := fmt.Sprintf("%v", getParams)

	// Get request origin
	origin := ""
	if originValue := c.Request.Header.Get("Origin"); originValue != "" {
		origin = originValue
	}

	// MENGISI VALUE APILOG
	apiLog := models.ApiLog{
		ID:              uuid.New(),
		Duration:        int64(time.Since(startTime).Milliseconds()),
		EndPoint:        c.Request.URL.Path,
		HttpStatus:      strconv.Itoa(c.Writer.Status()),
		Method:          c.Request.Method,
		TsStart:         startTime.Format("2006-01-02 15:04:05"),
		TsEnd:           endTime.Format("2006-01-02 15:04:05"),
		UserAgent:       c.Request.UserAgent(),
		ReqID:           uuid.New().String(),
		Origin:          origin,
		ReqSize:         c.Request.ContentLength,
		InstanceName:    "APM1",
		ClientIP:        c.ClientIP(),
		UpstreamElapsed: 50,
		RemoteHost:      "127.0.0.1",
		UserName:        "user123",
		Params:          getParamsString,
		ResSize:         int64(c.Writer.Size()),
		SendToUpstream:  0,
	}

	// MENULIS KE TABLE APILOG
	if err := models.DB.Create(&apiLog).Error; err != nil {
		panic("failed to insert record: " + err.Error())
	}
}
