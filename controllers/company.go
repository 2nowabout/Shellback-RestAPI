package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"Shellback.nl/Restapi/database"
	"Shellback.nl/Restapi/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompanyRepo struct {
	Db *gorm.DB
}

func NewCompanyRepo() *CompanyRepo {
	db := database.InitDb()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Company{})
	return &CompanyRepo{Db: db}
}

// create Company
func (repository *CompanyRepo) CreateCompany(c *gin.Context) {
	var Company models.Company
	c.BindJSON(&Company)
	err := models.CreateCompany(repository.Db, &Company)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Company)
}

// get Companys
func (repository *CompanyRepo) GetCompanys(c *gin.Context) {
	var Company []models.Company
	err := models.GetCompanys(repository.Db, &Company)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Company)
}

// get Company by id
func (repository *CompanyRepo) GetCompany(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var Company models.Company
	err := models.GetCompany(repository.Db, &Company, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Company)
}

func (repository *CompanyRepo) UpdateCompanyName(c *gin.Context) {
	var Company models.Company
	c.BindJSON(&Company)
	fmt.Println(&Company)
	err := models.UpdateCompanyName(repository.Db, &Company)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Company)
}

// delete Company
func (repository *CompanyRepo) DeleteCompany(c *gin.Context) {
	var Company models.Company
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteCompany(repository.Db, &Company, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Company deleted successfully"})
}

func (repository *CompanyRepo) UpdateLastActive(c *gin.Context) {
	id := c.Param("ip")
	err := models.UpdateActive(repository.Db, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Active time updated succesfully"})
}
