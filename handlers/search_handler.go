package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gorilla/mux"
	"github.com/storefinder/pkg/elastic"
	"github.com/storefinder/pkg/models"
)

const serverError = "An error occurred. Please try again"

var (
	config elastic.ProxyConfig
)

func init() {
	elasticURL, _ := url.Parse(os.Getenv("ELASTIC_URL"))
	//uName = os.Getenv("USERNAME")
	//pwd = os.Getenv("PASSWORD")
	config = elastic.ProxyConfig{
		ElasticURL: elasticURL,
	}
}

//Search handler for store locator queries
func Search() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e models.Error
		var errors []models.Error
		var httpStatus = http.StatusOK
		var payload []byte
		var err error
		var queryRequest models.StoreQueryRequest
		var response = &models.StoreQueryResponse{}
		var esProxy *elastic.Proxy

		params := mux.Vars(r)
		indexName := params["indexName"]

		if len(indexName) == 0 {
			log.Println("Couldn't parse index name from path")
			httpStatus = http.StatusBadRequest
			e = models.Error{
				Message: "Couldn't parse index name from path",
			}
			errors = append(errors, e)
			goto done
		}
		defer r.Body.Close()
		payload, err = ioutil.ReadAll(r.Body)
		if err != nil {
			httpStatus = http.StatusBadRequest
			log.Printf("Error loading payload from request %v", err)
			e = models.Error{
				Message: serverError,
			}
			goto done
		}
		if err = json.Unmarshal(payload, &queryRequest); err != nil {
			httpStatus = http.StatusBadRequest
			log.Printf("Couldn't deserialize JSON to type : %v", err)
			e = models.Error{
				Message: serverError,
			}
			errors = append(errors, e)
			goto done
		}
		esProxy = elastic.NewProxy(config)
		if response, err = esProxy.Search(queryRequest, indexName); err != nil {
			httpStatus = http.StatusInternalServerError
			log.Println(err.Error())
			e = models.Error{
				Message: serverError,
			}
			errors = append(errors, e)
		}
	done:
		if len(errors) > 0 {
			response.Errors = errors
		}
		jr := JSONResponse{
			status: httpStatus,
			data:   response,
		}
		jr.Write(w)
	}
}
