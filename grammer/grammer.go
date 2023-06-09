package main

import (
	"log"
	"math"
	"os"
	"reflect"
	"sync"
	"time"

	"kmoh.com/pkglib"
)

type context struct{
	st_name string
	st_score int
}

func create() {

	ctx := context { st_name:"st test", st_score: 20}

	log.Printf("st_name:%s,st_score:%d p:%x\n", ctx.st_name, ctx.st_score, ctx)

	px := new(context)
	px.st_name = "newed";
	px.st_score = 90;

	log.Printf("NEW st_name:%s,st_score:%d p:%x\n", px.st_name, px.st_score, px)
}

func variable() {
	var a string = "var a"
	var b int = 20
	var c,d int = 30,40
	e := "var e"
	var f string
	var g int

	log.Printf("varialble type a:%s, b:%d c:%d, d:%d, e:%s f:%s, g:%d", a,b,c,d,e, f,g)

	//const 상수
	const cons_h string = "const h"
	const cons_i int = 200000

	log.Print(cons_h, cons_i);
}

func loop() {
	i := 1
	//while
	for i <= 3 {
		log.Println(i)
		i++
	}

	for j := 7; j<=9 ; j++ {
		log.Println(j)
	}

	for {
		log.Println("whilte loop")
		break
	}
}

func if_parse() {

   //statment , extress 절이 모두 가능하다
	 //go 에는 삼항 조건문이 없다
	 //조건절에는 괄호가 없다

	if num := 9; num < 0 {
		log.Println(num, "is negative")
	} else if num < 10 {
		log.Println(num, "has 1 digit")
	} else {
		log.Println(num, "has multi digit")
	}
}

