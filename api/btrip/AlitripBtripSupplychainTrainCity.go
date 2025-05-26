package btrip

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/btrip"
)

// AlitripBtripSupplychainTrainCity 火车站数据查询
// alitrip.btrip.supplychain.train.city
//
// 火车站数据查询
func AlitripBtripSupplychainTrainCity(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripSupplychainTrainCityAPIRequest, resp *btrip.AlitripBtripSupplychainTrainCityAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
