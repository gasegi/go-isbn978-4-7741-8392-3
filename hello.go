package main

import (
	"bytes"
	"fmt"
	"io"
	"mime"
	"net/http"
	"path"
	"strconv"
)

// 以下のコメントが go generate で使用される

// go:generate go-bindata -prefix data ./data

func staticHandler(w http.ResponseWriter, req *http.Request) {
	p := req.URL.Path
	if p == "" {
		p = "index.html"
	}

	// リクエストパスから静的コンテンツを得る
	bs, err := Asset(p)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// 拡張子から Content-Type を調べる
	if ctype := mime.TypeByExtension(path.Ext(p)); ctype != "" {
		w.Header().Set("Content-Type", ctype)

	}
	w.Header().Set("Content-Length", strconv.Itoa(len(bs)))
	io.Copy(w, bytes.NewBuffer(bs))
}

func main() {
	http.Handle("/", http.StripPrefix("/", http.HandlerFunc(staticHandler)))
	http.ListenAndServe(":3000", nil)
	fmt.Printf("Hello, world.\n")
}
