学习笔记
## Q&A

Q: 我们在数据库操作的时候，比如 `dao` 层中当遇到一个 `sql.ErrNoRows` 的时候，是否应该 `Wrap` 这个 `error`，抛给上层。为什么？应该怎么做请写出代码

A: 首先要知道上层怎么返回error的，比如直接返回error，那么是可以Wrap这个error的（符合仅handle error一次的要求）。这种不需要打印堆栈信息的错误，定义对应业务的错误码进行处理就可以了。

```go
var (
	ErrUserNotFound = errors.New("user not fuond")
)
func QueryByUserID(id int) (user *model.User, err error) {
	row := db.QueryRow("select id, name, username, password from user where id = ? ", 1)

	var u model.User
	err = row.Scan(&u.Id, &u.Name, &u.Username, &u.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, errors.Wrap(err, "query by user id error.")
	}
	return &u, nil
}
```

