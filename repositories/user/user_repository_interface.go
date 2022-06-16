package user

import "alterra/entities"

type UserRepositoryInterface interface {

	/*
	 * Find User
	 * -------------------------------
	 * Mencari user berdasarkan ID
	 */
	FindAllUser(limit int, offset int, filters []map[string]string, sorts []map[string]interface{}) ([]entities.User, error)

	/*
	 * Find User
	 * -------------------------------
	 * Mencari user berdasarkan ID
	 */
	Find(id int) (entities.User, error)

	/*
	 * Find By Column
	 * -------------------------------
	 * Mencari user tunggal berdasarkan column dan value
	 */
	FindBy(field string, value string) (entities.User, error)

	/*
	 * Store
	 * -------------------------------
	 * Menambahkan user tunggal kedalam database
	 */
	Store(user entities.User) (entities.User, error)

	/*
	 * Update User
	 * -------------------------------
	 * Mengedit user tunggal
	 */
	Update(user entities.User) (entities.User, error)

	/*
	 * Delete
	 * -------------------------------
	 * Delete user tunggal berdasarkan ID
	 */
	Delete(id int) error
}
