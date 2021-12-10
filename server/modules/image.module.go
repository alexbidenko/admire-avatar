package modules

import (
	"admire-avatar/config"
	"admire-avatar/entities"
)

type ImageModule struct {
}

func (i *ImageModule) Get(userID uint) []entities.Image {
	images := make([]entities.Image, 0)
	config.DB.Model(&entities.Image{}).Where("user_id = ?", userID).Find(&images)
	return images
}

func (i *ImageModule) Find(userID uint, imageID string) (entities.Image, error) {
	var image entities.Image
	err := config.DB.Model(&entities.Image{}).Where("user_id = ?", userID).Where("id = ?", imageID).First(&image).Error
	return image, err
}

func (i *ImageModule) Create(image *entities.Image) {
	config.DB.Model(&entities.Image{}).Create(image)
}

func (i *ImageModule) Delete(id uint) {
	config.DB.Delete(&entities.Image{}, id)
}

func (i *ImageModule) CreateAvatar(userID uint, imageID string) {
	config.DB.Model(&entities.Image{}).Where("user_id = ?", userID).Update("main", false)
	config.DB.Model(&entities.Image{}).Where("user_id = ?", userID).Where("id = ?", imageID).Update("main", true)
}
