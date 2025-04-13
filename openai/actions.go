package openai

import (
	"bufio"
	"encoding/json"
	"errors"
	"log"
	"math"
	"os"
)

var Actions []*Action

func parseAction(line string) (err error) {
	if line == "" {
		return nil
	}
	var action Action
	err = json.Unmarshal([]byte(line), &action)
	if err != nil {
		return err
	}
	Actions = append(Actions, &action)
	return nil
}

func LoadActions() {
	log.Println("loading actions...")

	file, err := os.Open("./actions.jsonl")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		err = parseAction(line)
		if err != nil {
			log.Println(err)
			return
		}
	}

	if err = scanner.Err(); err != nil {
		log.Println(err)
	}
	log.Printf("%d actions loaded", len(Actions))
}

func SendEmbeddingRequest(input string) ([]float64, error) {
	params := EmbedParams{
		Input: input,
		Model: "text-embedding-3-small",
	}

	paramsData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	response, err := SendRequest("POST", "/v1/embeddings", paramsData)

	var embeddingResponse EmbeddingResponse
	err = json.Unmarshal(response, &embeddingResponse)
	if err != nil {
		return nil, err
	}

	if len(embeddingResponse.Data) == 0 {
		return nil, errors.New("no embeddings found")
	}

	vectors := embeddingResponse.Data[0].Embedding

	return vectors, nil
}

func GetActionsDistance(vectors []float64, actionVectors []float64) float64 {
	var distance float64
	distance = 0.0
	for i := 0; i < len(vectors); i++ {
		distance += math.Pow(vectors[i]-actionVectors[i], 2)
	}
	return math.Sqrt(distance)
}

func GetBestAction(vectors []float64) *Action {
	var bestAction *Action
	bestDistance := math.MaxFloat64
	for _, action := range Actions {
		dist := GetActionsDistance(vectors, action.Vector)
		if dist < bestDistance {
			bestDistance = dist
			bestAction = action
		}
	}
	// not good enough
	if bestDistance > 1.2 {
		return nil
	}
	return bestAction
}
