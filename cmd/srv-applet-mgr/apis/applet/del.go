package applet

import (
	"context"

	"github.com/iotexproject/Bumblebee/kit/httptransport/httpx"

	"github.com/iotexproject/w3bstream/cmd/srv-applet-mgr/apis/middleware"

	"github.com/iotexproject/w3bstream/pkg/modules/applet"
)

type RemoveApplet struct {
	httpx.MethodDelete
	applet.RemoveAppletReq
}

func (r *RemoveApplet) Path() string { return "/:projectID" }

func (r *RemoveApplet) Output(ctx context.Context) (interface{}, error) {
	a := middleware.CurrentAccountFromContext(ctx)
	if _, err := a.ValidateProjectPerm(ctx, r.ProjectID); err != nil {
		return nil, err
	}

	return nil, applet.RemoveApplet(ctx, &r.RemoveAppletReq)
}
