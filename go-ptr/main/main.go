package main

import (
	"fmt"
	"myself/go-ptr/p"
	"unsafe"
)

func main() {
	// 初始化内存，返回p.V类型的一个指针
	v := new(p.V)

	// v 的地址也是v.i的地址
	fmt.Println(fmt.Sprintf("new ptr:%p", v))
	v.PutIPtr()
	i := (*int)(unsafe.Pointer(v))
	*i = 66

	// uintptr(unsafe.Sizeof(int(2))) 获取int类型的指针长度
	// v 指针地址的基础上加上int类型长度，为v.j的地址
	j := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + uintptr(unsafe.Sizeof(int(2)))))
	*j = "this is string"
	k := (*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(j)) + uintptr(unsafe.Sizeof(string("22")))))
	*k = 99
	v.PutI()
	v.PutJ()
	v.PutK()
}
