package bowald

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"time"
)

type Exam struct {
	NotPassed int `json:notPassed`
	Three int `json:"three"`
	Four int `json:"four"`
	Five int `json:"five"`
	Date time.Time `json:"date"`
}
type Course struct {
	Id string `json:"_id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Owner string `json:"owner"`
	Exams []Exam `json:"exams"`
}

func Search(term string) ([]Course, error) {
	url := fmt.Sprintf("http://tenta.bowald.se/api/courses?searchterm=%s", term)
	res, err := http.Get(url)

	if err != nil {
		return nil,err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil,err
	}

	response := []Course{}
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println(err)
		return nil,err
	}

	for _, course := range response {
		sort.Slice(course.Exams, func(i,j int) bool {
			return course.Exams[i].Date.After(course.Exams[j].Date)
		})
	}

	return response, nil
}