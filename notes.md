## Hardcore Go

* In a **struct** have the largest field first, this helps with memory management. Ref: https://play.golang.org/p/TAX6NpPaEu
* No relative paths for imports. Only absolute
* Using build is better than run since you then dont have to define exceptions, like not running test cases, also if you have to run 8 files simultaneously you can just create a single .exe file instead of running 8
* The value of a pointer variable is ALWAYS an address
* Constants are compile type, they only exists at compile time.
* The constant can be of a type, and can then no longer be implicitly converted
* There are no constructors in GO. Finalizers are also not recommended. The factory pattern is used a LOT
* Statements: if, for, switch, while declare their own blocks with their own scope.
* ?? Difference between len and cap in case of slices
	* For the first 1000, for every single length increment/append the array's capacity doubles and then its a % value
	* Anytime the length and capacity are ot the same, the capacity can be brought in to length, during append operations. Cap is always > length
* rune : is an alias for int32
* "Every array in Go is a slice waiting to happen"
* Never name a receiver 'this' or 'self'
* ?? What is escape analysis, how is it determined if a value goes in stack or in heap??
* When a function has a receivers its now called and treated as a method.
	* The receivers is really the first parameter to the method
* Interfaces should only have one method. Always.
* We save concrete type values that satisfy the interface in interface
* ?? When to implement interface using pointers then only pointers receivers satisfy the interface, but if you use value then both value and pointers satisfy ??
* A Go program is constructed as a "package", which may in turn use facilities from other packages.
	* The basic unit of compilation in Go is a package
	* Two packages cannot include each other. ?? (this is weird!)
	* Every standalone Go program contains a package called main and its main function, after initialization, this is where execution starts. The main.main function takes no arguments and returns no value.
* You can also have a func init function along with main, the init is always called before the main function. You can also have muleiple init functions, if the need be Butt very rare, depends if your app is that complicated
	* If you have multiple init() functions, they are evaluated in lexical ordering
* If you ever see something like this:
```
import (
	"fmt"
	_ "github...""
)
```
that means that you may not need the github.. package but coz there is a init function somewhere in that file/package that may need that import pacakge initialized
* https://github.com/luciotato/golang-notes/blob/master/OOP.md
* Variable that starts with an underscore it is considered as an unexported identifier
* ?? What are Go concrete types ??
* Type assertions are a runtime and not compile time construct.
	* You cannot do type assertions on unexported values.
	* All error interfaces should be using pointer receivers during implementation:
	```
	// http://golang.org/src/pkg/errors/errors.go
	func (e *errorString) Error() string {
		return e.s
	}
	```
* Go needs trailing comma before newline in composite literal. Eg:
```
techfirm := company{
		administrator: admins.popAdmin(),
		developer:     devs.popDev()
	}
```
Above will generate error. It needs to be:
```
techfirm := company{
		administrator: admins.popAdmin(),
		developer:     devs.popDev(),
	}
```
* An error in Go is nothing more than a value
	* And so an error can be anything that you need it to be
	* Only have custom error types when they are absolutely necessary
	* Anytime you are dealing with errors, you are actually dealing with an interface. This is 	how it is internally:
	```
	// http://golang.org/pkg/builtin/#error
	type error interface {
		Error() string
	}
	```
	* If an error
* 'nil' auto takes the value acc to what data type that we are comparing it against
* Look into type assertions between interface types
* **Concurrency** - managing a lot of things at once
* **Parallelism** - not necessarily managing but doing a lot of things at once
* Problem with **Global Variables** - Compiler can put them anywhere it wants, and chances are that they'll end up on the same cache line, causing speed issues!
* **Data race** when goroutines are not synchronized and are fighting with each other. Go has a race detector, ```go build -race``` if this detects a race, you definitely have a race, if no race detected you may still ahve a race! On detection of race, it would generate a panic and o/p the stack race
* **Race conditions** occur when two or more **Goroutines** try to use a piece of shared data at the same time, but the result of their operations is dependent on the exact order that the scheduler executes their instructions. While Go's concurrency mechanisms make it easy to write clean concurrent code, they don't prevent race conditions.
* Go provides a number of tools to help us avoid data races.
	* These include Channels for communicating data between Goroutines, a Race Detector for monitoring unsynchronized access to memory at runtime, and a variety of 'locking' features in the Atomic and Sync packages.
	* One of these features are Mutual Exclusion locks, or **mutexes**. For complex state we can use a mutex to safely access data across multiple goroutines.
* There are 2 types of channels: **Buffered** and **Unbuffered***
	* **Unbuffered** :
	* **Buffered** channel comes with data loss, and is not 100% predictable
* Use testing as a way to validate the API
	* Test for what the user would be using
	* You should only test **Exported** values of a package. So your test package should only worry about exported packages.
	* Your test function must start with **'T'** caps T, must be exported! The second word after Test should also start with capslock. So eg: **TestDownload** or else it wont execute.
	* ```go test``` runs all your test functions in teh package. However if your really need to run one specific test function: ```go test -v -run TestFile ```


#### Receiver Choices
##### How to decide if something needs to be shared or not, to have a value reciever or a pointer reciever in a method
* Never share primitive data types, it not the nature of that type. Like int, string, bool etc
	* So no pointers!!!
	* They dont need to be shared, a copy of it should be available to all
	* So declaring a new type using one of these primitive datatypes, then that is also a 	primitive type and does not need to be shared. Like type bill string
	* Same with reference data type, like following,
	```
	type IP []byte
	```
	* Exception, if a function is implementing an interface on unmarshalling method, you will need to use a pointer reciever.
	* When you have a factory function that is providing you a pointer,it is a good sign that you should not be making copies of the referenced value being returned. You should just share it. Ref: https://play.golang.org/p/xD6PCx--GG