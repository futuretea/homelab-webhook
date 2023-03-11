package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST(uptimeKumaServerChanPath, uptimeKumaServerChanHandler)
	r.POST(harvesterServerChanPath, harvesterServerChanHandler)
	if err := r.Run(); err != nil {
		panic(err)
	}
}
