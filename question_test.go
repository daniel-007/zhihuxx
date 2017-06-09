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
	"path/filepath"
	"strings"
	"testing"
)

func TestCatchA(t *testing.T) {
	cookie := `q_c1=ab5eab2e40524b4e90aae70065efb5ba|1496835483000|1496835483000; r_cap_id="OTIxNzY0OTkzZjgyNDM0ZGJkZGNiZGYwN2I1MDUwN2M=|1496835486|0f4561035485d66a94a39ad505b3e944eed2ad16"; cap_id="ZDc3NmU5NTUxOWY3NDZlZmI3YjAwNGY0MjNhOGU5ODM=|1496835486|72e747a6ae31b38aae7c74f36d1f6829168bf5b6"; d_c0="AEBCmeip4AuPTuPzZwZLKjXPTWAq5lT-sXs=|1496835487"; _zap=b9ac213e-63aa-4020-8250-0e050afb3f56; _xsrf=2fc4811def8cd9f358465e4ea418b23b; __utma=51854390.576976127.1496835487.1496835487.1496973515.2; __utmb=51854390.0.10.1496973515; __utmc=51854390; __utmz=51854390.1496973515.2.2.utmcsr=zhihu.com|utmccn=(referral)|utmcmd=referral|utmcct=/; __utmv=51854390.100-1|2=registration_date=20150209=1^3=entry_date=20150209=1; z_c0=Mi4wQUJEQk9vSjVtd2NBUUVLWjZLbmdDeGNBQUFCaEFsVk51SEpmV1FBZmpxSS1KMHMxV3BOUWZuRWtiMHpKT2dnNURR|1496973514|4e40191547ed419cebdda17a7d14d2c74dbc987e`
	Baba.SetHeaderParm("Cookie", strings.TrimSpace(cookie))
	b, e := CatchA(Q("28467579"), 1)
	if e != nil {
		fmt.Println(e.Error())
	} else {
		util.SaveToFile(filepath.Join(util.CurDir(), "data/question.json"), []byte(b))
	}
}

func TestStructQ(t *testing.T) {
	body, err := util.ReadfromFile(filepath.Join(util.CurDir(), "data/question.json"))
	if err != nil {
		panic(err.Error())
	}
	temp, err := StructA(body)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("总数:%#v\n", temp.Page.Totals)
		fmt.Printf("%#v\n", temp.Page.IsEnd)
		fmt.Printf("%#v\n", temp.Page.IsStart)
		fmt.Printf("下一个:%#v\n", temp.Page.NextUrl)
		fmt.Printf("回答：%#v\n", temp.Data[0].Content)
	}

}

func TestOutputHtml(t *testing.T) {
	body, err := util.ReadfromFile(filepath.Join(util.CurDir(), "data/question.json"))
	if err != nil {
		panic(err.Error())
	}
	temp, err := StructA(body)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		answer := temp.Data[0]
		qid, title, aid, who, html := OutputHtml(answer)
		fmt.Println(qid, aid, who, title)
		util.SaveToFile(filepath.Join(util.CurDir(), "data/question.html"), []byte(html))
	}

}

func TestSavePicture(t *testing.T) {
	body, err := util.ReadfromFile(filepath.Join(util.CurDir(), "data/question.html"))
	dir := filepath.Join(util.CurDir(), "data/00/00")
	if err != nil {
		panic(err.Error())
	}
	SavePicture(dir, body)
}
