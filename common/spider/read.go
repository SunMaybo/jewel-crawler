package spider

import (
	"bytes"
	"io"
	"unicode/utf8"
)

//先看一下Buffer的定义，有帮助下面理解
type Buffer struct {
	buf       []byte            // 数据存放在 buf[off : len(buf)]
	off       int               // 从&buf[off]开始读, 从&buf[len(buf)]开始写
	runeBytes [utf8.UTFMax]byte // avoid allocation of slice on each WriteByte or Rune
	bootstrap [64]byte
}

func ReadAll(r io.Reader, size int) (buff []byte, err error) {
	b := &Buffer{buf: make([]byte, 0)}
	return b.readAll(r, size)
}

func (b *Buffer) readAll(r io.Reader, size int) (buff []byte, err error) {
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		//buf太大会返回相应错误
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = b.ReadFrom(r, size) //关键就是这个家伙
	return b.buf, err
}

func (b *Buffer) ReadFrom(r io.Reader, size int) (n int64, err error) {
	if b.off >= len(b.buf) {
		b.Reset() //还没有写就想读，清空buf
	}
	for {
		if len(b.buf) > size {
			panic(bytes.ErrTooLarge)
		}
		if free := cap(b.buf) - len(b.buf); free < 512 {
			// free的大小是总容量 - 现在占有长度
			newBuf := b.buf
			if b.off+free < 512 {
				//分配更大空间，分配失败会报错
				newBuf = makeSlice(2*cap(b.buf) + 512)
			}
			//把读的内容b.buf[b.off:]拷贝到newbuf前面去
			copy(newBuf, b.buf[b.off:])
			//读写之间的差距就是应该读的buf
			b.buf = newBuf[:len(b.buf)-b.off]
			b.off = 0
		}
		//把io.Reader内容写到buf的free中去
		m, e := r.Read(b.buf[len(b.buf):cap(b.buf)])
		//重新调整buf的大小
		b.buf = b.buf[0 : len(b.buf)+m]
		n += int64(m)
		//读到尾部就返回
		if e == io.EOF {
			break
		}
		if e != nil {
			return n, e
		}
	}
	return n, nil // err is EOF, so return nil explicitly
}

// makeSlice allocates a slice of size n. If the allocation fails, it panics
// with ErrTooLarge.
func makeSlice(n int) []byte {
	// If the make fails, give a known error.
	defer func() {
		if recover() != nil {
			panic(bytes.ErrTooLarge)
		}
	}()
	return make([]byte, n)
}

func (b *Buffer) Read(p []byte) (n int, err error) {
	if b.off >= len(b.buf) {
		// Buffer is empty, reset to recover space.
		b.Reset()
		if len(p) == 0 {
			return
		}
		return 0, io.EOF
	}
	//就是这里咯，把b.buf[b.off:]的值写到p中去，记住copy(s1,s2)是s2写到s1中去，不要弄反咯
	//而且此Buffer其实是io.ReadCloser接口转化的类型
	n = copy(p, b.buf[b.off:])
	b.off += n
	return
}
func (b *Buffer) Reset() {
	b.buf = b.buf[:0]
	b.off = 0
}
