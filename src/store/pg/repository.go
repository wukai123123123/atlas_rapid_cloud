package pg

type Repository[T any] struct {
	db *DataBase
}
