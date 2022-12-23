package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

var db *gorm.DB
var err error

type User struct { //创建一个用户表

	//用户序号
	UserId int `gorm:"column:userId;primaryKey;autoIncrement" json:"userId"`
	//用户名
	Name string `gorm:"column:name" json:"name"`
	//用户密码
	Password string `gorm:"column:password" json:"password"`
}

type Problems struct { //创建一个问题表

	//问题序号
	ProblemId int `gorm:"column:problemId;primaryKey;autoIncrement" json:"problemId"`
	//提问人用户名
	Name string `gorm:"column:name" json:"name"`
	//问题内容
	Problem string `gorm:"columbn:problem" json:"problem"`
}
type Answers struct { //创建一个评论表
	//评论序号
	AnswerId int `gorm:"column:answerId;primaryKey;autoIncrement" json:"answerId"`
	//问题序号
	ProblemId int `gorm:"column:problemId" json:"problemId"`
	//评论内容
	Answer string `gorm:"column:answer" json:"answer"`
	//评论人用户名
	Answerer string `gorm:"column:answerer" json:"answerer"`
}

func register(c *gin.Context) {
	var u User
	u.Name = c.Query("name")
	u.Password = c.Query("password")
	db.Create(&u)
	var u2 User
	db.Where("name=?", u2.Name).First(&u2)
	if len(u2.Name) == 0 {
		c.String(200, "无匹配用户")
	} else {
		c.String(http.StatusOK, "注册成功")
	}
}
func login(c *gin.Context) {
	var u1 User
	u1.Name = c.Query("name")
	u1.Password = c.Query("password")
	var u2 User
	db.Where("name=?", u1.Name).First(&u2)
	if len(u2.Name) == 0 {
		c.String(200, "无匹配用户")
	}
	if u1.Password == u2.Password {
		c.String(200, "登录成功")
	} else {
		c.String(200, "密码错误")
	}

}
func problem(c *gin.Context) {
	var p Problems

	p.Problem = c.Query("problem")
	p.Name = c.Query("name")
	db.Create(&p)
	var p1 Problems
	db.Where(" problemId=?", p.ProblemId).First(&p1)
	if p1.Problem == p.Problem {
		e := strconv.Itoa(p1.ProblemId)
		c.JSON(http.StatusOK, gin.H{
			e: p1.Problem,
		})
	} else {
		c.String(200, "提问失败")
	}

}

func answer(c *gin.Context) {
	var a Answers
	b := c.Query("problemId")
	a.ProblemId, _ = strconv.Atoi(b)
	a.Answer = c.Query("answer")
	a.Answerer = c.Query("answerer") //评论时输入自己的用户名
	db.Create(&a)
	var a1 Answers
	db.Where("problemId = ?", a.ProblemId).First(&a1)
	if a1.Answer == a.Answer {
		e := strconv.Itoa(a1.AnswerId)
		c.JSON(http.StatusOK, gin.H{
			e: a1.Answer,
		})
	} else {
		c.String(200, "回答失败")
	}
}
func findProblem(c *gin.Context) {
	var p Problems
	p.Name = c.Query("name")
	var p1 Problems
	db.Where("name = ?", p.Name).First(&p1)
	if len(p1.Problem) != 0 {
		a := strconv.Itoa(p1.ProblemId)
		c.JSON(http.StatusOK, gin.H{
			a: p1.Problem,
		})
	} else {
		c.String(http.StatusOK, "您没有提出任何问题")
	}
}
func findAnswer(c *gin.Context) {
	var a Answers
	a.Answerer = c.Query("answerer")
	var a1 Answers
	db.Where("answerer=?", a.Answerer).First(&a1)
	if len(a1.Answer) != 0 {
		e := strconv.Itoa(a1.ProblemId)
		c.JSON(http.StatusOK, gin.H{
			e: a1.Answer,
		})
	} else {
		c.String(http.StatusOK, "您没有回答任何问题")
	}
}
func updateProblem(c *gin.Context) {
	var p Problems
	b := c.Query("problemId")
	p.ProblemId, _ = strconv.Atoi(b)
	p.Problem = c.Query("problem") //更改后的问题
	db.Model(&Problems{}).Where("problemId = ?", p.ProblemId).Update("problem", p.Problem)
	e := strconv.Itoa(p.ProblemId)
	c.JSON(http.StatusOK, gin.H{
		e: p.Problem,
	})
}
func updateAnswer(c *gin.Context) {
	var a Answers
	b := c.Query("answerId")
	a.AnswerId, _ = strconv.Atoi(b)
	a.Answer = c.Query("answer")
	db.Model(&Answers{}).Where("answerId = ?", a.AnswerId).Update("answer", a.Answer)
	e := strconv.Itoa(a.AnswerId)
	c.JSON(http.StatusOK, gin.H{
		e: a.Answer,
	})
}
func deleteProblem(c *gin.Context) {
	var p Problems
	b := c.Query("problemId")
	p.ProblemId, _ = strconv.Atoi(b)
	var p2 Problems
	db.Where("problemId=?", p.ProblemId).First(&p2)
	if len(p2.Problem) == 0 {
		c.String(http.StatusOK, "找不到相关问题")
	} else {
		db.Where("problemId=?", p.ProblemId).Delete(&p)
		var p1 Problems
		db.Where("problemId=?", p.ProblemId).First(&p1)
		if len(p1.Problem) == 0 {
			c.String(http.StatusOK, "删除成功")
		} else {
			c.String(http.StatusOK, "删除失败")
		}
	}
}
func deleteAnswer(c *gin.Context) {
	var a Answers
	b := c.Query("answerId")
	a.AnswerId, _ = strconv.Atoi(b)
	var a2 Answers
	db.Where("answerId=?", a.AnswerId).First(&a2)
	if len(a2.Answer) == 0 {
		c.String(http.StatusOK, "找不到相关评论")
	} else {
		db.Where("answerId=?", a.AnswerId).Delete(&a)
		var a1 Answers
		db.Where("answerId=?", a.AnswerId).First(&a1)
		if len(a1.Answer) == 0 {
			c.String(http.StatusOK, "删除成功")
		} else {
			c.String(http.StatusOK, "删除失败")
		}
	}
}
func main() {
	dsn := "root:yrh200406@tcp(127.0.0.1:3306)/U?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//创建表，自动迁移（把结构体和表进行对应）
	err = db.AutoMigrate(&User{})
	if err != nil {
		return
	}
	if err != nil {
		log.Println(err)
	}
	err = db.AutoMigrate(&Problems{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&Answers{})
	if err != nil {
		return
	}

	r := gin.Default()
	r.GET("/register", register)
	r.GET("/login", login)
	r.GET("/problem", problem)
	r.GET("/answer", answer)
	r.GET("/findProblem", findProblem)
	r.GET("/findAnswer", findAnswer)
	r.GET("/updateProblem", updateProblem)
	r.GET("/updateAnswer", updateAnswer)
	r.GET("/deleteProblem", deleteProblem)
	r.GET("/deleteAnswer", deleteAnswer)
	err = r.Run()
	if err != nil {
		return
	}
}
