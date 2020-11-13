package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

/*
// 1. 使用 net/http 便准库进行web开发
func sayHello(w http.ResponseWriter, r *http.Request) {
	// 1.
	//_, _ = fmt.Fprintln(w, "<h1>Hello Web!</h1>")

	// 2. 把返回的数据放到单独的文件中
	bytes, _ := ioutil.ReadFile("./hello.txt")
	fmt.Fprintln(w, string(bytes))
}

func main() {
	http.HandleFunc("/hello", sayHello)
	//log.Fatal(http.ListenAndServe(":8080", nil))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}*/

/*
// 2. 使用 gin 框架进行web开发
func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Gin!",
	})
}

func main() {

	r := gin.Default() // 返回默认的路由引擎，也可以自定以路由引擎

	// 指定用户使用GET请求访问 /hello 时，执行sayHello这个函数
	r.GET("/hello", sayHello)

	// RESUful 风格
	// 1. 查询图书
	r.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})

	// 2. 创建图书记录
	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})

	// 3. 更新图书信息
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})

	// 4. 删除图书信息
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})

	// 启动
	_ = r.Run(":9090") // 默认的端口是8080

}*/

// net/http 使用模板开发

//type User struct {
//	Name   string
//	Gender string
//	Age    int
//}
//
//func main() {
//	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
//		// 2.解析模板
//
//		// 自定义函数  要么只有一个返回值，要么两个返回值第二个必须是error
//		//f1 := func(name string) (string, error) {
//		//	return "欢迎" + name + "使用自定义函数f1", nil
//		//}
//
//		//t, err := template.ParseFiles("./hello.tmpl")
//
//		//创建一个模板对象，名字是 f.tmpl，创建的名字一定是接下来解析的文件的名字，如果解析多个，改名字一定要是其中的一个
//		t := template.New("t.tmpl")
//		// 在解析模板之前，调用Func函数添加自定义的f1（功能是打招呼）函数
//		//t.Funcs(template.FuncMap{
//		//	"hello": f1,
//		//})
//		_, err := t.ParseFiles("./t.tmpl", "./ul.tmpl")
//		if err != nil {
//			fmt.Printf("Prase template failed, err:%v\n", err)
//			return
//		}
//
//		// 3.渲染模板
//		/*
//		u1 := User{
//			"小明",
//			"男",
//			18,
//		}
//		m1 := map[string]interface{}{
//			"name": "小红",
//			"gender": "女",
//			"age":18,
//		}
//
//		hobbyList := []string{
//			"篮球",
//			"足球",
//			"电影",
//		}
//		err = t.Execute(w, map[string]interface{}{
//			"u1": u1,
//			"m1": m1,
//			"hobby": hobbyList,
//		}) //把模板渲染之后发给响应w
//		 */
//		err = t.Execute(w, "小明")
//		if err != nil {
//			fmt.Printf("Rander template failed, err:%v\n", err)
//			return
//		}
//
//	})
//
//	err := http.ListenAndServe(":8080", nil)
//	if err != nil {
//		fmt.Printf("HTTP server start failed, err:%v\n", err)
//		return
//	}
//}

// block模板继承
/*
func index(w http.ResponseWriter, r *http.Request) {
	// 定义模板

	// 解析模板
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("Parse template dailed, err:%v", err)
	}
	// 渲染模板
	_ = t.Execute(w, "小明")
}

func home(w http.ResponseWriter, r *http.Request) {
	// 定义模板

	// 解析模板
	t := template.New("home.tmpl")
	_, err := t.ParseFiles("./home.tmpl")
	if err != nil {
		fmt.Printf("Parse template dailed, err:%v", err)
	}
	// 渲染模板
	_ = t.Execute(w, "小明")
}

func index2(w http.ResponseWriter, r *http.Request) {
	// 定义模板（继承模板的方式）
	// 解析模板
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/index.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed, err:%v", err)
	}
	// 渲染模板
	msg := "小明"
	err = t.ExecuteTemplate(w, "index.tmpl", msg)
	if err != nil {
		fmt.Printf("Execute template failed, err:%v", err)
	}
}

func home2(w http.ResponseWriter, r *http.Request) {
	// 定义模板（继承模板的方式）
	// 解析模板
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/home.tmpl")
	if err != nil {
		fmt.Printf("Prase templates failed, err:%v", err)
	}
	// 渲染模板
	msg := "小红"
	_ = t.ExecuteTemplate(w, "home.tmpl", msg)
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("HTTP start dailed, err:%v", err)
	}
}
*/

