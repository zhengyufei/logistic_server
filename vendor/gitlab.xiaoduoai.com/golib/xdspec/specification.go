package xdspec

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

var (
	Version   string
	GitCommit string
	BuildTime string
	AdminPort string
)

var gServer *http.Server

func init() {
	http.Handle("/version", http.HandlerFunc(showVersion))
	if AdminPort != "" {
		StartAdminOn(AdminPort)
	}
}

func showVersion(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "version: %s\ncommit: %s\nbuild: %s\n", Version, GitCommit, BuildTime)
}

func startAdmin(addr string) {
	if gServer != nil {
		fmt.Println("close default server on ", gServer.Addr)
		gServer.Close()
		gServer = nil
	}
	fmt.Printf("start admin server on %s\n", addr)
	gServer = &http.Server{Addr: addr, Handler: nil}
	gServer.ListenAndServe()
}

// StartAdminOn addr
func StartAdminOn(addr string) {
	if addr == "" {
		fmt.Printf("bad admin addr(%s)\n", addr)
		return
	}
	go startAdmin(addr)
}
