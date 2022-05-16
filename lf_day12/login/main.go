package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"net/http"
)

var db *sqlx.DB

func main() {
	if err := initDB(); err != nil {
		fmt.Printf("初始化 数据库连接失败 err: %s\n", err)
		panic(err)
	}
	r := gin.Default()
	r.LoadHTMLFiles("./login.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	// 接收请求并处理请求
	r.POST("/login", loginHandler)
	// 启动服务
	r.Run(":8888")
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// loginHandler 处理登陆请求的函数
func loginHandler(c *gin.Context) {
	// 1. 从请求中获取用户的请求数据
	// form 表单提交, json格式提交
	var reqDate Login
	if err := c.ShouldBind(&reqDate); err != nil {
		fmt.Println(err)
		// 从请求里解析数据出错
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "请求参数错误",
		})
		return
	}
	fmt.Printf("reqDate: %#v\n", reqDate)
	//c.JSON(http.StatusOK, reqDate)
	fmt.Println("....")
	// 2. 对数据进行校验
	// 去数据校验
	if u, err := QueryUser(reqDate.Username, reqDate.Password); err == nil {
		// 登陆成功
		fmt.Println(u)
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "hello " + u.Username,
			"date": u,
		})
	} else {
		// 登陆失败
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "用户名或密码错误",
		})
	}
	// 3. 返回响应

}

func initDB() (err error) {
	dsn := "root:mysql57@tcp(192.168.31.24:3306)/login?charset=utf8mb4&parseTime=True"
	//dsn := "user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	return
}

type User struct {
	Id       int    `db:"id" json:"-"`
	Username string `db:"username" json:"name"`
	Desc     string `json:"desc,omitempty"`
}

func QueryUser(username, password string) (*User, error) {
	// 查询
	var sqlStr = "select id,username from user where username=? and password=?"
	var u User
	err := db.Get(&u, sqlStr, username, password)
	if err != nil {
		fmt.Println(errors.Is(err, sql.ErrNoRows)) // 没有查询到记录到情况
		fmt.Printf("get failed, err: %v\n", err)
		return nil, err
	}
	return &u, nil
}
