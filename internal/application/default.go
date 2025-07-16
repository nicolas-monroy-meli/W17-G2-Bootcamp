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

	buyRepo := repo.NewBuyerRepo(db)
	purRepo := repo.NewPurchaseOrderRepo(db)
	empRepo := repo.NewEmployeeRepo(db)
	inbRepo := repo.NewInboundRepo(db)
	secRepo := repo.NewSectionRepo(db)
	pbRepo := repo.NewProductBatchRepo(db)
	prdRepo := repo.NewProductRepo(db)
	prdRcRepo := repo.NewProductRecordRepo(db)
	selRepo := repo.NewSellerRepo(db)
	locRepo := repo.NewLocalityRepo(db)
	wrhRepo := repo.NewWarehouseRepository(db)
	carrRepo := repo.NewCarryRepository(db)

	//instancing service layer
	buyServ := serv.NewBuyerService(buyRepo)
	purServ := serv.NewPurchaseOrderService(purRepo)
  empServ := serv.NewEmployeeService(empRepo)
	inbServ := serv.NewInboundService(inbRepo)
	secServ := serv.NewSectionService(secRepo)
	pbServ := serv.NewProductBatchRepository(pbRepo)
	prdServ := serv.NewProductService(prdRepo)
	prdRcServ := serv.NewProductRecordService(prdRcRepo, prdRepo)
	selServ := serv.NewSellerService(selRepo)
	locServ := serv.NewLocalityService(locRepo)
	wrhServ := serv.NewWarehouseService(wrhRepo)
	carrServ := serv.NewCarryService(carrRepo)

	//instancing handler layer
	buyHand := hand.NewBuyerHandler(buyServ)
	purHand := hand.NewPurchaseOrderHandler(purServ)
	empHand := hand.NewEmployeeHandler(empServ)
	inbHand := hand.NewInboundHandler(inbServ)
  secHand := hand.NewSectionHandler(secServ)
	pbHand := hand.NewProductBatchHandler(pbServ)
	prdHand := hand.NewProductHandler(prdServ)
	prdRcHand := hand.NewProductRecordHandler(prdRcServ)
	selHand := hand.NewSellerHandler(selServ)
	locHand := hand.NewLocalityHandler(locServ)
	wrhHand := hand.NewWarehouseHandler(wrhServ)
	carrHand := hand.NewCarryHandler(carrServ)


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
	rt.Route("/v1/warehouses", func(r chi.Router) {
		r.Get("/", wrhHand.GetAll())
		r.Get("/{id}", wrhHand.GetByID())
		r.Post("/", wrhHand.Create())
		r.Put("/{id}", wrhHand.Update())
		r.Delete("/{id}", wrhHand.Delete())
	})

	/// - Carries
	rt.Route("/v1/carries", func(r chi.Router) {
		r.Post("/", carrHand.Create()) // Crea un nuevo carry
	})

	// - sections

	rt.Route("/v1/sections", func(rt chi.Router) {
		rt.Get("/", secHand.GetAll())
		rt.Get("/{id}", secHand.GetByID())
		rt.Delete("/{id}", secHand.Delete())
		rt.Post("/", secHand.Create())
		rt.Patch("/{id}", secHand.Update())
		rt.Get("/reportProducts", secHand.ReportProducts())
	})

	rt.Route("/v1/productBatches", func(rt chi.Router) {
		rt.Get("/", pbHand.GetAll())
		rt.Post("/", pbHand.Create())
	})

	// - localities
	rt.Route("/v1/localities", func(rt chi.Router) {
		rt.Post("/", locHand.Create())
		rt.Get("/", locHand.GetSelByLoc())
		rt.Get("/reportSellers", locHand.GetSelByLocID())

	})

	// - products
	rt.Route("/v1/products", func(rt chi.Router) {
		rt.Get("/", prdHand.GetAll())
		rt.Get("/{id}", prdHand.GetByID())
		rt.Post("/", prdHand.Create())
		rt.Patch("/{id}", prdHand.Update())
		rt.Delete("/{id}", prdHand.Delete())
		rt.Get("/reportRecords", prdRcHand.GetRecords())
	})

	// - product records
	rt.Route("/v1/productRecords", func(rt chi.Router) {
		rt.Post("/", prdRcHand.CreateRecord())
	})

	//// - employees
	rt.Route("/v1/employees", func(rt chi.Router) {
		rt.Get("/", empHand.GetAll())
		rt.Get("/reportInboundOrders", inbHand.GetOrdersByEmployee())
		rt.Get("/{id}", empHand.GetById())
		rt.Post("/", empHand.Create())
		rt.Patch("/{id}", empHand.Edit())
		rt.Delete("/{id}", empHand.Delete())
	})

	rt.Route("/v1/inboundOrders", func(rt chi.Router) {
		rt.Post("/", inbHand.Create())
	})
	//
	//// - buyers
	rt.Route("/v1/purchaseOrders", func(rt chi.Router) {
		rt.Post("/", purHand.Create())
	})

	rt.Route("/v1/buyers", func(rt chi.Router) {
		rt.Get("/", buyHand.GetAll())
		rt.Get("/{id}", buyHand.GetByID())
		rt.Get("/reportPurchaseOrders", buyHand.GetReport())
		rt.Post("/", buyHand.Create())
		rt.Patch("/{id}", buyHand.Update())
		rt.Delete("/{id}", buyHand.Delete())
	})

	//run
	err = http.ListenAndServe(d.Address, rt)
	if err != nil {
		return err
	}
	return nil
}
