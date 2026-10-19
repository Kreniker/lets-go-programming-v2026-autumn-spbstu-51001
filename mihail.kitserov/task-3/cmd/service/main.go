package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
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

type ValuteJSON struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
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

func SortValutesByValueDesc(valCurs ValCurs) []Valute {
	sorted := make([]Valute, len(valCurs.Valutes))
	copy(sorted, valCurs.Valutes)

	sort.Slice(sorted, func(i, j int) bool {
		vi, errI := parseValue(sorted[i].Value)
		vj, errJ := parseValue(sorted[j].Value)

		if errI != nil || errJ != nil {
			return false
		}

		return vi > vj
	})

	return sorted
}

func SaveValutesJSON(valutes []Valute, outputPath string) {
	result := make([]ValuteJSON, 0, len(valutes))
	for _, v := range valutes {

		numCode, err := strconv.Atoi(v.NumCode)
		if err != nil {
			panic(fmt.Sprintf("Incorrect NumCode %q для %s: %v", v.NumCode, v.CharCode, err))
		}

		value, err := parseValue(v.Value)
		if err != nil {
			panic(fmt.Sprintf("Incorrect Value %q для %s: %v", v.Value, v.CharCode, err))
		}

		result = append(result, ValuteJSON{
			NumCode:  numCode,
			CharCode: v.CharCode,
			Value:    value,
		})
	}

	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		panic(fmt.Sprintf("error marshaling JSON: %v", err))
	}

	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(fmt.Sprintf("error creating directory %q: %v", dir, err))
	}

	if err := os.WriteFile(outputPath, data, 0644); err != nil {
		panic(fmt.Sprintf("error writing file %q: %v", outputPath, err))
	}
}

func (v *Valute) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
        for _, attr := range start.Attr {
        switch strings.ToLower(attr.Name.Local) {
        case "id":
            v.ID = attr.Value
        case "numcode", "vcode", "code":
            v.NumCode = strings.TrimSpace(attr.Value)
        case "charcode", "vchcode", "chcode":
            v.CharCode = strings.TrimSpace(attr.Value)
        case "nominal", "vnom", "nom":
            v.Nominal, _ = strconv.Atoi(strings.TrimSpace(attr.Value))
        case "name", "vname":
            v.Name = attr.Value
        case "value", "vcurs", "curs":
            v.Value = strings.TrimSpace(attr.Value)
        }
    }

		for {
        tok, err := d.Token()
        if err != nil {
            return err
        }
        switch t := tok.(type) {
        case xml.StartElement:
            var content string
            if err := d.DecodeElement(&content, &t); err != nil {
                return err
            }
            content = strings.TrimSpace(content)
            switch strings.ToLower(t.Name.Local) {
            case "numcode", "vcode", "code":
                v.NumCode = content
            case "charcode", "vchcode", "chcode":
                v.CharCode = content
            case "nominal", "vnom", "nom":
                v.Nominal, _ = strconv.Atoi(content)
            case "name", "vname":
                v.Name = content
            case "value", "vcurs", "curs":
                v.Value = content
            }
        case xml.EndElement:
            if t.Name == start.Name {
                return nil
            }
        }
    }
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

	sortedValutes := SortValutesByValueDesc(valCurs)

	SaveValutesJSON(sortedValutes, cfg.OutputFile)
}
