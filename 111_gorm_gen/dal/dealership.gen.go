// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"learn/111_gorm_gen/model"
)

func newDealership(db *gorm.DB, opts ...gen.DOOption) dealership {
	_dealership := dealership{}

	_dealership.dealershipDo.UseDB(db, opts...)
	_dealership.dealershipDo.UseModel(&model.Dealership{})

	tableName := _dealership.dealershipDo.TableName()
	_dealership.ALL = field.NewAsterisk(tableName)
	_dealership.ID = field.NewInt64(tableName, "id")
	_dealership.Name = field.NewString(tableName, "name")
	_dealership.Area = field.NewString(tableName, "area")
	_dealership.Province = field.NewString(tableName, "province")
	_dealership.City = field.NewString(tableName, "city")
	_dealership.Address = field.NewString(tableName, "address")
	_dealership.Longitude = field.NewFloat64(tableName, "longitude")
	_dealership.Latitude = field.NewFloat64(tableName, "latitude")
	_dealership.Phone = field.NewString(tableName, "phone")
	_dealership.Status = field.NewString(tableName, "status")
	_dealership.LastUpdatedUsername = field.NewString(tableName, "last_updated_username")
	_dealership.LastUpdatedAt = field.NewTime(tableName, "last_updated_at")

	_dealership.fillFieldMap()

	return _dealership
}

type dealership struct {
	dealershipDo

	ALL                 field.Asterisk
	ID                  field.Int64
	Name                field.String  // 经销商名称
	Area                field.String  // 所处地区
	Province            field.String  // 所在省份
	City                field.String  // 所在城市
	Address             field.String  // 详细地址
	Longitude           field.Float64 // 经度
	Latitude            field.Float64 // 纬度
	Phone               field.String  // 联系人电话
	Status              field.String  // 该店状态
	LastUpdatedUsername field.String  // 最近操作人
	LastUpdatedAt       field.Time    // 最近操作时间

	fieldMap map[string]field.Expr
}

func (d dealership) Table(newTableName string) *dealership {
	d.dealershipDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d dealership) As(alias string) *dealership {
	d.dealershipDo.DO = *(d.dealershipDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *dealership) updateTableName(table string) *dealership {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewInt64(table, "id")
	d.Name = field.NewString(table, "name")
	d.Area = field.NewString(table, "area")
	d.Province = field.NewString(table, "province")
	d.City = field.NewString(table, "city")
	d.Address = field.NewString(table, "address")
	d.Longitude = field.NewFloat64(table, "longitude")
	d.Latitude = field.NewFloat64(table, "latitude")
	d.Phone = field.NewString(table, "phone")
	d.Status = field.NewString(table, "status")
	d.LastUpdatedUsername = field.NewString(table, "last_updated_username")
	d.LastUpdatedAt = field.NewTime(table, "last_updated_at")

	d.fillFieldMap()

	return d
}

func (d *dealership) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *dealership) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 12)
	d.fieldMap["id"] = d.ID
	d.fieldMap["name"] = d.Name
	d.fieldMap["area"] = d.Area
	d.fieldMap["province"] = d.Province
	d.fieldMap["city"] = d.City
	d.fieldMap["address"] = d.Address
	d.fieldMap["longitude"] = d.Longitude
	d.fieldMap["latitude"] = d.Latitude
	d.fieldMap["phone"] = d.Phone
	d.fieldMap["status"] = d.Status
	d.fieldMap["last_updated_username"] = d.LastUpdatedUsername
	d.fieldMap["last_updated_at"] = d.LastUpdatedAt
}

func (d dealership) clone(db *gorm.DB) dealership {
	d.dealershipDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d dealership) replaceDB(db *gorm.DB) dealership {
	d.dealershipDo.ReplaceDB(db)
	return d
}

type dealershipDo struct{ gen.DO }

