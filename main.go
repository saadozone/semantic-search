package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saad/semantic-search/internal/db"
	"github.com/saad/semantic-search/internal/faiss"
)

func main() {
	// Test FAISS
	faiss.TestFAISS()

	// Init DB
	database := db.InitDB("semantic.db")

	// Setup Gin router
	r := gin.Default()

	// Health endpoint
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Test endpoint: return all docs
	r.GET("/docs", func(c *gin.Context) {
		rows, err := database.Query("SELECT id, title, content FROM documents")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		type Doc struct {
			ID      int    `json:"id"`
			Title   string `json:"title"`
			Content string `json:"content"`
		}

		var docs []Doc
		for rows.Next() {
			var d Doc
			rows.Scan(&d.ID, &d.Title, &d.Content)
			docs = append(docs, d)
		}

		c.JSON(http.StatusOK, docs)
	})

	// Insert a test doc
	_, err := database.Exec("INSERT INTO documents (title, content) VALUES (?, ?)",
		"Hello World", "This is a test document")
	if err != nil {
		fmt.Println("Insert error:", err)
	}

	// Run server
	r.Run(":8080")
}
