package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	
	port := os.Getenv("PORT")
	nl, err := net.Listen("tcp", "port")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer nl.Close() //await
	log.Printf("server started on heroku")

	for {

		conn, err := nl.Accept()
		if err != nil {
			fmt.Println(err.Error())
			//continue
		}

		//fmt.Println(conn.RemoteAddr().String())
		bs := make([]byte, 1024)
		n, e := conn.Read(bs)
		if e != nil {
			fmt.Println(e.Error())
		}

		//0-5
		//H e l l 0
		//0 1 2 3 4
		//bs[0:5]

		//fmt.Println(n)
		//fmt.Println(bs)
		reqstr := string(bs[:n])
		fmt.Println(reqstr)
		recvTime := time.Now().Format("2006-01-02 15:04:05")
		msg := fmt.Sprintf(`Your message: %s, received at %s`, reqstr, recvTime)
		conn.Write([]byte(msg))

	}
}