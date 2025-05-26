package util

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/util"
)

// AlibabaRetailShorturlGet 短链接获取
// alibaba.retail.shorturl.get
//
// 短链接获取
func AlibabaRetailShorturlGet(ctx context.Context, clt *core.SDKClient, req *util.AlibabaRetailShorturlGetAPIRequest, resp *util.AlibabaRetailShorturlGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
