package main

import (
	"context"
	"fmt"
	"net/url"

	"github.com/mnasruul/moodleClient"
)

func main() {
	ctx := context.Background()
	serviceURL, err := url.Parse("https://learning.pt-ssss.com")
	if err != nil {
		panic(err)
	}
	Client, err := moodleClient.NewClientWithLogin(
		ctx,
		serviceURL,
		"xxx",
		"xxxxxx",
	)
	if err != nil {
		panic(err)
	}

	siteInfo, err := Client.SiteAPI.GetSiteInfo(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", siteInfo)

	courses, err := Client.CourseAPI.GetEnrolledCoursesByTimelineClassification(
		ctx,
		moodleClient.CourseClassificationInProgress,
	)
	if err != nil {
		panic(err)
	}

	for _, c := range courses {
		fmt.Printf("%#v\n", c)
	}
}
