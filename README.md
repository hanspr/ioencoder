# Wrapper for text/encoding

Allow to select correct encoder based on an encoder name

For latter use on Reading or Writing files, convert strings or byte array

Basic examples:

```go
import (
    "github.com/hanspr/ioencoder"
)

func main() {
    // Create object
    enc := ioencoder.New()

    // Get a io.Reader and wrap with encoder
    file, err := os.Open(filename)
    utf8reader, err = enc.GetReader("iso-8859-1", file)
    // Code to read e Latin1 file and decode to UTF8

    // Encode - Decode a string
    txtbig5 := enc.EncodeString("BIG5", "這裡，可算是我的隨意空間吧")
    txtutf8 := enc.DecodeString("BIG5", txtbig5)

    // Get io.Wrtier and wrap with encoder
    if wfile, err = os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644); err == nil {
        writerutf8, err = enc.GetWriter("SHIFT-JIS", wfile)
        // Code to write to file encoded in SHIFT-JIS
    }
}
```

```go
import (
    "github.com/hanspr/ioencoder"
)

func main() {
    // Creat object
    enc := ioencoder.New()
    enc.SetEncoding("ISO88592")

    // Use the selected encoder properties directly
    txtlatin1 := enc.encoder.String("latin1 string: áéíóúÑ")
    txtutf8 := enc.decoder.String(txtlatin1)
}
```
