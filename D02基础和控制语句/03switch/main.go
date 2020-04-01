package main

import "fmt"

//switch
//简化大量判断(一个变量和具体的值作比较)

func main() {
	//var n int = 5
	// switch n := 6; n {
	// case 1:
	// 	fmt.Println("大")
	// case 2:
	// 	fmt.Println("食")
	// case 3:
	// 	fmt.Println("中")
	// case 4:
	// 	fmt.Println("无")
	// case 5:
	// 	fmt.Println("小")
	// default:
	// 	fmt.Println("无效值")

	switch n := 5; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	//	fallthrough  //执行满足条件的case的下个case,为兼容C语言的case而设计
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")
	default:
		fmt.Println("无效值")
	}

	n := 6
	switch {
	case n > 6:
		fmt.Println("true")
	case n < 6:
		fmt.Println("false")
	default:
		fmt.Println("=6")
	}

	//goto 无条件跳转至标志位

	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto XX
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
XX:
	fmt.Println("over")
}
/* 
算数运算符
i ++   //单独的语句
i --   //单独的语句

关系运算符
==   //GO是强类型语言，相同类型的变量才能比较
!=
>=
<=
>
<

逻辑运算符
&& //并且
|| //或者
not //取反

位运算符 //位指的是二进制位
& //按位与    5的二进制(101) & 2的二进制(10)  =  101 & 010 = 000
| //按位或    101 | 010 = 111 = 7
^ //按位异或   两位不一样则为1   5 ^ 2 = 111
<< //将二进制数向左移动指定位数 5 << 1  = 101 << 1 =1010 =10
>> //将二进制数向右移动指定位数 5 >> 2  = 101 >> 2 =1 =1

赋值运算 //给变量赋值

+= -= *= /= %=  <<=  >>=  &= |= ^=
*/