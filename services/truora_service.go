package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/ThaliaAC/labora-wallet/models"
	"github.com/joho/godotenv"
)

const (
	ContentType = "application/x-www-form-urlencoded"
	baseUrl     = "https://api.checks.truora.com/v1/checks"
)

var client = &http.Client{}
var apiKey string

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error loading .env file")
	}
	apiKey = string(os.Getenv("API_KEY"))
}

func getApiKey() string {
	return apiKey
}

func makeTruoraRequest(method, url string, payload *strings.Reader) ([]byte, error) {

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, fmt.Errorf("Error making request to API: %w", err)
	}

	req.Header.Add("Truora-API-Key", getApiKey())
	req.Header.Add("Content-Type", ContentType)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error making request to API: %w", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading the response API: %w", err)
	}

	return body, nil
}

func postTruoraRequest(person models.Person) (string, error) {
	truoraRequest := models.Api_Request_To_Truora{
		National_id:    person.National_id,
		Country:        person.Country,
		Type:           "person",
		UserAuthorized: true,
	}

	urlFull, data, err := getPostUrl(truoraRequest)
	urlStr := urlFull.String()
	payload := strings.NewReader(data.Encode())
	method := "POST"

	body, err := makeTruoraRequest(method, urlStr, payload)
	if err != nil {
		return "", fmt.Errorf("Error, failed to make POST request to API: %w", err)
	}

	var Response models.TruoraPostResponse
	err = json.Unmarshal(body, &Response)
	if err != nil {
		return "", fmt.Errorf("Error decoding the POST response API: %w", err)
	}

	checkID := Response.Check.CheckID

	return checkID, nil
}

func getTruoraScore(checkID string) (int, error) {
	url := baseUrl + checkID
	method := "GET"
	payload := strings.NewReader("")

	body, err := makeTruoraRequest(method, url, payload)
	if err != nil {
		return -1, fmt.Errorf("Error, failed to make GET request to API: %w", err)
	}

	var Response models.TruoraGetResponse
	err = json.Unmarshal(body, &Response)
	if err != nil {
		return -1, fmt.Errorf("Error decoding the GET response API: %w", err)
	}

	score := Response.Check.Score
	return score, err
}

func GetApproval(score int) (bool, error) {
	var truoraPostResponse models.TruoraPostResponse
	score, err := getTruoraScore(truoraPostResponse.Check.CheckID)

	if err != nil {

		return false, fmt.Errorf("Error,score request failed: %w", err)
	}

	return score == 1, err
}

func getPostUrl(truoraRequest models.Api_Request_To_Truora) (*url.URL, url.Values, error) {
	data := url.Values{}
	data.Set("national_id", truoraRequest.National_id)
	data.Set("country", truoraRequest.Country)
	data.Set("type", truoraRequest.Type)
	data.Set("userAuthorized", strconv.FormatBool(truoraRequest.UserAuthorized))

	urlFull, err := url.ParseRequestURI(baseUrl)
	if err != nil {
		return nil, nil, err
	}
	return urlFull, data, err
}
