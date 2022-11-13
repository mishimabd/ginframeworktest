package service

import "gitlab.com/pragmaticreviews/golang-gin-poc/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}
type videoService struct {
	videos []entity.Video
}

func (v videoService) Save(video entity.Video) entity.Video {
	v.videos = append(v.videos, video)
	return v.videos
}

func (v videoService) FindAll() []entity.Video {
	return v.videos
}

func New() VideoService {
	return &videoService{}
}
