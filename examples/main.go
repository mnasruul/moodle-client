package main

import (
	"context"
	"fmt"
	"net/url"

	"github.com/mnasruul/moodleClient"
)

func main() {
	ctx := context.Background()
	serviceURL, err := url.Parse("https://xx.xxx.com")
	if err != nil {
		panic(err)
	}
	Client, err := moodleClient.NewClient(
		ctx,
		serviceURL,
		"x",
		moodleClient.WithDebugEnabled(),
	)
	if err != nil {
		panic(err)
	}

	// siteInfo, err := Client.SiteAPI.GetSiteInfo(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("%#v\n", siteInfo)

	courses, err := Client.CourseAPI.GetCoursesByField(
		ctx,
		moodleClient.CoursesByFieldOptions{CourseFieldValues: &moodleClient.CourseFieldValues{{Field: moodleClient.CourseFieldCategory, Value: "2"}}},
	)
	// courses, err := Client.CourseAPI.GetEnrolledCoursesByTimelineClassification(
	// 	ctx,
	// 	moodleClient.CourseClassificationFuture)
	// id :=  make(moodleClient.CourseIds)}
	// courses, err := Client.CourseAPI.GetCourses(
	// 	ctx,
	// 	moodleClient.CoursesOptions{Options: moodleClient.CourseIds{3, 4}},
	// )
	if err != nil {
		panic(err)
	}

	for _, c := range courses {
		fmt.Printf("%#v\n", c)
	}

	param := []*moodleClient.UserRequest{
		{
			Id:       41,
			Password: "#Perang96",
		},
	}

	// fmt.Printf("%+v", resutl.Encode())
	res, err := Client.UserAPI.UpdateUsers(ctx, moodleClient.UserOptions{Users: param})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", res)

	paramEnrol := moodleClient.EnrolMentOptions{
		EnrolMents: moodleClient.ManualEnrolUsers{
			&moodleClient.ManualEnrolUser{
				RoleId:   5,
				UserId:   41,
				CourseId: 3,
			},
		},
	}

	err = Client.UserAPI.ManualEnrolUsers(ctx, paramEnrol)
	if err != nil {
		panic(err)
	}
}
