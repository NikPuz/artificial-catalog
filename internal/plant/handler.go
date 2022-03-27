package plant

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type plantHandler struct {
	plantService PlantService
	logger       *zap.Logger
}

func RegisterPlantHandlers(r *mux.Router, service PlantService, logger *zap.Logger) {
	plantHandler := plantHandler{
		plantService: service,
		logger:       logger,
	}

	r.HandleFunc("/dai", plantHandler.GetPage).Methods("GET")
	r.HandleFunc("/image/{imageName}", plantHandler.GetImage).Methods("GET")
}

func (s plantHandler) GetPage(w http.ResponseWriter, r *http.Request) {
	plants, err := s.plantService.GetPage(r.Context())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		s.logger.Error("GetPage", zap.Error(err))
		return
	}

	req, err := json.Marshal(plants)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		s.logger.Error("GetPage Marshal", zap.Error(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "message: "+string(req))
}

func (s plantHandler) GetImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageNamge := vars["imageName"]

	image, err := s.plantService.GetImage(r.Context(), imageNamge)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		s.logger.Error("GetImage", zap.Error(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(image)
}
