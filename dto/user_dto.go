package dto

import "gin/model"

//向前台返回数据的格式
type UserDto struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:  user.Name,
		Phone: user.Phone,
	}
}
