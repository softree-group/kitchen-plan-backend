package persistance

import (
	"github.com/jackc/pgx"
	"github.com/softree-group/kitchen-plan-backend/domain/entity"
)

func chooseSql(selection entity.Selection) (string, []interface{}) {
	switch {
	case selection.Type != "":
		return chooseByType(selection)
	case selection.Title != "":
		return chooseByTitle(selection)
	case len(selection.Ingredients) != 0:
		return chooseByIngredients(selection)
	}

	args := pgx.QueryArgs{}
	if selection.Limit != 0 {
		args.Append(selection.Limit)
		return allLimitReceipts, args
	}

	return allReceipts, args
}

func chooseByType(selection entity.Selection) (string, []interface{}) {
	var args []interface{}
	args = append(args, selection.Type)

	if selection.Title != "" {
		args.Append(selection.Title)

		if selection.Limit != 0 {
			args.Append(selection.Limit)
			return typeTitleLimitReceipts, args
		}
		return typeTitleReceipts, args
	}

	if selection.Limit != 0 {
		args.Append(selection.Limit)
		return typeLimitReceipts, args
	}

	return typeReceipts, args
}

func chooseByTitle(selection entity.Selection) (string, pgx.QueryArgs) {
	args := pgx.QueryArgs{}
	args.Append(selection.Title)

	if selection.Limit != 0 {
		args.Append(selection.Limit)
		return titleLimitReceipts, args
	}

	return titleReceipts, args
}

func chooseByIngredients(selection entity.Selection) (string, pgx.QueryArgs) {
	args := pgx.QueryArgs{}
	args.Append(selection.Ingredients)

	if selection.Type != "" {
		args.Append(selection.Type)

		if selection.Limit != 0 {
			args.Append(selection.Limit)
			return ingredientsTypeLimitReceipts, args
		}
		return ingredientsTypeReceipts, args
	}

	if selection.Limit != 0 {
		args.Append(selection.Limit)
		return ingredientsLimitReceipts, args
	}

	return ingredientsReceipts, args
}
