package moodleClient

import (
	"time"
)

type CourseClassification string

type CourseFieldType string

const (
	CourseClassificationPast       CourseClassification = "past"
	CourseClassificationInProgress CourseClassification = "inprogress"
	CourseClassificationFuture     CourseClassification = "future"
	CourseFieldCategory            CourseFieldType      = "category"
)

type Course struct {
	ID              int
	FullName        string
	ShortName       string
	Summary         string
	SummaryFormat   int
	StartDate       time.Time
	EndDate         time.Time
	Visible         bool
	FullNameDisplay string
	ViewURL         string
	CourseImage     string
	Progress        int
	HasProgress     bool
	IsSavourite     bool
	Hidden          bool
	ShowShortName   bool
	CourseCategory  string
}
