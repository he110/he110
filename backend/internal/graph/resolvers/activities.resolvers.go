package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"He110/PersonalWebSite/internal/generated"
	"He110/PersonalWebSite/internal/graph/model"
)

var typeMap = map[string]model.ActivityType{
	generated.ActivityType_ARTICLE.String(): model.ActivityTypeArticle,
	generated.ActivityType_PODCAST.String(): model.ActivityTypePodcast,
	generated.ActivityType_FACT.String(): model.ActivityTypeFact,
}

func (r *queryResolver) Activities(ctx context.Context) ([]*model.ActivityItem, error) {
	items, err := r.activityManager.FindAll()
	if err != nil {
		return nil, err
	}

	var activities []*model.ActivityItem
	for _, i := range items {
		itemType, exists := typeMap[i.Type.String()]
		if !exists {
			itemType = model.ActivityTypeFact
		}
		activity := model.ActivityItem{
			Title:       i.Title,
			ImageURL:    &i.ImageUrl,
			Description: &i.Description,
			Type:        itemType,
			Link:        i.Link,
		}
		activities = append(activities, &activity)
	}

	return activities, nil
}
