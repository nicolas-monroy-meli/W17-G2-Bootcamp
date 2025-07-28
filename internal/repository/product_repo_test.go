package repository_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	repository "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/repository"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	prodData "github.com/smartineztri_meli/W17-G2-Bootcamp/tests/data"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ProductRepoTestSuite struct {
	prodData.TestSuite
	repo *repository.ProductDB
}

// FindAll returns all products from the database - TESTED
func (suite *ProductRepoTestSuite) TestProducts_FindAll() {
	t := suite.T()
	t.Run("#1 - All Success", func(t *testing.T) {
		// given
		suite.SetupTest("products")
		suite.MockDb.ExpectQuery("SELECT `id`, `product_code`, `description`, `height`, `length`, `width`, `net_weight`, `expiration_rate`, `freezing_rate`, `recommended_freezing_temperature`, `product_type_id`, `seller_id` FROM products;").
			WillReturnRows(suite.TestTable)
		suite.repo = repository.NewProductRepo(suite.TestDb)

		// when
		products, err := suite.repo.FindAll()

		// then
		expected := []mod.Product{
			{ID: 1, ProductCode: "P001", Description: "Product 1", Height: 10.0, Length: 20.0, Width: 5.0, Weight: 2.0, ExpirationRate: 0.1, FreezingRate: 0.05, RecomFreezTemp: -18.0, ProductTypeID: 1, SellerID: 101},
			{ID: 2, ProductCode: "P002", Description: "Product 2", Height: 15.0, Length: 25.0, Width: 7.0, Weight: 3.0, ExpirationRate: 0.2, FreezingRate: 0.06, RecomFreezTemp: -20.0, ProductTypeID: 2, SellerID: 102},
			{ID: 3, ProductCode: "P003", Description: "Product 3", Height: 12.0, Length: 22.0, Width: 6.0, Weight: 2.5, ExpirationRate: 0.15, FreezingRate: 0.07, RecomFreezTemp: -19.0, ProductTypeID: 1, SellerID: 103},
		}
		require.NoError(t, err)
		require.Len(t, products, len(expected))
		require.Equal(t, expected, products)
	})

	t.Run("#2 - Unable to parse DB info", func(t *testing.T) {
		// given
		suite.SetupTest("products")
		suite.MockDb.ExpectQuery("SELECT `id`, `product_code`, `description`, `height`, `length`, `width`, `net_weight`, `expiration_rate`, `freezing_rate`, `recommended_freezing_temperature`, `product_type_id`, `seller_id` FROM products;").
			WillReturnRows(suite.TestTable.AddRow(1, "P001", "Product 1", 10.0, 20.0, 5.0, 2.0, 0.1, 0.05, -18.0, 1, nil))
		suite.repo = repository.NewProductRepo(suite.TestDb)

		// when
		products, err := suite.repo.FindAll()

		// then
		require.Error(t, err)
		require.Nil(t, products)
	})

	t.Run("#3 - All Query is malformed", func(t *testing.T) {
		// given
		suite.SetupTest("products")
		suite.MockDb.ExpectQuery("SELECT `id`, `product_code`, `description`, `height`, `length`, `width`, `net_weight`, `expiration_rate`, `freezing_rate`, `recommended_freezing_temperature`, `product_type_id`, `seller_id` FROM products;").
			WillReturnError(e.ErrQueryError)
		suite.repo = repository.NewProductRepo(suite.TestDb)

		// when
		products, err := suite.repo.FindAll()

		// then
		require.Error(t, err)
		require.Nil(t, products)
	})

	t.Run("#4 - rows.Err() returns error after iteration", func(t *testing.T) {
		// given
		suite.SetupTest("products")
		mockRows := sqlmock.NewRows([]string{
			"id", "product_code", "description", "height", "length", "width", "net_weight",
			"expiration_rate", "freezing_rate", "recommended_freezing_temperature",
			"product_type_id", "seller_id",
		}).AddRow(1, "P001", "Product 1", 10.0, 20.0, 5.0, 2.0, 0.1, 0.05, -18.0, 1, 101)
		// Simula error en rows.Err()
		mockRows.RowError(0, e.ErrQueryError)
		suite.MockDb.ExpectQuery("SELECT `id`, `product_code`, `description`, `height`, `length`, `width`, `net_weight`, `expiration_rate`, `freezing_rate`, `recommended_freezing_temperature`, `product_type_id`, `seller_id` FROM products;").
			WillReturnRows(mockRows)
		suite.repo = repository.NewProductRepo(suite.TestDb)

		// when
		products, err := suite.repo.FindAll()

		// then
		require.Error(t, err)
		require.Nil(t, products)
	})

	t.Run("#5 - Query returns zero products", func(t *testing.T) {
		// given
		suite.SetupTest("products")
		emptyRows := sqlmock.NewRows([]string{
			"id", "product_code", "description", "height", "length", "width", "net_weight",
			"expiration_rate", "freezing_rate", "recommended_freezing_temperature",
			"product_type_id", "seller_id",
		})
		suite.MockDb.ExpectQuery("SELECT `id`, `product_code`, `description`, `height`, `length`, `width`, `net_weight`, `expiration_rate`, `freezing_rate`, `recommended_freezing_temperature`, `product_type_id`, `seller_id` FROM products;").
			WillReturnRows(emptyRows)
		suite.repo = repository.NewProductRepo(suite.TestDb)

		// when
		products, err := suite.repo.FindAll()

		// then
		require.Error(t, err)
		require.Nil(t, products)
	})
}

