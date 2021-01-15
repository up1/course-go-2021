package main

import "fmt"

// File
type file struct {
	name string
}

func (file) read(b []byte) (int, error) {
	s := "<rss><channel><title>Going Go Programming</title></channel></rss>"
	copy(b, s)
	return len(s), nil
}

// Pipe
type pipe struct {
	name string
}

func (pipe) read(b []byte) (int, error) {
	s := `{name: "bill", title: "developer"}`
	copy(b, s)
	return len(s), nil
}

// Main
func main() {

	// Create two values one of type file and one of type pipe.
	f := file{"data.json"}
	p := pipe{"cfg_service"}

	// Call each retrieve function for each concrete type.
	retrieveFile(f)
	retrievePipe(p)
}

// retrieveFile can read from a file and process the data.
func retrieveFile(f file) error {
	data := make([]byte, 100)

	len, err := f.read(data)
	if err != nil {
		return err
	}

	fmt.Println(string(data[:len]))
	return nil
}

// retrievePipe can read from a pipe and process the data.
func retrievePipe(p pipe) error {
	data := make([]byte, 100)

	len, err := p.read(data)
	if err != nil {
		return err
	}

	fmt.Println(string(data[:len]))
	return nil
}
