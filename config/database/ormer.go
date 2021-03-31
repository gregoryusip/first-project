package database

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type Ormer interface {
	orm.DB
	Begin() (*pg.Tx, error)
	BeginTxn() (TransactionOrmer, error)
}

type TransactionOrmer interface {
	Ormer
	Commit() error
	Rollback() error
}

type ormer struct {
	*pg.DB
}

func NewOrmer(db *pg.DB) Ormer {
	return &ormer{
		db,
	}
}

func (o *ormer) BeginTxn() (TransactionOrmer, error) {
	txn, err := o.Begin()
	if err != nil {
		return nil, err
	}
	return NewTransactionOrmer(txn), nil
}

type transactionOrmer struct {
	*pg.Tx
}

func NewTransactionOrmer(tx *pg.Tx) TransactionOrmer {
	return &transactionOrmer{
		tx,
	}
}

func (o *transactionOrmer) BeginTxn() (TransactionOrmer, error) {
	return o, nil
}
