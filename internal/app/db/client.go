package db

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/Sergei3232/Getting-image-data/internal/app/datastruct"
	_ "github.com/lib/pq"
)

type Repository interface {
	GettingIdImageFileStorage(mapIdItems map[int64]datastruct.DataSCV)
	GetImageHightWidth(mapIdItems map[int64]datastruct.ImageFileCSV) error
}

type repository struct {
	db *sql.DB
	qb sq.StatementBuilderType
}

func NewDbConnectClient(sqlConnect string) (Repository, error) {
	bd, err := sql.Open("postgres", sqlConnect) //postgres
	if err != nil {
		return nil, err
	}
	return &repository{bd, sq.StatementBuilder.PlaceholderFormat(sq.Dollar)}, nil
}
