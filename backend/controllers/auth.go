package controllers

import (
	"context"
	"strconv"
	"time"

	"kenshilabs_demo/database"
	"kenshilabs_demo/helper"
	"kenshilabs_demo/jwt"
	"kenshilabs_demo/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *fiber.Ctx) error {
	payload := new(model.User)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Result{
			Status:  false,
			Message: err.Error(),
		})
	}

	usersCollection := database.GetClientDb().Database(database.GetDataBase()).Collection("users")
	if helper.CheckDataExistsByKeyValue("PhoneNumber", payload.PhoneNumber, usersCollection) {
		return c.Status(fiber.StatusBadRequest).JSON(model.Result{
			Status:  false,
			Message: "phone number already exist",
		})
	}

	hash, hasherr := HashPassword(payload.Password)
	if hasherr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Result{
			Status:  false,
			Message: hasherr.Error(),
		})
	}

	newUser := &model.User{
		ID:          primitive.NewObjectID(),
		PhoneNumber: payload.PhoneNumber,
		Password:    hash,
		CreatedAt:   time.Now(),
	}

	result, err := usersCollection.InsertOne(context.TODO(), newUser)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Result{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(model.Result{
		Status: true,
		Data:   result,
	})
}

func SignIn(c *fiber.Ctx) error {
	payload := new(model.User)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Result{
			Status:  false,
			Message: err.Error(),
		})
	}

	usersCollection := database.GetClientDb().Database(database.GetDataBase()).Collection("users")

	filter := bson.M{"PhoneNumber": payload.PhoneNumber}
	var result model.User
	err := usersCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Result{
			Status:  false,
			Message: err.Error(),
		})
	}

	if result.ID.IsZero() {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Result{
			Status:  false,
			Message: "phoneNumber not found",
		})
	}

	checkPass := CheckPasswordHash(payload.Password, result.Password)
	if !checkPass {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Result{
			Status:  false,
			Message: "wrong password",
		})
	}

	token, err := jwt.GenerateJWT(strconv.Itoa(payload.PhoneNumber))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Result{
			Status:  false,
			Message: err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		MaxAge:   10 * 60,
		Secure:   false,
		HTTPOnly: true,
		SameSite: "None", // Adjust SameSite attribute as needed

		// Domain:   ".localhost",
	})

	return c.Status(fiber.StatusOK).JSON(model.Result{
		Status: true,
		Data:   token,
	})
}

func SignOut(c *fiber.Ctx) error {
	expired := time.Now().Add(-time.Hour * 24)
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: expired,
	})

	return c.Status(fiber.StatusOK).JSON(model.Result{
		Status: true,
		Data:   "Sign Out Successful",
	})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
