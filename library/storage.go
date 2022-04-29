package library

import "database/sql"

type SQLStorage struct {
	db *sql.DB
}

func (s *SQLStorage) GetAllBooks() ([]Book, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SQLStorage) GetBooksByAuthor(author string) ([]Book, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SQLStorage) GetBooksByName(name string) ([]Book, error) {
	query := `SELECT id, name, author, cnt FROM books WHERE name = $1`

	rows, err := s.db.Query(query, name)
	if err != nil {
		return nil, err
	}

	var res = make([]Book, 0)

	for rows.Next() {
		var book Book
		err = rows.Scan(&book.ID, &book.Name, &book.Author, &book.Count)
		if err != nil {
			return nil, err
		}
		res = append(res, book)
	}

	return res, nil
}

func (s *SQLStorage) Get(id int) Book {
	//TODO implement me
	panic("implement me")
}

func (s *SQLStorage) Save(book Book) (Book, error) {
	//TODO implement me
	panic("implement me")
}
