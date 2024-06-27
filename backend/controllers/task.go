package controllers

import (
	"context"
	"time"

	"kenshilabs_demo/database"
	"kenshilabs_demo/helper"
	"kenshilabs_demo/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTask(c *fiber.Ctx) error {
	payload := new(model.Task)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Result{
			Status:  false,
			Message: err.Error(),
		})
	}

	taskCollection := database.GetClientDb().Database(database.GetDataBase()).Collection("tasks")

	newUser := &model.Task{
		TaskID:        primitive.NewObjectID(),
		TaskName:      payload.TaskName,
		TaskAssinedBy: payload.TaskAssinedBy,
		TaskAssinedTo: payload.TaskAssinedTo,
		CreatedAt:     time.Now(),
	}

	result, err := taskCollection.InsertOne(context.TODO(), newUser)

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

func GetTask(c *fiber.Ctx) error {
	var results []model.Task
	var taskCollection = database.GetClientDb().Database(database.GetDataBase()).Collection("tasks")

	filter := bson.M{"_id": bson.M{"$ne": ""}}
	cursor, err := taskCollection.Find(context.TODO(), filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Result{
			Status:  false,
			Message: err.Error(),
		})
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Result{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Result{
		Status: true,
		Data:   results,
	})
}

func GetTaskById(c *fiber.Ctx) error {
	taskId := c.Params("id")
	var taskCollection = database.GetClientDb().Database(database.GetDataBase()).Collection("tasks")

	if !helper.CheckDataExistsByID(taskId, taskCollection) {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Result{
			Status:  false,
			Message: "data not found",
		})
	}

	var results model.Task
	formObjectId, formiderr := primitive.ObjectIDFromHex(taskId)
	if formiderr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Result{
			Status:  false,
			Message: formiderr.Error(),
		})
	}

	filter := bson.D{{Key: "_id", Value: formObjectId}}
	err := taskCollection.FindOne(context.TODO(), filter).Decode(&results)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Result{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Result{
		Status: true,
		Data:   results,
	})
}

func DeleteTask(c *fiber.Ctx) error {
	taskId := c.Params("id")
	var taskCollection = database.GetClientDb().Database(database.GetDataBase()).Collection("tasks")

	if !helper.CheckDataExistsByID(taskId, taskCollection) {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Result{
			Status:  false,
			Message: "data not found",
		})
	}

	formObjectId, formiderr := primitive.ObjectIDFromHex(taskId)
	if formiderr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Result{
			Status:  false,
			Message: formiderr.Error(),
		})
	}

	filter := bson.D{{Key: "_id", Value: formObjectId}}
	results, err := taskCollection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Result{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Result{
		Status: true,
		Data:   results,
	})
}

func UpdateTask(c *fiber.Ctx) error {
	taskId := c.Params("id")

	payload := new(model.Task)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Result{
			Status:  false,
			Message: err.Error(),
		})
	}

	var taskCollection = database.GetClientDb().Database(database.GetDataBase()).Collection("tasks")

	if !helper.CheckDataExistsByID(taskId, taskCollection) {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Result{
			Status:  false,
			Message: "data not found",
		})
	}

	formObjectId, formiderr := primitive.ObjectIDFromHex(taskId)
	if formiderr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Result{
			Status:  false,
			Message: formiderr.Error(),
		})
	}

	filter := bson.D{{Key: "_id", Value: formObjectId}}
	update := bson.D{{Key: "$set", Value: bson.M{
		"TaskName":      payload.TaskName,
		"TaskAssinedBy": payload.TaskAssinedBy,
		"TaskAssinedTo": payload.TaskAssinedTo,
		"UpdatedAt":     time.Now(),
	}}}

	results, err := taskCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Result{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Result{
		Status: true,
		Data:   results,
	})
}
