package model

const InvalidFileType = ""

const (
	b = 1 << (iota * 10)
	Kb
	Mb
	Gb
)

const AvatarMaxSize = 5 * Mb
const InvalidFileName = ""

var ValidAvatarType = [...]string{
	"jpg", "png", "bmp",
}
