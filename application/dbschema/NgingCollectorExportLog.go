// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"time"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/param"
)

type Slice_NgingCollectorExportLog []*NgingCollectorExportLog

func (s Slice_NgingCollectorExportLog) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingCollectorExportLog) RangeRaw(fn func(m *NgingCollectorExportLog) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingCollectorExportLog) GroupBy(keyField string) map[string][]*NgingCollectorExportLog {
	r := map[string][]*NgingCollectorExportLog{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*NgingCollectorExportLog{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_NgingCollectorExportLog) KeyBy(keyField string) map[string]*NgingCollectorExportLog {
	r := map[string]*NgingCollectorExportLog{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_NgingCollectorExportLog) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_NgingCollectorExportLog) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

func (s Slice_NgingCollectorExportLog) FromList(data interface{}) Slice_NgingCollectorExportLog {
	values, ok := data.([]*NgingCollectorExportLog)
	if !ok {
		for _, value := range data.([]interface{}) {
			row := &NgingCollectorExportLog{}
			row.FromRow(value.(map[string]interface{}))
			s = append(s, row)
		}
		return s
	}
	s = append(s, values...)

	return s
}

func NewNgingCollectorExportLog(ctx echo.Context) *NgingCollectorExportLog {
	m := &NgingCollectorExportLog{}
	m.SetContext(ctx)
	return m
}

// NgingCollectorExportLog 导出日志
type NgingCollectorExportLog struct {
	base    factory.Base
	objects []*NgingCollectorExportLog

	Id       uint64 `db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	PageId   uint   `db:"page_id" bson:"page_id" comment:"页面规则ID" json:"page_id" xml:"page_id"`
	ExportId uint   `db:"export_id" bson:"export_id" comment:"导出方案ID" json:"export_id" xml:"export_id"`
	Result   string `db:"result" bson:"result" comment:"结果" json:"result" xml:"result"`
	Status   string `db:"status" bson:"status" comment:"状态" json:"status" xml:"status"`
	Created  uint   `db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
}

// - base function

func (a *NgingCollectorExportLog) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *NgingCollectorExportLog) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *NgingCollectorExportLog) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *NgingCollectorExportLog) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *NgingCollectorExportLog) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *NgingCollectorExportLog) Context() echo.Context {
	return a.base.Context()
}

func (a *NgingCollectorExportLog) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *NgingCollectorExportLog) SetNamer(namer func(factory.Model) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *NgingCollectorExportLog) Namer() func(factory.Model) string {
	return a.base.Namer()
}

func (a *NgingCollectorExportLog) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *NgingCollectorExportLog) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

// - current function

func (a *NgingCollectorExportLog) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *NgingCollectorExportLog) Objects() []*NgingCollectorExportLog {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *NgingCollectorExportLog) XObjects() Slice_NgingCollectorExportLog {
	return Slice_NgingCollectorExportLog(a.Objects())
}

func (a *NgingCollectorExportLog) NewObjects() factory.Ranger {
	return &Slice_NgingCollectorExportLog{}
}

func (a *NgingCollectorExportLog) InitObjects() *[]*NgingCollectorExportLog {
	a.objects = []*NgingCollectorExportLog{}
	return &a.objects
}

func (a *NgingCollectorExportLog) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *NgingCollectorExportLog) Short_() string {
	return "nging_collector_export_log"
}

func (a *NgingCollectorExportLog) Struct_() string {
	return "NgingCollectorExportLog"
}

func (a *NgingCollectorExportLog) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *NgingCollectorExportLog) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *NgingCollectorExportLog) Get(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	base := a.base
	if !a.base.Eventable() {
		err = a.Param(mw, args...).SetRecv(a).One()
		a.base = base
		return
	}
	queryParam := a.Param(mw, args...).SetRecv(a)
	if err = DBI.FireReading(a, queryParam); err != nil {
		return
	}
	err = queryParam.One()
	a.base = base
	if err == nil {
		err = DBI.FireReaded(a, queryParam)
	}
	return
}

