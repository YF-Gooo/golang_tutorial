package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

type Data struct {
	A int `json:"a"`
	B int `json:"b"`
}

func main() {
	r := gin.Default()
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/incr", func(c *gin.Context) {
		session := sessions.Default(c)
		// var count int
		var data Data
		v := session.Get("count9")
		if v == nil {
			data = Data{A: 1, B: 1}
		} else {
			json.Unmarshal(v.([]byte), &data)
			data.A += 1
			data.B += 1
		}
		fmt.Print(data)
		res, _ := json.Marshal(data)
		fmt.Print(string(res))
		session.Set("count9", res)
		session.Save()
		c.JSON(200, gin.H{"count8": string(res)})
	})
	r.Run(":8000")
}
