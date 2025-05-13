package repository

import (
	"fmt"
	"math"
	"reflect"
	"strings"

	"github.com/AlekSi/pointer"
	helpers "github.com/zercle/gofiber-helpers"
	queryInterface "github.com/zercle/kku-qf-services/pkg/common/models"
	"github.com/zercle/kku-qf-services/pkg/models"
	"gorm.io/gorm"
)

func (r commonRepository) GetList(
	queryTb queryInterface.StatementInterface,
	dest interface{},
	options *models.PaginationOptions,
	joinTB ...string) (err error) {
	db := r.MainDbConn
	if db == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}

	db = db.Model(queryTb)

	for _, tb := range joinTB {
		db.Preload(tb, func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at ASC")
		})
	}

	if options != nil && options.Search != nil && len(options.SearchFields) > 0 {
		for _, tb := range joinTB {
			db.Joins(tb)
		}
		searchString := "%" + strings.ToLower(pointer.GetString(options.Search)) + "%"
		args := make([]interface{}, len(options.SearchFields))
		for i := range options.SearchFields {
			args[i] = searchString
			options.SearchFields[i] = "LOWER(" + options.SearchFields[i] + ") LIKE ?"
		}
		statement := strings.Join(options.SearchFields, " OR ")
		db.Where("("+statement+")", args...)
	}

	db.Where(queryTb)

	if hasIDAttribute(queryTb) {
		db.Not(map[string]interface{}{queryTb.TableName() + ".id": uint(0)})
	}

	if options != nil {
		if options.Order != nil {
			db.Order(pointer.GetString(options.Order))
		}
		db.Count(options.Total)
		db.Scopes(Paginate(options))
	}

	if err := db.Find(dest).Error; err != nil {
		return err
	}
	return nil
}

func (r commonRepository) GetListWithNilSearch(
	queryTb queryInterface.StatementInterface,
	dest interface{},
	options *models.PaginationOptions,
	joinTB ...string) (err error) {
	db := r.MainDbConn
	if db == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}

	db = db.Model(queryTb)

	for _, tb := range joinTB {
		db.Preload(tb, func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at ASC")
		})
	}

	if options != nil && options.Search != nil && len(options.SearchFields) > 0 {
		for _, tb := range joinTB {
			db.Joins(tb)
		}
		searchString := "%" + strings.ToLower(pointer.GetString(options.Search)) + "%"
		args := make([]interface{}, len(options.SearchFields))
		for i := range options.SearchFields {
			args[i] = searchString
			options.SearchFields[i] = "LOWER(" + options.SearchFields[i] + ") LIKE ?"
		}
		statement := strings.Join(options.SearchFields, " OR ")
		db.Where("("+statement+")", args...)
	}

	if reflect.ValueOf(queryTb).Kind() == reflect.Ptr {
		interfaceValue := reflect.ValueOf(queryTb).Elem()
		interfaceType := reflect.Indirect(reflect.ValueOf(queryTb)).Type()
		for index := 0; index < interfaceValue.NumField(); index++ {
			interfaceSubValue := interfaceValue.Field(index)
			if !interfaceSubValue.IsZero() {
				if interfaceSubElement := interfaceSubValue.Elem(); interfaceSubElement.IsZero() {
					gormTag := interfaceType.Field(index).Tag.Get("gorm")
					db.Where(ExtractColumnName(gormTag), nil)
					interfaceSubValue.Set(reflect.Zero(interfaceSubValue.Type()))
				}
			}
		}
	}

	db.Where(queryTb)

	if hasIDAttribute(queryTb) {
		db.Not(map[string]interface{}{queryTb.TableName() + ".id": uint(0)})
	}

	if options != nil {
		if options.Order != nil {
			db.Order(pointer.GetString(options.Order))
		}
		db.Count(options.Total)
		db.Scopes(Paginate(options))
	}

	if err := db.Find(dest).Error; err != nil {
		return err
	}
	return nil
}

func (r commonRepository) GetFirst(
	tb interface{},
	joinTB ...string) (err error) {
	db := r.MainDbConn
	if db == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}

	db = db.Model(tb)
	for _, tb := range joinTB {
		db.Preload(tb, func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at ASC")
		})
	}

	db.Where(tb)

	if err := db.First(tb).Error; err != nil {
		return err
	}
	return nil
}

func (r commonRepository) GetFirstOrCreate(
	tb interface{},
	joinTB ...string) (err error) {
	db := r.MainDbConn
	if db == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}

	db = db.Model(tb)
	for _, tb := range joinTB {
		db.Preload(tb, func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at ASC")
		})
	}

	db.Where(tb)

	if err = db.FirstOrCreate(tb).Error; err != nil {
		return err
	}
	return nil
}

func Paginate(options *models.PaginationOptions) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if options.GetLimit() == 0 {
			options.TotalPage = pointer.ToInt(1)
			return db
		}
		options.DefaultPage(1)
		options.DefaultLimit(10)

		offset := (options.GetPage() - 1) * options.GetLimit()
		totalPage := int(math.Ceil(float64(options.GetTotal()) / float64(options.GetLimit())))
		options.TotalPage = &totalPage
		return db.Offset(offset).Limit(options.GetLimit()).Order(options.Order)
	}
}

func ExtractColumnName(str string) string {
	parts := strings.Split(str, ";")
	for _, part := range parts {
		if strings.HasPrefix(part, "column:") {
			return strings.TrimPrefix(part, "column:")
		}
	}
	return ""
}

func (r commonRepository) SelfReferenceJoin(times int, joinTB string) []string {
	var result []string
	for i := 1; i <= times; i++ {
		result = append(result, joinTB)
		joinTB += "." + joinTB
	}
	return result
}

func hasIDAttribute(a interface{}) bool {
	v := reflect.ValueOf(a)

	// ถ้าเป็น pointer ให้ดึงค่าออกมาก่อน
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Check ว่า struct มีฟิลด์ ID หรือไม่
	if v.Kind() == reflect.Struct {
		field := v.FieldByName("ID")
		return field.IsValid()
	}

	// Return false ถ้าไม่ใช่ struct หรือไม่มีฟิลด์ ID
	return false
}
