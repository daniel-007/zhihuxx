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
	"strings"
)

var (
	QuestionUrl   = "https://www.zhihu.com/question/%d"
	PeopleUrl     = "https://www.zhihu.com/people/%s"
	AnswerUrl     = "https://www.zhihu.com/question/%d/answer/%d"
	CollectionUrl = "https://www.zhihu.com/collection/%d?page=%d"
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
