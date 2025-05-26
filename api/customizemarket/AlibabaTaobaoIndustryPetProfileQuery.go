package customizemarket

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/customizemarket"
)

// AlibabaTaobaoIndustryPetProfileQuery 用户宠物列表查询
// alibaba.taobao.industry.pet.profile.query
//
// 用户宠物列表查询
func AlibabaTaobaoIndustryPetProfileQuery(ctx context.Context, clt *core.SDKClient, req *customizemarket.AlibabaTaobaoIndustryPetProfileQueryAPIRequest, resp *customizemarket.AlibabaTaobaoIndustryPetProfileQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
