package handler

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"

	"SimpleFizzBuzz/services"
	"SimpleFizzBuzz/services/request"
)

var semaphore = make(chan struct{}, 1000)

type ApiHandler struct {
	apiService services.Apis
}

func NewApiHandler() *ApiHandler {
	return &ApiHandler{
		apiService: services.NewApisService(),
	}
}

func (ah *ApiHandler) SemaFizzBuzzRange(c *gin.Context) {
	semaphore <- struct{}{}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer func() {
			// Release the semaphore slot
			<-semaphore
		}()
		ah.FizzBuzzRange(c, &wg)
	}()

	wg.Wait()

}

func (ah *ApiHandler) FizzBuzzRange(c *gin.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	var form = &request.FizzBuzzRequest{
		FromStr: c.Query("from"),
		ToStr:   c.Query("to"),
	}

	res, err := ah.apiService.HandleFizzBuzz(c, form)
	if err != nil {
		log.Printf("Bad request: %v\n", err)
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"result": res})
}
