package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
)

func main() {

	// 1.查看还有什么其他错误
	//err := ReadFileLeftErrInfo("/Users/blj/Downloads/skybai/learn/6os/error.log-20230422_501", "/Users/blj/Downloads/skybai/learn/6os/leftErr.log")
	//if err != nil {
	//	fmt.Printf("read file error:%v\n", err)
	//	return
	//}

	//err := ReadFileErrInfo("/Users/blj/Downloads/skybai/learn/6os/error.log-20230422_501")
	//if err != nil {
	//	fmt.Printf("read file error:%v\n", err)
	//	return
	//}

	//s, err := ReadFile3("/Users/blj/Downloads/skybai/learn/6os/test14.log")
	//if err != nil {
	//	fmt.Printf("read file error:%v\n", err)
	//	return
	//}

	//s1, err := ReadFileErrWid("/Users/blj/Downloads/skybai/learn/6os/out.txt")
	//if err != nil {
	//	fmt.Printf("read file error:%v\n", err)
	//	return
	//}
	//
	//fmt.Println("s1", s1)

	//
	err := ReadFile3("/Users/blj/Downloads/skybai/learn/6os/log.text")
	if err != nil {
		fmt.Printf("read file error:%v\n", err)
		return
	}

	//wid2, err := ReadFileErrWid2("/Users/blj/Downloads/skybai/learn/6os/out.txt")
	//if err != nil {
	//	fmt.Printf("read file error:%v\n", err)
	//	return
	//}

	//fmt.Println("wid2", wid2)
}

func ReadFile3(path string) (err error) {
	// 1.读文件
	fileHandle, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer fileHandle.Close()

	scanner := bufio.NewScanner(fileHandle)
	s := make([]string, 0)

	// 按行处理txt
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, "[STT") { // 截取每一行第三个引号到第四个引号之间内容的语句 统计重复的行出现的次数 打印出行内容和次数
			m := make(map[string]int)

			mx := sync.Mutex{}
			if strings.Contains(line, "/") && strings.Contains(line, "|") {
				if strings.Contains(line, "?") {
					start := strings.Index(line, "/")
					end := strings.Index(line, "?")
					api := line[start:end]
					mx.Lock()
					m[api]++
					mx.Unlock()
					if api == "/" {
						continue
					}
					s = append(s, api)
					continue
				}

				start := strings.Index(line, "/")
				end := strings.Index(line, "|")
				api := line[start:end]
				mx.Lock()
				m[api]++
				mx.Unlock()
				s = append(s, api)
			}
		}
	}

	// 统计每个元素出现的次数
	counts := make(map[string]int)
	for _, elem := range s {
		counts[elem]++
	}

	// 将出现次数排序
	type kv struct {
		Key   string
		Value int
	}
	var sortedCounts []kv
	for k, v := range counts {
		sortedCounts = append(sortedCounts, kv{k, v})
	}
	sort.Slice(sortedCounts, func(i, j int) bool {
		return sortedCounts[i].Value > sortedCounts[j].Value
	})
	fmt.Printf("统计出一共有%d个接口\n", len(sortedCounts))
	// 打印出现次数前10的元素和次数
	for i := 0; i < 260 && i < len(sortedCounts); i++ {
		fmt.Printf("%v 出现了 %v 次\n", sortedCounts[i].Key, sortedCounts[i].Value)
	}

	var deviceCounts []kv
	// find设备相关的前20接口
	for i := 0; i < 20 && i < len(sortedCounts); i++ {
		if strings.Contains(sortedCounts[i].Key, "device") || strings.Contains(sortedCounts[i].Key, "vehicle") {
			deviceKv := kv{sortedCounts[i].Key, sortedCounts[i].Value}
			deviceCounts = append(deviceCounts, deviceKv)
		}
	}

	sort.Slice(deviceCounts, func(i, j int) bool {
		return deviceCounts[i].Value > deviceCounts[j].Value
	})
	fmt.Println("下面是与设备相关的接口")
	fmt.Printf("统计出一共有%d个接口\n", len(deviceCounts))
	// 打印出现次数前10的元素和次数

	for i := 0; i < 260 && i < len(deviceCounts); i++ {
		fmt.Printf("%v 出现了 %v 次\n", deviceCounts[i].Key, deviceCounts[i].Value)
	}
	return nil
}

