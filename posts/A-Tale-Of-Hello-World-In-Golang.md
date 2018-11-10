# A Tale Of Hello World In Golang

In this post we are going to explore what happens when you call `fmt.Println` to print something in your terminal window. So let's breakdown some basic examples!

### Common Example:

After searching about golang on search engines, you may have seen this example:

```Go
package main

import (
   "fmt"
)

func main(){
	fmt.Print("Hello World!")
}

```

### Simpler Example:

So why don't we be a bit simpler! Let's be simpler. This example doesn't event need a package to import! ;)

```Go
package main

func main(){
	print("Hello World!")
}
```

### Explained Example:

Now it's time to be complex and here is a bit more explained example:

```Go
package main

import (
   "os"
)

func main(){
	os.Stdout.Write([]byte("Hello World!"))
}

```

### How the code above works!

Let's think with a flashed brain. Now answer some question.

**What's the program giving to you?**

*Output.*

**What's the output?**

*"Hello World!"*

**Where are you getting the output?**

*In my terminal/standard-output window*

#### Now breakdown the code:

Somehow somewhere the `fmt.Println("Hello World!")` is telling your computer's terminal window to print `Hello World!`. So in the end `fmt.Println()` == `os.Stdout.Write()`. 

**Now talk about `os.Stdout`**

`os.Stdout` defines your os's standard output (aka: console, terminal, bla bla bla) methods. So when you call `fmt.Println` to print `Hello World!` it asks `os.Stdout` to print those things, then `os.Stdout` finds the best output method for you os/machine, and you get the `Hello World!`

```Txt
Hello World!
```

