package persistance

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/softree-group/kitchen-plan-backend/domain/entity"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	id = 1
	image = "/path/to/image/" + string(rune(id)) + ".jpg"
	title = "title"
	timeToCook = 2
	typeOfReceipt = "type"
)

func TestGetReceiptsSimpleSuccess(t *testing.T) {
	conn, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer func() {
		err = conn.Close()
		require.NoError(t, err)
	}()

	rows := sqlmock.NewRows([]string{"id", "image", "title", "time", "type"})
	expected := []entity.Receipt{
		{Id: id, Image: image, Title: title, TimeToCook: timeToCook, Type: typeOfReceipt},
	}
	rows = rows.AddRow(
		expected[0].Id,
		expected[0].Image,
		expected[0].Title,
		expected[0].TimeToCook,
		expected[0].Type,
	)

	selection := entity.Selection{}

	mock.ExpectBegin()
	mock.ExpectQuery("select r.id, image, title, time, type from recipes r").
		WithArgs().
		WillReturnRows(rows)
	mock.ExpectCommit()
	mock.ExpectClose()

	storage := NewPostgresStorage(conn)
	out, err := storage.GetReceipts(selection)
	require.NoError(t, err)

	fmt.Println(out)
}

func TestGetReceiptsIngredientsSuccess(t *testing.T) {
	conn, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer func() {
		err = conn.Close()
		require.NoError(t, err)
	}()

	rows := sqlmock.NewRows([]string{"id", "image", "title", "time", "type"})
	expected := []entity.Receipt{
		{Id: id, Image: image, Title: title, TimeToCook: timeToCook, Type: typeOfReceipt},
	}
	rows = rows.AddRow(
		expected[0].Id,
		expected[0].Image,
		expected[0].Title,
		expected[0].TimeToCook,
		expected[0].Type,
	)

	selection := entity.Selection{Ingredients: []int{708, 375}}

	mock.ExpectBegin()
	mock.ExpectQuery("select r.id, image, title, time, type from recipes r" +
		" join recipes_ingridients ri on r.id=ri.recept_id" +
		" where ri.ingridient_id in ").
		WithArgs().
		WillReturnRows(rows)
	mock.ExpectCommit()
	mock.ExpectClose()

	storage := NewPostgresStorage(conn)
	out, err := storage.GetReceipts(selection)
	require.NoError(t, err)

	fmt.Println(out)
}
