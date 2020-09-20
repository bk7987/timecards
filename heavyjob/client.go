package heavyjob

import (
	"net/http"

	"github.com/bk7987/timecards/config"
	"github.com/gofiber/fiber/v2"
	"github.com/imroc/req"
)

// Client represents the Heavyjob API client.
type Client struct {
	RootURL string
	Token   string
	Request *req.Req
}

const clientLocal string = "heavyjob_client"

// newClient returns a configured HttpClient for making requests to the HeavyJob API.
func newClient() *Client {
	token := GetIdentity().AuthToken
	request := req.New()

	return &Client{
		RootURL: config.HeavyjobRootURL(),
		Token:   token,
		Request: request,
	}
}

// SetClient returns a middleware that creates an http client and stores it in the fiber context.
func SetClient() func(*fiber.Ctx) {
	client := newClient()
	return func(ctx *fiber.Ctx) {
		ctx.Locals(clientLocal, client)
		ctx.Next()
	}
}

// GetClient returns the Heavyjob Client from the current context
func GetClient(ctx *fiber.Ctx) *Client {
	return ctx.Locals(clientLocal).(*Client)
}

func (c *Client) get(path string, v interface{}) (*http.Response, error) {
	u := c.RootURL + path
	headers := req.Header{
		"Accept":        "application/json",
		"Authorization": "Bearer " + c.Token,
	}

	res, err := c.Request.Get(u, headers)
	res.ToJSON(&v)
	return res.Response(), err
}
