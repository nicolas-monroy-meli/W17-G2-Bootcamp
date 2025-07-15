package server

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
	buyers := docs.ReadFileToMap[mod.Buyer](a.loaderFilePath + "buyers.json")
	//employees := docs.ReadFileToMap[mod.Employee](a.loaderFilePath + "employees.json")
	products := docs.ReadFileToMap[mod.Product](a.loaderFilePath + "products.json")
	//sections := docs.ReadFileToMap[mod.Section](a.loaderFilePath + "sections.json")
	sellers := docs.ReadFileToMap[mod.Seller](a.loaderFilePath + "sellers.json")
	warehouses := docs.ReadFileToMap[mod.Warehouse](a.loaderFilePath + "warehouses.json")

	// - repositories
	buyrp := repo.NewBuyerRepo(buyers)
	//emprp := repo.NewEmployeeRepo(employees)
	prdrp := repo.NewProductRepo(products)
	//secrp := repo.NewSectionRepo(sections)
	selrp := repo.NewSellerRepo(sellers)
	wrhrp := repo.NewWarehouseRepo(warehouses)

	// - services
	buysv := srv.NewBuyerService(buyrp)
	//empsv := srv.NewEmployeeService(emprp)
	prdsv := srv.NewProductService(prdrp)
	//secsv := srv.NewSectionService(secrp)
	selsv := srv.NewSellerService(selrp)
	wrhsv := srv.NewWarehouseService(wrhrp)

	// - handlers
	buyhd := cont.NewBuyerHandler(buysv)
	//emphd := cont.NewEmployeeHandler(empsv)
	prdhd := cont.NewProductHandler(prdsv)
	//sechd := cont.NewSectionHandler(secsv)
	selhd := cont.NewSellerHandler(selsv)
	wrhhd := cont.NewWarehouseHandler(wrhsv)

	// router
	rt := chi.NewRouter()

	// check connection

	rt.Get("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.GoodResponse(w, 200, "pong!", nil)
	}))

	/** - sellers
	rt.Route("/v1/sellers", func(rt chi.Router) {
		rt.Get("/", selhd.GetAll())

		rt.Get("/{id}", selhd.GetByID())
		rt.Post("/", selhd.Create())
		rt.Patch("/{id}", selhd.Update())
		rt.Delete("/{id}", selhd.Delete())

	})**/

	// - warehouses
	rt.Route("/v1/warehouses", func(r chi.Router) {
		r.Get("/", wrhhd.GetAll())
		r.Get("/{id}", wrhhd.GetByID())
		r.Post("/", wrhhd.Create())
		r.Put("/{id}", wrhhd.Update())
		r.Delete("/{id}", wrhhd.Delete())
	})

	// - sections

	/*rt.Route("/v1/sections", func(rt chi.Router) {
		rt.Get("/", sechd.GetAll())
		rt.Get("/{id}", sechd.GetByID())
		rt.Delete("/{id}", sechd.Delete())
		rt.Post("/", sechd.Create())
		rt.Patch("/{id}", sechd.Update())
	})*/

	/** - products
	rt.Route("/v1/products", func(rt chi.Router) {
		rt.Get("/", prdhd.GetAll())
		rt.Get("/{id}", prdhd.GetByID())
		rt.Post("/", prdhd.Create())
		rt.Patch("/{id}", prdhd.Update())
		rt.Delete("/{id}", prdhd.Delete())
	})

	// - employees
	//rt.Route("/v1/employees", func(rt chi.Router) {
	//	rt.Get("/", emphd.GetAll())
	//	rt.Get("/{id}", emphd.GetById())
	//	rt.Post("/", emphd.Create())
	//	rt.Patch("/{id}", emphd.Edit())
	//	rt.Delete("/{id}", emphd.Delete())
	//})

	// - buyers
	rt.Route("/api/v1/buyers", func(rt chi.Router) {
		rt.Get("/", buyhd.GetAll())
		rt.Get("/{id}", buyhd.GetByID())
		rt.Post("/", buyhd.Create())
		rt.Patch("/{id}", buyhd.Update())
		rt.Delete("/{id}", buyhd.Delete())
	})**/

	// run
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
