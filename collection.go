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
	"github.com/PuerkitoBio/goquery"
	"github.com/hunterhug/GoSpider/query"
	"strings"
)

// 抓取收藏夹第几页列表
func CatchCoolection(id, page int) ([]byte, error) {
	Baba.SetUrl(fmt.Sprintf(CollectionUrl, id, page))
	return Baba.Get()
}


// 抓取收藏夹全部问题ID
func CatchAllCollection(id int) []string {
	returns := []string{}
	i := 1
	for {
		body, err := CatchCoolection(id, i)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Printf("抓取收藏夹第%d页\n", i)
		i = i + 1
		lists := ParseCollection(body)
		if len(lists) == 0 {
			break
		}
		returns = append(returns, lists...)
	}
	return returns
}

func ParseCollection(body []byte) []string {
	returns := []string{}
	doc, _ := query.QueryBytes(body)
	//zm-item-title
	doc.Find(".zm-item-title").Each(func(num int, node *goquery.Selection) {
		q, ok := node.Find("a").Attr("href")
		if ok {
			returns = append(returns, strings.Replace(q, "/question/", "", -1))
		}
	})
	return returns
}
