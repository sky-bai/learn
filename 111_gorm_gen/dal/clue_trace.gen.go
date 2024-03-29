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

func newClueTrace(db *gorm.DB, opts ...gen.DOOption) clueTrace {
	_clueTrace := clueTrace{}

	_clueTrace.clueTraceDo.UseDB(db, opts...)
	_clueTrace.clueTraceDo.UseModel(&model.ClueTrace{})

	tableName := _clueTrace.clueTraceDo.TableName()
	_clueTrace.ALL = field.NewAsterisk(tableName)
	_clueTrace.ID = field.NewInt64(tableName, "id")
	_clueTrace.CreatedAt = field.NewTime(tableName, "created_at")
	_clueTrace.UpdatedAt = field.NewTime(tableName, "updated_at")
	_clueTrace.DeletedAt = field.NewField(tableName, "deleted_at")
	_clueTrace.Status = field.NewInt32(tableName, "status")
	_clueTrace.TraceContent = field.NewString(tableName, "trace_content")
	_clueTrace.ClueID = field.NewInt64(tableName, "clue_id")
	_clueTrace.TraceType = field.NewString(tableName, "trace_type")
	_clueTrace.AccountID = field.NewInt64(tableName, "account_id")
	_clueTrace.EnterpriseID = field.NewString(tableName, "enterprise_id")
	_clueTrace.SubsidiaryID = field.NewInt64(tableName, "subsidiary_id")

	_clueTrace.fillFieldMap()

	return _clueTrace
}

type clueTrace struct {
	clueTraceDo

	ALL          field.Asterisk
	ID           field.Int64  // 主键Id
	CreatedAt    field.Time   // 创建时间
	UpdatedAt    field.Time   // 修改时间
	DeletedAt    field.Field  // 删除时间
	Status       field.Int32  // 线索跟进状态
	TraceContent field.String // 跟进内容
	ClueID       field.Int64  // 线索ID
	TraceType    field.String // 跟进类型，查阅车主信息、跟进记录
	AccountID    field.Int64  // 账号ID
	EnterpriseID field.String // 集团id,主系统wid
	SubsidiaryID field.Int64  // 所属4S店

	fieldMap map[string]field.Expr
}

func (c clueTrace) Table(newTableName string) *clueTrace {
	c.clueTraceDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c clueTrace) As(alias string) *clueTrace {
	c.clueTraceDo.DO = *(c.clueTraceDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *clueTrace) updateTableName(table string) *clueTrace {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.Status = field.NewInt32(table, "status")
	c.TraceContent = field.NewString(table, "trace_content")
	c.ClueID = field.NewInt64(table, "clue_id")
	c.TraceType = field.NewString(table, "trace_type")
	c.AccountID = field.NewInt64(table, "account_id")
	c.EnterpriseID = field.NewString(table, "enterprise_id")
	c.SubsidiaryID = field.NewInt64(table, "subsidiary_id")

	c.fillFieldMap()

	return c
}

func (c *clueTrace) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *clueTrace) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 11)
	c.fieldMap["id"] = c.ID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["status"] = c.Status
	c.fieldMap["trace_content"] = c.TraceContent
	c.fieldMap["clue_id"] = c.ClueID
	c.fieldMap["trace_type"] = c.TraceType
	c.fieldMap["account_id"] = c.AccountID
	c.fieldMap["enterprise_id"] = c.EnterpriseID
	c.fieldMap["subsidiary_id"] = c.SubsidiaryID
}

func (c clueTrace) clone(db *gorm.DB) clueTrace {
	c.clueTraceDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c clueTrace) replaceDB(db *gorm.DB) clueTrace {
	c.clueTraceDo.ReplaceDB(db)
	return c
}

type clueTraceDo struct{ gen.DO }

