# バルクインサート

## sqlx

```go
q := "INSERT INTO role(type,value) VALUES (:type, :value)"

pp := []map[string]interface{}{
	{"type": "type1", "value": "val1"},
	{"type": "type2", "value": "val2"},
	{"type": "type3", "value": "val3"},
}

res, err := db.NamedExec(q, pp)
```

`:type`, `:value`が大事

## GORM

スライス入れるだけでバルクインサートになるらしい

```go
var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
db.Create(&users)
```

## SQL Boiler

https://qiita.com/touyu/items/4b25fbf12804f12778b7

## その他

バルクインサートがサポートされてない場合はどうしようもないのでSprintfで頑張る

### バルクインサートがサポートされてないやつ
- kyleconroy/sqlc
- go-xorm/xorm
- go-gorp/gorp

※自力で実装する場合バインドできる引数の上限値を気にする

- 例1）PostgreSQL・MySQLは65535らしい
- 例2）SQLite3は外から指定できるっぽい: https://github.com/mattn/go-sqlite3/issues/704


### バルクインサートがサポートされてるやつ
- [beego/beego](https://github.com/beego/beedoc/blob/master/en-US/mvc/model/object.md#:~:text=for%20auto%20fields.-,InsertMulti,-Insert%20multiple%20objects)
- [ent/ent](https://entgo.io/ja/docs/crud/#create-many)
- [go-pg/pg](https://qiita.com/bubu_suke/items/8be0177e4da03cb153a2#bulk%E5%87%A6%E7%90%86)

### 謎なやつ
- xo/xo

---
ソース：https://zenn.dev/ryoneko/articles/4c1267d7d0e0ca （xo/xoまで調査）


## ISUCONのこれまでのORM

- ISUCON9からは予選・本選ともにずっとsqlxが使われている
- ISUCON8以前は直叩きしてるっぽい？