/* -.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.

* File Name : reader.go

* Purpose :

* Creation Date : 08-22-2016

* Last Modified : Mon 22 Aug 2016 11:31:34 PM UTC

* Created By : Kiyor

_._._._._._._._._._._._._._._._._._._._._.*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/djimenez/iconv-go"
	"io"
	"log"
	"os"
	"strings"
)

var (
	from      = flag.String("f", "GB2312", "from encoding")
	to        = flag.String("t", "UTF-8", "to encoding")
	converter *iconv.Converter
)

func init() {
	flag.Parse()
	log.SetFlags(19)
	var err error
	converter, err = iconv.NewConverter(*from, *to)
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	if len(flag.Args()) == 0 {
		reader(os.Stdin)
	} else {
		for _, v := range flag.Args() {
			f, err := os.Open(v)
			if err != nil {
				log.Panic(err)
			}
			reader(f)
		}
	}
}

func reader(f *os.File) {
	reader := bufio.NewReader(f)
	for {
		l, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				return
			} else {
				log.Println(err.Error())
				os.Exit(1)
			}
		} else {
			processing(l)
		}
	}
}

func processing(line string) {
	line = strings.Trim(line, "\n")
	out, _ := converter.ConvertString(line)
	fmt.Println(out)
}
