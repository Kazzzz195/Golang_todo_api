package repositories

import (
	"database/sql"
	"time"

	"github.com/Kazzzz195/GoProject/models"
	"github.com/google/uuid"
)

func GetTodoByID(db *sql.DB, id int) (models.Todo, error) {
	const sqlStr = `SELECT id, title, body, due_date, complete_at, created_at, update_at FROM todos WHERE id = ?`
	var todo models.Todo
	var completeAt sql.NullTime
	var updateAt sql.NullTime

	err := db.QueryRow(sqlStr, id).Scan(&todo.ID, &todo.Title, &todo.Body, &todo.DueDate, &completeAt, &todo.CreatedAt, &updateAt)
	if err != nil {
		if err == sql.ErrNoRows {
			// レコードが見つからない場合
			return models.Todo{}, nil
		}
		return models.Todo{}, err
	}

	// completeAtが有効かどうかをチェックし、適切なデフォルト値を設定
	if completeAt.Valid {
		todo.CompleteAt = completeAt.Time
	} else {
		todo.CompleteAt = time.Time{} // デフォルト値としてゼロ値を使用
	}

	// updateAtが有効かどうかをチェックし、適切なデフォルト値を設定
	if updateAt.Valid {
		todo.UpdateAt = updateAt.Time
	} else {
		todo.UpdateAt = time.Time{} // デフォルト値としてゼロ値を使用
	}

	return todo, nil
}

func InsertTodo(db *sql.DB, todo models.Todo) (models.Todo, error) {

	id := createUUID()

	const sqlStr = `INSERT INTO todos (id, title, body, due_date, created_at) VALUES (?, ?, ?, ?, ?);`
	result, err := db.Exec(sqlStr, id, todo.Title, todo.Body, todo.DueDate, todo.CreatedAt)
	if err != nil {
		return models.Todo{}, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return models.Todo{}, err
	}

	todo.ID = int(lastInsertId)
	return todo, nil
}

func GetAllTodos(db *sql.DB) ([]models.Todo, error) {
	const sqlStr = `SELECT * FROM todos ORDER BY created_at DESC`
	rows, err := db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		var completeAt sql.NullTime
		var updateAt sql.NullTime

		err := rows.Scan(&todo.ID, &todo.Title, &todo.Body, &todo.DueDate, &completeAt, &todo.CreatedAt, &updateAt)
		if err != nil {
			return nil, err
		}
		//NULL の場合エラー出るので
		// completeAtが有効かどうかをチェックし、適切なデフォルト値を設定
		if completeAt.Valid {
			todo.CompleteAt = completeAt.Time
		} else {
			todo.CompleteAt = time.Time{} // デフォルト値としてゼロ値を使用
		}

		// updateAtが有効かどうかをチェックし、適切なデフォルト値を設定
		if updateAt.Valid {
			todo.UpdateAt = updateAt.Time
		} else {
			todo.UpdateAt = time.Time{} // デフォルト値としてゼロ値を使用
		}

		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

// UpdateTodo 指定されたIDのTodoを更新する関数
func UpdateTodo(db *sql.DB, todo models.Todo) (models.Todo, error) {

	const sqlStr = `
		UPDATE todos
		SET title = ?, body = ?, due_date = ?, complete_at = ?, update_at = ?
		WHERE id = ?;
	`

	_, err := db.Exec(sqlStr, todo.Title, todo.Body, todo.DueDate, todo.CompleteAt, todo.UpdateAt, todo.ID)
	if err != nil {
		return models.Todo{}, err
	}

	// 更新されたTodoを再度取得して返す
	return GetTodoByID(db, todo.ID)
}

func DeleteTodoById(db *sql.DB, id int) error {

	const sqlStr = `DELETE FROM todos WHERE id = ?`

	_, err := db.Exec(sqlStr, id)
	if err != nil {
		return err
	}

	return nil
}

func CompleteTodo(db *sql.DB, todoId int, currentTime time.Time) (models.Todo, error) {

	const sqlStr = `
		UPDATE todos
		SET  complete_at = ?
		WHERE id = ?;
	`

	_, err := db.Exec(sqlStr, currentTime, todoId)
	if err != nil {
		return models.Todo{}, err
	}
	// 更新されたTodoを再度取得して返す
	updatedTodo, err := GetTodoByID(db, todoId)
	if err != nil {
		return models.Todo{}, err
	}
	return updatedTodo, err
}

func SearchTodosByBody(db *sql.DB, body string) ([]models.Todo, error) {
	const sqlStr = `SELECT * FROM todos WHERE body LIKE ? ORDER BY created_at DESC`
	rows, err := db.Query(sqlStr, "%"+body+"%")
	if err != nil {
		return []models.Todo{}, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Body, &todo.DueDate, &todo.CompleteAt, &todo.CreatedAt, &todo.UpdateAt); err != nil {
			return []models.Todo{}, err
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return []models.Todo{}, err
	}

	return todos, nil
}

func SearchTodosByTitle(db *sql.DB, title string) ([]models.Todo, error) {
	const sqlStr = `SELECT * FROM todos WHERE title LIKE ? ORDER BY created_at DESC`
	rows, err := db.Query(sqlStr, "%"+title+"%")
	if err != nil {
		return []models.Todo{}, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Body, &todo.DueDate, &todo.CompleteAt, &todo.CreatedAt, &todo.UpdateAt); err != nil {
			return []models.Todo{}, err
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return []models.Todo{}, err
	}

	return todos, nil
}

func SearchCompletedTodos(db *sql.DB, currentTime time.Time) ([]models.Todo, error) {

	const sqlStr = `SELECT * FROM todos WHERE complete_at IS NOT NULL ORDER BY created_at DESC`
	rows, err := db.Query(sqlStr, currentTime)
	if err != nil {
		return []models.Todo{}, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Body, &todo.DueDate, &todo.CompleteAt, &todo.CreatedAt, &todo.UpdateAt); err != nil {
			return []models.Todo{}, err
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return []models.Todo{}, err
	}

	return todos, nil
}
func SearchOngoingTodos(db *sql.DB, currentTime time.Time) ([]models.Todo, error) {

	const sqlStr = `SELECT * FROM todos WHERE due_date > ? ORDER BY created_at DESC`
	rows, err := db.Query(sqlStr, currentTime)
	if err != nil {
		return []models.Todo{}, err
	}

	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Body, &todo.DueDate, &todo.CompleteAt, &todo.CreatedAt, &todo.UpdateAt); err != nil {
			return []models.Todo{}, err
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return []models.Todo{}, err
	}

	return todos, nil
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}
