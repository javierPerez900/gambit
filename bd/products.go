package bd

import (
	"database/sql"
	"fmt"

	"strconv"
	// "strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/javier/gambit/models"
	"github.com/javier/gambit/tools"
)

func InsertProduct(p models.Product) (int64,error) {
	fmt.Println("Comienza Registro")

	err := DbConnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	sentencia := "INSERT INTO products (Prod_Title "

	if len(p.ProdDescription)>0 {
		sentencia += ", Prod_Description"
	}
	if p.ProdPrice > 0 {
		sentencia += ", Prod_Price"
	}
	if p.ProdCategId > 0 {
		sentencia += ", Prod_CategoryId"
	}
	if p.ProdStock > 0 {
		sentencia += ", Prod_Stock"
	}
	if len(p.ProdPath) > 0 {
		sentencia += ", Prod_Path"
	}

	sentencia += ") VALUES ('" + tools.EscapeString(p.ProdTitle) + "'"

	if len(p.ProdDescription)>0 {
		sentencia += ", '" + tools.EscapeString(p.ProdDescription) + "'"
	}
	if p.ProdPrice > 0 {	
		sentencia += ", " + strconv.FormatFloat(p.ProdPrice, 'e', -1, 64)
	}
	if p.ProdCategId > 0 {
		sentencia += ", " + strconv.Itoa(p.ProdCategId)
	}
	if p.ProdStock > 0 {
		sentencia += ", " + strconv.Itoa(p.ProdStock)
	}
	if len(p.ProdPath) > 0 {
		sentencia += ", '" + tools.EscapeString(p.ProdPath) + "'"
	}

	sentencia += ")"

	var result sql.Result
	result, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	LastInsertId, err2 := result.LastInsertId()
	if err2 != nil {
		return 0, err2
	}

	fmt.Println("Insert Product > Ejecución Exitosa")
	return LastInsertId, nil
}

func UpdateProduct(p models.Product) error {
	fmt.Println("Comienza UpdateProduct")

	err := DbConnect()
	if err!=nil {
		return err
	}
	defer Db.Close()

	sentencia := "UPDATE products SET "

	sentencia = tools.ArmoSentencia(sentencia, "Prod_Title", "S", 0, 0, p.ProdTitle)
	sentencia = tools.ArmoSentencia(sentencia, "Prod_Description", "S", 0, 0, p.ProdDescription)
	sentencia = tools.ArmoSentencia(sentencia, "Prod_Price", "F", 0, p.ProdPrice, "")
	sentencia = tools.ArmoSentencia(sentencia, "Prod_CategoryId", "N", p.ProdCategId, 0, "")
	sentencia = tools.ArmoSentencia(sentencia, "Prod_Stock", "N", p.ProdStock, 0, "")
	sentencia = tools.ArmoSentencia(sentencia, "Prod_Path", "S", 0, 0, p.ProdPath)

	sentencia += " WHERE Prod_id = " + strconv.Itoa(p.ProdId)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update Product > Ejecución Exitosa")
	return nil
}

func DeleteProduct(id int) error {
	fmt.Println("Comienza Registro de DeleteProduct")

	err := DbConnect()
	if err!=nil {
		return err
	}
	defer Db.Close()

	sentencia := "DELETE FROM products WHERE Prod_Id = " + strconv.Itoa(id)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(sentencia)
	fmt.Println("Delete Product > Ejecución Exitosa")
	return nil
}

// func SelectCategories(CategId int, Slug string) ([]models.Category, error) {
// 	fmt.Println("Comienza SelectCategories")

// 	var Categ []models.Category

// 	err := DbConnect()
// 	if err!=nil {
// 		return Categ, err
// 	}
// 	defer Db.Close()

// 	sentencia := "SELECT Categ_Id, Categ_Name, Categ_Path FROM category"

// 	if CategId > 0 {
// 		sentencia += " WHERE Categ_Id = " + strconv.Itoa(CategId)
// 	} else {
// 		if len(Slug) > 0 {
// 			sentencia += " WHERE Categ_Path LIKE '%" + Slug + "%'"
// 		}
// 	}

// 	fmt.Println(sentencia)

// 	var rows *sql.Rows
// 	rows, err = Db.Query(sentencia)

// 	for rows.Next() {
// 		var c models.Category
// 		var categId sql.NullInt32
// 		var categName sql.NullString
// 		var categPath sql.NullString
	
// 		err := rows.Scan(&categId, &categName, &categPath)
// 		if err != nil {
// 			return Categ, err
// 		}

// 		c.CategID = int(categId.Int32)
// 		c.CategName = categName.String
// 		c.CategPath = categPath.String

// 		Categ = append(Categ, c)
// 	}

// 	fmt.Println("Select Category > Ejecución Exitosa")
// 	return Categ, nil
// }