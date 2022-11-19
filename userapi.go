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
	CreateUsers(ctx context.Context, param UserOptions) ([]userResponse, error)
	UpdateUsers(ctx context.Context, param UserOptions) (*updateUserResponse, error)
	ManualEnrolUsers(ctx context.Context, param EnrolMentOptions) error
}

type userAPI struct {
	*apiClient
}

func newUserAPI(apiClient *apiClient) *userAPI {
	return &userAPI{apiClient}
}

type UserRequest struct {
	Id             int64  `json:"id"`
	Createpassword int    `json:"createpassword"`
	Username       string `json:"username"`
	Auth           string `json:"auth"`
	Password       string `json:"password"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	Email          string `json:"email"`
	// Maildisplay       int         `json:"maildisplay"`
	// City              string      `json:"city"`
	// Country           string      `json:"country"`
	// Timezone          string      `json:"timezone"`
	// Description       string      `json:"description"`
	// Firstnamephonetic string      `json:"firstnamephonetic"`
	// Lastnamephonetic  string      `json:"lastnamephonetic"`
	// Middlename        string      `json:"middlename"`
	// Alternatename     string      `json:"alternatename"`
	// Interests         string      `json:"interests"`
	// Idnumber          string      `json:"idnumber"`
	// Institution       string      `json:"institution"`
	// Department        string      `json:"department"`
	// Phone1            string      `json:"phone1"`
	// Phone2            string      `json:"phone2"`
	// Address           string      `json:"address"`
	// Lang              string      `json:"lang"`
	// Calendartype      string      `json:"calendartype"`
	// Theme             string      `json:"theme"`
	// Mailformat        int         `json:"mailformat"`
	// Customfields      interface{} `json:"customfields"`
}
type UserRequests []*UserRequest

type UserOptions struct {
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
			if val != "<nil>" && val != "" && val != "0" {
				v.Set(fmt.Sprintf("%s[%d][%s]", key, i, urlutil.GetStructTag(field, "json")), res.Get(subKey))
			}
		}
	}
	return nil
}

// param is Encoded of type UserRequests []*UserRequest
//
//	param := []*moodleClient.UserRequest{
//		{
//			Username:  "xxxx",
//			Firstname: "xxx",
//			Lastname:  "xxx",
//			Email:     "xxx@gmail.com",
//			Password:  "@123xxxx",
//			Auth:      "manual",
//		},
//	}
//
// resutl, err := query.Values(moodleClient.Options{Users: param})
// res, err := Client.UserAPI.CreateUsers(ctx, resutl.Encode())
func (u *userAPI) CreateUsers(ctx context.Context, param UserOptions) ([]userResponse, error) {
	qparam, err := query.Values(param)
	if err != nil {
		return nil, err
	}
	res := []userResponse{}
	err = u.callMoodleFunctionPost(ctx, &res, qparam.Encode(), map[string]string{
		"wsfunction": "core_user_create_users",
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

type updateUserResponse struct {
	Attempt  *userResponse `json:"attempt,omitempty"`
	Warnings Warnings      `json:"warnings,omitempty"`
}

// param is Encoded of type UserRequests []*UserRequest
//
//	param := []*moodleClient.UserRequest{
//		{
//			Username:  "xxxx",
//			Firstname: "xxx",
//			Lastname:  "xxx",
//			Email:     "xxx@gmail.com",
//			Password:  "@123xxxx",
//			Auth:      "manual",
//		},
//	}
//
// resutl, err := query.Values(moodleClient.Options{Users: param})
// res, err := Client.UserAPI.CreateUsers(ctx, resutl.Encode())
func (u *userAPI) UpdateUsers(ctx context.Context, param UserOptions) (*updateUserResponse, error) {
	qparam, err := query.Values(param)
	if err != nil {
		return nil, err
	}
	res := updateUserResponse{}
	err = u.callMoodleFunctionPost(ctx, &res, qparam.Encode(), map[string]string{
		"wsfunction": "core_user_update_users",
	})
	if err != nil {
		return nil, err
	}
	if len(res.Warnings) > 0 {
		return nil, res.Warnings
	}
	return &res, nil
}

type ManualEnrolUser struct {
	RoleId    int  `json:"roleid" url:"roleid"`       //Role to assign to the user
	UserId    int  `json:"userid" url:"userid"`       //The user that is going to be enrolled
	CourseId  int  `json:"courseid" url:"courseid"`   //The course to enrol the user role in
	TimeStart *int `json:"timestart" url:"timestart"` //Timestamp when the enrolment start
	TimeEnd   *int `json:"timeend" url:"timeend"`     //Timestamp when the enrolment end
	Suspend   *int `json:"suspend" url:"suspend"`     //set to 1 to suspend the enrolment
}

type ManualEnrolUsers []*ManualEnrolUser

type EnrolMentOptions struct {
	EnrolMents ManualEnrolUsers `url:"enrolments"`
}

func (us ManualEnrolUsers) EncodeValues(key string, v *url.Values) error {
	for i, u := range us {
		res, err := query.Values(u)
		if err != nil {
			return err
		}
		for subKey, subVal := range res {
			_ = subVal
			val := res.Get(subKey)
			if val != "<nil>" && val != "" && val != "0" {
				v.Set(fmt.Sprintf("%s[%d][%s]", key, i, subKey), res.Get(subKey))
			}
		}
	}
	return nil
}

// enrol_manual_enrol_users
func (u *userAPI) ManualEnrolUsers(ctx context.Context, param EnrolMentOptions) error {
	qparam, err := query.Values(param)
	if err != nil {
		return err
	}
	res := updateUserResponse{}
	err = u.callMoodleFunctionPost(ctx, &res, qparam.Encode(), map[string]string{
		"wsfunction": "enrol_manual_enrol_users",
	})
	if err != nil {
		return err
	}
	if len(res.Warnings) > 0 {
		return res.Warnings
	}
	return nil
}
