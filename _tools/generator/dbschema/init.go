//Do not edit this file, which is automatically generated by the generator.
package dbschema

import (
	"github.com/webx-top/db/lib/factory"
)

func init(){
	factory.Fields=map[string]map[string]*factory.FieldInfo{"config":map[string]*factory.FieldInfo{"updated":&factory.FieldInfo{Name:"updated", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"更新时间", GoType:"uint", GoName:"Updated"}, "id":&factory.FieldInfo{Name:"id", DataType:"int", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"主键ID", GoType:"uint", GoName:"Id"}, "key":&factory.FieldInfo{Name:"key", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:60, Options:[]string{}, DefaultValue:"", Comment:"配置项", GoType:"string", GoName:"Key"}, "val":&factory.FieldInfo{Name:"val", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:200, Options:[]string{}, DefaultValue:"", Comment:"配置值", GoType:"string", GoName:"Val"}}, "link":map[string]*factory.FieldInfo{"logo":&factory.FieldInfo{Name:"logo", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:200, Options:[]string{}, DefaultValue:"", Comment:"LOGO", GoType:"string", GoName:"Logo"}, "show":&factory.FieldInfo{Name:"show", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"Y", "N"}, DefaultValue:"N", Comment:"是否显示", GoType:"string", GoName:"Show"}, "created":&factory.FieldInfo{Name:"created", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"创建时间", GoType:"uint", GoName:"Created"}, "updated":&factory.FieldInfo{Name:"updated", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"更新时间", GoType:"uint", GoName:"Updated"}, "sort":&factory.FieldInfo{Name:"sort", DataType:"int", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"排序", GoType:"int", GoName:"Sort"}, "id":&factory.FieldInfo{Name:"id", DataType:"int", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"主键ID", GoType:"uint", GoName:"Id"}, "name":&factory.FieldInfo{Name:"name", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:30, Options:[]string{}, DefaultValue:"", Comment:"名称", GoType:"string", GoName:"Name"}, "url":&factory.FieldInfo{Name:"url", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:200, Options:[]string{}, DefaultValue:"", Comment:"网址", GoType:"string", GoName:"Url"}, "verified":&factory.FieldInfo{Name:"verified", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"验证时间", GoType:"uint", GoName:"Verified"}, "catid":&factory.FieldInfo{Name:"catid", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"分类", GoType:"uint", GoName:"Catid"}}, "ocontent":map[string]*factory.FieldInfo{"content":&factory.FieldInfo{Name:"content", DataType:"", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{}, DefaultValue:"", Comment:"博客原始内容", GoType:"string", GoName:"Content"}, "etype":&factory.FieldInfo{Name:"etype", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"markdown"}, DefaultValue:"markdown", Comment:"编辑器类型", GoType:"string", GoName:"Etype"}, "id":&factory.FieldInfo{Name:"id", DataType:"int", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"ID", GoType:"uint", GoName:"Id"}, "rc_id":&factory.FieldInfo{Name:"rc_id", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"关联ID", GoType:"uint", GoName:"RcId"}, "rc_type":&factory.FieldInfo{Name:"rc_type", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:30, Options:[]string{}, DefaultValue:"post", Comment:"关联类型", GoType:"string", GoName:"RcType"}}, "album":map[string]*factory.FieldInfo{"created":&factory.FieldInfo{Name:"created", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"创建时间", GoType:"uint", GoName:"Created"}, "comments":&factory.FieldInfo{Name:"comments", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"评论次数", GoType:"uint", GoName:"Comments"}, "display":&factory.FieldInfo{Name:"display", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"ALL", "SELF", "FRIEND", "PWD"}, DefaultValue:"ALL", Comment:"显示", GoType:"string", GoName:"Display"}, "deleted":&factory.FieldInfo{Name:"deleted", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"删除时间", GoType:"uint", GoName:"Deleted"}, "content":&factory.FieldInfo{Name:"content", DataType:"", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{}, DefaultValue:"", Comment:"正文", GoType:"string", GoName:"Content"}, "tags":&factory.FieldInfo{Name:"tags", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:255, Options:[]string{}, DefaultValue:"", Comment:"标签", GoType:"string", GoName:"Tags"}, "likes":&factory.FieldInfo{Name:"likes", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"喜欢次数", GoType:"uint", GoName:"Likes"}, "title":&factory.FieldInfo{Name:"title", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:180, Options:[]string{}, DefaultValue:"", Comment:"标题", GoType:"string", GoName:"Title"}, "catid":&factory.FieldInfo{Name:"catid", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"分类ID", GoType:"uint", GoName:"Catid"}, "id":&factory.FieldInfo{Name:"id", DataType:"int", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"ID", GoType:"uint", GoName:"Id"}, "updated":&factory.FieldInfo{Name:"updated", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"编辑时间", GoType:"uint", GoName:"Updated"}, "views":&factory.FieldInfo{Name:"views", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"浏览次数", GoType:"uint", GoName:"Views"}, "allow_comment":&factory.FieldInfo{Name:"allow_comment", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"Y", "N"}, DefaultValue:"Y", Comment:"是否允许评论", GoType:"string", GoName:"AllowComment"}, "description":&factory.FieldInfo{Name:"description", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:200, Options:[]string{}, DefaultValue:"", Comment:"简介", GoType:"string", GoName:"Description"}}, "attathment":map[string]*factory.FieldInfo{"path":&factory.FieldInfo{Name:"path", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:255, Options:[]string{}, DefaultValue:"", Comment:"保存路径", GoType:"string", GoName:"Path"}, "size":&factory.FieldInfo{Name:"size", DataType:"bigint", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:20, Options:[]string{}, DefaultValue:"", Comment:"文件尺寸", GoType:"uint64", GoName:"Size"}, "uid":&factory.FieldInfo{Name:"uid", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"UID", GoType:"uint", GoName:"Uid"}, "deleted":&factory.FieldInfo{Name:"deleted", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"被删除时间", GoType:"uint", GoName:"Deleted"}, "rc_id":&factory.FieldInfo{Name:"rc_id", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"关联id", GoType:"uint", GoName:"RcId"}, "rc_type":&factory.FieldInfo{Name:"rc_type", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:30, Options:[]string{}, DefaultValue:"", Comment:"关联类型", GoType:"string", GoName:"RcType"}, "tags":&factory.FieldInfo{Name:"tags", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:255, Options:[]string{}, DefaultValue:"", Comment:"标签", GoType:"string", GoName:"Tags"}, "id":&factory.FieldInfo{Name:"id", DataType:"int", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"ID", GoType:"uint", GoName:"Id"}, "name":&factory.FieldInfo{Name:"name", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:100, Options:[]string{}, DefaultValue:"", Comment:"文件名", GoType:"string", GoName:"Name"}, "extension":&factory.FieldInfo{Name:"extension", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:5, Options:[]string{}, DefaultValue:"", Comment:"扩展名", GoType:"string", GoName:"Extension"}, "type":&factory.FieldInfo{Name:"type", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"image", "media", "other"}, DefaultValue:"image", Comment:"文件类型", GoType:"string", GoName:"Type"}, "created":&factory.FieldInfo{Name:"created", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"创建时间", GoType:"uint", GoName:"Created"}, "audited":&factory.FieldInfo{Name:"audited", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"审核时间", GoType:"uint", GoName:"Audited"}}, "category":map[string]*factory.FieldInfo{"updated":&factory.FieldInfo{Name:"updated", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"更新时间", GoType:"uint", GoName:"Updated"}, "sort":&factory.FieldInfo{Name:"sort", DataType:"int", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"排序", GoType:"int", GoName:"Sort"}, "tmpl":&factory.FieldInfo{Name:"tmpl", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:100, Options:[]string{}, DefaultValue:"", Comment:"模板", GoType:"string", GoName:"Tmpl"}, "id":&factory.FieldInfo{Name:"id", DataType:"int", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"ID", GoType:"uint", GoName:"Id"}, "pid":&factory.FieldInfo{Name:"pid", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"上级分类", GoType:"uint", GoName:"Pid"}, "haschild":&factory.FieldInfo{Name:"haschild", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"Y", "N"}, DefaultValue:"N", Comment:"是否有子分类", GoType:"string", GoName:"Haschild"}, "name":&factory.FieldInfo{Name:"name", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:30, Options:[]string{}, DefaultValue:"", Comment:"分类名称", GoType:"string", GoName:"Name"}, "description":&factory.FieldInfo{Name:"description", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:200, Options:[]string{}, DefaultValue:"", Comment:"说明", GoType:"string", GoName:"Description"}, "rc_type":&factory.FieldInfo{Name:"rc_type", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:30, Options:[]string{}, DefaultValue:"post", Comment:"关联类型", GoType:"string", GoName:"RcType"}}, "user":map[string]*factory.FieldInfo{"id":&factory.FieldInfo{Name:"id", DataType:"int", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"UID", GoType:"uint", GoName:"Id"}, "uname":&factory.FieldInfo{Name:"uname", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:30, Options:[]string{}, DefaultValue:"", Comment:"用户名", GoType:"string", GoName:"Uname"}, "passwd":&factory.FieldInfo{Name:"passwd", DataType:"char", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:64, Options:[]string{}, DefaultValue:"", Comment:"密码", GoType:"string", GoName:"Passwd"}, "login_time":&factory.FieldInfo{Name:"login_time", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"上次登录时间", GoType:"uint", GoName:"LoginTime"}, "avatar":&factory.FieldInfo{Name:"avatar", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:200, Options:[]string{}, DefaultValue:"", Comment:"头像", GoType:"string", GoName:"Avatar"}, "salt":&factory.FieldInfo{Name:"salt", DataType:"char", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:64, Options:[]string{}, DefaultValue:"", Comment:"盐值", GoType:"string", GoName:"Salt"}, "email":&factory.FieldInfo{Name:"email", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:100, Options:[]string{}, DefaultValue:"", Comment:"邮箱", GoType:"string", GoName:"Email"}, "mobile":&factory.FieldInfo{Name:"mobile", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:15, Options:[]string{}, DefaultValue:"", Comment:"手机号", GoType:"string", GoName:"Mobile"}, "login_ip":&factory.FieldInfo{Name:"login_ip", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:40, Options:[]string{}, DefaultValue:"", Comment:"上次登录IP", GoType:"string", GoName:"LoginIp"}, "created":&factory.FieldInfo{Name:"created", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"创建时间", GoType:"uint", GoName:"Created"}, "updated":&factory.FieldInfo{Name:"updated", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"更新时间", GoType:"uint", GoName:"Updated"}, "active":&factory.FieldInfo{Name:"active", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"Y", "N"}, DefaultValue:"Y", Comment:"激活", GoType:"string", GoName:"Active"}}, "comment":map[string]*factory.FieldInfo{"uid":&factory.FieldInfo{Name:"uid", DataType:"bigint", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:20, Options:[]string{}, DefaultValue:"0", Comment:"发布者id", GoType:"uint64", GoName:"Uid"}, "uname":&factory.FieldInfo{Name:"uname", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:30, Options:[]string{}, DefaultValue:"", Comment:"发布者用户名", GoType:"string", GoName:"Uname"}, "rc_id":&factory.FieldInfo{Name:"rc_id", DataType:"bigint", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:20, Options:[]string{}, DefaultValue:"0", Comment:"关联内容ID", GoType:"uint64", GoName:"RcId"}, "for_uname":&factory.FieldInfo{Name:"for_uname", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:30, Options:[]string{}, DefaultValue:"", Comment:"被评人用户名", GoType:"string", GoName:"ForUname"}, "r_type":&factory.FieldInfo{Name:"r_type", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"reply", "append"}, DefaultValue:"", Comment:"关联类型（reply-回复，append-追加）", GoType:"string", GoName:"RType"}, "down":&factory.FieldInfo{Name:"down", DataType:"bigint", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:20, Options:[]string{}, DefaultValue:"0", Comment:"被踩次数", GoType:"uint64", GoName:"Down"}, "updated":&factory.FieldInfo{Name:"updated", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"更新时间", GoType:"uint", GoName:"Updated"}, "for_uid":&factory.FieldInfo{Name:"for_uid", DataType:"bigint", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:20, Options:[]string{}, DefaultValue:"0", Comment:"被评人id", GoType:"uint64", GoName:"ForUid"}, "quote":&factory.FieldInfo{Name:"quote", DataType:"", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{}, DefaultValue:"", Comment:"引用内容", GoType:"string", GoName:"Quote"}, "etype":&factory.FieldInfo{Name:"etype", DataType:"char", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"html", Comment:"编辑器类型", GoType:"string", GoName:"Etype"}, "root_id":&factory.FieldInfo{Name:"root_id", DataType:"bigint", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:20, Options:[]string{}, DefaultValue:"0", Comment:"", GoType:"uint64", GoName:"RootId"}, "r_id":&factory.FieldInfo{Name:"r_id", DataType:"bigint", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:20, Options:[]string{}, DefaultValue:"0", Comment:"关联本表的id", GoType:"uint64", GoName:"RId"}, "content":&factory.FieldInfo{Name:"content", DataType:"", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{}, DefaultValue:"", Comment:"内容", GoType:"string", GoName:"Content"}, "root_times":&factory.FieldInfo{Name:"root_times", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"根节点下的所有回复次数", GoType:"uint", GoName:"RootTimes"}, "status":&factory.FieldInfo{Name:"status", DataType:"int", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:1, Options:[]string{}, DefaultValue:"0", Comment:"状态", GoType:"int", GoName:"Status"}, "rc_type":&factory.FieldInfo{Name:"rc_type", DataType:"char", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:30, Options:[]string{}, DefaultValue:"post", Comment:"关联内容类型", GoType:"string", GoName:"RcType"}, "id":&factory.FieldInfo{Name:"id", DataType:"bigint", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:20, Options:[]string{}, DefaultValue:"", Comment:"主键", GoType:"uint64", GoName:"Id"}, "related_times":&factory.FieldInfo{Name:"related_times", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"本身被回复次数", GoType:"uint", GoName:"RelatedTimes"}, "up":&factory.FieldInfo{Name:"up", DataType:"bigint", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:20, Options:[]string{}, DefaultValue:"0", Comment:"被顶次数", GoType:"uint64", GoName:"Up"}, "created":&factory.FieldInfo{Name:"created", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"创建时间", GoType:"uint", GoName:"Created"}}, "post":map[string]*factory.FieldInfo{"uname":&factory.FieldInfo{Name:"uname", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:30, Options:[]string{}, DefaultValue:"", Comment:"作者名", GoType:"string", GoName:"Uname"}, "comments":&factory.FieldInfo{Name:"comments", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"被评论次数", GoType:"uint", GoName:"Comments"}, "deleted":&factory.FieldInfo{Name:"deleted", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"被删除时间", GoType:"uint", GoName:"Deleted"}, "tags":&factory.FieldInfo{Name:"tags", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:255, Options:[]string{}, DefaultValue:"", Comment:"标签", GoType:"string", GoName:"Tags"}, "content":&factory.FieldInfo{Name:"content", DataType:"", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{}, DefaultValue:"", Comment:"内容", GoType:"string", GoName:"Content"}, "updated":&factory.FieldInfo{Name:"updated", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"修改时间", GoType:"uint", GoName:"Updated"}, "passwd":&factory.FieldInfo{Name:"passwd", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:64, Options:[]string{}, DefaultValue:"", Comment:"访问密码", GoType:"string", GoName:"Passwd"}, "month":&factory.FieldInfo{Name:"month", DataType:"tinyint", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:1, Options:[]string{}, DefaultValue:"", Comment:"归档月份", GoType:"uint", GoName:"Month"}, "title":&factory.FieldInfo{Name:"title", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:180, Options:[]string{}, DefaultValue:"", Comment:"标题", GoType:"string", GoName:"Title"}, "description":&factory.FieldInfo{Name:"description", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:200, Options:[]string{}, DefaultValue:"", Comment:"简介", GoType:"string", GoName:"Description"}, "uid":&factory.FieldInfo{Name:"uid", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"UID", GoType:"uint", GoName:"Uid"}, "views":&factory.FieldInfo{Name:"views", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"被浏览次数", GoType:"uint", GoName:"Views"}, "year":&factory.FieldInfo{Name:"year", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:5, Options:[]string{}, DefaultValue:"", Comment:"归档年份", GoType:"uint", GoName:"Year"}, "etype":&factory.FieldInfo{Name:"etype", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"html", "markdown"}, DefaultValue:"html", Comment:"编辑器类型", GoType:"string", GoName:"Etype"}, "display":&factory.FieldInfo{Name:"display", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"ALL", "SELF", "FRIEND", "PWD"}, DefaultValue:"ALL", Comment:"显示", GoType:"string", GoName:"Display"}, "likes":&factory.FieldInfo{Name:"likes", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"被喜欢次数", GoType:"uint", GoName:"Likes"}, "allow_comment":&factory.FieldInfo{Name:"allow_comment", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"Y", "N"}, DefaultValue:"Y", Comment:"是否允许评论", GoType:"string", GoName:"AllowComment"}, "catid":&factory.FieldInfo{Name:"catid", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"分类ID", GoType:"uint", GoName:"Catid"}, "id":&factory.FieldInfo{Name:"id", DataType:"int", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"ID", GoType:"uint", GoName:"Id"}, "created":&factory.FieldInfo{Name:"created", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"创建时间", GoType:"uint", GoName:"Created"}}, "tag":map[string]*factory.FieldInfo{"rc_type":&factory.FieldInfo{Name:"rc_type", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:30, Options:[]string{}, DefaultValue:"post", Comment:"关联类型", GoType:"string", GoName:"RcType"}, "id":&factory.FieldInfo{Name:"id", DataType:"int", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"ID", GoType:"uint", GoName:"Id"}, "name":&factory.FieldInfo{Name:"name", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:30, Options:[]string{}, DefaultValue:"", Comment:"标签名", GoType:"string", GoName:"Name"}, "uid":&factory.FieldInfo{Name:"uid", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"创建者", GoType:"uint", GoName:"Uid"}, "created":&factory.FieldInfo{Name:"created", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"创建时间", GoType:"uint", GoName:"Created"}, "times":&factory.FieldInfo{Name:"times", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"使用次数", GoType:"uint", GoName:"Times"}}}

}

