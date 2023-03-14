package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string
	Title  string
	Artist string
	Year   int
}

var albums = []album{
	{ID: "1", Title: "Cancion1", Artist: "S.Landry", Year: 2010},
	{ID: "2", Title: "Cancion2", Artist: "Tiesto", Year: 2012},
	{ID: "3", Title: "Cancion3", Artist: "T.Trumpet", Year: 2016},
	{ID: "4", Title: "Cancion4", Artist: "O.Heldens", Year: 2019},
	{ID: "5", Title: "Cancion5", Artist: "Gordon", Year: 2022},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item no encontrado"})
}

func welcome(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "welcome to backend in go")
}

// FUNCION MADRE
func main() {

	router := gin.Default()

	router.GET("/", welcome)
	router.GET("/albums", getAlbums)
	router.GET("/album/:id", getAlbumById)
	router.POST("/albums", postAlbums)
	router.Run("localhost:4000")
}
