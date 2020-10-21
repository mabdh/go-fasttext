package fasttext

// #cgo CXXFLAGS: -std=c++11 -march=native -g -Wall -I/usr/local/include/fastText
// #cgo LDFLAGS: -L/usr/local/lib -lfasttext
// #include "gofasttext.h"
// #include <stdlib.h>
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

type Result struct {
	Label string
	Prob  float32
}

type Model interface {
	Predict(word string, k int, threshold float32) []Result
	Free()
}

type model struct {
	ft C.GoFastText
}

func New(path string) (Model, error) {
	var m model
	pathChar := C.CString(path)
	defer C.free(unsafe.Pointer(pathChar))
	m.ft = C.GoFastTextInit(pathChar)
	if m.ft == nil {
		return nil, errors.New("Error create fast-text model")
	}
	fmt.Println(m.ft)
	return m, nil
}

func (m model) Free() {
	C.GoFastTextFree(unsafe.Pointer(m.ft))
}

func (m model) Predict(word string, k int, threshold float32) []Result {
	wordChar := C.CString(word)
	defer C.free(unsafe.Pointer(wordChar))

	var cResultLen C.int
	var res *C.go_fast_text_pair_t = C.GoFastTextPredict(unsafe.Pointer(m.ft), wordChar, C.int(k), C.float(threshold), &cResultLen)

	resultLen := int(cResultLen)
	fmt.Println(resultLen)

	pairSlice := (*[1 << 28]C.go_fast_text_pair_t)(unsafe.Pointer(res))[:resultLen:resultLen]
	return pairResultCToGo(pairSlice)
}

func pairResultCToGo(pairSlice []C.go_fast_text_pair_t) []Result {
	predictionResult := []Result{}
	for _, cstruct := range pairSlice {
		r := Result{
			Label: C.GoString(cstruct.label),
			Prob:  float32(cstruct.prob),
		}
		predictionResult = append(predictionResult, r)
	}
	return predictionResult
}
