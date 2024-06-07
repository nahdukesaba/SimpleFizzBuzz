package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"

	"SimpleFizzBuzz/handler"
)

type Server struct {
	router *gin.Engine
}

func NewServer(router *gin.Engine) *Server {
	return &Server{
		router: router,
	}
}

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(
		func(params gin.LogFormatterParams) string {
			logs := map[string]interface{}{
				"status_code":   params.StatusCode,
				"path":          params.Path,
				"method":        params.Method,
				"start_time":    params.TimeStamp.Format("2006/01/02 - 15:04:05"),
				"remote_addr":   params.ClientIP,
				"response_time": params.Latency.String(),
			}
			s, _ := json.Marshal(logs)
			return string(s) + "\n"
		})
}

func (s *Server) RegisterRoutes() {
	apisHandler := handler.NewApiHandler()
	s.router.GET("/range-fizzbuzz", apisHandler.SemaFizzBuzzRange)

}

func (s *Server) Run() {
	//s.router.Run(":9000")
	serv := &http.Server{
		Addr:           ":9000",
		Handler:        s.router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Println("receive interrupt signal")
		if err := serv.Close(); err != nil {
			log.Fatal("Server Close:", err)
		}
	}()

	if err := serv.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			log.Println("Server closed under request")
		} else {
			log.Fatal("Server closed unexpect")
		}
	}
}
