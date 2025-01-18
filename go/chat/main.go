package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/glebarez/go-sqlite"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"passwd"`
}

type Channel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Message struct {
	ID        int    `json:"id"`
	ChannelID int    `json:"channel_id"`
	UserID    int    `json:"user_id"`
	UserName  string `json:"username"`
	Text      string `json:"message"`
}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Working Directory", wd)

	db, err := sql.Open("sqlite", wd+"/database.db")

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	r := gin.Default()
	if err != nil {
		log.Fatal(err)
	}

	// Creation Endpoints
	r.POST("/users", func(c *gin.Context) { createUser(c, db) })
	r.POST("/channels", func(c *gin.Context) { createChannel(c, db) })
	r.POST("/messages", func(c *gin.Context) { createMessage(c, db) })

	// Listing Endpoints
	r.GET("/channels", func(c *gin.Context) { listChannels(c, db) })
	r.GET("/messages", func(c *gin.Context) { listMessages(c, db) })

	// Login Endpoints
	r.POST("/login", func(c *gin.Context) { login(c, db) })

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func createUser(c *gin.Context, db *sql.DB) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func login(c *gin.Context, db *sql.DB) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	row := db.QueryRow("SELECT if FROM users WHERE username = ? AND password = ?", user.Username, user.Password)

	var id int
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func createChannel(c *gin.Context, db *sql.DB) {
	var channel Channel

	if err := c.ShouldBindJSON(&channel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := db.Exec("INSERT INTO channels (name), VALUES (?)", channel.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func listChannels(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT id, name FROM channels")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var channels []Channel

	for rows.Next() {
		var channel Channel

		err := rows.Scan(&channel.ID, &channel.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		channels = append(channels, channel)
	}

	c.JSON(http.StatusOK, channels)
}

func createMessage(c *gin.Context, db *sql.DB) {
	var message Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := db.Exec("INSERT INTO messages (channel_id, user_id, message) VALUES (?, ?, ?)", message.ChannelID, message.UserID, message.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func listMessages(c *gin.Context, db *sql.DB) {
	channelID, err := strconv.Atoi(c.Query("channelID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 100
	}

	lastMessageID, err := strconv.Atoi(c.Query("lastMessageID"))
	if err != nil {
		lastMessageID = 0
	}

	rows, err := db.Query("SELECT m.id, channel_id, user_id, u.username AS user_name, message FROM messages m LEFT JOIN users u ON u.id = m.user_id WHERE channel_id = ? AND m.id > ? ORDER BY m.id ASC LIMIT ?", channelID, lastMessageID, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var messages []Message

	for rows.Next() {
		var message Message

		err := rows.Scan(&message.ID, &message.ChannelID, &message.UserID, &message.UserName, &message.Text)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		messages = append(messages, message)
	}

	c.JSON(http.StatusOK, messages)
}
