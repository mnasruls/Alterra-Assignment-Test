package user

import (
	"alterra/entities"
	"alterra/entities/web"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
	db *gorm.DB
}

// Constructor
func NewUserRepository(database *gorm.DB) *UserRepository {
	return &UserRepository{
		db: database,
	}
}

func (repo UserRepository) FindAllUser(limit int, offset int, filters []map[string]string, sorts []map[string]interface{}) ([]entities.User, error) {
	users := []entities.User{}
	builder := repo.db.Limit(limit).Offset(offset)
	// Where filters
	for _, filter := range filters {
		builder.Where(filter["field"]+" "+filter["operator"]+" ?", filter["value"])
	}
	// OrderBy Filters
	for _, sort := range sorts {
		builder.Order(clause.OrderByColumn{Column: clause.Column{Name: sort["field"].(string)}, Desc: sort["desc"].(bool)})
	}
	tx := builder.Where("role=?", "user").Find(&users)
	if tx.Error != nil {
		return []entities.User{}, web.WebError{Code: 500, Message: tx.Error.Error()}
	}
	return users, nil
}

/*
 * Find User by ID
 * -------------------------------
 * Mencari user berdasarkan ID
 */
func (repo UserRepository) Find(id int) (entities.User, error) {

	// Get user dari database
	user := entities.User{}
	tx := repo.db.Find(&user, id)
	if tx.Error != nil {

		// Return error dengan code 500
		return entities.User{}, web.WebError{Code: 500, Message: tx.Error.Error()}
	} else if tx.RowsAffected <= 0 {

		// Return error dengan code 400 jika tidak ditemukan
		return entities.User{}, web.WebError{Code: 400, Message: "cannot get user data with specified id"}
	}
	return user, nil
}

/*
 * Find By Column
 * -------------------------------
 * Mencari user tunggal berdasarkan column dan value
 */
func (repo UserRepository) FindBy(field string, value string) (entities.User, error) {

	// Get user dari database
	user := entities.User{}
	tx := repo.db.Where(field+" = ?", value).Find(&user)
	if tx.Error != nil {

		// return kode 500 jika terjadi error
		return entities.User{}, web.WebError{Code: 500, Message: tx.Error.Error()}
	} else if tx.RowsAffected <= 0 {

		// return kode 400 jika tidak ditemukan
		return entities.User{}, web.WebError{Code: 400, Message: "doesn't match with any record"}
	}
	return user, nil
}

/*
 * Store
 * -------------------------------
 * Menambahkan user tunggal kedalam database
 */
func (repo UserRepository) Store(user entities.User) (entities.User, error) {

	// insert user ke database
	tx := repo.db.Create(&user)
	if tx.Error != nil {

		// return kode 500 jika error
		return entities.User{}, web.WebError{Code: 500, Message: tx.Error.Error()}
	}
	return user, nil
}

/*
 * Update User
 * -------------------------------
 * Mengedit user tunggal berdasarkan ID
 */
func (repo UserRepository) Update(user entities.User) (entities.User, error) {

	// Update database
	tx := repo.db.Save(&user)
	if tx.Error != nil {

		// return Kode 500 jika error
		return entities.User{}, web.WebError{Code: 500, Message: tx.Error.Error()}
	}
	return user, nil
}

/*
 * Delete
 * -------------------------------
 * Delete user tunggal berdasarkan ID
 */
func (repo UserRepository) Delete(id int) error {

	// Delete from database
	tx := repo.db.Delete(&entities.User{}, id)
	if tx.Error != nil {

		// return kode 500 jika error
		return web.WebError{Code: 500, Message: tx.Error.Error()}
	}
	return nil
}
