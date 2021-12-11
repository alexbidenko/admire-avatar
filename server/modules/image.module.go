package modules

import (
	"admire-avatar/config"
	"admire-avatar/entities"
)

type ImageModule struct {
}

func (i *ImageModule) Get(userID uint, imageType string) []entities.Image {
	images := make([]entities.Image, 0)
	config.DB.Model(&entities.Image{}).Where("user_id = ?", userID).Where("type = ?", imageType).Order("created_at desc").Find(&images)
	return images
}

func (i *ImageModule) Paginate(userID uint, imageType string, offset int, limit int) []entities.Image {
	images := make([]entities.Image, 0)
	config.DB.Model(&entities.Image{}).
		Where("user_id = ?", userID).
		Where("type = ?", imageType).
		Offset(offset).
		Limit(limit).
		Order("created_at desc").
		Find(&images)
	return images
}

func (i *ImageModule) GetAvatar(userID uint) (entities.Image, error) {
	var image entities.Image
	err := config.DB.Model(&entities.Image{}).Where("user_id = ?", userID).Where("main = ?", true).Find(&image).Error
	return image, err
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

func (i *ImageModule) Clear(userID uint, imageType string) {
	config.DB.Where("user_id = ?", userID).Where("type = ?", imageType).Delete(&entities.Image{})
}

func (i *ImageModule) CreateAvatar(userID uint, imageID string) {
	config.DB.Model(&entities.Image{}).Where("user_id = ?", userID).Update("main", false)
	config.DB.Model(&entities.Image{}).Where("user_id = ?", userID).Where("id = ?", imageID).Update("main", true)
}

func (i *ImageModule) PrintToAvatar(userID uint, imageID string) {
	config.DB.Model(&entities.Image{}).Where("id = ?", imageID).Update("type", "avatar")
}
