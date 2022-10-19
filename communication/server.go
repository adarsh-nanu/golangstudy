package communication

import (
	"fmt"
	"io"
	"net"
)

func Sever() {
	socket, err := net.Listen("tcp", ":2051")
	if err != nil {
		fmt.Println("some error in listen ", err)
		return
	}
	fmt.Println("listening..")
	connection, err := socket.Accept()
	if err != nil {
		fmt.Println("some error in accept ", err)
	}
	fmt.Println("accepted..", connection.RemoteAddr())
	/*
		bufferedreader := bufio.NewReader(connection)
		data, err := bufferedreader.ReadBytes('0')
		if err != nil {
			fmt.Println("some error in ReadByte ", err)
		}
		fmt.Println("Length ", len(data))
		for i := 0; i < len(data)-1; i++ {
			fmt.Printf("%0xf\n", data[i])
		}
	*/

	/*var arraylength []byte
	for {
		connection.Read(arraylength)
		//var arraydata []byte
	}
	*/
	tcpiplenarray := make([]byte, 2)
	//reader := bufio.NewReaderSize(connection, 2)
	//reader.Read(array)
	io.ReadFull(connection, tcpiplenarray)

	for i := 0; i < len(tcpiplenarray)-1; i++ {
		fmt.Printf("%0xf\n", tcpiplenarray[i])
	}
}
