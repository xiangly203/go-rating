package main

func main() {
	gin.DisableConsoleColor()
	//记录日志
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router = gin.Default()
	router.POST("/getCode", postCode)
	router.POST("/login", postLogin)
	router.Run("localhost:8080")
}

func getCode(c *gin.Context) {
	var req LoginReq
	var message string
	if err := c.BindJSON(&newAlbum); err != nil {
		message := err
		return
	}

}
