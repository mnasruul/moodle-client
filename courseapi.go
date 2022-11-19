package moodleClient

import (
	"context"
	"fmt"
	"net/url"
	"reflect"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/mnasruul/moodleClient/pkg/urlutil"
)

type CourseAPI interface {
	GetEnrolledCoursesByTimelineClassification(ctx context.Context, classification CourseClassification) ([]*Course, error)
	GetCoursesByField(ctx context.Context, param CoursesByFieldOptions) ([]*Course, error)
	GetCourses(ctx context.Context, param CoursesOptions) ([]*Course, error)
}

type courseAPI struct {
	*apiClient
}

func newCourseAPI(apiClient *apiClient) *courseAPI {
	return &courseAPI{apiClient}
}

type courseResponse struct {
	ID            int    `json:"id"`
	FullName      string `json:"fullname"`
	ShortName     string `json:"shortname"`
	Summary       string `json:",omitempty"`
	SummaryFormat int    `json:"summaryformat"`
	StartDateUnix int64  `json:"startdate"`
	EndDateUnix   int64  `json:"enddate"`
	// Visible         bool   `json:"visible"`
	FullNameDisplay string `json:"fullnamedisplay"`
	ViewURL         string `json:"viewurl"`
	CourseImage     string `json:"courseimage"`
	Progress        int    `json:"progress"`
	HasProgress     bool   `json:"hasprogress"`
	IsSavourite     bool   `json:"isfavourite"`
	Hidden          bool   `json:"hidden"`
	ShowShortName   bool   `json:"showshortname"`
	CourseCategory  string `json:"coursecategory"`
}

type getEnrolledCoursesByTimelineClassificationResponse struct {
	Courses    []*courseResponse `json:"courses"`
	NextOffset int               `json:"nextoffset"`
}

type CourseByField struct {
	Id        string `json:"id,omitempty" url:"id"`
	Ids       string `json:"ids,omitempty" url:"ids"`
	ShortName string `json:"shortname,omitempty" url:"shortname"`
	IdNumber  string `json:"idnumber,omitempty" url:"idnumber"`
	Category  string `json:"category,omitempty" url:"category"`
}

func (c *courseAPI) GetEnrolledCoursesByTimelineClassification(ctx context.Context, classification CourseClassification) ([]*Course, error) {
	res := getEnrolledCoursesByTimelineClassificationResponse{}
	err := c.callMoodleFunction(ctx, &res, map[string]string{
		"wsfunction":     "core_course_get_enrolled_courses_by_timeline_classification",
		"classification": string(classification),
	})
	if err != nil {
		return nil, err
	}
	return mapToCourseList(res.Courses), nil
}

func mapToCourseList(courseResList []*courseResponse) []*Course {
	courses := make([]*Course, 0, len(courseResList))
	for _, courseRes := range courseResList {
		courses = append(courses, mapToCourse(courseRes))
	}
	return courses
}

func mapToCourse(courseRes *courseResponse) *Course {
	return &Course{
		ID:            courseRes.ID,
		FullName:      courseRes.FullName,
		ShortName:     courseRes.ShortName,
		Summary:       courseRes.Summary,
		SummaryFormat: courseRes.SummaryFormat,
		StartDate:     time.Unix(courseRes.StartDateUnix, 0),
		EndDate:       time.Unix(courseRes.StartDateUnix, 0),
		// Visible:         courseRes.Visible,
		FullNameDisplay: courseRes.FullName,
		ViewURL:         courseRes.ViewURL,
		CourseImage:     courseRes.CourseImage,
		Progress:        courseRes.Progress,
		HasProgress:     courseRes.HasProgress,
		IsSavourite:     courseRes.IsSavourite,
		Hidden:          courseRes.Hidden,
		ShowShortName:   courseRes.ShowShortName,
		CourseCategory:  courseRes.CourseCategory,
	}
}

type CourseFieldValues []*CourseFieldValue

type CourseFieldValue struct {
	Field CourseFieldType `json:"field"`
	Value string          `json:"value"`
}
type CoursesByFieldOptions struct {
	*CourseFieldValues
}

func (cs CourseFieldValues) EncodeValues(key string, v *url.Values) error {
FirstLoop:
	for _, u := range cs {
		res, err := query.Values(u)
		if err != nil {
			return err
		}
		element := make(url.Values)
		for subKey, subVal := range res {
			_ = subVal
			val := res.Get(subKey)
			field, _ := reflect.TypeOf(u).Elem().FieldByName(subKey)
			tagJsonName := urlutil.GetStructTag(field, "json")
			if len(val) == 0 {
				continue FirstLoop
			}
			element.Set(tagJsonName, val)
		}
		for key, val := range element {
			v.Set(key, val[0])
		}
	}
	return nil
}

func (c *courseAPI) GetCoursesByField(ctx context.Context, param CoursesByFieldOptions) ([]*Course, error) {
	resutl, _ := query.Values(param)
	res := getEnrolledCoursesByTimelineClassificationResponse{}
	err := c.callMoodleFunctionPost(ctx, &res, resutl.Encode(), map[string]string{
		"wsfunction": "core_course_get_courses_by_field",
	})
	if err != nil {
		return nil, err
	}
	return mapToCourseList(res.Courses), nil
}

type CourseId int
type CourseIds []CourseId
type CoursesOptions struct {
	Options CourseIds `url:"options"`
}

func (us CourseIds) EncodeValues(key string, v *url.Values) error {
	for i, u := range us {
		res, err := query.Values(u)
		if err != nil {
			return err
		}
		for subKey, subVal := range res {
			_ = subVal
			val := res.Get(subKey)
			v.Set(fmt.Sprintf("%s[%s][%d]", key, "ids", i), val)
		}
	}
	return nil
}

func (c *courseAPI) GetCourses(ctx context.Context, param CoursesOptions) ([]*Course, error) {
	resutl, _ := query.Values(param)
	res := getEnrolledCoursesByTimelineClassificationResponse{}
	err := c.callMoodleFunctionPost(ctx, &res, resutl.Encode(), map[string]string{
		"wsfunction": "core_course_get_courses",
	})
	if err != nil {
		return nil, err
	}
	return mapToCourseList(res.Courses), nil
}
