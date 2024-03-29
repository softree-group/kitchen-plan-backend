package persistence

import (
	"fmt"
	"github.com/softree-group/kitchen-plan-backend/src/domain/entity"
	"strconv"
	"strings"
)

func genSQLFilter(filter *entity.ReceiptFilter) (string, []interface{}) {
	var args []interface{}

	sql := fmt.Sprintf("select id, title, image, time_to_cook from recipes where id > %d", filter.Since)

	if filter.Type != "" {
		sql += fmt.Sprintf(" and type = $1")
		args = append(args, filter.Type)
	}

	if filter.Title != "" {
		searchTitle := strings.ToLower(filter.Title) + ":*|" + strings.Title(filter.Title) + ":*"
		args = append(args, searchTitle)

		sql += fmt.Sprintf(" and to_tsvector(title) @@ to_tsquery('russian',$%d)", len(args))
	}

	if len(filter.ForRecipes) > 0 {
		strIds := joinIntArr(filter.ForRecipes)
		sql += fmt.Sprintf(" and id in (%s)" , strIds)
	}

	if len(filter.Ingredients) > 0 {
		strIds := joinIntArr(filter.Ingredients)
		sql += fmt.Sprintf(" and ARRAY[%s] || ARRAY(select distinct parent_ingredient_id from" +
			" overall_ingredients where ingredient_id in (%s)) @>"+
			" ARRAY(select ingredient_id from recipes_ingredients where receipt_id = recipes.id)",
			strIds, strIds)
	}

	sql += " order by id"

	if filter.Limit == 0 {
		filter.Limit = 50
	}

	sql += fmt.Sprintf(" limit %d", filter.Limit)
	return sql, args
}

func joinIntArr(ids []int) string {
	var arr []string
	for _, id := range ids {
		arr = append(arr, strconv.Itoa(id))
	}
	return strings.Join(arr, ",")
}
