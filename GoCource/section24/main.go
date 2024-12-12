package main

import (
	"fmt"
	"log"
	"net"
)

//############## ------First Lecture --------##############//
// // visible inside the package only
// // type person struct {
// // 	fname string
// // 	lname string
// // }

// // type secretAgent struct {
// // 	person
// // 	kill bool
// // }

// // for visible outside the package take first letter capital of variables.
// //
// //	type person struct {
// //		Fname string
// //		Lname string
// //	}
// // type human interface {
// // 	write()
// // 	// speak()

// // }

// type calculateArea interface {
// 	calculate()
// }

// type circle struct {
// 	r float32
// 	b int
// }

// type square struct {
// 	side1 float32
// }

// func main() {
// 	// xi := []int{2, 3, 4, 5, 6}
// 	// fmt.Println(xi)

// 	// m := map[string]int{
// 	// 	"joe": 41,
// 	// 	"toe": 51,
// 	// }
// 	// fmt.Println(m)

// 	// p1 := person{
// 	// 	"mayur",
// 	// 	"jadhav",
// 	// }

// 	// sa1 := secretAgent{
// 	// 	person{
// 	// 		"James",
// 	// 		"bond",
// 	// 	},
// 	// 	true,
// 	// }
// 	radius := circle{
// 		10,
// 		12,
// 	}
// 	square := square{
// 		3.4,
// 	}

// 	// fmt.Println(p1)
// 	// p1.write()
// 	// sa1.speck()
// 	// saySomething(sa1)
// 	// saySomething(p1)
// 	calculateBothArea(radius)
// 	calculateBothArea(square)
// 	// radius.calculate()
// 	// square.calculate()

// }

// // functions in go
// // func (p person) write() {
// // 	fmt.Println(p.fname, `Hii how are you`)
// // }

// // func (sa secretAgent) write() {
// // 	fmt.Println(sa.fname, sa.lname, "will kill you", sa.kill)
// // }
// // func saySomething(h human) {
// // 	// h.speak()
// // 	h.write()
// // }

// // Method to calculate circumference
// func (r circle) calculate() {
// 	const pie = 3.14
// 	circumference := 2 * pie * r.r
// 	fmt.Printf("Circumference of the circle: %.2f\n", circumference)
// }

// func (s1 square) calculate() {
// 	value := s1.side1 * s1.side1
// 	fmt.Printf("Circumference of the circle: %.2f\n", value)
// }

// func calculateBothArea(r calculateArea) {
// 	r.calculate()
// }

//############## ------Third Lecture Server Creation  --------##############//

func main() {
	// li, err := net.Listen("tcp", ":8080")
	// if err != nil {
	// 	log.Panic(err)
	// }
	// defer li.Close()
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()
	fmt.Fprintln(conn, "I dialed you")
	// for {
	// 	conn, err := li.Accept()
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	// go handle(conn)
	// 	// io.WriteString(conn, "\nHello TCP Server Mayur is here ky mhanato bhava")
	// 	// fmt.Fprintln(conn, "How are you connection")
	// 	// fmt.Fprintf(conn, "%v", "Well Done")
	// 	// conn.Close()
	// }
}

// func handle(conn net.Conn) {
// 	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
// 	if err != nil {
// 		log.Println("CONN TIMEOUT")
// 	}
// 	scanner := bufio.NewScanner(conn)
// 	for scanner.Scan() {
// 		ln := scanner.Text()
// 		fmt.Println(ln)
// 		fmt.Fprintf(conn, "I heard you say : %s\n", ln)
// 	}
// 	defer conn.Close()
// 	fmt.Println("Code is here ")
// }
