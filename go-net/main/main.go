package main

type MyError struct {
	Message string
	Err     error
}

func (e *MyError) Error() string {
	return e.Message
}

func (e *MyError) Unwrap() error {
	return e.Err
}

type Item struct {
	A int `json:"a"`
}

type ItemMap map[string]Item

func main() {}
