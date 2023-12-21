package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

type mutateHandler[T any] func() (T, error)
type responseHandler[T any] func(resp T) error

func doMutate[T any](mutate mutateHandler[T], verify responseHandler[T]) (T, error) {
	resp, err := mutate()
	if err != nil {
		return *new(T), err
	}

	err = verify(resp)
	if err != nil {
		return *new(T), err
	}

	return resp, nil
}

func mutateError(localMessage string, serverCode string, serverMessage string) error {
	return fmt.Errorf("%s. code: %s message: %s", localMessage, serverCode, serverMessage)
}

func Ptr[T any](v T) *T {
	return &v
}

func ConvertObject[T any](from any) (*T, error) {
	b, err := json.Marshal(from)
	if err != nil {
		return nil, err
	}

	var result T
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func GetObjectFromFile[T any](file string) (*T, error) {
	inputJson, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	var obj T
	if err = json.Unmarshal(inputJson, &obj); err != nil {
		return nil, err
	}

	return &obj, nil
}

// A debugging function which dumps the HTTP response to stdout.
func DumpResponse(resp *http.Response) {
	fmt.Printf("response status: %s\n", resp.Status)
	dump, err := httputil.DumpResponse(resp, true)

	if err != nil {
		log.Printf("error dumping response: %s", err)
		return
	}

	log.Printf("response body: %s\n\n", string(dump))
}

// A debugging function which dumps the HTTP request to stdout.
func DumpRequest(req *http.Request) {
	if req.Body == nil {
		return
	}

	dump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Printf("error dumping request: %s", err)
		return
	}

	log.Printf("request body: %s\n\n", string(dump))
}

func sliceValueExists(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
