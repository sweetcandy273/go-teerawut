package handlers

import (
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/sweetcandy273/go-teerawut/pkg/handlers/context"
	"github.com/sweetcandy273/go-teerawut/pkg/handlers/render"
)

// ResponseObject handle response object
func ResponseObject(c *fiber.Ctx, fn interface{}, request interface{}) error {
	ctx := context.New(c)
	err := ctx.BindValue(request, true)
	if err != nil {
		logrus.Errorf("bind value error: %s", err)
		return err
	}

	out := reflect.ValueOf(fn).Call([]reflect.Value{
		reflect.ValueOf(request),
	})
	errObj := out[1].Interface()
	if errObj != nil {
		logrus.Errorf("call service error: %s", errObj)
		return errObj.(error)
	}
	return render.JSON(c, out[0].Interface())
}

// ResponseObjectWithoutRequest handle response object without request
func ResponseObjectWithoutRequest(c *fiber.Ctx, fn interface{}) error {
	out := reflect.ValueOf(fn).Call([]reflect.Value{})
	errObj := out[1].Interface()
	if errObj != nil {
		logrus.Errorf("call service error: %s", errObj)
		return errObj.(error)
	}
	return render.JSON(c, out[0].Interface())
}

// ResponseSuccess handle response success
func ResponseSuccess(c *fiber.Ctx, fn interface{}, request interface{}) error {
	ctx := context.New(c)
	err := ctx.BindValue(request, true)
	if err != nil {
		logrus.Errorf("bind value error: %s", err)
		return err
	}

	out := reflect.ValueOf(fn).Call([]reflect.Value{
		reflect.ValueOf(request),
	})
	errObj := out[0].Interface()
	if errObj != nil {
		logrus.Errorf("call service error: %s", errObj)
		return errObj.(error)
	}
	return render.JSON(c, nil)
}

// ResponseSuccessWithoutRequest handle response success without request
func ResponseSuccessWithoutRequest(c *fiber.Ctx, fn interface{}) error {
	out := reflect.ValueOf(fn).Call([]reflect.Value{})
	errObj := out[0].Interface()
	if errObj != nil {
		logrus.Errorf("call service error: %s", errObj)
		return errObj.(error)
	}
	return render.JSON(c, nil)
}

// ResponseByte handle response object
func ResponseByte(c *fiber.Ctx, fn interface{}, request interface{}) error {
	ctx := context.New(c)
	err := ctx.BindValue(request, true)
	if err != nil {
		logrus.Errorf("bind value error: %s", err)
		return err
	}

	out := reflect.ValueOf(fn).Call([]reflect.Value{
		reflect.ValueOf(request),
	})

	errObj := out[1].Interface()
	if errObj != nil {
		logrus.Errorf("call service error: %s", errObj)
		return errObj.(error)
	}

	bytes, _ := out[0].Interface().([]byte)

	return render.Byte(c, bytes)
}
