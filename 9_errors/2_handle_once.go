package main

import (
	"encoding/json"
	"io"
	"log"
)

type Config struct {
}

// 错误只处理一次 处理错误意味着检查错误值并做出单一决定。

// WriteAllDemo writes the contents of buf to the supplied writer.
func WriteAllDemo(w io.Writer, buf []byte) {
	w.Write(buf)
}

func WriteAll(w io.Writer, buf []byte) error {
	_, err := w.Write(buf)
	if err != nil {
		log.Println("unable to write:", err) // annotated error goes to log file
		return err                           // unannotated error returned to caller
	}
	return nil
}

func WriteConfig(w io.Writer, conf *Config) error {
	buf, err := json.Marshal(conf)
	if err != nil {
		log.Printf("could not marshal configs: %v", err)
		return err // err后一定要返回
	}
	if err := WriteAll(w, buf); err != nil {
		log.Println("could not write configs: %v", err)
		return err
	}
	return nil
}

func main() {

	//err := WriteConfig(f, &conf)
	//fmt.Println(err) // io.EOF
}

// 这里会处理两次错误，一次是WriteAll，一次是WriteConfig 如果都打印到日志文件中就会有两条记录

// 在这个例子中，作者检查了错误，记录了它，但忘了返回。这就引起了一个微妙的错误。
//
//Go 语言中的错误处理规定，如果出现错误，你不能对其他返回值的内容做出任何假设。由于 JSON 解析失败，buf 的内容未知，可能它什么都没有，但更糟的是它可能包含解析的 JSON 片段部分。
//
//由于程序员在检查并记录错误后忘记返回，因此损坏的缓冲区将传递给 WriteAll，这可能会成功，因此配置文件将被错误地写入。但是，该函数会正常返回，并且发生问题的唯一日志行是有关 JSON 解析错误，而与写入配置失败有关。
