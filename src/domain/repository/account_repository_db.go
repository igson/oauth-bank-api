package repository

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/igson/banking/src/datasources/banking"
	"github.com/igson/banking/src/domain/models"
	"github.com/igson/banking/src/errors"
	"github.com/igson/banking/src/interfaces"
)

const (
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES (?,?,?,?,?,?);"
	queryGetAccountById         = "SELECT account_id, customer_id, opening_date, account_type, amount, status FROM accounts  WHERE account_id = ?"
	queryUpdateUser             = "UPDATE users set first_name=?, last_name=?, email=? WHERE id = ?"
	queryDeleteUser             = "DELETE from users WHERE id = ?"
	queryFindUser               = "select id, first_name, last_name, email, date_created, status FROM users WHERE status=?"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=? AND status=?"
)

//AccountRepository acesso ao repositório
type accountRepository struct {
	client *sqlx.DB
}

//NewAccountRepository acesso ao repositório
func NewAccountRepository(dbClient *sqlx.DB) interfaces.IAccountRepository {
	return &accountRepository{
		client: dbClient,
	}
}

func (repo *accountRepository) CriarConta(string) (*models.Account, *errors.RestErroAPI) {
	fmt.Println("Repository ----> Criar Conta")
	return nil, nil
}

func (repo *accountRepository) BuscarContaPorID(accountId int64) (*models.Account, *errors.RestErroAPI) {

	stmt, erro := banking.Conexao.Prepare(queryGetAccountById)

	if erro != nil {
		fmt.Println(erro)
		return nil, errors.NewInternalServerError("Database error")
	}

	defer stmt.Close()

	var account models.Account

	resultado := stmt.QueryRow(accountId)

	if sqlErro := resultado.Scan(&account.Id, &account.CustomerId, &account.OpeningDate, &account.AccountType, &account.Amount, &account.Status); sqlErro != nil {

		if sqlErro == sql.ErrNoRows {
			return nil, errors.NewNotFoundErro("Nenhuma conta encontrada com os dados informados.")
		}

		return nil, errors.NewInternalServerError("Erro parsing database response")

	}

	return &account, nil
}

func (repo *accountRepository) RegistrarTransacao(t models.Transaction) (*models.Transaction, *errors.RestErroAPI) {

	tx, err := repo.client.Begin()

	if err != nil {
		//	logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		fmt.Println(err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	// inserting bank account transaction
	result, _ := tx.Exec(`INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) 
											values (?, ?, ?, ?)`, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	// updating account balance
	if t.IsWithdrawal() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - ? where account_id = ?`, t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + ? where account_id = ?`, t.Amount, t.AccountId)
	}

	// in case of error Rollback, and changes from both the tables will be reverted
	if err != nil {
		fmt.Println("BBBBBBBBBBBBBBBBBBBBB")
		tx.Rollback()
		//logger.Error("Error while saving transaction: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}
	// commit the transaction when all is good
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Println("CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC")
		//logger.Error("Error while commiting transaction for bank account: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}
	// getting the last transaction ID from the transaction table
	transactionId, err := result.LastInsertId()

	if err != nil {
		//logger.Error("Error while getting the last transaction id: " + err.Error())
		fmt.Println("DDDDDDDDDDDDDDDDDDDDDDDDDDDDDD")
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	// Getting the latest account information from the accounts table
	account, appErr := repo.BuscarContaPorID(t.AccountId)
	if appErr != nil {
		fmt.Println("EEEEEEEEEEEEEEEEEEEEEEEEEEEE")
		return nil, appErr
	}

	t.Id = transactionId

	// updating the transaction struct with the latest balance

	t.Amount = account.Amount

	return &t, nil

}
