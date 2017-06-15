/*
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
*/
package zhihuxx

import (
	"fmt"
	"github.com/hunterhug/GoSpider/util"
	"strings"
	"time"
)

func Input(say, defaults string) string {
	fmt.Println(say)
	var str string
	fmt.Scanln(&str)
	if strings.TrimSpace(str) == "" {
		if strings.TrimSpace(defaults) != "" {
			return defaults
		} else {
			fmt.Println("不能为空！")
			return Input(say, defaults)
		}
	}
	//fmt.Println("--" + str + "--")
	return str
}

// 输出友好格式HTML，返回问题ID,回答ID，标题，作者，还有HTML
func OutputHtml(answer DataInfo) (qid, aid int, title, who, html string) {
	answer.Content = strings.Replace(answer.Content, "src", "xx", -1)
	if PublishToWeb {
		answer.Content = strings.Replace(answer.Content, "data-original", "data-src", -1)
	} else {
		answer.Content = strings.Replace(answer.Content, "data-original", "src", -1)
	}
	b := `
		<!DOCTYPE html>
		<html>
		<head>
		<meta charset="utf-8" />
		<title>%s:%d</title>
		<style>
		body{
		margin:20px 15%%
		}
		img {width:60%%;
		display:block;
		text-align:center}

		.link{
		margin:20px;
		height:10px
		}
		</style>
		%s
		</head>
		<body>
		<div id="author">
		%s
		</div>
		<div class="link">
		###link###
		</div>
				<div>
  跳页: <input type="number" id="page" min="1" max="500" value="3" style="width:100px">
  <input type="submit" onclick="var a=document.getElementById('page').value;location.href=a+'.html' "></div>
		<div id="answer">
		<hr/>
		正文:
		%s
		</div>
		<div class="link">
		###link###
		</div>
				<div>
  跳页: <input type="number" id="page" min="1" max="500" value="3" style="width:100px">
  <input type="submit" onclick="var a=document.getElementById('page').value;location.href=a+'.html' "></div>
		</body>
		</html>
		`

	sex := "男"
	if answer.Author.Sex == 0 {
		sex = "女"
	}
	purl := fmt.Sprintf(PeopleUrl, answer.Author.UrlToken)
	qurl := fmt.Sprintf(QuestionUrl, answer.Question.Qid)
	aurl := fmt.Sprintf(AnswerUrl, answer.Question.Qid, answer.Aid)
	ct := time.Unix(int64(answer.CreateTime), 0).Format("2006-01-02 03:04:05 PM")
	ut := time.Unix(int64(answer.UpdateTime), 0).Format("2006-01-02 03:04:05 PM")
	about := fmt.Sprintf(`
		名字:<a href="%s">%s</a> 性别:%s<br/>
		<img data-src="%s" width="400" height="500" /><br/>
		介绍:%s<br/>
		<a href="%s">问题</a><br/>
		<a href="%s">答案</a>新建于:%s，更新于%s

		<br/>
		`, purl, answer.Author.Name, sex, strings.Replace(answer.Author.Image, "{size}", "xll", -1), answer.Author.About, qurl, aurl, ct, ut)

	if !PublishToWeb {
		about = strings.Replace(about, "data-src", "src", -1)
	}
	JsScript := ""
	if PublishToWeb {
		JsScript = "<script type='application/ecmascript' async='' src='../" + JsName + "'></script>"
	}
	content := fmt.Sprintf(b, answer.Question.Title, answer.Aid, JsScript, about, answer.Content)
	return answer.Question.Qid, answer.Aid, answer.Question.Title, answer.Author.UrlToken, content
}

// 遇到返回的JSON中有中文乱码，请转意
func JsonBack(body []byte) ([]byte, error) {
	return util.JsonBack(body)
}
