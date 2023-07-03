package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Artist struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Hobby string `json:"hobby"`
}

type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist Artist
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Alolo Sayang", Artist: Artist{Name: "Happy Asmara", Age: 26, Hobby: "Singing"}, Price: 56.00},
	{ID: "2", Title: "Satru 3", Artist: Artist{Name: "Denny Caknan", Age: 28, Hobby: "Singing"}, Price: 40.00},
	{ID: "3", Title: "ILU IMU", Artist: Artist{Name: "Woro", Age: 28, Hobby: "Singing"}, Price: 80.00},
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
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumsById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Album Not Found"})
}

func updateAlbums(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			
		}
	}
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumsById)

	router.Run()
}
