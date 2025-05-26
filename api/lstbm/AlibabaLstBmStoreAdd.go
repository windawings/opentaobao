package lstbm

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/lstbm"
)

// AlibabaLstBmStoreAdd 导入品牌商自有门店
// alibaba.lst.bm.store.add
//
// 导入品牌商自有门店
func AlibabaLstBmStoreAdd(ctx context.Context, clt *core.SDKClient, req *lstbm.AlibabaLstBmStoreAddAPIRequest, resp *lstbm.AlibabaLstBmStoreAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
