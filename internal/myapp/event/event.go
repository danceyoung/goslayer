package event

import (
	"errors"
)

//Events implement biz logic and wrap response data
func Events() ([]map[string]interface{}, error) {
	return events()
}

//query events from db,eg:mysql
func events() ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	// wrapping data by using mysql
	// rows, err := db.NewMysql().Query("sql statement")
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	var temp string
	// 	err = rows.Scan(&temp)
	// }
	result = append(result, map[string]interface{}{"id": 1, "event_name": "dancing competition"}, map[string]interface{}{"id": 1, "event_name": "singing competition"})
	return result, nil

}

type Member struct {
	Name  string
	Email string
}

func JoinAEvent(eventid string, member Member) error {
	if len(eventid) == 0 || len(member.Name) == 0 || len(member.Email) == 0 {
		return errors.New("parmas are not enough")
	}
	if err := joinAEvent(eventid, member); err != nil {
		return errors.New("join a event occurring a error: " + err.Error())
	}
	return nil
}

//insert a record into db
func joinAEvent(eventid string, member Member) error {
	return nil
}
