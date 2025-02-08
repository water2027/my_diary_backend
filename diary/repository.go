package diary

import (
	"errors"
	"my_diary/database"
)

type diaryHandler interface {
	createDiary() error
	getDiaryById() error
	updateDiary() error
	deleteDiary() error
}

func (d *diaryModel) createDiary() error {
	db := database.GetMysqlDB()
	stmt, err := db.Prepare("INSERT INTO diaries(userId, title, content, isPublic) VALUES(?, ?, ? ,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(d.userId, d.title, d.content, d.isPublic)
	if err != nil {
		return err
	}
	return nil
}

func getPublicDiaryNum() (int, error) {
	db := database.GetMysqlDB()
	count := 0
	err := db.QueryRow("SELECT COUNT(*) FROM diaries WHERE isPublic = true").Scan(&count)
	return count, err
}

func getUserDiaryNum(userId int) (int, error) {
	db := database.GetMysqlDB()
	count := 0
	err := db.QueryRow("SELECT COUNT(*) FROM diaries WHERE userId = ?", userId).Scan(&count)
	return count, err
}

func getPublicDiaryList(offset, limit int) ([]Diary, error) {
	db := database.GetMysqlDB()
	stmt, err := db.Prepare("SELECT diaryId, userId, title, content, isPublic, createdAt, updatedAt FROM diaries WHERE isPublic = true LIMIT ?, ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	diaries := []Diary{}
	for rows.Next() {
		var d Diary
		err := rows.Scan(&d.DiaryId, &d.UserId, &d.Title, &d.Content, &d.IsPublic, &d.CreatedAt, &d.UpdatedAt)
		if err != nil {
			return nil, err
		}
		diaries = append(diaries, d)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return diaries, nil
}

func getUserDiaryList(userId, offset, limit int) ([]Diary, error) {
	db := database.GetMysqlDB()
	stmt, err := db.Prepare("SELECT diaryId, userId, title, content, isPublic, createdAt, updatedAt FROM diaries WHERE userId = ? LIMIT ?, ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	diaries := []Diary{}
	for rows.Next() {
		var d Diary
		err := rows.Scan(&d.DiaryId, &d.UserId, &d.Title, &d.Content, &d.IsPublic, &d.CreatedAt, &d.UpdatedAt)
		if err != nil {
			return nil, err
		}
		diaries = append(diaries, d)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return diaries, nil
}

func (d *diaryModel) getDiaryById() error {
	db := database.GetMysqlDB()
	stmt, err := db.Prepare("SELECT * FROM diaries WHERE diaryId = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(d.diaryId).Scan(&d.diaryId, &d.userId, &d.title, &d.content, &d.createdAt, &d.updatedAt, &d.isPublic)
	if err != nil {
		return errors.New("diary not found")
	}
	return nil
}

func (d *diaryModel) updateDiary() error {
	db := database.GetMysqlDB()
	stmtString := "UPDATE diaries SET "
	args := []interface{}{}
	hasSet := false
	if d.title != "" {
		stmtString += "title = ?, "
		args = append(args, d.title)
		hasSet = true
	}
	if d.content != "" {
		stmtString += "content = ?, "
		args = append(args, d.content)
		hasSet = true
	}
	if d.isPublic {
		stmtString += "isPublic = ?, "
		args = append(args, d.isPublic)
		hasSet = true
	}
	if !hasSet {
		return errors.New("no update field")
	}
	stmtString = stmtString[:len(stmtString)-2]
	stmtString += " WHERE diaryId = ?"
	args = append(args, d.diaryId)
	stmt, err := db.Prepare(stmtString)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(args...)
	if err != nil {
		return err
	}
	return nil
}

func (d *diaryModel) deleteDiary() error {
	db := database.GetMysqlDB()
	stmt, err := db.Prepare("DELETE FROM diaries WHERE diaryId = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(d.diaryId)
	if err != nil {
		return err
	}
	return nil
}


