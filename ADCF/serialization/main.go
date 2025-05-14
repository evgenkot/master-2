package main

import (
	"bytes"
	"encoding/csv"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/vmihailenco/msgpack/v5"
	"gopkg.in/yaml.v3"
	"google.golang.org/protobuf/proto"
	"github.com/hamba/avro/v2"
	"serialization/datastruct"
)

type DataStruct struct {
	StringField string
	IntField    int
	FloatField  float64
	ArrayField  []string
	MapField    map[string]int
}

type FormatResult struct {
	Format          string
	Size            int
	SerializeTime   time.Duration
	DeserializeTime time.Duration
}

// Avro serialization
const avroSchema = `{
    "type": "record",
    "name": "DataStruct",
    "fields": [
        {"name": "StringField", "type": "string"},
        {"name": "IntField", "type": "int"},
        {"name": "FloatField", "type": "double"},
        {"name": "ArrayField", "type": {"type": "array", "items": "string"}},
        {"name": "MapField", "type": {"type": "map", "values": "int"}}
    ]
}`

// XML helper types
type XMLMapEntry struct {
	Key   string `xml:"Key"`
	Value int    `xml:"Value"`
}

type XMLDataStruct struct {
	StringField string        `xml:"StringField"`
	IntField    int           `xml:"IntField"`
	FloatField  float64       `xml:"FloatField"`
	ArrayField  []string      `xml:"ArrayField>element"`
	MapField    []XMLMapEntry `xml:"MapField>Entry"`
}

// Avro serialization
func avroSerialize(data DataStruct) ([]byte, error) {
	schema := avro.MustParse(avroSchema)
	return avro.Marshal(schema, data)
}

func avroDeserialize(b []byte) (DataStruct, error) {
	var d DataStruct
	schema := avro.MustParse(avroSchema)
	err := avro.Unmarshal(schema, b, &d)
	return d, err
}


// JSON serialization
func jsonSerialize(data DataStruct) ([]byte, error) {
	return json.Marshal(data)
}

func jsonDeserialize(b []byte) (DataStruct, error) {
	var d DataStruct
	err := json.Unmarshal(b, &d)
	return d, err
}

// XML serialization
func xmlSerialize(data DataStruct) ([]byte, error) {
	xmlMap := make([]XMLMapEntry, 0, len(data.MapField))
	for k, v := range data.MapField {
		xmlMap = append(xmlMap, XMLMapEntry{Key: k, Value: v})
	}

	xmlData := XMLDataStruct{
		StringField: data.StringField,
		IntField:    data.IntField,
		FloatField:  data.FloatField,
		ArrayField:  data.ArrayField,
		MapField:    xmlMap,
	}
	return xml.Marshal(xmlData)
}

func xmlDeserialize(b []byte) (DataStruct, error) {
	var xmlData XMLDataStruct
	if err := xml.Unmarshal(b, &xmlData); err != nil {
		return DataStruct{}, err
	}

	dataMap := make(map[string]int, len(xmlData.MapField))
	for _, entry := range xmlData.MapField {
		dataMap[entry.Key] = entry.Value
	}

	return DataStruct{
		StringField: xmlData.StringField,
		IntField:    xmlData.IntField,
		FloatField:  xmlData.FloatField,
		ArrayField:  xmlData.ArrayField,
		MapField:    dataMap,
	}, nil
}

// YAML serialization
func yamlSerialize(data DataStruct) ([]byte, error) {
	return yaml.Marshal(data)
}

func yamlDeserialize(b []byte) (DataStruct, error) {
	var d DataStruct
	err := yaml.Unmarshal(b, &d)
	return d, err
}

// MessagePack serialization
func msgpackSerialize(data DataStruct) ([]byte, error) {
	return msgpack.Marshal(data)
}

func msgpackDeserialize(b []byte) (DataStruct, error) {
	var d DataStruct
	err := msgpack.Unmarshal(b, &d)
	return d, err
}

// Gob serialization
func gobSerialize(data DataStruct) ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func gobDeserialize(b []byte) (DataStruct, error) {
	var d DataStruct
	buf := bytes.NewBuffer(b)
	decoder := gob.NewDecoder(buf)
	if err := decoder.Decode(&d); err != nil {
		return d, err
	}
	return d, nil
}

