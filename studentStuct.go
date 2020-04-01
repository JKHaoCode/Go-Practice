package main

import (
	"fmt"
)

type student struct {
	id   int64
	name string
}

type studentMgr struct {
	allStudent map[int64]student
}

func (s studentMgr) showStudents() {

}

func (s studentMgr) addStudents() {

}

func (s studentMgr) editStudents() {

}

func (s studentMgr) deleteStudents() {

}

func main() {

}

func showMenu() {
	fmt.Println("welcome sms")
	fmt.Println(`
        1.查询所有学生
        2.新增学生
        3.编辑学生
        3.删除学生
        4.退出
    `)
}
