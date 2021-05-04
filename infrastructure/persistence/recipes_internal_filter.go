package persistence

import (
	"fmt"
	"github.com/softree-group/kitchen-plan-backend/domain/entity"
	"strconv"
	"strings"
)

func genSQLFilter(filter *entity.ReceiptFilter) string {
	sql := fmt.Sprintf("select id, title, image, time_to_cook from recipes where id > %d", filter.Since)
	if filter.Type != "" {
		sql += fmt.Sprintf(" and type = '%s'", filter.Type)
	}

	if filter.Title != "" {
		searchTitle := strings.ToLower(filter.Title) + ":*|" + strings.Title(filter.Title) + ":*"
		sql += fmt.Sprintf(" and and to_tsvector(title) @@ to_tsquery('russian','%s')", searchTitle)
	}

	if len(filter.Ingredients) > 0 {
		sql += fmt.Sprintf(" and ARRAY[%s] @>"+
			" ARRAY(select ingredient_id from recipes_ingredients where receipt_id = id)",
			joinIntArr(filter.Ingredients))
	}

	sql += " order by id"

	var limit int
	if filter.Limit != 0 {
		limit = filter.Limit
	}

	sql += fmt.Sprintf(" limit %d", limit)
	return sql
}

func joinIntArr(ids []int) string {
	var arr []string
	for id := range ids {
		arr = append(arr, strconv.Itoa(id))
	}
	return strings.Join(arr, ",")
}
