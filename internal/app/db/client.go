package db

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
)

type Repository interface {
}

type repository struct {
	db *sql.DB
	qb sq.StatementBuilderType
}

//func NewDbConnectClient(sqlConnect string) (Repository, error) {
//	bd, err := sql.Open("postgres", sqlConnect) //postgres
//	if err != nil {
//		return nil, err
//	}
//	return &repository{bd, sq.StatementBuilder.PlaceholderFormat(sq.Dollar)}, nil
//}
//
//func (r *repository) GetListData(conf config.Config, indent, startId uint64) ([]FileDataBase, error) {
//	listDb := make([]FileDataBase, 0, 1000)
//
//	queryGetUserTelegram, args, err := r.qb.
//		Select("id, path, created_at").
//		From("file").
//		Where(sq.Lt{"id": startId}).
//		OrderBy("id desc").
//		Offset(indent).
//		Limit(conf.PortionDb).
//		ToSql()
//
//	if err != nil {
//		return nil, err
//	}
//
//	rows, errDB := r.db.Query(queryGetUserTelegram, args...)
//	defer rows.Close()
//	if errDB != nil {
//		return nil, errDB
//	}
//
//	for rows.Next() {
//		recordDb := FileDataBase{}
//		errScan := rows.Scan(&recordDb.Id, &recordDb.File, &recordDb.Date)
//		if errScan != nil {
//			return nil, errScan
//		}
//		listDb = append(listDb, recordDb)
//
//	}
//
//	return listDb, nil
//}
