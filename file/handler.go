package file

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Readfile(path string) error {
	file, er := os.Open(path)
	if er != nil {
		log.Println(er)
		return er
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		det := strings.Split(scanner.Text(), ",")
		c := Csv{
			det[0],
			det[1],
			det[2],
			det[3],
			det[4],
			det[5],
		}
		fmt.Println("County name: ", c.CountryCode)
		fmt.Println("All : ", c)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
