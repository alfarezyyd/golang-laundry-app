package helper

import (
	"database/sql"
)

func CommitOrRollback(tx *sql.Tx) {
	errorTransaction := recover()
	if errorTransaction == nil {
		errorCommit := tx.Commit()
		PanicIfError(errorCommit)
	} else {
		errorRollback := tx.Rollback()
		PanicIfError(errorRollback)
		panic(errorTransaction)
	}
}
