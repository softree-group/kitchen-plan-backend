package persistance

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/softree-group/kitchen-plan-backend/domain/entity"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPostgresStorage_GetReceipts(t *testing.T) {
	conn, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer func() {
		err = conn.Close()
		require.NoError(t, err)
	}()

	rows := sqlmock.NewRows([]string{"id", "image", "title", "time", "type"})
	expected := []entity.Receipt{
		{Id: 1, Image: "image", Title: "title", TimeToCook: 2, Type: "type"},
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
	mock.ExpectQuery("select id, image, title, time, type from recipes").
		WithArgs().
		WillReturnRows(rows)
	mock.ExpectClose()

	storage := NewPostgresStorage(conn)
	out, err := storage.GetReceipts(selection)
	require.NoError(t, err)

	fmt.Println(out)
}
