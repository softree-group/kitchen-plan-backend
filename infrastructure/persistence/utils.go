package persistence

import (
	"github.com/jackc/pgx"
	"github.com/softree-group/kitchen-plan-backend/domain/entity"
)

func endTx(tx *pgx.Tx, err error) {
	if tx.Status() == -1 {
		return
	}
	if err != nil {
		_ = tx.Rollback()
		return
	}
	_ = tx.Commit()
}

func toEntityError(err error) error {
	switch true {
	case err.Error() == "no rows in result set":
		return entity.ErrNotFound
	default:
		return err
	}
}
