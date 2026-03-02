package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/te-fa-bene/api-go/internal/app/database"
	"github.com/te-fa-bene/api-go/internal/app/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type seedEmployee struct {
	Name  string
	Email string
	Role  string
}

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("db connect error: %v", err)
	}

	storeName := getenv("SEED_STORE_NAME", "Te Fa Bene")
	rawPassword := getenv("SEED_EMPLOYEE_PASSWORD", "mvptfb")

	store, created, err := ensureStore(db, storeName)
	if err != nil {
		log.Fatalf("ensure store error: %v", err)
	}
	log.Printf("store: %s id=%s (created=%v)\n", store.Name, store.ID, created)

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("bcrypt error: %v", err)
	}
	passwordHash := string(hashBytes)

	employees := []seedEmployee{
		{Name: "Waiter Demo", Email: "waiter@tefabene.dev", Role: "waiter"},
		{Name: "Kitchen Demo", Email: "kitchen@tefabene.dev", Role: "kitchen"},
		{Name: "Cashier Demo", Email: "cashier@tefabene.dev", Role: "cashier"},
		{Name: "Manager Demo", Email: "manager@tefabene.dev", Role: "manager"},
	}

	for _, e := range employees {
		emp, createdEmp, err := ensureEmployee(db, store.ID, e, passwordHash)
		if err != nil {
			log.Fatalf("ensure employee error (%s): %v", e.Email, err)
		}
		log.Printf("employee: %s id=%s role=%s email=%s (created=%v)\n", emp.Name, emp.ID, emp.Role, emp.Email, createdEmp)
	}

	fmt.Println()
	fmt.Println("Seed complete ✅")
	fmt.Printf("Store ID: %s\n", store.ID)
	fmt.Printf("Demo password: %s\n", rawPassword)
}

func ensureStore(db *gorm.DB, name string) (domain.Store, bool, error) {
	var store domain.Store

	err := db.
		Where("name = ? AND deleted_at IS NULL", name).
		First(&store).Error

	if err == nil {
		return store, false, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.Store{}, false, err
	}

	store = domain.Store{Name: name}
	if err := db.Create(&store).Error; err != nil {
		return domain.Store{}, false, err
	}

	return store, true, nil
}

func ensureEmployee(db *gorm.DB, storeID string, e seedEmployee, passwordHash string) (domain.Employee, bool, error) {
	var emp domain.Employee

	err := db.
		Where("store_id = ? AND email = ? AND deleted_at IS NULL", storeID, e.Email).
		First(&emp).Error

	if err == nil {
		return emp, false, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.Employee{}, false, err
	}

	emp = domain.Employee{
		StoreID:      storeID,
		Name:         e.Name,
		Email:        e.Email,
		PasswordHash: passwordHash,
		Role:         e.Role,
		IsActive:     true,
	}

	if err := db.Create(&emp).Error; err != nil {
		return domain.Employee{}, false, err
	}

	return emp, true, nil
}

func getenv(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}
