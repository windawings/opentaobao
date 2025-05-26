package alihouse

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghouseHouseBase （租房）同步房屋信息
// alibaba.alihouse.existinghouse.house.base
//
// 房屋信息上翻
func AlibabaAlihouseExistinghouseHouseBase(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghouseHouseBaseAPIRequest, resp *alihouse.AlibabaAlihouseExistinghouseHouseBaseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
