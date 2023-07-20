package twitter

import (
	"fmt"
	"net/http"
)

// Authorizer will add the authorization to the HTTP request
type Authorizer interface {
	Add(req *http.Request)
}

// AuthorizeReq is the global request token struct
type AuthorizeReq struct {
	Token string
}

// Add adds the bearer token to the request
func (a AuthorizeReq) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}
