package simba

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/simba"
)

// TaobaoSubwayKeywordOfflineLayeredfind 查询关键词离线报表30天转化周期
// taobao.subway.keyword.offline.layeredfind
//
// 获取关键词离线报表
func TaobaoSubwayKeywordOfflineLayeredfind(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSubwayKeywordOfflineLayeredfindAPIRequest, resp *simba.TaobaoSubwayKeywordOfflineLayeredfindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
