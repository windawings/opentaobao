package {{ .Pkg }}

import (
    "context"

    "github.com/windawings/opentaobao/core"
    "github.com/windawings/opentaobao/model/{{ .Pkg }}"
)

// {{ .Name }} {{ .ChineseName }} 
// {{ .ApiName }}
//
{{ html .Desc }}
func {{ .Name }}(ctx context.Context, clt *core.SDKClient, req *{{ .Pkg }}.{{ .Name }}APIRequest, resp *{{ .Pkg }}.{{ .Name }}APIResponse, session string) error {
    return clt.Post(ctx, req, resp, session)
}
