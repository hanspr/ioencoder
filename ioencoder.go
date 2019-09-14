package ioencoder

import (
	"errors"
	"io"
	"strconv"
	"strings"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
)

type IoEncoder struct {
	encodername string
	decoder     *encoding.Decoder
	encoder     *encoding.Encoder
}

func New() *IoEncoder {
	e := new(IoEncoder)
	return e
}

func (e *IoEncoder) GetAvailableEncodingsAsString(separator string) string {
	str := ""
	encs := e.GetAvailableEncodings()
	for _, e := range encs {
		if str == "" {
			str = e
		} else {
			str = str + separator + e
		}
	}
	return str
}

func (e *IoEncoder) GetAvailableEncodings() []string {
	var encs []string

	encs = append(encs, "BIG5")
	encs = append(encs, "EUCJP")
	encs = append(encs, "EUCKR")
	encs = append(encs, "GB18030")
	encs = append(encs, "GBK")
	encs = append(encs, "IBM855")
	encs = append(encs, "IBM866")
	encs = append(encs, "ISO2022JP")
	for i := 1; i < 17; i++ {
		if i == 11 || i == 12 {
			continue
		}
		encs = append(encs, "ISO8859"+strconv.Itoa(i))
	}
	encs = append(encs, "KOI8R")
	encs = append(encs, "KOI8U")
	encs = append(encs, "MACINTOSH")
	encs = append(encs, "SHIFTJIS")
	for i := 1250; i < 1259; i++ {
		encs = append(encs, "WINDOWS"+strconv.Itoa(i))
	}
	return encs
}

func (e *IoEncoder) DecodeBytes(encodername string, b []byte) []byte {
	err := e.setEncoding(encodername)
	if err == nil {
		newb, err := e.decoder.Bytes(b)
		if err == nil {
			return newb
		}
	}
	return b
}

func (e *IoEncoder) DecodeString(encodername, text string) string {
	err := e.setEncoding(encodername)
	if err == nil {
		newstr, err := e.decoder.String(text)
		if err == nil {
			return newstr
		}
	}
	return text
}

func (e *IoEncoder) GetReader(encodername string, r io.Reader) (io.Reader, error) {
	err := e.setEncoding(encodername)
	if err == nil {
		return e.decoder.Reader(r), nil
	}
	return r, err
}

func (e *IoEncoder) GetWriter(encodername string, w io.Writer) (io.Writer, error) {
	err := e.setEncoding(encodername)
	if err == nil {
		return e.encoder.Writer(w), nil
	}
	return w, err
}

