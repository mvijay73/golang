package main

// Run checks overwrite flag and generates logs with given options
func main() {
	option := &Option{
		Format: "Common",
		Output: "/tmp/sample.log",
		Type:   "log",
		Number: 10000,
	}
	Generate(option)
}
