package main

import (
	"my-app/discord"
	"my-app/middleware"
	"my-app/module/news/transport/ginnews"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// gormDB, err := db.ConnectMysql()
	// if err != nil {
	// 	fmt.Println("Error connecting to database:", err)
	// 	return
	// }
	// appContext := appctx.NewAppContext(gormDB, os.Getenv("JWT_SECRET"))

	// open discord bot
	go discord.InitDiscord()

	// gin start
	r := gin.Default()
	r.Use(middleware.CORS())

	r.Use(middleware.Recover())
	r.Static("/static", "./static")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")

	// v1.POST("/register", ginuser.Resgiter(appContext))
	// v1.POST("/authenticate", ginuser.Login(appContext))
	// v1.GET("/userinfo", middleware.RequireAuth(appContext), ginuser.UserInfo(appContext))
	v1.GET("/getnews", ginnews.GetNews())

	// restaurantRoute := v1.Group("restaurants", middleware.RequireAuth(appContext))
	// restaurantRoute.GET("/:id", ginrestaurant.GetOneRestaurant(appContext))
	// restaurantRoute.GET("", ginrestaurant.GetRestaurant(appContext))
	// restaurantRoute.POST("", ginrestaurant.CreateRestaurant(appContext))
	// restaurantRoute.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// defer db.CloseDB(gormDB)
}
