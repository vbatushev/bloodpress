// main is main package
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"time"
)

var (
	inFile string
)

// BloodPressure - ...
type BloodPressure struct {
	Measuries []Measure `json:"bp_history"`
}

// Measure - ...
type Measure struct {
	ID          int   `json:"id"`
	MeasureTime int64 `json:"measureTime"`
	DIA         int   `json:"dia"`
	SYS         int   `json:"sys"`
	Pulse       int   `json:"pulse"`
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	inFile = os.Args[1]
	data, err := os.ReadFile(inFile)
	if err != nil {
		log.Fatal(err)
	}

	var bp = BloodPressure{}
	if err := json.Unmarshal(data, &bp); err != nil {
		log.Fatal(err)
	}

	sort.Slice(bp.Measuries[:], func(i, j int) bool {
		return bp.Measuries[i].ID < bp.Measuries[j].ID
	})

	start := time.UnixMilli(int64(bp.Measuries[0].MeasureTime))
	end := time.UnixMilli(int64(bp.Measuries[len(bp.Measuries)-1].MeasureTime))
	days := math.Round(end.Sub(start).Hours() / 24)
	fmt.Println(fmt.Sprintf("Средние значения за %v дней", days))

	var diaval int
	var sysval int
	var pulseval int
	size := len(bp.Measuries)

	for _, m := range bp.Measuries {
		diaval += m.DIA
		sysval += m.SYS
		pulseval += m.Pulse
	}

	fmt.Println(fmt.Sprintf("Верхнее: %d, нижнее: %d, пульс: %d", sysval/size, diaval/size, pulseval/size))
}
