package gziphandler

type preferType int

const (
	preferGzip preferType = iota
	preferBrotli
	preferClientThenGzip
	preferClientThenBrotli
)

// returns true if we should try gzip first
func (p preferType) negotiate(a acceptsType) bool {
	if p == preferClientThenGzip || p == preferClientThenBrotli {
		switch a {
		case acceptsBrotliThenGzip, acceptsBrotli:
			return false
		case acceptsGzipThenBrotli, acceptsGzip:
			return true
		case acceptsGzipAndBrotli:
			return p == preferClientThenGzip
		}
	}
	return p == preferGzip
}
