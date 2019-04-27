package main

var memberTemplate = "\t%v\t%v\t`db:\"%v\" bson:\"%v\" comment:\"%v\" json:\"%v\" xml:\"%v\"`"
var replaces = &map[string]string{
	"packageName":  "",
	"imports":      "",
	"structName":   "",
	"attributes":   "",
	"reset":        "",
	"asMap":        "",
	"tableName":    "",
	"beforeInsert": "",
	"afterInsert":  "",
	"beforeUpdate": "",
	"setUpdatedAt": "",
	"beforeDelete": "",
}
var structFuncs = map[string]string{
	`Trans`:        `Trans`,
	`Use`:          `Use`,
	`Objects`:      `Objects`,
	`NewObjects`:   `NewObjects`,
	`NewParam`:     `NewParam`,
	`SetParam`:     `SetParam`,
	`Get`:          `Get`,
	`List`:         `List`,
	`ListByOffset`: `ListByOffset`,
	`Add`:          `Add`,
	`Edit`:         `Edit`,
	`Upsert`:       `Upsert`,
	`Delete`:       `Delete`,
	`Count`:        `Count`,
}
var structTemplate = `//Do not edit this file, which is automatically generated by the generator.

package {{packageName}}

import (
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	{{imports}}
)

type {{structName}} struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*{{structName}}
	namer   func(string) string
	
{{attributes}}
}

func (this *{{structName}}) Trans() *factory.Transaction {
	return this.trans
}

func (this *{{structName}}) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *{{structName}}) Objects() []*{{structName}} {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *{{structName}}) NewObjects() *[]*{{structName}} {
	this.objects = []*{{structName}}{}
	return &this.objects
}

func (this *{{structName}}) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *{{structName}}) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *{{structName}}) Name_() string {
	if this.namer != nil {
		return this.namer("{{tableName}}")
	}
	return "{{tableName}}"
}

func (this *{{structName}}) FullName_() string {
	return factory.DefaultFactory.Table(this.Name_())
}

func (this *{{structName}}) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *{{structName}}) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *{{structName}}) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *{{structName}}) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *{{structName}}) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *{{structName}}) Add() (pk interface{}, err error) {
	{{beforeInsert}}
	pk, err = this.Param().SetSend(this).Insert()
	{{afterInsert}}
	return
}

func (this *{{structName}}) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	{{beforeUpdate}}
	return this.Setter(mw, args...).SetSend(this).Update()
}

func (this *{{structName}}) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *{{structName}}) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) error {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *{{structName}}) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) error {
	{{setUpdatedAt}}
	return this.Setter(mw, args...).SetSend(kvset).Update()
}

func (this *{{structName}}) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		{{beforeUpdate}}
	},func(){
		{{beforeInsert}}
	})
	{{afterInsert}}
	return 
}

func (this *{{structName}}) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	{{beforeDelete}}
	return this.Param().SetArgs(args...).SetMiddleware(mw).Delete()
}

func (this *{{structName}}) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *{{structName}}) Reset() *{{structName}} {
{{reset}}
	return this
}

func (this *{{structName}}) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
{{asMap}}
	return r
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

var initFileTemplate = `//Do not edit this file, which is automatically generated by the generator.

package {{packageName}}

import (
	"github.com/webx-top/db/lib/factory"
)

func init(){
	{{initCode}}
}

`
