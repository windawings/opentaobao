package alihouse

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeShopconfigAstoreSync 天猫好房店铺装修-Astore上翻
// alibaba.alihouse.newhome.shopconfig.astore.sync
//
// 天猫好房店铺装修-Astore上翻
func AlibabaAlihouseNewhomeShopconfigAstoreSync(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
