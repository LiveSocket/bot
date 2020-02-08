package migrations

import (
	"github.com/jmoiron/sqlx"
)

func CreateModChatTable(tx *sqlx.Tx) error {
	tx.MustExec("CREATE TABLE `mod_chat` (`id` INT UNSIGNED NOT NULL AUTO_INCREMENT,`channel` varchar(255) NOT NULL,`message` text NOT NULL,`name` varchar(255) NOT NULL,`timestamp` datetime DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY(`id`))")
	return nil
}
