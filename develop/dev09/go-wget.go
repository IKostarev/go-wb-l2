package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"
)

type myWget struct {
	visited sync.Map
	proto   string
	src     string
	domain  string
	delim   string
	http.Client
	wg          sync.WaitGroup
	uris        chan string
	lastNewLink struct {
		time.Time
		sync.RWMutex
	}
}

var wget myWget

func checkProto(uri string) string {
	reg := regexp.MustCompile(`[\w]+://`)
	return reg.FindString(uri)
}

func osDelim() string {
	if runtime.GOOS == "windows" {
		return "\\"
	}
	return "/"
}

func correctDelims(in string) string {
	in = strings.ReplaceAll(in, "\\", wget.delim)
	in = strings.ReplaceAll(in, "/", wget.delim)
	reg := regexp.MustCompile(`[\\]+`)
	in = reg.ReplaceAllString(in, "\\")
	reg = regexp.MustCompile(`[\/]+`)
	in = reg.ReplaceAllString(in, "/")
	if strings.HasPrefix(in, wget.delim) {
		in = in[:len(in)-1]
	}
	return in
}

func fileWrite(path string, body []byte) error {
	path = wget.src + wget.delim + path
	path = correctDelims(path)
	folder := filepath.Dir(path)
	fmt.Println(path, ":", folder)
	os.MkdirAll(folder, os.ModePerm)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = f.Write(body)
	if err != nil {
		return err
	}
	f.Close()
	return nil
}

func trimWWW(uri string) string {
	return strings.ReplaceAll(uri, "www.", "")
}

func sameHost(r *http.Response) bool {
	if trimWWW(r.Request.URL.Host) == wget.domain {
		return true
	}
	return false
}

func correctLinks(links []string) []string {
	for i := range links {
		if strings.HasPrefix(links[i], "href") {
			links[i] = links[i][6 : len(links[i])-1]
			continue
		}
		if strings.HasPrefix(links[i], "src") {
			links[i] = links[i][5 : len(links[i])-1]
			continue
		}
		if strings.HasPrefix(links[i], "background=") {
			links[i] = links[i][12 : len(links[i])-1]
			continue
		}
	}
	return links
}

func removeProto(in string) string {
	if strings.HasPrefix(in, "https://") {
		return in[8:]
	}
	if strings.HasPrefix(in, "http://") {
		return in[7:]
	}
	return in
}

func findLinks(webpage string) []string {
	var res []string
	reg := regexp.MustCompile(`href="(.*?)+"`)
	hrefs := reg.FindAllString(webpage, -1)
	res = append(res, hrefs...)
	reg = regexp.MustCompile(`src="(.*?)+"`)
	srcs := reg.FindAllString(webpage, -1)
	res = append(res, srcs...)
	reg = regexp.MustCompile(`background="(.*?)+"`)
	backgrounds := reg.FindAllString(webpage, -1)
	res = append(res, backgrounds...)
	return correctLinks(res)
}

func prepLink(link string, currentPage string) string {
	if strings.Contains(link, "#") {
		return ""
	}
	if strings.HasPrefix(link, "//") {
		return wget.proto + link[2:]
	}
	if (strings.HasPrefix(link, "https://")) || (strings.HasPrefix(link, "http://")) {
		if strings.HasPrefix(trimWWW(removeProto(link)), wget.domain) {
			return link
		}
		return ""
	}
	if strings.HasPrefix(link, wget.domain) {
		return link
	}
	if strings.HasPrefix(link, "../") {
		tmp := path.Dir(currentPage)
		if tmp == "" {
			tmp = wget.proto + wget.domain
		}
		newLink := link[2:]
		return prepLink(newLink, currentPage)
	}
	if strings.HasPrefix(link, "./") {
		return wget.proto + wget.domain + link[2:]
	}
	if strings.HasPrefix(link, "/") {
		return wget.proto + path.Dir(removeProto(currentPage)) + link
	}
	pathd := path.Dir(removeProto(currentPage))
	if (pathd == "") || (pathd == ".") {
		pathd = wget.domain
	}
	return wget.proto + pathd + "/" + link
}

func crawl(uri string) {
	if _, ok := wget.visited.LoadOrStore(uri, struct{}{}); ok {
		return
	}
	fmt.Println("Посещаем страницу:", uri)
	resp, err := wget.Get(uri)
	if err != nil {
		log.Println(err)
		return
	}
	if !sameHost(resp) {
		return
	}
	wget.lastNewLink.Lock()
	wget.lastNewLink.Time = time.Now()
	wget.lastNewLink.Unlock()
	body, _ := ioutil.ReadAll(resp.Body)
	if !strings.Contains(resp.Header.Get("Content-Type"), "text") {
		fileWrite(resp.Request.URL.Path, body)
		return
	}
	save := resp.Request.URL.Path
	if save == "" {
		save = "index"
	}
	if !(strings.HasSuffix(save, ".html")) && !(strings.HasSuffix(save, ".htm")) {
		if strings.HasSuffix(save, "/") {
			save = save + "index.html"
		} else {
			save = save + ".html"
		}
	}
	links := findLinks(string(body))
	for _, l := range links {
		link := prepLink(l, uri)
		if link == "" {
			continue
		}
		if _, ok := wget.visited.Load(link); ok == true {
			continue
		}
		go func() {
			fmt.Println("Я на странице", uri, "отправляю следующую ссылку:", link)
			wget.uris <- prepLink(l, uri)
		}()
	}
	fileWrite(save, body)
}

func WGet(uri string) {
	wget.uris = make(chan string)
	wget.Client = http.Client{Transport: &http.Transport{}}
	wget.proto = checkProto(uri)
	wget.delim = osDelim()
	curDir, _ := os.Getwd()
	u, err := url.Parse(uri)
	if err != nil {
		log.Fatal(err)
	}
	wget.domain = trimWWW(u.Hostname())
	wget.src = curDir + wget.delim + wget.domain
	go func() {
		wget.uris <- uri
	}()
}

func crawler(uris chan string) {
	for uri := range uris {
		crawl(uri)
		time.Sleep(1000 * time.Millisecond)
	}
}

func workPool(uris chan string, n int) {
	for i := 0; i < n; i++ {
		fmt.Println("Запускаем воркера №", i)
		wget.wg.Add(1)
		go crawler(uris)
	}
	wget.wg.Wait()
}

func main() {
	uri := flag.String("u", "", "введите URL")
	flag.Parse()

	if checkProto(*uri) == "" {
		flag.Usage()
		os.Exit(1)
	}

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGHUP)
	signal.Notify(gracefulStop, syscall.SIGQUIT)
	signal.Notify(gracefulStop, syscall.SIGINT)

	go func() {
		<-gracefulStop
		fmt.Println("Graceful stop!")
		time.Sleep(1 * time.Second)
		os.Exit(1)
	}()

	go func() {
		for {
			time.Sleep(10 * time.Second)
			if time.Since(wget.lastNewLink.Time) >= 10*time.Second {
				gracefulStop <- syscall.SIGINT
			}
		}
	}()
	
	WGet(*uri)
	workPool(wget.uris, 10)

}
