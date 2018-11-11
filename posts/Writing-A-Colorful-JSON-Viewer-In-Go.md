# Writing A Colorful JSON Viewer In Go

In this post we are going to create a `jq` like JSON highlighter for viewing colorful JSON in terminal!

## Before we begin:
Before we begin, please be sure that you have enough knowledge in JSON, Go, Terminal, Terminal Colors e.t.c. If you don't have enough knowledge, don't worry about it! 

### Let's be cool without any actual code:

Run this code in your terminal:

```sh
$ echo -e "\e[101mI am light red"
```

Can you see the text with a red background? That's what we are going to make!

Now try to change the value of the `\e[101`, example:

```sh
$ echo -e "\e[102mI am green"
```

Also 


```sh
$ echo -e "\e[103mI am yellow"
```

### Colors in Go:

We've seen how to get colors in terminal.But we've done with `bash`'s `echo`. But now we should give a try with Go's `fmt.Println`.

```Go
package main

import (
    "fmt"
)

func main(){
    fmt.Println("\033[95mHello Colors!")
}
```

The output should be `Hello Colors!` in magenta color!

## The JSON things:

We've done our first step *Colors*. Now we have to do the second step *JSON parsing*.

We have to parse JSON data so that we can detect the keys and value separately. And now it's time to do that!

Here is the sample data we will work on:

```JSON
{"Name":"Anik"}
```

And we've to do it like:

```TXT
{"(START_MAGENTA_COLOR)Name(END_MAGENTA_COLOR)":"(START_BLUE_COLOR)Anik(END_OF_BLUE_COLOR)"}
```

Looks complex? Ah, ok! So the task is:

* Find a key.
* Start magenta color.
* Print the key. (`Name`)
* Stop coloring.
* Find a value.
* Start blue color.
* Print the value. (`Anik`)
* End coloring.

## Let's code and comment!

Here is the final code we are working on! We're taking the cyan color as blue :p

Define the colors

```Go
const Magenta = 95
const Blue = 94
EndColor = "\033[0m" // This will take us to the default white color or no color
```

Now define some color functions:

```Go
// This function will start the color by printing the color values we've set before
func ColorStart(color uint8) string {
        return fmt.Sprintf("\033[%dm", color)
}

// This function will print the colored text
// Ex: Color_Value+JSON_Value+End_Coloring
func Color(str string, color uint8) string {
        return fmt.Sprintf("%s%s%s", ColorStart(color), str, EndColor)
}
```

Ok! so now we have done with colors, let's parse json!

We are going to check each character one by one. Let's code it.

```Go
func ColorIt(str string) string {
        var rsli []rune
        var key, val, startcolor, endcolor, startsemicolon bool
        var prev rune
	// loop through each character
        for _, char := range []rune(str) {
                switch char {
                case ' ':
                        rsli = append(rsli, char)
                case '{':
                        startcolor = true
                        key = true
                        val = false
                        rsli = append(rsli, char)
                case '}':
                        startcolor = false
                        endcolor = false
                        key = false
                        val = false
                        rsli = append(rsli, char)
                case '"':
                        if startsemicolon && prev == '\\' {
                                rsli = append(rsli, char)
                        } else {
                                if startcolor {
                                        rsli = append(rsli, char)
                                        if key {
                                            rsli = append(rsli, []rune(ColorStart(Magenta))...)
                                        } else if val {
                                            rsli = append(rsli, []rune(ColorStart(Cyan))...)
                                        }
                                        startsemicolon = true
                                        key = false
                                        val = false
                                        startcolor = false
                                } else {
                                        rsli = append(rsli, []rune(EndColor)...)
                                        rsli = append(rsli, char)
                                        endcolor = true
                                        startsemicolon = false
                                }
                        }
                case ',':
                        if !startsemicolon {
                                startcolor = true
                                key = true
                                val = false
                                if !endcolor {
                                        rsli = append(rsli, []rune(EndColor)...)
                                        endcolor = true
                                }
                        }
                        rsli = append(rsli, char)
                case ':':
                        if !startsemicolon {
                                key = false
                                val = true
                                startcolor = true
                                if !endcolor {
                                        rsli = append(rsli, []rune(EndColor)...)
                                        endcolor = true
                                }
                        }
                        rsli = append(rsli, char)
                case '\n', '\r', '[', ']':
                        rsli = append(rsli, char)
                default:
                        if !startsemicolon {
                                if key && startcolor {
                                        rsli = append(rsli, []rune(ColorStart(Magenta))...)
                                        key = false
                                        startcolor = false
                                        endcolor = false
                                }
                                if val && startcolor {
                                        rsli = append(rsli, []rune(ColorStart(Cyan))...)
                                        val = false
                                        startcolor = false
                                        endcolor = false
                                }
                        }
                        rsli = append(rsli, char)
                }
                prev = char
        }
        return string(rsli)
}
```

And we have done!

## Full code:

Here is the full version of the code. Taken from `astaxie/bat`.

```Go
package main

import (
	"fmt"
)

const (
	Gray = uint8(iota + 90)
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White

	EndColor = "\033[0m"
)

func Color(str string, color uint8) string {
	return fmt.Sprintf("%s%s%s", ColorStart(color), str, EndColor)
}

func ColorStart(color uint8) string {
	return fmt.Sprintf("\033[%dm", color)
}

func ColorIt(str string) string {
	var rsli []rune
	var key, val, startcolor, endcolor, startsemicolon bool
	var prev rune
	for _, char := range []rune(str) {
		switch char {
		case ' ':
			rsli = append(rsli, char)
		case '{':
			startcolor = true
			key = true
			val = false
			rsli = append(rsli, char)
		case '}':
			startcolor = false
			endcolor = false
			key = false
			val = false
			rsli = append(rsli, char)
		case '"':
			if startsemicolon && prev == '\\' {
				rsli = append(rsli, char)
			} else {
				if startcolor {
					rsli = append(rsli, char)
					if key {
						rsli = append(rsli, []rune(ColorStart(Magenta))...)
					} else if val {
						rsli = append(rsli, []rune(ColorStart(Cyan))...)
					}
					startsemicolon = true
					key = false
					val = false
					startcolor = false
				} else {
					rsli = append(rsli, []rune(EndColor)...)
					rsli = append(rsli, char)
					endcolor = true
					startsemicolon = false
				}
			}
		case ',':
			if !startsemicolon {
				startcolor = true
				key = true
				val = false
				if !endcolor {
					rsli = append(rsli, []rune(EndColor)...)
					endcolor = true
				}
			}
			rsli = append(rsli, char)
		case ':':
			if !startsemicolon {
				key = false
				val = true
				startcolor = true
				if !endcolor {
					rsli = append(rsli, []rune(EndColor)...)
					endcolor = true
				}
			}
			rsli = append(rsli, char)
		case '\n', '\r', '[', ']':
			rsli = append(rsli, char)
		default:
			if !startsemicolon {
				if key && startcolor {
					rsli = append(rsli, []rune(ColorStart(Magenta))...)
					key = false
					startcolor = false
					endcolor = false
				}
				if val && startcolor {
					rsli = append(rsli, []rune(ColorStart(Cyan))...)
					val = false
					startcolor = false
					endcolor = false
				}
			}
			rsli = append(rsli, char)
		}
		prev = char
	}
	return string(rsli)
}

func main() {
	input := `{"Name":"Anik"}`
	fmt.Println(ColorIt(input))
}
```

The colored output should be:

```TXT
{"Name":"Anik"}
```

*My blog doesn't support colored view :p*
