package service

import (
	"github.com/abdoohossamm/golang-gin-first/entity"
)

type VideoService interface {
	Save(video entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func (service videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service videoService) FindAll() []entity.Video {
	return service.videos
}

func New() VideoService {
	return &videoService{}
}
