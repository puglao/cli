package version

import "fmt"

var (
	tag string = ""
	sha string = ""
)

func Tag() string {
	if tag == "" {
		return "dev"
	}
	return tag
}

func Sha() string {
	if sha == "" {
		return "unknown"
	}
	return sha
}

func Info() string {
	return fmt.Sprintf("%s (%s)\n", Tag(), Sha())
}