func ReadFileErrWid(path string) (y []string, err error) {
	fileHanle, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}

	defer fileHanle.Close()

	scanner := bufio.NewScanner(fileHanle)
	s := make([]string, 0)
	i := 0
	newLine := ""
	// 按行处理txt
	for scanner.Scan() {
		i++
		line := strings.TrimSpace(scanner.Text())

		//m := make(map[string]int)

		//mx := sync.Mutex{}
		// 读取空格之后的字符串
		index := strings.Index(line, " ")
		resultLine := line[index+1:]

		number := line[:index]

		newLine := newLine + number + " " + resultLine + "" + "wid" + "\n"
		fmt.Println("newLine:", newLine)

	}
	err = os.WriteFile("/Users/blj/Downloads/skybai/learn/6os/test50.log", []byte(newLine), 0666)
	if err != nil {
		fmt.Println("write file error:", err)
		return nil, err
	}
	fmt.Println("i:", i)
	return s, nil
}

func ReadFileErrWid2(path string) (y []string, err error) {
	fileHanle, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}

	defer fileHanle.Close()

	scanner := bufio.NewScanner(fileHanle)
	s := make([]string, 0)
	i := 0
	newLine := ""
	flag := os.O_CREATE | os.O_WRONLY | os.O_APPEND
	perm := os.FileMode(0666)
	file, err := os.OpenFile("/Users/blj/Downloads/skybai/learn/6os/test50.log", flag, perm)

	w := bufio.NewWriter(file)

	// 按行处理txt
	for scanner.Scan() {
		i++
		line := strings.TrimSpace(scanner.Text())

		// 读取空格之后的字符串
		index := strings.Index(line, " ")
		resultLine := line[index+1:]

		number := line[:index]

		newLine := newLine + number + " " + resultLine + " " + "wid" + "\n"
		fmt.Println("newLine:", newLine)

		_, err := w.WriteString(newLine)
		if err != nil {
			fmt.Println("write string error:", err)
			i--
			continue
		}
		err = w.Flush()
		if err != nil {
			fmt.Println("flush error:", err)
			i--
		}
	}

	fmt.Println("i:", i)
	return s, nil
}

// ReadFileLeftErrInfo 读取文件，将剩下的错误写入新文件中
func ReadFileLeftErrInfo(readFilepath, writeErrFilePath string) (err error) {
	// 1.读文件
	fileHandle, err := os.OpenFile(readFilepath, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer fileHandle.Close()
	scanner := bufio.NewScanner(fileHandle)

	// 2.创建leftErr文件
	file1, err := os.OpenFile(writeErrFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(0666))
	w := bufio.NewWriter(file1)

	// 按行处理txt
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.Contains(line, "no live upstreams while connecting to upstream") {
			continue
		} else if strings.Contains(line, "104: Connection reset by peer") {
			continue
		} else if strings.Contains(line, "110: Connection timed out") {
			continue
		} else if strings.Contains(line, "SSL_do_handshake") {
			continue
		} else if strings.Contains(line, "client sent invalid chunked body") {
			continue
		} else if strings.Contains(line, "worker_connections are not enough") {
			continue
		} else if strings.Contains(line, "client intended to send too large body") {
			continue
		} else if strings.Contains(line, "upstream prematurely closed connection while reading response header from upstream,") {
			continue
		} else if strings.Contains(line, "client intended to send too large chunked body") {
			continue
		}

		_, err := w.WriteString(line)
		if err != nil {
			return err
		}
		err = w.Flush()
		if err != nil {
			return err
		}
	}

	return nil
}

