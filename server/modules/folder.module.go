package modules

import (
	"admire-avatar/config"
	"admire-avatar/entities"
)

type FolderModule struct {
}

func (f *FolderModule) Get(userID uint) []entities.Folder {
	folders := make([]entities.Folder, 0)
	config.DB.Model(&entities.Folder{}).Where("user_id = ?", userID).Order("name asc").Find(&folders)
	return folders
}

func (f *FolderModule) GetPublic() []entities.Folder {
	folders := make([]entities.Folder, 0)
	config.DB.Model(&entities.Folder{}).Where("public = ?", true).Order("name asc").Find(&folders)
	return folders
}

func (f *FolderModule) GetByName(userID uint, name string) (entities.Folder, error) {
	var folder entities.Folder
	err := config.DB.Model(&entities.Folder{}).Where("user_id = ?", userID).Where("name = ?", name).First(&folder).Error
	return folder, err
}

func (f *FolderModule) Find(id string) (entities.Folder, error) {
	var folder entities.Folder
	err := config.DB.Model(&entities.Folder{}).First(&folder, id).Error
	return folder, err
}

func (f *FolderModule) Update(userID uint, id string, folder *entities.Folder) {
	config.DB.Model(&entities.Folder{}).Where("user_id = ?", userID).Where("id = ?", id).Update("name", folder.Name).Update("public", folder.Public)
}

func (f *FolderModule) Create(folder *entities.Folder) {
	config.DB.Model(&entities.Folder{}).Create(folder)
}

func (f *FolderModule) Delete(id string) {
	config.DB.Model(&entities.Image{}).Where("folder_id = ?", id).Update("folder_id", 0)
	config.DB.Delete(&entities.Folder{}, id)
}
