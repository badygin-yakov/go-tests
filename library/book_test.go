package library

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestBookEqual(t *testing.T) {
	type args struct {
		b1, b2 *Book
	}
	type want struct {
		res bool
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "test equal true",
			args: args{
				b1: &Book{Name: "example 1", Author: "author 1"},
				b2: &Book{Name: "example 1", Author: "author 1"},
			},
			want: want{
				res: true,
			},
		},
		{
			name: "test equal false",
			args: args{
				b1: &Book{Name: "example 1", Author: "author 1"},
				b2: &Book{Name: "example 2", Author: "author 1"},
			},
			want: want{
				res: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BookEqual(tt.args.b1, tt.args.b2)
			if got != tt.want.res {
				t.Logf("got: %v, want: %v", got, tt.want.res)
				t.FailNow()
			}
		})
	}
}

//type mockStorage struct {
//	books []Book
//}
//
//func (s *mockStorage) SetBooks(books []Book) {
//	s.books = books
//}
//
//var _ Storage = (*mockStorage)(nil)
//
//func (s *mockStorage) GetAllBooks() ([]Book, error) {
//	return s.books, nil
//}

func TestBookService_GetAll(t *testing.T) {
	type want struct {
		books []Book
		err   error
	}

	tests := []struct {
		name    string
		prepare func(*MockStorage)
		want    want
	}{
		{
			name: "get all books success",
			prepare: func(s *MockStorage) {
				s.EXPECT().
					GetAllBooks().
					Times(1).
					Return([]Book{{Name: "example 1", Author: "author 1"}}, nil)
			},
			want: want{
				books: []Book{
					{Name: "example 1", Author: "author 1"},
				},
			},
		},
		{
			name: "get all books success two",
			prepare: func(s *MockStorage) {
				s.EXPECT().
					GetAllBooks().
					Times(1).
					Return([]Book{
						{Name: "example 1", Author: "author 1"},
						{Name: "example 2", Author: "author 1"},
					}, nil)
			},
			want: want{
				books: []Book{
					{Name: "example 1", Author: "author 1"},
					{Name: "example 2", Author: "author 1"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockStorage := NewMockStorage(ctrl)

			service := &BookService{
				storage: mockStorage,
			}
			tt.prepare(mockStorage)

			got, err := service.GetAll()
			if err != nil && !errors.Is(err, tt.want.err) {
				t.Errorf("got error, but not expected: %v", err)
				t.FailNow()
			}
			if !reflect.DeepEqual(got, tt.want.books) {
				t.Errorf("got isn't equat want: %v, %v", got, tt.want.books)
			}
		})
	}
}

func TestBookService_GetByAuthor(t *testing.T) {
	type args struct {
		author string
	}
	type want struct {
		books []Book
		err   error
	}

	tests := []struct {
		name    string
		prepare func(*MockStorage)
		args    args
		want    want
	}{
		{
			name: "get all books success",
			prepare: func(s *MockStorage) {
				s.EXPECT().
					GetBooksByAuthor("author 1").
					Times(1).
					Return([]Book{{Name: "example 1", Author: "author 1"}}, nil)
			},
			args: args{
				author: "Author 1",
			},
			want: want{
				books: []Book{
					{Name: "example 1", Author: "author 1"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockStorage := NewMockStorage(ctrl)

			service := &BookService{
				storage: mockStorage,
			}
			tt.prepare(mockStorage)

			got, err := service.GetByAuthor(tt.args.author)
			if err != nil && !errors.Is(err, tt.want.err) {
				t.Errorf("got error, but not expected: %v", err)
				t.FailNow()
			}
			if !reflect.DeepEqual(got, tt.want.books) {
				t.Errorf("got isn't equat want: %v, %v", got, tt.want.books)
			}
		})
	}
}

func TestBookService_GetByID(t *testing.T) {
	type args struct {
		id int
	}
	type want struct {
		book Book
	}

	tests := []struct {
		name    string
		prepare func(*MockStorage)
		args    args
		want    want
	}{
		{
			name: "get all books success",
			prepare: func(s *MockStorage) {
				s.EXPECT().Get(1).
					Times(1).
					Return(Book{
						ID: 2, Count: 10,
						Name: "example 1", Author: "author 1",
					})

				s.EXPECT().Save(Book{
					ID: 2, Count: 11,
					Name: "example 1", Author: "author 1",
				}).Return(
					Book{
						ID: 1, Count: 10,
						Name: "example 1", Author: "author 1",
					}, nil)
			},
			args: args{
				id: 1,
			},
			want: want{
				book: Book{
					ID: 1, Count: 10,
					Name: "example 1", Author: "author 1",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockStorage := NewMockStorage(ctrl)

			service := &BookService{
				storage: mockStorage,
			}
			tt.prepare(mockStorage)

			got := service.GetByID(tt.args.id)
			if !reflect.DeepEqual(got, tt.want.book) {
				t.Errorf("got isn't equat want: %v, %v", got, tt.want.book)
			}
		})
	}
}
