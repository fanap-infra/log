package log

import (
	"bytes"
	"fmt"
	"sync"
	"time"

	influxdb "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

const measurement = "logs"

type Influx struct {
	pool     sync.Pool
	client   influxdb.Client
	writeAPI api.WriteAPI
	app      string
}

func InfluxWriter(serverURL string, authToken string, org string, bucket string, app string, caller bool, stack EnablerFunc, enabler EnablerFunc) *Writer {
	i := &Influx{
		pool: sync.Pool{New: func() interface{} {
			b := bytes.NewBuffer(make([]byte, 150)) // buffer init with 150 size
			b.Reset()
			return b
		}},
		app: app,
	}
	i.Connect(serverURL, authToken, org, bucket)
	return newWriter(enabler, stack, caller, i)
}

func (i *Influx) Connect(serverURL string, authToken string, org string, bucket string) {
	i.client = influxdb.NewClient(serverURL, authToken)
	i.writeAPI = i.client.WriteAPI(org, bucket) // https://docs.influxdata.com/influxdb/v2.0/write-data/
}

func (i *Influx) close() {
	// Force all unwritten data to be sent
	i.writeAPI.Flush()
	// Ensures background processes finishes
	i.client.Close()
}

func (i *Influx) getBuffer() *bytes.Buffer {
	return i.pool.Get().(*bytes.Buffer)
}

func (i *Influx) putBuffer(b *bytes.Buffer) {
	b.Reset()
	i.pool.Put(b)
}

func (i *Influx) Print(l Level, s string, caller string, stacks []string, message string) {
	// create point
	p := influxdb.NewPoint(
		measurement,
		map[string]string{
			"app":   i.app,
			"scope": s,
			"level": levelText[l],
		},
		map[string]interface{}{
			"message": message,
			"caller":  caller,
			"stacks":  stacks,
		},
		time.Now())

	// write asynchronously
	i.writeAPI.WritePoint(p)
}

func (i *Influx) Printv(l Level, s string, caller string, stacks []string, message string, keysValues []interface{}) {
	fields := i.getBuffer()
	defer i.putBuffer(fields)

	ln := len(keysValues)
	if ln > 1 {
		fields.WriteByte('{')
		i := 0
		for {
			if key, ok := keysValues[i].(string); ok {
				fields.WriteString(`"` + key + `":"` + fmt.Sprint(keysValues[i+1]) + `"`)
			}

			i += 2
			if i == ln {
				break
			}
			fields.WriteByte(',')
		}
		fields.WriteByte('}')
	}

	// create point
	p := influxdb.NewPoint(
		measurement,
		map[string]string{
			"app":   i.app,
			"scope": s,
			"level": levelText[l],
		},
		map[string]interface{}{
			"message": message,
			"caller":  caller,
			"stacks":  stacks,
			"data":    fields,
		},
		time.Now())

	// write asynchronously
	i.writeAPI.WritePoint(p)
}
