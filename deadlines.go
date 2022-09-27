package main

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

func GetDeadlines() ([]Deadline, error) {
	result := make([]Deadline, 0)
	rows, err := db.Query("SELECT `Id`, `SubjectName`, `Deadline`, `DeadlineTypeName`, `Topic`, `Comments` FROM DEADLINES LEFT JOIN SUBJECTS ON DEADLINES.SubjectId = SUBJECTS.SubjectKey LEFT JOIN DEADLINETYPES ON DEADLINES.TypeId = DEADLINETYPES.DeadlineTypeId WHERE `Deadline` > NOW() ORDER BY `Deadline`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var item Deadline
	for rows.Next() {
		rows.Scan(&item.Id, &item.Subject, &item.Deadline, &item.Type, &item.Topic, &item.Comments)
		result = append(result, item)
	}
	return result, nil
}

func GetDeadline(id int) (*Deadline, error) {
	result := &Deadline{}
	//todo: sql ha valaha használnám ezt
	rows, err := db.Query("SELECT * FROM DEADLINES WHERE `Id`=?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&result.Id, &result.Subject, &result.Deadline, &result.Type, &result.Topic, &result.Comments)
	} else {
		return nil, errors.New("item not found")
	}
	return result, nil
}

func CreateDeadline(d NewDeadline) error {
	_, err := db.Exec("INSERT INTO `DEADLINES` (`SubjectId`, `Deadline`, `TypeId`, `Topic`, `Comments`) VALUES (?,?,?,?,?)", d.SubjectId, d.Deadline, d.TypeId, d.Topic, d.Comments)
	return err
}

func UpdateDeadline(d NewDeadline, id int) error {
	_, err := db.Exec("UPDATE DEADLINES SET `SubjectId`=?, `Deadline`=?, `TypeId`=?, `Topic`=?, `Comments`=? WHERE `Id`=?", d.SubjectId, d.Deadline, d.TypeId, d.Topic, d.Comments, id)
	return err
}

func DeleteDeadline(id int) error {
	_, err := db.Exec("DELETE FROM DEADLINES WHERE `Id`=?", id)
	return err
}
