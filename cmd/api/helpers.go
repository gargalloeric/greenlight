package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type envelop map[string]any

func (app *application) readIDParam(r *http.Request) (int64, error) {
	paramId := r.PathValue("id")

	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelop, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
