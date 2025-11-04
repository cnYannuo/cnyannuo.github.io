package mods

import (
	"log"
	"net/http"
	"strconv"

	mydb "test_bi/db"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db    *gorm.DB
	cache *redis.Client
)

func InitDB() {
	dsn := "root:yannuo1996@tcp(127.0.0.1:3307)/bi_test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	db.AutoMigrate(&mydb.User{})
}

func SetupRoutes(router *gin.Engine) {
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUser)
	router.POST("/users", createUser)
	router.PUT("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)
	router.GET("/cache-test", cacheTest)
}

func getUsers(c *gin.Context) {
	users, err := GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func getUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func createUser(c *gin.Context) {
	var user mydb.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效数据格式"})
		return
	}
	user, err := CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": user})
}

func updateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user mydb.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效数据格式"})
		return
	}
	user, err := UpdateUser(id, user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func deleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除失败"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func cacheTest(c *gin.Context) {
	val, err := CacheTest()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": val})
}
