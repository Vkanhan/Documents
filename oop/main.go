package main

import "oop/employee"

func main() {
	// Creating a new employee using the New factory function.
	e := employee.New("Sam", "Adolf", 30, 20)
	e.LeaveRemaining()
}



//wrong
// package main

// import "oop/employee"

// func main() {
// 	e := employee.Employee {
// 		FirstName: "Sam",
// 		LastName: "Adolf",
// 		TotalLeaves: 30,
// 		LeavesTaken: 20,
// 	}
// 	e.LeavesRemaining()
// }
