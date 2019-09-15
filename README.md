# Wrapper for text/encoding

Allow to select correct encoder based on an encoder name

For latter use on Reading or Writing files, convert strings or byte array

Basic examples:

```go
import (
    "github.com/hanspr/ioencoder"
)

func main() {
    enc := ioencoder.New()
    file, err := os.Open(filename)
    utf8reader, err = enc.GetReader("iso-8859-1", file)
    // Code to read e Latin1 file and decode to UTF8

    // Decode a string
	txtutf8 := enc.DecodeString("BIG5", "這裡，可算是我的隨意空間吧")

    if wfile, err = os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644); err == nil {
        writerutf8, err = enc.GetWriter("SHIFT-JIS", wfile)
        // Code to wrte to file encoded in SHIFT-JIS
	}
}

```
