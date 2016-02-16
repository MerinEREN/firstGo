package main

import (
	"fmt"
	"github.com/MerinEREN/firstGo/stringutil"
	//"io/ioutil"
	//"log"
	//"github.com/zenazn/goji"
	//"github.com/zenazn/goji/web"
	"math"
	//"net/http"
	//"os"
	"sort"
	"strconv"
	"sync"
	"time"
)

var dummySlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
var dummyArray = [5]float64{1.1, 2.2, 3.3, 4, 5}

var pizzaNumber = 0
var pizzaName = ""

/*var pizzaNumber int = 0
var pizzaName string = ""*/

func main() {
	sOne, iOne := "merin", 5
	fmt.Println(sOne, iOne)

	dummySlice2 := make([]int, 6)
	copy(dummySlice2, dummySlice)
	dummySlice2 = append(dummySlice2, -9, 13, 8, 9, 10)

	for i, value := range dummySlice {
		fmt.Println(i, value)
	}

	defer func() {
		for i, value := range dummySlice2 {
			fmt.Println(i, value)
		}
		fmt.Println(len(dummySlice2))
	}()

	for _, value := range dummyArray {
		fmt.Printf("%0.7f\n", value)
	}

	presAge := make(map[string]int)
	presAge["mrn"] = 31
	presAge["mtn"] = 56
	presAge["hry"] = 77
	delete(presAge, "mrn")
	for i, value := range presAge {
		fmt.Println(i, value)
	}
	fmt.Println(len(presAge))

	num1, num2 := addThemUp(5)
	fmt.Println(num1, num2)

	fmt.Println(substractAndAddThem(dummySlice2))

	closureResultString := func() int {
		closureHolderInt := 0
		for _, value := range presAge {
			closureHolderInt += value
		}
		return closureHolderInt
	}
	fmt.Println(closureResultString)

	facValue := 5
	fmt.Println(Factorial(facValue))

	fmt.Println(safeDiv(3, 0))

	demPanic()

	valX := 0
	fmt.Println(changeValX(&valX))
	fmt.Println(valX)

	yPtr := new(int)
	fmt.Println(changeValY(yPtr))

	rect1 := Rectangle{0, 50, 10, 10}
	circle1 := Circle{10}
	fmt.Println(getArea(rect1))
	fmt.Println(getArea(circle1))
	fmt.Println(getPerimeter(rect1))
	fmt.Println(getPerimeter(circle1))

	fmt.Printf(stringutil.Reverse("!oG ,olleH"))

	//GuessingGame()

	/*file, err := os.Create("sample.go")
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("Sample text file createt.\nHell Yeah !!!!!!!!!!!\n")
	file.Close()
	//os.Remove("sample.go")
	os.Rename("sample.go", "dummy.txt")

	stream, err := ioutil.ReadFile("dummy.txt")
	if err != nil {
		fmt.Println("\nHeHeHeHeHe")
		log.Fatal(err)
	}

	fmt.Println(string(stream))*/

	/*http.HandleFunc("/", rootHandler)
	http.HandleFunc("/England", englandHandler)
	http.ListenAndServe(":8080", nil)*/
	//goji.Get("/hello/:name", hello)
	//goji.Serve()

	stringChan := make(chan string)
	for i := 0; i < 5; i++ {
		go makeDough(stringChan)
		go addSouce(stringChan)
		go addToppings(stringChan)

		time.Sleep(time.Millisecond * 1000)
	}

	//FUNCTIONS !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	// FUNCTIONS ARE VALUES TOO. THEY CAN BE PASSED AROUND JUST LIKE OTHER
	// VALUES.
	// FUNCTION VALUES MAY BE USED AS FUNCTION ARGUMANTS AND RETURN VALUES.

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
	// FONCTION CLOSURES
	// Go functions may be closures. A closure is a function value that
	// references variables from outside its body. The function may access
	// and assign to the referenced variables; in this sense the function
	// is "bound" to the variables.
	// For example, the adder function returns a closure. Each closure is
	// bound to its own sum variable.
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i))
		fmt.Println(neg(-2 * i))
	}

	if err := run(); err != nil {
		fmt.Println(err)
	}

	// IMPORTANT EXAMPLE OF CREATING CUSTOM ERROR VIA error INTERFACE !!!!!
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1) // INCREMENT OUR WAIT GROUP
		x := i
		go func() { // AFTER A LITTLE NAP, OUR GOROUTINES WILL
			// DECREMENT THE WAIT GROUP
			time.Sleep(time.Second * 2)
			fmt.Println(x)
			wg.Done()
		}()
	}
	wg.Wait() // WAIT FOR OUR WAIT GROUP'S INTERNAL COUNTER TO HIT ZERO
}

func addThemUp(val int) (int, int) {
	return val + 1, val + 2
}

/*func addThemUp(val int) (x int, y int) {
	x = val + 1
	y = val + 2

	return
}*/

func substractAndAddThem(array []int) (int, int) {
	sub := 0
	add := 0
	for _, value := range array {
		sub -= value
		add += value
	}
	return sub, add
}

func Factorial(v int) int {
	if v == 0 {
		return 1
	}
	return v * Factorial(v-1)
}

func safeDiv(v1, v2 int) int {
	defer func() {
		fmt.Println(recover())
	}()
	return v1 / v2
}

func demPanic() {
	defer func() {
		fmt.Println(recover())
	}()
	panic("PANIC!!!!!")
}

func changeValX(valX *int) (*int, int) {
	*valX = 17
	return valX, *valX
}

func changeValY(p *int) (*int, int) {
	*p = 100
	return p, *p
}

type Rectangle struct {
	leftX  float64
	topY   float64
	height float64
	width  float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Shape interface {
	area() float64
	perimeter() float64
}

func getArea(s Shape) float64 {
	return s.area()
}

func getPerimeter(s Shape) float64 {
	return s.perimeter()
}

func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}

/*func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Earth !!!!!!!\n")
}

func englandHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, England !!!!!!!\n")
}*/

/*func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}*/

func makeDough(c chan string) {
	pizzaNumber++
	pizzaName = "This is Pizza #" + strconv.Itoa(pizzaNumber)
	fmt.Println(pizzaName + " send to souce adding\n")
	c <- pizzaName
	//time.Sleep(time.Millisecond * 10)
}

func addSouce(c chan string) {
	pizza := <-c
	fmt.Println(pizza + " send to topping adding\n")
	c <- pizzaName
	//time.Sleep(time.Millisecond * 10)
}

func addToppings(c chan string) {
	pizza := <-c
	fmt.Println(pizza + " send to the customer\n")
	//time.Sleep(time.Millisecond * 10)
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

type MyError struct {
	When time.Time
	What string
}

func (me MyError) Error() string {
	return fmt.Sprintf("When: %v\nWhat: %s", me.When, me.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"It didn't work!!!",
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("can not Sqrt negatine number: %f", float64(e))
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return f, ErrNegativeSqrt(f)
	} else {
		return f, nil
	}
}