func switch_parse() {

	// case에서 복수개의 조건을 처리 할 수 있다.
	// default optional

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday :
		log.Println(" saturuday sunday");
	case time.Thursday :
		log.Println(" thursday");
	default:
		log.Println(" other day");
	}

	// switch에 expression을 사용하지 않을 수 있다. if 조건문 처럼 사용이 가능하다.
	t := time.Now()
	switch {
	case t.Hour() > 15:
		log.Println(" over 15")
	default :
		log.Println(" before 15")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			log.Println("I'm a bool")
		case int:
			log.Println("I'm an int")
		default:
			log.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

func array_parse() {

	var a[5]int // 0 값 설정
	log.Println(a);

	b := [5]int {1,2,3,4,5}
	log.Println(b);

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	log.Println(twoD);
}

func slice_parse() {

	// slice 초기화 1  (배열 차이는 []에 크기 미지정)
	var a []int  //nil splice 초기화, a == nil : true
	log.Println("a == nil", nil == a)
	a = [] int{1,2,3}
	a[1] = 10
	log.Println(a, len(a), cap(a), nil == a)

	// slice 초기화 2 make 로 구성 (default zero value)
	s := make([]string, 5)
	s[0] = "first"
	log.Println(s, len(s), cap(s));
	s = append(s, "last");  //cap5로 지정했지만 append 는 size 뒤에 추가됨 cap 10으로 변경됨
	log.Println(s, len(s), cap(s));

}

func subslice_parse() {
	s := []string{"a","b","c","d"}
	p := s[1:4];
	var k []string = s[3:]

	log.Println(p, len(s), cap(s))
	log.Println(k, len(k), cap(k))

	p[1] = "chg"
	k[0] = "last"
	s[1] = "orichg"

	k = append(k, "final")  //slice가 복제된다.

	k[0] = "first"

	log.Println(s, len(s), cap(s))
	log.Println(p, len(p), cap(p))
	log.Println(k, len(k), cap(k))

	sliceA := [] int{1,2,3}
	sliceB := [] int{4,5,6}

	log.Println(sliceA, len(sliceA), cap(sliceA))

	sliceC := append(sliceA, sliceB...)  //slice가 복제된다.

	sliceA[0] = 100

	log.Println(sliceA, len(sliceA), cap(sliceA))
	log.Println(sliceC, len(sliceC), cap(sliceC))

}

func map_parse() {
	//초기화 1
	var idMap map[int]string
	//초기화 2
	var keyMap map[string]string
	log.Println(idMap, keyMap, "nil == idMap :", nil == idMap )

	keyMap = map[string]string{
		"first" : "a",
		"second" : "b",
		"last" : "z",  //마지막까지 콤마 구성
	}

	keyMap["first"] = "one"
	keyMap["new"] = ""
	delete(keyMap, "last")
	noneElement := keyMap["none"]

	log.Println(keyMap, "noneElement type:", reflect.TypeOf(noneElement), "noneElement empty string:", noneElement == "")

	//key checking
	val, exist := keyMap["second"]
	val2, exist2 := keyMap["oh"]
	val3, exist3 := keyMap["new"]
	log.Println(keyMap, val, exist, val2, exist2, val3, exist3)

	for key, val := range(keyMap) {
		log.Println(key, val)
	}

}

func package_parse() {
	pkglib.PublicTest()
}

type person struct {
	name string
	age int
}

func newPerson() *person {
	p := person{"Chg", 50}
	return &p
}

func struct_parse() {
	//  init 1
	p := person{}
	p.name = "Oh"
	p.age = 10

	log.Println(p)

	// init 2
	var p1 person
	p1 = person{"Kim", 30}
	p2 := person{name:"Ou", age: 50}

	p3 := new(person)
	p3.name = "wow"
	p3.age = 20

	p4 := newPerson()
	p4.name = "changed"

	log.Println(p1, p2,p3, p4)
}


type Rect struct {
	width, height int
}

// value receiver
func (r Rect) area() int {
	r.width = 100
	return r.width * r.height
}

//point receiver
func (r* Rect) square() int {
	r.width = 300;
	return r.width * r.height * 2
}

func method_parse() {
	rect := Rect{10,20}
	area := rect.area()
	log.Println(area, rect)
	//pointer receiver는 rect 를 변경한다.
	square := rect.square()
	log.Println(square, rect)

}


type Shape interface {
	area() float64
	square() float64
}

type Circle struct {
	radius float64
}

func (r Circle) area() float64 {
	r.radius = 20
	return math.Pi *r.radius
}

func (r Circle) square() float64 {
	r.radius = 20;
	return math.Pi * r.radius
}

type CircleDual struct {
	radius float64
}

func (r CircleDual) area() float64 {
	r.radius = 20
	return math.Pi *r.radius * 2
}

func (r CircleDual) square() float64 {
	return math.Pi * r.radius * 2
}


func interface_parse() {
	// TOOD : interface 구현시에는 value receiver만 적용된다. pointer receiver를 사용하는 법을 연구하자.
	// r := Rect{10,20}
	c := Circle{10}
	cd := CircleDual{50}
	showArea(c,cd)

	//empty interface
	var ei interface{}
	ei = c

	showEmptyInterface(ei)

	//empty interface를 이용한 type assertion
	var ta interface{} = "type assertion"
	i := ta
	j := ta.(string)
	//j := ta.(int) // ~~~ runtime exception

	log.Println(i, j, "ta type : ", reflect.TypeOf(ta), "j type : ", reflect.TypeOf(j))

}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area()
		log.Println(a)
	}
}

func showEmptyInterface(ei interface{}) {
	log.Println(ei)
}

func error_phase() {
	f, err := os.Open("./go.modx")

	if err != nil {
		log.Fatal(err.Error())
	} else {
		log.Println(f.Name());
	}

	defer f.Close()

	log.Println("close");

}

func defer_phase() {

	defer log.Println("call parent defer")

	deferPanicTest("start panic test")

}

func deferPanicTest(fn string) {

	log.Println("end panic")
	defer log.Println("call defer")

	// 모든 defer를 호출한다. caller의 defer도 호출하고 에러 종료 처리
	panic(fn)

	log.Println("after defer")
}

func defer_recover_phase() {

	defer log.Println("call parent defer")

	deferRecoverTest("start recover test")

	log.Println("call parent end")

}

func deferRecoverTest(fn string) {

	log.Println("end panic")
	defer func () {
		log.Println("call defer")
		r := recover();
		log.Println("call recover", r)
	}()

	// 모든 defer를 호출한다. caller의 defer도 호출하고 에러 종료 처리
	panic(fn)

	log.Println("after defer")
}

