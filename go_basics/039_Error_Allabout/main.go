package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

type ValidationError struct {
    Field string
    Msg   string
}

func (e *ValidationError) Error() string {
    return e.Field + ": " + e.Msg
}

func findUser(id int) error {
    if id < 0 {
        return &ValidationError{Field: "id", Msg: "must be positive"}
    }
    return ErrNotFound
}

func getUser(id int) error {
    if err := findUser(id); err != nil {
        return fmt.Errorf("getUser %d: %w", id, err)
    }
    return nil
}

func main() {
    err := getUser(5)
    fmt.Println(err)                          // getUser 5: not found
    fmt.Println(err == ErrNotFound)           // false
    fmt.Println(errors.Is(err, ErrNotFound))  // true

    err = getUser(-1)
    fmt.Println(err)                          // getUser -1: id: must be positive

    var ve *ValidationError
    if errors.As(err, &ve) {
        fmt.Println("bad field:", ve.Field)   // id
    }
}
