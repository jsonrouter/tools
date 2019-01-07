package ds

import (
	www "net/http"
	//
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	datastore "cloud.google.com/go/datastore"
	datastoreAE "google.golang.org/appengine/datastore"
	//
	"github.com/jsonrouter/core/http"
)

func (client *Client) RunKeysQuery(req http.Request, query *datastore.Query) ([]*datastore.Key, error) {

	keys, err := client.GetAll(context.Background(), query, nil)
	if err != nil {
		return nil, err
	}
	return keys, nil
}

func (client *Client) RunKeysQueryAE(req http.Request, query *datastoreAE.Query) ([]*datastoreAE.Key, error) {

	ctx := appengine.NewContext(req.R().(*www.Request))
	keys, err := query.GetAll(ctx, nil)
	if err != nil {
		return nil, err
	}
	return keys, nil
}

func (client *Client) RunQuery(req http.Request, q, dst interface{}) error {

	_, ok := req.(*http.MockRequest)
	if ok {
		return nil
	}

	switch query := q.(type) {

		case *datastore.Query:

			_, err := client.GetAll(context.Background(), query, dst)
			if err != nil {
				return err
			}

		case *datastoreAE.Query:

			ctx := appengine.NewContext(req.R().(*www.Request))
			_, err := query.GetAll(ctx, dst)
			if err != nil {
				return err
			}

	}

	return nil
}
