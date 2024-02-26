// by convention, we name our package the package name of the file
package mystrings

// Reverse reverses a string left to righ
// Notice that we need to capitalize the first letter
// If we don't then we won't be able to access this
// mystrings package

func Reverse(s string) string {
	result := ""
	for _, v := range s {
		result = string(v) + result
	}
	return result
}
