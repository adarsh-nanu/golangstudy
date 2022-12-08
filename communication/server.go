package communication

import (
	"fmt"
	"io"
	"net"
)

func Sever() {
	fmt.Println("Inside server ")
	serversocket, err := net.Listen("tcp", ":2051")
	if err != nil {
		fmt.Println("some error in listen ", err)
		return
	}
	fmt.Println("listening..")
	connection, err := serversocket.Accept()
	if err != nil {
		fmt.Println("some error in accept ", err)
	}
	fmt.Println("accepted..", connection.RemoteAddr())

	tcpiplenarray := make([]byte, 2)
	io.ReadFull(connection, tcpiplenarray)

	for i := 0; i <= len(tcpiplenarray)-1; i++ {
		fmt.Printf("%0x\n", tcpiplenarray[i])
	}
	serversocket.Close()
}

func Client() {
	sendBuffer := make([]byte, 4)
	sendBuffer[0] = 0x31
	sendBuffer[1] = 0x33
	clientsocket, err := net.Dial("tcp", "127.0.0.1:2051")
	if err != nil {
		fmt.Println("some error in dial ", err)
		panic("error")
	}
	clientsocket.Write(sendBuffer)
	clientsocket.Close()
}
