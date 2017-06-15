# 项目：知乎xx API

已实现功能： 

1. 通过单个问题id获取批量答案
2. 通过集合id获取批量问题后获取批量答案
3. 关注别人（风险大容易被封杀去除,xxxx）
4. 登录(验证码问题去除,xxxx)

待实现功能：

1. 通过答案id获取单个回答
2. 根据用户唯一域名id获取她（它）他的全部回答（有用，优先级高）
3. 根据用户唯一域名id获取其关注的人，和关注她的人

## 一.小白指南

Golang开发的爬虫，小白用户请下载`main`文件夹下的`zhihu_windows_amd64.exe`，并在同一目录下新建一个`cookie.txt`文件，

打开火狐浏览器后人工登录知乎，按F12，点击网络，刷新一下首页，然后点击第一个出现的`GET /`，找到消息头请求头，复制Cookie，然后粘贴到cookie.txt

![](data/cookie.png)

点击EXE后,可选JS解决防盗链（这个是你要发布到自己的网站如：[减肥成功是什么感觉？给生活带来哪些改变？](http://www.lenggirl.com/zhihu/26613082-html/1.html)）
我们自己本地看的话就不要选择防盗链了！回答个数已经限制不大于500个。如果没有答案证明Cookie失效，请重新按照上述方法手动修改`cookie.txt`。

单问题模式：

```bash
zhihu_linux_x86_64 

        -----------------
        知乎问题信息小助手
        功能:
        1. 可选抓取图片
        2. 抓取答案
        3. 可选关注小伙伴

        选项:
        1. 从收藏夹https://www.zhihu.com/collection/78172986批量获取很多问题答案
        2. 从问题https://www.zhihu.com/question/28853910批量获取一个问题很多答案

        请您按提示操作（Enter）！答案保存在data文件夹下！

        因为知乎防盗链，放在你的网站上是看不见图片的！
        但是本地查看是没问题的！可选择防盗链生成HTML

        如果什么都没抓到请往exe同级目录cookie.txt
        增加cookie，手动增加cookie见说明

        你亲爱的萌萌~
        太阳萌飞了~~~
        -----------------
        
萌萌：你要发布到自己的网站上吗(JS解决防盗链)Y/N(默认N)
y
萌萌：要抓取图片吗Y/N(默认N)
n
萌萌：从收藏夹获取按1，从问题获取按2(默认)
2
萌萌说亲爱的，因为回答实在太多，请限制获取的回答个数:30（默认)
499
萌萌：请输入问题ID:
57000057
预抓取第一个回答！
开始处理答案:文末更新了！ 啊我被你们叫小姐姐叫的心都化了！ -- 我174cm 49kg 题主你这个身材超招人羡慕了好吗！ 个子高走什么风格不重要，主要是要简单，简单，简单。 以大面积纯色为主，过多的花纹和图案都会让觉得“巨婴”“傻大个”。款式也是越简单越好，有一些设…
哦，这个问题是:个子较高的女生怎么穿搭？
保存答案成功:data/57000057/chen-jian-guo-he-li-zi-165635365/chen-jian-guo-he-li-zi-165635365的回答.html
批量抓取答案，默认N(Y/N)
y
开始处理答案:5.10号 微博：Chilli-M 这么久了还没沉底儿既然这样大噶多多点赞关注好不啦～ 你们看看排我后头的赞都比我多两倍带拐弯儿！ 争气啊朋友们！ 实在没什么穿搭发 发最近的一点日常 …………………………………………………更新～～ ～～～～～～～～～～～～～…
保存答案成功:data/57000057/ma-tian-jiao-92-155865780/ma-tian-jiao-92-155865780的回答.html

```

上帝模式

```bash
萌萌：从收藏夹获取按1，从问题获取按2(默认)
1
萌萌说亲爱的，因为回答实在太多，请限制获取的回答个数:30（默认)
499
萌萌：请输入集合ID:
78172986
开启上帝模式吗(一路抓到底)，默认N(Y/N)?
y
抓取收藏夹第1页
抓取收藏夹第2页
抓取收藏夹第3页
...
```

结果：

![](data/1.png)
![](data/2.png)

## 二.API说明

下载

```bash
go get -u -v github.com/hunterhug/zhihuxx
```

此包在哥哥封装的爬虫包基础上开发：[土拨鼠（tubo）](https://github.com/hunterhug/GoSpider)，请进入`main`文件夹运行成品程序，`IDE`开发模式下，运行路径是不一样的，请在`IDE`项目根目录放`cookie.txt`文件

二次开发时你只需`import`本包。

```go
import zhihu "github.com/hunterhug/zhihuxx"
```

API如下：

```go
// 设置cookie，需传入文件位置，文件中放cookie
func SetCookie(file string) error 

// 构造问题链接，返回url
func Question(id string) string

// 抓答案，需传入页数，返回一堆数据
func CatchAnswer(url string, page int) ([]byte, error)

// 结构化回答，返回一个结构体
func StructAnswer(body []byte) (*Answer, error)

// 抓取收藏夹第几页列表
func CatchCoolection(id, page int) ([]byte, error)

// 抓取全部收藏夹页数,并返回问题ID和标题
func CatchAllCollection(id int) map[string]string 

// 解析收藏夹，返回问题ID和标题
func ParseCollection(body []byte) map[string]string

// 输出HTML选择防盗链方式
func SetPublishToWeb(put bool)

// 输出友好格式HTML，返回问题ID,回答ID，标题，作者，还有HTML
func OutputHtml(answer DataInfo) (qid, aid int, title, who, html string)

// 抓取图片前需要设置true
func SetSavePicture(catch bool) 

// 抓取html中的图片，保存图片在dir下
func SavePicture(dir string, body []byte) 

// 遇到返回的JSON中有中文乱码，请转意
func JsonBack(body []byte) ([]byte, error)

// 设置爬虫调试日志级别，开发可用:debug,info
func SetLogLevel(level string) 

// 设置爬虫暂停时间
func SetWaitTime(w int)
```

还差某些API，需逐步优化。

使用时需要先`SetCookie()`，再根据具体进行开发，使用如下：

```go
package main

import (
	"fmt"
	zhihu "github.com/hunterhug/zhihuxx"
	"strings"
)

// API使用说明
func main() {
	//  1. 设置爬虫暂停时间，可选
	zhihu.SetWaitTime(1)

	// 2. 调试模式设置为debug，可选
	zhihu.SetLogLevel("info")

	// 3. 需先传入cookie，必须
	e := zhihu.SetCookie("./cookie.txt")
	if e != nil {
		panic(e.Error())
	}

	// 4.构建问题，url差页数
	q := zhihu.Question("28467579")
	fmt.Println(q)

	// 5.抓取问题回答，按页数，传入页数是为了补齐url，策略是循环抓，直到抓不到可认为页数已完
	page := 1
	body, e := zhihu.CatchAnswer(q, page)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	if strings.Contains(string(body), "error") { //可能cookie失效
		b, _ := zhihu.JsonBack(body)
		fmt.Println(string(b))
	}

	// 6.结构化回答
	answers, e := zhihu.StructAnswer(body)
	if e != nil {
		fmt.Println(e.Error())
	} else {
		// 就不打出来了
		//fmt.Printf("%#v\n", answers.Page)
		//fmt.Printf("%#v\n", answers.Data)
	}

	// 7. 选择OutputHtml不要防盗链，因为回答输出的html经过了处理，所以我们进行过滤出好东西
	zhihu.SetPublishToWeb(false)
	qid,aid,t,who,html:=zhihu.OutputHtml(answers.Data[0])
	fmt.Println(qid)
	fmt.Println(aid)
	fmt.Println(t)
	fmt.Println(who)

	// 8. 抓图片
	zhihu.SetSavePicture(false)
	zhihu.SavePicture("test", []byte(html))

	// 9. 抓集合，第2页
	b, e := zhihu.CatchCoolection(78172986, 2)
	if e != nil {
		fmt.Println(e.Error())
	} else {
		// 解析集合
		fmt.Printf("%#v",zhihu.ParseCollection(b))
	}
}
```

登录待破解验证码：

```go
// 登录，验证码突破不了，请采用SetCookie
func Login(email, password string) ([]byte, error)
```

![](data/ca.png)

```
_xsrf:2fc4811def8cd9f358465e4ea418b23b
password:z13112502886
captcha:{"img_size":[200,44],"input_points":[[19.2969,28],[40.2969,28],[68.2969,27],[89.2969,31],[112.297,34],[138.297,15],[161.297,27]]}
captcha_type:cn
email:wefwefwefwef@qq.com
```

## 三.编译执行文件方式

### Linux下跨平台编译

Linux二进制

```bash
cd main
go build -ldflags "-s -w" -v -o zhihu_linux_x86_64 main.go
```

Windows二进制

```bash
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -x -o zhihu_windows_amd64.exe main.go 
```

### Windows编译

```bash
go build -o zhihu.exe main.go
```

如果你觉得项目帮助到你,欢迎请我喝杯咖啡

微信
![微信](https://raw.githubusercontent.com/hunterhug/hunterhug.github.io/master/static/jpg/wei.png)

支付宝
![支付宝](https://raw.githubusercontent.com/hunterhug/hunterhug.github.io/master/static/jpg/ali.png)

## 四.环境配置

### Ubuntu安装

[云盘](https://yun.baidu.com/s/1jHKUGZG)下载源码解压.下载IDE也是解压设置环境变量.

```bash
vim /etc/profile.d/myenv.sh

export GOROOT=/app/go
export GOPATH=/home/jinhan/code
export GOBIN=$GOROOT/bin
export PATH=.:$PATH:/app/go/bin:$GOPATH/bin:/home/jinhan/software/Gogland-171.3780.106/bin

source /etc/profile.d/myenv.sh
```

### Windows安装

[云盘](https://yun.baidu.com/s/1jHKUGZG) 选择后缀为msi安装如1.6

环境变量设置：

```bash
Path G:\smartdogo\bin
GOBIN G:\smartdogo\bin
GOPATH G:\smartdogo
GOROOT C:\Go\
```

### docker安装

我们的库可能要使用各种各样的工具，配置连我这种专业人员有时都搞不定，而且还可能会损坏，所以用docker方式随时随地开发。

先拉镜像

```bash
docker pull golang:1.8
```

Golang环境启动：

```bash
docker run --rm --net=host -it -v /home/jinhan/code:/go --name mygolang golang:1.8 /bin/bash

root@27214c6216f5:/go# go env
GOARCH="amd64"
```

其中`/home/jinhan/code`为你自己的本地文件夹（虚拟GOPATH），你在docker内`go get`产生在`/go`的文件会保留在这里，容器死掉，你的`/home/jinhan/code`还在，你可以随时修改文件配置。

启动后你就可以在里面开发了。


# LICENSE

欢迎加功能(PR/issues),请遵循Apache License协议(即可随意使用但每个文件下都需加此申明）

```
Copyright 2017 by GoSpider author.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License
```