func (e *IoEncoder) setEncoding(encodername string) error {
	name := strings.ReplaceAll(strings.ToUpper(encodername), "-", "")
	name = strings.ReplaceAll(name, "_", "")
	e.encodername = encodername

	if name == "ISO88591" {
		e.decoder = charmap.ISO8859_1.NewDecoder()
		e.encoder = charmap.ISO8859_1.NewEncoder()
	} else if name == "ISO88592" {
		e.decoder = charmap.ISO8859_2.NewDecoder()
		e.encoder = charmap.ISO8859_2.NewEncoder()
	} else if name == "ISO88593" {
		e.decoder = charmap.ISO8859_3.NewDecoder()
		e.encoder = charmap.ISO8859_3.NewEncoder()
	} else if name == "ISO88594" {
		e.decoder = charmap.ISO8859_4.NewDecoder()
		e.encoder = charmap.ISO8859_4.NewEncoder()
	} else if name == "ISO88595" {
		e.decoder = charmap.ISO8859_5.NewDecoder()
		e.encoder = charmap.ISO8859_5.NewEncoder()
	} else if name == "ISO88596" {
		e.decoder = charmap.ISO8859_6.NewDecoder()
		e.encoder = charmap.ISO8859_6.NewEncoder()
	} else if name == "ISO88597" {
		e.decoder = charmap.ISO8859_7.NewDecoder()
		e.encoder = charmap.ISO8859_7.NewEncoder()
	} else if name == "ISO88598" {
		e.decoder = charmap.ISO8859_8.NewDecoder()
		e.encoder = charmap.ISO8859_8.NewEncoder()
	} else if name == "ISO88599" {
		e.decoder = charmap.ISO8859_9.NewDecoder()
		e.encoder = charmap.ISO8859_9.NewEncoder()
	} else if name == "ISO885910" {
		e.decoder = charmap.ISO8859_10.NewDecoder()
		e.encoder = charmap.ISO8859_10.NewEncoder()
	} else if name == "ISO885913" {
		e.decoder = charmap.ISO8859_13.NewDecoder()
		e.encoder = charmap.ISO8859_13.NewEncoder()
	} else if name == "ISO885914" {
		e.decoder = charmap.ISO8859_14.NewDecoder()
		e.encoder = charmap.ISO8859_14.NewEncoder()
	} else if name == "ISO885915" {
		e.decoder = charmap.ISO8859_15.NewDecoder()
		e.encoder = charmap.ISO8859_15.NewEncoder()
	} else if name == "ISO885916" {
		e.decoder = charmap.ISO8859_16.NewDecoder()
		e.encoder = charmap.ISO8859_16.NewEncoder()
	} else if name == "WINDOWS1250" {
		e.decoder = charmap.Windows1250.NewDecoder()
		e.encoder = charmap.Windows1250.NewEncoder()
	} else if name == "WINDOWS1251" {
		e.decoder = charmap.Windows1251.NewDecoder()
		e.encoder = charmap.Windows1251.NewEncoder()
	} else if name == "WINDOWS1252" {
		e.decoder = charmap.Windows1252.NewDecoder()
		e.encoder = charmap.Windows1252.NewEncoder()
	} else if name == "WINDOWS1253" {
		e.decoder = charmap.Windows1253.NewDecoder()
		e.encoder = charmap.Windows1253.NewEncoder()
	} else if name == "WINDOWS1254" {
		e.decoder = charmap.Windows1254.NewDecoder()
		e.encoder = charmap.Windows1254.NewEncoder()
	} else if name == "WINDOWS1255" {
		e.decoder = charmap.Windows1255.NewDecoder()
		e.encoder = charmap.Windows1255.NewEncoder()
	} else if name == "WINDOWS1256" {
		e.decoder = charmap.Windows1256.NewDecoder()
		e.encoder = charmap.Windows1256.NewEncoder()
	} else if name == "WINDOWS1257" {
		e.decoder = charmap.Windows1257.NewDecoder()
		e.encoder = charmap.Windows1257.NewEncoder()
	} else if name == "WINDOWS1258" {
		e.decoder = charmap.Windows1258.NewDecoder()
		e.encoder = charmap.Windows1258.NewEncoder()
	} else if name == "KOI8R" {
		e.decoder = charmap.KOI8R.NewDecoder()
		e.encoder = charmap.KOI8R.NewEncoder()
	} else if name == "KOI8U" {
		e.decoder = charmap.KOI8U.NewDecoder()
		e.encoder = charmap.KOI8U.NewEncoder()
	} else if name == "IBM855" {
		e.decoder = charmap.CodePage855.NewDecoder()
		e.encoder = charmap.CodePage855.NewEncoder()
	} else if name == "IBM866" {
		e.decoder = charmap.CodePage866.NewDecoder()
		e.encoder = charmap.CodePage866.NewEncoder()
	} else if name == "SHIFTJIS" {
		e.decoder = japanese.ShiftJIS.NewDecoder()
		e.encoder = japanese.ShiftJIS.NewEncoder()
	} else if name == "ISO2022JP" {
		e.decoder = japanese.ISO2022JP.NewDecoder()
		e.encoder = japanese.ISO2022JP.NewEncoder()
	} else if name == "EUCJP" {
		e.decoder = japanese.EUCJP.NewDecoder()
		e.encoder = japanese.EUCJP.NewEncoder()
	} else if name == "BIG5" {
		e.decoder = traditionalchinese.Big5.NewDecoder()
		e.encoder = traditionalchinese.Big5.NewEncoder()
	} else if name == "GB18030" {
		e.decoder = simplifiedchinese.GB18030.NewDecoder()
		e.encoder = simplifiedchinese.GB18030.NewEncoder()
	} else if name == "GBK" {
		e.decoder = simplifiedchinese.GBK.NewDecoder()
		e.encoder = simplifiedchinese.GBK.NewEncoder()
	} else if name == "EUCKR" {
		e.decoder = korean.EUCKR.NewDecoder()
		e.encoder = korean.EUCKR.NewEncoder()
	} else if name == "MACINTOSH" {
		e.decoder = charmap.Macintosh.NewDecoder()
		e.encoder = charmap.Macintosh.NewEncoder()
	} else {
		return errors.New("Unknownd encoding")
	}
	return nil
}
