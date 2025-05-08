package db

type DB interface {
	DBSave(t any) error
	DBGet(t any, offset int64, limit int64) (any, error)
}
