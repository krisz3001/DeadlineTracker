package main

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

func GetDeadlineTypes() ([]DeadlineTypes, error) {
	result := make([]DeadlineTypes, 0)
	rows, err := db.Query("SELECT * FROM `DEADLINETYPES`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var item DeadlineTypes
	for rows.Next() {
		rows.Scan(&item.DeadlineTypeId, &item.DeadlineTypeName)
		result = append(result, item)
	}
	return result, nil
}

func GetDeadlineType(id int) (*DeadlineTypes, error) {
	result := &DeadlineTypes{}

	rows, err := db.Query("SELECT * FROM `DEADLINETYPES` WHERE `DeadlineTypeId`=?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&result.DeadlineTypeId, &result.DeadlineTypeName)
	} else {
		return nil, errors.New("item not found")
	}
	return result, nil
}

func CreateDeadlineType(d NewDeadlineType) error {
	_, err := db.Exec("INSERT INTO `DEADLINETYPES` (`DeadlineTypeName`) VALUES (?)", d.DeadlineTypeName)
	return err
}

func UpdateDeadlineType(d NewDeadlineType, id int) error {
	_, err := db.Exec("UPDATE `DEADLINETYPES` SET `DeadlineTypeName`=? WHERE `DeadlineTypeId`=?", d.DeadlineTypeName, id)
	return err
}

func DeleteDeadlineType(id int) error {
	_, err := db.Exec("DELETE FROM `DEADLINETYPES` WHERE `DeadlineTypeId`=?", id)
	return err
}
