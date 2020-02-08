package migrations

import "github.com/jmoiron/sqlx"

const createChannelTable = "CREATE TABLE `channels` (`name` varchar(255) NOT NULL, `bot_name` varchar(255) NOT NULL, `notes` text NULL, `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP , `deleted_at` datetime NULL, PRIMARY KEY(`name`))"

func CreateChannelTable(tx *sqlx.Tx) error {
	_, err := tx.Exec(createChannelTable)
	return err
}
