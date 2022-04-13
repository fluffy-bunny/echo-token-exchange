package utils

import "strings"

func ToCanonical(src string) string {
	var replacer = strings.NewReplacer("\\", "/")
	str := replacer.Replace(src)
	return "file:///" + str
}

func TrimLeftAndRight(src string) string {
	return strings.TrimLeft(strings.TrimRight(src, " "), " ")
}
