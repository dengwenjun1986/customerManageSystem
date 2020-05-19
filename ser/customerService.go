package ser

import (
	"fmt"

	"github.com/dengwenjun1986/customerManage/mode"
)

//该结构体完成对Customer的操作，包括增删改查
type CustomerService struct {
	customers []mode.Customer
	//声明一个字段，表示当前切片含有多少个客户
	customerNum int
}

//编写一个方法，可以返回*CustomerService
func NewCustomerService() *CustomerService {
	//为了能够看到有客户在切片中，初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := mode.NewCustomer(1, "张三", "男", 20, "112", "dwj@sohu.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

//返回客户切片
func (this *CustomerService) List() []mode.Customer {
	return this.customers
}

//添加客户到customers切片
func (this *CustomerService) Add(customer mode.Customer) bool {
	//确定一个分配Id的规则，
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

//根据id查找客户在切片中对应下标，如果没有该客户，返回-1
func (this *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(this.customers); i++ {
		if id == this.customers[i].Id {
			index = i
		}
	}
	return index
}

//根据FindById函数查找到的客户id,删除客户函数
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		fmt.Println("输入的id不存在，请重新输入.....")
		return false
	}
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	fmt.Println("删除客户成功.......")
	return true

}

//根据FindById函数查找到的客户id,更新客户数据函数
func (this *CustomerService) Update(id int, name string, gender string, age int, phone string,
	email string) bool {
	index := this.FindById(id)
	if index == -1 {
		fmt.Println("输入的id不存在，请重新输入.....")
		return false
	}
	this.customers[index].Name = name
	this.customers[index].Gender = gender
	this.customers[index].Age = age
	this.customers[index].Phone = phone
	this.customers[index].Email = email
	fmt.Println("更新客户数据完成....")
	return true
}
