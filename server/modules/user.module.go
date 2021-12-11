package modules

import (
	"admire-avatar/config"
	"admire-avatar/entities"
)

type UserModule struct {
}

func (userModule *UserModule) All() []entities.BaseUser {
	var users []entities.BaseUser
	config.DB.Model(&entities.User{}).Order("created_at asc").Find(&users)
	return users
}

func (userModule *UserModule) Search(email string, exclude uint) []entities.BaseUser {
	var users []entities.BaseUser
	config.DB.Model(&entities.User{}).Where("email like ?", "%"+email+"%").Where("id <> ?", exclude).Offset(0).Limit(20).Order("created_at asc").Find(&users)
	return users
}

func (userModule *UserModule) Create(user *entities.User) {
	config.DB.Model(&entities.User{}).Create(&user)
}

func (userModule *UserModule) Update(id uint, user *entities.User) {
	config.DB.Model(&entities.User{}).Where("id = ?", id).Updates(user).First(user, "id = ?", id)
}

func (userModule *UserModule) Delete(id string) {
	config.DB.Delete(&entities.User{}, id)
}

func (userModule *UserModule) GetByEmail(email string) (entities.User, error) {
	var user entities.User
	err := config.DB.Model(&entities.User{}).First(&user, "email = ?", email).Error
	return user, err
}

func (userModule *UserModule) GetByHash(emailHash string) (entities.User, error) {
	var user entities.User
	err := config.DB.Model(&entities.User{}).First(&user, "email_hash = ?", emailHash).Error
	return user, err
}

func (userModule *UserModule) Find(ID uint) (entities.BaseUser, error) {
	var user entities.BaseUser
	err := config.DB.Model(&entities.User{}).First(&user, "id = ?", ID).Error
	return user, err
}
