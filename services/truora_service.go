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
	"time"

	"github.com/ThaliaAC/labora-wallet/models"
	"github.com/joho/godotenv"
)

const ContentType = "application/x-www-form-urlencoded"

func makeRequest(method, url string, payload *strings.Reader) ([]byte, error) {
	client := &http.Client{}
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

func getCheckID(truoraRequest models.Api_Request_To_Truora) (string, error) {
	urlFull, data := getPostUrl(truoraRequest)
	urlStr := urlFull.String()
	payload := strings.NewReader(data.Encode())
	method := "POST"

	body, err := makeRequest(method, urlStr, payload)
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

func CreateRequest(person models.Person) (int, models.Wallet, error) {
	var wallet models.Wallet
	var log models.Log

	truoraRequest := models.Api_Request_To_Truora{
		National_id:    person.National_id,
		Country:        person.Country,
		Type:           "person",
		UserAuthorized: true,
	}

	autorization, err := getApproval(truoraRequest)
	if err != nil {
		return http.StatusInternalServerError, models.Wallet{}, fmt.Errorf("API request failed %w", err)
	}
	wallet.National_id = truoraRequest.National_id
	wallet.Country = truoraRequest.Country
	wallet.RequestDate = time.Now()

	if !autorization {
		log.National_id = wallet.National_id
		log.Country = wallet.Country
		log.Status = "Denied"
		log.RequestType = "CREATE WALLET"
		err = (*WalletService).CreateLog(log)
		if err != nil {

			return http.StatusInternalServerError, models.Wallet{}, fmt.Errorf("Error creating the log: %w", err)
		}

		return http.StatusConflict, models.Wallet{}, nil
	}

	log.National_id = wallet.National_id
	log.Country = wallet.Country
	log.Status = "Approved"
	log.RequestType = "CREATE WALLET"

	wallet, err = (*WalletService).CreateWallet(wallet, log)
	if err != nil {

		return http.StatusInternalServerError, models.Wallet{}, fmt.Errorf("Error creating the wallet %w", err)
	}

	return http.StatusOK, wallet, nil
}

func getTruoraAPIRequest(checkID string) (string, error) {
	url := "https://api.checks.truora.com/v1/checks/" + checkID
	method := "GET"
	payload := strings.NewReader("")

	body, err := makeRequest(method, url, payload)
	if err != nil {

		return "-1", fmt.Errorf("Error, failed to make GET request to API: %w", err)
	}

	var Response models.TruoraGetResponse
	err = json.Unmarshal(body, &Response)
	if err != nil {

		return "-1", fmt.Errorf("Error decoding the GET response API: %w", err)
	}

	score := strconv.Itoa(Response.Check.Score)

	return score, nil
}

func getBackgroundCheck(truoraRequest models.Api_Request_To_Truora) (string, error) {
	checkID, err := getCheckID(truoraRequest)
	if err != nil {

		return "-1", fmt.Errorf("Post request failed: %w", err)
	}

	time.Sleep(5 * time.Second)

	score, err := getTruoraAPIRequest(checkID)
	if err != nil {

		return "-1", fmt.Errorf("Get request failed: %w", err)
	}

	return score, nil
}

func getApproval(truoraRequest models.Api_Request_To_Truora) (bool, error) {
	score, err := getBackgroundCheck(truoraRequest)

	if err != nil {

		return false, fmt.Errorf("Error,score request failed: %w", err)
	}

	return score == "1", err

}

func getPostUrl(truoraRequest models.Api_Request_To_Truora) (*url.URL, url.Values) {
	apiTruoraUrl := "https://api.checks.truora.com"
	resource := "/v1/checks"
	data := url.Values{}
	data.Set("national_id", truoraRequest.National_id)
	data.Set("country", truoraRequest.Country)
	data.Set("type", truoraRequest.Type)
	data.Set("userAuthorized", strconv.FormatBool(truoraRequest.UserAuthorized))

	urlFull, _ := url.ParseRequestURI(apiTruoraUrl)
	urlFull.Path = resource
	return urlFull, data
}

func getApiKey() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error loading .env file")
	}
	apiKey := string(os.Getenv("API_KEY"))
	return apiKey
}
