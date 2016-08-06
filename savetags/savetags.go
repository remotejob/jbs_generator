package savetags

import (
	"encoding/csv"
	"log"
	"os"
)

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func Saveinfile(filestr string, tags []string) {

	file, err :=os.OpenFile(filestr,os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)

	for _, value := range tags {
		arr_value :=[]string{value}
		err := writer.Write(arr_value)
		checkError("Cannot write to file", err)
	}

	defer writer.Flush()

}
