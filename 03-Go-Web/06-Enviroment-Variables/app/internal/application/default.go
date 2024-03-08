package application

import (
	"bufio"
	"os"
	"time"

	"github.com/fatih/color"
)

// ConfigDefaultApp is the configuration for the default application
type ConfigDefaultApp struct {
	Title    string `json:"title" yaml:"title"`
	Color    string `json:"color" yaml:"color"`
	FilePath string `json:"file_path" yaml:"file_path"`
}

// NewDefaultApp creates a new default application
func NewDefaultApp(cfg ConfigDefaultApp) *DefaultApp {
	// default the title
	if cfg.Title == "" {
		cfg.Title = "Default Application"
	}

	return &DefaultApp{
		title:    cfg.Title,
		color:    cfg.Color,
		filePath: cfg.FilePath,
	}
}

// DefaultApp is the default application
type DefaultApp struct {
	// title of the application
	title    string
	// color of the output
	color    string
	// file path to read
	filePath string
}

func (d *DefaultApp) Run() (err error) {
	// display the title
	println("Title:", d.title)

	// open the file
	file, err := os.Open(d.filePath)
	if err != nil {
		return
	}
	defer file.Close()

	// read the file
	rd := bufio.NewReader(file)
	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return err
		}
		lineString := string(line)

		// simulate processing
		time.Sleep(1 * time.Second)

		// display the line
		switch d.color {
		case "red":
			color.Red(lineString)
		case "green":
			color.Green(lineString)
		default:
			println(lineString)
		}
	}

	return
}