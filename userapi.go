package moodleClient

import (
	"context"
	"fmt"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
	"github.com/mnasruul/moodleClient/pkg/urlutil"
)

type UserAPI interface {
	CreateUsers(ctx context.Context, param string) ([]userResponse, error)
}

type userAPI struct {
	*apiClient
}

func newUserAPI(apiClient *apiClient) *userAPI {
	return &userAPI{apiClient}
}

type userResponse struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
}

type UserRequest struct {
	Createpassword    int         `json:"createpassword"`
	Username          string      `json:"username"`
	Auth              string      `json:"auth"`
	Password          string      `json:"password"`
	Firstname         string      `json:"firstname"`
	Lastname          string      `json:"lastname"`
	Email             string      `json:"email"`
	Maildisplay       int         `json:"maildisplay"`
	City              string      `json:"city"`
	Country           string      `json:"country"`
	Timezone          string      `json:"timezone"`
	Description       string      `json:"description"`
	Firstnamephonetic string      `json:"firstnamephonetic"`
	Lastnamephonetic  string      `json:"lastnamephonetic"`
	Middlename        string      `json:"middlename"`
	Alternatename     string      `json:"alternatename"`
	Interests         string      `json:"interests"`
	Idnumber          string      `json:"idnumber"`
	Institution       string      `json:"institution"`
	Department        string      `json:"department"`
	Phone1            string      `json:"phone1"`
	Phone2            string      `json:"phone2"`
	Address           string      `json:"address"`
	Lang              string      `json:"lang"`
	Calendartype      string      `json:"calendartype"`
	Theme             string      `json:"theme"`
	Mailformat        int         `json:"mailformat"`
	Customfields      interface{} `json:"customfields"`
}
type UserRequests []*UserRequest

type Options struct {
	Users UserRequests `url:"users"`
}

func (us UserRequests) EncodeValues(key string, v *url.Values) error {

	for i, u := range us {
		res, err := query.Values(u)
		if err != nil {
			return err
		}
		for subKey, subVal := range res {
			_ = subVal
			val := res.Get(subKey)
			field, _ := reflect.TypeOf(u).Elem().FieldByName(subKey)
			if val != "<nil>" {
				v.Set(fmt.Sprintf("%s[%d][%s]", key, i, urlutil.GetStructTag(field, "json")), res.Get(subKey))
			}
		}
	}
	return nil
}

func (u *userAPI) CreateUsers(ctx context.Context, param string) ([]userResponse, error) {
	res := []userResponse{}
	err := u.callMoodleFunctionPost(ctx, &res, param, map[string]string{
		"wsfunction": "core_user_create_users",
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