// Protobuf serialization
func protoSerialize(data DataStruct) ([]byte, error) {
	p := &datastruct.ProtoDataStruct{
		StringField: data.StringField,
		IntField:    int32(data.IntField),
		FloatField:  data.FloatField,
		ArrayField:  data.ArrayField,
		MapField:    make(map[string]int32, len(data.MapField)),
	}
	for k, v := range data.MapField {
		p.MapField[k] = int32(v)
	}
	return proto.Marshal(p)
}

func protoDeserialize(b []byte) (DataStruct, error) {
	p := &datastruct.ProtoDataStruct{}
	if err := proto.Unmarshal(b, p); err != nil {
		return DataStruct{}, err
	}

	dataMap := make(map[string]int, len(p.MapField))
	for k, v := range p.MapField {
		dataMap[k] = int(v)
	}

	return DataStruct{
		StringField: p.StringField,
		IntField:    int(p.IntField),
		FloatField:  p.FloatField,
		ArrayField:  p.ArrayField,
		MapField:    dataMap,
	}, nil
}

func benchmarkSerialize(data DataStruct, serializeFunc func(DataStruct) ([]byte, error), iterations int) (time.Duration, int, error) {
	var totalTime time.Duration
	var size int

	for range make([]int, iterations) {
		start := time.Now()
		bytes, err := serializeFunc(data)
		if err != nil {
			return 0, 0, err
		}
		totalTime += time.Since(start)
		size = len(bytes)
	}

	return totalTime / time.Duration(iterations), size, nil
}

func benchmarkDeserialize(data DataStruct, serializeFunc func(DataStruct) ([]byte, error), deserializeFunc func([]byte) (DataStruct, error), iterations int) (time.Duration, error) {
	bytes, err := serializeFunc(data)
	if err != nil {
		return 0, err
	}

	var totalTime time.Duration
	for range make([]int, iterations) {
		start := time.Now()
		_, err := deserializeFunc(bytes)
		if err != nil {
			return 0, err
		}
		totalTime += time.Since(start)
	}

	return totalTime / time.Duration(iterations), nil
}

func writeCSV(results []FormatResult) error {
	file, err := os.Create("report.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Format", "Size (bytes)", "Serialize (ns)", "Deserialize (ns)"}
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, res := range results {
		record := []string{
			res.Format,
			strconv.Itoa(res.Size),
			strconv.FormatInt(res.SerializeTime.Nanoseconds(), 10),
			strconv.FormatInt(res.DeserializeTime.Nanoseconds(), 10),
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	// Create large test data
	arrayField := make([]string, 100)
	for i := range arrayField {
		arrayField[i] = fmt.Sprintf("element%d", i)
	}

	mapField := make(map[string]int, 100)
	for i := range make([]int, 100) {
		mapField[fmt.Sprintf("key%d", i)] = i
	}

	data := DataStruct{
		StringField: string(bytes.Repeat([]byte("a"), 1024)),
		IntField:    42,
		FloatField:  3.14159,
		ArrayField:  arrayField,
		MapField:    mapField,
	}

	formats := []struct {
		name string
		ser  func(DataStruct) ([]byte, error)
		des  func([]byte) (DataStruct, error)
	}{
		{"JSON", jsonSerialize, jsonDeserialize},
		{"XML", xmlSerialize, xmlDeserialize},
		{"YAML", yamlSerialize, yamlDeserialize},
		{"MsgPack", msgpackSerialize, msgpackDeserialize},
		{"Gob", gobSerialize, gobDeserialize},
		{"Protobuf", protoSerialize, protoDeserialize},
		{"Avro", avroSerialize, avroDeserialize},
	}

	const iterations = 1000
	var results []FormatResult

	for _, fmt := range formats {
		serializeTime, size, err := benchmarkSerialize(data, fmt.ser, iterations)
		if err != nil {
			log.Printf("Serialization error for %s: %v", fmt.name, err)
			continue
		}

		deserializeTime, err := benchmarkDeserialize(data, fmt.ser, fmt.des, iterations)
		if err != nil {
			log.Printf("Deserialization error for %s: %v", fmt.name, err)
			continue
		}

		results = append(results, FormatResult{
			Format:          fmt.name,
			Size:            size,
			SerializeTime:   serializeTime,
			DeserializeTime: deserializeTime,
		})
	}

	if err := writeCSV(results); err != nil {
		log.Fatalf("CSV write error: %v", err)
	}

	fmt.Println("Benchmark results saved to report.csv")
}
