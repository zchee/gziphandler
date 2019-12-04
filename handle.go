package gziphandler

import "mime"

type handleType int

const (
	handleNone handleType = iota
	handleGzip
	handleBrotli
)

// returns how we should handle the request
func handleContentType2(gzipContentTypes, brotliContentTypes []parsedContentType, ct string, accept acceptsType) handleType {
	preferGzip := true
	if preferGzip {
		if accept.gzip() && handleContentType(gzipContentTypes, ct) {
			return handleGzip
		}
		if accept.brotli() && handleContentType(brotliContentTypes, ct) {
			return handleBrotli
		}
	} else {
		if accept.brotli() && handleContentType(brotliContentTypes, ct) {
			return handleBrotli
		}
		if accept.gzip() && handleContentType(gzipContentTypes, ct) {
			return handleGzip
		}
	}
	return handleNone
}

// returns true if we've been configured to compress the specific content type.
func handleContentType(contentTypes []parsedContentType, ct string) bool {
	// If contentTypes is empty we handle all content types.
	if len(contentTypes) == 0 {
		return true
	}

	mediaType, params, err := mime.ParseMediaType(ct)
	if err != nil {
		return false
	}

	for _, c := range contentTypes {
		if c.equals(mediaType, params) {
			return true
		}
	}

	return false
}
