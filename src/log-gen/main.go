package main

// Run checks overwrite flag and generates logs with given options
func main() {
	option := &Option{
		Format: "Common",
		Output: "sample.log",
		Type:   "log",
		Number: 1000,
	}
	Generate(option)
}
