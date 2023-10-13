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

func newRoleAPI(db *gorm.DB, opts ...gen.DOOption) roleAPI {
	_roleAPI := roleAPI{}

	_roleAPI.roleAPIDo.UseDB(db, opts...)
	_roleAPI.roleAPIDo.UseModel(&model.RoleAPI{})

	tableName := _roleAPI.roleAPIDo.TableName()
	_roleAPI.ALL = field.NewAsterisk(tableName)
	_roleAPI.ID = field.NewInt64(tableName, "id")
	_roleAPI.CreatedAt = field.NewTime(tableName, "created_at")
	_roleAPI.UpdatedAt = field.NewTime(tableName, "updated_at")
	_roleAPI.DeletedAt = field.NewField(tableName, "deleted_at")
	_roleAPI.Role = field.NewString(tableName, "role")
	_roleAPI.URI = field.NewString(tableName, "uri")
	_roleAPI.Method = field.NewString(tableName, "method")

	_roleAPI.fillFieldMap()

	return _roleAPI
}

type roleAPI struct {
	roleAPIDo

	ALL       field.Asterisk
	ID        field.Int64  // 主键Id
	CreatedAt field.Time   // 创建时间
	UpdatedAt field.Time   // 修改时间
	DeletedAt field.Field  // 删除时间
	Role      field.String // 角色
	URI       field.String // 允许的请求uri
	Method    field.String // 允许的请求方法

	fieldMap map[string]field.Expr
}

func (r roleAPI) Table(newTableName string) *roleAPI {
	r.roleAPIDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r roleAPI) As(alias string) *roleAPI {
	r.roleAPIDo.DO = *(r.roleAPIDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *roleAPI) updateTableName(table string) *roleAPI {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt64(table, "id")
	r.CreatedAt = field.NewTime(table, "created_at")
	r.UpdatedAt = field.NewTime(table, "updated_at")
	r.DeletedAt = field.NewField(table, "deleted_at")
	r.Role = field.NewString(table, "role")
	r.URI = field.NewString(table, "uri")
	r.Method = field.NewString(table, "method")

	r.fillFieldMap()

	return r
}

func (r *roleAPI) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *roleAPI) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 7)
	r.fieldMap["id"] = r.ID
	r.fieldMap["created_at"] = r.CreatedAt
	r.fieldMap["updated_at"] = r.UpdatedAt
	r.fieldMap["deleted_at"] = r.DeletedAt
	r.fieldMap["role"] = r.Role
	r.fieldMap["uri"] = r.URI
	r.fieldMap["method"] = r.Method
}

func (r roleAPI) clone(db *gorm.DB) roleAPI {
	r.roleAPIDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r roleAPI) replaceDB(db *gorm.DB) roleAPI {
	r.roleAPIDo.ReplaceDB(db)
	return r
}

type roleAPIDo struct{ gen.DO }

type IRoleAPIDo interface {
	gen.SubQuery
	Debug() IRoleAPIDo
	WithContext(ctx context.Context) IRoleAPIDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IRoleAPIDo
	WriteDB() IRoleAPIDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IRoleAPIDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IRoleAPIDo
	Not(conds ...gen.Condition) IRoleAPIDo
	Or(conds ...gen.Condition) IRoleAPIDo
	Select(conds ...field.Expr) IRoleAPIDo
	Where(conds ...gen.Condition) IRoleAPIDo
	Order(conds ...field.Expr) IRoleAPIDo
	Distinct(cols ...field.Expr) IRoleAPIDo
	Omit(cols ...field.Expr) IRoleAPIDo
	Join(table schema.Tabler, on ...field.Expr) IRoleAPIDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IRoleAPIDo
	RightJoin(table schema.Tabler, on ...field.Expr) IRoleAPIDo
	Group(cols ...field.Expr) IRoleAPIDo
	Having(conds ...gen.Condition) IRoleAPIDo
	Limit(limit int) IRoleAPIDo
	Offset(offset int) IRoleAPIDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IRoleAPIDo
	Unscoped() IRoleAPIDo
	Create(values ...*model.RoleAPI) error
	CreateInBatches(values []*model.RoleAPI, batchSize int) error
	Save(values ...*model.RoleAPI) error
	First() (*model.RoleAPI, error)
	Take() (*model.RoleAPI, error)
	Last() (*model.RoleAPI, error)
	Find() ([]*model.RoleAPI, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RoleAPI, err error)
	FindInBatches(result *[]*model.RoleAPI, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.RoleAPI) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IRoleAPIDo
	Assign(attrs ...field.AssignExpr) IRoleAPIDo
	Joins(fields ...field.RelationField) IRoleAPIDo
	Preload(fields ...field.RelationField) IRoleAPIDo
	FirstOrInit() (*model.RoleAPI, error)
	FirstOrCreate() (*model.RoleAPI, error)
	FindByPage(offset int, limit int) (result []*model.RoleAPI, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IRoleAPIDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r roleAPIDo) Debug() IRoleAPIDo {
	return r.withDO(r.DO.Debug())
}

func (r roleAPIDo) WithContext(ctx context.Context) IRoleAPIDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r roleAPIDo) ReadDB() IRoleAPIDo {
	return r.Clauses(dbresolver.Read)
}

