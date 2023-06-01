package services

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/ThaliaAC/labora-wallet/models"
)

func (p *PostgresDBHandler) CreateWalletInTx(wallet models.Wallet, tx *sql.Tx) (models.Wallet, error) {
	//Validation
	if wallet.National_id == "" || wallet.Country == "" {
		log.Fatal("Empty fields")
	}

	query := `INSERT INTO walletsTable (national_id, country, request_date, balance)
	VALUES ($1, $2, $3, $4) RETURNING *`
	row := tx.QueryRow(query, &wallet.National_id, &wallet.Country, time.Now(), 100)

	err := row.Scan(&wallet.ID, &wallet.National_id, &wallet.Country, &wallet.RequestDate, &wallet.Balance)
	if err != nil {
		return models.Wallet{}, fmt.Errorf("Error creating the wallet in the transaction: %w", err)
	}

	return wallet, nil
}

func (p *PostgresDBHandler) CreateWallet(wallet models.Wallet, log models.Log) (models.Wallet, error) {
	// Implementar la lógica para crear un artículo en la base de datos PostgreSQL
	// Start a transaction
	tx, err := p.Begin()
	if err != nil {
		tx.Rollback()

		return models.Wallet{}, fmt.Errorf("Error at the beginning of the transaction: %w", err)
	}

	wallet, err = p.CreateWalletInTx(wallet, tx)
	if err != nil {
		tx.Rollback()

		return models.Wallet{}, fmt.Errorf("Error trying to create the wallet in the transaction: %w", err)
	}

	err = p.CreateLogInTx(wallet.National_id, wallet.Country, "Approved", "CREATE WALLET", tx)
	if err != nil {
		tx.Rollback()

		return models.Wallet{}, fmt.Errorf("Error trying to create the log in the transaction: %w", err)
	}

	// Commit the transaction if no errors occur
	err = tx.Commit()
	if err != nil {
		tx.Rollback()

		return models.Wallet{}, fmt.Errorf("Error committing the transaction: %w", err)
	}

	return wallet, nil
}

func (p *PostgresDBHandler) UpdateWallet(id int, wallet models.Wallet) (models.Wallet, error) {
	// Implementar la lógica para obtener un artículo de la base de datos PostgreSQL
	return models.Wallet{}, nil
}
func (Db *PostgresDBHandler) searchWalletByIdInTx(id int, tx *sql.Tx) (models.Wallet, error) {
	var wallet models.Wallet
	query := "SELECT * FROM walletsTable WHERE id=$1"

	err := tx.QueryRow(query, id).Scan(&wallet.ID, &wallet.National_id, &wallet.Country, &wallet.RequestDate, &wallet.Balance)
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return models.Wallet{}, fmt.Errorf("Error querying database for wallet: %w", err)
	}

	return wallet, nil
}

func (p *PostgresDBHandler) DeleteWallet(id int, log models.Log) error {
	// Implementar la lógica para actualizar un artículo en la base de datos PostgreSQL
	// Start a transaction
	tx, err := p.Begin()
	if err != nil {
		tx.Rollback()

		return fmt.Errorf("Error at the beginning of the transaction: %w", err)
	}

	wallet, err := p.searchWalletByIdInTx(id, tx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Error trying to search the wallet in the transaction: %w", err)
	}

	National_id, Status, Country, RequestType, err := p.DeleteWalletInTx(wallet, tx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Error trying to delete the wallet in the transaction: %w", err)
	}

	err = p.CreateLogInTx(National_id, Status, Country, RequestType, RequestDate, tx)
	if err != nil {
		tx.Rollback()

		return fmt.Errorf("Error trying to create the log in the transaction: %w", err)
	}

	// Commit the transaction if no errors occur
	err = tx.Commit()
	if err != nil {
		tx.Rollback()

		return fmt.Errorf("Error committing the transaction: %w", err)
	}

	return nil
}

func (p *PostgresDBHandler) DeleteWalletInTx(wallet models.Wallet, tx *sql.Tx) (string, string, string, string, error) {
	query := "DELETE FROM walletsTable WHERE id = $1"

	_, err := tx.Exec(query, wallet.ID)
	if err != nil {
		tx.Rollback()
		return "", "", "", "", fmt.Errorf("error executing delete query: %w", err)
	}

	National_id := wallet.National_id
	Country := wallet.Country
	Status := "Deleted"
	request_type := "DELETE WALLET"

	return National_id, Country, Status, request_type, nil
}

func (p *PostgresDBHandler) WalletStatus(id int) (string, error) {
	// Implementar la lógica para eliminar un artículo de la base de datos PostgreSQL
	return "", nil
}
func (p *PostgresDBHandler) CreateLog(log models.Log) error {
	// Implementar la lógica para crear un artículo de la base de datos PostgreSQL
	return nil
}

func (p *PostgresDBHandler) CreateLogInTx(National_id, Country, Status, RequestType string, RequestDate time.Time, tx *sql.Tx) error {
	var log models.Log
	//Validation
	if National_id == "" || Country == "" || Status == "" || RequestType == "" {
		fmt.Println("empty fields")
	}

	// Insert the new log in the database
	query := `INSERT INTO logsTable (national_id, country, status, request_date, request_type)
                        VALUES ($1, $2, $3, $4, $5) RETURNING *`
	row := tx.QueryRow(query, National_id, Status, Country, time.Now(), RequestType)

	err := row.Scan(&log.National_id, &log.Country, &log.Status, &log.RequestDate, &log.RequestType)
	if err != nil {

		return fmt.Errorf("Error creating the log: %w", err)
	}

	return nil
}

func (p *PostgresDBHandler) GetLogs(log models.Log) (models.Log, error) {
	// Implementar la lógica para eliminar un artículo de la base de datos PostgreSQL
	return models.Log{}, nil
}