func say(s string) {
	for i :=0; i < 100000; i++ {
		log.Println(s, "****", i)
	}
}

func routine_phase() {
	//say("Sync")

	go say("ASync1")
	go say("ASync2")
	go say("ASync3")

	time.Sleep(time.Second * 100)
}

func anonymous_func_routin_phase() {

	var wait sync.WaitGroup
	wait.Add(2)  //go rutine 기다는 수

	go func() {
		defer wait.Done()  //이 부분의 Add에 지정된 count임
		log.Println("Hello")
	}()

	go func(msg string) {
		defer wait.Done()
		log.Println(msg)
	}("Hi")

	go func() {
		defer wait.Done()
		log.Println("check 3")
	}()


	wait.Wait()
}


func channel_phase() {
	ch := make(chan int)

	go func() {
		ch <- 123
	}()

	i := <- ch
	log.Println(i)

	done := make(chan bool)
	go func() {
		for i :=0; i < 10; i++ {
			log.Println(i)
		}
		done <- true
		log.Println("go routine done")
	}()

	// 위의 Go루틴이 끝날 때까지 대기
	log.Println("running")
	<-done
	log.Println("main done")
}



func channel_buffer_phase() {
	//unbuffered channel : 수신자가 데이터를 받을때가지 송신자는 데이터를 보내는 채널에 묶여있는다.
	//  상위 channel 은 unbufferd channel

	//buffered channel : 수신자가 데이터를 수신 여부와 상관없이 송신자는 데이터를 보내고 다른 일을 한다.
		//ch := make(chan int)
		//ch <- 1 //수신 받을 go routine이 없어 데드락 발생 fatal error: all goroutines are asleep - deadlock!
		//log.Println(<-ch)

		ch := make(chan string, 3)
		ch <- "start channel"
		log.Println(<-ch)  //수신 받을 곳이 없어도진행됨.

		//channel 을 func parameter로 넘기는 방법
		//send channel : chan<-
		func(ch chan<- string) {
				ch <- "send chennel"
		}(ch)

		ch <- "last channel"

		//receive channel : chan<-
		func(ch <-chan string) {
			log.Println(<-ch, "received")
		}(ch)

		//중간에 채널을 닫는다. ch을 닫으면 송신은 안되도 buffer에 따라 수신은 가능하다.
		close(ch)
		if data, success := <-ch; !success {
			log.Println(success, "NOTING CHENNAL")
		} else {
			log.Println(success, data)
		}

		ch2 := make(chan int, 5)
		ch2 <- 1
		ch2 <- 2
		ch2 <- 3

		close(ch2)

		// channel range 방법 1
		// for {
		// 	// change roop
		// 	if data, success := <-ch2; success {
		// 		log.Println(success, data)
		// 	} else {
		// 		log.Println(success, "NOTING CHENNAL DATA")
		// 		break;
		// 	}
		// }

		// channel range 방법 2
		for i := range ch2 {
			log.Println(i)
		}
}

func run1(done chan bool) {
	time.Sleep(1 * time.Second)
	done <- true
}

func run2(done chan bool) {
	time.Sleep(2 * time.Second)
	done <- true
}

func channel_case_phase() {
	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)


EXIT:
	for {
		//switch에서 channel 지정시 모든 채널이 완료 될 때가지 대기를 탄다.
		select {
		case <-done1:
			log.Println("run1 complete")
			close(done2) //done2를 close해도 수신은 가능하여 대기탐
		case <-done2:
			log.Println("run2 complete")
			break EXIT
		}
	}

	//이 케이스로 go rutine의 순차 종료 처리가 가능해짐.
	log.Println("END")
}







func main() {
	// create()
	//variable()
	//loop()
	//if_parse()
	//switch_parse()
	//array_parse()
	//slice_parse()
	//subslice_parse()
	//map_parse()
	//package_parse()
	//struct_parse()
	//method_parse()
	//interface_parse()
	//error_phase()
	//defer_phase()
	//defer_recover_phase()
	//routine_phase();
	//anonymous_func_routin_phase()
	//channel_phase()
	//channel_buffer_phase()
	channel_case_phase()
}

