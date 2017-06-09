# 一.说明

Golang开发的爬虫，入口在main文件夹下，你需要放一个cookie.txt在里面

打开火狐浏览器后人工登录知乎，按F12 ，点击网络，刷新一下首页，然后点击第一个出现的GET /，找到消息头请求头，复制Cookie，然后粘贴到cookie.txt

![](data/cookie.png)

```
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
```

# 二.使用

下载

```
go get -u -v github.com/hunterhug/zhihuxx
```

运行

```
cd main
go run main.go
```

已经为非程序员编译好exe文件，点击即可！

# 三.编译成执行文件

## Linux下跨平台编译

Linux二进制
```
cd main
go build -o zhihu_linux_x86_64 main.go 
```

Windows二进制
```
GOOS=windows GOARCH=amd64 go build -x -o zhihu_windows_amd64.exe main.go 
```

## Windows编译

```
go build -o zhihu.exe main.go
```

# 四.环境配置

## Ubuntu安装

下载源码解压.下载IDE也是解压设置环境变量.

```
vim /etc/profile.d/myenv.sh

export GOROOT=/app/go
export GOPATH=/home/jinhan/code
export GOBIN=$GOROOT/bin
export PATH=.:$PATH:/app/go/bin:$GOPATH/bin:/home/jinhan/software/Gogland-171.3780.106/bin

source /etc/profile.d/myenv.sh
```

## Windows安装

[](https://yun.baidu.com/s/1jHKUGZG) 选择后缀为msi安装如1.6

环境变量设置：

```
Path G:\smartdogo\bin
GOBIN G:\smartdogo\bin
GOPATH G:\smartdogo
GOROOT C:\Go\
```

## docker安装

我们的库可能要使用各种各样的工具，配置连我这种专业人员有时都搞不定，而且还可能会损坏，所以用docker方式随时随地开发。

先拉镜像

```
docker pull golang:1.8
```

Golang环境启动：

```
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