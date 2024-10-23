Go

Goroutine: A goroutine is a lightweight thread managed by the Go runtime. Goroutines run concurrently with other goroutines and are extremely cheap compared to traditional threads.
 Simple example of grouting
func sayHello() {
	fmt.Println("Hello, World!")
	time.Sleep(2 * time.Second)
}

func main() {
	go sayHello() // Start a new goroutine
}  The above program won’t print anything because the main() function execution completes before the execution of the sayHello() function 
 Channel: Channels are a way to communicate between goroutines and synchronize their execution. Channels can be thought of as pipes that connect concurrent goroutines, allowing them to send and receive values. 
Simple example of channel func sum(numbers []int, result chan int) {
    sum := 0
    for _, number := range numbers {
        sum += number
    }
    result <- sum  // send sum to channel
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    result := make(chan int)

    go sum(numbers, result)

    total := <-result  // receive from channel
    fmt.Println("Sum:", total)
}   Defer: The defer statement schedules a function call to be run after the function completes. This is commonly used for cleanup purposes, such as closing a file or releasing a lock.  Unbuffered channel: An unbuffered channel is a channel with no capacity to hold values; it has a capacity of zero. When a value is sent to an unbuffered channel, the send operation blocks until another goroutine receives the value from the channel. Similarly, a receive operation on an unbuffered channel blocks until another goroutine sends a value.  Buffered channel:  A buffered channel has a specified capacity to hold values. When a value is sent to a buffered channel, the send operation blocks only if the channel is full. Conversely, a receive operation blocks only if the channel is empty.  Garbage Collection: Garbage collection (GC) in Go is designed to manage memory efficiently by automatically reclaiming memory that is no longer in use, thus preventing memory leaks and other related issues. Go's garbage collector is concurrent, meaning it runs alongside the program, which helps minimise pauses and latency.  Interface: Interfaces in Go provide a way to specify the behaviour expected of an object. an interface is a type that specifies a set of method signatures. If a type provides definitions for all the methods in the interface, it implicitly implements that interface.  Example of Interfaces for Mocking
Consider you have a service that interacts with an external API. You can define an interface for the API client and then create a mock implementation for testing.
type APIClient interface {
    GetData(id string) (string, error)
}

type RealAPIClient struct{}

func (c RealAPIClient) GetData(id string) (string, error) {
    // Actual API call
    return "real data", nil
}

type MockAPIClient struct{}

func (c MockAPIClient) GetData(id string) (string, error) {
    // Mocked data
    return "mock data", nil
}

func FetchData(client APIClient, id string) (string, error) {
    return client.GetData(id)
}

func main() {
    realClient := RealAPIClient{}
    mockClient := MockAPIClient{}

    data, _ := FetchData(realClient, "123")
    fmt.Println("Real client data:", data)

    data, _ = FetchData(mockClient, "123")
    fmt.Println("Mock client data:", data)
}


Closure: Closure is a function value that references variables from outside its body. The function can access and modify these variables even after the surrounding function has returned.  // MakeFilter returns a closure that checks if a number is greater than the given threshold
func MakeFilter(threshold int) func(int) bool {
    return func(n int) bool {
        return n > threshold
    }
}

func main() {
    filter := MakeFilter(10)

    numbers := []int{5, 10, 15, 20}
    for _, number := range numbers {
        if filter(number) {
            fmt.Println(number) // Output: 15 20
        }
    }
}

Slice: Slice is a flexible, dynamic data structure that represents a segment of an array. Unlike arrays, slices can grow and shrink dynamically.  Mutex in golang: sync.Mutex  JSON marshalling and unmarshalling are essential concepts in Go programming. They are used to convert Go data structures into JSON format and vice versa.  

