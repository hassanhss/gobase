package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("forms\\login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		//请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("username:", r.Form["username"][0])
		fmt.Println("password:", r.Form["password"][0])
		fmt.Println("password:", r.Form.Get("username"))
		if len(r.Form["username"][0]) == 0 {
			//为空的处理
		}
		strconv.Atoi(r.Form.Get("age"))
	}
}

func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
