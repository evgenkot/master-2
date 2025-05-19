package main

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

// anecdote represents all data about a anecdote.
type anecdote struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Author string  `json:"author"`
    Text string  `json:"text"`
    Likes  int     `json:"likes"`
    Dislikes int   `json:"dislikes"`
}

// anecdotes slice to seed record anecdote data.
var anecdotes = []anecdote{
    {ID: "1", Title: "Анекдот про сгоревшего медведя", Author: "Неизвестен", Text: "Шёл Медведев по лесу. Видит: машина горит. Сел в нее и сгорел.", Likes: 20000, Dislikes: 2000},
    {ID: "2", Title: "Мужик и шляпа", Author: "Неизвестен", Text: "Купил мужик шляпу, а она ему как раз.", Likes: 20000, Dislikes: 20000},
    {ID: "3", Title: "Без названия", Author: "Неизвестен", Text: "Неизвестен - Без названия...", Likes: 10000, Dislikes: 1},
}

func main() {
    router := gin.Default()
    router.GET("/api/anecdotes", getAnecdote)
    router.GET("/api/anecdotes/:id", getAnecdoteByID)
    router.POST("/api/anecdotes", postAnecdote)
    router.PUT("/api/anecdotes/:id", updateAnecdote)
    router.PATCH("/api/anecdotes/:id", patchAnecdote)
    router.DELETE("/api/anecdotes/:id", deleteAnecdoteByID)

    router.POST("/api/anecdotes/:id/like", likeAnecdote)
    router.POST("/api/anecdotes/:id/dislike", dislikeAnecdote)

    router.Run("0.0.0.0:8080")
}

// getAnecdote responds with the list of all anecdotes as JSON.
func getAnecdote(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, anecdotes)
}

// postAnecdote adds an anecdote from JSON received in the request body.
func postAnecdote(c *gin.Context) {
    var newAnecdote anecdote

    // Call BindJSON to bind the received JSON to
    // newAnecdote.
    if err := c.BindJSON(&newAnecdote); err != nil {
        return
    }

    // Add the new anecdote to the slice.
    anecdotes = append(anecdotes, newAnecdote)
    c.IndentedJSON(http.StatusCreated, newAnecdote)
}

// getAnecdoteByID locates the anecdote whose ID value matches the id
// parameter sent by the client, then returns that anecdote as a response.
func getAnecdoteByID(c *gin.Context) {
    id := c.Param("id")

    // Loop through the list of anecdotes, looking for
    // an anecdote whose ID value matches the parameter.
    for _, a := range anecdotes {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "anecdote not found"})
}


// deleteAnecdoteByID removes the anecdote from the slice based on the ID provided.
func deleteAnecdoteByID(c *gin.Context) {
    id := c.Param("id")

    // Loop through the list of anecdotes to find the index of the anecdote to delete.
    for i, a := range anecdotes {
        if a.ID == id {
            // Remove the anecdote from the slice.
            anecdotes = slices.Delete(anecdotes, i, i+1)
            c.IndentedJSON(http.StatusOK, gin.H{"message": "anecdote deleted"})
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "anecdote not found"})
}

// updateAnecdote replaces the anecdote with the specified ID with the new anecdote data.
func updateAnecdote(c *gin.Context) {
    id := c.Param("id")
    var updatedAnecdote anecdote

    if err := c.BindJSON(&updatedAnecdote); err != nil {
        return
    }

    for i, a := range anecdotes {
        if a.ID == id {
            updatedAnecdote.ID = id // Ensure the ID remains the same
            anecdotes[i] = updatedAnecdote
            c.IndentedJSON(http.StatusOK, updatedAnecdote)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "anecdote not found"})
}

// patchAnecdote updates the specified fields of the anecdote with the given ID.
func patchAnecdote(c *gin.Context) {
    id := c.Param("id")
    var updatedFields anecdote

    if err := c.BindJSON(&updatedFields); err != nil {
        return
    }

    for i, a := range anecdotes {
        if a.ID == id {
            if updatedFields.Title != "" {
                a.Title = updatedFields.Title
            }
            if updatedFields.Author != "" {
                a.Author = updatedFields.Author
            }
            if updatedFields.Text != "" {
                a.Text = updatedFields.Text
            }
            if updatedFields.Likes != 0 {
                a.Likes = updatedFields.Likes
            }
            if updatedFields.Dislikes != 0 {
                a.Dislikes = updatedFields.Dislikes
            }
            anecdotes[i] = a
            c.IndentedJSON(http.StatusOK, anecdotes[i])
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "anecdote not found"})
}

// likeAnecdote adds 1 to anecdote likes
func likeAnecdote(c *gin.Context) {
    id := c.Param("id")

    for i, a := range anecdotes {
        if a.ID == id {
            anecdotes[i].Likes++
            c.IndentedJSON(http.StatusOK, anecdotes[i])
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "anecdote not found"})
}

// dislikeAnecdote adds 1 to anecdote dislikes
func dislikeAnecdote(c *gin.Context) {
    id := c.Param("id")

    for i, a := range anecdotes {
        if a.ID == id {
            anecdotes[i].Dislikes++
            c.IndentedJSON(http.StatusOK, anecdotes[i])
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "anecdote not found"})
}
