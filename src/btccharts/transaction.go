// Package provides functions and structs for working with data from telnet
// btccharts.com
package btccharts

import "regexp"
import "strings"
import "io"
import "log"
import "encoding/json"

// Represents the a json messsage row from btccharts
type Message struct {
	Symbol, Volume, Id, Timestamp, Price string
}

// This map hardcodes a trader string to a float:
var TraderIdMap = map[string]float32{
	"bit2cILS":    0.01,
	"bitfloorUSD": 0.02,
	"bitstampUSD": 0.03,
	"btcdeEUR":    0.04,
	"btceRUR":     0.05,
	"btceUSD":     0.06,
	"btcnCNY":     0.07,
	"mtgoxEUR":    0.08,
	"mtgoxPLN":    0.09,
	"mtgoxUSD":    0.10,
	"rippleXRP":   0.11,
	"virwoxSLL":   0.12,
}


/*}*/

/*const input = "{\"volume\": 4.0, \"timestamp\": 1365812301, \"price\": 114.0, \"symbol\": \"virtexCAD\", \"id\": 21913359}"*/

// Parses a string from btccharts to json. We have to do some hacky reggex to
// make all values strings (fuck float64) before parsing it with std json go lib
func GetMessage(data string) Message {
	var tmp string
	tmp = regexp.MustCompile(`: `).ReplaceAllString(data, ": \"")
	tmp = regexp.MustCompile(`[,]`).ReplaceAllString(tmp, "\",")
	tmp = regexp.MustCompile(`[}]`).ReplaceAllString(tmp, "\"}")
	tmp = regexp.MustCompile(`["]{2}`).ReplaceAllString(tmp, "\"")
	log.Println(tmp)
	dec := json.NewDecoder(strings.NewReader(tmp))
	var m Message
	for {
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}

	return m
}
