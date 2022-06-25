package daos

import (
	"rest/database"
	"rest/models"
)

type UsersDao struct{}

func (u UsersDao) CreateUser(user models.User) (models.User, error) {
	Db, DbContext := database.GetDatabaseAndContext()
	_, err := Db.NewInsert().Model(&user).Exec(DbContext)
	return user, err
}

func (u UsersDao) GetUser(credentials map[string]string) (models.User, error) {
	Db, DbContext := database.GetDatabaseAndContext()
	user := models.User{}
	err := Db.NewSelect().Model(&user).Where("username = ? OR email = ?", credentials["id"], credentials["id"]).Scan(DbContext)
	return user, err
}