// ReadFileErrInfo 读取文件，将剩下的错误文件写入新文件中
func ReadFileErrInfo(errFilepath, path string) (err error) {
	// 1.创建文件
	fileHandle, err := os.OpenFile(errFilepath, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}

	defer fileHandle.Close()

	scanner := bufio.NewScanner(fileHandle)

	// 2.创建文件

	// err_no_live_upstreams_while_connecting_to_upstream
	file1, err := os.OpenFile(path+"/err_no_live_upstreams_while_connecting_to_upstream.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(0666))
	err1 := bufio.NewWriter(file1)

	// 104:_Connection_reset_by_peer
	file2, err := os.OpenFile(path+"/err_104:_Connection_reset_by_peer.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(0666))
	err2 := bufio.NewWriter(file2)

	// 110: Connection timed out
	file3, err := os.OpenFile(path+"/err_110:_Connection_timed_out.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(0666))
	err3 := bufio.NewWriter(file3)

	// SSL_do_handshake
	file4, err := os.OpenFile(path+"/err_SSL_do_handshake.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(0666))
	err4 := bufio.NewWriter(file4)

	// client sent invalid chunked body
	file5, err := os.OpenFile(path+"/err_client_sent_invalid_chunked_body.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(0666))
	err5 := bufio.NewWriter(file5)

	// worker_connections are not enough
	file6, err := os.OpenFile(path+"/err_worker_connections_are_not_enough.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(0666))
	err6 := bufio.NewWriter(file6)

	// client intended to send too large body
	file7, err := os.OpenFile(path+"/err_client_intended_to_send_too_large_body.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(0666))
	err7 := bufio.NewWriter(file7)

	// err_upstream_prematurely_closed_connection_while_reading_response_header_from_upstream
	file8, err := os.OpenFile(path+"/err_upstream_prematurely_closed_connection_while_reading_response_header_from_upstream.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(0666))
	err8 := bufio.NewWriter(file8)

	// err_client intended_to_send_too_large_chunked_body
	file9, err := os.OpenFile(path+"/err_client_intended_to_send_too_large_chunked_body.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(0666))
	err9 := bufio.NewWriter(file9)

	// 按行处理txt
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.Contains(line, "no live upstreams while connecting to upstream") {
			// 1.按照空格分割字符串
			arrLine := strings.Split(line, " ")
			// awk -F '[/: ]' '{if (($4 < 4 || ($4 == 4 && $5 >= 30)) || ($4 > 4)) {print}}' .log > .log
			newArr := strings.Split(arrLine[1], ":")
			if newArr[0] < "04" || (newArr[0] == "04" && newArr[1] >= "30") || newArr[0] > "04" {
				_, err := err1.WriteString(line + "\n")
				if err != nil {
					return err
				}
				err = err1.Flush()
				if err != nil {
					return err
				}
			}
		} else if strings.Contains(line, "104: Connection reset by peer") {
			// 1.按照空格分割字符串
			arrLine := strings.Split(line, " ")
			// awk -F '[/: ]' '{if (($4 < 4 || ($4 == 4 && $5 >= 30)) || ($4 > 4)) {print}}' .log > .log
			newArr := strings.Split(arrLine[1], ":")
			if newArr[0] < "04" || (newArr[0] == "04" && newArr[1] >= "30") || newArr[0] > "04" {
				_, err := err2.WriteString(line + "\n")
				if err != nil {
					return err
				}
				err = err2.Flush()
				if err != nil {
					return err
				}
			}
		} else if strings.Contains(line, "110: Connection timed out") {
			// 1.按照空格分割字符串
			arrLine := strings.Split(line, " ")
			// awk -F '[/: ]' '{if (($4 < 4 || ($4 == 4 && $5 >= 30)) || ($4 > 4)) {print}}' .log > .log
			newArr := strings.Split(arrLine[1], ":")
			if newArr[0] < "04" || (newArr[0] == "04" && newArr[1] >= "30") || newArr[0] > "04" {
				_, err := err3.WriteString(line + "\n")
				if err != nil {
					return err
				}
				err = err3.Flush()
				if err != nil {
					return err
				}
			}
		} else if strings.Contains(line, "SSL_do_handshake") {
			// 1.按照空格分割字符串
			arrLine := strings.Split(line, " ")
			// awk -F '[/: ]' '{if (($4 < 4 || ($4 == 4 && $5 >= 30)) || ($4 > 4)) {print}}' .log > .log
			newArr := strings.Split(arrLine[1], ":")
			if newArr[0] < "04" || (newArr[0] == "04" && newArr[1] >= "30") || newArr[0] > "04" {
				_, err := err4.WriteString(line + "\n")
				if err != nil {
					return err
				}
				err = err4.Flush()
				if err != nil {
					return err
				}
			}
		} else if strings.Contains(line, "client sent invalid chunked body") {
			// 1.按照空格分割字符串
			arrLine := strings.Split(line, " ")
			// awk -F '[/: ]' '{if (($4 < 4 || ($4 == 4 && $5 >= 30)) || ($4 > 4)) {print}}' .log > .log
			newArr := strings.Split(arrLine[1], ":")
			if newArr[0] < "04" || (newArr[0] == "04" && newArr[1] >= "30") || newArr[0] > "04" {
				_, err := err5.WriteString(line + "\n")
				if err != nil {
					return err
				}
				err = err5.Flush()
				if err != nil {
					return err
				}
			}
		} else if strings.Contains(line, "worker_connections are not enough") {
			// 1.按照空格分割字符串
			arrLine := strings.Split(line, " ")
			// awk -F '[/: ]' '{if (($4 < 4 || ($4 == 4 && $5 >= 30)) || ($4 > 4)) {print}}' .log > .log
			newArr := strings.Split(arrLine[1], ":")
			if newArr[0] < "04" || (newArr[0] == "04" && newArr[1] >= "30") || newArr[0] > "04" {
				_, err := err6.WriteString(line + "\n")
				if err != nil {
					return err
				}
				err = err6.Flush()
				if err != nil {
					return err
				}
			}
		} else if strings.Contains(line, "client intended to send too large body") {
			// 1.按照空格分割字符串
			arrLine := strings.Split(line, " ")
			// awk -F '[/: ]' '{if (($4 < 4 || ($4 == 4 && $5 >= 30)) || ($4 > 4)) {print}}' .log > .log
			newArr := strings.Split(arrLine[1], ":")
			if newArr[0] < "04" || (newArr[0] == "04" && newArr[1] >= "30") || newArr[0] > "04" {
				_, err := err7.WriteString(line + "\n")
				if err != nil {
					return err
				}
				err = err7.Flush()
				if err != nil {
					return err
				}
			}
		} else if strings.Contains(line, "upstream prematurely closed connection while reading response header from upstream,") {
			// 1.按照空格分割字符串
			arrLine := strings.Split(line, " ")
			// awk -F '[/: ]' '{if (($4 < 4 || ($4 == 4 && $5 >= 30)) || ($4 > 4)) {print}}' .log > .log
			newArr := strings.Split(arrLine[1], ":")
			if newArr[0] < "04" || (newArr[0] == "04" && newArr[1] >= "30") || newArr[0] > "04" {
				_, err := err8.WriteString(line + "\n")
				if err != nil {
					return err
				}
				err = err8.Flush()
				if err != nil {
					return err
				}
			}
		} else if strings.Contains(line, "client intended to send too large chunked body") {
			// 1.按照空格分割字符串
			arrLine := strings.Split(line, " ")
			// awk -F '[/: ]' '{if (($4 < 4 || ($4 == 4 && $5 >= 30)) || ($4 > 4)) {print}}' .log > .log
			newArr := strings.Split(arrLine[1], ":")
			if newArr[0] < "04" || (newArr[0] == "04" && newArr[1] >= "30") || newArr[0] > "04" {
				_, err := err9.WriteString(line + "\n")
				if err != nil {
					return err
				}
				err = err9.Flush()
				if err != nil {
					return err
				}
			}
		}

	}

	return nil
}
