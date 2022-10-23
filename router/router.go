package router

import (
	"fmt"
	"kpt_api/controller"

	"github.com/gorilla/mux"
)

func Router(port string) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/gpus", controller.GetGpuData).Methods("GET")
	router.HandleFunc("/api/gpuFilter", controller.GetFilterData).Methods("GET")
	router.HandleFunc("/api/deleteGpuCollectionData", controller.DeleteGpuData).Methods("DELETE")
	router.HandleFunc("/api/createGpuCollectionData", controller.CreateGpuData).Methods("POST")
	//router.HandleFunc("/gpu/{id}", controller.DeleteAGpu).Methods("DELETE")

	fmt.Println("listening on port " + port + "")

	return router
}
