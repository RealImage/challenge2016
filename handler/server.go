package handler

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	inmemory "github.com/nikhilsiwach28/Cinema-Distribution-System/inMemory"
	svc "github.com/nikhilsiwach28/Cinema-Distribution-System/service"
)

type apiRequest interface {
	Parse(*http.Request) error
}

type apiResponse interface {
	Write(http.ResponseWriter) error
}

// CustomHandler is your custom handler interface that includes the repository as a dependency.
type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// ServerConfig is your server configuration interface.
type ServerConfig interface {
	GetAddress() string
	GetPort() string
}

type APIServer struct {
	Config   ServerConfig
	Router   *mux.Router
	Handlers map[string]Handler
}

func NewServer(config ServerConfig) *APIServer {
	return &APIServer{
		Config:   config,
		Router:   mux.NewRouter(), // Initialize Gorilla Mux router
		Handlers: make(map[string]Handler),
	}
}

func (s *APIServer) run() {
	address := s.Config.GetAddress()
	port := s.Config.GetPort()
	addr := fmt.Sprintf("%s:%s", address, port)

	fmt.Printf("APIServer is running on http://%s\n", addr)

	for route, handler := range s.Handlers {
		s.Router.Handle(route, handler)
	}

	http.Handle("/", s.Router) // Set the Gorilla Mux router as the default handler

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func StartHttpServer(config ServerConfig) {
	server := NewServer(config)
	inmemory := inmemory.NewInMemory()
	server.initRoutes(inmemory)
	server.run()
}

func (s *APIServer) initRoutes(inMemory inmemory.InMemory) {
	// Initialize routes and handlers for different entities
	distributor := NewDistributorHandler(svc.NewDistributorService(inMemory))
	s.Handlers["/distributor"] = distributor

	splitDistributor := NewSplitDistributorHandler(svc.NewSplitDistributorService(inMemory))
	s.Handlers["/split-distribution"] = splitDistributor
}

func initDB() map[string]map[string][]string {
	// Load regions data from CSV
	file, err := os.Open("./cities.csv")
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	// Create maps for country and state
	countryMap := make(map[string]map[string][]string) // Country -> State -> Cities

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read records from the CSV
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV: %s", err)
	}

	// Process each record
	for _, record := range records {
		countryName := record[5]
		stateName := record[4]
		cityName := record[3]

		// Add city to country map
		if _, ok := countryMap[countryName]; !ok {
			countryMap[countryName] = make(map[string][]string)
		}
		if _, ok := countryMap[countryName][stateName]; !ok {
			countryMap[countryName][stateName] = []string{}
		}
		countryMap[countryName][stateName] = append(countryMap[countryName][stateName], cityName)

	}
	return countryMap
}