/*
// 修改默认标识符
func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// 定义模板
		// 解析模板
		//t := template.New("a.tmpl")
		//t.Delims("{[", "]}")
		//_, err := t.ParseFiles("./templates/a.tmpl")
		//if err != nil {
		//	fmt.Printf("Parse template failed, err:%v", err)
		//}

		// 解析之前定义一个safe函数
		t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
			"safe": func(str string) template.HTML {
				return template.HTML(str)
			},
		}).ParseFiles("./templates/xss.tmpl")
		if err != nil {
				fmt.Printf("Parse template failed, err:%v", err)
		}

		// 渲染模板

		xss := "<script></script>"
		safeStr := "<script>alert('这是被信任的脚本注入!')</script>"
		err = t.Execute(w, map[string]string{
			"xss": xss,
			"safeStr": safeStr,
		})
		if err != nil {
			fmt.Printf("Execute template failed, err:%v", err)
		}
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, err:%v", err)
	}
}

*/

// Gin

//type UserInfo struct {
//	// 常用 tag 记录
//	// form: gin框架中对应的前端的数据字段名
//	// json: json序列化或反序列化时字段的名称
//	// db: sqlx模块中对应的数据库字段名
//	// binding: 搭配 form 使用, 默认如果没查找到结构体中的某个字段则不报错值为空, binding为 required 代表没找到返回错误给前端
//	Username string `form:"username" json:"name"`
//	Password string `form:"password" json:"pwd"`
//}
//
//// 定义中间件
//func m1(c *gin.Context) {
//	fmt.Println("m1 in...")
//	start := time.Now() // 获取当前时间
//	c.Next()            // 调用后续的处理函数
//	//c.Abort() // 组织调用后续的处理函数
//	cost := time.Since(start) // start到现在的耗时
//	fmt.Printf("cost:%v\nm1 out...", cost)
//}
//
//func m2(c *gin.Context) {
//	fmt.Println("m2 in...")
//	// 跨中间件存取值
//	// 在这里做一些操作，传入一些参数，在m2后面执行的函数可以拿到该信息
//	c.Set("name", "小明")
//	c.Next()
//	fmt.Println("m2 out...")
//}
//
//// 一般利用闭包来创建中间件，这样可以在中间件逻辑开始之前做一些其他的事情，如下
//func authMiddleware(doCheck bool) gin.HandlerFunc {
//	// 连接数据库
//	// 做一些其他的准备工作
//	return func(c *gin.Context) {
//		if doCheck { // 中间件的开关，在注册的时候传入true才会开启检测
//			// 存放具体的逻辑
//			// 是否登录的判断
//			// if 是登录用户
//			// c.Next()
//			// else
//			// c.Abort()
//		} else {
//			c.Next()
//		}
//	}
//}
//
//// Gin使用模板
//func main() {
//	// 返回默认路由引擎
//	r := gin.Default()
//	//r = gin.New()
//
//	// 全局注册中间件   在该语句下面的路由函数才生效
//	//r.Use(m1, m2)
//
//	// 在解析模板之前加载静态文件
//	// 第一个参数是在html引入静态文件时的前缀，第二个参数是静态文件存放的位置
//	r.Static("/xxx", "./statics")
//
//	// 在模板解析之前添加自定义模板函数
//	r.SetFuncMap(template.FuncMap{
//		"safe": func(str string) template.HTML {
//			return template.HTML(str)
//		},
//	})
//
//	//r.LoadHTMLFiles("templates/posts/index.tmpl", "templates/users/index.tmpl") // 解析模板
//	r.LoadHTMLGlob("templates/**/*")
//	r.LoadHTMLFiles("templates/login.html", "templates/index.html",
//		"templates/formIndex.html", "templates/upload.html")
//	r.GET("/index", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{ // 渲染模板
//			"title": "Gin 框架渲染模板",
//		})
//	})
//
//	r.GET("users/index", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{ // 渲染模板
//			"title": "<a href='https://liwenzhou.com'>李文周的博客</a>",
//		})
//	})
//
//	// 返回JSON
//	r.GET("/json", func(c *gin.Context) {
//
//		// 1. 传入map
//		//data := map[string]interface{}{
//		//	"name": "小明",
//		//}
//		// gin.H 就是一个map[string]interface{}类型的
//		//c.JSON(http.StatusOK, data)
//
//		// 2.传入结构体
//		var msg struct {
//			Name    string `json:"user_name"`
//			Message string `json:"message"`
//			Age     int    `json:"age"`
//		}
//
//		msg.Name = "小红"
//		msg.Message = "hello world!"
//		msg.Age = 18
//
//		c.JSON(http.StatusOK, msg)
//	})
//
//	// 获取query string 参数
//	// GET请求 URL ?后面的是query string参数
//	// key=value格式，多个key-value用 & 连接
//	// eg: /web/query=小明&age=18
//	r.GET("/web", func(c *gin.Context) {
//		// 获取浏览器发请求带的 query string 参数
//		name := c.Query("query") // 获取 query 的值，获取不到返回空
//		age := c.Query("age")
//		//当查询不到key对应的value时返回空，当查询不到key时返回后面的默认值
//		//name := c.DefaultQuery("query", "defaultValue")
//
//		//取到返回值和true，取不到返回""和false bool类型的返回值我们通常用ok来接收
//		//name, ok := c.GetQuery("query")
//		//if !ok { // 取不到
//		//	name = "defaultValue"
//		//}
//
//		c.JSON(http.StatusOK, gin.H{
//			"name": name,
//			"age":  age,
//		})
//	})
//
//	// 获取form表单参数
//	r.GET("/login", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "login.html", nil)
//	})
//
//	// 表单提交时的请求处理
//	r.POST("/login", func(c *gin.Context) {
//		// 获取浏览器请求表单提交的数据  key就是input输入框的name,value就是输入框的值
//		//username := c.PostForm("username")
//		//password := c.PostForm("password")
//
//		// 使用DefaultPostForm时，输入框内不填写不是不存在，是会返回空
//		// 是当找不到第一个参数这个key时，是找不到，会显示默认值，即key和input的name不同时
//		//username := c.DefaultPostForm("username", "defaultValue-somebody")
//		//password := c.DefaultPostForm("password", "***")
//
//		// 这里同DefaultPostForm函数一样，找不到key时才不存在，此时会返回false，否则返回值(输入框不填写返回空)和true
//		username, ok := c.GetPostForm("username")
//		if !ok {
//			username = "somebody"
//		}
//		password, ok := c.GetPostForm("password")
//		if !ok {
//			password = "***"
//		}
//		c.HTML(http.StatusOK, "index.html", gin.H{
//			"name":     username,
//			"password": password,
//		})
//	})
//
//	// 获取URI路径参数
//	// 获取的path(URI)参数，返回的都是string类型的
//	// 注意URI的匹配不要冲突
//	r.GET("/test/:name/:age", func(c *gin.Context) {
//		username := c.Param("name")
//		age := c.Param("age")
//		c.JSON(http.StatusOK, gin.H{
//			"username": username,
//			"age":      age,
//		})
//	})
//
//	// 这里使用时不能直接使用 /:year/:month   这样会和上面的冲突，会把test和:year匹配 month和:name匹配
//	// 所以要在前面进行区分开加个/blog
//	r.GET("/blog/:year/:month", func(c *gin.Context) {
//		year := c.Param("year") // 在路径里写的是数字，但是返回的是string
//		month := c.Param("month")
//		c.JSON(http.StatusOK, gin.H{
//			"year":  year,
//			"month": month,
//		})
//	})
//
//	// 使用参数绑定
//	// 绑定query string参数
//	r.GET("/user", func(c *gin.Context) {
//		// 创建结构体接收的形式
//		//username := c.Query("username")
//		//password := c.Query("password")
//		//u := UserInfo{
//		//	username,
//		//	password,
//		//}
//		//fmt.Printf("%#v\n", u)
//
//		// 使用参数绑定的方式
//		var u UserInfo
//		err := c.ShouldBind(&u)
//		if err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"error": err.Error(),
//			})
//		} else {
//			fmt.Printf("%#v\n", u)
//			c.JSON(http.StatusOK, gin.H{
//				"status": "ok",
//			})
//		}
//
//	})
//	// 绑定form参数
//	r.GET("/formIndex", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "formIndex.html", nil)
//	})
//	r.POST("/form", func(c *gin.Context) {
//		var u UserInfo
//		err := c.ShouldBind(&u)
//		if err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"error": err.Error(),
//			})
//		} else {
//			fmt.Printf("%#v\n", u)
//			c.JSON(http.StatusOK, gin.H{
//				"status": "ok",
//			})
//		}
//	})
//	// 绑定json参数
//	r.POST("/json", func(c *gin.Context) {
//		var u UserInfo
//		err := c.ShouldBind(&u)
//		if err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"error": err.Error(),
//			})
//		} else {
//			fmt.Printf("%#v\n", u)
//			c.JSON(http.StatusOK, gin.H{
//				"status": "ok",
//			})
//		}
//	})
//
//	// 文件上传
//	// 单文件上传
//	r.GET("/upload", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "upload.html", nil)
//	})
//	r.POST("/upload", func(c *gin.Context) {
//		// 从请求中读取文件
//		f1, err := c.FormFile("f1") // form 中 name为f1的
//		if err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{
//				"error": err.Error(),
//			})
//			return
//		} else {
//			// 将读取的文件保存到本地（服务端）
//			dst := fmt.Sprintf("./%s", f1.Filename)
//			//dst = path.Join("./", f1.Filename)
//			err := c.SaveUploadedFile(f1, dst)
//			if err != nil {
//				fmt.Printf("Save file failed, err:%v", err)
//				return
//			}
//			c.JSON(http.StatusOK, gin.H{
//				"status": "ok",
//			})
//		}
//
//	})
//	//多文件上传
//
//	//重定向
//	// HTTP重定向
//	r.GET("/redirect", func(c *gin.Context) {
//		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
//	})
//	// 路由重定向
//	r.GET("/a", func(c *gin.Context) {
//		c.Request.URL.Path = "/b" // 把请求的URI修改
//		r.HandleContext(c)        // 继续后续的处理
//	})
//	r.GET("/b", func(c *gin.Context) {
//		c.JSON(http.StatusOK, gin.H{
//			"msg": "hello 欢迎重定向路由到/b",
//		})
//	})
//	// 路由组的使用
//	// 为路由组注册中间件
//	// 写法1.
//	//shopGroup := r.Group("/shop", m1, m2)
//	// 写法2.
//	shopGroup := r.Group("/shop")
//	shopGroup.Use(m1, m2)
//	{
//		shopGroup.GET("/home", func(c *gin.Context) {
//			c.JSON(http.StatusOK, gin.H{
//				"msg": "首页",
//			})
//		})
//		shopGroup.GET("/aaa", func(c *gin.Context) {
//			c.JSON(http.StatusOK, gin.H{
//				"msg": "aaa页面",
//			})
//		})
//
//		// 嵌套路由组
//		userGroup := shopGroup.Group("/user")
//		{
//			userGroup.GET("/car", func(c *gin.Context) {
//				c.JSON(http.StatusOK, gin.H{
//					"msg": "您的购物车页面 /user/car",
//				})
//			})
//		}
//	}
//
//	// 中间件
//	// 为/middleware路由单独注册
//	r.GET("/middleware", m1, m2, func(c *gin.Context) {
//		name, _ := c.Get("name")
//		c.JSON(http.StatusOK, gin.H{
//			"status":  "ok",
//			"m2-name": name,
//		})
//	})
//	// 定义中间件 在上面定义
//	// 注册中间件
//	// 全局注册 在该语句下面才生效
//	//r.Use(m1, m2)
//
//	// 启动server
//	err := r.Run(":8080")
//	if err != nil {
//		fmt.Printf("HTTP server start failed, err:%v", err)
//	}
//}

