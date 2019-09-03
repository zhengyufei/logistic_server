package main

//
//import (
//	"fmt"
//	"github.com/globalsign/mgo/bson"
//	"gitlab.xiaoduoai.com/ecrobot/logistic_server/models"
//)
//
//type TestStruct struct {
//	Name string
//	ID   int32
//}
//
//func main() {
//	fmt.Println("start")
//	data, err := bson.Marshal(&models.Logistic2{})//
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("%q", data)
//
//	value := models.Logistic2{}
//	err2 := bson.Unmarshal(data, &value)
//	if err2 != nil {
//		panic(err)
//	}
//	fmt.Println("value:", value)
//
//	mmap := bson.M{}
//	err3 := bson.Unmarshal(data, mmap)
//	if err3 != nil {
//		panic(err)
//	}
//	fmt.Println("mmap:", mmap)
//}
