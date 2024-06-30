package main

import (
	"github.com/joshbrgs/my-portfolio/internal/handler"
	"github.com/joshbrgs/my-portfolio/pkg/server"
)

func main() {
	options := server.DefaultServerOptions()
	server := server.NewServer(options)

	// Serve static files from the "static" directory
	server.Static("/static", "static")

	indexHandler := handler.IndexHandler{}
	aboutHandler := handler.AboutHandler{}

	server.GET("/", indexHandler.HandleIndex)
	server.GET("/about", aboutHandler.HandleAbout)

	if err := server.Start(":" + options.Port); err != nil {
		server.Logger.Panic(err)
	}
	server.Logger.Info("server started on http://localhost:" + options.Port)
}
