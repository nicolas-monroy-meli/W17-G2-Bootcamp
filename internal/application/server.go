package server

/*
import (
	"net/http"

	"github.com/smartineztri_meli/W17-G2-Bootcamp/docs"
	cont "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/handler"
	repo "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/repository"
	srv "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/service"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils"

	"github.com/go-chi/chi/v5"
)

// ConfigServerChi is a struct that represents the configuration for ServerChi
type ConfigServerChi struct {
	// ServerAddress is the address where the server will be listening
	ServerAddress string
	// LoaderFilePath is the path to the file that contains the vehicles
	LoaderFilePath string
}

// NewServerChi is a function that returns a new instance of ServerChi
func NewServerChi(cfg *ConfigServerChi) *ServerChi {
	// default values
	defaultConfig := &ConfigServerChi{
		ServerAddress: ":8080",
	}
	if cfg != nil {
		if cfg.ServerAddress != "" {
			defaultConfig.ServerAddress = cfg.ServerAddress
		}
		if cfg.LoaderFilePath != "" {
			defaultConfig.LoaderFilePath = cfg.LoaderFilePath
		}
	}

	return &ServerChi{
		serverAddress:  defaultConfig.ServerAddress,
		loaderFilePath: defaultConfig.LoaderFilePath,
	}
}

// ServerChi is a struct that implements the Application interface
type ServerChi struct {
	// serverAddress is the address where the server will be listening
	serverAddress string
	// loaderFilePath is the path to the file that contains the vehicles
	loaderFilePath string
}

// Run is a method that runs the server
func (a *ServerChi) Run() (err error) {
	// dependencies

	// - database loaders
	//sections := docs.ReadFileToMap[mod.Section](a.loaderFilePath + "sections.json")
	sellers := docs.ReadFileToMap[mod.Seller](a.loaderFilePath + "sellers.json")

	// - repositories
	//secrp := repo.NewSectionRepo(sections)
	//selrp := repo.NewSellerRepo(sellers)

	// - services
	//secsv := srv.NewSectionService(secrp)
	//selsv := srv.NewSellerService(selrp)

	// - handlers
	//sechd := cont.NewSectionHandler(secsv)
	//selhd := cont.NewSellerHandler(selsv)

	// router
	rt := chi.NewRouter()

	// check connection

	rt.Get("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.GoodResponse(w, 200, "pong!", nil)
	}))

	// - sellers
	rt.Route("/v1/sellers", func(rt chi.Router) {
		rt.Get("/", selhd.GetAll())

		rt.Get("/{id}", selhd.GetByID())
		rt.Post("/", selhd.Create())
		rt.Patch("/{id}", selhd.Update())
		rt.Delete("/{id}", selhd.Delete())

	})

	// - sections

	/*rt.Route("/v1/sections", func(rt chi.Router) {
		rt.Get("/", sechd.GetAll())
		rt.Get("/{id}", sechd.GetByID())
		rt.Delete("/{id}", sechd.Delete())
		rt.Post("/", sechd.Create())
		rt.Patch("/{id}", sechd.Update())
	})

	// run
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
*/
