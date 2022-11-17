package moodleClient

import "context"

type UserAPI interface {
	CreateUsers(ctx context.Context, param []UserRequest) ([]userResponse, error)
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
	Password          string      `json:"passowrd"`
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

func (u *userAPI) CreateUsers(ctx context.Context, param []UserRequest) ([]userResponse, error) {
	res := []userResponse{}
	err := u.callMoodleFunctionPost(ctx, &res, param, map[string]string{
		"wsfunction": "core_user_create_users",
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
