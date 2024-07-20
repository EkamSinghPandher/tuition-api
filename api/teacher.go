package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"tuition-api/data"
	"tuition-api/utils"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Content struct {
	Parts []string `json:"Parts"`
	Role  string   `json:"Role"`
}
type Candidates struct {
	Content *Content `json:"Content"`
}
type ContentResponse struct {
	Candidates *[]Candidates `json:"Candidates"`
}

/*
Get the response from the Teacher, an LLM.
We format the student query and pass it to the teacher, returning the reponse to the request.
*/
func GetResponseFromTeacher(query data.StudentQuery) (ContentResponse, *utils.ErrorHandler) {
	var response ContentResponse

	apiKey, envApiKeyExists := os.LookupEnv("GEMINI_API_KEY")

	if !envApiKeyExists {
		var envErr = utils.ErrorHandler{
			Message: "Something went wrong parsing api key",
			Code:    500,
		}
		return response, &envErr
	}

	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	model.GenerateContent(context.Background())
	// requestBodyJSON, err := json.Marshal(teacherQuery)

	// if err != nil {
	// 	var marshalErr = utils.ErrorHandler{
	// 		Message: "Something went wrong marshalling the data",
	// 		Code:    500,
	// 	}
	// 	return response, &marshalErr
	// }

	// req, err := http.NewRequest(http.MethodPost, "https://api.anthropic.com/v1/messages", bytes.NewBuffer(requestBodyJSON))

	// fmt.Println(req)
	// if err != nil {
	// 	var createReqErr = utils.ErrorHandler{
	// 		Message: "Something went wrong in creating the request",
	// 		Code:    500,
	// 	}
	// 	return response, &createReqErr
	// }

	// req.Header.Add("x-api-key", apiKey)
	// req.Header.Add("anthropic-version", "2023-06-01")
	// req.Header.Add("content-type", "application/json")

	// res, err := http.DefaultClient.Do(req)

	// if err != nil {
	// 	var respErr = utils.ErrorHandler{
	// 		Message: "Something went wrong with the request to GPT",
	// 		Code:    500,
	// 	}
	// 	return response, &respErr
	// }

	// defer res.Body.Close()

	// fmt.Println(res)

	var prompt = []genai.Part{
		genai.Text("You are a teacher at a school, explain this to me in simple and concise terms."),
		genai.Text(query.StudentQuestion),
	}

	resp, err := model.GenerateContent(context.Background(), prompt...)
	if err != nil {
		log.Fatal(err)
	}
	marshalResponse, _ := json.MarshalIndent(resp, "", "  ")

	fmt.Println(resp)

	if err := json.Unmarshal(marshalResponse, &response); err != nil {
		log.Fatal(err)
	}

	return response, nil
}

func getTeacherQueryRequestFromStudentQuery(studentQuery data.StudentQuery) data.TeacherQueryRequest {
	var queryContent = []data.TeacherQueryContent{
		{Role: "system",
			Content: "You are a teacher, explain this question in a simple and informative way"},
		{Role: "user",
			Content: studentQuery.StudentQuestion},
	}

	var query = data.TeacherQueryRequest{
		Model:    "claude-3-5-sonnet-20240620",
		Messages: queryContent,
	}

	return query
}