// FindByID returns a product from the database by its id
func (suite *ProductRepoTestSuite) TestProducts_FindByID() {
	t := suite.T()
	t.Run("#1 - Producto encontrado", func(t *testing.T) {
		// given
		suite.SetupTest("products")
		expected := mod.Product{ID: 1, ProductCode: "P001", Description: "Product 1", Height: 10.0, Length: 20.0, Width: 5.0, Weight: 2.0, ExpirationRate: 0.1, FreezingRate: 0.05, RecomFreezTemp: -18.0, ProductTypeID: 1, SellerID: 101}
		rows := sqlmock.NewRows([]string{
			"id", "product_code", "description", "height", "length", "width", "net_weight",
			"expiration_rate", "freezing_rate", "recommended_freezing_temperature",
			"product_type_id", "seller_id",
		}).AddRow(1, "P001", "Product 1", 10.0, 20.0, 5.0, 2.0, 0.1, 0.05, -18.0, 1, 101)
		suite.MockDb.ExpectQuery("SELECT `id`, `product_code`, `description`, `height`, `length`, `width`, `net_weight`, `expiration_rate`, `freezing_rate`, `recommended_freezing_temperature`, `product_type_id`, `seller_id` FROM frescos_db.products WHERE id = \\?;").
			WithArgs(1).
			WillReturnRows(rows)
		suite.repo = repository.NewProductRepo(suite.TestDb)

		// when
		product, err := suite.repo.FindByID(expected.ID)

		// then
		require.NoError(t, err)
		require.Equal(t, expected, product)
	})

	t.Run("#2 - Producto no existe", func(t *testing.T) {
		// given
		suite.SetupTest("products")
		rows := sqlmock.NewRows([]string{
			"id", "product_code", "description", "height", "length", "width", "net_weight",
			"expiration_rate", "freezing_rate", "recommended_freezing_temperature",
			"product_type_id", "seller_id",
		}) // sin filas
		suite.MockDb.ExpectQuery("SELECT `id`, `product_code`, `description`, `height`, `length`, `width`, `net_weight`, `expiration_rate`, `freezing_rate`, `recommended_freezing_temperature`, `product_type_id`, `seller_id` FROM frescos_db.products WHERE id = ?;").
			WithArgs(999).
			WillReturnRows(rows)
		suite.repo = repository.NewProductRepo(suite.TestDb)

		// when
		product, err := suite.repo.FindByID(999)

		// then
		require.Error(t, err)
		require.Equal(t, mod.Product{}, product)
	})

	t.Run("#3 - Error en el scan", func(t *testing.T) {
		// given
		suite.SetupTest("products")
		rows := sqlmock.NewRows([]string{
			"id", "product_code", "description", "height", "length", "width", "net_weight",
			"expiration_rate", "freezing_rate", "recommended_freezing_temperature",
			"product_type_id", "seller_id",
		}).AddRow(nil, "P001", "Product 1", 10.0, 20.0, 5.0, 2.0, 0.1, 0.05, -18.0, 1, 101) // id nil
		suite.MockDb.ExpectQuery("SELECT `id`, `product_code`, `description`, `height`, `length`, `width`, `net_weight`, `expiration_rate`, `freezing_rate`, `recommended_freezing_temperature`, `product_type_id`, `seller_id` FROM frescos_db.products WHERE id = ?;").
			WithArgs(1).
			WillReturnRows(rows)
		suite.repo = repository.NewProductRepo(suite.TestDb)

		// when
		product, err := suite.repo.FindByID(1)

		// then
		require.Error(t, err)
		require.Equal(t, mod.Product{}, product)
	})

	t.Run("#4 - Producto con ID 0", func(t *testing.T) {
		// given
		suite.SetupTest("products")
		rows := sqlmock.NewRows([]string{
			"id", "product_code", "description", "height", "length", "width", "net_weight",
			"expiration_rate", "freezing_rate", "recommended_freezing_temperature",
			"product_type_id", "seller_id",
		}).AddRow(0, "P001", "Product 1", 10.0, 20.0, 5.0, 2.0, 0.1, 0.05, -18.0, 1, 101)
		suite.MockDb.ExpectQuery("SELECT `id`, `product_code`, `description`, `height`, `length`, `width`, `net_weight`, `expiration_rate`, `freezing_rate`, `recommended_freezing_temperature`, `product_type_id`, `seller_id` FROM frescos_db.products WHERE id = \\?;").
			WithArgs(123). // El valor puede ser cualquiera, pero debe coincidir con el argumento del método
			WillReturnRows(rows)
		suite.repo = repository.NewProductRepo(suite.TestDb)

		// when
		product, err := suite.repo.FindByID(123)

		// then
		require.Error(t, err)
		require.Equal(t, mod.Product{}, product)
	})
}

