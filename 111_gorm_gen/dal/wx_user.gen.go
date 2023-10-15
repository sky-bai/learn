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

func newWxUser(db *gorm.DB, opts ...gen.DOOption) wxUser {
	_wxUser := wxUser{}

	_wxUser.wxUserDo.UseDB(db, opts...)
	_wxUser.wxUserDo.UseModel(&model.WxUser{})

	tableName := _wxUser.wxUserDo.TableName()
	_wxUser.ALL = field.NewAsterisk(tableName)
	_wxUser.ID = field.NewInt64(tableName, "id")
	_wxUser.CreatedAt = field.NewTime(tableName, "created_at")
	_wxUser.UpdatedAt = field.NewTime(tableName, "updated_at")
	_wxUser.DeletedAt = field.NewField(tableName, "deleted_at")
	_wxUser.OpenID = field.NewString(tableName, "open_id")
	_wxUser.Customer = field.NewString(tableName, "customer")
	_wxUser.FollowStatus = field.NewInt32(tableName, "follow_status")
	_wxUser.HeadURL = field.NewString(tableName, "head_url")
	_wxUser.LastLoginTime = field.NewTime(tableName, "last_login_time")
	_wxUser.Nickname = field.NewString(tableName, "nickname")
	_wxUser.Sex = field.NewInt32(tableName, "sex")
	_wxUser.UnionID = field.NewString(tableName, "union_id")

	_wxUser.fillFieldMap()

	return _wxUser
}

type wxUser struct {
	wxUserDo

	ALL           field.Asterisk
	ID            field.Int64 // 主键Id
	CreatedAt     field.Time  // 创建时间
	UpdatedAt     field.Time  // 修改时间
	DeletedAt     field.Field // 删除时间
	OpenID        field.String
	Customer      field.String
	FollowStatus  field.Int32
	HeadURL       field.String
	LastLoginTime field.Time
	Nickname      field.String
	Sex           field.Int32
	UnionID       field.String

	fieldMap map[string]field.Expr
}

func (w wxUser) Table(newTableName string) *wxUser {
	w.wxUserDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w wxUser) As(alias string) *wxUser {
	w.wxUserDo.DO = *(w.wxUserDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *wxUser) updateTableName(table string) *wxUser {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.CreatedAt = field.NewTime(table, "created_at")
	w.UpdatedAt = field.NewTime(table, "updated_at")
	w.DeletedAt = field.NewField(table, "deleted_at")
	w.OpenID = field.NewString(table, "open_id")
	w.Customer = field.NewString(table, "customer")
	w.FollowStatus = field.NewInt32(table, "follow_status")
	w.HeadURL = field.NewString(table, "head_url")
	w.LastLoginTime = field.NewTime(table, "last_login_time")
	w.Nickname = field.NewString(table, "nickname")
	w.Sex = field.NewInt32(table, "sex")
	w.UnionID = field.NewString(table, "union_id")

	w.fillFieldMap()

	return w
}

func (w *wxUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *wxUser) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 12)
	w.fieldMap["id"] = w.ID
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["deleted_at"] = w.DeletedAt
	w.fieldMap["open_id"] = w.OpenID
	w.fieldMap["customer"] = w.Customer
	w.fieldMap["follow_status"] = w.FollowStatus
	w.fieldMap["head_url"] = w.HeadURL
	w.fieldMap["last_login_time"] = w.LastLoginTime
	w.fieldMap["nickname"] = w.Nickname
	w.fieldMap["sex"] = w.Sex
	w.fieldMap["union_id"] = w.UnionID
}

func (w wxUser) clone(db *gorm.DB) wxUser {
	w.wxUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w wxUser) replaceDB(db *gorm.DB) wxUser {
	w.wxUserDo.ReplaceDB(db)
	return w
}

type wxUserDo struct{ gen.DO }

