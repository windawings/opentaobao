package singletreasure

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/singletreasure"
)

// TaobaoSingletreasureActivityQuery 查询活动列表接口
// taobao.singletreasure.activity.query
//
// 查询活动列表接口
func TaobaoSingletreasureActivityQuery(ctx context.Context, clt *core.SDKClient, req *singletreasure.TaobaoSingletreasureActivityQueryAPIRequest, resp *singletreasure.TaobaoSingletreasureActivityQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
