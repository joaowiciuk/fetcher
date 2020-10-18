package parser

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/wiciuk-dev/fetcher/pkg/model"
)

// Parser abstracts a parser that can extract information from any source.
type Parser interface {
	Parse(source io.Reader) ([]model.Book, error)
}

type JSONParser struct {
}

func (jp *JSONParser) Parse(source io.Reader) ([]model.Book, error) {
	var books []model.Book
	err := json.NewDecoder(source).Decode(&books)
	if err != nil {
		return nil, fmt.Errorf("error decoding source: %w", err)
	}
	return books, nil
}

type CSVParser struct {
}

func (cp *CSVParser) Parse(source io.Reader) ([]model.Book, error) {
	r := csv.NewReader(source)
	books := make([]model.Book, 0)
	header, _ := r.Read()
	if len(header) < 4 {
		return nil, errors.New("error parsing books: invalid source")
	}
	indexes := make(map[string]int)
	for i, h := range header {
		indexes[strings.ToLower(h)] = i
	}
	for {
		records, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("error reading from source: %w", err)
		}

		if len(records) != len(indexes) {
			return nil, errors.New("error parsing books: invalid csv")
		}

		id, err := strconv.ParseInt(records[indexes["id"]], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing id: %w", err)
		}

		likes, err := strconv.ParseInt(records[indexes["likes"]], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing likes: %w", err)
		}

		book := model.Book{
			ID:    int(id),
			Title: records[indexes["title"]],
			ISBN:  records[indexes["isbn"]],
			Likes: int(likes),
		}
		books = append(books, book)
	}
	return books, nil
}
