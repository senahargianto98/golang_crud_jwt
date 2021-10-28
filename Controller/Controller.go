package Controller

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"math"

	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"skb/Config"
	"skb/Model"
)

func Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename
	out, err := os.Create("public/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	filepath := "https://api.belajaryuk.xyz/file/" + filename
	c.JSON(http.StatusOK, gin.H{"filepath": filepath})
}

func Paginate(c *gin.Context) {
	db, _ := Config.GetDB()
	var artikel []Model.Artikel
	sql := "SELECT * FROM artikels ORDER BY id DESC"
	id, _ := strconv.Atoi(c.Params.ByName(`id`))
	page := 4
	var total int64
	sql = fmt.Sprintf("%s LIMIT %d  OFFSET %d", sql, page, (id-1)*page)
	db.Raw(sql).Scan(&artikel).Count(&total)
	c.JSON(200, gin.H{"data": artikel, "thisPage": id, "last_page": math.Ceil(float64(total / int64(page)))})
}

func GetUser(c *gin.Context) {
	db, _ := Config.GetDB()
	var user []Model.User
	sql := "SELECT * FROM users"
	if err := db.Raw(sql).Scan(&user).Error; err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, user)
	}
}

func GetAllArtikel(c *gin.Context) {
	db, _ := Config.GetDB()
	var artikel []Model.Artikel
	if err := db.Find(&artikel).Error; err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, artikel)
	}
}

func CreateArtikel(c *gin.Context) {
	db, _ := Config.GetDB()
	var artikel Model.Artikel
	artikel.Title = c.Request.FormValue("title")
	artikel.Kontent = c.Request.FormValue("kontent")
	db.Create(&artikel)
	c.JSON(200, artikel)
}

func EditArtikel(c *gin.Context) {
	db, _ := Config.GetDB()
	var artikel Model.Artikel
	id := c.Params.ByName(`id`)
	if err := db.Where(`id = ?`, id).First(&artikel).Error; err != nil {
		c.AbortWithStatus(404)
	}
	artikel.Title = c.Request.FormValue("title")
	artikel.Kontent = c.Request.FormValue("kontent")
	db.Save(&artikel)
	c.JSON(200, artikel)
}

func DetailArtikel(c *gin.Context) {
	db, _ := Config.GetDB()
	var artikel Model.Artikel
	id := c.Params.ByName(`id`)
	if err := db.Where(`id = ?`, id).First(&artikel).Error; err != nil {
		c.AbortWithStatus(404)
	}
	db.Save(&artikel)
	c.JSON(200, artikel)
}

func DeleteArtikel(c *gin.Context) {
	db, _ := Config.GetDB()
	id := c.Params.ByName(`id`)
	var artikel Model.Artikel
	db.Where(`id = ?`, id).Delete(&artikel)
	c.JSON(200, gin.H{`id #` + id: `delete`})
}

func QueryUser(username string) Model.User {
	var users = Model.User{}
	db, _ := Config.GetDB()
	_ = db.Raw(`
		SELECT id, 
		username, 
		password 
		FROM users WHERE username=?
		`, username).
		Scan(
			&users,
		)
	return users
}

func Register(c *gin.Context) {
	db, _ := Config.GetDB()
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user := Model.User{Username: username, Password: string(bytes)}
	db.Create(&user)
	c.JSON(200, user)
}
