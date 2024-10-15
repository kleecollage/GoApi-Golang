package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MODEL
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// DATA
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 38.45},
}

// CONTROLLER | HANDLER
func getAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, albums)
}

// CONTROLLER | HANDLER
func getAlbumByID(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, a := range albums {
		if a.ID == id {
			ctx.IndentedJSON(http.StatusOK, a)
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album no encontrado"})
}

// CONTROLLER | HANDLER
func postAlbums(ctx *gin.Context) {
	var newAlbum album

	if err := ctx.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	// router.PUT("/albums/:id", updateAlbum)
	// router.DELETE("albums/:id", deleteAlbum)

	router.Run("localhost:8080")
}
