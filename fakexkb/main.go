package main

import (
	"os"
	"bytes"
	"errors"
	"regexp"
	"fmt"
	"io/fs"
	"os/exec"
	"strings"
)

var re *regexp.Regexp = regexp.MustCompile(`(KEYMAP|KEYTABLE)=("|)(.*?)("|\n)`)

var layouts map[string]string = make(map[string]string)
var keymap bytes.Buffer

func main() {
	loadLayouts()

	// get the keymap name
	keymapName := getKeymapName()
	// get the folder it's in
	keymapFolder := getKeyboardFolder(keymapName)
	keymapFile := "/usr/share/keymaps/i386/"+keymapFolder+"/"+keymapName+".kmap.gz"
	fmt.Println("Loading "+keymapFile)
	// Make sure that actually exists.
	if _, err := os.ReadFile(keymapFile); err != nil {
		// If it doesn't, we're not on DSL, so default to the qwerty layout on modern 64-bit (arch) Linux.
		keymapFile = "/usr/share/kbd/keymaps/i386/qwerty/us.map.gz"
		fmt.Println("Nevermind, loading "+keymapFile)
		// if THAT doesn't work, panic.
		if _, err := os.ReadFile(keymapFile); err != nil {
			panic(err)
		}
	}
	// Write the appropriate keymap into /tmp/keymap.h
	writeKeymapFile(keymapFile)
}

func getKeymapName() (string) {
	// Open the keymap file.
	file, err := os.ReadFile("/etc/sysconfig/keyboard")
	if(err != nil) {
		// If it "doesn't exist", don't even print an error. It's expected.
		if !errors.Is(err, fs.ErrNotExist) {
			fmt.Println(err)
		}
		return "us"
	}
	// Next get the value for the keymap or keytable
	matches := re.FindAllSubmatch(file, 1)
	// Unless the user has a weird and broken system, there's almost
	// always gonna be 1 match, and it's almost always gonna be the 4th submatch.
	// That said, if there is an error...just default to us.
	if(len(matches) < 0 || len(matches[0]) < 4) {
		return "us"
	} else {
		keymap := matches[0][3]
		return string(keymap)
	}
}

func getKeyboardFolder(keymap string) (string) {
	// We have a function for this so that we can save on memory; instead of
	// filling the rest of the folder map with qwerty we can just make it a fallback
	// using this function.
	if elm, ok := layouts[keymap]; ok {
		return elm
	} else {
		return "qwerty"
	}
}

func writeKeymapFile(filename string) {
	filename = strings.Replace(filename,";","",-1)
	// loadkeys -m <file.map.gz>
	out1, err := exec.Command("loadkeys","-m",filename).Output()
	if(err != nil) {panic(err)}
	// Hang until the command successfully executes.
	for(string(out1) == "") {} 
	// write that to a .c file
	file, err := os.OpenFile("/tmp/keymap.h", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if(err != nil) {panic(err)}
	err = file.Truncate(0)
	if(err != nil) {panic(err)}
	_, err = file.Write(out1)
	if(err != nil) {panic(err)}
}

func loadLayouts() {
// fill the layout map according to what's on the DSL filesystem.
	// todo: please god there has to be something i'm missing this is such an awful solution 
	layouts["azerty"] = "azerty"
	layouts["be-latin1"] = "azerty"
	layouts["be2-latin1"] = "azerty"
	layouts["fr-latin0"] = "azerty"
	layouts["fr-latin1"] = "azerty"
	layouts["fr-latin9"] = "azerty"
	layouts["fr-pc"] = "azerty"
	layouts["fr-x11"] = "azerty"
	layouts["fr"] = "azerty"

	layouts["ANSI-dvorak"] = "dvorak"
	layouts["dvorak-classic"] = "dvorak"
	layouts["dvorak"] = "dvorak"
	layouts["pc-dvorak-latin1"] = "dvorak"

	layouts["croat"] = "qwertz"
	layouts["cz-us-qwertz"] = "qwertz"
	layouts["de-lat1-nd"] = "qwertz"
	layouts["de-latin1-nodeadkeys"] = "qwertz"
	layouts["de-latin1"] = "qwertz"
	layouts["de"] = "qwertz"
	layouts["fr_CH-latin1"] = "qwertz"
	layouts["fr_CH"] = "qwertz"
	layouts["hu"] = "qwertz"
	layouts["mac-usb-de-latin1-nodeadkeys"] = "qwertz"
	layouts["mac-usb-de-latin1"] = "qwertz"
	layouts["mac-usb-de_CH"] = "qwertz"
	layouts["mac-usb-fr_CH-latin1"] = "qwertz"
	layouts["sg-latin1-lk450"] = "qwertz"
	layouts["sg-latin1"] = "qwertz"
	layouts["sk-prog-qwertz"] = "qwertz"
	layouts["sk-qwertz"] = "qwertz"
	layouts["slovene"] = "qwertz"
	layouts["sr"] = "qwertz"

	// Every other country uses qwerty, and actually the list is too long to even fit in DSL's terminal.
	// so that'll just be the fallback
}