type IDealershipDo interface {
	gen.SubQuery
	Debug() IDealershipDo
	WithContext(ctx context.Context) IDealershipDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDealershipDo
	WriteDB() IDealershipDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDealershipDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDealershipDo
	Not(conds ...gen.Condition) IDealershipDo
	Or(conds ...gen.Condition) IDealershipDo
	Select(conds ...field.Expr) IDealershipDo
	Where(conds ...gen.Condition) IDealershipDo
	Order(conds ...field.Expr) IDealershipDo
	Distinct(cols ...field.Expr) IDealershipDo
	Omit(cols ...field.Expr) IDealershipDo
	Join(table schema.Tabler, on ...field.Expr) IDealershipDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDealershipDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDealershipDo
	Group(cols ...field.Expr) IDealershipDo
	Having(conds ...gen.Condition) IDealershipDo
	Limit(limit int) IDealershipDo
	Offset(offset int) IDealershipDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDealershipDo
	Unscoped() IDealershipDo
	Create(values ...*model.Dealership) error
	CreateInBatches(values []*model.Dealership, batchSize int) error
	Save(values ...*model.Dealership) error
	First() (*model.Dealership, error)
	Take() (*model.Dealership, error)
	Last() (*model.Dealership, error)
	Find() ([]*model.Dealership, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Dealership, err error)
	FindInBatches(result *[]*model.Dealership, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Dealership) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDealershipDo
	Assign(attrs ...field.AssignExpr) IDealershipDo
	Joins(fields ...field.RelationField) IDealershipDo
	Preload(fields ...field.RelationField) IDealershipDo
	FirstOrInit() (*model.Dealership, error)
	FirstOrCreate() (*model.Dealership, error)
	FindByPage(offset int, limit int) (result []*model.Dealership, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDealershipDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d dealershipDo) Debug() IDealershipDo {
	return d.withDO(d.DO.Debug())
}

func (d dealershipDo) WithContext(ctx context.Context) IDealershipDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d dealershipDo) ReadDB() IDealershipDo {
	return d.Clauses(dbresolver.Read)
}

func (d dealershipDo) WriteDB() IDealershipDo {
	return d.Clauses(dbresolver.Write)
}

func (d dealershipDo) Session(config *gorm.Session) IDealershipDo {
	return d.withDO(d.DO.Session(config))
}

func (d dealershipDo) Clauses(conds ...clause.Expression) IDealershipDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d dealershipDo) Returning(value interface{}, columns ...string) IDealershipDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d dealershipDo) Not(conds ...gen.Condition) IDealershipDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d dealershipDo) Or(conds ...gen.Condition) IDealershipDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d dealershipDo) Select(conds ...field.Expr) IDealershipDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d dealershipDo) Where(conds ...gen.Condition) IDealershipDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d dealershipDo) Order(conds ...field.Expr) IDealershipDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d dealershipDo) Distinct(cols ...field.Expr) IDealershipDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d dealershipDo) Omit(cols ...field.Expr) IDealershipDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d dealershipDo) Join(table schema.Tabler, on ...field.Expr) IDealershipDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d dealershipDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDealershipDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d dealershipDo) RightJoin(table schema.Tabler, on ...field.Expr) IDealershipDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d dealershipDo) Group(cols ...field.Expr) IDealershipDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d dealershipDo) Having(conds ...gen.Condition) IDealershipDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d dealershipDo) Limit(limit int) IDealershipDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d dealershipDo) Offset(offset int) IDealershipDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d dealershipDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDealershipDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d dealershipDo) Unscoped() IDealershipDo {
	return d.withDO(d.DO.Unscoped())
}

func (d dealershipDo) Create(values ...*model.Dealership) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d dealershipDo) CreateInBatches(values []*model.Dealership, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d dealershipDo) Save(values ...*model.Dealership) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d dealershipDo) First() (*model.Dealership, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Dealership), nil
	}
}

func (d dealershipDo) Take() (*model.Dealership, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Dealership), nil
	}
}

func (d dealershipDo) Last() (*model.Dealership, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Dealership), nil
	}
}

func (d dealershipDo) Find() ([]*model.Dealership, error) {
	result, err := d.DO.Find()
	return result.([]*model.Dealership), err
}

func (d dealershipDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Dealership, err error) {
	buf := make([]*model.Dealership, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d dealershipDo) FindInBatches(result *[]*model.Dealership, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d dealershipDo) Attrs(attrs ...field.AssignExpr) IDealershipDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d dealershipDo) Assign(attrs ...field.AssignExpr) IDealershipDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d dealershipDo) Joins(fields ...field.RelationField) IDealershipDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d dealershipDo) Preload(fields ...field.RelationField) IDealershipDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d dealershipDo) FirstOrInit() (*model.Dealership, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Dealership), nil
	}
}

func (d dealershipDo) FirstOrCreate() (*model.Dealership, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Dealership), nil
	}
}

func (d dealershipDo) FindByPage(offset int, limit int) (result []*model.Dealership, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d dealershipDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d dealershipDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d dealershipDo) Delete(models ...*model.Dealership) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *dealershipDo) withDO(do gen.Dao) *dealershipDo {
	d.DO = *do.(*gen.DO)
	return d
}
