package main

import (
	"bytes"
	"io"
	"io/ioutil"
)

func handleUpload(u io.Reader) (err error) {
	// capture all bytes from upload
	b, err := ioutil.ReadAll(u)
	if err != nil {
		return
	}

	// wrap the bytes in a ReadSeeker
	r := bytes.NewReader(b)
	// 把这些字节重新构造成一个可读的reader

	// process the meta data
	//err = processMetaData(r)
	if err != nil {
		return
	}

	// rewind the reader back to the start
	r.Seek(0, 0)

	// upload the data
	//err = uploadFile(r)
	if err != nil {
		return
	}

	return nil
}

// 1.将可读的reader读出来 就是字节了
// 2.将这些字节又构造bytes.NewReader
// 3.通过seek方法再读出来

// 需求：需要流里面的东西重复读
// 优点：不用一次性读完，可以分批读
// 缺点：需要seek
// 适用场景：读取的数据量比较小，如果数据量比较大，例如上传文件或者是图片或者是下载文件，那么就不适合使用这种方式了，
// 因为这种方式需要将所有的数据都读取到内存中，如果数据量比较大，那么就会导致内存溢出，所以这种方式适合读取的数据量比较小的场景。

// io读出来就不能再读了
