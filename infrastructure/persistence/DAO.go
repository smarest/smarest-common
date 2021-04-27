package persistence

import (
	"database/sql"

	"gopkg.in/gorp.v3"
)

type DAOImpl struct {
	Table       string
	DbMap       *gorp.DbMap
	Transaction *gorp.Transaction
}

func NewDAOImpl(table string, dbMap *gorp.DbMap) *DAOImpl {
	return &DAOImpl{Table: table, DbMap: dbMap}
}

func (dao *DAOImpl) FindByID(id int64, holder interface{}) error {
	return dao.DbMap.SelectOne(holder, "SELECT * FROM "+dao.Table+" WHERE id=?", id)
}

func (dao *DAOImpl) FindAll(holder interface{}) ([]interface{}, error) {
	return dao.DbMap.Select(holder, "SELECT * FROM "+dao.Table)

}

func (dao *DAOImpl) Select(holder interface{}, sql string, args ...interface{}) ([]interface{}, error) {
	return dao.DbMap.Select(holder, sql, args...)
}

func (dao *DAOImpl) SelectOnce(holder interface{}, sql string, args ...interface{}) error {
	return dao.DbMap.SelectOne(holder, sql, args...)
}

func (dao *DAOImpl) Exec(queryString string, args ...interface{}) (sql.Result, error) {
	if dao.Transaction != nil {
		return dao.Transaction.Exec(queryString, args...)
	} else {
		return dao.DbMap.Exec(queryString, args...)
	}
}

func (dao *DAOImpl) InsertBySQL(queryString string, args ...interface{}) (int64, error) {
	result, err := dao.Exec(queryString, args...)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (dao *DAOImpl) UpdateBySQL(queryString string, args ...interface{}) (int64, error) {
	result, err := dao.Exec(queryString, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (dao *DAOImpl) DeleteBySQL(queryString string, args ...interface{}) (int64, error) {
	var result sql.Result
	var err error
	if dao.Transaction != nil {
		result, err = dao.Transaction.Exec(queryString, args...)
	} else {
		result, err = dao.Exec(queryString, args...)
	}
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (dao *DAOImpl) Insert(args ...interface{}) error {
	if dao.Transaction != nil {
		return dao.Transaction.Insert(args...)
	}
	return dao.DbMap.Insert(args...)
}
func (dao *DAOImpl) Update(args ...interface{}) (int64, error) {
	if dao.Transaction != nil {
		return dao.DbMap.Update(args...)
	}
	return dao.DbMap.Update(args...)
}
func (dao *DAOImpl) UpdateColumns(colFilter gorp.ColumnFilter, args ...interface{}) (int64, error) {
	if dao.Transaction != nil {
		return dao.DbMap.UpdateColumns(colFilter, args...)
	}
	return dao.DbMap.UpdateColumns(colFilter, args...)
}
func (dao *DAOImpl) Delete(args ...interface{}) (int64, error) {
	if dao.Transaction != nil {
		return dao.DbMap.Delete(args...)
	}
	return dao.DbMap.Delete(args...)
}

func (dao *DAOImpl) Commit() error {
	err := dao.Transaction.Commit()
	dao.Transaction = nil
	return err
}
func (dao *DAOImpl) Rollback() error {
	err := dao.Transaction.Rollback()
	dao.Transaction = nil
	return err
}

func (dao *DAOImpl) Begin() error {
	var err error
	dao.Transaction, err = dao.DbMap.Begin()
	return err
}
