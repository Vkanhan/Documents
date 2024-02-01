package employee

import (
	"fmt"
)

type employee struct {  
    firstName   string
    LastName    string
    TotalLeave int
    LeaveTaken int
}

// New is a factory function to create and initialize a new employee instance.
func New(firstName string, LastName string, TotalLeave int, LeaveTaken int) employee {
    // Creating a new employee instance with provided values.
    e := employee {firstName, LastName, TotalLeave, LeaveTaken}
    return e
}

func (e employee) LeaveRemaining() {  
    fmt.Printf("%s %s has %d leave remaining\n", e.firstName, e.LastName, (e.TotalLeave - e.LeaveTaken))
}
