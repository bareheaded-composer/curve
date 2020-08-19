package model

const PhotoDirName = "photo_pool"
const PhotoMaxSize = 5 * Mb

var ValidPhotoType = []string{
	"jpg", "png", "bmp",
}

const ThumbnailWidth = 200
const ThumbnailHeight = 0