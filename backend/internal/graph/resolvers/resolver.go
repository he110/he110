package resolvers

import "He110/PersonalWebSite/internal/providers/manager/activity_manager"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	activityManager *activity_manager.ActivityManager
}

func NewResolver(am *activity_manager.ActivityManager) *Resolver {
	return &Resolver{
		activityManager: am,
	}
}
