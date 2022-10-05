package main

import (
	"github.com/MarcosOliveir4/nt-quiz/server/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/questions", controllers.Questions)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
