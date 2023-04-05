package main

import (
	"embed"
	"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

// Embed a directory
//go:embed dist/go-ui/*
var embedDirStatic embed.FS

func main() {
    app := fiber.New()
    app.Use("/", filesystem.New(filesystem.Config{
        Root: http.FS(embedDirStatic),
        PathPrefix: "dist/go-ui",
        Browse: true,
    }))
    app.Listen(":3000")
}