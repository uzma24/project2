package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/project2/internal/utils"
	"github.com/project2/models"
)

func GetFrequency(c *gin.Context) (models.OutputJSON, error) {
	response := models.OutputJSON{}
	text, err := ioutil.ReadFile("Golang_Test.txt")
	if err != nil {
		fmt.Printf("Error found while file reading: %v", err)
		return response, err
	}

	var input_in_JSON models.InputText
	input_in_JSON.Text = string(text)
	payload, err := json.Marshal(input_in_JSON)

	if err != nil {
		fmt.Printf("Error in marshalling text: %v", err)
		return response, err
	}
	data, statusCode, err := callGetFrequencyService(payload)

	if err != nil {
		return response, &utils.InternalServerError{ErrMessage: err.Error()}
	}
	if statusCode != http.StatusOK {
		return response, &utils.InternalServerError{ErrMessage: "Status code error"}
	}

	var outputJson models.OutputJSON
	err = json.Unmarshal(data, &outputJson)
	if err != nil {
		fmt.Printf("error while unmarshalling: %v", err)
		return response, err
	}
	response = outputJson
	return response, nil
}

func callGetFrequencyService(payload []byte) ([]byte, int, error) {
	host := "http://localhost:8080"
	txnURI := "/text"
	url := host + txnURI
	return utils.GetHttp(url, payload)
}
