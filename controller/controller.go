package controller

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/sshaunn/train-go/ms1/service"
)

func init() {
	fmt.Println("controller init...")
}

// type MessageEvent interface{}

func Upload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(500, err)
	}
	files := form.File["file"]
	wg := &sync.WaitGroup{}
	ch := make(chan interface{}, len(files))
	wg.Add(len(files))
	slice := make([]interface{}, len(files))
	service.ConcurrencyWorkflow(service.UploadFiles(files, ch, wg), service.ProduceMessage(ch, slice))
	wg.Wait()
	c.JSON(200, slice)
}
