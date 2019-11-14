package main

import (
	"os"
    "./service"
	flag "github.com/spf13/pflag"
)

const (
	PORT string = "8080" //设置默认端口为8080
)

func main() {
	//如果没有设置端口，则启用默认端口8080
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	//用户使用-p xxxx来设置访问页面的端口
	pPort := flag.StringP("port", "p", "PORT", "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}
	
	//运行
	service.NewServer(port) 
}