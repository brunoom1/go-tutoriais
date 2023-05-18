package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Album struct {
	ID     string  `json:"ID"`
	Title  string  `json:"Title"`
	Artist string  `json:"Artist"`
	Price  float32 `json:"Price"`
}

var albums = []Album{
	{ID: "1", Title: "Album 1", Artist: "Ruan Carlos", Price: 10.2},
	{ID: "2", Title: "Album 2", Artist: "Juliano Medeiros", Price: 123.00},
	{ID: "3", Title: "Album 3", Artist: "Juliano Medeiros", Price: 234.00},
}

// routes

func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

func setAlbums(c *gin.Context) {
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbum(c *gin.Context) {
	id, ok := c.Params.Get("id")

	if ok {
		for _, album := range albums {
			if album.ID == id {
				c.JSON(http.StatusOK, album)
				return
			}
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

// main
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", setAlbums)
	router.GET("/albums/:id", getAlbum)

	router.Run()
}
