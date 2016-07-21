package mgenerator

import (

	"bytes"
	//	"fmt"
	"io"
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {

	buf := bytes.NewBuffer(nil)

	f, _ := os.Open("/tmp/book.txt")

	io.Copy(buf, f) // Error handling elided for brevity.
	f.Close()


	Generate(buf.Bytes())
	Generate(buf.Bytes())	
	

	//	fmt.Println(result)

}
