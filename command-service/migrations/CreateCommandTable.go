package migrations

import (
	"github.com/jmoiron/sqlx"
)

func CreateCommandTable(tx *sqlx.Tx) error {
	_, err := tx.Exec("CREATE TABLE `commands` (`channel` varchar(255),`name` varchar(255),`response` varchar(255) NOT NULL,`enabled` tinyint UNSIGNED DEFAULT 1 NOT NULL,`restricted` tinyint UNSIGNED DEFAULT 0 NOT NULL,`cooldown` tinyint UNSIGNED DEFAULT 0 NOT NULL,`description` varchar(255),`schedule` int UNSIGNED DEFAULT 0 NOT NULL,`updated_by` varchar(255) NOT NULL,`updated_at` datetime DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY(`channel`,`name`))")
	return err
}
