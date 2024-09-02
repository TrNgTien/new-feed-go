package controller

import "github.com/TrNgTien/new-feed-go/internal/services"

type FeedController struct {
	Service services.AuthService
}

func NewFeedController(service services.AuthService) *FeedController {
	return &FeedController{Service: service}
}
