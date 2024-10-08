package main

import (
	//	"net/rpc"
	"flag"
	"net/rpc"
	"uk.ac.bris.cs/distributed2/secretstrings/stubs"

	//	"bufio"
	//	"os"
	//	"uk.ac.bris.cs/distributed2/secretstrings/stubs"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/**func questions() []string {
	f, err := os.Open("wordlist.csv")
	check(err)
	reader := csv.NewReader(f)
	table, err := reader.ReadAll()
	check(err)
	var questions []string
	for _, row := range table {
		questions = append(questions, row[0])
	}
	return questions
}**/

func main() {
	server := flag.String("server", "127.0.0.1:8030", "IP:port string to connect to as server")
	flag.Parse()
	fmt.Println("Server: ", *server)
	//TODO: connect to the RPC server and send the request(s)
	client, _ := rpc.Dial("tcp", *server)
	defer client.Close()
	//qs := questions()
	/**for _, q := range qs {


	}**/
	request := stubs.Request{"Hello"}
	response := new(stubs.Response)
	client.Call(stubs.PremiumReverseHandler, request, response)
	fmt.Println("Responded: ", response.Message)
}
