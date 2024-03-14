package sq

import (
	"errors"
)

func InsertPassWord(name string, word string) error {
	var u PassWord
	db.Limit(1).Find(&u)
	if u.ID != 0 {
		return errors.New("账户已经存在")
	}
	db.Last(&u)
	u.ID, u.Name, u.PassWord = u.ID+1, name, word
	db.Create(&u)
	return nil
}

func InsertUser(name string, word string) error {
 
	return nil

}
