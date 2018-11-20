package main

import (
	"os"
	"os/user"
	"github.com/zserge/lorca"
)

func main() {
	user, err := user.Current()
	if err != nil {
			panic(err)
	}
	dir := user.HomeDir + "/.lorca-app"
	os.MkdirAll(dir, os.ModePerm);
	ui, _ := lorca.New("https://www.google.com/intl/ja/gmail/about/", dir, 2000, 1000, "--start-maximized")
	defer ui.Close()
	<-ui.Done()
}