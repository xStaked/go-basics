package message

import "fmt"

// privateFunc is a private function
func privateFunc() {
	fmt.Println("This is a private function")
}

// PublicFunc is a public function
func PublicFunc() {
	fmt.Println("This is a public function")
	// privateFunc()
}
