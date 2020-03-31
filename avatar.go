package main

import (
	"./AvatarConfig"
	"./AvatarEmail"
	"./AvatarScreenOperation"
	"fmt"
	"os"

	//"github.com/go-vgo/robotgo"
	"gopkg.in/ini.v1"
)

/*
Version info
*/

var VERSION = 0.1

func main() {
	//first, let us handle the config file and see if version matches
	cfg, err := ini.Load("avatar.ini")

	//as per the document of gopkg.in/ini.v1, set BlockMode to false to improve efficiency
	cfg.BlockMode = false

	if err != nil {
		fmt.Printf("Error in ini.Load: %v", err)
		os.Exit(1)
	}

	//miss my c++ days
	mGeneralConfig := new(AvatarConfig.General_cfg)
	mEmailConfig := new(AvatarConfig.Email_cfg)
	mScreenOp := new(AvatarConfig.Screen_cfg)

	if cfg.Section("General").MapTo(mGeneralConfig) != nil {
		fmt.Printf("Error in Reading General Section: %v", err)
		os.Exit(1)
	}

	if cfg.Section("Email").MapTo(mEmailConfig) != nil {
		fmt.Printf("Error in Reading Email Section: %v", err)
		os.Exit(1)
	}

	if cfg.Section("ScreenOperation").MapTo(mScreenOp) != nil {
		fmt.Printf("Error in Reading Screen Op Section: %v", err)
		os.Exit(1)
	}

	if mGeneralConfig.Version == VERSION {
		//Just be good to check version, but nothing serious happens really
		go AvatarEmail.AvatarEmail(mEmailConfig)
		go AvatarScreenOperation.AvatarScreenOperation(mScreenOp)
	}

	select {}

}
