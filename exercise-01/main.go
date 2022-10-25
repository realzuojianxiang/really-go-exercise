package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

type admin struct {
	person user
	level  string
}

func (u *user) notify() {
	fmt.Printf("Sending User Email To %s<%s> \n", u.name, u.email)
}

func (u *user) changEmail(email string) {
	u.email = email
}

type IP []byte
type IPAddr struct {
	IP IP
}

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip.IP[0], ip.IP[1], ip.IP[2], ip.IP[3])
}

type notifier interface {
	notify()
	changEmail(string)
}

func main() {
	var bill user
	bill.name = "Bill"
	bill.email = ""
	bill.ext = 123
	bill.privileged = true

	bill.notify()

	lisa := user{
		name:       "Lisa",
		email:      "",
		ext:        123,
		privileged: true,
	}

	fmt.Println(bill)
	fmt.Println(lisa)

	// 通过嵌入类型的方式，将user类型的值赋值给admin类型的值
	fred := admin{
		person: user{
			name:       "Fred",
			email:      "",
			ext:        123,
			privileged: true,
		},
		level: "super",
	}

	fmt.Println(fred)

	duration := time.Duration(1)
	format := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(duration)
	fmt.Println(format)

	//当前文件路径
	dir, _ := os.Getwd()
	fmt.Println(dir)
	fileName := dir + "exercise-01\\test.txt"
	fmt.Println(fileName)
	//获打开本地txt文件
	file, _ := os.Open("exercise-01\\test.txt")

	//获得文件信息
	info, _ := file.Stat()
	//获得文件大小
	size := info.Size()
	//获得文件名
	name := info.Name()
	//获得文件修改时间
	modTime := info.ModTime()
	fmt.Println(size)
	fmt.Println(name)
	fmt.Println(modTime)

	//关闭文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	resp, err := http.Get("https://www.sina.com.cn")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(resp *http.Response) {
		err := resp.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp)

	//获得响应状态码
	status := resp.StatusCode
	//获得响应头
	header := resp.Header
	//获得响应体
	body := resp.Body
	s, _ := io.ReadAll(body)

	fmt.Println("status:", status)
	fmt.Println("header:", header)
	fmt.Println(string(s))

	//写入文件
	err = os.WriteFile("exercise-01\\test.txt", s, 0666)

	n := notifier(&bill)
	m := notifier(&lisa)
	n.notify()
	n.changEmail("1@1.com")
	m.notify()
	m.changEmail("2@2.com")

	sendNotification(&bill)

}
func sendNotification(n notifier) {
	n.notify()
}
