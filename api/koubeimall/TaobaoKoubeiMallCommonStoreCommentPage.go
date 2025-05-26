package koubeimall

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/koubeimall"
)

// TaobaoKoubeiMallCommonStoreCommentPage 分页查询门店评论详情信息
// taobao.koubei.mall.common.store.comment.page
//
// 查询口碑综合体内的门店评论信息
func TaobaoKoubeiMallCommonStoreCommentPage(ctx context.Context, clt *core.SDKClient, req *koubeimall.TaobaoKoubeiMallCommonStoreCommentPageAPIRequest, resp *koubeimall.TaobaoKoubeiMallCommonStoreCommentPageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
