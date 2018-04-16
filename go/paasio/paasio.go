package paasio

import (
	"io"
	"sync"
)

type counter struct {
	rwBytes int64
	nops    int
	lock    sync.Mutex
}

type WC struct {
	writer io.Writer
	counter
}

type RC struct {
	reader io.Reader
	counter
}

type RWC struct {
	*WC
	*RC
	counter
}

func (c *counter) increment(n int64) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.nops++
	c.rwBytes += n
}

func (w WC) WriteCount() (int64, int) {
	return w.counter.rwBytes, w.counter.nops
}

func (w RC) ReadCount() (n int64, nops int) {
	return w.counter.rwBytes, w.counter.nops
}

func (w WC) Write(p []byte) (n int, err error) {
	return w.writer.Write(p)
}

func (w *RC) Read(p []byte) (n int, err error) {

	n, err = w.reader.Read(p)
	w.counter.increment(int64(n))
	return
}

func NewWriteCounter(w io.Writer) WriteCounter {
	wc := &WC{writer: w}

	return wc
}

func NewReadCounter(r io.Reader) ReadCounter {
	wc := &RC{reader: r}

	return wc
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	wc := &WC{writer: rw}
	rc := &RC{reader: rw}
	rwc := &RWC{WC: wc, RC: rc}

	return rwc
}
