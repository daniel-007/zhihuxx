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
package main

import (
	"fmt"
	"github.com/hunterhug/GoSpider/util"
	zhihu "github.com/hunterhug/zhihuxx"
	"os"
	"strings"
	"time"
)

var Limit = 30 //限制回答个数
var Follow = false

// 抓取一个问题的全部信息

func help() {
	fmt.Println(`
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
	`)
}

// 应该替换为本地照片！待做
func main() {
	help()

	haha, err := util.ReadfromFile("cookie.txt")
	if err != nil {
		fmt.Println("请您一定要保证cookie.txt存在哦：" + err.Error())
		time.Sleep(50 * time.Second)
		os.Exit(0)
	}
	cookie := string(haha)
	zhihu.Baba.SetHeaderParm("Cookie", strings.TrimSpace(cookie))

	js := strings.ToLower(zhihu.Input("萌萌：你要发布到自己的网站上吗(JS解决防盗链)Y/N(默认N)", "n"))
	if strings.Contains(js, "y") {
		zhihu.PublishToWeb = true
		zhihu.InitJs()
		util.SaveToFile("data/"+zhihu.JsName, []byte(zhihu.Js))
	}
	tu := strings.ToLower(zhihu.Input("萌萌：要抓取图片吗Y/N(默认N)", "n"))
	if strings.Contains(tu, "y") {
		zhihu.CatchP = true
	}
	choice := zhihu.Input("萌萌：从收藏夹获取按1，从问题获取按2(默认)", "2")
	for {
		ll := zhihu.Input("萌萌说亲爱的，因为回答实在太多，请限制获取的回答个数:30（默认)", "30")
		temp, errr := util.SI(ll)
		if errr != nil {
			fmt.Println("萌萌表示输入的应该是数字哦")
			continue
		}
		if temp <= 0 || temp > 500 {
			fmt.Println("萌萌表示不抓取，哼")
			continue
		}
		Limit = temp
		break
	}
	//ff := util.ToLower(zhihu.Input("酱~关注下答案中的小姐姐/小哥哥吧，默认N（Y/N）", "n"))
	ff := "n"
	if strings.Contains(ff, "y") {
		Follow = true
	}
	if choice == "1" {
		Many()
	} else {
		Base()
	}

}

