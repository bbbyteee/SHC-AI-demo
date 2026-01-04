package image

import (
	"log"
	"net/http"
	"shc-ai-demo/common/code"
	"shc-ai-demo/controller"
	"shc-ai-demo/service/image"

	"github.com/gin-gonic/gin"
)

type (
	RecognizeImageResponse struct {
		ClassName string `json:"class_name,omitempty"`
		controller.Response
	}
)

func RecognizeImage(c *gin.Context) {
	res := new(RecognizeImageResponse)

	file, err := c.FormFile("image")
	if err != nil {
		log.Println("FormFile fail", err)
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}
	className, err := image.RecognizeImage(file)
	if err != nil {
		log.Println("RecognizeImage fail", err)
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	res.Success()
	res.ClassName = className
	c.JSON(http.StatusOK, res)
}
