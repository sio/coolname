package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/sio/coolname/data"
)

// Convert upstream data files to Golang source code
func convert() (err error) {
	pack := make(datapack, len(data.UpstreamLists))
	for i := 0; i < len(data.UpstreamLists); i++ {
		pack[i], err = load(data.UpstreamLists[i])
		if err != nil {
			return err
		}
	}
	fmt.Printf("Loaded %d lists with a total of %d words\n", len(pack), pack.Size())

	tmpl, err := template.ParseFiles("codegen/data.go.template")
	if err != nil {
		return err
	}

	output, err := os.Create("data.go")
	if err != nil {
		return err
	}
	defer output.Close()

	err = tmpl.Execute(output, pack)
	if err != nil {
		return err
	}
	fmt.Printf("Word lists successfully written to %s\n", output.Name())

	fmt.Println("Removing intermediate *.txt files")
	for _, name := range data.UpstreamLists {
		err = os.Remove(name + ".txt")
		if err != nil {
			return err
		}
	}
	return nil
}

type datapack []*dataset

func (pack *datapack) Size() uint {
	var sum uint
	for _, ds := range *pack {
		sum += ds.Size()
	}
	return sum
}

type dataset struct {
	Name  string
	Words []string
}

func (ds *dataset) Size() uint {
	return uint(len(ds.Words))
}

func (ds *dataset) String() string {
	return fmt.Sprintf("dataset{Name: %s, Words: %d}", ds.Name, ds.Size())
}

// Load a single dataset into memory
func load(name string) (result *dataset, err error) {
	filename := name + ".txt"
	input, err := os.Open(filename)
	if err != nil {
		return &dataset{}, err
	}
	defer input.Close()

	result = &dataset{Name: name}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue // skip empty lines
		}
		if strings.HasPrefix(line, "#") {
			continue // skip comments
		}
		if strings.Contains(line, "=") {
			continue // skip inline configuration
		}
		result.Words = append(result.Words, line)
	}
	if err = scanner.Err(); err != nil {
		return &dataset{}, err
	}
	return result, nil
}
