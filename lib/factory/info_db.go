package factory

import (
	"strings"

	"github.com/webx-top/db"
)

var (
	databases = map[string]*DBI{
		DefaultDBKey: DefaultDBI,
	}
)

func DBIRegister(dbi *DBI, keys ...string) {
	key := DefaultDBKey
	if len(keys) > 0 {
		key = keys[0]
	}
	if _, y := databases[key]; y {
		panic(`DBI key already exists, please do not duplicate registrations`)
	}
	databases[key] = dbi
}

func DBIGet(keys ...string) *DBI {
	if len(keys) > 0 {
		return databases[keys[0]]
	}
	return databases[DefaultDBKey]
}

func DBIExists(key string) bool {
	_, ok := databases[key]
	return ok
}

func NewDBI() *DBI {
	return &DBI{
		Fields:      map[string]map[string]*FieldInfo{},
		Columns:     map[string][]string{},
		Models:      ModelInstancers{},
		TableNamers: map[string]func(obj interface{}) string{},
		Events:      NewEvents(),
	}
}

// DBI 数据库信息
type DBI struct {
	// Fields {table:{field:FieldInfo}}
	Fields FieldValidator
	// Columns {table:[field1,field2]}
	Columns map[string][]string
	// Models {StructName:ModelInstancer}
	Models ModelInstancers
	// TableNamers {table:NewName}
	TableNamers TableNamers
	Events      Events
}

func (d *DBI) TableName(structName string) string {
	m, ok := d.Models[structName]
	if ok {
		return m.Short
	}
	return ``
}

func (d *DBI) TableComment(structName string) string {
	m, ok := d.Models[structName]
	if ok {
		return m.Comment
	}
	return ``
}

func (d *DBI) TableColumns(tableName string) []string {
	cols, ok := d.Columns[tableName]
	if ok {
		return cols
	}
	return nil
}

func (d *DBI) Fire(event string, model Model, mw func(db.Result) db.Result, args ...interface{}) error {
	return d.Events.Call(event, model, nil, mw, args...)
}

func (d *DBI) FireUpdate(event string, model Model, editColumns []string, mw func(db.Result) db.Result, args ...interface{}) error {
	return d.Events.Call(event, model, editColumns, mw, args...)
}

func (d *DBI) FireReading(model Model, param *Param, rangers ...Ranger) error {
	return d.Events.CallRead(`reading`, model, param, rangers...)
}

func (d *DBI) FireReaded(model Model, param *Param, rangers ...Ranger) error {
	return d.Events.CallRead(`readed`, model, param, rangers...)
}

func (d *DBI) ParseEventNames(event string) []string{
	switch event {
	case `w+`:
		return AllAfterWriteEvents
	case `+w`:
		return AllBeforeWriteEvents
	case `r+`:
		return AllAfterReadEvents
	case `+r`:
		return AllBeforeReadEvents
	default:
		return strings.Split(event, ",")
	}
}

// - 注册写(CUD)事件

func (d *DBI) On(event string, h interface{}, tableName ...string) {
	var table string
	if len(tableName) > 0 {
		table = tableName[0]
	} else {
		set := strings.SplitN(event, ":", 2)
		switch len(set) {
		case 2:
			event = set[1]
			fallthrough
		case 1:
			table = set[0]
		}
	}
	for _, evt := range d.ParseEventNames(event) {
		switch evt {
		case `reading`,`readed`:
			d.Events.OnRead(evt, h.(EventReadHandler), table)
		default:
			d.Events.On(evt, h.(EventHandler), table)
		}
	}
}

func (d *DBI) OnAsync(event string, h interface{}, tableName ...string) {
	var table string
	if len(tableName) > 0 {
		table = tableName[0]
	} else {
		set := strings.SplitN(event, ":", 2)
		switch len(set) {
		case 2:
			event = set[1]
			fallthrough
		case 1:
			table = set[0]
		}
	}
	for _, evt := range d.ParseEventNames(event) {
		switch evt {
		case `reading`,`readed`:
			d.Events.OnRead(evt, h.(EventReadHandler), table, true)
		default:
			d.Events.On(evt, h.(EventHandler), table, true)
		}
	}
}

// - 注册读(R)事件

func (d *DBI) OnRead(event string, h EventReadHandler, tableName ...string) {
	var table string
	if len(tableName) > 0 {
		table = tableName[0]
	} else {
		set := strings.SplitN(event, ":", 2)
		switch len(set) {
		case 2:
			event = set[1]
			fallthrough
		case 1:
			table = set[0]
		}
	}
	for _, evt := range d.ParseEventNames(event) {
		d.Events.OnRead(evt, h, table)
	}
}

func (d *DBI) OnReadAsync(event string, h EventReadHandler, tableName ...string) {
	var table string
	if len(tableName) > 0 {
		table = tableName[0]
	} else {
		set := strings.SplitN(event, ":", 2)
		switch len(set) {
		case 2:
			event = set[1]
			fallthrough
		case 1:
			table = set[0]
		}
	}
	for _, evt := range d.ParseEventNames(event) {
		d.Events.OnRead(evt, h, table, true)
	}
}
