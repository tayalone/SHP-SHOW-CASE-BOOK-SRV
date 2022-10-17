package services

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/core/domains"
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/core/ports"
)

type mockRepo struct{}

var desc = "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s"

var tmpBook = domains.Book{
	ID:        1,
	Title:     "Lorem",
	Desc:      &desc,
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

func (m *mockRepo) GetByPk(id uint) (domains.Book, error) {
	if id == 1 {
		return tmpBook, nil
	}
	return domains.Book{}, errors.New("Book does not exist")
}

var mr = &mockRepo{}

func TestNew(t *testing.T) {
	type args struct {
		r ports.BookRpstr
	}
	tests := []struct {
		name string
		args args
		want ports.BookSrv
	}{
		// TODO: Add test cases.
		{
			name: "Case: Define Book Service Instant",
			args: args{
				r: mr,
			},
			want: New(mr),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBookService_GetByID(t *testing.T) {
	type fields struct {
		repo ports.BookRpstr
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domains.Book
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &BookService{
				repo: tt.fields.repo,
			}
			got, err := s.GetByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("BookService.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BookService.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
