package db

import (
	"tina/packages/types"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UsersModel() *gorm.DB {
	return DB.Model(types.User{})
}

func FindUser(query interface{}, args ...interface{}) (*types.User, error) {
	var user types.User

	result := DB.Model(types.User{}).Where(query, args...).Find(&user)

	if result.RowsAffected == 0 {
		return nil, result.Error
	} else {
		return &user, nil
	}
}

func InsertUser(user *types.User) error {
	result := DB.Model(types.User{}).Create(user)
	return result.Error
}
func UpdateUser(user *types.User) error {
	result := DB.Model(types.User{}).Where("id = ?", user.Id).Updates(user)
	return result.Error
}

func FindUsers(query interface{}, offset int, limit int, args ...interface{}) ([]types.User, error) {
	var users []types.User
	result := DB.Model(types.User{}).Select("*").Where(query, args...).Offset(offset).Limit(limit).Find(users)
	return users, result.Error
}

func RegisterUser(username string, nexusId uuid.UUID, defaultUsage uint64, status types.UserStatus) (*types.User, error) {
	user := types.User{
		Id:             uuid.New(),
		Username:       username,
		NexusId:        nexusId,
		Status:         status,
		Credit:         0,
		UsageSoftLimit: 500000,
		UsageHardLimit: 10000000,
		Usage:          defaultUsage,
	}
	err := InsertUser(&user)
	return &user, err
}

func UpdateUserUsage(username string, nexusId uuid.UUID, usage uint64, status types.UserStatus) error {
	user, err := FindUser("nexus_id = ?", nexusId)
	if err != nil {
		return err
	}
	if user == nil {
		user, err = RegisterUser(username, nexusId, usage, types.UserStatusActive)
	}
	if err != nil {
		return err
	}
	user.Usage += usage
	user.Credit = types.CreditFromUsage(user.Usage)
	result := DB.Save(&user)
	return result.Error
}
