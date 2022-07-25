package service

import (
	"fmt"
	"mime/multipart"
	"sync"

	"github.com/sshaunn/train-go/ms1/config"
)

var (
	s3service    *config.S3Service    = new(config.S3Service).NewS3Service()
	kafkaService *config.KafkaService = new(config.KafkaService).NewKafkaService()
)

type Executable func()

type MessageEvent struct {
	Filename string `json:"filename"`
	Filepath string `json:"filepath"`
}

func UploadFiles(files []*multipart.FileHeader, ch chan interface{}, wg *sync.WaitGroup) func() {
	// s3service.Save(name, f)

	return func() {
		for _, fileHeader := range files {
			file, err := fileHeader.Open()
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()
			go func(name string, f multipart.File) {
				defer wg.Done()
				err = s3service.Save(name, f)
				if err != nil {
					fmt.Println(err)
				}
				ch <- &MessageEvent{
					Filename: name,
					Filepath: "https://" + config.BUCKET_NAME + ".s3.amazonaws.com/" + name,
				}
			}(fileHeader.Filename, file)
		}
	}
}

func ProduceMessage(ch chan interface{}, s []interface{}) func() {
	return func() {
		// for msg := range ch {
		go func(ss []interface{}) {
			for msg := range ch {
				go kafkaService.Produce(msg)
				// ss = append(ss, msg)
			}
			defer close(ch)
		}(s)
	}
}

// f1 -> upload to s3 -> if successful -> produce msg to kafka
// f2 -> upload to s3 -> if successful -> produce msg to kafka
// f3 -> upload to s3 -> if successful -> produce msg to kafka
// func fakeProducer(msg interface{}) {
// 	fmt.Println(msg)
// }

// fmt.Println()
func ConcurrencyWorkflow(fns ...Executable) {
	fmt.Println(len(fns), "fns in flow")
	// wg.Add(len(files))
	for _, fn := range fns {
		fn()
	}
}