func (r roleAPIDo) WriteDB() IRoleAPIDo {
	return r.Clauses(dbresolver.Write)
}

func (r roleAPIDo) Session(config *gorm.Session) IRoleAPIDo {
	return r.withDO(r.DO.Session(config))
}

func (r roleAPIDo) Clauses(conds ...clause.Expression) IRoleAPIDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r roleAPIDo) Returning(value interface{}, columns ...string) IRoleAPIDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r roleAPIDo) Not(conds ...gen.Condition) IRoleAPIDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r roleAPIDo) Or(conds ...gen.Condition) IRoleAPIDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r roleAPIDo) Select(conds ...field.Expr) IRoleAPIDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r roleAPIDo) Where(conds ...gen.Condition) IRoleAPIDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r roleAPIDo) Order(conds ...field.Expr) IRoleAPIDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r roleAPIDo) Distinct(cols ...field.Expr) IRoleAPIDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r roleAPIDo) Omit(cols ...field.Expr) IRoleAPIDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r roleAPIDo) Join(table schema.Tabler, on ...field.Expr) IRoleAPIDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r roleAPIDo) LeftJoin(table schema.Tabler, on ...field.Expr) IRoleAPIDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r roleAPIDo) RightJoin(table schema.Tabler, on ...field.Expr) IRoleAPIDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r roleAPIDo) Group(cols ...field.Expr) IRoleAPIDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r roleAPIDo) Having(conds ...gen.Condition) IRoleAPIDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r roleAPIDo) Limit(limit int) IRoleAPIDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r roleAPIDo) Offset(offset int) IRoleAPIDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r roleAPIDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IRoleAPIDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r roleAPIDo) Unscoped() IRoleAPIDo {
	return r.withDO(r.DO.Unscoped())
}

func (r roleAPIDo) Create(values ...*model.RoleAPI) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r roleAPIDo) CreateInBatches(values []*model.RoleAPI, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r roleAPIDo) Save(values ...*model.RoleAPI) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r roleAPIDo) First() (*model.RoleAPI, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.RoleAPI), nil
	}
}

func (r roleAPIDo) Take() (*model.RoleAPI, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.RoleAPI), nil
	}
}

func (r roleAPIDo) Last() (*model.RoleAPI, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.RoleAPI), nil
	}
}

func (r roleAPIDo) Find() ([]*model.RoleAPI, error) {
	result, err := r.DO.Find()
	return result.([]*model.RoleAPI), err
}

func (r roleAPIDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RoleAPI, err error) {
	buf := make([]*model.RoleAPI, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r roleAPIDo) FindInBatches(result *[]*model.RoleAPI, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r roleAPIDo) Attrs(attrs ...field.AssignExpr) IRoleAPIDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r roleAPIDo) Assign(attrs ...field.AssignExpr) IRoleAPIDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r roleAPIDo) Joins(fields ...field.RelationField) IRoleAPIDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r roleAPIDo) Preload(fields ...field.RelationField) IRoleAPIDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r roleAPIDo) FirstOrInit() (*model.RoleAPI, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.RoleAPI), nil
	}
}

func (r roleAPIDo) FirstOrCreate() (*model.RoleAPI, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.RoleAPI), nil
	}
}

func (r roleAPIDo) FindByPage(offset int, limit int) (result []*model.RoleAPI, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r roleAPIDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r roleAPIDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r roleAPIDo) Delete(models ...*model.RoleAPI) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *roleAPIDo) withDO(do gen.Dao) *roleAPIDo {
	r.DO = *do.(*gen.DO)
	return r
}
