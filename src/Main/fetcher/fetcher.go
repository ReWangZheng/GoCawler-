package fetcher

import (
	"net/http"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding"
	"io/ioutil"
	"bufio"
	"golang.org/x/net/html/charset"
	"fmt"
	"golang.org/x/text/encoding/unicode"
	"log"
)

func Getdocument(url string) ([]byte,error){
	request, e := http.NewRequest(http.MethodGet, url, nil)
	if e != nil {
		return nil,e
	}

	request.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36");

	client:=http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close() //关闭输出流
	//如果状态码不为2xx
	if resp.StatusCode != http.StatusOK {
		return nil,fmt.Errorf("StatusCode=%d",resp.StatusCode)
	}
	bufioreader := bufio.NewReader(resp.Body)
	encodemethd := determineEncoding(bufioreader)//得到编码方式
	//得到固定编码格式的读入器
	reader := transform.NewReader(resp.Body, encodemethd.NewDecoder())
	return ioutil.ReadAll(reader)
}

//编码处理器
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err :=r.Peek(1024)
	if err != nil {
		log.Printf("fetcher error %v",err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
