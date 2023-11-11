// Code generated by hertz generator.

package main

import (
	"Blog/dao"
	"Blog/timedtask"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/cors"
)

func main() {
	dao.Init()
	timedtask.Init()
	h := server.Default(
		server.WithHostPorts(":8080"),
	)
	h.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowHeaders:     []string{"Origin", "accept", "x-requested-with", "Content-Type", "X-Custom-Header"},
		AllowMethods:     []string{"PUT", "PATCH", "OPTIONS"},
	}))
	h.StaticFS("/static/", &app.FS{Root: "./", GenerateIndexPages: true})
	register(h)
	h.Spin()
}
