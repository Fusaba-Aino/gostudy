package modernc_org_sqlite

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

// 链接数据库
func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./test.db")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

// 关闭数据库
func CloseDB(db *sql.DB) {
	if db != nil {
		db.Close()
	}
}

// 创建表
func CreateTable(db *sql.DB) error {
	sql := `CREATE TABLE IF NOT EXISTS user (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		age INTEGER
	);`

	_, err := db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

// 插入数据
func InsertUser(db *sql.DB, name string, age int) error {
	sql := `INSERT INTO user (name, age) VALUES (?, ?);`

	_, err := db.Exec(sql, name, age)
	if err != nil {
		return err
	}

	return nil
}

// 查询数据
func QueryUsers(db *sql.DB) ([]map[string]interface{}, error) {
	sql := `SELECT id, name, age FROM user;`

	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []map[string]interface{}
	for rows.Next() {
		var id int
		var name string
		var age int

		err := rows.Scan(&id, &name, &age)
		if err != nil {
			return nil, err
		}

		user := map[string]interface{}{
			"id":   id,
			"name": name,
			"age":  age,
		}
		users = append(users, user)
	}

	return users, nil
}

// 更新数据
func UpdateUserAge(db *sql.DB, id int, age int) error {
	sql := `UPDATE user SET age = ? WHERE id = ?;`

	_, err := db.Exec(sql, age, id)
	if err != nil {
		return err
	}

	return nil
}

// 删除数据
func DeleteUser(db *sql.DB, id int) error {
	sql := `DELETE FROM user WHERE id = ?;`

	_, err := db.Exec(sql, id)
	if err != nil {
		return err
	}

	return nil
}
