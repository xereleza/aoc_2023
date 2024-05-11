package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("input") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buf := make([]byte, 1)
	//scanner := bufio.NewScanner(file)
	var temp []byte
	var final int
	for {
		_, err := file.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		if err == io.EOF {
			break
		}

		if !bytes.Equal([]byte("\n"), buf) {
			temp = append(temp, buf[0])
		} else {
			res, _ := proccess_line(temp)
			/* if err != nil {
				log.Fatal(err)
			} */
			fmt.Println("adding " + strconv.Itoa(res) + " to result...")
			final += res
			temp = make([]byte, 0)
		}
	}
	res, _ := proccess_line(temp)
	/* if err != nil {
		log.Fatal(err)
	} */
	fmt.Println("adding " + strconv.Itoa(res) + " to result...")
	final += res
	fmt.Println("final result: " + strconv.Itoa(final))
}

func check_word(chars []byte) (string, bool) {
	INTS := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	//fmt.Printf("checking word: " + string(chars) + "\n")
	valor, ok := INTS[string(chars)]
	return valor, ok

}

func proccess_line(line []byte) (int, error) {
	fmt.Print("proccess line " + string(line) + "...\n")
	res := make([]string, 2)
	l := 0
	for l < len(line) {
		if line[l] >= 48 && line[l] <= 57 { // range dos int
			res[0] = string(line[l])
			break
		}
		if line[l] >= 97 && line[l] <= 122 { // range dos char
			r := l
			for r < len(line)-1 {
				if line[r+1] >= 97 && line[r+1] <= 122 {
					//fmt.Print(string(line[r+1]) + "\n")
					valor, ok := check_word(line[l : r+2])
					if ok {
						res[0] = valor
						break
					}
				}
				r++
			}
			if res[0] != "" {
				break
			}
		}
		l++
	}
	r := len(line) - 1
	for r >= l {
		if line[r] >= 48 && line[r] <= 57 { // range dos int
			res[1] = string(line[r])
			break
		}
		if line[r] >= 97 && line[r] <= 122 { // range dos char
			l2 := r
			for l2 > l {
				if line[l2-1] >= 97 && line[l2-1] <= 122 {
					//fmt.Print(string(line[r+1]) + "\n")
					valor, ok := check_word(line[l2-1 : r+1])
					if ok {
						res[1] = valor
						break
					}
				}
				l2--
			}
			if res[1] != "" {
				break
			}
		}
		r--
	}
	fmt.Println(res)
	return strconv.Atoi(res[0] + res[1])
}
