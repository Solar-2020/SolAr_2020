package group

import (
	"encoding/json"
	"errors"
	"github.com/valyala/fasthttp"
	"net/http"
	"strconv"
)

type Client interface {
	GetUserRole(userID, groupID int) (roleID int, err error)
	CheckPermission(userID, groupId, actionID int) (err error)
}

type client struct {
	host   string
	secret string
}

func NewClient(host string, secret string) Client {
	return &client{host: host, secret: secret}
}

type httpError struct {
	Error string `json:"error"`
}

type UserRole struct {
	UserID   int    `json:"userID"`
	GroupID  int    `json:"groupID"`
	RoleID   int    `json:"roleID"`
	RoleName string `json:"roleName"`
}

func (c *client) GetUserRole(userID, groupID int) (roleID int, err error) {

	req, err := http.NewRequest(http.MethodGet, c.host+"/api/internal/group/permission", nil)
	if err != nil {
		return
	}
	q := req.URL.Query()
	q.Add("user_id", strconv.Itoa(userID))
	q.Add("group_id", strconv.Itoa(groupID))
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Authorization", c.secret)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var response UserRole
		err = json.NewDecoder(resp.Body).Decode(&response)
		return response.RoleID, err
	case http.StatusBadRequest:
		var httpErr httpError
		err = json.NewDecoder(resp.Body).Decode(&httpErr)
		if err != nil {
			return
		}
		return roleID, errors.New(httpErr.Error)
	default:
		return roleID, errors.New("Unexpected Server Error")
	}
}



func (c *client) CheckPermission(userID, groupId, actionID int) (err error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.URI().SetScheme("http")
	req.URI().SetHost(c.host)
	req.URI().SetPath("api/internal/group/check-permission")

	req.Header.Set("Authorization", c.secret)
	req.Header.SetMethod(fasthttp.MethodPost)

	body, err := json.Marshal(createRequest)
	if err != nil {
		return
	}

	req.SetBody(body)

	err = fasthttp.Do(req, resp)
	if err != nil {
		return
	}
}
