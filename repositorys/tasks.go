package repositorys

import (
	"database/sql"
	"taskmanager/models"
)

func SelectTask(db *sql.DB, accountID int) (resTasks []models.Task, err error) {
	const sqlStr = ("SELECT task_id, title, due_date, progress FROM tasks WHERE account_id=$1;")
	rows, err := db.Query(sqlStr, accountID)
	if err != nil {
		return []models.Task{}, err
	}
	var resTask = models.Task{AccountID: accountID}
	for rows.Next() {
		rows.Scan(&resTask.TaskID, &resTask.Title, &resTask.DueDate, &resTask.Progress)
		resTasks = append(resTasks, resTask)
	}
	return resTasks, nil
}

func InsertTask(db *sql.DB, reqTask models.Task) (resTask models.Task, err error) {
	const sqlStr = ("INSERT INTO tasks (account_id, title, due_date, progress) VALUES ($1, $2, $3, $4) RETURNING task_id;")
	var taskID int
	err = db.QueryRow(sqlStr, reqTask.AccountID, reqTask.Title, reqTask.DueDate, reqTask.Progress).Scan(&taskID)
	if err != nil {
		return models.Task{}, err
	}
	resTask = reqTask
	resTask.TaskID = taskID
	return resTask, nil
}

func UpdateTask(db *sql.DB, reqTask models.Task) (resTask models.Task, err error) {
	const sqlStr = ("UPDATE tasks SET title = $1, due_date = $2, progress = $3 WHERE task_id = $4;")
	_, err = db.Exec(sqlStr, reqTask.Title, reqTask.DueDate, reqTask.Progress, reqTask.TaskID)
	if err != nil {
		return models.Task{}, err
	}
	resTask = reqTask
	return resTask, nil
}

func DeleteTask(db *sql.DB, taskID int) (err error) {
	const sqlStr = ("DELETE FROM tasks WHERE task_id = $1;")
	_, err = db.Exec(sqlStr, taskID)
	return err
}
