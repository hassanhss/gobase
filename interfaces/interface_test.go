package interfaces

import "testing"
/*
1,单元测试代码的go文件必须以_test.go结尾，Go语言测试工具只认符合这个规则的文件
2,单元测试的函数名必须以Test开头，是可导出公开的函数。备注：函数名最好是Test+要测试的方法函数名
3,测试函数的签名必须接收一个指向testing.T类型的指针作为参数，并且该测试函数不能返回任何值
 */
func TestInterfaces(t *testing.T) {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}
	mike.SayHi()
	mike.Sing("good day")
	mike.BorrowMoney(1000)

	tom.SayHi()
	tom.Sing("good study")
	tom.SpendSalary(2000)
}
