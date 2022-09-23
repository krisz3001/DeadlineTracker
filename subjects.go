package main

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

func GetSubjects() ([]Subject, error) {
	result := make([]Subject, 0)
	rows, err := db.Query("SELECT * FROM SUBJECTS")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var item Subject
	for rows.Next() {
		rows.Scan(&item.SubjectKey, &item.SubjectName)
		result = append(result, item)
	}
	return result, nil
}

func GetSubject(id int) (*Subject, error) {
	result := &Subject{}

	rows, err := db.Query("SELECT * FROM SUBJECTS WHERE `SubjectKey`=?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&result.SubjectKey, &result.SubjectName)
	} else {
		return nil, errors.New("item not found")
	}
	return result, nil
}

func DoesSubjectExist(s NewSubject) (bool, error) {
	rows, err := db.Query("SELECT * FROM `SUBJECTS` WHERE `SubjectName`=?", s.SubjectName)
	if err != nil {
		return true, err
	}
	defer rows.Close()
	if rows.Next() {
		return true, nil
	} else {
		return false, nil
	}
}

func CreateSubject(s NewSubject) error {
	_, err := db.Exec("INSERT INTO `SUBJECTS` (`SubjectName`) VALUES (?)", s.SubjectName)
	return err
}

func UpdateSubject(s NewSubject, id int) error {
	_, err := db.Exec("UPDATE SUBJECTS SET `SubjectName`=? WHERE `SubjectKey`=?", s.SubjectName, id)
	return err
}

func DeleteSubject(id int) error {
	_, err := db.Exec("DELETE FROM SUBJECTS WHERE `Id`=?", id)
	return err
}
