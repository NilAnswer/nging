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

type Slice_NgingCloudBackup []*NgingCloudBackup

func (s Slice_NgingCloudBackup) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingCloudBackup) RangeRaw(fn func(m *NgingCloudBackup) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingCloudBackup) GroupBy(keyField string) map[string][]*NgingCloudBackup {
	r := map[string][]*NgingCloudBackup{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*NgingCloudBackup{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_NgingCloudBackup) KeyBy(keyField string) map[string]*NgingCloudBackup {
	r := map[string]*NgingCloudBackup{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_NgingCloudBackup) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_NgingCloudBackup) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

func (s Slice_NgingCloudBackup) FromList(data interface{}) Slice_NgingCloudBackup {
	values, ok := data.([]*NgingCloudBackup)
	if !ok {
		for _, value := range data.([]interface{}) {
			row := &NgingCloudBackup{}
			row.FromRow(value.(map[string]interface{}))
			s = append(s, row)
		}
		return s
	}
	s = append(s, values...)

	return s
}

func NewNgingCloudBackup(ctx echo.Context) *NgingCloudBackup {
	m := &NgingCloudBackup{}
	m.SetContext(ctx)
	return m
}

// NgingCloudBackup 云备份
type NgingCloudBackup struct {
	base    factory.Base
	objects []*NgingCloudBackup

	Id                uint   `db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Name              string `db:"name" bson:"name" comment:"配置名" json:"name" xml:"name"`
	SourcePath        string `db:"source_path" bson:"source_path" comment:"源" json:"source_path" xml:"source_path"`
	IgnoreRule        string `db:"ignore_rule" bson:"ignore_rule" comment:"忽略文件路径(正则表达式)" json:"ignore_rule" xml:"ignore_rule"`
	WaitFillCompleted string `db:"wait_fill_completed" bson:"wait_fill_completed" comment:"是否等待文件填充结束" json:"wait_fill_completed" xml:"wait_fill_completed"`
	IgnoreWaitRule    string `db:"ignore_wait_rule" bson:"ignore_wait_rule" comment:"忽略等待文件完成的规则" json:"ignore_wait_rule" xml:"ignore_wait_rule"`
	Delay             uint   `db:"delay" bson:"delay" comment:"延后秒数" json:"delay" xml:"delay"`
	DestStorage       uint   `db:"dest_storage" bson:"dest_storage" comment:"目标存储ID" json:"dest_storage" xml:"dest_storage"`
	DestPath          string `db:"dest_path" bson:"dest_path" comment:"目标存储路径" json:"dest_path" xml:"dest_path"`
	Result            string `db:"result" bson:"result" comment:"运行结果" json:"result" xml:"result"`
	LastExecuted      uint   `db:"last_executed" bson:"last_executed" comment:"最近运行时间" json:"last_executed" xml:"last_executed"`
	Status            string `db:"status" bson:"status" comment:"运行状态" json:"status" xml:"status"`
	Disabled          string `db:"disabled" bson:"disabled" comment:"是否(Y/N)禁用" json:"disabled" xml:"disabled"`
	Created           uint   `db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Updated           uint   `db:"updated" bson:"updated" comment:"更新时间" json:"updated" xml:"updated"`
}

// - base function

func (a *NgingCloudBackup) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *NgingCloudBackup) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *NgingCloudBackup) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *NgingCloudBackup) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *NgingCloudBackup) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *NgingCloudBackup) Context() echo.Context {
	return a.base.Context()
}

func (a *NgingCloudBackup) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *NgingCloudBackup) SetNamer(namer func(factory.Model) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *NgingCloudBackup) Namer() func(factory.Model) string {
	return a.base.Namer()
}

func (a *NgingCloudBackup) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *NgingCloudBackup) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

// - current function

func (a *NgingCloudBackup) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *NgingCloudBackup) Objects() []*NgingCloudBackup {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *NgingCloudBackup) XObjects() Slice_NgingCloudBackup {
	return Slice_NgingCloudBackup(a.Objects())
}

func (a *NgingCloudBackup) NewObjects() factory.Ranger {
	return &Slice_NgingCloudBackup{}
}

func (a *NgingCloudBackup) InitObjects() *[]*NgingCloudBackup {
	a.objects = []*NgingCloudBackup{}
	return &a.objects
}

func (a *NgingCloudBackup) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *NgingCloudBackup) Short_() string {
	return "nging_cloud_backup"
}

func (a *NgingCloudBackup) Struct_() string {
	return "NgingCloudBackup"
}

