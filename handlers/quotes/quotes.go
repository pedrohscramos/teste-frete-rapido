package concursos

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	services "github.com/pedrohscramos/teste-frete-rapido/services/quote"
)

type QuoteInput struct {
	Recipient struct {
		Address struct {
			Zipcode string `json:"zipcode"`
		} `json:"address"`
	} `json:"recipient"`
	Volumes []struct {
		Amount        int     `json:"amount"`
		Category      int     `json:"category"`
		UnitaryWeight float64 `json:"unitary_weight"`
		Price         float64 `json:"price"`
		SKU           string  `json:"sku"`
		Height        float64 `json:"height"`
		Width         float64 `json:"width"`
		Length        float64 `json:"length"`
	} `json:"volumes"`
}

type QuoteRequest struct {
	Recipient struct {
		Zipcode int    `json:"zipcode"`
		Country string `json:"country"`
		Type    int    `json:"type"`
	} `json:"recipient"`

	Shipper struct {
		RegisteredNumber string `json:"registered_number"`
		Platform_code    string `json:"platform_code"`
		Token            string `json:"token"`
	} `json:"shipper"`
	Dispatchers    []interface{} `json:"dispatchers"`
	SimulationType []interface{} `json:"simulation_type"`
}

type QuoteResponse struct {
	Carrier []struct {
		Name     string  `json:"name"`
		Service  string  `json:"service"`
		Deadline string  `json:"deadline"`
		Price    float64 `json:"price"`
	} `json:"carrier"`
}

const (
	country            = "BRA"
	typeRecipient      = 0
	cnpjRemetente      = "25438296000158"
	tokenAutenticacao  = "1d52a9b6b78cf07b08586152459a5c90"
	codigoPlataforma   = "5AKVkHqCn"
	cepDispatchers     = 29161376
	unitaryPrice       = 1.0
	freteRapidoBaseURL = "https://sp.freterapido.com/api/v3/quote/simulate"
)

type QuoteHandler struct {
	service services.QuoteService
}

func NewQuoteHandler(service services.QuoteService) *QuoteHandler {
	return &QuoteHandler{
		service: service,
	}
}

func (handler *QuoteHandler) InsertQuote(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da solicitação", http.StatusBadRequest)
		return
	}

	var quoteInput QuoteInput
	err = json.Unmarshal(body, &quoteInput)
	if err != nil {
		http.Error(w, "Erro ao decodificar o JSON de entrada", http.StatusBadRequest)
		return
	}

	quoteRequestComplementado := complementarDadosFreteRapido(quoteInput)

	requestJSON, err := json.Marshal(quoteRequestComplementado)
	fmt.Println(string(requestJSON))
	if err != nil {
		http.Error(w, "Erro ao codificar a solicitação para a API da Frete Rápido", http.StatusInternalServerError)
		return
	}

	response, err := http.NewRequest("POST", freteRapidoBaseURL, bytes.NewBuffer(requestJSON))
	response.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(response)
	if err != nil {
		panic(err)
	}

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	bdd, _ := io.ReadAll(resp.Body)
	fmt.Println("response Body:", string(bdd))

	defer response.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Erro ao ler a resposta da API da Frete Rápido", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responseBody)
}

func complementarDadosFreteRapido(input QuoteInput) QuoteRequest {

	var request QuoteRequest

	zipCode, _ := strconv.Atoi(input.Recipient.Address.Zipcode)
	request.Recipient.Zipcode = zipCode

	documents := make([]interface{}, len(input.Volumes))

	for i := range input.Volumes {

		documents[i] = struct {
			Category      string  `json:"category"`
			Amount        int     `json:"amount"`
			UnitaryWeight float64 `json:"unitary_weight"`
			UnitaryPrice  float64 `json:"unitary_price"`
			SKU           string  `json:"sku"`
			Height        float64 `json:"height"`
			Width         float64 `json:"width"`
			Length        float64 `json:"length"`
		}{
			Category:      strconv.Itoa(input.Volumes[i].Category),
			Amount:        input.Volumes[i].Amount,
			UnitaryWeight: input.Volumes[i].UnitaryWeight,
			UnitaryPrice:  input.Volumes[i].Price,
			SKU:           input.Volumes[i].SKU,
			Height:        input.Volumes[i].Height,
			Width:         input.Volumes[i].Width,
			Length:        input.Volumes[i].Length,
		}

	}

	dispatchers := make([]interface{}, 1)

	dispatchers[0] = struct {
		RegisteredNumber string        `json:"registered_number"`
		Zipcode          int           `json:"zipcode"`
		Volumes          []interface{} `json:"volumes"`
	}{
		RegisteredNumber: cnpjRemetente,
		Zipcode:          cepDispatchers,
		Volumes:          documents,
	}

	simulation := make([]interface{}, 1)
	simulation[0] = []interface{}{0}
	request.SimulationType = simulation
	request.Dispatchers = dispatchers

	request.Shipper.RegisteredNumber = cnpjRemetente

	request.Shipper.Platform_code = codigoPlataforma
	request.Shipper.Token = tokenAutenticacao

	request.Recipient.Country = country
	request.Recipient.Type = typeRecipient
	request.SimulationType = []interface{}{0}

	return request
}
