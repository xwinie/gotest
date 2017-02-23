package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

var (
	Client = http.Client{
		CheckRedirect: func(_ *http.Request, _ []*http.Request) error { return fmt.Errorf("Redirect Disabled") },
	}
	lock = sync.Mutex{}
)

func ShowRequest(r *http.Request) {
	lock.Lock()
	fmt.Println("---> ", r.Method, " ", r.URL.String())
	for i, j := range r.Header {
		fmt.Print("    ", i, " : ")
		for _, k := range j {
			fmt.Print(k, " ")
		}
		fmt.Print("\n")
	}
	lock.Unlock()
}

func ShowResponse(r *http.Response) {
	lock.Lock()
	fmt.Println("<--- ", r.StatusCode, " ", r.Request.URL.String())
	for i, j := range r.Header {
		fmt.Print("    ", i, " : ")
		for _, k := range j {
			fmt.Print(k, " ")
		}
		fmt.Print("\n")
	}
	lock.Unlock()
}

func handler(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	r.RequestURI = ""

	ShowRequest(r)

	m, err := http.DefaultClient.Do(r) // 页面可以看到缓冲，但是不能播放
	// Client.Do(r) www.bilibili.tv视频页面会不停请求视频数据（看不到response）
	if err != nil {
		http.NotFound(w, r)
		return
	}
	defer m.Body.Close()

	ShowResponse(m)

	for i, j := range m.Header {
		for _, k := range j {
			w.Header().Add(i, k)
		}
	}
	w.WriteHeader(m.StatusCode)

	_, ok := m.Header["Content-Length"]
	if !ok {
		data, err := ioutil.ReadAll(m.Body)
		if err != nil && err != io.EOF {
			return
		}
		w.Write(data)
	} else {
		data := make([]byte, 8192)
		for i, j := int64(0), 0; i < m.ContentLength; i += int64(j) {
			j, err = m.Body.Read(data)
			if err == nil {
				w.Write(data)
			}
		}
		if err == io.EOF {
			w.Write(data)
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Start serving on port 12345")
	http.ListenAndServe(":12345", nil)
	os.Exit(0)
}