func Base() {
	for {
		page := 1
		//28467579
		id := zhihu.Input("萌萌：请输入问题ID:", "")
		q := zhihu.Q(id)
		//fmt.Println(q)

		// 第一个答案
		body, err := zhihu.CatchA(q, page)
		fmt.Println("预抓取第一个回答！")
		if err != nil {
			fmt.Println("a" + err.Error())
			continue
		}

		temp, err := zhihu.StructA(body)
		if err != nil {
			fmt.Println("b" + err.Error())
			s, _ := util.JsonBack(body)
			fmt.Println(string(s))
			continue
		}
		if len(temp.Data) == 0 {
			fmt.Println("没有答案！")
			continue
		}

		fmt.Println("开始处理答案:" + temp.Data[0].Excerpt)
		qid, aid, title, who, html := zhihu.OutputHtml(temp.Data[0])
		fmt.Println("哦，这个问题是:" + title)
		filename := fmt.Sprintf("data/%d/%s-%d/%s-%d的回答.html", qid, who, aid, who, aid)
		util.MakeDirByFile(filename)
		if zhihu.PublishToWeb {
			util.SaveToFile(fmt.Sprintf("data/%d/%s", qid, zhihu.JsName), []byte(zhihu.Js))
		}
		util.SaveToFile(fmt.Sprintf("data/%d-%s.xx", qid, util.ValidFileName(title)), []byte(""))
		err = util.SaveToFile(filename, []byte(html))

		// html
		util.MakeDir(fmt.Sprintf("data/%d-html", qid))
		link := ""
		if page == 1 {
			link = fmt.Sprintf(`<a href="%d.html" style="float:right">Next下一页</a>`, page+1)
		} else {
			link = fmt.Sprintf(`<a href="%d.html" style="float:left">Pre上一页</a><a href="%d.html" style="float:right">Next下一页</a>`, page-1, page+1)
		}
		html = strings.Replace(html, "###link###", link, -1)
		util.SaveToFile(fmt.Sprintf("data/%d-html/%d.html", qid, page), []byte(html))

		if Follow {
			zhihu.Follow(who)
		}
		if err == nil {
			fmt.Println("保存答案成功:" + filename)
		} else {
			fmt.Println("保存答案失败:" + err.Error())
			continue
		}
		zhihu.SavePicture(fmt.Sprintf("data/%d/%s-%d", qid, who, aid), []byte(html))

		all := util.ToLower(zhihu.Input("批量抓取答案，默认N(Y/N)", "N"))
		for {
			if temp.Page.IsEnd {
				fmt.Println("回答已经结束！")
				break
			}
			if strings.Contains(all, "n") {
				yes := util.ToLower(zhihu.Input("抓取下一个答案，默认Y(Y/N)", "Y"))
				if strings.Contains(yes, "n") {
					break
				}
			}
			//fmt.Println(temp.Page.NextUrl)
			if page+1 > Limit {
				fmt.Println("萌萌：答案超出个数了哦，哦耶~")
				break
			}
			body, err = zhihu.CatchA(q, page+1)
			if err != nil {
				fmt.Println("抓取答案失败：" + err.Error())
				continue
			} else {
				page = page + 1
			}
			//util.SaveToFile("data/question.json", body)

			temp1, err := zhihu.StructA(body)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			if len(temp1.Data) == 0 {
				fmt.Println("没有答案！")
				s, _ := util.JsonBack(body)
				fmt.Println(string(s))
				continue
			}

			// 成功后再来
			temp = temp1

			fmt.Println("开始处理答案:" + temp.Data[0].Excerpt)
			qid, aid, _, who, html := zhihu.OutputHtml(temp.Data[0])
			filename := fmt.Sprintf("data/%d/%s-%d/%s-%d的回答.html", qid, who, aid, who, aid)
			util.MakeDirByFile(filename)
			err = util.SaveToFile(filename, []byte(html))
			// html
			util.MakeDir(fmt.Sprintf("data/%d-html", qid))
			link := ""
			if page == 1 {
				link = fmt.Sprintf(`<a href="%d.html" style="float:right">Next下一页</a>`, page+1)
			} else {
				link = fmt.Sprintf(`<a href="%d.html" style="float:left">Pre上一页</a><a href="%d.html" style="float:right">Next下一页</a>`, page-1, page+1)
			}
			html = strings.Replace(html, "###link###", link, -1)
			util.SaveToFile(fmt.Sprintf("data/%d-html/%d.html", qid, page), []byte(html))

			if Follow {
				zhihu.Follow(who)
			}
			if err == nil {
				fmt.Println("保存答案成功:" + filename)
			} else {
				fmt.Println("保存答案失败:", err.Error())
				continue
			}
			zhihu.SavePicture(fmt.Sprintf("data/%d/%s-%d", qid, who, aid), []byte(html))
		}
	}
}

