package main

import (

	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"os"
	"strings"
	
)
func main()  {
	// how to use
if len(os.Args) != 2{
	fmt.Println("----------")
	fmt.Println("Error! Please enter a valid link.")
	fmt.Println("----------")
	fmt.Println("Usage:")
    fmt.Println(" example:")
	fmt.Println("go run main.go <link>")
	fmt.Println("example:go run main.go example.com ")
	fmt.Println("example:go run main.go http://example.com")
}
link := os.Args[1]
fmt.Println("link:",link)

if !strings.Contains(link, "://") {
    link = "http://" + link
}
parsedURL,err := url.Parse(link)
if err != nil {
	log.Fatalf("cann process %v",err)
}


	// Information extraction
	host := parsedURL.Host
	path := parsedURL.Path
	if path == "" {
		path = "/"
	}
		// Port identification
	port := "80"
	if parsedURL.Scheme == "https" {
		port = "443"
	}

		if strings.Contains(host, ":") {
    parts := strings.SplitN(host, ":", 2)
    host = parts[0]
    port = parts[1]
    }

	fmt.Println("Server :", host)
	fmt.Println("Path :", path)
	fmt.Println("port :", port)



conn,err := net.Dial("tcp", host+":"+port)
if err != nil {
	fmt.Println("cann connaced", err)
	return
}


defer conn.Close()
fmt.Println(" Connected to", host)


	request := fmt.Sprintf("GET %s HTTP/1.1\r\n", path)
	request += fmt.Sprintf("Host: %s\r\n", host)
	request += "User-Agent: Go-HTTP-Client/1.0\r\n"
	request += "Accept: */*\r\n"
	request += "Connection: close\r\n"
	request += "\r\n"
    
    fmt.Println(" The request is being processed...")
    fmt.Println("-Start order-")
    fmt.Print(request)
    fmt.Println("-Order End-\n")

    // Submit the request
    _, err = conn.Write([]byte(request))
    if err != nil {
        log.Fatalf(" Transmission failed!!: %v", err)
    }

    // Print Reply
    fmt.Println(" Server reply:")
    fmt.Println("-Beginning reply-")
    _, err = io.Copy(os.Stdout, conn)
    if err != nil {
        log.Fatalf("Reading failure: %v", err)
    }
    fmt.Println("\n-End reply-")
}
