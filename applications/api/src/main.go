package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type photo struct {
	AlbumId int    `json:"albumId"`
	Id      int    `json:"id"`
	Title   string `json:"Title"`
	Url     string `json:"url"`
}

// albums slice to seed record album data.
var photos = []photo{
	{AlbumId: 1, Id: 1, Title: "accusamus beatae ad facilis cum similique qui sunt", Url: "https://via.placeholder.com/600/92c952"},
	{AlbumId: 1, Id: 2, Title: "reprehenderit est deserunt velit ipsam", Url: "https://via.placeholder.com/600/771796"},
	{AlbumId: 1, Id: 3, Title: "officia porro iure quia iusto qui ipsa ut modi", Url: "https://via.placeholder.com/600/24f355"},
}

// getAlbums responds with the list of all albums as JSON.
func getPhotos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, photos)
}

func main() {
	router := gin.Default()
	router.GET("/photos", getPhotos)

	router.Run("0.0.0.0:8080")
}