// Save saves a product into the database - TESTED
func (suite *ProductRepoTestSuite) TestProducts_Save() {
	t := suite.T()
	t.Run("#1 - Producto duplicado", func(t *testing.T) {
		// given
		suite.SetupTest("products")
		product := &mod.Product{ID: 1, ProductCode: "P001", Description: "Product 1", Height: 10.0, Length: 20.0, Width: 5.0, Weight: 2.0, ExpirationRate: 0.1, FreezingRate: 0.05, RecomFreezTemp: -18.0, ProductTypeID: 1, SellerID: 101}
		// Simula que FindByID retorna producto existente
		suite.repo = repository.NewProductRepo(suite.TestDb)
		suite.MockDb.ExpectQuery("SELECT `id`, `product_code`, `description`, `height`, `length`, `width`, `net_weight`, `expiration_rate`, `freezing_rate`, `recommended_freezing_temperature`, `product_type_id`, `seller_id` FROM frescos_db.products WHERE id = \\?;").
			WithArgs(product.ID).
			WillReturnRows(sqlmock.NewRows([]string{
				"id", "product_code", "description", "height", "length", "width", "net_weight",
				"expiration_rate", "freezing_rate", "recommended_freezing_temperature",
				"product_type_id", "seller_id",
			}).AddRow(product.ID, product.ProductCode, product.Description, product.Height, product.Length, product.Width, product.Weight, product.ExpirationRate, product.FreezingRate, product.RecomFreezTemp, product.ProductTypeID, product.SellerID))

		// when
		err := suite.repo.Save(product)

		// then
		require.ErrorIs(t, err, e.ErrProductRepositoryDuplicated)
	})

	t.Run("#2 - Error de clave foránea", func(t *testing.T) {
		// given
		suite.SetupTest("products")
		product := &mod.Product{ID: 999, ProductCode: "P999"}
		// Simula que FindByID retorna not found
		suite.MockDb.ExpectQuery("SELECT `id`, `product_code`, `description`, `height`, `length`, `width`, `net_weight`, `expiration_rate`, `freezing_rate`, `recommended_freezing_temperature`, `product_type_id`, `seller_id` FROM frescos_db.products WHERE id = \\?;").
			WithArgs(product.ID).
			WillReturnRows(sqlmock.NewRows([]string{"id"}))
		// Simula error de clave foránea en el insert
		suite.MockDb.ExpectExec("INSERT INTO frescos_db.products").
			WithArgs(product.ProductCode, product.Description, product.Height, product.Length, product.Width, product.Weight, product.ExpirationRate, product.FreezingRate, product.RecomFreezTemp, product.ProductTypeID, product.SellerID).
			WillReturnError(fmt.Errorf("foreign key constraint fails: products_ibfk_1"))
		suite.repo = repository.NewProductRepo(suite.TestDb)

		// when
		err := suite.repo.Save(product)

		// then
		require.ErrorIs(t, err, e.ErrSellerRepositoryNotFound)
	})

	t.Run("#3 - Error genérico al insertar", func(t *testing.T) {
		// given
		suite.SetupTest("products")
		product := &mod.Product{ID: 999, ProductCode: "P999"}
		suite.MockDb.ExpectQuery("SELECT `id`, `product_code`, `description`, `height`, `length`, `width`, `net_weight`, `expiration_rate`, `freezing_rate`, `recommended_freezing_temperature`, `product_type_id`, `seller_id` FROM frescos_db.products WHERE id = \\?;").
			WithArgs(product.ID).
			WillReturnRows(sqlmock.NewRows([]string{"id"}))
		suite.MockDb.ExpectExec("INSERT INTO frescos_db.products").
			WithArgs(product.ProductCode, product.Description, product.Height, product.Length, product.Width, product.Weight, product.ExpirationRate, product.FreezingRate, product.RecomFreezTemp, product.ProductTypeID, product.SellerID).
			WillReturnError(fmt.Errorf("error genérico"))
		suite.repo = repository.NewProductRepo(suite.TestDb)

		// when
		err := suite.repo.Save(product)

		// then
		require.ErrorContains(t, err, "error genérico")
	})

	t.Run("#4 - Error en LastInsertId", func(t *testing.T) {
		// given
		suite.SetupTest("products")
		product := &mod.Product{ID: 999, ProductCode: "P999"}
		suite.MockDb.ExpectQuery("SELECT `id`, `product_code`, `description`, `height`, `length`, `width`, `net_weight`, `expiration_rate`, `freezing_rate`, `recommended_freezing_temperature`, `product_type_id`, `seller_id` FROM frescos_db.products WHERE id = \\?;").
			WithArgs(product.ID).
			WillReturnRows(sqlmock.NewRows([]string{"id"}))
		// Simula error en LastInsertId
		result := sqlmock.NewErrorResult(fmt.Errorf("error lastinsertid"))
		suite.MockDb.ExpectExec("INSERT INTO frescos_db.products").
			WithArgs(product.ProductCode, product.Description, product.Height, product.Length, product.Width, product.Weight, product.ExpirationRate, product.FreezingRate, product.RecomFreezTemp, product.ProductTypeID, product.SellerID).
			WillReturnResult(result)
		suite.repo = repository.NewProductRepo(suite.TestDb)

		// when
		err := suite.repo.Save(product)

		// then
		require.ErrorContains(t, err, "error lastinsertid")
	})

	t.Run("#5 - Inserción exitosa", func(t *testing.T) {
		// given
		suite.SetupTest("products")
		product := &mod.Product{ID: 999, ProductCode: "P999"}
		suite.MockDb.ExpectQuery("SELECT `id`, `product_code`, `description`, `height`, `length`, `width`, `net_weight`, `expiration_rate`, `freezing_rate`, `recommended_freezing_temperature`, `product_type_id`, `seller_id` FROM frescos_db.products WHERE id = \\?;").
			WithArgs(product.ID).
			WillReturnRows(sqlmock.NewRows([]string{"id"}))
		// Simula insert exitoso
		suite.MockDb.ExpectExec("INSERT INTO frescos_db.products").
			WithArgs(product.ProductCode, product.Description, product.Height, product.Length, product.Width, product.Weight, product.ExpirationRate, product.FreezingRate, product.RecomFreezTemp, product.ProductTypeID, product.SellerID).
			WillReturnResult(sqlmock.NewResult(123, 1))
		suite.repo = repository.NewProductRepo(suite.TestDb)

		// when
		err := suite.repo.Save(product)

		// then
		require.NoError(t, err)
		require.Equal(t, 123, product.ID)
	})
}

