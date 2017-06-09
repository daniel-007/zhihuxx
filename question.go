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
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/hunterhug/GoSpider/query"
	"github.com/hunterhug/GoSpider/util"
	"strings"
	"time"
)

var (
	// 问题链接
	Qurl      = "https://www.zhihu.com/api/v4/questions/%s/answers?"
	Qurlquery = "sort_by=default&include=data[*].%s&limit=1&offset=" //不要太暴力，每次一个答案就可以
	// 各种参数，问题获取到的JSON字段意思
	Qurlparm = []string{
		"is_normal",           // 是否正常
		"is_collapsed",        // 是否折叠
		"collapse_reason",     // 折叠理由
		"is_sticky",           // 无
		"collapsed_by",        // 5
		"suggest_edit",        // 建议
		"comment_count",       //评论数
		"can_comment",         // 能否评论
		"content",             //内容 重要
		"editable_content",    //5
		"voteup_count",        // 投票数?
		"reshipment_settings", //?
		"comment_permission",  // 可否评论
		"mark_infos",          //5
		"created_time",
		"updated_time",
		"relationship.is_authorized,is_author,voting,is_thanked,is_nothelp", // 关系？
		"upvoted_followees;data[*].author.follower_count,badge[?(type=best_answerer)].topics",
	}
)

type Answer struct {
	Page PageInfo   `json:"paging"`
	Data []DataInfo `json:"data"`
}

type PageInfo struct {
	IsEnd   bool   `json:"is_end"`
	Totals  int    `json:"totals"`
	PreUrl  string `json:"previous"`
	IsStart bool   `json:"is_start"`
	NextUrl string `json:"next"`
}

type DataInfo struct {
	Excerpt    string       `json:"excerpt"`
	Author     AuthorInfo   `json:"author"`
	Question   QuestionInfo `json:"question"`
	Content    string       `json:"content"`
	Aid        int          `json:"id"`
	CreateTime int          `json:"created_time"`
	UpdateTime int          `json:"updated_time"`
}

type AuthorInfo struct {
	About    string `json:"headline"`
	UrlToken string `json:"url_token"`
	Name     string `json:"name"`
	Sex      int    `json:"gender"`
	Image    string `json:"avatar_url_template"`
}
type QuestionInfo struct {
	CreateTime int    `json:"created`
	Title      string `json:"title"`
	UpdateTime int    `json:"updated_time"`
	Qid        int    `json:"id"`
}

// 构造问题链接
func Q(id string) string {
	return fmt.Sprintf(Qurl, id) + fmt.Sprintf(Qurlquery, strings.Join(Qurlparm, ","))
}

// 抓答案，需传入页数
func CatchA(url string, page int) ([]byte, error) {
	if page < 1 {
		page = 1
	}
	Baba.SetUrl(url + util.IS((page-1)*1))
	body, err := Baba.Get()
	if err != nil {

	} else {
		//body, err = util.JsonBack(body)
	}
	return body, err
}

// 结构化返回数据
func StructA(body []byte) (*Answer, error) {
	temp := new(Answer)
	err := json.Unmarshal(body, temp)
	return temp, err
}

// 输出HTML
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

func SavePicture(dir string, body []byte) {
	if !CatchP {
		return
	}
	util.MakeDir(dir)
	docm, err := query.QueryBytes(body)
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		//fmt.Println(string(content))
		docm.Find("img").Each(func(num int, node *goquery.Selection) {
			img, e := node.Attr("src")
			if e == false {
				img, e = node.Attr("data-src")
			}
			if e && img != "" {
				//fmt.Println("原始文件：" + img)
				temp := img
				filename := util.ValidFileName(temp)
				if util.FileExist(dir + "/" + filename) {
					//fmt.Println("文件存在：" + dir + "/" + filename)
				} else {
					//fmt.Println("下载:" + temp)
					Baba.SetUrl(temp)
					imgsrc, e := Baba.Get()
					if e != nil {
						fmt.Println("下载出错" + temp + ":" + e.Error())
						return
					}
					e = util.SaveToFile(dir+"/"+filename, imgsrc)
					if e == nil {
						fmt.Println("成功保存在" + dir + "/" + filename)
					}
					util.Sleep(1)
					//fmt.Println("暂停1秒")
				}
			}
		})
	}
}