package go_learn

import (
    "database/sql"
    "log"
)

/**
  不 wrap 这个 error，并不抛错给 上层（调用者）

  为空 不等于 错误、可能是判断 rows 是否存在的情况
  调用者在调用的时候应该考虑为空的情况

  在处理 结果集 的时候，只判断 err 为 nil 的情况下进行处理
*/

func (this *DB) Query(sqlStr string, vals ...interface{}) (result []map[string]string, err error) {
    var rows *sql.Rows

    rows, err = this.Query(sqlStr, vals...)

    // 处理结果集
    if err == nil {
    }

    this.err = err
    return
}

func (this *DB) errorHandler() {
    if this.err != nil && this.err != sql.ErrNoRows {
        log.Fatal(this.err)
    }
}
