package main

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// Book represents a textbook
type Book struct {
	Title    string `json:"title"`
	Filename string `json:"filename"`
}

// SemesterBooks maps semester names to a list of books
var SemesterBooks map[string][]Book

func loadBooks() error {
	// Correct path for Render (relative to project root)
	f, err := os.Open("backend/data/books.json")
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewDecoder(f).Decode(&SemesterBooks)
}

func main() {
	r := gin.Default()

	// Correct paths for serving static files from the project root
	r.Static("/static", "frontend")
	r.GET("/", func(c *gin.Context) {
		c.File("frontend/index.html")
	})

	// API group
	api := r.Group("/api")
	{
		api.GET("/semesters", func(c *gin.Context) {
			semesters := make([]string, 0, len(SemesterBooks))
			for s := range SemesterBooks {
				semesters = append(semesters, s)
			}
			c.JSON(http.StatusOK, gin.H{"semesters": semesters})
		})

		api.GET("/books/:semester", func(c *gin.Context) {
			semester := c.Param("semester")
			books, ok := SemesterBooks[semester]
			if !ok {
				c.JSON(http.StatusNotFound, gin.H{"error": "Semester not found"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"books": books})
		})
	}

	// Download endpoint with corrected path
	r.GET("/download/:semester/:filename", func(c *gin.Context) {
		semester := c.Param("semester")
		filename := c.Param("filename")
		// Correct path for Render (relative to project root)
		filePath := filepath.Join("backend/books", semester, filename)
		c.FileAttachment(filePath, filename)
	})

	// Load books data
	if err := loadBooks(); err != nil {
		panic("Failed to load books.json: " + err.Error())
	}

	// Use PORT environment variable for Render, with a fallback for local dev
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
