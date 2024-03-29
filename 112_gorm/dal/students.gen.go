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

	"learn/112_gorm/model"
)

func newStudent(db *gorm.DB, opts ...gen.DOOption) student {
	_student := student{}

	_student.studentDo.UseDB(db, opts...)
	_student.studentDo.UseModel(&model.Student{})

	tableName := _student.studentDo.TableName()
	_student.ALL = field.NewAsterisk(tableName)
	_student.ID = field.NewInt64(tableName, "id")
	_student.Name = field.NewString(tableName, "name")
	_student.Age = field.NewInt64(tableName, "age")

	_student.fillFieldMap()

	return _student
}

type student struct {
	studentDo studentDo

	ALL  field.Asterisk
	ID   field.Int64
	Name field.String
	Age  field.Int64

	fieldMap map[string]field.Expr
}

func (s student) Table(newTableName string) *student {
	s.studentDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s student) As(alias string) *student {
	s.studentDo.DO = *(s.studentDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *student) updateTableName(table string) *student {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.Name = field.NewString(table, "name")
	s.Age = field.NewInt64(table, "age")

	s.fillFieldMap()

	return s
}

func (s *student) WithContext(ctx context.Context) *studentDo { return s.studentDo.WithContext(ctx) }

func (s student) TableName() string { return s.studentDo.TableName() }

func (s student) Alias() string { return s.studentDo.Alias() }

func (s student) Columns(cols ...field.Expr) gen.Columns { return s.studentDo.Columns(cols...) }

func (s *student) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *student) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 3)
	s.fieldMap["id"] = s.ID
	s.fieldMap["name"] = s.Name
	s.fieldMap["age"] = s.Age
}

func (s student) clone(db *gorm.DB) student {
	s.studentDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s student) replaceDB(db *gorm.DB) student {
	s.studentDo.ReplaceDB(db)
	return s
}

type studentDo struct{ gen.DO }

func (s studentDo) Debug() *studentDo {
	return s.withDO(s.DO.Debug())
}

func (s studentDo) WithContext(ctx context.Context) *studentDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s studentDo) ReadDB() *studentDo {
	return s.Clauses(dbresolver.Read)
}

func (s studentDo) WriteDB() *studentDo {
	return s.Clauses(dbresolver.Write)
}

func (s studentDo) Session(config *gorm.Session) *studentDo {
	return s.withDO(s.DO.Session(config))
}

func (s studentDo) Clauses(conds ...clause.Expression) *studentDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s studentDo) Returning(value interface{}, columns ...string) *studentDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s studentDo) Not(conds ...gen.Condition) *studentDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s studentDo) Or(conds ...gen.Condition) *studentDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s studentDo) Select(conds ...field.Expr) *studentDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s studentDo) Where(conds ...gen.Condition) *studentDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s studentDo) Order(conds ...field.Expr) *studentDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s studentDo) Distinct(cols ...field.Expr) *studentDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s studentDo) Omit(cols ...field.Expr) *studentDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s studentDo) Join(table schema.Tabler, on ...field.Expr) *studentDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s studentDo) LeftJoin(table schema.Tabler, on ...field.Expr) *studentDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s studentDo) RightJoin(table schema.Tabler, on ...field.Expr) *studentDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s studentDo) Group(cols ...field.Expr) *studentDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s studentDo) Having(conds ...gen.Condition) *studentDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s studentDo) Limit(limit int) *studentDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s studentDo) Offset(offset int) *studentDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s studentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *studentDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s studentDo) Unscoped() *studentDo {
	return s.withDO(s.DO.Unscoped())
}

func (s studentDo) Create(values ...*model.Student) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s studentDo) CreateInBatches(values []*model.Student, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s studentDo) Save(values ...*model.Student) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s studentDo) First() (*model.Student, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Student), nil
	}
}

func (s studentDo) Take() (*model.Student, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Student), nil
	}
}

func (s studentDo) Last() (*model.Student, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Student), nil
	}
}

func (s studentDo) Find() ([]*model.Student, error) {
	result, err := s.DO.Find()
	return result.([]*model.Student), err
}

func (s studentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Student, err error) {
	buf := make([]*model.Student, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s studentDo) FindInBatches(result *[]*model.Student, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s studentDo) Attrs(attrs ...field.AssignExpr) *studentDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s studentDo) Assign(attrs ...field.AssignExpr) *studentDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s studentDo) Joins(fields ...field.RelationField) *studentDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s studentDo) Preload(fields ...field.RelationField) *studentDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s studentDo) FirstOrInit() (*model.Student, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Student), nil
	}
}

func (s studentDo) FirstOrCreate() (*model.Student, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Student), nil
	}
}

func (s studentDo) FindByPage(offset int, limit int) (result []*model.Student, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s studentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s studentDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s studentDo) Delete(models ...*model.Student) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *studentDo) withDO(do gen.Dao) *studentDo {
	s.DO = *do.(*gen.DO)
	return s
}
