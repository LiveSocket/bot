package migrations

import (
	"github.com/jmoiron/sqlx"
)

func CreateModTable(tx *sqlx.Tx) error {
	tx.MustExec("CREATE TABLE `mods` (`channel` varchar(255),`username` varchar(255),`created_at` datetime DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY(`channel`,`username`))")
	return nil
}
