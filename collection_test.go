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

func TestCatchCoolection(t *testing.T) {
	cookie := `q_c1=902510c4493740aca0c12964714d21a9|1496466243000|1496466243000; cap_id="MzNmOTgzYWFiNGFiNGZhNmFmNzJmNjA5NWM1ZDQwZGQ=|1496564819|afe4a7dcd180e9426c727c130354e202896def55"; d_c0="AJACQuYp2wuPTgukFS3cyRKIs-9xlHIj7yo=|1496466385"; _zap=21c848aa-1c3f-4ce4-ac2d-6368660733ef; __utma=51854390.1031072116.1496559262.1496559262.1496564189.2; __utmz=51854390.1496564189.2.2.utmcsr=zhihu.com|utmccn=(referral)|utmcmd=referral|utmcct=/oauth/callback/sina; aliyungf_tc=AQAAAC3Fq2HHJQsAodFbcR0GDyNZQbBe; acw_tc=AQAAABMu8x8RVAsAodFbcbgVKDVnqlW9; AffiliatedA=1; _xsrf=bbb860f76247a68ca0a56e76f3e783fa; __utmc=51854390; __utmv=51854390.100--|2=registration_date=20170604=1^3=entry_date=20170603=1; capsion_ticket="2|1:0|10:1496564171|14:capsion_ticket|44:MGZjNDJhMTlmZDFlNDkxZWFiY2FjOTIzZDVmMGFmNTU=|36587999e00c35d87dd75753c7531bffe9698f39132f67af451b5ac65325c4ec"; auth_type="c2luYQ==|1496564197|07ffd2d0382b2cdc2c12735388eb69fd0e7fe050"; atoken=2.00xbIHGCEA722D3f606a091cU2ZEXB; atoken_expired_in=2630603; token="Mi4wMHhiSUhHQ0VBNzIyRDNmNjA2YTA5MWNVMlpFWEI=|1496564197|aa3d82667383fb85b8ce8c4565f8c27b292e8f10"; client_id="MTkyMjYyNTA4MQ==|1496564197|e0641a36db7ec0413ffd811fcdbd64742fa70231"; __utmb=51854390.0.10.1496564189; r_cap_id="YjM0MDU3NGIxZmUxNDUyMGE2ZDIyMTFkMWM2MGI1ZmM=|1496564734|69bcb9a0d26829b78bfeaf2a6a94ea6525b4617c"; z_c0=Mi4wQUNCQ054YWYzQXNBa0FKQzVpbmJDeGNBQUFCaEFsVk5ZMUZiV1FDV01YSVp3TUh6d2lET3pGQmdNay1tM3RKRC1R|1496564841|ae99b5a75fdb55b818f849d8e7533c1b58d6e6c2`
	Baba.SetHeaderParm("Cookie", strings.TrimSpace(cookie))
	b, e := CatchCoolection(78172986, 2)
	if e != nil {
		fmt.Println(e.Error())
	} else {
		util.SaveToFile(filepath.Join(util.CurDir(), "data/collection.html"), []byte(b))
	}
}

func TestParseCollection(t *testing.T) {
	body, err := util.ReadfromFile(filepath.Join(util.CurDir(), "data/collection.html"))
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%#v.", ParseCollection(body))

}

func TestCatchAllCollection(t *testing.T) {
	fmt.Printf("%#v,", CatchAllCollection(78172986))
}
