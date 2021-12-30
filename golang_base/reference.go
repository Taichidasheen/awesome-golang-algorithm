package main

import "fmt"

type Foo interface {
	foo()
	setName(string)
}

type MyFoo struct {
	name string
}

func (mf MyFoo) foo() {
	fmt.Println(mf.name)
}

func (mf MyFoo) setName(n string) {
	mf.name = n
}

func main() {
	// map的引用测试
	mp1 := make(map[int]int)
	mp2 := mp1
	mp1[1] = 1
	if n, ok := mp2[1]; ok {
		fmt.Println(n)
	} else {
		fmt.Println("NO")
	}

	// slice的引用测试
	s1 := make([]int, 10)
	s2 := s1
	s1[0] = 10101
	fmt.Println("slice s2", s2[0])

	// slice的引用测试
	numbers:=[]int{0,1,2,3,4,5,6,7,8,9,10}
	printSlice(numbers)
	numbers1 := numbers[1:3]
	numbers2 := numbers[1:]
	printSlice(numbers1)
	printSlice(numbers2)
	numbers[2] = 222//注意：numbers2的操作会影响numbers和numbers1的结果
	numbers1 = append(numbers1, 333)//numbers1的操作会影响numbers和numbers2的结果
	printSlice(numbers)
	printSlice(numbers1)
	printSlice(numbers2)

	numbers2 = append(numbers2, 11)//numbers2的操作不会影响numbers的结果，因为这个时候numbers2重新开辟了一块内存空间
	numbers2[6] = 66//由于numbers2重新开辟了内存空间，所以这里也不会影响numbers的结果
	numbers1[1] = 11//numbers1和numbers还是共享同一份内存，所以会相互影响
	printSlice(numbers)
	printSlice(numbers1)
	printSlice(numbers2)


	// 数组测试
	a1 := [3]int{1,2,3}
	a2 := a1
	a1[0] = 0
	fmt.Println("array a1:", a1)
	fmt.Println("array a2:", a2)

	// channel的引用测试
	ch1 := make(chan int, 10)
	ch2 := ch1
	ch1 <- 1
	ch2 <- 2  // 这里仅仅是为了防止编译器检测死锁
	fmt.Println(<-ch2)

	// interface的引用测试
	var foo1 Foo
	var foo2 Foo
	mf := MyFoo{name:"foo"}
	foo1 = mf
	foo2 = foo1
	foo1.setName("foo1")
	foo1.foo()//注意：输出结果为foo，而不是foo1，原因是方法接收者不是指针类型
	foo2.foo()//注意：输出结果为foo，而不是foo1，原因是方法接收者不是指针类型
}

func printSlice(x []int) {
	fmt.Printf("len=%d  cap=%d   slice=%v\n",len(x),cap(x),x)
}