type IClueTraceDo interface {
	gen.SubQuery
	Debug() IClueTraceDo
	WithContext(ctx context.Context) IClueTraceDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IClueTraceDo
	WriteDB() IClueTraceDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IClueTraceDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IClueTraceDo
	Not(conds ...gen.Condition) IClueTraceDo
	Or(conds ...gen.Condition) IClueTraceDo
	Select(conds ...field.Expr) IClueTraceDo
	Where(conds ...gen.Condition) IClueTraceDo
	Order(conds ...field.Expr) IClueTraceDo
	Distinct(cols ...field.Expr) IClueTraceDo
	Omit(cols ...field.Expr) IClueTraceDo
	Join(table schema.Tabler, on ...field.Expr) IClueTraceDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IClueTraceDo
	RightJoin(table schema.Tabler, on ...field.Expr) IClueTraceDo
	Group(cols ...field.Expr) IClueTraceDo
	Having(conds ...gen.Condition) IClueTraceDo
	Limit(limit int) IClueTraceDo
	Offset(offset int) IClueTraceDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IClueTraceDo
	Unscoped() IClueTraceDo
	Create(values ...*model.ClueTrace) error
	CreateInBatches(values []*model.ClueTrace, batchSize int) error
	Save(values ...*model.ClueTrace) error
	First() (*model.ClueTrace, error)
	Take() (*model.ClueTrace, error)
	Last() (*model.ClueTrace, error)
	Find() ([]*model.ClueTrace, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ClueTrace, err error)
	FindInBatches(result *[]*model.ClueTrace, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ClueTrace) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IClueTraceDo
	Assign(attrs ...field.AssignExpr) IClueTraceDo
	Joins(fields ...field.RelationField) IClueTraceDo
	Preload(fields ...field.RelationField) IClueTraceDo
	FirstOrInit() (*model.ClueTrace, error)
	FirstOrCreate() (*model.ClueTrace, error)
	FindByPage(offset int, limit int) (result []*model.ClueTrace, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IClueTraceDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c clueTraceDo) Debug() IClueTraceDo {
	return c.withDO(c.DO.Debug())
}

func (c clueTraceDo) WithContext(ctx context.Context) IClueTraceDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c clueTraceDo) ReadDB() IClueTraceDo {
	return c.Clauses(dbresolver.Read)
}

func (c clueTraceDo) WriteDB() IClueTraceDo {
	return c.Clauses(dbresolver.Write)
}

func (c clueTraceDo) Session(config *gorm.Session) IClueTraceDo {
	return c.withDO(c.DO.Session(config))
}

func (c clueTraceDo) Clauses(conds ...clause.Expression) IClueTraceDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c clueTraceDo) Returning(value interface{}, columns ...string) IClueTraceDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c clueTraceDo) Not(conds ...gen.Condition) IClueTraceDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c clueTraceDo) Or(conds ...gen.Condition) IClueTraceDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c clueTraceDo) Select(conds ...field.Expr) IClueTraceDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c clueTraceDo) Where(conds ...gen.Condition) IClueTraceDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c clueTraceDo) Order(conds ...field.Expr) IClueTraceDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c clueTraceDo) Distinct(cols ...field.Expr) IClueTraceDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c clueTraceDo) Omit(cols ...field.Expr) IClueTraceDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c clueTraceDo) Join(table schema.Tabler, on ...field.Expr) IClueTraceDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c clueTraceDo) LeftJoin(table schema.Tabler, on ...field.Expr) IClueTraceDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c clueTraceDo) RightJoin(table schema.Tabler, on ...field.Expr) IClueTraceDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c clueTraceDo) Group(cols ...field.Expr) IClueTraceDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c clueTraceDo) Having(conds ...gen.Condition) IClueTraceDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c clueTraceDo) Limit(limit int) IClueTraceDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c clueTraceDo) Offset(offset int) IClueTraceDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c clueTraceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IClueTraceDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c clueTraceDo) Unscoped() IClueTraceDo {
	return c.withDO(c.DO.Unscoped())
}

func (c clueTraceDo) Create(values ...*model.ClueTrace) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c clueTraceDo) CreateInBatches(values []*model.ClueTrace, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c clueTraceDo) Save(values ...*model.ClueTrace) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c clueTraceDo) First() (*model.ClueTrace, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ClueTrace), nil
	}
}

func (c clueTraceDo) Take() (*model.ClueTrace, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ClueTrace), nil
	}
}

func (c clueTraceDo) Last() (*model.ClueTrace, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ClueTrace), nil
	}
}

func (c clueTraceDo) Find() ([]*model.ClueTrace, error) {
	result, err := c.DO.Find()
	return result.([]*model.ClueTrace), err
}

func (c clueTraceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ClueTrace, err error) {
	buf := make([]*model.ClueTrace, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c clueTraceDo) FindInBatches(result *[]*model.ClueTrace, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c clueTraceDo) Attrs(attrs ...field.AssignExpr) IClueTraceDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c clueTraceDo) Assign(attrs ...field.AssignExpr) IClueTraceDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c clueTraceDo) Joins(fields ...field.RelationField) IClueTraceDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c clueTraceDo) Preload(fields ...field.RelationField) IClueTraceDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c clueTraceDo) FirstOrInit() (*model.ClueTrace, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ClueTrace), nil
	}
}

func (c clueTraceDo) FirstOrCreate() (*model.ClueTrace, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ClueTrace), nil
	}
}

func (c clueTraceDo) FindByPage(offset int, limit int) (result []*model.ClueTrace, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c clueTraceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c clueTraceDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c clueTraceDo) Delete(models ...*model.ClueTrace) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *clueTraceDo) withDO(do gen.Dao) *clueTraceDo {
	c.DO = *do.(*gen.DO)
	return c
}