func (a *NgingCollectorExportLog) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv).List()
	}
	queryParam := a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv)
	if err := DBI.FireReading(a, queryParam); err != nil {
		return nil, err
	}
	cnt, err := queryParam.List()
	if err == nil {
		switch v := recv.(type) {
		case *[]*NgingCollectorExportLog:
			err = DBI.FireReaded(a, queryParam, Slice_NgingCollectorExportLog(*v))
		case []*NgingCollectorExportLog:
			err = DBI.FireReaded(a, queryParam, Slice_NgingCollectorExportLog(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingCollectorExportLog) GroupBy(keyField string, inputRows ...[]*NgingCollectorExportLog) map[string][]*NgingCollectorExportLog {
	var rows Slice_NgingCollectorExportLog
	if len(inputRows) > 0 {
		rows = Slice_NgingCollectorExportLog(inputRows[0])
	} else {
		rows = Slice_NgingCollectorExportLog(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *NgingCollectorExportLog) KeyBy(keyField string, inputRows ...[]*NgingCollectorExportLog) map[string]*NgingCollectorExportLog {
	var rows Slice_NgingCollectorExportLog
	if len(inputRows) > 0 {
		rows = Slice_NgingCollectorExportLog(inputRows[0])
	} else {
		rows = Slice_NgingCollectorExportLog(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *NgingCollectorExportLog) AsKV(keyField string, valueField string, inputRows ...[]*NgingCollectorExportLog) param.Store {
	var rows Slice_NgingCollectorExportLog
	if len(inputRows) > 0 {
		rows = Slice_NgingCollectorExportLog(inputRows[0])
	} else {
		rows = Slice_NgingCollectorExportLog(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *NgingCollectorExportLog) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv).List()
	}
	queryParam := a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv)
	if err := DBI.FireReading(a, queryParam); err != nil {
		return nil, err
	}
	cnt, err := queryParam.List()
	if err == nil {
		switch v := recv.(type) {
		case *[]*NgingCollectorExportLog:
			err = DBI.FireReaded(a, queryParam, Slice_NgingCollectorExportLog(*v))
		case []*NgingCollectorExportLog:
			err = DBI.FireReaded(a, queryParam, Slice_NgingCollectorExportLog(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingCollectorExportLog) Add() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	if len(a.Status) == 0 {
		a.Status = "idle"
	}
	if a.base.Eventable() {
		err = DBI.Fire("creating", a, nil)
		if err != nil {
			return
		}
	}
	pk, err = a.Param(nil).SetSend(a).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint64(v)
		}
	}
	if err == nil && a.base.Eventable() {
		err = DBI.Fire("created", a, nil)
	}
	return
}

func (a *NgingCollectorExportLog) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if len(a.Status) == 0 {
		a.Status = "idle"
	}
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

func (a *NgingCollectorExportLog) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingCollectorExportLog) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["status"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["status"] = "idle"
		}
	}
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

func (a *NgingCollectorExportLog) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		if len(a.Status) == 0 {
			a.Status = "idle"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
		if len(a.Status) == 0 {
			a.Status = "idle"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("creating", a, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint64(v)
		}
	}
	if err == nil && a.base.Eventable() {
		if pk == nil {
			err = DBI.Fire("updated", a, mw, args...)
		} else {
			err = DBI.Fire("created", a, nil)
		}
	}
	return
}

func (a *NgingCollectorExportLog) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

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

func (a *NgingCollectorExportLog) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *NgingCollectorExportLog) Exists(mw func(db.Result) db.Result, args ...interface{}) (bool, error) {
	return a.Param(mw, args...).Exists()
}

func (a *NgingCollectorExportLog) Reset() *NgingCollectorExportLog {
	a.Id = 0
	a.PageId = 0
	a.ExportId = 0
	a.Result = ``
	a.Status = ``
	a.Created = 0
	return a
}

func (a *NgingCollectorExportLog) AsMap(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["Id"] = a.Id
		r["PageId"] = a.PageId
		r["ExportId"] = a.ExportId
		r["Result"] = a.Result
		r["Status"] = a.Status
		r["Created"] = a.Created
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "Id":
			r["Id"] = a.Id
		case "PageId":
			r["PageId"] = a.PageId
		case "ExportId":
			r["ExportId"] = a.ExportId
		case "Result":
			r["Result"] = a.Result
		case "Status":
			r["Status"] = a.Status
		case "Created":
			r["Created"] = a.Created
		}
	}
	return r
}

func (a *NgingCollectorExportLog) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint64(value)
		case "page_id":
			a.PageId = param.AsUint(value)
		case "export_id":
			a.ExportId = param.AsUint(value)
		case "result":
			a.Result = param.AsString(value)
		case "status":
			a.Status = param.AsString(value)
		case "created":
			a.Created = param.AsUint(value)
		}
	}
}

func (a *NgingCollectorExportLog) Set(key interface{}, value ...interface{}) {
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
		case "Id":
			a.Id = param.AsUint64(vv)
		case "PageId":
			a.PageId = param.AsUint(vv)
		case "ExportId":
			a.ExportId = param.AsUint(vv)
		case "Result":
			a.Result = param.AsString(vv)
		case "Status":
			a.Status = param.AsString(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		}
	}
}

func (a *NgingCollectorExportLog) AsRow(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["id"] = a.Id
		r["page_id"] = a.PageId
		r["export_id"] = a.ExportId
		r["result"] = a.Result
		r["status"] = a.Status
		r["created"] = a.Created
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "id":
			r["id"] = a.Id
		case "page_id":
			r["page_id"] = a.PageId
		case "export_id":
			r["export_id"] = a.ExportId
		case "result":
			r["result"] = a.Result
		case "status":
			r["status"] = a.Status
		case "created":
			r["created"] = a.Created
		}
	}
	return r
}

func (a *NgingCollectorExportLog) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return DBI.Fields.BatchValidate(a.Short_(), kvset)
}

func (a *NgingCollectorExportLog) Validate(field string, value interface{}) error {
	return DBI.Fields.Validate(a.Short_(), field, value)
}
