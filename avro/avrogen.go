package avro

import (
	"bytes"

	"github.com/elodina/go-avro"
	schemaregistry "github.com/landoop/schema-registry"
)

type MessageEvent struct {
	Filename string `json:"filename"`
	Filepath string `json:"filepath"`
}

var rawSchema = `{
	"type": "record",
	"name": "MessageEvent",
	"namespace": "ms1.avro",
	"fields": [
		{
		  "name": "filename",
		  "type": "string",
		  "doc": "this is the file name info",
		  "default": "null"
		},
		{
		  "name": "filepath",
		  "type": "string",
		  "doc": "this is the file path info from aws s3 bucket url",
		  "default": "null"
		}
	  ]
	}`

func AvroGen(name string, path string) {
	// schema := avro.MustParseSchema(rawSchema)
	schemas := avro.LoadSchemas("schemas/")
	schema, err := avro.ParseSchemaWithRegistry(rawSchema, schemas)
	if err != nil {
		panic(err)
	}
	record := avro.NewGenericRecord(schema)
	record.Set("filename", name)
	record.Set("filepath", path)

	writer := avro.NewGenericDatumWriter()
	writer.SetSchema(schema)
	buffer := new(bytes.Buffer)
	encoder := avro.NewBinaryEncoder(buffer)
	writer.Write(record, encoder)
	// record := avro.NewGenericRecord(schema)
	// record.Set("")
}

func SchemaRegistry() []string {
	client, _ := schemaregistry.NewClient(schemaregistry.DefaultURL)
	subjects, _ := client.Subjects()
	return subjects
}
