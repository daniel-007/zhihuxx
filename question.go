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
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/hunterhug/GoSpider/query"
	"github.com/hunterhug/GoSpider/util"
	"strings"
)

var (
	// 问题链接
	Qurl      = "https://www.zhihu.com/api/v4/questions/%s/answers?"
	Qurlquery = "sort_by=default&include=data[*].%s"
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

var tempParm = fmt.Sprintf(Qurlquery, strings.Join(Qurlparm, ","))

// 构造问题链接，返回url,你可以通过Qurlparm拼出另外一个url
func Question(id string) string {
	return fmt.Sprintf(Qurl, id) + tempParm + "&limit=%d&offset=%d"
}

// 抓答案，需传入限制和页数,每次最多抓20个答案
func CatchAnswer(url string, limit, page int) ([]byte, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 1
	}
	if limit > 20 {
		limit = 20
	}
	uurl := fmt.Sprintf(url, limit, (page-1)*limit)

	//fmt.Println(uurl)
	Baba.SetUrl(uurl)

	body, err := Baba.Get()
	if err != nil {

	} else {
		if strings.Contains(string(body), "AuthenticationInvalid") {
			data, _ := JsonBack(body)
			return data, errors.New("AuthenticationInvalid cookie fail")
		}
	}
	return body, err
}

// 结构化回答，返回一个结构体
func StructAnswer(body []byte) (*Answer, error) {
	temp := new(Answer)
	err := json.Unmarshal(body, temp)
	return temp, err
}

// 抓取图片前需要设置true
func SetSavePicture(catch bool) {
	CatchP = catch
}

// 抓取html中的图片，保存图片在dir下
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
					//util.Sleep(1)
					//fmt.Println("暂停1秒")
				}
			}
		})
	}
}
