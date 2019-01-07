package ds

import (
	www "net/http"
	//
	"golang.org/x/net/context"
	"google.golang.org/appengine"
//	datastore "cloud.google.com/go/datastore"
	datastoreAE "google.golang.org/appengine/datastore"
	//
	"github.com/jsonrouter/core/http"
)

func (client *Client) RunInTransaction(req http.Request, f func (context.Context) error) error {

	ctx := appengine.NewContext(req.R().(*www.Request))

	if client.appEngine {
		return datastoreAE.RunInTransaction(
			ctx,
			f,
			nil,
		)
	}

	return f(ctx)
}
