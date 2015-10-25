package main

import (
	"errors"
	"flag"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"text/template"
	"time"
)

var js string
var css string
var port string
var tmplFile string

type Item struct {
	Name string
	Url  string
	Type string
}

type Data struct {
	Style  string
	Script string
	Items  []Item
	Dirs   []string
}

var fs http.Handler

func main() {
	var baseDir string
	flag.StringVar(&baseDir, "base-dir", ".", "base direcotry")
	flag.StringVar(&css, "css", "", "your custom css file")
	flag.StringVar(&js, "js", "", "your custom js file")
	flag.StringVar(&port, "port", "8000", "listen port number")
	flag.StringVar(&tmplFile, "template", "", "template file")
	flag.Parse()

	if css != "" {
		css, _ = filepath.Abs(css)
	}
	if js != "" {
		js, _ = filepath.Abs(js)
	}
	if tmplFile != "" {
		tmplFile, _ = filepath.Abs(tmplFile)
	}

	os.Chdir(baseDir)
	fs = http.FileServer(http.Dir("./"))
	listener, ch := server(":" + port)
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT)
	go func() {
		log.Println(<-sig)
		listener.Close()
	}()
	log.Println(<-ch)
}

func server(addr string) (listener net.Listener, ch chan error) {
	ch = make(chan error)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", showItems)
		ch <- http.Serve(listener, mux)
	}()
	return
}

func responseBinData(w http.ResponseWriter, r *http.Request, file string) (err error) {
	if _, err = os.Stat(file); err != nil {
		http.NotFound(w, r)
		return err
	}

	ext := filepath.Ext(file)
	switch ext {
	case ".mp4", ".jpg", ".jpeg", ".png", ".gif", ".bmp":
		w.Header().Set("Cache-Control", "max-age=600")
		fs.ServeHTTP(w, r)
	case ".html", ".js", ".css":
		b, err := ioutil.ReadFile(file)
		if err != nil {
			Error(err)
			http.NotFound(w, r)
			return err
		}
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Content-Length", fmt.Sprint(len(b)))
		w.Header().Set("Accept-Ranges", "bytes")
		w.Header().Set("Cache-Control", "max-age=600")
		w.Write(b)
	default:
		http.NotFound(w, r)
	}
	return nil
}

func showItems(w http.ResponseWriter, r *http.Request) {
	info := "Time:" + fmt.Sprint(time.Now()) + "\t"
	info = info + "Url:" + r.Host + r.URL.String() + "\t"
	info = info + "Proto:" + r.Proto + "\t"
	info = info + "RemoteAddr:" + r.RemoteAddr + "\t"
	info = info + "Header:" + fmt.Sprint(r.Header) + "\t"
	if r.PostForm != nil {
		info = info + "PostForm:" + r.PostForm.Encode() + "\t"
	}
	if r.Trailer != nil {
		info = info + "Trailer:" + fmt.Sprint(r.Trailer) + "\t"
	}
	fmt.Println(info)
	name := r.URL.Path
	defer r.Body.Close()
	if name[len(name)-1] == '/' {
		w.Header().Set("Content-Type", "text/html")
		var files []string
		var dirs []string
		if name == "/" {
			files, dirs, _ = listFiles(".")
		} else {
			files, dirs, _ = listFiles(name[1 : len(name)-1])
		}
		if len(files) == 0 {
			http.NotFound(w, r)
			Error(errors.New("file not found in " + name))
			return
		}

		var style []byte
		var err error
		if css != "" {
			style, err = ioutil.ReadFile(css)
		} else {
			style, err = Asset("base.css")
		}
		if err != nil {
			Error(err)
			style = nil
		}

		var script []byte
		if js != "" {
			script, err = ioutil.ReadFile(js)
		} else {
			script, err = Asset("base.js")
		}
		if err != nil {
			Error(err)
			script = nil
		}

		s, e := getQueryValue(r, len(files))
		if s == e {
			http.NotFound(w, r)
			return
		}
		data := Data{
			Items:  classifyItem(files, s, e),
			Dirs:   dirs,
			Style:  fmt.Sprintf("%s", string(style)),
			Script: fmt.Sprintf("%s", string(script)),
		}
		var tmplValue []byte
		if tmplValue, err = ioutil.ReadFile(name[1:] + "index.html"); err == nil {
		} else if tmplFile != "" {
			tmplValue, err = ioutil.ReadFile(tmplFile)
		} else {
			tmplValue, err = Asset("base.html")
		}
		if err != nil {
			Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl := template.Must(template.New("").Parse(string(tmplValue)))
		if err := tmpl.ExecuteTemplate(w, "", data); err != nil {
			Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		err := responseBinData(w, r, name[1:])
		if err != nil {
			Error(err)
			return
		}
		return
	}
}

func classifyItem(files []string, s int, e int) (items []Item) {
	items = make([]Item, 0)
	for _, f := range files[s:e] {
		f = html.EscapeString(f)
		uri := strings.Replace(strings.Replace(url.QueryEscape(f), "+", "%20", -1), "%2F", "/", -1)
		switch filepath.Ext(f) {
		case ".jpg", ".jpeg", ".gif", ".png", ".bmp":
			items = append(items, Item{Name: f, Url: uri, Type: "image"})
		case ".mp4":
			items = append(items, Item{Name: f, Url: uri, Type: "video"})
		case ".html":
			items = append(items, Item{Name: f, Url: uri, Type: "html"})
		}
	}
	return items
}

func getQueryValue(r *http.Request, length int) (s int, e int) {
	p := r.URL.Query().Get("s")
	s = 0
	c := 50
	if p != "" {
		s, _ = strconv.Atoi(p)
		if s < 0 {
			s = 0
		} else if s >= length {
			return 0, 0
		}
	}
	p = r.URL.Query().Get("c")
	if p != "" {
		c, _ = strconv.Atoi(p)
		if c < 0 {
			c = 50
		}
	}
	e = s + c
	if length < s {
		s = length - c
		if s < 0 {
			s = 0
		}
	}
	if length < e {
		e = length
	}
	return s, e
}

func listFiles(path string) ([]string, []string, error) {
	var files []string
	var dirs []string
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		Error(err)
		return nil, nil, err
	}
	for i := 0; i < len(fileInfos); i++ {
		name := fileInfos[i].Name()
		if name[0] == '.' || name == "index.html" {
			continue
		}
		if fileInfos[i].IsDir() {
			children, d, err := listFiles(path + "/" + fileInfos[i].Name())
			if err != nil {
				continue
			}
			var tmp = make([]string, len(children))
			for i, v := range children {
				tmp[i] = name + "/" + v
			}
			files = append(files, tmp...)
			dirs = append(dirs, name+"/")
			tmp = make([]string, len(d))
			for i, v := range d {
				tmp[i] = name + "/" + v
			}
			dirs = append(dirs, tmp...)
		} else {
			switch filepath.Ext(fileInfos[i].Name()) {
			case ".jpg", ".jpeg", ".gif", ".png", ".bmp", ".mp4", ".html":
				files = append(files, name)
			default:
			}
		}
	}
	return files, dirs, nil
}

func Error(e error) {
	fmt.Fprintln(os.Stderr, "Time:"+fmt.Sprint(time.Now())+"\t"+fmt.Sprint(e))
}
