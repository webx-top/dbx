//Do not edit this file, which is automatically generated by the generator.
package dbschema

import (
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	
	"time"
)

type User struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*User
	
	Id        	uint    	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"UID" json:"id" xml:"id"`
	Uname     	string  	`db:"uname" bson:"uname" comment:"用户名" json:"uname" xml:"uname"`
	Passwd    	string  	`db:"passwd" bson:"passwd" comment:"密码" json:"passwd" xml:"passwd"`
	Salt      	string  	`db:"salt" bson:"salt" comment:"盐值" json:"salt" xml:"salt"`
	Email     	string  	`db:"email" bson:"email" comment:"邮箱" json:"email" xml:"email"`
	Mobile    	string  	`db:"mobile" bson:"mobile" comment:"手机号" json:"mobile" xml:"mobile"`
	LoginTime 	uint    	`db:"login_time" bson:"login_time" comment:"上次登录时间" json:"login_time" xml:"login_time"`
	LoginIp   	string  	`db:"login_ip" bson:"login_ip" comment:"上次登录IP" json:"login_ip" xml:"login_ip"`
	Created   	uint    	`db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Updated   	uint    	`db:"updated" bson:"updated" comment:"更新时间" json:"updated" xml:"updated"`
	Active    	string  	`db:"active" bson:"active" comment:"激活" json:"active" xml:"active"`
	Avatar    	string  	`db:"avatar" bson:"avatar" comment:"头像" json:"avatar" xml:"avatar"`
}

func (this *User) Trans() *factory.Transaction {
	return this.trans
}

func (this *User) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *User) Objects() []*User {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *User) NewObjects() *[]*User {
	this.objects = []*User{}
	return &this.objects
}

func (this *User) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetTrans(this.trans).SetCollection("user").SetModel(this)
}

func (this *User) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *User) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *User) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *User) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *User) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *User) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	pk, err = this.Param().SetSend(this).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		}
	}
	return
}

func (this *User) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	this.Updated = uint(time.Now().Unix())
	return this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Update()
}

func (this *User) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		this.Updated = uint(time.Now().Unix())
	},func(){
		this.Created = uint(time.Now().Unix())
	this.Id = 0
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		}
	}
	return 
}

func (this *User) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetMiddleware(mw).Delete()
}

