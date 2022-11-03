package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

const _HOST = "localhost"
const _PORT = 3333
const _CLIENT_DELAY = 0 * time.Millisecond
const _CLIENT_COUNT = 10

var startTime = time.Now()

var echoCount = 0 // number of echos
var echoMutex sync.Mutex

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func sender() {

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", _HOST, _PORT))
	panicIf(err)

	for {

		// thing to echo
		var id = uuid.New().String()

		// send
		_, err := conn.Write([]byte(id + "\n"))
		panicIf(err)

		// receive
		var reader = bufio.NewReader(conn)
		response, err := reader.ReadString('\n')
		panicIf(err)

		// cleanup
		response = strings.TrimSpace(response)

		// check if its working
		if response != id {
			panic(fmt.Sprintf(">%s< != >%s<", response, id))
		}

		echoMutex.Lock()
		echoCount++
		echoMutex.Unlock()

		time.Sleep(_CLIENT_DELAY)

	}

}

func main() {

	for i := 0; i < _CLIENT_COUNT; i++ {
		go sender()
	}

	for {
		rps := float64(echoCount) / time.Since(startTime).Seconds()

		log.Printf("client count: %d, delay: %v, echos per second: %f\n",
			_CLIENT_COUNT, _CLIENT_DELAY, rps)

		time.Sleep(time.Second)
	}

}
