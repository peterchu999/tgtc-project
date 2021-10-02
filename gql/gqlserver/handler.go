package gqlserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

type postData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operationName"`
	Variables map[string]interface{} `json:"variables"`
}

func Handle() http.Handler {
	return http.HandlerFunc( func(rw http.ResponseWriter, r *http.Request) {
			fmt.Print("Got Herer")
			var p postData
			if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
				rw.WriteHeader(400)
				// rw.WriteHeader(400)
				fmt.Print(err)
				return
			}
			fmt.Print(p)
			result := graphql.Do(graphql.Params{
				Schema:         Init(),
				RequestString:  p.Query,
				VariableValues: p.Variables,
				OperationName:  p.Operation,
			})
			if len(result.Errors) > 0 {
				log.Println("[GQLHandler][Handle] there were some errors, errs: ", result.Errors)
			}

			rw.Header().Set("Content-Type", "application/json")
			json.NewEncoder(rw).Encode(result)
		},
	)
}