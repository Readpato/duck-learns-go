# Hello, world!

## Executing a function and testing it

This is the first exercise in the course and it' was actually pretty fun!

Coming from the JavaScript/TypeScript ecosystem it's interesing to see how other languages handle their own syntax and interactions with the language.

Given that Go is a language that normally is used to write Backend or CLI apps, you spend a lot of time in the terminal, which for me sounds amazing!

Firstly, we started our first file with the following contents:

```go
package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
```

Here we can se already very new things:

```go
package main
```

This line of code is a declaration of the package name. In Go, every file belongs to a package, and the package name is used to organize and group related code. Since this is our main file (and the only one), we name it main

```go
import "fmt"
```

Here we are importing the fmt package, which comes from the standard [Go library](https://pkg.go.dev/fmt). It provides functions for formatting and printing text. We use this package to print the output of our Hello function.

```go
const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
```

Here we have our `main` function. This is the entry point of our program. When we run our program, the `main` function will be executed. Which will call the `Hello` function with the argument "world" and print: `"Hello, world" to the console.

We have also declared a constant `englishHelloPrefix` which is used to prefix the English greeting. Since we know that the greeting in English will never change, we declare it as a constant `englishHelloPrefix`. As any other language, the constant cannot be modified.

If you would like to run the program, then you would now execute `go run hello.go` in the folder where it's located. And you would see something like

```zsh
$ go run hello.go
Hello, world
```

We just run our first Go program!

Now, as the title of the course states, we will now test the function. From this point onwards the course we are told we will be testing first and then writing code following the TDD (Test-Driven Development) approach.

```go
package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, world' when there's no provided name", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
```

Woah, so many new shiny things! Let's start by analyzing each piece of it:

```go
package main

import "testing"
```

Here we are naming our package `main` because it allows us to access functions and variables within that package!

Then, we import another standard library module [`testing`](https://pkg.go.dev/testing), which allows us to enter all the baked in goodies from the Go language to test.

Then, we have our nice test function:

```go
func TestHello(t *testing.T) {
	t.Run("Say hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, world' when there's no provided name", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
```

Our main function is called `TestHello`, which takes only one parameter `t` of type `*testing.T` which comes directly from the testing library. This was called an `interface` in the course. I wonder if it will behave the same way as an `interface` does in TypeScript. We will found out in the future :)

Next, we see that we are immediately utilizing our `t` parameter by calling `t.Run` this allows us to run subtests inside of our main test function. This comes super handy when we want to test multiple scenarios within a single test function.

```go
	t.Run("Say 'Hello, world' when there's no provided name", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
```

Inside of our function we have two declared variables `got` and `want` which are of type `string`. These variables are used to store the result of the `Hello` function and the expected output respectively. We then call the `assertCorrectMessage` function to compare the two values and assert that they are equal.

```go
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
```

In this tiny helper function we have so many good things to notice! Firstly, we are using the `t.Helper()` function which is a helper function that allows us to set the helper flag on the test function. This is useful when we want to call a helper function from within a test function. Without it we would not be able to see the correct line number in our test output when an error occurs.

Then, we have the `t.Errorf` function which is a helper function that allows us to print an error message to the test output. This is useful when we want to print an error message to the test output when an error occurs. Noticed the `%q`? Keen eye, this means that the value will be printed in quotes! If you want to know more of the different types of formatting, check it out [here](https://pkg.go.dev/fmt)!

### Next steps

If you check out the current `hello.go` file, you can see that the contents have changed and I have expanded on it given the course instructions, give it a try if you think you would like to expand the function from the state in this README to the one in the file and more importantly, dont forget the [TDD and course rules](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world#discipline):

- Write a test
- Make the compiler pass
- Run the test, see that it fails and check the error message is meaningful
- Write enough code to make the test pass
- Refactor

### Cool Go things that I loved

- Declaring a variable inline and assigning a value to it `got := "Hello, world"`
- Creating a variable inline in a function! (Here's `prefix` of type string)

```go
	func greetingPrefix(language string) (prefix string) {
```

- Declaring multiple constants like this:

```go
const (
	french  = "french"
	german  = "german"
	spanish = "spanish"
	...
)
```

- Having an integrated testing framework! That's so crazy. Coming from the JS/TS ecosystem, having a default way of doing things sounds amazing.
