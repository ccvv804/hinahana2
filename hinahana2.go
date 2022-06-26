package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"strings"
)

func hinahana(r string) {
	nulldir := false
	dat, err := ioutil.ReadFile(r)
	if err != nil {
		fmt.Println(err.Error())
		//fmt.Println("File not found")
		return
	}
	nametest := dat[0:4]
	hanaheader := []byte{0x48, 0x41, 0x4E, 0x41}
	hananumber := reflect.DeepEqual(nametest, hanaheader)
	if !hananumber {
		fmt.Println("File magic number NG")
		return
	}
	rr := path.Base(r)
	rrr := strings.TrimSuffix(rr, ".bin")
	rrrr := strings.TrimSuffix(rrr, ".BIN")
	
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = os.Mkdir(dir+"/"+rrrr, 0755)
	if err != nil {
		if os.IsExist(err) {
			fmt.Println("warning! " + rrrr + " folder already exists")
			return
		} else {
			fmt.Println(err.Error())
			return
		}
	}
	err = os.Chdir(dir + "/" + rrrr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	filenumbertest := dat[8:12]
	//filesizetest := dat[12:16]
	filenumber := binary.LittleEndian.Uint32(filenumbertest)
	//fmt.Println(filenumber)
	filenumbercounter := filenumber
	for filenumbercounter > 0 {
		fakeseek := 16 + 48*(filenumber-filenumbercounter)
		exfileinfo := dat[fakeseek : fakeseek+48]
		exfilename := string(bytes.Trim(exfileinfo[0:14], "\x00"))
		exfiledisk := string(bytes.Trim(exfileinfo[14:16], "\x00"))
		if exfiledisk != "c:" {
			fmt.Println(exfiledisk)
			fmt.Println("Sorry I didn't make this feature.")
			fmt.Println("Please contact the developer.")
			return
		}
		//fmt.Println(exfilename)
		exfileaddress := string(bytes.Trim(exfileinfo[16:36], "\x00"))
		exfilesize := binary.LittleEndian.Uint32(exfileinfo[36:40])
		exfilebinaddress := binary.LittleEndian.Uint32(exfileinfo[40:44])
		//fmt.Println(exfileaddress)
		//fmt.Println(exfilesize)
		//fmt.Println(exfilebinaddress)
		blankstr := ""

		if exfileaddress != blankstr {
			err = os.MkdirAll(dir+"/"+rrrr+exfileaddress, 0755)
			if err != nil {
				if os.IsExist(err) {
				} else {
					fmt.Println(err.Error())
					return
				}
			}
		} else {
			nulldir = true
		}
		if nulldir {
			nulldir = false
		} else {
			err = os.Chdir(dir + "/" + rrrr + exfileaddress)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
		exfiledata := dat[exfilebinaddress : exfilebinaddress+exfilesize]
		headexfiledata := exfiledata[0:4]
		if strings.HasSuffix(exfilename, ".bin") && reflect.DeepEqual(headexfiledata, hanaheader) || strings.HasSuffix(exfilename, ".BIN") && reflect.DeepEqual(headexfiledata, hanaheader) {
			exr := path.Base(exfilename)
			exrr := strings.TrimSuffix(exr, ".bin")
			exrrr := strings.TrimSuffix(exrr, ".BIN")
			err = os.Chdir(dir + "/" + rrrr)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			hinahanamini(exfiledata, exrrr)
			err = os.Chdir(dir + "/" + rrrr)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		} else if strings.HasSuffix(exfilename, ".KBN") || strings.HasSuffix(exfilename, ".kbn") {
			xm := string(bytes.Trim(exfiledata[16:32], "\x00"))
			xdata := exfiledata[144:]
			err := ioutil.WriteFile(xm, xdata, 0755)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		} else {
			err := ioutil.WriteFile(exfilename, exfiledata, 0755)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
		filenumbercounter = filenumbercounter - 1
	}
}

func hinahanamini(dat []byte, r string) {
	nulldir := false
	nametest := dat[0:4]
	hanaheader := []byte{0x48, 0x41, 0x4E, 0x41}
	hananumber := reflect.DeepEqual(nametest, hanaheader)
	if !hananumber {
		fmt.Println("File magic number NG")
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	filenumbertest := dat[8:12]
	filenumber := binary.LittleEndian.Uint32(filenumbertest)
	filenumbercounter := filenumber
	for filenumbercounter > 0 {
		fakeseek := 16 + 48*(filenumber-filenumbercounter)
		exfileinfo := dat[fakeseek : fakeseek+48]
		exfilename := string(bytes.Trim(exfileinfo[0:14], "\x00"))
		exfiledisk := string(bytes.Trim(exfileinfo[14:16], "\x00"))
		if exfiledisk != "c:" {
			fmt.Println(exfiledisk)
			fmt.Println("Sorry I didn't make this feature.")
			fmt.Println("Please contact the developer.")
			return
		}
		exfileaddress := string(bytes.Trim(exfileinfo[16:36], "\x00"))
		exfilesize := binary.LittleEndian.Uint32(exfileinfo[36:40])
		exfilebinaddress := binary.LittleEndian.Uint32(exfileinfo[40:44])
		blankstr := ""
		if exfileaddress != blankstr {
			err = os.MkdirAll(dir+exfileaddress, 0755)
			if err != nil {
				if os.IsExist(err) {
				} else {
					fmt.Println(err.Error())
					return
				}
			}
		} else {
			nulldir = true
		}
		if nulldir {
			nulldir = false
		} else {
			err = os.Chdir(dir + exfileaddress)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
		exfiledata := dat[exfilebinaddress : exfilebinaddress+exfilesize]
		if strings.HasSuffix(exfilename, ".KBN") || strings.HasSuffix(exfilename, ".kbn") {
			xm := string(bytes.Trim(exfiledata[16:32], "\x00"))
			xdata := exfiledata[144:]
			err := ioutil.WriteFile(xm, xdata, 0755)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		} else {
			err := ioutil.WriteFile(exfilename, exfiledata, 0755)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
		filenumbercounter = filenumbercounter - 1
	}
}

func main() {
	file := flag.String("file", "F10.bin", "Input file")
	flag.Parse()
	hinahana(*file)
}
