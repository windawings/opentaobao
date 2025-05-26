package tmallcar

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmallcar"
)

// TmallCarXcarSynchronizeCarModelData 爱车车型数据同步
// tmall.car.xcar.synchronize.car.model.data
//
// 爱车汽车车型数据同步到天猫
func TmallCarXcarSynchronizeCarModelData(ctx context.Context, clt *core.SDKClient, req *tmallcar.TmallCarXcarSynchronizeCarModelDataAPIRequest, resp *tmallcar.TmallCarXcarSynchronizeCarModelDataAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
