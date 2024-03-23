package http

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"time"
)

const MIN = 1
const MAX = 100

func HandleConnection(c net.Conn) {

	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	//fmt.Println("here2")
	for {
		scanner := bufio.NewScanner(c)
		//rd := bufio.NewReader(c)

		var data []string = make([]string, 0)
		// eof := false
		// for !eof {
		// 	line, err := rd.ReadString('\n')
		// 	if err == io.EOF {
		// 		err = nil
		// 		eof = true
		// 	} else if err != nil {
		// 		break
		// 	}
		// 	data = append(data, line)
		// 	//fmt.Println("data:::")
		// }

		for scanner.Scan() {
			//fmt.Println("data:::")
			tmp := scanner.Text()
			if tmp != "" {

				data = append(data, tmp)

			} else {
				break
			}

		}
		for _, d := range data {
			fmt.Println("data:::" + d)

		}

		//result := strconv.Itoa(random()) + "\n"
		res := "HTTP/1.1 200 OK\r\n" +
			"Content-Type: text/html;charset=UTF-8\r\n" +
			// "Content-Length: 27\r\n" +
			"Date: Sat, 25 Mar 2023 16:20:33 GMT\r\n" +
			"Keep-Alive: timeout=60\r\n" +
			"Connection: keep-alive\r\n\r\n" +
			"<!DOCTYPE html>\n" +
			"<html lang='en'>\n" +
			"<body>\n" +
			"<h1>test</h1>\n" +
			"</body>\n" +
			"</html>\n"

		//res2 := "HTTP/1.1 200 OK\r\n"
		lt, err := c.Write([]byte(res))
		//c.Write([]byte(string("Greetings from Spring Boot!")))
		if err == nil {
			fmt.Println("reading standard input:", lt)
			break
		}

	}
	fmt.Println("closing conn")
	c.Close()

}

func Start() {
	lst, err := net.Listen("tcp4", ":1443")
	if err != nil {
		fmt.Println("closing ....")
		return
	}
	defer lst.Close()

	rand.Seed(time.Now().Unix())

	for {
		//fmt.Println("calling....")
		c, err := lst.Accept()
		//fmt.Println("calling.d...")
		if err != nil {
			fmt.Println(err, "main loop")
			return
		}
		fmt.Println("calling handler....")
		go HandleConnection(c)
		//fmt.Println("done")
	}
}