type IWxUserDo interface {
	gen.SubQuery
	Debug() IWxUserDo
	WithContext(ctx context.Context) IWxUserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWxUserDo
	WriteDB() IWxUserDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWxUserDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWxUserDo
	Not(conds ...gen.Condition) IWxUserDo
	Or(conds ...gen.Condition) IWxUserDo
	Select(conds ...field.Expr) IWxUserDo
	Where(conds ...gen.Condition) IWxUserDo
	Order(conds ...field.Expr) IWxUserDo
	Distinct(cols ...field.Expr) IWxUserDo
	Omit(cols ...field.Expr) IWxUserDo
	Join(table schema.Tabler, on ...field.Expr) IWxUserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWxUserDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWxUserDo
	Group(cols ...field.Expr) IWxUserDo
	Having(conds ...gen.Condition) IWxUserDo
	Limit(limit int) IWxUserDo
	Offset(offset int) IWxUserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWxUserDo
	Unscoped() IWxUserDo
	Create(values ...*model.WxUser) error
	CreateInBatches(values []*model.WxUser, batchSize int) error
	Save(values ...*model.WxUser) error
	First() (*model.WxUser, error)
	Take() (*model.WxUser, error)
	Last() (*model.WxUser, error)
	Find() ([]*model.WxUser, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WxUser, err error)
	FindInBatches(result *[]*model.WxUser, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WxUser) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWxUserDo
	Assign(attrs ...field.AssignExpr) IWxUserDo
	Joins(fields ...field.RelationField) IWxUserDo
	Preload(fields ...field.RelationField) IWxUserDo
	FirstOrInit() (*model.WxUser, error)
	FirstOrCreate() (*model.WxUser, error)
	FindByPage(offset int, limit int) (result []*model.WxUser, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWxUserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w wxUserDo) Debug() IWxUserDo {
	return w.withDO(w.DO.Debug())
}

func (w wxUserDo) WithContext(ctx context.Context) IWxUserDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w wxUserDo) ReadDB() IWxUserDo {
	return w.Clauses(dbresolver.Read)
}

func (w wxUserDo) WriteDB() IWxUserDo {
	return w.Clauses(dbresolver.Write)
}

func (w wxUserDo) Session(config *gorm.Session) IWxUserDo {
	return w.withDO(w.DO.Session(config))
}

func (w wxUserDo) Clauses(conds ...clause.Expression) IWxUserDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w wxUserDo) Returning(value interface{}, columns ...string) IWxUserDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w wxUserDo) Not(conds ...gen.Condition) IWxUserDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w wxUserDo) Or(conds ...gen.Condition) IWxUserDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w wxUserDo) Select(conds ...field.Expr) IWxUserDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w wxUserDo) Where(conds ...gen.Condition) IWxUserDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w wxUserDo) Order(conds ...field.Expr) IWxUserDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w wxUserDo) Distinct(cols ...field.Expr) IWxUserDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w wxUserDo) Omit(cols ...field.Expr) IWxUserDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w wxUserDo) Join(table schema.Tabler, on ...field.Expr) IWxUserDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w wxUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWxUserDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w wxUserDo) RightJoin(table schema.Tabler, on ...field.Expr) IWxUserDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w wxUserDo) Group(cols ...field.Expr) IWxUserDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w wxUserDo) Having(conds ...gen.Condition) IWxUserDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w wxUserDo) Limit(limit int) IWxUserDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w wxUserDo) Offset(offset int) IWxUserDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w wxUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWxUserDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w wxUserDo) Unscoped() IWxUserDo {
	return w.withDO(w.DO.Unscoped())
}

func (w wxUserDo) Create(values ...*model.WxUser) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w wxUserDo) CreateInBatches(values []*model.WxUser, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w wxUserDo) Save(values ...*model.WxUser) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w wxUserDo) First() (*model.WxUser, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WxUser), nil
	}
}

func (w wxUserDo) Take() (*model.WxUser, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WxUser), nil
	}
}

func (w wxUserDo) Last() (*model.WxUser, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WxUser), nil
	}
}

func (w wxUserDo) Find() ([]*model.WxUser, error) {
	result, err := w.DO.Find()
	return result.([]*model.WxUser), err
}

func (w wxUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WxUser, err error) {
	buf := make([]*model.WxUser, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w wxUserDo) FindInBatches(result *[]*model.WxUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w wxUserDo) Attrs(attrs ...field.AssignExpr) IWxUserDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w wxUserDo) Assign(attrs ...field.AssignExpr) IWxUserDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w wxUserDo) Joins(fields ...field.RelationField) IWxUserDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w wxUserDo) Preload(fields ...field.RelationField) IWxUserDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w wxUserDo) FirstOrInit() (*model.WxUser, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WxUser), nil
	}
}

func (w wxUserDo) FirstOrCreate() (*model.WxUser, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WxUser), nil
	}
}

func (w wxUserDo) FindByPage(offset int, limit int) (result []*model.WxUser, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w wxUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w wxUserDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w wxUserDo) Delete(models ...*model.WxUser) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *wxUserDo) withDO(do gen.Dao) *wxUserDo {
	w.DO = *do.(*gen.DO)
	return w
}
