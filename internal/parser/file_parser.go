package fileparsers

import (
	"bufio"
	"errors"
	"net/url"
	"os"
	"project/internal/models"
	"strings"
)

func Parse(filePath string) ([]models.Task, error) { // TODO : aggregate from map to slice
	file, ok := os.Open(filePath)

	if ok != nil {
		return nil, ok
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	allInf := []models.Task{}

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())

		if len(words) < 2 {
			return nil, errors.New("There is must be at least Url and one tittle")
		}

		for i := range words {
			words[i] = strings.ReplaceAll(words[i], ",", "")
		}

		targetUrl, err := url.Parse(words[0])
		if err != nil {
			return nil, err
		}

		allInf = append(allInf, models.Task{Url: *targetUrl, Tittles: words[1:]})
	}

	return allInf, nil
}
