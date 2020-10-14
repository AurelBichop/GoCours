package main

import (
	"flag"
	"fmt"
	"os"

	"training.go/GenCert/cert"
	"training.go/GenCert/cert/html"
	"training.go/GenCert/pdf"
)

func main() {
	inputFileCsv := flag.String("file", "", "Input Csv for parse.")
	outputType := flag.String("type", "pdf", "Output Type of the certificate.")
	flag.Parse()

	if len(*inputFileCsv) <= 0 {
		fmt.Printf("Invalid File. got=%v\n", *inputFileCsv)
		os.Exit(1)
	}

	var saver cert.Saver
	var err error

	switch *outputType {
	case "html":
		saver, err = html.New("outpout")
	case "pdf":
		saver, err = pdf.New("outpout")
	default:
		fmt.Printf("Unknown output type. got=%v\n", *outputType)
	}

	if err != nil {
		fmt.Printf("Could not create the generator : %v", err)
		os.Exit(1)
	}

	//c, err := cert.New("Golang programming", "Bob dylan", "2018-06-21")
	certs, err := cert.ParseCSV(*inputFileCsv)
	if err != nil {
		fmt.Printf("Could not parse Csv file : %v", err)
		os.Exit(1)
	}

	for _, c := range certs {
		err = saver.Save(*c)
		if err != nil {
			fmt.Printf("Could not save cert. get=%v\n", err)
		}
	}
}
