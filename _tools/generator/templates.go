package main

import (
	"fmt"
)

type structField struct {
	field   string
	typ     string
	dbTag   string
	bsonTag string
	jsonTag string
	xmlTag  string
	comment string
}

func (f *structField) String() string {
	return fmt.Sprintf(memberTemplate, f.field, f.typ, f.dbTag, f.bsonTag, f.comment, f.jsonTag, f.xmlTag)
}

var memberTemplate = "\t%v\t%v\t`db:\"%v\" bson:\"%v\" comment:\"%v\" json:\"%v\" xml:\"%v\"`"
var replaces = &map[string]string{
	"packageName":   "",
	"imports":       "",
	"structName":    "",
	"structComment": "",
	"attributes":    "",
	"reset":         "",
	"asMap":         "",
	"asRow":         "",
	"fromRowCase":   "",
	"setCase":       "",
	"tableName":     "",
	"beforeInsert":  "",
	"afterInsert":   "",
	"beforeUpdate":  "",
	"setUpdatedAt":  "",
	"beforeDelete":  "",
}
var structFuncs = map[string]string{
	`Trans`:         `Trans`,
	`Use`:           `Use`,
	`Objects`:       `Objects`,
	`NewObjects`:    `NewObjects`,
	`NewParam`:      `NewParam`,
	`SetParam`:      `SetParam`,
	`Get`:           `Get`,
	`List`:          `List`,
	`ListByOffset`:  `ListByOffset`,
	`Add`:           `Add`,
	`Edit`:          `Edit`,
	`Upsert`:        `Upsert`,
	`Delete`:        `Delete`,
	`Count`:         `Count`,
	`GroupBy`:       `GroupBy`,
	`KeyBy`:         `KeyBy`,
	`Setter`:        `Setter`,
	`SetField`:      `SetField`,
	`SetFields`:     `SetFields`,
	`Reset`:         `Reset`,
	`AsMap`:         `AsMap`,
	`AsRow`:         `AsRow`,
	`BatchValidate`: `BatchValidate`,
	`Validate`:      `Validate`,
}
var structTemplate = `// @generated Do not edit this file, which is automatically generated by the generator.

package {{packageName}}

import (
	"fmt"
	{{imports}}

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/param"
)

type Slice_{{structName}} []*{{structName}}

func (s Slice_{{structName}}) Range(fn func(m factory.Model) error ) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_{{structName}}) RangeRaw(fn func(m *{{structName}}) error ) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

// {{structName}} {{structComment}}
type {{structName}} struct {
	base    factory.Base
	objects []*{{structName}}
	
{{attributes}}
}

// - base function

func (a *{{structName}}) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *{{structName}}) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *{{structName}}) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *{{structName}}) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *{{structName}}) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *{{structName}}) Context() echo.Context {
	return a.base.Context()
}

func (a *{{structName}}) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *{{structName}}) SetNamer(namer func (string) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *{{structName}}) Namer() func(string) string {
	return a.base.Namer()
}

func (a *{{structName}}) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *{{structName}}) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

// - current function

func (a *{{structName}}) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName,a.base.ConnID()).Use(a.base.Trans())
}

func (a *{{structName}}) Objects() []*{{structName}} {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *{{structName}}) NewObjects() factory.Ranger {
	return &Slice_{{structName}}{}
}

func (a *{{structName}}) InitObjects() *[]*{{structName}} {
	a.objects = []*{{structName}}{}
	return &a.objects
}

func (a *{{structName}}) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *{{structName}}) Short_() string {
	return "{{tableName}}"
}

func (a *{{structName}}) Struct_() string {
	return "{{structName}}"
}

func (a *{{structName}}) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *{{structName}}) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *{{structName}}) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	base := a.base
	err := a.Param(mw, args...).SetRecv(a).One()
	a.base = base
	return err
}

func (a *{{structName}}) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	return a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv).List()
}

func (a *{{structName}}) GroupBy(keyField string, inputRows ...[]*{{structName}}) map[string][]*{{structName}} {
	var rows []*{{structName}}
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = a.Objects()
	}
	r := map[string][]*{{structName}}{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*{{structName}}{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (a *{{structName}}) KeyBy(keyField string, inputRows ...[]*{{structName}}) map[string]*{{structName}} {
	var rows []*{{structName}}
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = a.Objects()
	}
	r := map[string]*{{structName}}{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (a *{{structName}}) AsKV(keyField string, valueField string, inputRows ...[]*{{structName}}) map[string]interface{} {
	var rows []*{{structName}}
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = a.Objects()
	}
	r := map[string]interface{}{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (a *{{structName}}) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	return a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv).List()
}

func (a *{{structName}}) Add() (pk interface{}, err error) {
	{{beforeInsert}}
	if a.base.Eventable() {
		err = DBI.Fire("creating", a, nil)
		if err != nil {
			return
		}
	}
	pk, err = a.Param(nil).SetSend(a).Insert()
	{{afterInsert}}
	if err == nil && a.base.Eventable() {
		err = DBI.Fire("created", a, nil)
	}
	return
}

func (a *{{structName}}) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	{{beforeUpdate}}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(a).Update()
	}
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(a).Update(); err != nil {
		return
	}
	return DBI.Fire("updated", a, mw, args...)
}

func (a *{{structName}}) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *{{structName}}) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {
	{{setUpdatedAt}}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(kvset).Update()
	}
	m := *a
	m.FromRow(kvset)
	var editColumns []string
	for column := range kvset {
		editColumns = append(editColumns, column)
	}
	if err = DBI.FireUpdate("updating", &m, editColumns, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(kvset).Update(); err != nil {
		return
	}
	return DBI.FireUpdate("updated", &m, editColumns, mw, args...)
}

func (a *{{structName}}) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error { {{beforeUpdate}}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error { {{beforeInsert}}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("creating", a, nil)
	})
	{{afterInsert}}
	if err == nil && a.base.Eventable() {
		if pk == nil {
			err = DBI.Fire("updated", a, mw, args...)
		} else {
			err = DBI.Fire("created", a, nil)
		}
	} 
	return 
}

func (a *{{structName}}) Delete(mw func(db.Result) db.Result, args ...interface{})  (err error) {
	{{beforeDelete}}
	if !a.base.Eventable() {
		return a.Param(mw, args...).Delete()
	}
	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).Delete(); err != nil {
		return
	}
	return DBI.Fire("deleted", a, mw, args...)
}

func (a *{{structName}}) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *{{structName}}) Reset() *{{structName}} {
{{reset}}
	return a
}

func (a *{{structName}}) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
{{asMap}}
	return r
}

func (a *{{structName}}) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
{{fromRowCase}}
		}
	}
}

func (a *{{structName}}) Set(key interface{}, value ...interface{}) {
	switch k := key.(type) {
		case map[string]interface{}:
			for kk, vv := range k {
				a.Set(kk, vv)
			}
		default:
			var (
				kk string
				vv interface{}
			)
			if k, y := key.(string); y {
				kk = k
			} else {
				kk = fmt.Sprint(key)
			}
			if len(value) > 0 {
				vv = value[0]
			}
			switch kk {
{{setCase}}
			}
	}
}

func (a *{{structName}}) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
{{asRow}}
	return r
}

func (a *{{structName}}) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return factory.BatchValidate(a.Short_(), kvset)
}

func (a *{{structName}}) Validate(field string, value interface{}) error {
	return factory.Validate(a.Short_(), field, value)
}

`

var modelReplaces = &map[string]string{
	"packageName":       "",
	"imports":           "",
	"structName":        "",
	"schemaPackagePath": "",
	"schemaPackageName": "",
	"baseName":          "",
}
var modelBaseTemplate = `package {{packageName}}

import (
	"github.com/webx-top/echo"
	{{imports}}
)

type {{structName}} struct {
	echo.Context
}
`

var modelTemplate = `package {{packageName}}

import (
	"{{schemaPackagePath}}"
	"github.com/webx-top/echo"
	{{imports}}
)
				
func New{{structName}}(ctx echo.Context) *{{structName}} {
	return &{{structName}}{
		{{structName}}: &{{schemaPackageName}}.{{structName}}{},
		{{baseName}}:   &{{baseName}}{Context: ctx},
	}
}

type {{structName}} struct {
	*{{schemaPackageName}}.{{structName}}
	*{{baseName}}
}

`

var initFileTemplate = `// @generated Do not edit a file, which is automatically generated by the generator.

package {{packageName}}

import (
	"github.com/webx-top/db/lib/factory"
)

var WithPrefix = func(tableName string) string {
	return "{{prefix}}" + tableName
}

var DBI = factory.NewDBI()

func init(){
	{{initCode}}
}

`
