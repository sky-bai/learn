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

func newMiniProgram(db *gorm.DB, opts ...gen.DOOption) miniProgram {
	_miniProgram := miniProgram{}

	_miniProgram.miniProgramDo.UseDB(db, opts...)
	_miniProgram.miniProgramDo.UseModel(&model.MiniProgram{})

	tableName := _miniProgram.miniProgramDo.TableName()
	_miniProgram.ALL = field.NewAsterisk(tableName)
	_miniProgram.ID = field.NewInt64(tableName, "id")
	_miniProgram.CreatedAt = field.NewTime(tableName, "created_at")
	_miniProgram.UpdatedAt = field.NewTime(tableName, "updated_at")
	_miniProgram.DeletedAt = field.NewField(tableName, "deleted_at")
	_miniProgram.Customer = field.NewString(tableName, "customer")
	_miniProgram.Name = field.NewString(tableName, "name")
	_miniProgram.AppID = field.NewString(tableName, "app_Id")
	_miniProgram.AppSecret = field.NewString(tableName, "app_secret")
	_miniProgram.JsToken = field.NewString(tableName, "js_token")
	_miniProgram.Token = field.NewString(tableName, "token")
	_miniProgram.TokenUpdateTime = field.NewTime(tableName, "token_update_time")
	_miniProgram.URL = field.NewString(tableName, "url")
	_miniProgram.ServerToken = field.NewString(tableName, "server_token")
	_miniProgram.EncodingAesKey = field.NewString(tableName, "encoding_aes_key")
	_miniProgram.AppIDAlias = field.NewString(tableName, "app_Id_alias")

	_miniProgram.fillFieldMap()

	return _miniProgram
}

type miniProgram struct {
	miniProgramDo

	ALL             field.Asterisk
	ID              field.Int64 // 主键Id
	CreatedAt       field.Time  // 创建时间
	UpdatedAt       field.Time  // 修改时间
	DeletedAt       field.Field // 删除时间
	Customer        field.String
	Name            field.String
	AppID           field.String
	AppSecret       field.String
	JsToken         field.String
	Token           field.String
	TokenUpdateTime field.Time
	URL             field.String
	ServerToken     field.String
	EncodingAesKey  field.String
	AppIDAlias      field.String // appId的别名

	fieldMap map[string]field.Expr
}

func (m miniProgram) Table(newTableName string) *miniProgram {
	m.miniProgramDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m miniProgram) As(alias string) *miniProgram {
	m.miniProgramDo.DO = *(m.miniProgramDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *miniProgram) updateTableName(table string) *miniProgram {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt64(table, "id")
	m.CreatedAt = field.NewTime(table, "created_at")
	m.UpdatedAt = field.NewTime(table, "updated_at")
	m.DeletedAt = field.NewField(table, "deleted_at")
	m.Customer = field.NewString(table, "customer")
	m.Name = field.NewString(table, "name")
	m.AppID = field.NewString(table, "app_Id")
	m.AppSecret = field.NewString(table, "app_secret")
	m.JsToken = field.NewString(table, "js_token")
	m.Token = field.NewString(table, "token")
	m.TokenUpdateTime = field.NewTime(table, "token_update_time")
	m.URL = field.NewString(table, "url")
	m.ServerToken = field.NewString(table, "server_token")
	m.EncodingAesKey = field.NewString(table, "encoding_aes_key")
	m.AppIDAlias = field.NewString(table, "app_Id_alias")

	m.fillFieldMap()

	return m
}

func (m *miniProgram) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *miniProgram) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 15)
	m.fieldMap["id"] = m.ID
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
	m.fieldMap["deleted_at"] = m.DeletedAt
	m.fieldMap["customer"] = m.Customer
	m.fieldMap["name"] = m.Name
	m.fieldMap["app_Id"] = m.AppID
	m.fieldMap["app_secret"] = m.AppSecret
	m.fieldMap["js_token"] = m.JsToken
	m.fieldMap["token"] = m.Token
	m.fieldMap["token_update_time"] = m.TokenUpdateTime
	m.fieldMap["url"] = m.URL
	m.fieldMap["server_token"] = m.ServerToken
	m.fieldMap["encoding_aes_key"] = m.EncodingAesKey
	m.fieldMap["app_Id_alias"] = m.AppIDAlias
}

func (m miniProgram) clone(db *gorm.DB) miniProgram {
	m.miniProgramDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m miniProgram) replaceDB(db *gorm.DB) miniProgram {
	m.miniProgramDo.ReplaceDB(db)
	return m
}

type miniProgramDo struct{ gen.DO }

