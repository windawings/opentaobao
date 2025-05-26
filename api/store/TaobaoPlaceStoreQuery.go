package store

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/store"
)

// TaobaoPlaceStoreQuery 门店信息查询接口
// taobao.place.store.query
//
// 根据用户授权信息，获取用户的门店公开信息
func TaobaoPlaceStoreQuery(ctx context.Context, clt *core.SDKClient, req *store.TaobaoPlaceStoreQueryAPIRequest, resp *store.TaobaoPlaceStoreQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
