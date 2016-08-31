# factory
为了方便使用，将原包中的常用功能进行再次封装。支持的目标功能有：

1. 支持主从式读写分离，支持多主多从
2. 支持同时使用多种数据库且每一种数据库都支持读写分离
3. 支持缓存数据结果
4. 便捷的分页查询功能
5. 更加的易于使用

以上。

# 用法

```
package main
import (
	"fmt"
	"log"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/db/mysql"
)

var settings = mysql.ConnectionURL{
	Host:     "localhost",
	Database: "blog",
	User:     "root",
	Password: "root",
}

type Post struct {
    Id      int     `db:"id,omitempty"`
    Title   string  `db:"title"`
    Group   string  `db:"group"`
    Views   int     `db:"views"`
}

type PostCollection struct {
	Post *Post `db:",inline"`
	User *User `db:",inline"`
}

type User struct {
    Id      int     `db:"id,omitempty"`
    Name   string   `db:"name"`
}

func main() {
	database, err := mysql.Open(settings)
	if err != nil {
		log.Fatal(err)
	}
	factory.SetDebug(true)
	factory.AddDB(database).Cluster(0).SetPrefix(`webx_`)
	defer factory.Default().CloseAll()

	var posts []*Post

    err = factory.All(factory.NewParam().SetCollection(`post`).SetPage(1).SetSize(10).SetRecv(&posts))
	// 生成SQL：SELECT * FROM `webx_post` LIMIT 10

	if err != nil {
		log.Fatal(err)
	}

	for _, post := range posts {
		log.Printf("%q (ID: %d)\n", post.Title, post.Id)
	}
}
```

## 查询多行数据 (使用All方法)

### 方法 1.
```
err = factory.All(factory.NewParam().SetCollection(`post`).SetRecv(&posts))
```

也可以附加更多条件（后面介绍的所有方法均支持这种方式）：
```
    err = factory.NewParam().SetCollection(`post`).SetRecv(&posts).SetArgs(db.Cond{`title LIKE`:`%test%`}).SetMiddleware(func(r db.Result)db.Result{
        return r.OrderBy(`-id`).Group(`group`)
    }).All()
    // 生成SQL：SELECT * FROM `webx_post` WHERE (`title` LIKE "%test%") GROUP BY `group` ORDER BY `id` DESC
```

### 方法 2.
```
err = factory.NewParam().SetCollection(`post`).SetRecv(&posts).All()

```

### 关联查询
```
m := []*PostCollection{}
err = factory.NewParam().SetCollection(`post AS a`).SetCols(db.Raw(`a.*`)).AddJoin(`LEFT`, `user`, `b`, `b.id=a.id`).Select().All(&m)
```

## 查询分页数据 (使用List方法)

### 方法 1.
```
var countFn func()int64
countFn, err = factory.List(factory.NewParam().SetCollection(`post`).SetRecv(&posts).SetPage(1).SetSize(10))
```

### 方法 2.
```
countFn, err = factory.NewParam().SetCollection(`post`).SetRecv(&posts).SetPage(1).SetSize(10).List()
```

## 查询一行数据 (使用One方法)

### 方法 1.
```
var post Post
err = factory.One(factory.NewParam().SetCollection(`post`).SetRecv(&post))
```

### 方法 2.
```
var post Post
err = factory.NewParam().SetCollection(`post`).SetRecv(&post).One()
```

## 插入数据 (使用Insert方法)

### 方法 1.
```
var post Post
post=Post{
    Title:`test title`,
}
err = factory.Insert(factory.NewParam().SetCollection(`post`).SetSend(&post))
```

### 方法 2.
```
var post Post
post=Post{
    Title:`test title`,
}
err = factory.NewParam().SetCollection(`post`).SetSend(&post).Insert()
```

## 更新数据 (使用Update方法)

