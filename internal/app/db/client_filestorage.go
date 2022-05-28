package db

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/Sergei3232/Getting-image-data/internal/app/datastruct"
	"log"
)

func (r *repository) GetImageHightWidth(mapIdItems map[int64]datastruct.DataSCV) error {
	for key, val := range mapIdItems {
		sqlStatement, args, err := r.qb.Select("height,width").
			From("image").
			Where(sq.Eq{"id": val.IdFileStorage}).
			ToSql()
		if err != nil {
			log.Println(err)
			return err
		}

		rows, errDB := r.db.Query(sqlStatement, args...)
		if errDB != nil {
			log.Println(errDB)
		}

		for rows.Next() {
			var height, width int64
			if err := rows.Scan(&height, &width); err != nil {
				log.Println(err)
			}
			val.Width, val.Height = width, height
			mapIdItems[key] = val
		}
	}
	return nil
}
