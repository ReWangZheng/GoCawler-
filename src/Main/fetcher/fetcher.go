package fetcher

import (
	"net/http"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding"
	"io/ioutil"
	"io"
	"bufio"
	"golang.org/x/net/html/charset"
	"fmt"
	"golang.org/x/text/encoding/unicode"
	"log"
)

func Getdocument(url string) ([]byte,error){
	resp, err := http.Get(url)
	if err != nil {
		return nil,err
	}

	defer resp.Body.Close() //关闭输出流

	//如果状态码不为2xx
	if resp.StatusCode != http.StatusOK {
		return nil,fmt.Errorf("StatusCode=%d",resp.StatusCode)
	}

	encodemethd := determineEncoding(resp.Body)//得到编码方式

	//得到固定编码格式的读入器
	reader := transform.NewReader(resp.Body, encodemethd.NewDecoder())
	return ioutil.ReadAll(reader)
}

//编码处理器
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("fetcher error %v",err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