func (a *NgingCloudBackup) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *NgingCloudBackup) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *NgingCloudBackup) Get(mw func(db.Result) db.Result, args ...interface{}) (err error) {
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

func (a *NgingCloudBackup) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingCloudBackup:
			err = DBI.FireReaded(a, queryParam, Slice_NgingCloudBackup(*v))
		case []*NgingCloudBackup:
			err = DBI.FireReaded(a, queryParam, Slice_NgingCloudBackup(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingCloudBackup) GroupBy(keyField string, inputRows ...[]*NgingCloudBackup) map[string][]*NgingCloudBackup {
	var rows Slice_NgingCloudBackup
	if len(inputRows) > 0 {
		rows = Slice_NgingCloudBackup(inputRows[0])
	} else {
		rows = Slice_NgingCloudBackup(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *NgingCloudBackup) KeyBy(keyField string, inputRows ...[]*NgingCloudBackup) map[string]*NgingCloudBackup {
	var rows Slice_NgingCloudBackup
	if len(inputRows) > 0 {
		rows = Slice_NgingCloudBackup(inputRows[0])
	} else {
		rows = Slice_NgingCloudBackup(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *NgingCloudBackup) AsKV(keyField string, valueField string, inputRows ...[]*NgingCloudBackup) param.Store {
	var rows Slice_NgingCloudBackup
	if len(inputRows) > 0 {
		rows = Slice_NgingCloudBackup(inputRows[0])
	} else {
		rows = Slice_NgingCloudBackup(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *NgingCloudBackup) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingCloudBackup:
			err = DBI.FireReaded(a, queryParam, Slice_NgingCloudBackup(*v))
		case []*NgingCloudBackup:
			err = DBI.FireReaded(a, queryParam, Slice_NgingCloudBackup(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingCloudBackup) Add() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	if len(a.WaitFillCompleted) == 0 {
		a.WaitFillCompleted = "N"
	}
	if len(a.Status) == 0 {
		a.Status = "idle"
	}
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if a.base.Eventable() {
		err = DBI.Fire("creating", a, nil)
		if err != nil {
			return
		}
	}
	pk, err = a.Param(nil).SetSend(a).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint(v)
		}
	}
	if err == nil && a.base.Eventable() {
		err = DBI.Fire("created", a, nil)
	}
	return
}

func (a *NgingCloudBackup) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.WaitFillCompleted) == 0 {
		a.WaitFillCompleted = "N"
	}
	if len(a.Status) == 0 {
		a.Status = "idle"
	}
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
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

func (a *NgingCloudBackup) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingCloudBackup) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["wait_fill_completed"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["wait_fill_completed"] = "N"
		}
	}
	if val, ok := kvset["status"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["status"] = "idle"
		}
	}
	if val, ok := kvset["disabled"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["disabled"] = "N"
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

func (a *NgingCloudBackup) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		a.Updated = uint(time.Now().Unix())
		if len(a.WaitFillCompleted) == 0 {
			a.WaitFillCompleted = "N"
		}
		if len(a.Status) == 0 {
			a.Status = "idle"
		}
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
		if len(a.WaitFillCompleted) == 0 {
			a.WaitFillCompleted = "N"
		}
		if len(a.Status) == 0 {
			a.Status = "idle"
		}
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("creating", a, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint(v)
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

func (a *NgingCloudBackup) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

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

func (a *NgingCloudBackup) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *NgingCloudBackup) Exists(mw func(db.Result) db.Result, args ...interface{}) (bool, error) {
	return a.Param(mw, args...).Exists()
}

func (a *NgingCloudBackup) Reset() *NgingCloudBackup {
	a.Id = 0
	a.Name = ``
	a.SourcePath = ``
	a.IgnoreRule = ``
	a.WaitFillCompleted = ``
	a.IgnoreWaitRule = ``
	a.Delay = 0
	a.DestStorage = 0
	a.DestPath = ``
	a.Result = ``
	a.LastExecuted = 0
	a.Status = ``
	a.Disabled = ``
	a.Created = 0
	a.Updated = 0
	return a
}

func (a *NgingCloudBackup) AsMap(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["Id"] = a.Id
		r["Name"] = a.Name
		r["SourcePath"] = a.SourcePath
		r["IgnoreRule"] = a.IgnoreRule
		r["WaitFillCompleted"] = a.WaitFillCompleted
		r["IgnoreWaitRule"] = a.IgnoreWaitRule
		r["Delay"] = a.Delay
		r["DestStorage"] = a.DestStorage
		r["DestPath"] = a.DestPath
		r["Result"] = a.Result
		r["LastExecuted"] = a.LastExecuted
		r["Status"] = a.Status
		r["Disabled"] = a.Disabled
		r["Created"] = a.Created
		r["Updated"] = a.Updated
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "Id":
			r["Id"] = a.Id
		case "Name":
			r["Name"] = a.Name
		case "SourcePath":
			r["SourcePath"] = a.SourcePath
		case "IgnoreRule":
			r["IgnoreRule"] = a.IgnoreRule
		case "WaitFillCompleted":
			r["WaitFillCompleted"] = a.WaitFillCompleted
		case "IgnoreWaitRule":
			r["IgnoreWaitRule"] = a.IgnoreWaitRule
		case "Delay":
			r["Delay"] = a.Delay
		case "DestStorage":
			r["DestStorage"] = a.DestStorage
		case "DestPath":
			r["DestPath"] = a.DestPath
		case "Result":
			r["Result"] = a.Result
		case "LastExecuted":
			r["LastExecuted"] = a.LastExecuted
		case "Status":
			r["Status"] = a.Status
		case "Disabled":
			r["Disabled"] = a.Disabled
		case "Created":
			r["Created"] = a.Created
		case "Updated":
			r["Updated"] = a.Updated
		}
	}
	return r
}

func (a *NgingCloudBackup) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint(value)
		case "name":
			a.Name = param.AsString(value)
		case "source_path":
			a.SourcePath = param.AsString(value)
		case "ignore_rule":
			a.IgnoreRule = param.AsString(value)
		case "wait_fill_completed":
			a.WaitFillCompleted = param.AsString(value)
		case "ignore_wait_rule":
			a.IgnoreWaitRule = param.AsString(value)
		case "delay":
			a.Delay = param.AsUint(value)
		case "dest_storage":
			a.DestStorage = param.AsUint(value)
		case "dest_path":
			a.DestPath = param.AsString(value)
		case "result":
			a.Result = param.AsString(value)
		case "last_executed":
			a.LastExecuted = param.AsUint(value)
		case "status":
			a.Status = param.AsString(value)
		case "disabled":
			a.Disabled = param.AsString(value)
		case "created":
			a.Created = param.AsUint(value)
		case "updated":
			a.Updated = param.AsUint(value)
		}
	}
}

