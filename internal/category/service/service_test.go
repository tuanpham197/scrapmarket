package service

import (
	"context"
	"errors"
	"fmt"
	"sendo/internal/category/service/entity"
	"sendo/internal/category/service/request"
	"sendo/pkg/utils/paginations"
	"testing"

	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	catOne = entity.Category{
		Id:        uuid.MustParse("c9998fdc-6b18-4322-8e6c-93b0ddf71bbe"),
		Name:      "test 1",
		Thumbnail: "test",
	}

	catTwo = entity.Category{
		Id:        uuid.MustParse("cb457482-c26f-436d-a1b5-4fce3173673d"),
		Name:      "test 2",
		Thumbnail: "test",
	}

	emptyTranslation = entity.Category{}

	errorCreate = errors.New("create fail")
	ErrDB       = errors.New("db error")
)

type mockRepository struct{}

func (mockRepository) Insert(ctx context.Context, name, thumbnail, parentId string) (*entity.Category, error) {
	cat := entity.Category{
		Name:      name,
		Thumbnail: thumbnail,
	}
	return &cat, nil
}

func (mockRepository) GetList(ctx context.Context, queryParam request.QueryParam) (*paginations.Pagination, error) {
	dataSet := []entity.Category{catOne, catTwo}
	result := []entity.Category{}
	cnt := 0
	for _, v := range dataSet {
		if v.Name == queryParam.Name {
			result = append(result, v)
			cnt += 1
		}
	}

	dataPagi := paginations.Pagination{
		Records: result,
		Count:   int64(cnt),
	}
	if len(result) > 0 {
		return &dataPagi, nil
	}

	return nil, nil
}

func (mockRepository) UpdateThumbnail(ctx context.Context, thumbnailPath string, category *entity.Category) (bool, error) {
	panic("err")
}

func (mockRepository) FindOne(ctx context.Context, id string) (*entity.Category, error) {
	panic("err")
}

func Test_Insert_Category(t *testing.T) {

	// setup
	repo := mockRepository{}
	catSv := NewService(repo)

	Convey("Given some data to insert", t, func() {
		testTable := []struct {
			name        string
			thumbnail   string
			expectError error
			expectData  *entity.Category
		}{
			{
				name:        catOne.Name,
				thumbnail:   catOne.Thumbnail,
				expectError: nil,
				expectData:  &catOne,
			},
			{
				name:        catTwo.Name,
				thumbnail:   catTwo.Thumbnail,
				expectError: nil,
				expectData:  &catTwo,
			},
			{
				name:        "",
				thumbnail:   "",
				expectError: errorCreate,
				expectData:  nil,
			},
		}

		Convey("Start insert data", func() {
			for i, item := range testTable {
				cat, err := catSv.repository.Insert(context.Background(), item.name, item.thumbnail, "")
				conveySuit := fmt.Sprintf("compare with case %d", i)
				Convey(conveySuit, func() {
					So(err, ShouldEqual, nil)
					if item.name != "" {
						catExpect := *item.expectData
						So(cat.Name, ShouldEqual, catExpect.Name)
						So(cat.Thumbnail, ShouldEqual, catExpect.Thumbnail)
					}
				})
			}
		})
	})
}

// func Test_Insert_Category(t *testing.T) {

// }
