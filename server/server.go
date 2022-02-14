package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type Album struct {
	ID      string  `json:"id"`
	Title   string  `json:"title"`
	Artist  string  `json:"artist"`
	// A: Comment below line to test a change in Provider API
	// Price   float64 `json:"price"`
	Country string  `json:"country"`
}

// albums slice to seed record album data.
var albums = []Album{
	// A: Comment below line to test a change in Provider API. Run the CLIENT tests and see that 
	// will generate the consumer pact file with a missing "Price" field.
	// Now enable A. and run the Provider tests, it should still pass.
	// Since Consumer is concerned only about ID, Title, Artist and Country keys.
	// If there are any additional fields in the Provider, it shouldn't matter the Consumer.
	// {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99, Country: "United States"},
	// {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99, Country: "United Kingdom"},
	// {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99, Country: "India"},
	
	// B. Comment A. and run the Provider tests, it will FAIL as Consumer was expecting 'price' field 
	// But it was missing in the Provider response as Provider removed the support for it.
	// But if you run Consumer tests then the pact file will be updated and now if you run Provider tests again
	// it should pass.
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Country: "United States"},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Country: "United Kingdom"},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Country: "India"},
}

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	var newAlbum Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