func (a *NgingCloudBackup) Set(key interface{}, value ...interface{}) {
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
			a.Id = param.AsUint(vv)
		case "Name":
			a.Name = param.AsString(vv)
		case "SourcePath":
			a.SourcePath = param.AsString(vv)
		case "IgnoreRule":
			a.IgnoreRule = param.AsString(vv)
		case "WaitFillCompleted":
			a.WaitFillCompleted = param.AsString(vv)
		case "IgnoreWaitRule":
			a.IgnoreWaitRule = param.AsString(vv)
		case "Delay":
			a.Delay = param.AsUint(vv)
		case "DestStorage":
			a.DestStorage = param.AsUint(vv)
		case "DestPath":
			a.DestPath = param.AsString(vv)
		case "Result":
			a.Result = param.AsString(vv)
		case "LastExecuted":
			a.LastExecuted = param.AsUint(vv)
		case "Status":
			a.Status = param.AsString(vv)
		case "Disabled":
			a.Disabled = param.AsString(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "Updated":
			a.Updated = param.AsUint(vv)
		}
	}
}

func (a *NgingCloudBackup) AsRow(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["id"] = a.Id
		r["name"] = a.Name
		r["source_path"] = a.SourcePath
		r["ignore_rule"] = a.IgnoreRule
		r["wait_fill_completed"] = a.WaitFillCompleted
		r["ignore_wait_rule"] = a.IgnoreWaitRule
		r["delay"] = a.Delay
		r["dest_storage"] = a.DestStorage
		r["dest_path"] = a.DestPath
		r["result"] = a.Result
		r["last_executed"] = a.LastExecuted
		r["status"] = a.Status
		r["disabled"] = a.Disabled
		r["created"] = a.Created
		r["updated"] = a.Updated
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "id":
			r["id"] = a.Id
		case "name":
			r["name"] = a.Name
		case "source_path":
			r["source_path"] = a.SourcePath
		case "ignore_rule":
			r["ignore_rule"] = a.IgnoreRule
		case "wait_fill_completed":
			r["wait_fill_completed"] = a.WaitFillCompleted
		case "ignore_wait_rule":
			r["ignore_wait_rule"] = a.IgnoreWaitRule
		case "delay":
			r["delay"] = a.Delay
		case "dest_storage":
			r["dest_storage"] = a.DestStorage
		case "dest_path":
			r["dest_path"] = a.DestPath
		case "result":
			r["result"] = a.Result
		case "last_executed":
			r["last_executed"] = a.LastExecuted
		case "status":
			r["status"] = a.Status
		case "disabled":
			r["disabled"] = a.Disabled
		case "created":
			r["created"] = a.Created
		case "updated":
			r["updated"] = a.Updated
		}
	}
	return r
}

func (a *NgingCloudBackup) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return DBI.Fields.BatchValidate(a.Short_(), kvset)
}

func (a *NgingCloudBackup) Validate(field string, value interface{}) error {
	return DBI.Fields.Validate(a.Short_(), field, value)
}
