package mods

import (
	"context"
	"errors"
	mydb "test_bi/db"
)

func GetUsers() ([]mydb.User, error) {
	var users []mydb.User
	result := db.Find(&users)
	return users, result.Error
}

func GetUser(id int) (mydb.User, error) {
	var user mydb.User
	result := db.First(&user, id)
	if result.Error != nil {
		return user, errors.New("用户不存在")
	}
	return user, nil
}

func CreateUser(user mydb.User) (mydb.User, error) {
	result := db.Create(&user)
	return user, result.Error
}

func UpdateUser(id int, user mydb.User) (mydb.User, error) {
	var existingUser mydb.User
	if result := db.First(&existingUser, id); result.Error != nil {
		return user, errors.New("用户不存在")
	}
	db.Model(&existingUser).Updates(&user)
	return existingUser, nil
}

func DeleteUser(id int) error {
	result := db.Delete(&mydb.User{}, id)
	return result.Error
}

func CacheTest() (string, error) {
	val, err := cache.Get(context.Background(), "test_key").Result()
	if err != nil {
		err = cache.Set(context.Background(), "test_key", "缓存数据", 0).Err()
		if err != nil {
			return "", err
		}
		return "首次访问，缓存已创建", nil
	}
	return val, nil
}
