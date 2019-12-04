package gziphandler

import (
	"io"
	"sync"

	"github.com/andybalholm/brotli"
)

var brotliWriterPools [brotli.BestCompression - brotli.BestSpeed]sync.Pool

func init() {
	for i := brotli.BestSpeed; i <= brotli.BestCompression; i++ {
		addBrotliLevelPool(i)
	}
}

func brotliPoolIndex(level int) int {
	return level // duh
}

func addBrotliLevelPool(level int) {
	brotliWriterPools[brotliPoolIndex(level)].New = func() interface{} {
		return brotli.NewWriterLevel(nil, level)
	}
}

func getBrotliWriter(w io.Writer, level int, handle handleType) *brotli.Writer {
	bw, _ := brotliWriterPools[brotliPoolIndex(level)].Get().(*brotli.Writer)
	bw.Reset(w)
	// TODO: use BROTLI_MODE_TEXT and BROTLI_MODE_FONT
	return bw
}

func putBrotliWriter(bw *brotli.Writer, level int) {
	bw.Reset(nil)
	brotliWriterPools[brotliPoolIndex(level)].Put(bw)
}
