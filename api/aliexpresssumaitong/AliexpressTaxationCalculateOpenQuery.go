package aliexpresssumaitong

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/aliexpresssumaitong"
)

// AliexpressTaxationCalculateOpenQuery 关务所需的申报清关字段
// aliexpress.taxation.calculate.open.query
//
// 关务所需的申报清关字段
func AliexpressTaxationCalculateOpenQuery(ctx context.Context, clt *core.SDKClient, req *aliexpresssumaitong.AliexpressTaxationCalculateOpenQueryAPIRequest, resp *aliexpresssumaitong.AliexpressTaxationCalculateOpenQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
