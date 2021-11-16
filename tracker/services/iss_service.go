package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jefersondsgomes/iss-tracker/models"
)

func TrackISS() (response models.Response, err error) {
	apiResponse, err := http.Get("http://api.open-notify.org/iss-now.json")
	if err != nil {
		return response, err
	}

	defer apiResponse.Body.Close()
	body, err := ioutil.ReadAll(apiResponse.Body)
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return response, err
	}

	return response, nil
}

func FormatDate(d int) string {
	stringValue := strconv.Itoa(d)
	parsed, err := strconv.ParseInt(stringValue, 10, 64)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	result := time.Unix(parsed, 0)
	return result.Format("2006-01-02 15:04:05")
}
