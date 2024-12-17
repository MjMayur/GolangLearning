package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
		// io.WriteString(conn, "\nHello TCP Server Mayur is here ky mhanato bhava")
		// fmt.Fprintln(conn, "How are you connection")
		// fmt.Fprintf(conn, "%v", "Well Done")
		// conn.Close()
	}
}

func handle(conn net.Conn) {
	// err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	// if err != nil {
	// 	log.Println("CONN TIMEOUT")
	// }
	// scanner := bufio.NewScanner(conn)
	// for scanner.Scan() {
	// 	ln := scanner.Text()
	// 	fmt.Println(ln)
	// 	fmt.Fprintf(conn, "I heard you say : %s\n", ln)
	// }
	defer conn.Close()
	fmt.Println("Code is here ")
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			break
		}
		i++
	}
}

// #############-------TCP server Mux server
func mux(conn net.Conn, ln string) {
	m := strings.Fields(ln)[0]
	u := strings.Fields(ln)[1]
	fmt.Println("****Method", m)
	fmt.Println("****URI", u)

	//multiplexer
	if m == "GET" && u == "/" {
		index(conn)
	}
	if m == "GET" && u == "/about" {
		about(conn)
	}
	if m == "GET" && u == "/contact" {
		contact(conn)
	}
	if m == "GET" && u == "/apply" {
		apply(conn)
	}
	if m == "POST" && u == "/apply" {
		applyProcess(conn)
	}

}

func index(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta
charet="UTF-8"><title></title></head><body>
<strong>INDEX</strong><br>
<a href="/">index</a><br>
<a href="/about">about</a><br>
<a href="/contact">contact</a><br>
<a href="/apply">apply</a><br>
</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func about(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta
charet="UTF-8"><title></title></head><body>
<strong>About</strong><br>
<a href="/">index</a><br>
<a href="/about">about</a><br>
<a href="/contact">contact</a><br>
<a href="/apply">apply</a><br>
</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func contact(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta
charet="UTF-8"><title></title></head><body>
<strong>Contact</strong><br>
<a href="/">index</a><br>
<a href="/about">about</a><br>
<a href="/contact">contact</a><br>
<a href="/apply">apply</a><br>
</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func apply(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta
charet="UTF-8"><title></title></head><body>
<strong>APPLY</strong><br>
<a href="/">index</a><br>
<a href="/about">about</a><br>
<a href="/contact">contact</a><br>
<a href="/apply">apply</a><br>
<form method="post" action="/apply">
<input type="submit" value="apply">
</form>
</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func applyProcess(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta
charet="UTF-8"><title></title></head><body>
<strong>APPLY PROCESS</strong><br>
<a href="/">index</a><br>
<a href="/about">about</a><br>
<a href="/contact">contact</a><br>
<a href="/apply">apply</a><br>
</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
