// $ 6g echo.go && 6l -o echo echo.6
// $ ./echo
//
//  ~ in another terminal ~
//
// $ nc localhost 3540

package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

const PORT = 3540

type simulator struct {
	commands map[string]func() interface{}
	port int
}

func NewSimulator(port int) *simulator {
	p := new(simulator)
	p.port = port
	p.commands = make(map[string]func() interface{})
	p.commands["exit\n"] = p.done
	return p
}

func main() {
	/*
	server, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	if server == nil {
		panic(fmt.Sprintf("echo: failed to listen on %v: %v", server, err))
	}
	conns := clientConns(server)
	for {
		go handleConn(<-conns)
	}
	*/
	sim := NewSimulator(3540)
	sim.connect()

}

func (v *simulator) connect() {

	server, err := net.Listen("tcp", ":"+strconv.Itoa(v.port))
	if server == nil {
		panic(fmt.Sprintf("echo: failed to listen on %v: %v", server, err))
	}
	conns := v.clientConns(server)
	for {
		go v.handleConn(<-conns)
	}

}

func (v *simulator) clientConns(listener net.Listener) chan net.Conn {
	ch := make(chan net.Conn)
	i := 0
	go func() {
		for {
			client, err := listener.Accept()
			if client == nil {
				//fmt.Printf("couldn't accept: " + err.String())
				fmt.Printf("echo: failed to listen on %v: %v", client, err)
				continue
			}
			i++
			fmt.Printf("%d: %v <-> %v\n", i, client.LocalAddr(), client.RemoteAddr())
			ch <- client
		}
	}()
	return ch
	}

	func (v *simulator) handleConn(client net.Conn) {
		b := bufio.NewReader(client)
		for {
			line, err := b.ReadBytes('\n')
			if err != nil { // EOF, or worse
				break
			}

			fin, ok := v.commands[string(line)]
			if ok {
				ret := fin().(string)
				client.Write([]byte(ret))
				client.Close()
			}

			client.Write(line)
		}
	}

	func (v *simulator) done() interface{} {
		return "done"
	}

/*
func clientConns(listener net.Listener) chan net.Conn {
	ch := make(chan net.Conn)
	i := 0
	go func() {
		for {
			client, err := listener.Accept()
			if client == nil {
				//fmt.Printf("couldn't accept: " + err.String())
				fmt.Printf("echo: failed to listen on %v: %v", client, err)
				continue
			}
			i++
			fmt.Printf("%d: %v <-> %v\n", i, client.LocalAddr(), client.RemoteAddr())
			ch <- client
		}
	}()
	return ch
}

func handleConn(client net.Conn) {
	b := bufio.NewReader(client)
	for {
		line, err := b.ReadBytes('\n')
		if err != nil { // EOF, or worse
			break
		}
		client.Write(line)
	}
}
*/