type IMiniProgramDo interface {
	gen.SubQuery
	Debug() IMiniProgramDo
	WithContext(ctx context.Context) IMiniProgramDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMiniProgramDo
	WriteDB() IMiniProgramDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMiniProgramDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMiniProgramDo
	Not(conds ...gen.Condition) IMiniProgramDo
	Or(conds ...gen.Condition) IMiniProgramDo
	Select(conds ...field.Expr) IMiniProgramDo
	Where(conds ...gen.Condition) IMiniProgramDo
	Order(conds ...field.Expr) IMiniProgramDo
	Distinct(cols ...field.Expr) IMiniProgramDo
	Omit(cols ...field.Expr) IMiniProgramDo
	Join(table schema.Tabler, on ...field.Expr) IMiniProgramDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMiniProgramDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMiniProgramDo
	Group(cols ...field.Expr) IMiniProgramDo
	Having(conds ...gen.Condition) IMiniProgramDo
	Limit(limit int) IMiniProgramDo
	Offset(offset int) IMiniProgramDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMiniProgramDo
	Unscoped() IMiniProgramDo
	Create(values ...*model.MiniProgram) error
	CreateInBatches(values []*model.MiniProgram, batchSize int) error
	Save(values ...*model.MiniProgram) error
	First() (*model.MiniProgram, error)
	Take() (*model.MiniProgram, error)
	Last() (*model.MiniProgram, error)
	Find() ([]*model.MiniProgram, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MiniProgram, err error)
	FindInBatches(result *[]*model.MiniProgram, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.MiniProgram) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMiniProgramDo
	Assign(attrs ...field.AssignExpr) IMiniProgramDo
	Joins(fields ...field.RelationField) IMiniProgramDo
	Preload(fields ...field.RelationField) IMiniProgramDo
	FirstOrInit() (*model.MiniProgram, error)
	FirstOrCreate() (*model.MiniProgram, error)
	FindByPage(offset int, limit int) (result []*model.MiniProgram, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMiniProgramDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m miniProgramDo) Debug() IMiniProgramDo {
	return m.withDO(m.DO.Debug())
}

func (m miniProgramDo) WithContext(ctx context.Context) IMiniProgramDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m miniProgramDo) ReadDB() IMiniProgramDo {
	return m.Clauses(dbresolver.Read)
}

func (m miniProgramDo) WriteDB() IMiniProgramDo {
	return m.Clauses(dbresolver.Write)
}

func (m miniProgramDo) Session(config *gorm.Session) IMiniProgramDo {
	return m.withDO(m.DO.Session(config))
}

func (m miniProgramDo) Clauses(conds ...clause.Expression) IMiniProgramDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m miniProgramDo) Returning(value interface{}, columns ...string) IMiniProgramDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m miniProgramDo) Not(conds ...gen.Condition) IMiniProgramDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m miniProgramDo) Or(conds ...gen.Condition) IMiniProgramDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m miniProgramDo) Select(conds ...field.Expr) IMiniProgramDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m miniProgramDo) Where(conds ...gen.Condition) IMiniProgramDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m miniProgramDo) Order(conds ...field.Expr) IMiniProgramDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m miniProgramDo) Distinct(cols ...field.Expr) IMiniProgramDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m miniProgramDo) Omit(cols ...field.Expr) IMiniProgramDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m miniProgramDo) Join(table schema.Tabler, on ...field.Expr) IMiniProgramDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m miniProgramDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMiniProgramDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m miniProgramDo) RightJoin(table schema.Tabler, on ...field.Expr) IMiniProgramDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m miniProgramDo) Group(cols ...field.Expr) IMiniProgramDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m miniProgramDo) Having(conds ...gen.Condition) IMiniProgramDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m miniProgramDo) Limit(limit int) IMiniProgramDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m miniProgramDo) Offset(offset int) IMiniProgramDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m miniProgramDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMiniProgramDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m miniProgramDo) Unscoped() IMiniProgramDo {
	return m.withDO(m.DO.Unscoped())
}

func (m miniProgramDo) Create(values ...*model.MiniProgram) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m miniProgramDo) CreateInBatches(values []*model.MiniProgram, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m miniProgramDo) Save(values ...*model.MiniProgram) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m miniProgramDo) First() (*model.MiniProgram, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MiniProgram), nil
	}
}

func (m miniProgramDo) Take() (*model.MiniProgram, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MiniProgram), nil
	}
}

func (m miniProgramDo) Last() (*model.MiniProgram, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MiniProgram), nil
	}
}

func (m miniProgramDo) Find() ([]*model.MiniProgram, error) {
	result, err := m.DO.Find()
	return result.([]*model.MiniProgram), err
}

func (m miniProgramDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MiniProgram, err error) {
	buf := make([]*model.MiniProgram, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m miniProgramDo) FindInBatches(result *[]*model.MiniProgram, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m miniProgramDo) Attrs(attrs ...field.AssignExpr) IMiniProgramDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m miniProgramDo) Assign(attrs ...field.AssignExpr) IMiniProgramDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m miniProgramDo) Joins(fields ...field.RelationField) IMiniProgramDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m miniProgramDo) Preload(fields ...field.RelationField) IMiniProgramDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m miniProgramDo) FirstOrInit() (*model.MiniProgram, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MiniProgram), nil
	}
}

func (m miniProgramDo) FirstOrCreate() (*model.MiniProgram, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MiniProgram), nil
	}
}

func (m miniProgramDo) FindByPage(offset int, limit int) (result []*model.MiniProgram, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m miniProgramDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m miniProgramDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m miniProgramDo) Delete(models ...*model.MiniProgram) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *miniProgramDo) withDO(do gen.Dao) *miniProgramDo {
	m.DO = *do.(*gen.DO)
	return m
}
