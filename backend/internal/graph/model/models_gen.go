// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type ActivityItem struct {
	Title       string       `json:"title"`
	ImageURL    *string      `json:"imageUrl"`
	Description *string      `json:"description"`
	Type        ActivityType `json:"type"`
	Labels      []*string    `json:"labels"`
	Link        string       `json:"link"`
}

type ActivityType string

const (
	ActivityTypeArticle ActivityType = "ARTICLE"
	ActivityTypePodcast ActivityType = "PODCAST"
	ActivityTypeFact    ActivityType = "FACT"
)

var AllActivityType = []ActivityType{
	ActivityTypeArticle,
	ActivityTypePodcast,
	ActivityTypeFact,
}

func (e ActivityType) IsValid() bool {
	switch e {
	case ActivityTypeArticle, ActivityTypePodcast, ActivityTypeFact:
		return true
	}
	return false
}

func (e ActivityType) String() string {
	return string(e)
}

func (e *ActivityType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ActivityType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ActivityType", str)
	}
	return nil
}

func (e ActivityType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
