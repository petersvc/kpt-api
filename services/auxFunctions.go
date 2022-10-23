package services

import (
	"fmt"
	"kpt_api/model"
	"net/http"
	"strconv"
	"strings"
)

func GetFilters(r *http.Request) []model.Filters {
	requestedURL := r.URL.Query()
	// create a slice of arrays
	var parameters []model.Filters
	for k, v := range requestedURL {
		if k != "page" {
			// fmt.Println("key: ", k)
			// fmt.Println("value: ", v)
			parameter := model.Filters{Key: k} // }
			if strings.Contains(v[0], ",") {
				splitString := strings.Split(v[0], ",")
				// fmt.Println("splitString: ", splitString)
				// fmt.Println("parameter: ", parameter)
				parameter.Value = append(parameter.Value, splitString...)
			} else {
				parameter.Value = append(parameter.Value, v[0])
			}
			// fmt.Println("parameter: ", parameter)
			parameters = append(parameters, parameter)
		}
	}
	return parameters
}

func IndexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func GetLimit(r *http.Request) int64 {
	filter := r.URL.Query().Get("limit")
	limit, _ := strconv.Atoi(filter)

	if limit == 0 {
		limit = 30
	}

	fmt.Println("limit: ", limit)

	return int64(limit)
}

func GetPage(r *http.Request) int64 {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	// fmt.Println("page: ", page)
	return int64(page)
}

func GetSortOptions(r *http.Request) (string, int8) {
	key := r.URL.Query().Get("sort")
	typeSortFilter := r.URL.Query().Get("typeSort")
	typeSort, _ := strconv.Atoi(typeSortFilter)

	// if key == "" {
	// 	key = "priceInt"
	// }

	if typeSort == 0 {
		typeSort = 1
	}

	fmt.Println("sortKey: ", key)
	fmt.Println("typeSort: ", typeSort)
	// options.Find().SetSort(bson.D{{Key: sort, Value: 1}})
	return key, int8(typeSort)
}

func ValidateQuery(query string) {
	queryTrueCounter := 0

	if strings.Contains(query, "/api/gpus?") {
		// fmt.Println("validating step 1: success")
		if strings.Index(query, "/api/gpus?") == 0 {
			// fmt.Println("validating step 2: success")
			parameters := query[len("/api/gpus?") : len(query)-0]

			if strings.Contains(parameters, "limit=") {
				queryTrueCounter++
			}
			if strings.Contains(parameters, "page=") {
				queryTrueCounter++
			}
			if strings.Contains(parameters, "sort=") {
				queryTrueCounter++
			}
		}
		if queryTrueCounter < 1 {
			panic("invalid query")
		}
	}

	fmt.Println("query: ", query)
}

func GetError(err error) {
	if err != nil {
		panic(err)
	}
}

// uri := r.RequestURI
// page := GetPage(r)
// stringPage := strconv.Itoa(int(page))
// pageReplacer := "/api/gpus?page=" + stringPage + "&"
// uri = strings.Replace(uri, pageReplacer, "", -1)
// filters := strings.Split(uri, "&")
// fmt.Println("filters: ", filters)
// DbSettings.CollectionName = append(DbSettings.CollectionName, GPU_COLLECTION_NAME)

// parameters = append(parameters, parameter)
// kIndex := IndexOf(k, parameters)
// parameters = append(parameters[kIndex], splitString)
