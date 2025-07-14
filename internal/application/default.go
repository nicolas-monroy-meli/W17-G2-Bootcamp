package server

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-sql-driver/mysql"
	hand "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/handler"
	repo "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/repository"
	serv "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/service"
)

type SQLConfig struct {
	Database mysql.Config
	Address  string
}

func NewSQLConfig(cfg *SQLConfig) *SQLConfig {
	cfgDefault := &SQLConfig{
		Address: ":8080",
	}
	if cfg != nil {
		cfgDefault.Database = cfg.Database
		if cfg.Address != "" {
			cfgDefault.Address = cfg.Address
		}
	}
	return &SQLConfig{
		Database: cfgDefault.Database,
		Address:  cfgDefault.Address,
	}
}

func (d *SQLConfig) Run() (err error) {
	//open database connection
	db, err := sql.Open("mysql", d.Database.FormatDSN())
	if err != nil {
		return
	}
	defer db.Close()
	//test db connection
	err = db.Ping()
	if err != nil {
		return
	}
	// instancing repository layer
	//buyRepo := repo.NewBuyerRepo(db)
	//empRepo := repo.NewEmployeeRepo(db)
	//prdRepo := repo.NewProductRepo(db)
	secRepo := repo.NewSectionRepo(db)
	selRepo := repo.NewSellerRepo(db)
	locRepo := repo.NewLocalityRepo(db)
	//wrhRepo := repo.NewWarehouseRepo(db)

	//instancing service layer
	//buyServ := serv.NewBuyerService(buyRepo)
	//empServ := serv.NewEmployeeService(empRepo)
	//prdServ := serv.NewProductService(prdRepo)
	secServ := serv.NewSectionService(secRepo)
	selServ := serv.NewSellerService(selRepo)
	locServ := serv.NewLocalityService(locRepo)
	//wrhServ := serv.NewWarehouseService(wrhRepo)

	//instancing handler layer
	//buyHand := hand.NewBuyerHandler(buyServ)
	//empHand := hand.NewEmployeeHandler(empServ)
	//prdHand := hand.NewProductHandler(prdServ)
	secHand := hand.NewSectionHandler(secServ)
	selHand := hand.NewSellerHandler(selServ)
	locHand := hand.NewLocalityHandler(locServ)
	//wrhHand := hand.NewWarehouseHandler(wrhServ)

	//routing

	rt := chi.NewRouter()
	//middlewares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	//Routing
	// - sellers
	rt.Route("/v1/sellers", func(rt chi.Router) {
		rt.Get("/", selHand.GetAll())
		rt.Get("/{id}", selHand.GetByID())
		rt.Post("/", selHand.Create())
		rt.Patch("/{id}", selHand.Update())
		rt.Delete("/{id}", selHand.Delete())
	})
	//
	//// - warehouses
	//rt.Route("/v1/warehouses", func(r chi.Router) {
	//	r.Get("/", wrhHand.GetAll())
	//	r.Get("/{id}", wrhHand.GetByID())
	//	r.Post("/", wrhHand.Create())
	//	r.Put("/{id}", wrhHand.Update())
	//	r.Delete("/{id}", wrhHand.Delete())
	//})

	// - sections
	rt.Route("/v1/sections", func(rt chi.Router) {
		rt.Get("/", secHand.GetAll())
		rt.Get("/{id}", secHand.GetByID())
		rt.Delete("/{id}", secHand.Delete())
		rt.Post("/", secHand.Create())
		rt.Patch("/{id}", secHand.Update())
	})

	// - localities
	rt.Route("/v1/localities", func(rt chi.Router) {
		//rt.Post("/", secHand.GetAll())
		rt.Get("/reportSellers", locHand.GetSelByLoc())
		rt.Get("/reportSeller", locHand.GetSelByLocID())

	})

	//// - products
	//rt.Route("/v1/products", func(rt chi.Router) {
	//	rt.Get("/", prdHand.GetAll())
	//	rt.Get("/{id}", prdHand.GetByID())
	//	rt.Post("/", prdHand.Create())
	//	rt.Patch("/{id}", prdHand.Update())
	//	rt.Delete("/{id}", prdHand.Delete())
	//})
	//
	//// - employees
	//rt.Route("/v1/employees", func(rt chi.Router) {
	//	rt.Get("/", empHand.GetAllEmployees())
	//	rt.Get("/{id}", empHand.GetEmployeeById())
	//	rt.Post("/", empHand.CreateEmployee())
	//	rt.Patch("/{id}", empHand.EditEmployee())
	//	rt.Delete("/{id}", empHand.DeleteEmployee())
	//})
	//
	//// - buyers
	//rt.Route("/v1/buyers", func(rt chi.Router) {
	//	rt.Get("/", buyHand.GetAll())
	//	rt.Get("/{id}", buyHand.GetByID())
	//	rt.Post("/", buyHand.Create())
	//	rt.Patch("/{id}", buyHand.Update())
	//	rt.Delete("/{id}", buyHand.Delete())
	//})

	//run
	err = http.ListenAndServe(d.Address, rt)
	if err != nil {
		return err
	}
	return nil
}
