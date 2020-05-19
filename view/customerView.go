package main

import (
	"fmt"

	"github.com/dengwenjun1986/customerManage/mode"
	"github.com/dengwenjun1986/customerManage/ser"
)

type customerView struct {
	//定义必要的字段
	key  string //接收用户输入
	loop bool   //是否循环显示主菜单
	//增加一个字段customerService
	customerService *ser.CustomerService
}

//显示所有的客户信息
func (this *customerView) list() {
	//首先获取到当前所有的客户信息（切片）
	customers := this.customerService.List()
	fmt.Println("-------客户列表--------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("-------客户列表完成-------")
	fmt.Println()
}

func (this *customerView) add() {
	fmt.Println("-------添加客户-------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电子邮件:")
	email := ""
	fmt.Scanln(&email)
	//构建一个新的Customer实例
	//id号，系统分配
	customer := mode.NewCustomer2(name, gender, age, phone, email)
	if this.customerService.Add(customer) {
		fmt.Println("-------添加完成-------")
	} else {
		fmt.Println("--------添加失败-------")
	}

}

func (this *customerView) delete() {
	fmt.Println("--------删除客户--------")
	fmt.Println("请选择待删除的客户编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	choice := ""
	for {
		fmt.Println("请确认删除(Y/N)：")
		fmt.Scanln(&choice)
		if choice == "Y" || choice == "y" {
			this.customerService.Delete(id)
			fmt.Println("--------删除客户成功-------")
			break
		} else {
			fmt.Println("输入有误，请重新输入.....")
		}
	}
}
func (this *customerView) update() {
	fmt.Println("--------更新客户信息--------")
	fmt.Println("请选择需更新客户的编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电子邮件:")
	email := ""
	fmt.Scanln(&email)
	choice := ""
	for {
		fmt.Println("是否确定更新客户数据:y/n:")
		fmt.Scanln(&choice)
		if choice == "y" || choice == "Y" {
			this.customerService.Update(id, name, gender, age, phone, email)
			break
		} else if choice == "N" || choice == "n" {
			fmt.Println("放弃更新客户信息.....")
			break
		}
		fmt.Println("输入有误，请重新输入....")
	}
}

func (this *customerView) exit() {
	choice := ""
	for {
		fmt.Println("是否确定退出该客户管理系统：Y/y/N/n")
		fmt.Scanln(&choice)
		if choice == "Y" || choice == "y" {
			this.loop = false
			break
		} else if choice == "N" || choice == "n" {
			this.loop = true
			break
		}
		fmt.Println("输入有误，请重新输入Y/y/N/n....")
	}

}

//显示主菜单
func (cv *customerView) mainMenu() {
	for {
		fmt.Println("-------客户信息管理软件-------")
		fmt.Println("1 添加客户")
		fmt.Println("2 修改客户")
		fmt.Println("3 删除客户")
		fmt.Println("4 客户列表")
		fmt.Println("5 退出系统")
		fmt.Println("请选择(1-5): ")
		fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			cv.add()

		case "2":
			cv.update()
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			cv.exit()
		default:
			fmt.Println("输入有误")
		}
		if !cv.loop {
			break
		}
	}
	fmt.Println("你退出了客户管理系统")
}
func main() {
	//在main函数中创建一个customerView，
	customerView := customerView{
		key:  "",
		loop: true,
	}
	//完成对customerView结构体的customerService字段的初始化
	customerView.customerService = ser.NewCustomerService()
	//显示主菜单
	customerView.mainMenu()
}
