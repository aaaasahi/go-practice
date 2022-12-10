package foo

const (
	Max = 100 // 大文字Public
	min = 1   // 小文字はPrivate
)

func ReturnMin() int {
	return min
}
