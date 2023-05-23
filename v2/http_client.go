package binance

import "net/http"

//func (c *Client) Do(req *Request) (*Response, error) {
//	return c.do(req)
//}


type IHTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
