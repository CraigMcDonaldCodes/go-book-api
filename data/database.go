package data

type Db interface {
	Get(id int64) Book
	GetAll() []Book
	Delete(id int64) bool
	Update(id int64) bool
}

type fakeDb struct {
}

func NewFakeDb() Db {
	return nil
}
