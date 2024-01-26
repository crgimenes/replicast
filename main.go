package main

import (
	"io"
	"log"
	"os"
	"replicast/assets"
	"replicast/config"
	"replicast/luaengine"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)

	err := config.Load()
	if err != nil {
		log.Fatalf("error loading config: %s\n", err)
	}

	/////////////////////////////////////////////////

	// read file init.lua from assets

	luaInit := config.CFG.Path + "/" + config.CFG.InitFile

	_, err = os.Stat(luaInit)
	if err != nil && !os.IsNotExist(err) {
		log.Printf("error reading %q : %s\r\n", luaInit, err)
		return
	}

	if os.IsNotExist(err) && config.CFG.InitFile == "init.lua" {
		f, err := os.Create(luaInit)
		if err != nil {
			return
		}
		finit, err := assets.FS.Open("init.lua")
		if err != nil {
			log.Printf("error reading init.lua from assets: %s\r\n", err)
			return
		}
		_, err = io.Copy(f, finit)
		if err != nil {
			log.Printf("error writing init.lua: %s\r\n", err)
			return
		}
		f.Close()
	}

	err = luaengine.Startup(luaInit)
	if err != nil {
		return
	}

	// TODO: start server

}
