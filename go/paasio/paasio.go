package paasio

import (
	"io"
)

type WC struct {
	writer io.Writer
}
type RC struct {
	reader io.Reader
}
type RWC struct {
	readwriter io.ReadWriter
}

func (w WC) WriteCount() (n int64, nops int){
	return 0, 0
}

func (w RWC) WriteCount() (n int64, nops int){
	return 0, 0
}

func (w RC) ReadCount() (n int64, nops int) {
	return 0, 0
}

func (w RWC) ReadCount() (n int64, nops int) {
	return 0, 0
}

func (w WC) Write(p []byte)(n int, err error) {
	return w.writer.Write(p)
}

func (w RWC) Write(p []byte)(n int, err error) {
	return w.readwriter.Write(p)
}

func (w RC) Read(p []byte) (n int, err error) {
	return w.reader.Read(p)
}

func (w RWC) Read(p []byte) (n int, err error) {
	return w.readwriter.Read(p)
}

func NewWriteCounter(w io.Writer) WriteCounter {
	wc := &WC{writer: w}

	return *wc
}

func NewReadCounter(r io.Reader) ReadCounter {
	wc := &RC{reader: r}

	return *wc
}

func NewReadWriteCounter(rw io.ReadWriter) RWC {
	wc := &RWC{readwriter: rw}

	return *wc
}
