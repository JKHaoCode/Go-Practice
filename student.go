package main

import (
	"fmt"
	"os"
)

type student struct {
	id   uint64
	name string
}

var (
	allStudent map[uint64]*student
)

func newStudent(id uint64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func showAllStudent() {
	for k, v := range allStudent {
		fmt.Printf("学号：%d 姓名：%s\n", k, v.name)
	}
}

func addStudent() {
	var (
		id   uint64
		name string
	)
	fmt.Println("请输入学生学号：")
	fmt.Scanln(&id)
	fmt.Println("请输入学生姓名：")
	fmt.Scanln(&name)

	newStu := newStudent(id, name)
	allStudent[id] = newStu
}

func deleteStudent() {
	var (
		deleteId uint64
	)
	fmt.Println("请输入学生学号：")
	fmt.Scanln(&deleteId)

	delete(allStudent, deleteId)
}

func main() {
	allStudent = make(map[uint64]*student, 50)
	for {
		fmt.Println("欢迎学生管理系统")
		fmt.Println(`
        1.查询所有学生
        2.新增学生
        3.删除学生
        4.退出
    `)
		fmt.Print("输入要做啥：")

		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d这个选项\n", choice)

		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1) //退出
		default:
			fmt.Println("退出")
		}
	}
}
