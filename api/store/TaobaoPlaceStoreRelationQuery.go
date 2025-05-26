package store

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/store"
)

// TaobaoPlaceStoreRelationQuery 门店关系查询
// taobao.place.store.relation.query
//
// 查询门店关系
func TaobaoPlaceStoreRelationQuery(ctx context.Context, clt *core.SDKClient, req *store.TaobaoPlaceStoreRelationQueryAPIRequest, resp *store.TaobaoPlaceStoreRelationQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