### 方法 1.
```
var post Post
post=Post{
    Title:`test title`,
}
err = factory.Update(factory.NewParam().SetCollection(`post`).SetSend(&post).SetArgs("id",1))
```

### 方法 2.
```
var post Post
post=Post{
    Title:`test title`,
}
err = factory.NewParam().SetCollection(`post`).SetSend(&post).SetArgs("id",1).Update()
```

## 删除数据 (使用Delete方法)

### 方法 1.
```
err = factory.Delete(factory.NewParam().SetCollection(`post`).SetArgs("id",1))
```

### 方法 2.
```
err = factory.NewParam().SetCollection(`post`).SetArgs("id",1).Update()
```

## 使用事务

### 方法 1.
```
	param = factory.NewParam().SetCollection(`post`).SetTxMW(func(t *factory.Transaction) (err error) {
		param := factory.NewParam().SetCollection(`post`).SetSend(map[string]int{
			"views": 1,
		}).SetArgs("id", 1)
		err = t.Update(param)
		// err=fmt.Errorf(`failured`)
		// 当返回 nil 时，自动执行Commit，否则自动执行Rollback
		return
	})
	factory.Tx(param)
```

### 方法 2.
```
    param = factory.NewParam().Begin().SetCollection(`post`)
    err:=param.SetSend(map[string]int{"views": 1}).SetArgs("id", 1).Update()
    if err!=nil {
        param.End(err)
        return
    }
    err=factory.NewParam().TransFrom(param).SetCollection(`post`).SetSend(map[string]int{"views": 2}).SetArgs("id", 1).Update()
    if err!=nil {
        param.End(err)
        return
    }
    err=factory.NewParam().TransFrom(param).SetCollection(`post`).SetSend(map[string]int{"views": 3}).SetArgs("id", 1).Update()
    param.End(err)
```

# 自动生成数据表的结构体(struct)
进入目录`github.com/webx-top/db/_tools/generator`执行命令
```
go build -o generator.exe
generator.exe -u <数据库用户名> -p <数据库密码> -p <数据库主机名> -e <数据库类型> -d <数据库名> -o <文件保存目录> -pre <数据表前缀> -pkg <生成的包名>
```

支持的参数：

* -u <数据库用户名> 默认为`root`
* -p <数据库密码> 默认为空
* -p <数据库主机名> 默认为`localhost`
* -e <数据库类型> 默认为`mysql`
* -d <数据库名> 默认为`blog`
* -o <文件保存目录> 默认为`dbschema`
* -pre <数据表前缀> 默认为空
* -pkg <生成的包名> 默认为`dbschema`

本命令会自动生成各个表的结构体和所有表的相关信息(存放于子文件夹info内)

生成的文件范例：[blog结构体文件](https://github.com/webx-top/db/tree/master/_tools/generator/dbschema)

每一个生成的结构体中都自带了以下常用方法便于我们使用：

* 设置事务 `SetTrans(trans *factory.Transaction) *结构体名`
* 参数对象 `Param() *factory.Param` 
* 查询一行 `Get(mw func(db.Result) db.Result) error`
* 分页查询 `List(mw func(db.Result) db.Result, page, size int) ([]*结构体名, func() int64, error)`
* 根据偏移量查询 `ListByOffset(mw func(db.Result) db.Result, offset, size int) ([]*结构体名, func() int64, error)`
* 添加数据 `Add(args ...*结构体名) (interface{}, error)`
* 修改数据 `Edit(mw func(db.Result) db.Result, args ...*结构体名) error`
* 删除数据 `Delete(mw func(db.Result) db.Result) error`

我们还可以根据生成的数据表信息来验证表或字段的类型和合法性，我们可以使用生成的文件夹内的子包`info`中的代码来处理，比如：

* 验证表是否存在 `info.Fields.ValidTable(tableName)`
* 验证表中的字段是否存在 `info.Fields.ValidField(tableName,fieldName)`
* 获取某个表中某个字段的信息  `info.Fields[tableName][fieldName]`