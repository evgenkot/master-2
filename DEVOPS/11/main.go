package main

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
    Likes  int     `json:"likes"`
    Dislikes int   `json:"dislikes"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99, Likes: 2000, Dislikes: 123},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99, Likes: 1237, Dislikes: 12347},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99, Likes: 1299, Dislikes: 192},
}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)
    router.PUT("/albums/:id", updateAlbum)
    router.PATCH("/albums/:id", patchAlbum)
    router.DELETE("/albums/:id", deleteAlbumByID)

    router.POST("/albums/:id/like", likeAlbum)
    router.POST("/albums/:id/dislike", dislikeAlbum)

    router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album

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
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop through the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}


// deleteAlbumByID removes the album from the slice based on the ID provided.
func deleteAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop through the list of albums to find the index of the album to delete.
    for i, a := range albums {
        if a.ID == id {
            // Remove the album from the slice.
            albums = slices.Delete(albums, i, i+1)
            c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// updateAlbum replaces the album with the specified ID with the new album data.
func updateAlbum(c *gin.Context) {
    id := c.Param("id")
    var updatedAlbum album

    if err := c.BindJSON(&updatedAlbum); err != nil {
        return
    }

    for i, a := range albums {
        if a.ID == id {
            updatedAlbum.ID = id // Ensure the ID remains the same
            albums[i] = updatedAlbum
            c.IndentedJSON(http.StatusOK, updatedAlbum)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// patchAlbum updates the specified fields of the album with the given ID.
func patchAlbum(c *gin.Context) {
    id := c.Param("id")
    var updatedFields album

    if err := c.BindJSON(&updatedFields); err != nil {
        return
    }

    for i, a := range albums {
        if a.ID == id {
            if updatedFields.Title != "" {
                a.Title = updatedFields.Title
            }
            if updatedFields.Artist != "" {
                a.Artist = updatedFields.Artist
            }
            if updatedFields.Price != 0 {
                a.Price = updatedFields.Price
            }
            if updatedFields.Likes != 0 {
                a.Likes = updatedFields.Likes
            }
            if updatedFields.Dislikes != 0 {
                a.Dislikes = updatedFields.Dislikes
            }
            albums[i] = a
            c.IndentedJSON(http.StatusOK, albums[i])
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// likeAlbum adds 1 to album likes
func likeAlbum(c *gin.Context) {
    id := c.Param("id")

    for i, a := range albums {
        if a.ID == id {
            albums[i].Likes++
            c.IndentedJSON(http.StatusOK, albums[i])
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// dislikeAlbum adds 1 to album dislikes
func dislikeAlbum(c *gin.Context) {
    id := c.Param("id")

    for i, a := range albums {
        if a.ID == id {
            albums[i].Dislikes++
            c.IndentedJSON(http.StatusOK, albums[i])
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
