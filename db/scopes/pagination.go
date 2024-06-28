package scopes

import (
	"math"
	"net/http"
	"strconv"

	"demo.com/hello/core/http/utlis"
	"gorm.io/gorm"
)

func Paginate(model interface{}, metaData *utlis.Pagination, r *http.Request, tx *gorm.DB) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		var totalRows int64
		q := r.URL.Query()
		page, _ := strconv.Atoi(q.Get("page"))
		if page <= 0 {
			page = 1
		}
		pageSize, _ := strconv.Atoi(q.Get("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		metaData.Page = page
		metaData.PageSize = pageSize
		tx.Model(model).Count(&totalRows)
		metaData.TotalRows = totalRows
		totalPages := int(math.Ceil(float64(totalRows) / float64(metaData.PageSize)))
		metaData.TotalPages = totalPages
		if metaData.Page <= 0 {
			metaData.Page = 1
		}
		switch {
		case metaData.PageSize > 100:
			metaData.PageSize = 100
		case metaData.PageSize <= 0:
			metaData.PageSize = 10
		}
		offset := (metaData.Page - 1) * metaData.PageSize

		return db.Offset(offset).Limit(metaData.PageSize)
	}
}

//func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
//	return func(db *gorm.DB) *gorm.DB {
//		q := r.URL.Query()
//		page, _ := strconv.Atoi(q.Get("page"))
//		if page <= 0 {
//			page = 1
//		}
//
//		pageSize, _ := strconv.Atoi(q.Get("page_size"))
//		switch {
//		case pageSize > 100:
//			pageSize = 100
//		case pageSize <= 0:
//			pageSize = 10
//		}
//
//		offset := (page - 1) * pageSize
//		return db.Offset(offset).Limit(pageSize)
//	}
//}
//
