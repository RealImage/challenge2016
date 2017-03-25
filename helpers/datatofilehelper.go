package helpers

import (
	"os"
	"log"
)

func DataToFile(fileLocation, dataString string) {

	file, err := os.Create(fileLocation)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	file.WriteString(dataString)


}

