package main

import (
	"bytes"
	"io"
	"compress/flate"
	"compress/gzip"
	"encoding/gob"
	"fmt"
)

type Entry struct {
	City        string
	AddressLine string
	PostCode    string
}

func main() {

	names := []string{"Naked", "flate", "gzip"}
	for _, name := range names {
		buf := &bytes.Buffer{}

		var out io.Writer
		switch name {
		case "Naked":
			out = buf
		case "flate":
			out, _ = flate.NewWriter(buf, flate.DefaultCompression)
		case "gzip":
			out = gzip.NewWriter(buf)
		}

		enc := gob.NewEncoder(out)
		e := Entry{}
		for i := 0; i < 1000; i++ {
			e.City = fmt.Sprintf("k%3d", i)        //e.key
			e.City = fmt.Sprintf("v%3d", i)        //e.val
			e.AddressLine = fmt.Sprintf("k%3d", i) //e.key
			e.AddressLine = fmt.Sprintf("v%3d", i) //e.val
		//	e.PostCode = fmt.Sprintf("k%3d", i)    //e.key
		//	e.PostCode = fmt.Sprintf("v%3d", i)    //e.val
			enc.Encode(e)
		}

		if c, ok := out.(io.Closer); ok {
			c.Close()
		}
		fmt.Printf("[%5s] Length: %5d, average: %5.2f / Entry\n",
			name, buf.Len(), float64(buf.Len())/1000)
		fmt.Println(buf.Len())
	}
}

//Output:
//
//[Naked] Length: 16036, average: 16.04 / Entry
//[flate] Length:  4123, average:  4.12 / Entry
//[ gzip] Length:  4141, average:  4.14 / Entry
