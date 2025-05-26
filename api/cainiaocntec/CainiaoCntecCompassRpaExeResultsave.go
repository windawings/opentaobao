package cainiaocntec

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/cainiaocntec"
)

// CainiaoCntecCompassRpaExeResultsave rpa执行结果回传
// cainiao.cntec.compass.rpa.exe.resultsave
//
// rpa执行结果回传
func CainiaoCntecCompassRpaExeResultsave(ctx context.Context, clt *core.SDKClient, req *cainiaocntec.CainiaoCntecCompassRpaExeResultsaveAPIRequest, resp *cainiaocntec.CainiaoCntecCompassRpaExeResultsaveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