// GORM
// UserInfo ---> 数据表
type UserInfo struct {
	Id     uint
	Name   string
	Gender string
	Hobby  string
}

// 1.定义模型（Model）
type MyUser struct {
	gorm.Model        // 嵌套 ID、CreatedAt、UpdatedAt、DeletedAt
	Name       string `gorm:"default:'新用户QAQ'"`
	Age        *int64 `gorm:"default:18"`
	Birthday   time.Time
	Active     bool
}

func main() {
	// 2.连接MySQL数据库
	db, err := gorm.Open("mysql",
		"root:644315@(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 3.创建表 自动迁移（把结构体和数据表进行对应）
	//db.AutoMigrate(&UserInfo{})

	// new和make的区别
	// new返回的是所创建变量的指针类型
	// make返回的是所创建值的类型，make一般用来创建数组、切片、map
	db.AutoMigrate(&MyUser{})

	// 创建记录
	//u1 := UserInfo{
	//	1,
	//	"小明",
	//	"男",
	//	"写代码",
	//}
	//db.Create(u1)

	//// 查询
	//var u UserInfo
	//db.First(&u) // 查询第一条数据保存到u中
	//fmt.Printf("u:%#v\n", u)
	//
	//// 更新
	//db.Model(&u).Update("hobby", "看电影") // 更新u这条记录中的hobby为 看电影
	//
	//// 删除
	//db.Delete(&u) // 删除u这条记录

	// 创建记录
	//age := new(int64)
	//*age = 18
	//u := MyUser{ // 在代码层面创建一个结构体
	//	Name: "小明",
	//	Age: age,
	//	Birthday: time.Now(),
	//	Active: true,
	//}
	//u := MyUser{
	//	Name: "小红",
	//}
	//fmt.Println(db.NewRecord(&u)) // 判断主键是否为空（就是查看一下是不是有这条记录了），这时返回true
	//db.Create(&u) //传入一个结构体指针，因为有的结构体比较大，传入结构体变量拷贝会花费时间相对较多
	//fmt.Println(db.NewRecord(&u)) // 创建之后再查询，此时返回false

	// 查询记录

	//// 一般查询
	//var user1 MyUser // 声明表所对应的结构体类型变量user，来接收查询的结果
	//// user := new(MyUser) // new返回的是对应类型的指针，所以直接传递user进去就行，make返回的就是对应的类型
	//
	//// 根据主键查询第一条记录 Debug() 会打印执行的sql语句
	//db.Debug().First(&user1) // 这里要传递指针变量，因为传递至拷贝一份是没办法给上面我们声明的user赋值的
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL ORDER BY `my_users`.`id` ASC LIMIT 1;
	//fmt.Printf("\nFirst user:%#v\n", user1)
	//
	//// 随机获取一条数据
	//var user2 MyUser
	//db.Debug().Take(&user2)
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL LIMIT 1;
	//fmt.Printf("\nTake user:%#v\n", user2)
	//
	//// 根据主键查询最后提条记录
	//var user3 MyUser
	//db.Debug().Last(&user3)
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL ORDER BY `my_users`.`id` DESC LIMIT 1;
	//fmt.Printf("\nLast user:%#v\n", user3)
	//
	//// 查询所有记录
	//var users []MyUser
	//db.Debug().Find(&users)
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL;
	//fmt.Printf("\nusers:%#v\n", users)
	//
	//// 查询指定的某条记录（仅当主键为整型时可用）
	//var user4 MyUser
	//db.Debug().First(&user4, 4)
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL AND ((`my_users`.`id` = 4)) ORDER BY `my_ users`.`id` ASC LIMIT 1;
	//fmt.Printf("\nUser ID = 4:%#v\n", user4)

	// Where 条件

	// 普通sql
	//var user MyUser
	//var users []MyUser
	//// 获取第一个匹配的记录
	//db.Debug().Where("name=?", "小明").First(&user)
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL AND ((name='小明')) ORDER BY `my_users`. `id` ASC LIMIT 1
	//fmt.Printf("%#v\n", user)
	//
	//// 获取所有匹配的记录
	//db.Debug().Where("age=?", 18).Find(&users)
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL AND ((age=18))
	//fmt.Printf("%#v\n", users)
	//
	//// <> 获取所有比等于指定值的记录
	//db.Debug().Where("age<>?", 18).Find(&users)
	//// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL AND ((age<>18))
	//fmt.Printf("%#v\n", users)
	//
	//// IN
	//db.Debug().Where("name IN (?)", []string{"TheShy", "小红", "张三"}).Find(&users)
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL AND ((name IN ('TheShy','小红','张三')))
	//fmt.Printf("%#v\n", users)
	//
	//// LIKE
	//db.Debug().Where("name LIKE ?", "%小%").Find(&users)
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL AND ((name LIKE '%小%'))
	//fmt.Printf("%#v\n", users)
	//
	//// AND
	//db.Debug().Where("age = ? AND name IN (?)", 18, []string{"小明", "小红", "TheShy"}).Find(&users)
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL AND ((age = 18 AND name IN ('小明','小红' ,'TheShy')))
	//fmt.Printf("%#v\n", users)
	//
	//// Time
	//db.Debug().Where("updated_at < ?", time.Now()).Find(&users)
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL AND ((updated_at < '2020-07-24 23:27:16')
	//fmt.Printf("%#v\n", users)
	//
	//// BETWEEN
	//currentTime := time.Now()
	//oldTime := currentTime.AddDate(0, 0, -7)
	//db.Debug().Where("updated_at BETWEEN ? AND ?", oldTime, currentTime).Find(&users)
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL AND ((updated_at BETWEEN '2020-07-17 23:3 3:40' AND '2020-07-24 23:33:40'))
	//fmt.Printf("%#v\n", users)

	// Struct
	//age := new(int64)
	//*age = 18
	//db.Debug().Where(&MyUser{Name: "小明", Age: age}).Find(&user)
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL AND ((`my_users`.`name` = '小明') AND (` my_users`.`age` = 18))
	//fmt.Printf("%#v\n", user)
	//
	//// Map 注意这里的key是数据库中的列名
	//db.Debug().Where(map[string]interface{}{"name": "TheShy","age": 18}).Find(&users)
	//////  SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL AND ((`my_users`.`age` = 18) AND (`my_use rs`.`name` = 'TheShy'))
	//fmt.Printf("%#v\n", users)
	//
	//// 主键切片
	//db.Debug().Where([]int64{1,2,3}).Find(&users)
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL AND ((`my_users`.`id` IN (1,2,3)))
	//fmt.Printf("%#v\n", users)

	// NOT 条件 作用与Where类似
	//var user MyUser
	//var users []MyUser
	//
	//// 获取第一条不符合的记录
	//db.Debug().Not("name", "TheShy").First(&user)
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL AND ((`my_users`.`name` NOT IN ('TheShy' ))) ORDER BY `my_users`.`id` ASC LIMIT 1
	//fmt.Printf("%#v\n", user)
	//
	//// Not In
	//db.Debug().Not("name", []string{"小明", "TheShy"}).Find(&users)
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL AND ((`my_users`.`name` NOT IN ('小明','T heShy')))
	//fmt.Printf("%#v\n", users)
	//
	//// 不在主键切片中
	//db.Debug().Not([]int64{1,2}).Find(&users)
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL AND ((`my_users`.`id` NOT IN (1,2)))
	//fmt.Printf("%#v\n", users)
	//
	//db.Debug().Not([]int64{}).First(&user)
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL AND `my_users`.`id` = 1 ORDER BY `my_user s`.`id` ASC LIMIT 1
	//fmt.Printf("%#v\n", user)
	//
	//// 普通sql
	//db.Debug().Not("name = ?", "新用户QAQ").First(&user)
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL AND (NOT (name = '新用户QAQ')) ORDER BY `my_users`.`id` ASC LIMIT 1
	//fmt.Printf("%#v\n", user)
	//
	//// struct
	//db.Debug().Not(MyUser{Name: "小明"}).Find(&users)
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL AND ((`my_users`.`name` <> '小明'))
	//fmt.Printf("%#v\n", users)

	// Or 条件
	//var users []MyUser
	//
	//db.Debug().Where("age = ?", 18).Or("age = ?", 0).Find(&users)
	//////  SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL AND ((age = 18) OR (age = 0))
	//fmt.Printf("%#v\n", users)
	//
	//// Struct
	//age := new(int64)
	//*age = 18
	//db.Debug().Where("name = ?", "小红").Or(MyUser{Age: age}).Find(&users)
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL AND ((name = '小红') OR (`my_users`.`age` = 18))
	//fmt.Printf("%#v\n", users)
	//
	//// Map
	//db.Debug().Where("name = ?", "小明").Or(map[string]interface{}{"age":age}).Find(&users)
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL AND ((name = '小明') OR (`my_users`.`age` = 18))
	//fmt.Printf("%#v\n", users)

	// Inline Condition 内联条件
	// 当内联条件与 多个立即执行方法 一起使用时, 内联条件不会传递给后面的立即执行方法。
	//db.Debug().First(&user, 2)
	//db.Debug().First(&user, "age=?", 18)
	//
	//db.Debug().Find(&users, "age = ?", 18)

	// FirstOrInit
	// 获取匹配的第一条记录，否则根据给定的条件初始化一个新的对象 (仅支持 struct 和 map 条件)
	// 未找到，会给把查询条件赋值给user
	//db.Debug().FirstOrInit(&user, MyUser{Name: "FirstOrInit"})
	//fmt.Printf("%#v", user)
	// 找到，找到的记录赋值给user
	//db.Where("name = ?", "小明").FirstOrInit(&user)
	//db.Attrs(MyUser{Name: "hello"}).FirstOrInit(&user, "name = ?", "小明")
	//fmt.Printf("%#v", user)
	////

	// 高级查询
	// Select，指定你想从数据库中检索出的字段，默认会选择全部字段。

	//var users []MyUser
	//db.Debug().Select("name, age").Find(&users)
	////// SELECT name, age FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL
	//fmt.Printf("%#v\n", users)
	//
	//db.Debug().Select([]string{"name", "age"}).Find(&users)
	////// SELECT name, age FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL
	//fmt.Printf("%#v\n", users)
	//
	//db.Debug().Table("my_users").Select("COALESCE(age,?)", 18).Find(&users)
	////// SELECT COALESCE(age,18) FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL
	//fmt.Printf("%#v\n", users)

	// 排序
	// Order，指定从数据库中检索出记录的顺序。设置第二个参数 reorder 为 true ，可以覆盖前面定义的排序条件。

	//var users []MyUser
	//db.Debug().Order("age DESC, name").Find(&users)
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL ORDER BY age DESC, name
	//fmt.Printf("%#v\n", users)
	//
	//db.Debug().Order("age DESC").Order("name").Find(&users)
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL ORDER BY age DESC, name
	//fmt.Printf("%#v\n\n\n", users)
	//
	//var users2 []MyUser
	//db.Debug().Order("age DESC").Find(&users).Order("name").Find(&users2)
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL ORDER BY age DESC
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL ORDER BY age DESC,`name`
	//fmt.Printf("%#v\n%#v\n", users, users2)

	// 数量 Limit，指定从数据库检索出的最大记录数。
	//var users1 []MyUser
	//var users2 []MyUser
	//db.Debug().Limit(2).Find(&users1)
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL LIMIT 2
	//fmt.Printf("%#v\n", users1)
	//
	//// -1 取消 Limit 条件
	//db.Debug().Limit(2).Find(&users1).Limit(-1).Find(&users2)
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL LIMIT 2
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL
	//fmt.Printf("%#v\n%#v\n", users1, users2)

	// 偏移 Offset，指定开始返回记录前要跳过的记录数。
	//db.Debug().Limit(2).Offset(1).Find(&users1)
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL LIMIT 2 OFFSET 1
	//fmt.Printf("%#v\n", users1)
	//
	//// -1 取消 Offset 条件
	//db.Debug().Limit(2).Offset(1).Find(&users1).Offset(-1).Find(&users2)
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL LIMIT 2 OFFSET 1
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL LIMIT 2
	//fmt.Printf("%#v\n%#v\n", users1, users2)

	// 总数 Count，该 model 能获取的记录总数。
	//var count int
	//db.Debug().Where("id > ?", 2).Find(&users1).Count(&count)
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL AND ((id > 2))
	////// SELECT count(*) FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL AND ((id > 2))
	//fmt.Printf("%d", count)
	//
	//db.Debug().Model(&MyUser{}).Where("age = ?", 18).Count(&count)
	////// SELECT count(*) FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL AND ((age = 18))
	//fmt.Printf("%d\n",count)
	//
	//db.Debug().Table("my_users").Count(&count)
	////// SELECT count(*) FROM `my_users`
	//fmt.Printf("%#v\n", count)
	//
	//// 非重复总数
	//db.Debug().Table("my_users").Select("count(distinct(name))").Count(&count)
	////// SELECT count(distinct(name)) FROM `my_users`
	//fmt.Printf("%#v\n", count)

	// Pluck 查询 model 中的一个列作为切片，如果您想要查询多个列，您应该使用 Scan
	//var ages []int64
	//var users []MyUser
	//db.Debug().Find(&users).Pluck("age",&ages)
	////// SELECT age FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL
	//fmt.Printf("%#v\n ages:%v\n", users, ages)
	//
	//var names []string
	//db.Debug().Model(&MyUser{}).Pluck("name", &names)
	////// SELECT name FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL
	//fmt.Printf("%#v\n", names)
	//
	//var birthdays []time.Time
	//db.Table("my_users").Pluck("birthday", &birthdays)
	//fmt.Printf("%#v\n", birthdays[2].Format("2006/01/02"))

	// 扫面 Scan，扫描结果至一个 struct
	//type Result struct {
	//	Name string
	//	Age int64
	//}
	//var result Result
	//db.Debug().Table("my_users").Select("name, age").Where("name = ?", "TheShy").Scan(&result)
	////// SELECT name, age FROM `my_users`  WHERE (name = 'TheShy')
	//fmt.Printf("%#v\n", result)
	//
	//// 原生sql
	//db.Debug().Raw("SELECT name, age FROM my_users WHERE name = ?", "小明").Scan(&result)
	//fmt.Printf("%#v\n", result)

	//// 更新
	//var user MyUser
	//db.FirstOrInit(&user, MyUser{Name: "小公主"})
	//fmt.Printf("%d\n", user.ID)

	//user.Name = "小王子"
	//*user.Age = 18
	// Save会更新所有字段，即使你没有赋值
	//db.Debug().Save(&user)
	//// UPDATE `my_users` SET `created_at` = '2020-07-24 19:43:41', `updated_at` = '2020-07-26 14:18:04', `delete
	//d_at` = NULL, `name` = '小王子', `age` = 18, `birthday` = '0000-00-00 00:00:00', `active` = false  WHERE `my_users`.`deleted_at` IS NULL AND `my_users`.`id` =
	// 1

	// 如果你只希望更新指定字段，可以使用Update或者Updates
	// 更新单个属性，如果它有变化
	//db.Debug().Model(&user).Update("name", "小明")
	//// UPDATE `my_users` SET `name` = '小明', `updated_at` = '2020-07-26 14:21:53'  WHERE `my_users`.`deleted_at
	//` IS NULL AND `my_users`.`id` = 1

	// 根据给定的条件更新单个属性
	//db.Debug().Model(&user).Where("active = ?", false).Update("name", "小红")
	//// UPDATE `my_users` SET `name` = '小红', `updated_at` = '2020-07-26 14:27:26'  WHERE `my_users`.`deleted_at
	//` IS NULL AND `my_users`.`id` = 2 AND ((active = false))

	// 使用 map 更新多个属性，只会更新其中有变化的属性
	//db.Debug().Model(&user).Updates(map[string]interface{}{"name":"哆啦A梦", "age":18, "active":true})
	//// UPDATE `my_users` SET `active` = true, `age` = 18, `name` = '哆啦A梦', `updated_at` = '2020-07-26 14:30:1
	//2'  WHERE `my_users`.`deleted_at` IS NULL AND `my_users`.`id` = 3

	// 使用 struct 更新多个属性，只会更新其中有变化且为非零值的字段
	//a :=  new(int64)
	//*a = 15
	//db.Debug().Model(&user).Updates(MyUser{Name: "小王子", Age: a, Active: true})
	//// UPDATE `my_users` SET `active` = true, `age` = 15, `name` = '小王子', `updated_at` = '2020-07-26 14:35:34
	//'  WHERE `my_users`.`deleted_at` IS NULL AND `my_users`.`id` = 4

	// 警告：当使用 struct 更新时，GORM只会更新那些非零值的字段
	// 对于下面的操作，不会发生任何更新，"", 0, false 都是其类型的零值
	// 但是我们这里Age字段是指针类型的变量，所以会更新
	//*a = 0
	//db.Debug().Model(&user).Updates(MyUser{Name: "", Age: a, Active: false})

	// 更新选定字段 如果你想更新或忽略某些字段，你可以使用 Select，Omit
	//db.Debug().Model(&user).Select("age").Updates(map[string]interface{}{
	//	"name": "小公主", "age": 17, "active": false,
	//})
	//// UPDATE `my_users` SET `age` = 17, `updated_at` = '2020-07-26 15:11:27'  WHERE `my_users`.`deleted_at` IS
	//NULL AND `my_users`.`id` = 4

	//db.Debug().Model(&user).Omit("age").Updates(map[string]interface{}{
	//	"name": "小公主",
	//	"age": 15,
	//	"active": false,
	//})
	//// UPDATE `my_users` SET `active` = false, `name` = '小公主', `updated_at` = '2020-07-26 15:16:23'  WHERE `m
	//y_users`.`deleted_at` IS NULL AND `my_users`.`id` = 4

	// 无HOOKS更新 上面的更新操作会自动运行 model 的 BeforeUpdate, AfterUpdate 方法
	// 更新 UpdatedAt 时间戳, 在更新时保存其 Associations
	// 如果你不想调用这些方法，你可以使用 UpdateColumn， UpdateColumns
	// 更新单个属性，类似于 `Update`
	//db.Debug().Model(&user).UpdateColumn("name", "小公主")
	//// UPDATE `my_users` SET `name` = '小公主'  WHERE `my_users`.`deleted_at` IS NULL AND `my_users`.`id` = 1

	// 更新多个属性，类似于 `Updates`
	//db.Debug().Model(&user).UpdateColumns(map[string]interface{}{
	//	"name": "小红",
	//	"age": 18,
	//})
	//// UPDATE `my_users` SET `age` = 18, `name` = '小红'  WHERE `my_users`.`deleted_at` IS NULL AND `my_users`.`
	//id` = 2

	// 批量更新 批量更新时 Hooks 不会运行
	//db.Debug().Table("my_users").Where("id IN (?)", []int{2, 3}).Updates(map[string]interface{}{
	//	"age": 20,
	//})
	////  UPDATE `my_users` SET `age` = 20  WHERE (id IN (2,3))

	//a := new(int64)
	//*a = 21
	// 使用 struct 更新时，只会更新非零值字段，若想更新所有字段，请使用map[string]interface{}
	// 但是这里hooks运行了
	//db.Debug().Model(MyUser{}).Updates(MyUser{Age: a})
	//// UPDATE `my_users` SET `age` = 21, `updated_at` = '2020-07-26 15:37:16'  WHERE `my_users`.`deleted_at` IS NULL
	// 使用 `RowsAffected` 获取更新记录总数
	//rowsAffected := db.Debug().Model(MyUser{}).Updates(MyUser{Age: a}).RowsAffected
	//fmt.Printf("%d\n", rowsAffected)

	// 使用 SQL 表达式更新
	//db.Debug().Model(&user).Update("age", gorm.Expr("age - ?", 3))
	//// UPDATE `my_users` SET `age` = age - 3, `updated_at` = '2020-07-26 16:37:34'  WHERE `my_users`.`deleted_at
	//` IS NULL AND `my_users`.`id` = 3

	//db.Debug().Model(&user).Updates(map[string]interface{}{"age": gorm.Expr("age + ?", 3)})
	//// UPDATE `my_users` SET `age` = age + 3, `updated_at` = '2020-07-26 16:39:57'  WHERE `my_users`.`deleted_at
	//` IS NULL AND `my_users`.`id` = 3

	//db.Debug().Model(&user).UpdateColumn("age", gorm.Expr("age - ?", 3))
	//// UPDATE `my_users` SET `age` = age - 3  WHERE `my_users`.`deleted_at` IS NULL AND `my_users`.`id` = 3

	//db.Debug().Model(&user).Where("age > 18").UpdateColumn("age", gorm.Expr("age - ?", 3))
	//// UPDATE `my_users` SET `age` = age - 3  WHERE `my_users`.`deleted_at` IS NULL AND `my_users`.`id` = 1 AND ((age > 18))

	// 删除
	//var user MyUser
	//db.FirstOrInit(&user, MyUser{Name: "小公主"})
	//fmt.Printf("%d\n", user.ID)

	//db.Debug().Delete(&user)
	//// UPDATE `my_users` SET `deleted_at`='2020-07-27 02:11:26'  WHERE `my_users`.`deleted_at` IS NULL AND `my_users`.`id` = 1

	// 批量删除 删除全部匹配的记录
	//db.Debug().Where("name LIKE ?", "小王%").Delete(MyUser{})
	//// UPDATE `my_users` SET `deleted_at`='2020-07-27 09:19:55'  WHERE `my_users`.`deleted_at` IS NULL AND ((na
	//me LIKE '小王%'))
	//db.Debug().Delete(MyUser{}, "name IN (?)", []string{"小明", "小王子", "小红"})
	//// UPDATE `my_users` SET `deleted_at`='2020-07-27 09:25:56'  WHERE `my_users`.`deleted_at` IS NULL AND ((nam
	//e IN ('小明','小王子','小红')))

	// 软删除 如果一个 model 有 DeletedAt 字段，他将自动获得软删除的功能！ 当
	// 调用 Delete 方法时， 记录不会真正的从数据库中被删除， 只会将DeletedAt 字段的值会被设置为当前时间

	// 查询时会忽略软删除的记录
	//var users []MyUser
	//db.Debug().Where("age = ?", 18).Find(&users)
	////// SELECT * FROM `my_users`  WHERE `my_users`.`deleted_at` IS NULL AND ((age = 18))
	//fmt.Printf("%#v\n", users)

	// Unscoped 方法可以查询被软删除的记录
	//db.Debug().Unscoped().Where("age = ?", 18).Find(&users)
	////// SELECT * FROM `my_users`  WHERE (age = 18)
	//fmt.Printf("%#v\n", users)

	// 物理删除
	// Unscoped 方法可以物理删除记录
	//db.Debug().Unscoped().Delete(&user)
	//// DELETE FROM `my_users`  WHERE `my_users`.`id` = 7
}
