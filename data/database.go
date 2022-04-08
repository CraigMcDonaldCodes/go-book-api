package data

type Db interface {
	Get(id int64) Book
	GetAll() []Book
	Delete(id int64) bool
	Update(id int64) bool
}

type fakeDb struct {
}

func (db fakeDb) Get(id int64) Book {

	return Book{
		Id:        1,
		Isbn:      "foo",
		Title:     "bar",
		Author:    "foo",
		Publisher: "bar",
	}
}

func (db fakeDb) GetAll() []Book {

	books := []Book{
		{
			Id:        1,
			Isbn:      "foo",
			Title:     "bar",
			Author:    "foo",
			Publisher: "bar",
		},
		{
			Id:        2,
			Isbn:      "foo",
			Title:     "baz",
			Author:    "foo",
			Publisher: "baz",
		},
		{
			Id:        3,
			Isbn:      "foo",
			Title:     "bat",
			Author:    "foo",
			Publisher: "bat",
		},
	}

	return books
}

func (db *fakeDb) Delete(id int64) bool {
	return false
}

func (db *fakeDb) Update(id int64) bool {
	return false
}

func NewFakeDb() Db {
	return nil
}