func Many() {
	for {
		//78172986
		collectids := zhihu.Input("萌萌：请输入集合ID:", "")
		collectid, e := util.SI(collectids)
		if e != nil {
			fmt.Println("收藏夹ID错误")
			continue
		}

		god := util.ToLower(zhihu.Input("开启上帝模式吗(一路抓到底)，默认N(Y/N)?", "N"))
		skip := false
		if strings.Contains(god, "y") {
			skip = true
		}
		qids := zhihu.CatchAllCollection(collectid)
		if len(qids) == 0 {
			fmt.Println("收藏夹下没问题！")
			continue
		}
		fmt.Printf("总计有%d个问题:\n", len(qids))
		for _, id := range qids {
			page := 1
			q := zhihu.Q(id)
			//fmt.Println(q)

			// 第一个答案
			body, err := zhihu.CatchA(q, page)
			fmt.Println("预抓取第一个回答！")
			if err != nil {
				fmt.Println("a" + err.Error())
				continue
			}

			temp, err := zhihu.StructA(body)
			if err != nil {
				fmt.Println("b" + err.Error())
				s, _ := util.JsonBack(body)
				fmt.Println(string(s))
				continue
			}
			if len(temp.Data) == 0 {
				fmt.Println("没有答案！")
				continue
			}

			fmt.Println("开始处理答案:" + temp.Data[0].Excerpt)
			qid, aid, title, who, html := zhihu.OutputHtml(temp.Data[0])
			fmt.Println("哦，这个问题是:" + title)
			if util.FileExist(fmt.Sprintf("data/%d-%s.xx", qid, util.ValidFileName(title))) {
				fmt.Printf("已经存在：%s,跳过！\n", fmt.Sprintf("data/%d-%s.xx", qid, util.ValidFileName(title)))
				continue
			}

			if !skip {
				tiaotiao := util.ToLower(zhihu.Input("跳过这个问题吗，默认N(Y/N)?", "N"))
				if strings.Contains(tiaotiao, "y") {
					continue
				}
			}
			filename := fmt.Sprintf("data/%d/%s-%d/%s-%d的回答.html", qid, who, aid, who, aid)
			util.MakeDirByFile(filename)
			if zhihu.PublishToWeb {
				util.SaveToFile(fmt.Sprintf("data/%d/%s", qid, zhihu.JsName), []byte(zhihu.Js))
			}
			util.SaveToFile(fmt.Sprintf("data/%d-%s.xx", qid, util.ValidFileName(title)), []byte(""))
			err = util.SaveToFile(filename, []byte(html))
			// html
			util.MakeDir(fmt.Sprintf("data/%d-html", qid))
			link := ""
			if page == 1 {
				link = fmt.Sprintf(`<a href="%d.html" style="float:right">Next下一页</a>`, page+1)
			} else {
				link = fmt.Sprintf(`<a href="%d.html" style="float:left">Pre上一页</a><a href="%d.html" style="float:right">Next下一页</a>`, page-1, page+1)
			}
			html = strings.Replace(html, "###link###", link, -1)
			util.SaveToFile(fmt.Sprintf("data/%d-html/%d.html", qid, page), []byte(html))

			if Follow {
				zhihu.Follow(who)
			}
			if err == nil {
				fmt.Println("保存答案成功:" + filename)
			} else {
				fmt.Println("保存答案失败:" + err.Error())
				continue
			}
			zhihu.SavePicture(fmt.Sprintf("data/%d/%s-%d", qid, who, aid), []byte(html))

			all := "y"
			if !skip {
				all = util.ToLower(zhihu.Input("批量抓取这个问题的所有答案，默认N(Y/N)", "N"))
			}
			for {
				if temp.Page.IsEnd {
					fmt.Println("回答已经结束！")
					break
				}
				if strings.Contains(all, "n") {
					yes := util.ToLower(zhihu.Input("抓取下一个答案，默认Y(Y/N)", "Y"))
					if strings.Contains(yes, "n") {
						break
					}
				}
				//fmt.Println(temp.Page.NextUrl)
				if page+1 > Limit {
					fmt.Println("萌萌：答案超出个数了哦，哦耶~")
					break
				}
				body, err = zhihu.CatchA(q, page+1)
				if err != nil {
					fmt.Println("抓取答案失败：" + err.Error())
					continue
				} else {
					page = page + 1
				}
				//util.SaveToFile("data/question.json", body)

				temp1, err := zhihu.StructA(body)
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
				if len(temp1.Data) == 0 {
					fmt.Println("没有答案！")
					s, _ := util.JsonBack(body)
					fmt.Println(string(s))
					continue
				}

				// 成功后再来
				temp = temp1

				fmt.Println("开始处理答案:" + temp.Data[0].Excerpt)
				qid, aid, _, who, html := zhihu.OutputHtml(temp.Data[0])
				filename := fmt.Sprintf("data/%d/%s-%d/%s-%d的回答.html", qid, who, aid, who, aid)
				util.MakeDirByFile(filename)
				err = util.SaveToFile(filename, []byte(html))
				// html
				util.MakeDir(fmt.Sprintf("data/%d-html", qid))
				link := ""
				if page == 1 {
					link = fmt.Sprintf(`<a href="%d.html" style="float:right">Next下一页</a>`, page+1)
				} else {
					link = fmt.Sprintf(`<a href="%d.html" style="float:left">Pre上一页</a><a href="%d.html" style="float:right">Next下一页</a>`, page-1, page+1)
				}
				html = strings.Replace(html, "###link###", link, -1)
				util.SaveToFile(fmt.Sprintf("data/%d-html/%d.html", qid, page), []byte(html))

				if Follow {
					zhihu.Follow(who)
				}
				if err == nil {
					fmt.Println("保存答案成功:" + filename)
				} else {
					fmt.Println("保存答案失败:", err.Error())
					continue
				}
				zhihu.SavePicture(fmt.Sprintf("data/%d/%s-%d", qid, who, aid), []byte(html))
			}
		}
	}
}
