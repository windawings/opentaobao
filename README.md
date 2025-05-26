# 淘宝开放平台 golang SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/windawings/opentaobao.svg)](https://pkg.go.dev/github.com/windawings/opentaobao)
[![Go](https://github.com/windawings/opentaobao/actions/workflows/go.yml/badge.svg)](https://github.com/windawings/opentaobao/actions/workflows/go.yml)
[![goreleaser](https://github.com/windawings/opentaobao/actions/workflows/goreleaser.yml/badge.svg)](https://github.com/windawings/opentaobao/actions/workflows/goreleaser.yml)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/bububa/opentaobao.svg)](https://github.com/windawings/opentaobao)
[![GoReportCard](https://goreportcard.com/badge/github.com/windawings/opentaobao)](https://goreportcard.com/report/github.com/windawings/opentaobao)
[![GitHub license](https://img.shields.io/github/license/bububa/opentaobao.svg)](https://github.com/windawings/opentaobao/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/bububa/opentaobao.svg)](https://github.com/windawings/opentaobao/releases/)

## 编译工具链
```sh
make tools
```
- 在bin目录下会生成downloader和generator两个可执行文件 
- 因为使用embed特性，编译downloder和generator需要golang版本>= 1.16


## 下载淘宝开放平台ApiMetadata
```sh
./bin/downloader -meta=./metadata/assets/json
```

### downloader 参数
- meta: 下载metadata目录
- pkg: 指定下载特定包，如：-pkg=user, 仅下载"用户API"

## 根据metadata生成API Golang SDK
```sh
./bin/generator -meta=./metadata/assets/json -patch=./metadata/assets/patch
```

### generator 参数
- meta: metadata文件所在目录
- patch: metadata patch文件所在目录
- pkg: 指定生成特定API包 


### API 调用
```golang
package main

import (
    "context"
    "log"

    "github.com/windawings/opentaobao/core"
    userModel "github.com/windawings/opentaobao/model/user"
    userApi "github.com/windawings/opentaobao/api/user"
)

func main() {
    clt := core.NewSDKClient(APP_KEY, APP_SECRET)
    req := userModel.GetTaobaoUserAvatarGetAPIRequest()
    defer userModel.PutTaobaoUserAvatarGetAPIRequest(req)
    req.SetNick("nick")
    resp := userModel.GetTaobaoUserAvatarGetAPIResponse()
    defer userModel.PutTaobaoUserAvatarGetAPIResponse(resp)
    accessToken := ""
    if err := userApi.TaobaoUserAvatarGet(context.Background(), clt, req, resp, accessToken); err != nil {
        log.Fatalln(err)
    }
    log.Printf("%+v\n", resp)
}
```

## 常见问题

### API分包
为避免生成最终目录内文件过多，以及命名空间冲突问题，根据淘宝开放平台API分类原则对API分包。
分包设置在metadata/assets/package.json，"pkg"参数即为API包名，与"id"(API树节点ID)对应。
下载的metadata如果在package.json无法找到对应的包则不会下载保存。

### 字段类型映射
| 淘宝字段类型 | go SDK 类型 |
| ----------- | -----------|
| Number | int64 |
| Price | float64 |
| BigDecimal | float64 |
| Boolean | bool |
| String | string |
| Date | string |
| Json | string |
| Field List | []string |
| 其他 | struct指针 |

- 如果字段类型后带"[]", 则转换为slice
- 如果字段UsePointer=true, 则转换为指针
- 同一保重struct名称一样会对struct成员进行合并

淘宝API中对象命名空间非常混乱，即使在同一包中仍然会出现struct冲突的情况。如果出现struct命名空间冲突有两种解决办法:
- 修改metadata/assets/conflict_models.json文件，增加可能冲突的struct名。generator会对该struct前增加API名进行重命名，比如 
AlibabaWdkFinanceOrderBackflow API中的ApiResult会被重命名为AlibabaWdkFinanceOrderBackflowApiResult
- 也可增加对应的patch文件修改字段的Type。generator在生成API时会先查找对应的API有没有patch，如果有patch则使用patch文件替换metadata文件

### 缺少特定API
对于无法通过downloader下载的API metaddata, 可以参照已下载metadata文件在patch目录自行添加相应的metadata, generator会生成相应的API。

## API 列表
| 淘宝API分类 | 对应SDK package |
| ---------- | ---------- |
| [x] [用户API](https://open.taobao.com/API.htm?docType=2&docId=24938) | [github.com/windawings/opentaobao/api/user](https://pkg.go.dev/github.com/windawings/opentaobao/api/user) |
| [x] [类目API](https://open.taobao.com/API.htm?docType=2&docId=47659) | [github.com/windawings/opentaobao/api/category](https://pkg.go.dev/github.com/windawings/opentaobao/api/category) |
| [x] [商品API](https://open.taobao.com/API.htm?docType=2&docId=38593) | [github.com/windawings/opentaobao/api/product](https://pkg.go.dev/github.com/windawings/opentaobao/api/product) |
| [x] [交易API](https://open.taobao.com/API.htm?docType=2&docId=67523) | [github.com/windawings/opentaobao/api/trade](https://pkg.go.dev/github.com/windawings/opentaobao/api/trade) |
| [x] [评价API](https://open.taobao.com/API.htm?docType=2&docId=48881) | [github.com/windawings/opentaobao/api/traderate](https://pkg.go.dev/github.com/windawings/opentaobao/api/traderate) |
| [x] [物流API](https://open.taobao.com/API.htm?docType=2&docId=28329) | [github.com/windawings/opentaobao/api/logistic](https://pkg.go.dev/github.com/windawings/opentaobao/api/logistic) |
| [x] [店铺API](https://open.taobao.com/API.htm?docType=2&docId=30898) | [github.com/windawings/opentaobao/api/shop](https://pkg.go.dev/github.com/windawings/opentaobao/api/shop) |
| [x] [分销API](https://open.taobao.com/API.htm?docType=2&docId=62017) | [github.com/windawings/opentaobao/api/fenxiao](https://pkg.go.dev/github.com/windawings/opentaobao/api/fenxiao) |
| [x] [旺旺API](https://open.taobao.com/API.htm?docType=2&docId=66033) | [github.com/windawings/opentaobao/api/wangwang](https://pkg.go.dev/github.com/windawings/opentaobao/api/wangwang) |
| [x] [淘宝客API](https://open.taobao.com/API.htm?docType=2&docId=48340) | [github.com/windawings/opentaobao/api/tbk](https://pkg.go.dev/github.com/windawings/opentaobao/api/tbk) |
| [x] [工具API](https://open.taobao.com/API.htm?docType=2&docId=47640) | [github.com/windawings/opentaobao/api/util](https://pkg.go.dev/github.com/windawings/opentaobao/api/util) |
| [x] [物流宝API](https://open.taobao.com/API.htm?docType=2&docId=28394) | [github.com/windawings/opentaobao/api/wlb](https://pkg.go.dev/github.com/windawings/opentaobao/api/wlb) |
| [x] [直通车API](https://open.taobao.com/API.htm?docType=2&docId=10551) | [github.com/windawings/opentaobao/api/simba](https://pkg.go.dev/github.com/windawings/opentaobao/api/simba) |
| [x] [机票API](https://open.taobao.com/API.htm?docType=2&docId=64841) | [github.com/windawings/opentaobao/api/flight](https://pkg.go.dev/github.com/windawings/opentaobao/api/flight) |
| [x] [ONS消息服务](https://open.taobao.com/API.htm?docType=2&docId=25634) | [github.com/windawings/opentaobao/api/jms](https://pkg.go.dev/github.com/windawings/opentaobao/api/jms) |
| [x] [营销API](https://open.taobao.com/API.htm?docType=2&docId=50923) | [github.com/windawings/opentaobao/api/promotion](https://pkg.go.dev/github.com/windawings/opentaobao/api/promotion) |
| [x] [数据API](https://open.taobao.com/API.htm?docType=2&docId=46786) | [github.com/windawings/opentaobao/api/dt](https://pkg.go.dev/github.com/windawings/opentaobao/api/dt) |
| [x] [酒店API](https://open.taobao.com/API.htm?docType=2&docId=49115) | [github.com/windawings/opentaobao/api/xhotel](https://pkg.go.dev/github.com/windawings/opentaobao/api/xhotel) |
| [x] [聚划算API](https://open.taobao.com/API.htm?docType=2&docId=65912) | [github.com/windawings/opentaobao/api/ju](https://pkg.go.dev/github.com/windawings/opentaobao/api/ju) |
| [x] [店铺会员管理API](https://open.taobao.com/API.htm?docType=2&docId=25584) | [github.com/windawings/opentaobao/api/crm](https://pkg.go.dev/github.com/windawings/opentaobao/api/crm) |
| [x] [多媒体平台API](https://open.taobao.com/API.htm?docType=2&docId=62392) | [github.com/windawings/opentaobao/api/media](https://pkg.go.dev/github.com/windawings/opentaobao/api/media) |
| [x] [子账号管理API](https://open.taobao.com/API.htm?docType=2&docId=10820) | [github.com/windawings/opentaobao/api/subuser](https://pkg.go.dev/github.com/windawings/opentaobao/api/subuser) |
| [x] [服务平台API](https://open.taobao.com/API.htm?docType=2&docId=25687) | [github.com/windawings/opentaobao/api/servicecenter](https://pkg.go.dev/github.com/windawings/opentaobao/api/servicecenter) |
| [x] [退款API](https://open.taobao.com/API.htm?docType=2&docId=46784) | [github.com/windawings/opentaobao/api/refund](https://pkg.go.dev/github.com/windawings/opentaobao/api/refund) |
| [x] [质检品控API](https://open.taobao.com/API.htm?docType=2&docId=10902) | [github.com/windawings/opentaobao/api/qt](https://pkg.go.dev/github.com/windawings/opentaobao/api/qt) |
| [x] [天猫服务商品API](https://open.taobao.com/API.htm?docType=2&docId=49044) | [github.com/windawings/opentaobao/api/tmallservice](https://pkg.go.dev/github.com/windawings/opentaobao/api/tmallservice) |
| [x] [天猫精品库API](https://open.taobao.com/API.htm?docType=2&docId=21643) | [github.com/windawings/opentaobao/api/tmallitem](https://pkg.go.dev/github.com/windawings/opentaobao/api/tmallitem) |
| [x] [聚石塔API](https://open.taobao.com/API.htm?docType=2&docId=55254) | [github.com/windawings/opentaobao/api/jst](https://pkg.go.dev/github.com/windawings/opentaobao/api/jst) |
| [x] [电子物流API](https://open.taobao.com/API.htm?docType=2&docId=27156) | [github.com/windawings/opentaobao/api/eticket](https://pkg.go.dev/github.com/windawings/opentaobao/api/eticket) |
| [x] [彩票API](https://open.taobao.com/API.htm?docType=2&docId=21828) | [github.com/windawings/opentaobao/api/caipiao](https://pkg.go.dev/github.com/windawings/opentaobao/api/caipiao) |
| [x] [账务API](https://open.taobao.com/API.htm?docType=2&docId=21709) | [github.com/windawings/opentaobao/api/bill](https://pkg.go.dev/github.com/windawings/opentaobao/api/bill) |
| [x] [拍卖API](https://open.taobao.com/API.htm?docType=2&docId=64801) | [github.com/windawings/opentaobao/api/paimai](https://pkg.go.dev/github.com/windawings/opentaobao/api/paimai) |
| [x] [千牛接口](https://open.taobao.com/API.htm?docType=2&docId=27244) | [github.com/windawings/opentaobao/api/qianniu](https://pkg.go.dev/github.com/windawings/opentaobao/api/qianniu) |
| [x] [消息服务API](https://open.taobao.com/API.htm?docType=2&docId=53697) | [github.com/windawings/opentaobao/api/tmc](https://pkg.go.dev/github.com/windawings/opentaobao/api/tmc) |
| [x] [本地生活API](https://open.taobao.com/API.htm?docType=2&docId=57060) | [github.com/windawings/opentaobao/api/alsc](https://pkg.go.dev/github.com/windawings/opentaobao/api/alsc) |
| [x] [阿里云ocsAPI](https://open.taobao.com/API.htm?docType=2&docId=24698) | [github.com/windawings/opentaobao/api/aliyunocs](https://pkg.go.dev/github.com/windawings/opentaobao/api/aliyunocs) |
| [x] [淘宝同城API](https://open.taobao.com/API.htm?docType=2&docId=51645) | [github.com/windawings/opentaobao/api/cityretail](https://pkg.go.dev/github.com/windawings/opentaobao/api/cityretail) |
| [x] [YunOS](https://open.taobao.com/API.htm?docType=2&docId=33101) | [github.com/windawings/opentaobao/api/yunos](https://pkg.go.dev/github.com/windawings/opentaobao/api/yunos) |
| [x] [阿里云API](https://open.taobao.com/API.htm?docType=2&docId=24304) | [github.com/windawings/opentaobao/api/aliyun](https://pkg.go.dev/github.com/windawings/opentaobao/api/aliyun) |
| [x] [火车票API](https://open.taobao.com/API.htm?docType=2&docId=29812) | [github.com/windawings/opentaobao/api/train](https://pkg.go.dev/github.com/windawings/opentaobao/api/train) |
| [x] [Tanx API](https://open.taobao.com/API.htm?docType=2&docId=25062) | [github.com/windawings/opentaobao/api/tanx](https://pkg.go.dev/github.com/windawings/opentaobao/api/tanx) |
| [x] [手淘开放API](https://open.taobao.com/API.htm?docType=2&docId=26114) | [github.com/windawings/opentaobao/api/mtopopen](https://pkg.go.dev/github.com/windawings/opentaobao/api/mtopopen) |
| [x] [JAE](https://open.taobao.com/API.htm?docType=2&docId=30901) | [github.com/windawings/opentaobao/api/jae](https://pkg.go.dev/github.com/windawings/opentaobao/api/jae) |
| [x] [宝点API](https://open.taobao.com/API.htm?docType=2&docId=22800) | [github.com/windawings/opentaobao/api/baodian](https://pkg.go.dev/github.com/windawings/opentaobao/api/baodian) |
| [x] [天猫会员积分](https://open.taobao.com/API.htm?docType=2&docId=25758) | [github.com/windawings/opentaobao/api/tmallcms](https://pkg.go.dev/github.com/windawings/opentaobao/api/tmallcms) |
| [x] [汽车票API](https://open.taobao.com/API.htm?docType=2&docId=61877) | [github.com/windawings/opentaobao/api/bus](https://pkg.go.dev/github.com/windawings/opentaobao/api/bus) |
| [x] [码上淘API](https://open.taobao.com/API.htm?docType=2&docId=23660) | [github.com/windawings/opentaobao/api/ma](https://pkg.go.dev/github.com/windawings/opentaobao/api/ma) |
| [x] [游戏激励平台API](https://open.taobao.com/API.htm?docType=2&docId=25396) | [github.com/windawings/opentaobao/api/gameact](https://pkg.go.dev/github.com/windawings/opentaobao/api/gameact) |
| [x] [淘宝抽奖平台API](https://open.taobao.com/API.htm?docType=2&docId=23213) | [github.com/windawings/opentaobao/api/choujiang](https://pkg.go.dev/github.com/windawings/opentaobao/api/choujiang) |
| [x] [天猫国际API](https://open.taobao.com/API.htm?docType=2&docId=67532) | [github.com/windawings/opentaobao/api/tmallhk](https://pkg.go.dev/github.com/windawings/opentaobao/api/tmallhk) |
| [x] [司法拍卖](https://open.taobao.com/API.htm?docType=2&docId=40551) | [github.com/windawings/opentaobao/api/auction](https://pkg.go.dev/github.com/windawings/opentaobao/api/auction) |
| [x] [虾米API](https://open.taobao.com/API.htm?docType=2&docId=23270) | [github.com/windawings/opentaobao/api/xiami](https://pkg.go.dev/github.com/windawings/opentaobao/api/xiami) |
| [x] [天猫互动接口](https://open.taobao.com/API.htm?docType=2&docId=65074) | [github.com/windawings/opentaobao/api/interact](https://pkg.go.dev/github.com/windawings/opentaobao/api/interact) |
| [x] [DMP API](https://open.taobao.com/API.htm?docType=2&docId=61644) | [github.com/windawings/opentaobao/api/dmp](https://pkg.go.dev/github.com/windawings/opentaobao/api/dmp) |
| [x] [生活服务API](https://open.taobao.com/API.htm?docType=2&docId=50111) | [github.com/windawings/opentaobao/api/lifeservice](https://pkg.go.dev/github.com/windawings/opentaobao/api/lifeservice) |
| [x] [手机淘宝API](https://open.taobao.com/API.htm?docType=2&docId=23431) | [github.com/windawings/opentaobao/api/mtop](https://pkg.go.dev/github.com/windawings/opentaobao/api/mtop) |
| [x] [物联API](https://open.taobao.com/API.htm?docType=2&docId=28313) | [github.com/windawings/opentaobao/api/alink](https://pkg.go.dev/github.com/windawings/opentaobao/api/alink) |
| [x] [酒店导购API](https://open.taobao.com/API.htm?docType=2&docId=35634) | [github.com/windawings/opentaobao/api/hotel](https://pkg.go.dev/github.com/windawings/opentaobao/api/hotel) |
| [x] [保险API](https://open.taobao.com/API.htm?docType=2&docId=37741) | [github.com/windawings/opentaobao/api/baoxian](https://pkg.go.dev/github.com/windawings/opentaobao/api/baoxian) |
| [x] [天猫美妆API](https://open.taobao.com/API.htm?docType=2&docId=25013) | [github.com/windawings/opentaobao/api/mei](https://pkg.go.dev/github.com/windawings/opentaobao/api/mei) |
| [x] [会员卡](https://open.taobao.com/API.htm?docType=2&docId=46278) | [github.com/windawings/opentaobao/api/blackvip](https://pkg.go.dev/github.com/windawings/opentaobao/api/blackvip) |
| [x] [电子面单API](https://open.taobao.com/API.htm?docType=2&docId=27111) | [github.com/windawings/opentaobao/api/waybill](https://pkg.go.dev/github.com/windawings/opentaobao/api/waybill) |
| [x] [电影票API](https://open.taobao.com/API.htm?docType=2&docId=45077) | [github.com/windawings/opentaobao/api/film](https://pkg.go.dev/github.com/windawings/opentaobao/api/film) |
| [x] [阿里通信API](https://open.taobao.com/API.htm?docType=2&docId=31279) | [github.com/windawings/opentaobao/api/alicom](https://pkg.go.dev/github.com/windawings/opentaobao/api/alicom) |
| [x] [openimAPI](https://open.taobao.com/API.htm?docType=2&docId=25766) | [github.com/windawings/opentaobao/api/openim](https://pkg.go.dev/github.com/windawings/opentaobao/api/openim) |
| [x] [阿里车联网API](https://open.taobao.com/API.htm?docType=2&docId=24740) | [github.com/windawings/opentaobao/api/autonavi](https://pkg.go.dev/github.com/windawings/opentaobao/api/autonavi) |
| [x] [虚拟院线API](https://open.taobao.com/API.htm?docType=2&docId=35818) | [github.com/windawings/opentaobao/api/taotv](https://pkg.go.dev/github.com/windawings/opentaobao/api/taotv) |
| [x] [知识库API](https://open.taobao.com/API.htm?docType=2&docId=38771) | [github.com/windawings/opentaobao/api/kclub](https://pkg.go.dev/github.com/windawings/opentaobao/api/kclub) |
| [x] [反欺诈风控API](https://open.taobao.com/API.htm?docType=2&docId=25505) | [github.com/windawings/opentaobao/api/antifraud](https://pkg.go.dev/github.com/windawings/opentaobao/api/antifraud) |
| [x] [国际站外贸直通车API](https://open.taobao.com/API.htm?docType=2&docId=25200) | [github.com/windawings/opentaobao/api/scbp](https://pkg.go.dev/github.com/windawings/opentaobao/api/scbp) |
| [x] [天猫服务数据API](https://open.taobao.com/API.htm?docType=2&docId=53341) | [github.com/windawings/opentaobao/api/tmallsc](https://pkg.go.dev/github.com/windawings/opentaobao/api/tmallsc) |
| [x] [智能设备](https://open.taobao.com/API.htm?docType=2&docId=40674) | [github.com/windawings/opentaobao/api/iot](https://pkg.go.dev/github.com/windawings/opentaobao/api/iot) |
| [x] [百川](https://open.taobao.com/API.htm?docType=2&docId=31057) | [github.com/windawings/opentaobao/api/baichuan](https://pkg.go.dev/github.com/windawings/opentaobao/api/baichuan) |
| [x] [文本算法API](https://open.taobao.com/API.htm?docType=2&docId=26130) | [github.com/windawings/opentaobao/api/nlp](https://pkg.go.dev/github.com/windawings/opentaobao/api/nlp) |
| [x] [百川推送](https://open.taobao.com/API.htm?docType=2&docId=25038) | [github.com/windawings/opentaobao/api/cloudpush](https://pkg.go.dev/github.com/windawings/opentaobao/api/cloudpush) |
| [x] [国际站商品API](https://open.taobao.com/API.htm?docType=2&docId=25348) | [github.com/windawings/opentaobao/api/icbu](https://pkg.go.dev/github.com/windawings/opentaobao/api/icbu) |
| [x] [淘宝游戏API](https://open.taobao.com/API.htm?docType=2&docId=45035) | [github.com/windawings/opentaobao/api/game](https://pkg.go.dev/github.com/windawings/opentaobao/api/game) |
| [x] [聚安全API](https://open.taobao.com/API.htm?docType=2&docId=61781) | [github.com/windawings/opentaobao/api/security](https://pkg.go.dev/github.com/windawings/opentaobao/api/security) |
| [x] [喵街API](https://open.taobao.com/API.htm?docType=2&docId=55338) | [github.com/windawings/opentaobao/api/mos](https://pkg.go.dev/github.com/windawings/opentaobao/api/mos) |
| [x] [菜鸟配送API](https://open.taobao.com/API.htm?docType=2&docId=25659) | [github.com/windawings/opentaobao/api/cntms](https://pkg.go.dev/github.com/windawings/opentaobao/api/cntms) |
| [x] [菜鸟仓配API](https://open.taobao.com/API.htm?docType=2&docId=26245) | [github.com/windawings/opentaobao/api/wms](https://pkg.go.dev/github.com/windawings/opentaobao/api/wms) |
| [x] [网上法庭对外API](https://open.taobao.com/API.htm?docType=2&docId=27772) | [github.com/windawings/opentaobao/api/nazca](https://pkg.go.dev/github.com/windawings/opentaobao/api/nazca) |
| [x] [五道口API](https://open.taobao.com/API.htm?docType=2&docId=66186) | [github.com/windawings/opentaobao/api/wdk](https://pkg.go.dev/github.com/windawings/opentaobao/api/wdk) |
| [x] [阿里大于API](https://open.taobao.com/API.htm?docType=2&docId=40274) | [github.com/windawings/opentaobao/api/aliqin](https://pkg.go.dev/github.com/windawings/opentaobao/api/aliqin) |
| [x] [淘宝内容API](https://open.taobao.com/API.htm?docType=2&docId=27196) | [github.com/windawings/opentaobao/api/beehive](https://pkg.go.dev/github.com/windawings/opentaobao/api/beehive) |
| [x] [旅行用车API](https://open.taobao.com/API.htm?docType=2&docId=61845) | [github.com/windawings/opentaobao/api/car](https://pkg.go.dev/github.com/windawings/opentaobao/api/car) |
| [x] [门票-商品管理API](https://open.taobao.com/API.htm?docType=2&docId=27945) | [github.com/windawings/opentaobao/api/ticket](https://pkg.go.dev/github.com/windawings/opentaobao/api/ticket) |
| [x] [菜鸟无线API](https://open.taobao.com/API.htm?docType=2&docId=26056) | [github.com/windawings/opentaobao/api/guoguo](https://pkg.go.dev/github.com/windawings/opentaobao/api/guoguo) |
| [x] [奇门仓储API](https://open.taobao.com/API.htm?docType=2&docId=29364) | [github.com/windawings/opentaobao/api/qimen](https://pkg.go.dev/github.com/windawings/opentaobao/api/qimen) |
| [x] [yunos推送服务api](https://open.taobao.com/API.htm?docType=2&docId=49740) | [github.com/windawings/opentaobao/api/cmns](https://pkg.go.dev/github.com/windawings/opentaobao/api/cmns) |
| [x] [生活汇API](https://open.taobao.com/API.htm?docType=2&docId=26189) | [github.com/windawings/opentaobao/api/elife](https://pkg.go.dev/github.com/windawings/opentaobao/api/elife) |
| [x] [tv支付](https://open.taobao.com/API.htm?docType=2&docId=26205) | [github.com/windawings/opentaobao/api/tvpay](https://pkg.go.dev/github.com/windawings/opentaobao/api/tvpay) |
| [x] [菜鸟集货API](https://open.taobao.com/API.htm?docType=2&docId=60752) | [github.com/windawings/opentaobao/api/wlbimports](https://pkg.go.dev/github.com/windawings/opentaobao/api/wlbimports) |
| [x] [阿里健康医](https://open.taobao.com/API.htm?docType=2&docId=42645) | [github.com/windawings/opentaobao/api/alihealth](https://pkg.go.dev/github.com/windawings/opentaobao/api/alihealth) |
| [x] [地动仪](https://open.taobao.com/API.htm?docType=2&docId=37287) | [github.com/windawings/opentaobao/api/lbs](https://pkg.go.dev/github.com/windawings/opentaobao/api/lbs) |
| [x] [阿里健康药API](https://open.taobao.com/API.htm?docType=2&docId=50465) | [github.com/windawings/opentaobao/api/drug](https://pkg.go.dev/github.com/windawings/opentaobao/api/drug) |
| [x] [手淘分享](https://open.taobao.com/API.htm?docType=2&docId=32461) | [github.com/windawings/opentaobao/api/wirelessshare](https://pkg.go.dev/github.com/windawings/opentaobao/api/wirelessshare) |
| [x] [法务诉讼对外API](https://open.taobao.com/API.htm?docType=2&docId=53340) | [github.com/windawings/opentaobao/api/legalcase](https://pkg.go.dev/github.com/windawings/opentaobao/api/legalcase) |
| [x] [度假-商品管理API](https://open.taobao.com/API.htm?docType=2&docId=28132) | [github.com/windawings/opentaobao/api/travel](https://pkg.go.dev/github.com/windawings/opentaobao/api/travel) |
| [x] [酒店商品API](https://open.taobao.com/API.htm?docType=2&docId=67808) | [github.com/windawings/opentaobao/api/xhotelitem](https://pkg.go.dev/github.com/windawings/opentaobao/api/xhotelitem) |
| [x] [酒店在线预订API](https://open.taobao.com/API.htm?docType=2&docId=48716) | [github.com/windawings/opentaobao/api/xhotelonlineorder](https://pkg.go.dev/github.com/windawings/opentaobao/api/xhotelonlineorder) |
| [x] [酒店官网信用住API](https://open.taobao.com/API.htm?docType=2&docId=26073) | [github.com/windawings/opentaobao/api/xhotelofficial](https://pkg.go.dev/github.com/windawings/opentaobao/api/xhotelofficial) |
| [x] [酒店线下信用住API](https://open.taobao.com/API.htm?docType=2&docId=26074) | [github.com/windawings/opentaobao/api/xhoteloffline](https://pkg.go.dev/github.com/windawings/opentaobao/api/xhoteloffline) |
| [x] [全渠道API](https://open.taobao.com/API.htm?docType=2&docId=42927) | [github.com/windawings/opentaobao/api/omniorder](https://pkg.go.dev/github.com/windawings/opentaobao/api/omniorder) |
| [x] [国际机票政策API](https://open.taobao.com/API.htm?docType=2&docId=25494) | [github.com/windawings/opentaobao/api/itpolicy](https://pkg.go.dev/github.com/windawings/opentaobao/api/itpolicy) |
| [x] [国际机票订单API](https://open.taobao.com/API.htm?docType=2&docId=27905) | [github.com/windawings/opentaobao/api/ieagency](https://pkg.go.dev/github.com/windawings/opentaobao/api/ieagency) |
| [x] [国内机票订单API](https://open.taobao.com/API.htm?docType=2&docId=26565) | [github.com/windawings/opentaobao/api/jipiao](https://pkg.go.dev/github.com/windawings/opentaobao/api/jipiao) |
| [x] [度假&门票-交易管理API](https://open.taobao.com/API.htm?docType=2&docId=52211) | [github.com/windawings/opentaobao/api/traveltrade](https://pkg.go.dev/github.com/windawings/opentaobao/api/traveltrade) |
| [x] [阿里体育API](https://open.taobao.com/API.htm?docType=2&docId=40644) | [github.com/windawings/opentaobao/api/alisports](https://pkg.go.dev/github.com/windawings/opentaobao/api/alisports) |
| [x] [电子发票](https://open.taobao.com/API.htm?docType=2&docId=27764) | [github.com/windawings/opentaobao/api/einvoice](https://pkg.go.dev/github.com/windawings/opentaobao/api/einvoice) |
| [x] [阿里翻译API](https://open.taobao.com/API.htm?docType=2&docId=47571) | [github.com/windawings/opentaobao/api/seaking](https://pkg.go.dev/github.com/windawings/opentaobao/api/seaking) |
| [x] [国际站数据管](https://open.taobao.com/API.htm?docType=2&docId=26908) | [github.com/windawings/opentaobao/api/mydata](https://pkg.go.dev/github.com/windawings/opentaobao/api/mydata) |
| [x] [tmall-carcenter](https://open.taobao.com/API.htm?docType=2&docId=31209) | [github.com/windawings/opentaobao/api/tmallcarenter](https://pkg.go.dev/github.com/windawings/opentaobao/api/tmallcarenter) |
| [x] [阿里妈妈广告中心API](https://open.taobao.com/API.htm?docType=2&docId=58946) | [github.com/windawings/opentaobao/api/scs](https://pkg.go.dev/github.com/windawings/opentaobao/api/scs) |
| [x] [Yunos Account API](https://open.taobao.com/API.htm?docType=2&docId=27319) | [github.com/windawings/opentaobao/api/yunosaccount](https://pkg.go.dev/github.com/windawings/opentaobao/api/yunosaccount) |
| [x] [菜鸟裹裹API](https://open.taobao.com/API.htm?docType=2&docId=31197) | [github.com/windawings/opentaobao/api/cainiaolocker](https://pkg.go.dev/github.com/windawings/opentaobao/api/cainiaolocker) |
| [x] [度假-签证管理API](https://open.taobao.com/API.htm?docType=2&docId=44453) | [github.com/windawings/opentaobao/api/normalvisa](https://pkg.go.dev/github.com/windawings/opentaobao/api/normalvisa) |
| [x] [菜鸟农村物流](https://open.taobao.com/API.htm?docType=2&docId=27588) | [github.com/windawings/opentaobao/api/cainiaoncwl](https://pkg.go.dev/github.com/windawings/opentaobao/api/cainiaoncwl) |
| [x] [体检机构API](https://open.taobao.com/API.htm?docType=2&docId=41650) | [github.com/windawings/opentaobao/api/examination](https://pkg.go.dev/github.com/windawings/opentaobao/api/examination) |
| [x] [1688推客API](https://open.taobao.com/API.htm?docType=2&docId=27457) | [github.com/windawings/opentaobao/api/tuike](https://pkg.go.dev/github.com/windawings/opentaobao/api/tuike) |
| [x] [商户API](https://open.taobao.com/API.htm?docType=2&docId=50163) | [github.com/windawings/opentaobao/api/store](https://pkg.go.dev/github.com/windawings/opentaobao/api/store) |
| [x] [桌面API](https://open.taobao.com/API.htm?docType=2&docId=29910) | [github.com/windawings/opentaobao/api/ott](https://pkg.go.dev/github.com/windawings/opentaobao/api/ott) |
| [x] [电动车API](https://open.taobao.com/API.htm?docType=2&docId=30434) | [github.com/windawings/opentaobao/api/vms](https://pkg.go.dev/github.com/windawings/opentaobao/api/vms) |
| [x] [新零售API](https://open.taobao.com/API.htm?docType=2&docId=44781) | [github.com/windawings/opentaobao/api/newretail](https://pkg.go.dev/github.com/windawings/opentaobao/api/newretail) |
| [x] [SCM API](https://open.taobao.com/API.htm?docType=2&docId=36216) | [github.com/windawings/opentaobao/api/ascm](https://pkg.go.dev/github.com/windawings/opentaobao/api/ascm) |
| [x] [百川-ctg](https://open.taobao.com/API.htm?docType=2&docId=28682) | [github.com/windawings/opentaobao/api/baichuanctg](https://pkg.go.dev/github.com/windawings/opentaobao/api/baichuanctg) |
| [x] [汇金API](https://open.taobao.com/API.htm?docType=2&docId=65164) | [github.com/windawings/opentaobao/api/fundplatform](https://pkg.go.dev/github.com/windawings/opentaobao/api/fundplatform) |
| [x] [数娱媒资输出](https://open.taobao.com/API.htm?docType=2&docId=41681) | [github.com/windawings/opentaobao/api/wenyuvideo](https://pkg.go.dev/github.com/windawings/opentaobao/api/wenyuvideo) |
| [x] [商旅API](https://open.taobao.com/API.htm?docType=2&docId=33169) | [github.com/windawings/opentaobao/api/btrip](https://pkg.go.dev/github.com/windawings/opentaobao/api/btrip) |
| [x] [零售plus](https://open.taobao.com/API.htm?docType=2&docId=29320) | [github.com/windawings/opentaobao/api/nlife](https://pkg.go.dev/github.com/windawings/opentaobao/api/nlife) |
| [x] [天猫门店API](https://open.taobao.com/API.htm?docType=2&docId=41298) | [github.com/windawings/opentaobao/api/retail](https://pkg.go.dev/github.com/windawings/opentaobao/api/retail) |
| [x] [AILAB图像算法API](https://open.taobao.com/API.htm?docType=2&docId=37057) | [github.com/windawings/opentaobao/api/aiar](https://pkg.go.dev/github.com/windawings/opentaobao/api/aiar) |
| [x] [天猫供应](https://open.taobao.com/API.htm?docType=2&docId=65940) | [github.com/windawings/opentaobao/api/ascp](https://pkg.go.dev/github.com/windawings/opentaobao/api/ascp) |
| [x] [欢行开发平台API](https://open.taobao.com/API.htm?docType=2&docId=47501) | [github.com/windawings/opentaobao/api/happytrip](https://pkg.go.dev/github.com/windawings/opentaobao/api/happytrip) |
| [x] [跨境API](https://open.taobao.com/API.htm?docType=2&docId=28703) | [github.com/windawings/opentaobao/api/oversea](https://pkg.go.dev/github.com/windawings/opentaobao/api/oversea) |
| [x] [迎客松牌照审核接口](https://open.taobao.com/API.htm?docType=2&docId=30410) | [github.com/windawings/opentaobao/api/tvupadmin](https://pkg.go.dev/github.com/windawings/opentaobao/api/tvupadmin) |
| [x] [IoTI API](https://open.taobao.com/API.htm?docType=2&docId=46348) | [github.com/windawings/opentaobao/api/ioti](https://pkg.go.dev/github.com/windawings/opentaobao/api/ioti) |
| [x] [淘宝卡券平台](https://open.taobao.com/API.htm?docType=2&docId=29874) | [github.com/windawings/opentaobao/api/deliveryvoucher](https://pkg.go.dev/github.com/windawings/opentaobao/api/deliveryvoucher) |
| [x] [智慧门店](https://open.taobao.com/API.htm?docType=2&docId=39618) | [github.com/windawings/opentaobao/api/smartstore](https://pkg.go.dev/github.com/windawings/opentaobao/api/smartstore) |
| [x] [淘宝定制行业API](https://open.taobao.com/API.htm?docType=2&docId=62678) | [github.com/windawings/opentaobao/api/customizemarket](https://pkg.go.dev/github.com/windawings/opentaobao/api/customizemarket) |
| [x] [闲鱼](https://open.taobao.com/API.htm?docType=2&docId=67023) | [github.com/windawings/opentaobao/api/idle](https://pkg.go.dev/github.com/windawings/opentaobao/api/idle) |
| [x] [库存API](https://open.taobao.com/API.htm?docType=2&docId=31754) | [github.com/windawings/opentaobao/api/inventory](https://pkg.go.dev/github.com/windawings/opentaobao/api/inventory) |
| [x] [aliExpress](https://open.taobao.com/API.htm?docType=2&docId=54728) | [github.com/windawings/opentaobao/api/aliexpress](https://pkg.go.dev/github.com/windawings/opentaobao/api/aliexpress) |
| [x] [YunOS-广告](https://open.taobao.com/API.htm?docType=2&docId=29086) | [github.com/windawings/opentaobao/api/yunosad](https://pkg.go.dev/github.com/windawings/opentaobao/api/yunosad) |
| [x] [司法开放平台](https://open.taobao.com/API.htm?docType=2&docId=65037) | [github.com/windawings/opentaobao/api/legalsuit](https://pkg.go.dev/github.com/windawings/opentaobao/api/legalsuit) |
| [x] [全球速卖通](https://open.taobao.com/API.htm?docType=2&docId=56229) | [github.com/windawings/opentaobao/api/aliexpresssumaitong](https://pkg.go.dev/github.com/windawings/opentaobao/api/aliexpresssumaitong) |
| [x] [阿里健康追溯码](https://open.taobao.com/API.htm?docType=2&docId=64867) | [github.com/windawings/opentaobao/api/drugtrace](https://pkg.go.dev/github.com/windawings/opentaobao/api/drugtrace) |
| [x] [会员中心API](https://open.taobao.com/API.htm?docType=2&docId=29108) | [github.com/windawings/opentaobao/api/interactvip](https://pkg.go.dev/github.com/windawings/opentaobao/api/interactvip) |
| [x] [阿里健康-会员管理](https://open.taobao.com/API.htm?docType=2&docId=46015) | [github.com/windawings/opentaobao/api/alihealthcrm](https://pkg.go.dev/github.com/windawings/opentaobao/api/alihealthcrm) |
| [x] [品效API](https://open.taobao.com/API.htm?docType=2&docId=36460) | [github.com/windawings/opentaobao/api/brandhub](https://pkg.go.dev/github.com/windawings/opentaobao/api/brandhub) |
| [x] [小蜜API](https://open.taobao.com/API.htm?docType=2&docId=44065) | [github.com/windawings/opentaobao/api/alime](https://pkg.go.dev/github.com/windawings/opentaobao/api/alime) |
| [x] [大麦票务云分销API](https://open.taobao.com/API.htm?docType=2&docId=49119) | [github.com/windawings/opentaobao/api/maitix](https://pkg.go.dev/github.com/windawings/opentaobao/api/maitix) |
| [x] [天猫精灵开放API](https://open.taobao.com/API.htm?docType=2&docId=66106) | [github.com/windawings/opentaobao/api/tmallgenie](https://pkg.go.dev/github.com/windawings/opentaobao/api/tmallgenie) |
| [x] [ICBU卖家API](https://open.taobao.com/API.htm?docType=2&docId=51234) | [github.com/windawings/opentaobao/api/icbuseller](https://pkg.go.dev/github.com/windawings/opentaobao/api/icbuseller) |
| [x] [智慧园区API](https://open.taobao.com/API.htm?docType=2&docId=32226) | [github.com/windawings/opentaobao/api/campus](https://pkg.go.dev/github.com/windawings/opentaobao/api/campus) |
| [x] [酒店会员API](https://open.taobao.com/API.htm?docType=2&docId=64794) | [github.com/windawings/opentaobao/api/xhotelcrm](https://pkg.go.dev/github.com/windawings/opentaobao/api/xhotelcrm) |
| [x] [大麦API](https://open.taobao.com/API.htm?docType=2&docId=39475) | [github.com/windawings/opentaobao/api/damai](https://pkg.go.dev/github.com/windawings/opentaobao/api/damai) |
| [x] [换货API](https://open.taobao.com/API.htm?docType=2&docId=31202) | [github.com/windawings/opentaobao/api/exchange](https://pkg.go.dev/github.com/windawings/opentaobao/api/exchange) |
| [x] [商家营销中心API](https://open.taobao.com/API.htm?docType=2&docId=31352) | [github.com/windawings/opentaobao/api/singletreasure](https://pkg.go.dev/github.com/windawings/opentaobao/api/singletreasure) |
| [x] [ICBU－物流](https://open.taobao.com/API.htm?docType=2&docId=47959) | [github.com/windawings/opentaobao/api/icbulogistics](https://pkg.go.dev/github.com/windawings/opentaobao/api/icbulogistics) |
| [x] [ICBU－RFQ](https://open.taobao.com/API.htm?docType=2&docId=40805) | [github.com/windawings/opentaobao/api/icburfq](https://pkg.go.dev/github.com/windawings/opentaobao/api/icburfq) |
| [x] [ICBU－信保](https://open.taobao.com/API.htm?docType=2&docId=31686) | [github.com/windawings/opentaobao/api/icbuassurance](https://pkg.go.dev/github.com/windawings/opentaobao/api/icbuassurance) |
| [x] [飞猪POI数据API](https://open.taobao.com/API.htm?docType=2&docId=45315) | [github.com/windawings/opentaobao/api/alitrippoi](https://pkg.go.dev/github.com/windawings/opentaobao/api/alitrippoi) |
| [x] [飞猪行政区划API](https://open.taobao.com/API.htm?docType=2&docId=31375) | [github.com/windawings/opentaobao/api/alitripdivisions](https://pkg.go.dev/github.com/windawings/opentaobao/api/alitripdivisions) |
| [x] [零售终端API](https://open.taobao.com/API.htm?docType=2&docId=42299) | [github.com/windawings/opentaobao/api/nrt](https://pkg.go.dev/github.com/windawings/opentaobao/api/nrt) |
| [x] [手淘用户增](https://open.taobao.com/API.htm?docType=2&docId=66356) | [github.com/windawings/opentaobao/api/usergrowth](https://pkg.go.dev/github.com/windawings/opentaobao/api/usergrowth) |
| [x] [ALiOS应用中心](https://open.taobao.com/API.htm?docType=2&docId=35428) | [github.com/windawings/opentaobao/api/yunosappstore](https://pkg.go.dev/github.com/windawings/opentaobao/api/yunosappstore) |
| [x] [b2b认证平台api](https://open.taobao.com/API.htm?docType=2&docId=32288) | [github.com/windawings/opentaobao/api/b2bcert](https://pkg.go.dev/github.com/windawings/opentaobao/api/b2bcert) |
| [x] [企业运营平台-集团财](https://open.taobao.com/API.htm?docType=2&docId=61521) | [github.com/windawings/opentaobao/api/fpm](https://pkg.go.dev/github.com/windawings/opentaobao/api/fpm) |
| [x] [零售通智能POS开放](https://open.taobao.com/API.htm?docType=2&docId=32945) | [github.com/windawings/opentaobao/api/lstpos](https://pkg.go.dev/github.com/windawings/opentaobao/api/lstpos) |
| [x] [阿里巴巴供应链平台API](https://open.taobao.com/API.htm?docType=2&docId=32779) | [github.com/windawings/opentaobao/api/ascpqcc](https://pkg.go.dev/github.com/windawings/opentaobao/api/ascpqcc) |
| [x] [天猫新零售](https://open.taobao.com/API.htm?docType=2&docId=55242) | [github.com/windawings/opentaobao/api/tmallnr](https://pkg.go.dev/github.com/windawings/opentaobao/api/tmallnr) |
| [x] [酒店团购套餐API](https://open.taobao.com/API.htm?docType=2&docId=45432) | [github.com/windawings/opentaobao/api/tuanhotel](https://pkg.go.dev/github.com/windawings/opentaobao/api/tuanhotel) |
| [x] [阿里影业云智API](https://open.taobao.com/API.htm?docType=2&docId=42042) | [github.com/windawings/opentaobao/api/larkiot](https://pkg.go.dev/github.com/windawings/opentaobao/api/larkiot) |
| [x] [渠道中心API](https://open.taobao.com/API.htm?docType=2&docId=40706) | [github.com/windawings/opentaobao/api/tmallchannel](https://pkg.go.dev/github.com/windawings/opentaobao/api/tmallchannel) |
| [x] [阿里健康-疫苗API](https://open.taobao.com/API.htm?docType=2&docId=58793) | [github.com/windawings/opentaobao/api/vaccin](https://pkg.go.dev/github.com/windawings/opentaobao/api/vaccin) |
| [x] [教育账号 API](https://open.taobao.com/API.htm?docType=2&docId=35502) | [github.com/windawings/opentaobao/api/yunosdm](https://pkg.go.dev/github.com/windawings/opentaobao/api/yunosdm) |
| [x] [用户增长](https://open.taobao.com/API.htm?docType=2&docId=60596) | [github.com/windawings/opentaobao/api/usergrowth2](https://pkg.go.dev/github.com/windawings/opentaobao/api/usergrowth2) |
| [x] [AE-供应链](https://open.taobao.com/API.htm?docType=2&docId=56317) | [github.com/windawings/opentaobao/api/ascpffo](https://pkg.go.dev/github.com/windawings/opentaobao/api/ascpffo) |
| [x] [互动吧API](https://open.taobao.com/API.htm?docType=2&docId=44032) | [github.com/windawings/opentaobao/api/fans](https://pkg.go.dev/github.com/windawings/opentaobao/api/fans) |
| [x] [飞猪机票前台类目](https://open.taobao.com/API.htm?docType=2&docId=42575) | [github.com/windawings/opentaobao/api/flightuppc](https://pkg.go.dev/github.com/windawings/opentaobao/api/flightuppc) |
| [x] [平台治理API](https://open.taobao.com/API.htm?docType=2&docId=32353) | [github.com/windawings/opentaobao/api/sungari](https://pkg.go.dev/github.com/windawings/opentaobao/api/sungari) |
| [x] [天猫超市前台API](https://open.taobao.com/API.htm?docType=2&docId=45441) | [github.com/windawings/opentaobao/api/txcs](https://pkg.go.dev/github.com/windawings/opentaobao/api/txcs) |
| [x] [虾米开放平台](https://open.taobao.com/API.htm?docType=2&docId=36051) | [github.com/windawings/opentaobao/api/xiamiopen](https://pkg.go.dev/github.com/windawings/opentaobao/api/xiamiopen) |
| [x] [阿里健康医生](https://open.taobao.com/API.htm?docType=2&docId=43570) | [github.com/windawings/opentaobao/api/alihealthoutflow](https://pkg.go.dev/github.com/windawings/opentaobao/api/alihealthoutflow) |
| [x] [飞猪酒店标准库](https://open.taobao.com/API.htm?docType=2&docId=44900) | [github.com/windawings/opentaobao/api/hotelhstdf](https://pkg.go.dev/github.com/windawings/opentaobao/api/hotelhstdf) |
| [x] [AE-Dropshipper](https://open.taobao.com/API.htm?docType=2&docId=57188) | [github.com/windawings/opentaobao/api/aedropshiper](https://pkg.go.dev/github.com/windawings/opentaobao/api/aedropshiper) |
| [x] [银泰scm-openapi](https://open.taobao.com/API.htm?docType=2&docId=41072) | [github.com/windawings/opentaobao/api/moscm](https://pkg.go.dev/github.com/windawings/opentaobao/api/moscm) |
| [x] [信息平台-采购](https://open.taobao.com/API.htm?docType=2&docId=38501) | [github.com/windawings/opentaobao/api/pur](https://pkg.go.dev/github.com/windawings/opentaobao/api/pur) |
| [x] [零售通自动售货机](https://open.taobao.com/API.htm?docType=2&docId=37506) | [github.com/windawings/opentaobao/api/lstvending](https://pkg.go.dev/github.com/windawings/opentaobao/api/lstvending) |
| [x] [淘宝小程序API](https://open.taobao.com/API.htm?docType=2&docId=65356) | [github.com/windawings/opentaobao/api/miniapp](https://pkg.go.dev/github.com/windawings/opentaobao/api/miniapp) |
| [x] [ott支付](https://open.taobao.com/API.htm?docType=2&docId=46647) | [github.com/windawings/opentaobao/api/ottpay](https://pkg.go.dev/github.com/windawings/opentaobao/api/ottpay) |
| [x] [天猫汽车](https://open.taobao.com/API.htm?docType=2&docId=41631) | [github.com/windawings/opentaobao/api/tmallcar](https://pkg.go.dev/github.com/windawings/opentaobao/api/tmallcar) |
| [x] [ALIOS广告平台](https://open.taobao.com/API.htm?docType=2&docId=37669) | [github.com/windawings/opentaobao/api/admarket](https://pkg.go.dev/github.com/windawings/opentaobao/api/admarket) |
| [x] [业务平台新零售](https://open.taobao.com/API.htm?docType=2&docId=41646) | [github.com/windawings/opentaobao/api/uscesl](https://pkg.go.dev/github.com/windawings/opentaobao/api/uscesl) |
| [x] [大麦第三方商家接入API](https://open.taobao.com/API.htm?docType=2&docId=62313) | [github.com/windawings/opentaobao/api/damaiticklet](https://pkg.go.dev/github.com/windawings/opentaobao/api/damaiticklet) |
| [x] [ICBU商品api](https://open.taobao.com/API.htm?docType=2&docId=50558) | [github.com/windawings/opentaobao/api/icbuproduct](https://pkg.go.dev/github.com/windawings/opentaobao/api/icbuproduct) |
| [x] [神鲸应用API](https://open.taobao.com/API.htm?docType=2&docId=41196) | [github.com/windawings/opentaobao/api/shenjing](https://pkg.go.dev/github.com/windawings/opentaobao/api/shenjing) |
| [x] [阿里影业灯塔](https://open.taobao.com/API.htm?docType=2&docId=50713) | [github.com/windawings/opentaobao/api/dengta](https://pkg.go.dev/github.com/windawings/opentaobao/api/dengta) |
| [x] [激励API](https://open.taobao.com/API.htm?docType=2&docId=38249) | [github.com/windawings/opentaobao/api/degoperation](https://pkg.go.dev/github.com/windawings/opentaobao/api/degoperation) |
| [x] [优酷-媒资](https://open.taobao.com/API.htm?docType=2&docId=42057) | [github.com/windawings/opentaobao/api/youkuott](https://pkg.go.dev/github.com/windawings/opentaobao/api/youkuott) |
| [x] [优酷网盟](https://open.taobao.com/API.htm?docType=2&docId=39065) | [github.com/windawings/opentaobao/api/youkudsp](https://pkg.go.dev/github.com/windawings/opentaobao/api/youkudsp) |
| [x] [阿里健康API](https://open.taobao.com/API.htm?docType=2&docId=64037) | [github.com/windawings/opentaobao/api/alihealth2](https://pkg.go.dev/github.com/windawings/opentaobao/api/alihealth2) |
| [x] [新制造API](https://open.taobao.com/API.htm?docType=2&docId=65267) | [github.com/windawings/opentaobao/api/rhino](https://pkg.go.dev/github.com/windawings/opentaobao/api/rhino) |
| [x] [零售通小店智能设备](https://open.taobao.com/API.htm?docType=2&docId=45361) | [github.com/windawings/opentaobao/api/lstspeacker](https://pkg.go.dev/github.com/windawings/opentaobao/api/lstspeacker) |
| [x] [医知鹿-视频](https://open.taobao.com/API.htm?docType=2&docId=46849) | [github.com/windawings/opentaobao/api/alihealthmdeer](https://pkg.go.dev/github.com/windawings/opentaobao/api/alihealthmdeer) |
| [x] [五道口商品API](https://open.taobao.com/API.htm?docType=2&docId=30648) | [github.com/windawings/opentaobao/api/wdkitem](https://pkg.go.dev/github.com/windawings/opentaobao/api/wdkitem) |
| [x] [零售通订单履行](https://open.taobao.com/API.htm?docType=2&docId=53482) | [github.com/windawings/opentaobao/api/lstlogistics](https://pkg.go.dev/github.com/windawings/opentaobao/api/lstlogistics) |
| [x] [天猫线下大屏](https://open.taobao.com/API.htm?docType=2&docId=39240) | [github.com/windawings/opentaobao/api/tmallfcbox](https://pkg.go.dev/github.com/windawings/opentaobao/api/tmallfcbox) |
| [x] [AE-Oversea-Solution](https://open.taobao.com/API.htm?docType=2&docId=47378) | [github.com/windawings/opentaobao/api/aesolution](https://pkg.go.dev/github.com/windawings/opentaobao/api/aesolution) |
| [x] [阿里精灵基础能力接口](https://open.taobao.com/API.htm?docType=2&docId=39668) | [github.com/windawings/opentaobao/api/aligenie](https://pkg.go.dev/github.com/windawings/opentaobao/api/aligenie) |
| [x] [阿里健康新零售](https://open.taobao.com/API.htm?docType=2&docId=41387) | [github.com/windawings/opentaobao/api/healthnr](https://pkg.go.dev/github.com/windawings/opentaobao/api/healthnr) |
| [x] [闲鱼发布](https://open.taobao.com/API.htm?docType=2&docId=39867) | [github.com/windawings/opentaobao/api/idleitem](https://pkg.go.dev/github.com/windawings/opentaobao/api/idleitem) |
| [x] [飞猪酒店签约中心](https://open.taobao.com/API.htm?docType=2&docId=49754) | [github.com/windawings/opentaobao/api/hotelalliance](https://pkg.go.dev/github.com/windawings/opentaobao/api/hotelalliance) |
| [x] [人工智能实验室开放平台API](https://open.taobao.com/API.htm?docType=2&docId=45531) | [github.com/windawings/opentaobao/api/alilabs](https://pkg.go.dev/github.com/windawings/opentaobao/api/alilabs) |
| [x] [天猫新品创新中心API](https://open.taobao.com/API.htm?docType=2&docId=43329) | [github.com/windawings/opentaobao/api/tmic](https://pkg.go.dev/github.com/windawings/opentaobao/api/tmic) |
| [x] [MOZI账号API](https://open.taobao.com/API.htm?docType=2&docId=45946) | [github.com/windawings/opentaobao/api/mozi](https://pkg.go.dev/github.com/windawings/opentaobao/api/mozi) |
| [x] [亲橙里westcrmAPI](https://open.taobao.com/API.htm?docType=2&docId=41221) | [github.com/windawings/opentaobao/api/westcrm](https://pkg.go.dev/github.com/windawings/opentaobao/api/westcrm) |
| [x] [ICBU-橱](https://open.taobao.com/API.htm?docType=2&docId=40964) | [github.com/windawings/opentaobao/api/icbushowcase](https://pkg.go.dev/github.com/windawings/opentaobao/api/icbushowcase) |
| [x] [云码API](https://open.taobao.com/API.htm?docType=2&docId=47156) | [github.com/windawings/opentaobao/api/mc](https://pkg.go.dev/github.com/windawings/opentaobao/api/mc) |
| [x] [HOMEAI](https://open.taobao.com/API.htm?docType=2&docId=41831) | [github.com/windawings/opentaobao/api/homeai](https://pkg.go.dev/github.com/windawings/opentaobao/api/homeai) |
| [x] [MOZI权限API](https://open.taobao.com/API.htm?docType=2&docId=45484) | [github.com/windawings/opentaobao/api/moziacl](https://pkg.go.dev/github.com/windawings/opentaobao/api/moziacl) |
| [x] [飞猪商家平台](https://open.taobao.com/API.htm?docType=2&docId=57069) | [github.com/windawings/opentaobao/api/alitripmerchant](https://pkg.go.dev/github.com/windawings/opentaobao/api/alitripmerchant) |
| [x] [新零售POS](https://open.taobao.com/API.htm?docType=2&docId=42992) | [github.com/windawings/opentaobao/api/nrpos](https://pkg.go.dev/github.com/windawings/opentaobao/api/nrpos) |
| [x] [资质共享API](https://open.taobao.com/API.htm?docType=2&docId=42262) | [github.com/windawings/opentaobao/api/fivee](https://pkg.go.dev/github.com/windawings/opentaobao/api/fivee) |
| [x] [业务平台事业部-税务平台API](https://open.taobao.com/API.htm?docType=2&docId=41740) | [github.com/windawings/opentaobao/api/tax](https://pkg.go.dev/github.com/windawings/opentaobao/api/tax) |
| [x] [MOZI 租户](https://open.taobao.com/API.htm?docType=2&docId=51055) | [github.com/windawings/opentaobao/api/mozivds](https://pkg.go.dev/github.com/windawings/opentaobao/api/mozivds) |
| [x] [公益三小时公共](https://open.taobao.com/API.htm?docType=2&docId=59468) | [github.com/windawings/opentaobao/api/charity](https://pkg.go.dev/github.com/windawings/opentaobao/api/charity) |
| [x] [脚模API](https://open.taobao.com/API.htm?docType=2&docId=50510) | [github.com/windawings/opentaobao/api/foodscan](https://pkg.go.dev/github.com/windawings/opentaobao/api/foodscan) |
| [x] [阿里健康处方药平台](https://open.taobao.com/API.htm?docType=2&docId=44810) | [github.com/windawings/opentaobao/api/alidoc](https://pkg.go.dev/github.com/windawings/opentaobao/api/alidoc) |
| [x] [优酷播控幻影API](https://open.taobao.com/API.htm?docType=2&docId=43082) | [github.com/windawings/opentaobao/api/mirage](https://pkg.go.dev/github.com/windawings/opentaobao/api/mirage) |
| [x] [miniapp开放API](https://open.taobao.com/API.htm?docType=2&docId=51768) | [github.com/windawings/opentaobao/api/miniappopen](https://pkg.go.dev/github.com/windawings/opentaobao/api/miniappopen) |
| [x] [天天特价](https://open.taobao.com/API.htm?docType=2&docId=48644) | [github.com/windawings/opentaobao/api/tttm](https://pkg.go.dev/github.com/windawings/opentaobao/api/tttm) |
| [x] [iHome API](https://open.taobao.com/API.htm?docType=2&docId=43219) | [github.com/windawings/opentaobao/api/ihome](https://pkg.go.dev/github.com/windawings/opentaobao/api/ihome) |
| [x] [信息流API](https://open.taobao.com/API.htm?docType=2&docId=43248) | [github.com/windawings/opentaobao/api/feedflow](https://pkg.go.dev/github.com/windawings/opentaobao/api/feedflow) |
| [x] [淘宝C2M](https://open.taobao.com/API.htm?docType=2&docId=50720) | [github.com/windawings/opentaobao/api/c2m](https://pkg.go.dev/github.com/windawings/opentaobao/api/c2m) |
| [x] [新零售供应链API](https://open.taobao.com/API.htm?docType=2&docId=54020) | [github.com/windawings/opentaobao/api/ascpchannel](https://pkg.go.dev/github.com/windawings/opentaobao/api/ascpchannel) |
| [x] [AliOS支付API](https://open.taobao.com/API.htm?docType=2&docId=61621) | [github.com/windawings/opentaobao/api/aliospay](https://pkg.go.dev/github.com/windawings/opentaobao/api/aliospay) |
| [x] [菜鸟控制塔API](https://open.taobao.com/API.htm?docType=2&docId=43854) | [github.com/windawings/opentaobao/api/cainiaoecc](https://pkg.go.dev/github.com/windawings/opentaobao/api/cainiaoecc) |
| [x] [小程序API](https://open.taobao.com/API.htm?docType=2&docId=54992) | [github.com/windawings/opentaobao/api/yunosminiapp](https://pkg.go.dev/github.com/windawings/opentaobao/api/yunosminiapp) |
| [x] [菜鸟末端商业](https://open.taobao.com/API.htm?docType=2&docId=68008) | [github.com/windawings/opentaobao/api/cainiaocntec](https://pkg.go.dev/github.com/windawings/opentaobao/api/cainiaocntec) |
| [x] [阿里健康挂号号源直连](https://open.taobao.com/API.htm?docType=2&docId=55216) | [github.com/windawings/opentaobao/api/medicalbase](https://pkg.go.dev/github.com/windawings/opentaobao/api/medicalbase) |
| [x] [AE-UserOpen-Recommend](https://open.taobao.com/API.htm?docType=2&docId=45722) | [github.com/windawings/opentaobao/api/aeusergrowth](https://pkg.go.dev/github.com/windawings/opentaobao/api/aeusergrowth) |
| [x] [Efficient Tools](https://open.taobao.com/API.htm?docType=2&docId=45804) | [github.com/windawings/opentaobao/api/aetools](https://pkg.go.dev/github.com/windawings/opentaobao/api/aetools) |
| [x] [Data Reports](https://open.taobao.com/API.htm?docType=2&docId=45807) | [github.com/windawings/opentaobao/api/aedata](https://pkg.go.dev/github.com/windawings/opentaobao/api/aedata) |
| [x] [Promotion Creatives](https://open.taobao.com/API.htm?docType=2&docId=45801) | [github.com/windawings/opentaobao/api/aecreatives](https://pkg.go.dev/github.com/windawings/opentaobao/api/aecreatives) |
| [x] [菜鸟国际出口](https://open.taobao.com/API.htm?docType=2&docId=62511) | [github.com/windawings/opentaobao/api/cainiaohandover](https://pkg.go.dev/github.com/windawings/opentaobao/api/cainiaohandover) |
| [x] [全域会员通](https://open.taobao.com/API.htm?docType=2&docId=46937) | [github.com/windawings/opentaobao/api/alimember](https://pkg.go.dev/github.com/windawings/opentaobao/api/alimember) |
| [x] [交易猫API](https://open.taobao.com/API.htm?docType=2&docId=60445) | [github.com/windawings/opentaobao/api/jym](https://pkg.go.dev/github.com/windawings/opentaobao/api/jym) |
| [x] [影城自运营开放Api](https://open.taobao.com/API.htm?docType=2&docId=46972) | [github.com/windawings/opentaobao/api/filmtfavatar](https://pkg.go.dev/github.com/windawings/opentaobao/api/filmtfavatar) |
| [x] [商家应用](https://open.taobao.com/API.htm?docType=2&docId=46995) | [github.com/windawings/opentaobao/api/miniappcloud](https://pkg.go.dev/github.com/windawings/opentaobao/api/miniappcloud) |
| [x] [国际火车票API](https://open.taobao.com/API.htm?docType=2&docId=47364) | [github.com/windawings/opentaobao/api/rail](https://pkg.go.dev/github.com/windawings/opentaobao/api/rail) |
| [x] [天猫校园API](https://open.taobao.com/API.htm?docType=2&docId=61316) | [github.com/windawings/opentaobao/api/tmallcampus](https://pkg.go.dev/github.com/windawings/opentaobao/api/tmallcampus) |
| [x] [消息API](https://open.taobao.com/API.htm?docType=2&docId=56166) | [github.com/windawings/opentaobao/api/alimsg](https://pkg.go.dev/github.com/windawings/opentaobao/api/alimsg) |
| [x] [菜鸟发货工作台API](https://open.taobao.com/API.htm?docType=2&docId=47727) | [github.com/windawings/opentaobao/api/consignplatform](https://pkg.go.dev/github.com/windawings/opentaobao/api/consignplatform) |
| [x] [五道口-物流-自提API](https://open.taobao.com/API.htm?docType=2&docId=47575) | [github.com/windawings/opentaobao/api/wdklogistics](https://pkg.go.dev/github.com/windawings/opentaobao/api/wdklogistics) |
| [x] [企业订餐员工API](https://open.taobao.com/API.htm?docType=2&docId=48679) | [github.com/windawings/opentaobao/api/eleenterpriseemployee](https://pkg.go.dev/github.com/windawings/opentaobao/api/eleenterpriseemployee) |
| [x] [船票API](https://open.taobao.com/API.htm?docType=2&docId=48055) | [github.com/windawings/opentaobao/api/ship](https://pkg.go.dev/github.com/windawings/opentaobao/api/ship) |
| [x] [飞猪-综合交通api](https://open.taobao.com/API.htm?docType=2&docId=53101) | [github.com/windawings/opentaobao/api/alitripcar](https://pkg.go.dev/github.com/windawings/opentaobao/api/alitripcar) |
| [x] [企业订餐店铺接口](https://open.taobao.com/API.htm?docType=2&docId=49167) | [github.com/windawings/opentaobao/api/eleenterpriserestaurant](https://pkg.go.dev/github.com/windawings/opentaobao/api/eleenterpriserestaurant) |
| [x] [新零售开放API](https://open.taobao.com/API.htm?docType=2&docId=50091) | [github.com/windawings/opentaobao/api/nropen](https://pkg.go.dev/github.com/windawings/opentaobao/api/nropen) |
| [x] [企业订餐购物车API](https://open.taobao.com/API.htm?docType=2&docId=49011) | [github.com/windawings/opentaobao/api/eleenterprisecartnew](https://pkg.go.dev/github.com/windawings/opentaobao/api/eleenterprisecartnew) |
| [x] [企业订餐优惠券API](https://open.taobao.com/API.htm?docType=2&docId=49008) | [github.com/windawings/opentaobao/api/eleenterprisecoupon](https://pkg.go.dev/github.com/windawings/opentaobao/api/eleenterprisecoupon) |
| [x] [阿里健康算法](https://open.taobao.com/API.htm?docType=2&docId=48492) | [github.com/windawings/opentaobao/api/alihealthalgo](https://pkg.go.dev/github.com/windawings/opentaobao/api/alihealthalgo) |
| [x] [企业订餐订单API](https://open.taobao.com/API.htm?docType=2&docId=49014) | [github.com/windawings/opentaobao/api/eleenterpriseordernew](https://pkg.go.dev/github.com/windawings/opentaobao/api/eleenterpriseordernew) |
| [x] [云游戏API](https://open.taobao.com/API.htm?docType=2&docId=55012) | [github.com/windawings/opentaobao/api/cloudgame](https://pkg.go.dev/github.com/windawings/opentaobao/api/cloudgame) |
| [x] [国际虚拟API](https://open.taobao.com/API.htm?docType=2&docId=48538) | [github.com/windawings/opentaobao/api/globalvirtual](https://pkg.go.dev/github.com/windawings/opentaobao/api/globalvirtual) |
| [x] [交易开放](https://open.taobao.com/API.htm?docType=2&docId=63776) | [github.com/windawings/opentaobao/api/opentrade](https://pkg.go.dev/github.com/windawings/opentaobao/api/opentrade) |
| [x] [本地生活商户基础API](https://open.taobao.com/API.htm?docType=2&docId=64157) | [github.com/windawings/opentaobao/api/alscmerchant](https://pkg.go.dev/github.com/windawings/opentaobao/api/alscmerchant) |
| [x] [消息amp通道](https://open.taobao.com/API.htm?docType=2&docId=62215) | [github.com/windawings/opentaobao/api/msgamp](https://pkg.go.dev/github.com/windawings/opentaobao/api/msgamp) |
| [x] [口碑综合体API](https://open.taobao.com/API.htm?docType=2&docId=49834) | [github.com/windawings/opentaobao/api/koubeimall](https://pkg.go.dev/github.com/windawings/opentaobao/api/koubeimall) |
| [x] [IoT售后解决方案API](https://open.taobao.com/API.htm?docType=2&docId=50245) | [github.com/windawings/opentaobao/api/iotticket](https://pkg.go.dev/github.com/windawings/opentaobao/api/iotticket) |
| [x] [OpenMall-API](https://open.taobao.com/API.htm?docType=2&docId=50836) | [github.com/windawings/opentaobao/api/openmall](https://pkg.go.dev/github.com/windawings/opentaobao/api/openmall) |
| [x] [视觉开放API(viapi)](https://open.taobao.com/API.htm?docType=2&docId=50726) | [github.com/windawings/opentaobao/api/viapi](https://pkg.go.dev/github.com/windawings/opentaobao/api/viapi) |
| [x] [曲库开放平台歌曲API](https://open.taobao.com/API.htm?docType=2&docId=61549) | [github.com/windawings/opentaobao/api/xiamicontent](https://pkg.go.dev/github.com/windawings/opentaobao/api/xiamicontent) |
| [x] [阿里健康三方机构](https://open.taobao.com/API.htm?docType=2&docId=51290) | [github.com/windawings/opentaobao/api/alihealthmedical](https://pkg.go.dev/github.com/windawings/opentaobao/api/alihealthmedical) |
| [x] [阿里健康-检测检验-预约](https://open.taobao.com/API.htm?docType=2&docId=51444) | [github.com/windawings/opentaobao/api/alihealthlab](https://pkg.go.dev/github.com/windawings/opentaobao/api/alihealthlab) |
| [x] [天猫精灵供应链网关API](https://open.taobao.com/API.htm?docType=2&docId=53192) | [github.com/windawings/opentaobao/api/tmallgeniescp](https://pkg.go.dev/github.com/windawings/opentaobao/api/tmallgeniescp) |
| [x] [曲库开放平台TraceAPI](https://open.taobao.com/API.htm?docType=2&docId=51350) | [github.com/windawings/opentaobao/api/xiamitrace](https://pkg.go.dev/github.com/windawings/opentaobao/api/xiamitrace) |
| [x] [飞猪发票](https://open.taobao.com/API.htm?docType=2&docId=60461) | [github.com/windawings/opentaobao/api/alitripreceipt](https://pkg.go.dev/github.com/windawings/opentaobao/api/alitripreceipt) |
| [x] [互动开放API](https://open.taobao.com/API.htm?docType=2&docId=53888) | [github.com/windawings/opentaobao/api/jstinteractive](https://pkg.go.dev/github.com/windawings/opentaobao/api/jstinteractive) |
| [x] [同城零售全渠道](https://open.taobao.com/API.htm?docType=2&docId=53988) | [github.com/windawings/opentaobao/api/perfect](https://pkg.go.dev/github.com/windawings/opentaobao/api/perfect) |
| [x] [天猫好房工具API](https://open.taobao.com/API.htm?docType=2&docId=61783) | [github.com/windawings/opentaobao/api/alihouse](https://pkg.go.dev/github.com/windawings/opentaobao/api/alihouse) |
| [x] [五棱镜任务API](https://open.taobao.com/API.htm?docType=2&docId=52767) | [github.com/windawings/opentaobao/api/pentraprism](https://pkg.go.dev/github.com/windawings/opentaobao/api/pentraprism) |
| [x] [闲鱼兼职](https://open.taobao.com/API.htm?docType=2&docId=52622) | [github.com/windawings/opentaobao/api/idleparttime](https://pkg.go.dev/github.com/windawings/opentaobao/api/idleparttime) |
| [x] [聚石塔隐私号](https://open.taobao.com/API.htm?docType=2&docId=52636) | [github.com/windawings/opentaobao/api/jstsecret](https://pkg.go.dev/github.com/windawings/opentaobao/api/jstsecret) |
| [x] [零售通商品API](https://open.taobao.com/API.htm?docType=2&docId=47318) | [github.com/windawings/opentaobao/api/lsticitem](https://pkg.go.dev/github.com/windawings/opentaobao/api/lsticitem) |
| [x] [零售通交易API](https://open.taobao.com/API.htm?docType=2&docId=43690) | [github.com/windawings/opentaobao/api/lsttrade](https://pkg.go.dev/github.com/windawings/opentaobao/api/lsttrade) |
| [x] [零售通门店API](https://open.taobao.com/API.htm?docType=2&docId=43687) | [github.com/windawings/opentaobao/api/lstbm](https://pkg.go.dev/github.com/windawings/opentaobao/api/lstbm) |
| [x] [零售通履单API](https://open.taobao.com/API.htm?docType=2&docId=51631) | [github.com/windawings/opentaobao/api/lstlogistics2](https://pkg.go.dev/github.com/windawings/opentaobao/api/lstlogistics2) |
| [x] [零售通仓储API](https://open.taobao.com/API.htm?docType=2&docId=53981) | [github.com/windawings/opentaobao/api/lstwarehouse](https://pkg.go.dev/github.com/windawings/opentaobao/api/lstwarehouse) |
| [x] [零售通结算API](https://open.taobao.com/API.htm?docType=2&docId=46608) | [github.com/windawings/opentaobao/api/lstfundbill](https://pkg.go.dev/github.com/windawings/opentaobao/api/lstfundbill) |
| [x] [零售通营销API](https://open.taobao.com/API.htm?docType=2&docId=44185) | [github.com/windawings/opentaobao/api/lstmarketing](https://pkg.go.dev/github.com/windawings/opentaobao/api/lstmarketing) |
| [x] [AE任务开放平台](https://open.taobao.com/API.htm?docType=2&docId=53415) | [github.com/windawings/opentaobao/api/aetask](https://pkg.go.dev/github.com/windawings/opentaobao/api/aetask) |
| [x] [本地生活内容API](https://open.taobao.com/API.htm?docType=2&docId=54328) | [github.com/windawings/opentaobao/api/kbalgo](https://pkg.go.dev/github.com/windawings/opentaobao/api/kbalgo) |
| [x] [天猫新品平台API](https://open.taobao.com/API.htm?docType=2&docId=54824) | [github.com/windawings/opentaobao/api/tmalltrend](https://pkg.go.dev/github.com/windawings/opentaobao/api/tmalltrend) |
| [x] [国际化中台服务域保险](https://open.taobao.com/API.htm?docType=2&docId=53335) | [github.com/windawings/opentaobao/api/middleclaims](https://pkg.go.dev/github.com/windawings/opentaobao/api/middleclaims) |
| [x] [阿里健康-健康证](https://open.taobao.com/API.htm?docType=2&docId=54140) | [github.com/windawings/opentaobao/api/alihealthcert](https://pkg.go.dev/github.com/windawings/opentaobao/api/alihealthcert) |
| [x] [隐私保护API](https://open.taobao.com/API.htm?docType=2&docId=57628) | [github.com/windawings/opentaobao/api/topoaid](https://pkg.go.dev/github.com/windawings/opentaobao/api/topoaid) |
| [x] [ICBU-DropShipping](https://open.taobao.com/API.htm?docType=2&docId=54908) | [github.com/windawings/opentaobao/api/icbudropshipping](https://pkg.go.dev/github.com/windawings/opentaobao/api/icbudropshipping) |
| [x] [曲库开放平台艺人API](https://open.taobao.com/API.htm?docType=2&docId=55125) | [github.com/windawings/opentaobao/api/xiamiatrist](https://pkg.go.dev/github.com/windawings/opentaobao/api/xiamiatrist) |
| [x] [周期购API](https://open.taobao.com/API.htm?docType=2&docId=55500) | [github.com/windawings/opentaobao/api/zqs](https://pkg.go.dev/github.com/windawings/opentaobao/api/zqs) |
| [x] [闲鱼已验货](https://open.taobao.com/API.htm?docType=2&docId=49487) | [github.com/windawings/opentaobao/api/idleisv](https://pkg.go.dev/github.com/windawings/opentaobao/api/idleisv) |
| [x] [阿信-交易API](https://open.taobao.com/API.htm?docType=2&docId=55695) | [github.com/windawings/opentaobao/api/axintrade](https://pkg.go.dev/github.com/windawings/opentaobao/api/axintrade) |
| [x] [飞猪商业化API](https://open.taobao.com/API.htm?docType=2&docId=58555) | [github.com/windawings/opentaobao/api/alitripbp](https://pkg.go.dev/github.com/windawings/opentaobao/api/alitripbp) |
| [x] [阿信-基础数据](https://open.taobao.com/API.htm?docType=2&docId=56861) | [github.com/windawings/opentaobao/api/axindata](https://pkg.go.dev/github.com/windawings/opentaobao/api/axindata) |
| [x] [阿里健康公益线API](https://open.taobao.com/API.htm?docType=2&docId=56009) | [github.com/windawings/opentaobao/api/alihealthpw](https://pkg.go.dev/github.com/windawings/opentaobao/api/alihealthpw) |
| [x] [海南离岛对外API](https://open.taobao.com/API.htm?docType=2&docId=56257) | [github.com/windawings/opentaobao/api/dutyfree](https://pkg.go.dev/github.com/windawings/opentaobao/api/dutyfree) |
| [x] [淘宝交易API,](https://open.taobao.com/API.htm?docType=2&docId=62880) | [github.com/windawings/opentaobao/api/tbtrade](https://pkg.go.dev/github.com/windawings/opentaobao/api/tbtrade) |
| [x] [淘宝用户API,](https://open.taobao.com/API.htm?docType=2&docId=26303) | [github.com/windawings/opentaobao/api/tbuser](https://pkg.go.dev/github.com/windawings/opentaobao/api/tbuser) |
| [x] [淘宝商品API,](https://open.taobao.com/API.htm?docType=2&docId=54542) | [github.com/windawings/opentaobao/api/tbitem](https://pkg.go.dev/github.com/windawings/opentaobao/api/tbitem) |
| [x] [淘宝物流API,](https://open.taobao.com/API.htm?docType=2&docId=65697) | [github.com/windawings/opentaobao/api/tblogistics](https://pkg.go.dev/github.com/windawings/opentaobao/api/tblogistics) |
| [x] [淘宝退款API,](https://open.taobao.com/API.htm?docType=2&docId=64018) | [github.com/windawings/opentaobao/api/tbrefund](https://pkg.go.dev/github.com/windawings/opentaobao/api/tbrefund) |

