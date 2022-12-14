package main

/*
API-REST-GIN is a simple API REST that allows to GET, POST, PUT and DELETE some music albums 
using Gin framework, net/http package and POSTMAN to try the HTTP requests.
*/

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Creating the data structures
type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Year int `json:"year"`
}

// Creating the data of type []album who will contain the information.
var albums = []album{
	{ID: "1", Title: "El Madrileño", Artist: "C Tangana", Year: 2021},
	{ID: "2", Title: "Hybrid Theory", Artist: "Linkin Park", Year: 2002},
	{ID: "3", Title: "Ser Humano", Artist: "Tiro de Gracia", Year: 1997},
	{ID: "4", Title: "Un Verano sin Ti", Artist: "Bad Bunny", Year: 2022},
	{ID: "5", Title: "Master of Puppets", Artist: "Metallica", Year: 1986},
	{ID: "6", Title: "And Justice For All", Artist: "Metallica", Year: 1988},
	{ID: "7", Title: "Un Día en Suburbia", Artist: "Nach Scratch", Year: 2008},
	{ID: "8", Title: "Swimming", Artist: "Mac Miller", Year: 2008},
	{ID: "9", Title: "Mr Morale & The Big Steppers", Artist: "Kendrick Lamar", Year: 2022},
	{ID: "10", Title: "Muerte", Artist: "Canserbero", Year: 2012},
}

// getAlbums captures the client request and return a JSON data and a Status.
func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums creates and sends to the server a new album.
func postAlbums(c *gin.Context) {
	var newAlbum album

	 // We send the reference to the variable and bind to JSON.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// We check if the new album is repeated or have an existing ID 
	// and response with a 400 bad request status code.
	for _, a := range albums {
		if a.ID == newAlbum.ID {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"ID already exists."})
			return
		}
		if a.Title == newAlbum.Title {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Album already exists."})
			return
		}
	}
	// If everything went good, the new album is added to the list.
	albums = append(albums, newAlbum) 
	c.IndentedJSON(http.StatusCreated, albums) // Return to the cliente a Status that data were created and the new album. 
}

// getAlbumByID takes gin Context 
func getAlbumByID(c *gin.Context){
	id := c.Param("id")

	// Search for id in albums. If match, return statusOK and the album.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		} 
	}
	// If not found, return status and a message.
	c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Album not found"})	
}

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var newAlbums []album

	// Search for id into albums. If match, return statusOK and remove the album.
	for i, a := range albums {
		if a.ID == id {
			newAlbums = append(newAlbums, albums[:i]...)
			newAlbums = append(newAlbums, albums[i+1:]...)
			albums = newAlbums
			c.IndentedJSON(http.StatusOK, newAlbums)
			return
		} 
	}
	// If not found, return status and a message.
	c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Cannot delete album. Not found"})
}

func replaceAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var newAlbums []album 

	var newAlbum album

	 // We send the reference to the variable and bind to JSON.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	for i, a := range albums {
		if a.ID == id {
			newAlbums = append(newAlbums, albums[:i]...)
			newAlbums = append(newAlbums,newAlbum)
			newAlbums = append(newAlbums, albums[i+1:]...)
			albums = newAlbums
			c.IndentedJSON(http.StatusOK, newAlbums)
			return
		} 
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Cannot replace album. ID not found"})
}

func main(){
	
	router := gin.Default() // Initialize router

	// Methods, routes and functions.
	router.GET("/albums", getAlbums) // Request to the server the data from all Albums.
	router.GET("/albums/:id", getAlbumByID) //  Request to server especific album ID and return it.
	router.POST("/albums", postAlbums) // Create new album and add it to the music library.
	router.PUT("/albums/:id",replaceAlbumByID) // Replace one album with another if it match and existing ID.
	router.DELETE("/albums/:id", deleteAlbumByID) // Delete album by ID. If it matchs, is erased.
	router.Run("localhost:8080") // Executes the server in local machine.
}

