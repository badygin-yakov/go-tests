package library

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestSQLStorage_GetBooksByName(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		name string
	}

	type want struct {
		books []Book
		err   error
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name: "success get books by name",
			fields: fields{
				db: func(t *testing.T) *sql.DB {
					db, mock, err := sqlmock.New()
					if err != nil {
						t.FailNow()
					}

					mock.
						ExpectQuery(`SELECT id, name, author, cnt FROM books WHERE name = \$1`).
						WithArgs("example 1").
						WillReturnRows(
							sqlmock.NewRows([]string{"id", "name", "author", "cnt"}).
								AddRow(1, "example 1", "author 1", 11),
						)

					return db
				}(t),
			},
			args: args{
				name: "example 1",
			},
			want: want{
				books: []Book{{
					ID: 1, Count: 11,
					Name: "example 1", Author: "author 1",
				}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SQLStorage{
				db: tt.fields.db,
			}
			got, err := s.GetBooksByName(tt.args.name)
			if err != nil && !errors.Is(err, tt.want.err) {
				t.Errorf("GetBooksByName() error = %v, wantErr %v", err, tt.want.err)
				return
			}
			if !reflect.DeepEqual(got, tt.want.books) {
				t.Errorf("GetBooksByName() got = %v, want %v", got, tt.want)
			}
		})
	}
}
