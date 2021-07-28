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
	foo1.foo()
	foo2.foo()
}