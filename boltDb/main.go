package main

import (
	"execlt1/boltDb/operator"
	"execlt1/boltDb/student"
	"fmt"
	"time"
)

func main() {
	boltDB, err := operator.New("howlink", "test111")
	if err != nil {
		panic(err)
	}
	now := time.Now().Unix()
	after := time.Now().Add(3 * time.Second).Unix()
	err = boltDB.Put(fmt.Sprintf("student%d", now), student.Student{
		ID:        1,
		Name:      "mcc",
		Timestamp: now,
		Height:    153,
	}, string(operator.StudentHeight))
	if err != nil {
		panic(err)
	}
	err = boltDB.Put(fmt.Sprintf("student%d", now), student.Student{
		ID:        2,
		Name:      "1111",
		Timestamp: after,
		Height:    174,
	}, string(operator.StudentHeight))
	if err != nil {
		panic(err)
	}
	list, err := boltDB.ConditionalList("student", student.Student{}, func(obj interface{}) bool {
		if obj != nil {
			stu := obj.(student.Student)
			return stu.Height > 153
		}
		return false
	}, string(operator.StudentHeight))
	if err != nil {
		panic(err)
	}
	for _, item := range list {
		stu := item.(student.Student)
		fmt.Println(stu.Name)
	}
}
