package euler

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/Nooby/EulerGo/lib/data"
)

//go:generate esc -o data/static.go -ignore "static.go" -pkg data data

var (
	fs           http.FileSystem
	descriptions []string
	solutions    []string
	resources    map[int]interface{}
)

func init() {
	err := initProjectEuler()
	if err != nil {
		log.Fatal(err)
	}
}

func Verify(challenge int, solution string) (bool, error) {
	if len(solutions) < challenge {
		return false, fmt.Errorf("solution not available for challenge %v", challenge)
	}

	sol := solutions[challenge-1]
	return sol == solution, nil
}

func initProjectEuler() error {
	fs = data.FS(false)
	err := initDescriptions()
	if err != nil {
		return fmt.Errorf("problem descriptions not available: %v", err)
	}
	err = initSolutions()
	if err != nil {
		return fmt.Errorf("problem solutions not available: %v", err)
	}
	err = initResources()
	if err != nil {
		return fmt.Errorf("problem resources not available: %v", err)
	}
	return nil
}

func initDescriptions() error {
	file, err := fs.Open("/data/problems.txt")
	if err != nil {
		return fmt.Errorf("couldn't open problems descriptions: %v", err)
	}
	defer file.Close()

	text, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("couldn't read problems file: %v", err)
	}
	descriptions = strings.Split(string(text), "\n\n\n")
	log.Printf("%v Descriptions Parsed.", len(descriptions))

	return nil
}

func initSolutions() error {
	solutions = make([]string, 0)
	file, err := fs.Open("/data/solutions.txt")
	if err != nil {
		return fmt.Errorf("couldn't open solution file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), " ")
		solutions = append(solutions, splitLine[1])
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("couldn't read problems file: %v", err)
	}
	log.Printf("%v Solutions Parsed.", len(solutions))

	return nil
}

func initResources() error {
	file, err := fs.Open("/data/resources.json")
	if err != nil {
		return fmt.Errorf("couldn't open resources file: %v", err)
	}
	defer file.Close()

	text, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("couldn't read resources file: %v", err)
	}

	json.Unmarshal(text, &resources)
	log.Printf("%v Resources Parsed.", len(resources))

	for _, v := range [2]int{252, 255} {
		fmt.Printf("Key: %v, Value: %v\n", v, resources[v])
		var t interface{}
		t = resources[v]
		switch t := t.(type) {
		default:
			fmt.Printf("unexpected type %T\n", t) // %T prints whatever type t has
		case string:
			fmt.Printf("string %v\n", t)
		case []interface{}:
			fmt.Printf("[]string %v\n", t) // t has type int
			for i := range t {
				fmt.Printf("element %v\n", t[i])
			}

		}
	}

	return nil
}