// Update updates a product in the database - TESTED
func (suite *ProductRepoTestSuite) TestProducts_Update() {
	t := suite.T()
	t.Run("#1 - Actualización exitosa", func(t *testing.T) {
		// given
		suite.SetupTest("products")
		product := &mod.Product{ID: 1, ProductCode: "P001", Description: "Product 1", Height: 10.0, Length: 20.0, Width: 5.0, Weight: 2.0, ExpirationRate: 0.1, FreezingRate: 0.05, RecomFreezTemp: -18.0, ProductTypeID: 1, SellerID: 101}
		suite.MockDb.ExpectExec(regexp.QuoteMeta("UPDATE frescos_db.products SET `product_code` = ?, `description` = ?, `height` = ?, `length` = ?, `width` = ?, `net_weight` = ?, `expiration_rate` = ?, `freezing_rate` = ?, `recommended_freezing_temperature` = ?, `product_type_id` = ?, `seller_id` = ? WHERE id = ?;")).
			WithArgs(product.ProductCode, product.Description, product.Height, product.Length, product.Width, product.Weight, product.ExpirationRate, product.FreezingRate, product.RecomFreezTemp, product.ProductTypeID, product.SellerID, product.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		suite.repo = repository.NewProductRepo(suite.TestDb)

		// when
		err := suite.repo.Update(product)

		// then
		require.NoError(t, err)
	})

	t.Run("#2 - Error de clave foránea", func(t *testing.T) {
		// given
		suite.SetupTest("products")
		product := &mod.Product{ID: 2, ProductCode: "P002"}
		suite.MockDb.ExpectExec(regexp.QuoteMeta("UPDATE frescos_db.products SET `product_code` = ?, `description` = ?, `height` = ?, `length` = ?, `width` = ?, `net_weight` = ?, `expiration_rate` = ?, `freezing_rate` = ?, `recommended_freezing_temperature` = ?, `product_type_id` = ?, `seller_id` = ? WHERE id = ?;")).
			WithArgs(product.ProductCode, product.Description, product.Height, product.Length, product.Width, product.Weight, product.ExpirationRate, product.FreezingRate, product.RecomFreezTemp, product.ProductTypeID, product.SellerID, product.ID).
			WillReturnError(fmt.Errorf("foreign key constraint fails: products_ibfk_1"))
		suite.repo = repository.NewProductRepo(suite.TestDb)

		// when
		err := suite.repo.Update(product)

		// then
		require.ErrorIs(t, err, e.ErrSellerRepositoryNotFound)
	})

	t.Run("#3 - Error genérico en el update", func(t *testing.T) {
		// given
		suite.SetupTest("products")
		product := &mod.Product{ID: 3, ProductCode: "P003"}
		suite.MockDb.ExpectExec(regexp.QuoteMeta("UPDATE frescos_db.products SET `product_code` = ?, `description` = ?, `height` = ?, `length` = ?, `width` = ?, `net_weight` = ?, `expiration_rate` = ?, `freezing_rate` = ?, `recommended_freezing_temperature` = ?, `product_type_id` = ?, `seller_id` = ? WHERE id = ?;")).
			WithArgs(product.ProductCode, product.Description, product.Height, product.Length, product.Width, product.Weight, product.ExpirationRate, product.FreezingRate, product.RecomFreezTemp, product.ProductTypeID, product.SellerID, product.ID).
			WillReturnError(fmt.Errorf("error genérico"))
		suite.repo = repository.NewProductRepo(suite.TestDb)

		// when
		err := suite.repo.Update(product)

		// then
		require.ErrorContains(t, err, "error genérico")
	})
}

// Delete deletes a product from the database - TESTED
func (suite *ProductRepoTestSuite) TestProducts_Delete() {
	t := suite.T()
	t.Run("#1 - Producto no existe", func(t *testing.T) {
		// given
		suite.SetupTest("products")
		id := 999
		// Simula que FindByID retorna error
		suite.MockDb.ExpectQuery("SELECT `id`, `product_code`, `description`, `height`, `length`, `width`, `net_weight`, `expiration_rate`, `freezing_rate`, `recommended_freezing_temperature`, `product_type_id`, `seller_id` FROM frescos_db.products WHERE id = \\?;").
			WithArgs(id).
			WillReturnRows(sqlmock.NewRows([]string{"id"}))
		suite.repo = repository.NewProductRepo(suite.TestDb)

		// when
		err := suite.repo.Delete(id)

		// then
		require.ErrorIs(t, err, e.ErrProductRepositoryNotFound)
	})

	t.Run("#2 - Error al ejecutar el delete", func(t *testing.T) {
		// given
		suite.SetupTest("products")
		id := 1
		// Simula que FindByID retorna producto existente
		suite.MockDb.ExpectQuery("SELECT `id`, `product_code`, `description`, `height`, `length`, `width`, `net_weight`, `expiration_rate`, `freezing_rate`, `recommended_freezing_temperature`, `product_type_id`, `seller_id` FROM frescos_db.products WHERE id = \\?;").
			WithArgs(id).
			WillReturnRows(sqlmock.NewRows([]string{
				"id", "product_code", "description", "height", "length", "width", "net_weight",
				"expiration_rate", "freezing_rate", "recommended_freezing_temperature",
				"product_type_id", "seller_id",
			}).AddRow(id, "P001", "Product 1", 10.0, 20.0, 5.0, 2.0, 0.1, 0.05, -18.0, 1, 101))
		// Simula error en el delete
		suite.MockDb.ExpectExec(regexp.QuoteMeta("DELETE FROM frescos_db.products WHERE id = ?;")).
			WithArgs(id).
			WillReturnError(fmt.Errorf("error en delete"))
		suite.repo = repository.NewProductRepo(suite.TestDb)

		// when
		err := suite.repo.Delete(id)

		// then
		require.ErrorContains(t, err, "error en delete")
	})

	t.Run("#3 - Eliminación exitosa", func(t *testing.T) {
		// given
		suite.SetupTest("products")
		id := 1
		// Simula que FindByID retorna producto existente
		suite.MockDb.ExpectQuery("SELECT `id`, `product_code`, `description`, `height`, `length`, `width`, `net_weight`, `expiration_rate`, `freezing_rate`, `recommended_freezing_temperature`, `product_type_id`, `seller_id` FROM frescos_db.products WHERE id = \\?;").
			WithArgs(id).
			WillReturnRows(sqlmock.NewRows([]string{
				"id", "product_code", "description", "height", "length", "width", "net_weight",
				"expiration_rate", "freezing_rate", "recommended_freezing_temperature",
				"product_type_id", "seller_id",
			}).AddRow(id, "P001", "Product 1", 10.0, 20.0, 5.0, 2.0, 0.1, 0.05, -18.0, 1, 101))
		// Simula delete exitoso
		suite.MockDb.ExpectExec(regexp.QuoteMeta("DELETE FROM frescos_db.products WHERE id = ?;")).
			WithArgs(id).
			WillReturnResult(sqlmock.NewResult(0, 1))
		suite.repo = repository.NewProductRepo(suite.TestDb)

		// when
		err := suite.repo.Delete(id)

		// then
		require.NoError(t, err)
	})
}

func TestProductRepoTestSuite(t *testing.T) {
	suite.Run(t, new(ProductRepoTestSuite))
}
