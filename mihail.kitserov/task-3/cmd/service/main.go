package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"go.yaml.in/yaml/v3"
	"golang.org/x/text/encoding/charmap"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type Valute struct {
	ID       string `xml:"ID,attr"`
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nominal  int    `xml:"Nominal"`
	Name     string `xml:"Name"`
	Value    string `xml:"Value"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml:"Date,attr"`
	Valutes []Valute `xml:"Valute"`
}

func LoadConfig(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("error reading config file %q: %v", path, err))
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(fmt.Sprintf("error unmarshaling YAML config %q: %v", path, err))
	}

	if cfg.InputFile == "" {
		panic("in config not set field input-file")
	}

	if cfg.OutputFile == "" {
		panic("in config not set field output-file")
	}

	return cfg
}

func EnsureOutputPath(outputPath string) {
	dir := filepath.Dir(outputPath)

	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(fmt.Sprintf("error creating directory %q: %v", dir, err))
	}

	f, err := os.OpenFile(outputPath, os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(fmt.Sprintf("error creating file %q: %v", outputPath, err))
	}

	_ = f.Close()
}

func LoadValCurs(path string) ValCurs {
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("error opening input-file %q: %v", path, err))
	}

	defer file.Close()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		case "utf-8", "":
			return input, nil
		default:
			return nil, fmt.Errorf("error: unsupported charset: %s", charset)
		}
	}

	var valCurs ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		panic(fmt.Sprintf("error decoding XML from %q: %v", path, err))
	}

	return valCurs
}

func parseValue(s string) (float64, error) {
	normalized := strings.ReplaceAll(s, ",", ".")
	return strconv.ParseFloat(normalized, 64)
}

func main() {
	configPath := flag.String("config", "config.yaml", "path to the config file")
	flag.Parse()

	if *configPath == "" {
		panic("config path is empty")
	}

	cfg := LoadConfig(*configPath)
	EnsureOutputPath(cfg.OutputFile)

	valCurs := LoadValCurs(cfg.InputFile)
}
