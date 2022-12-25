package endpoint

import (
	"github.com/balqisgautama/jubelio-chat/constanta"
	"github.com/balqisgautama/jubelio-chat/dto/response"
	"github.com/balqisgautama/jubelio-chat/util"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"time"
)

type AbstractEndpoint struct {
	FileName string
	FuncName string
}

func (input AbstractEndpoint) ServeEndpoint(serveFunction func(*fiber.Request) (response.APIResponse, map[string]string, error), c *fiber.Ctx) {
	serveEndpoint(serveFunction, c.Response(), c.Request())
}

func serveEndpoint(serveFunction func(*fiber.Request) (response.APIResponse, map[string]string, error), responseWriter *fiber.Response, request *fiber.Request) {
	serve(serveFunction, responseWriter, request)
}

func serve(serveFunction func(*fiber.Request) (response.APIResponse, map[string]string, error), responseWriter *fiber.Response, request *fiber.Request) {
	var err error
	var output response.APIResponse
	var header map[string]string

	defer func() {
		if r := recover(); r != nil {
			util.Logger.Info("recovery")
		} else {
			if err != nil {
				util.Logger.Info("server", zap.String("details", err.Error()))
			}
		}

		finish(responseWriter, err, output)
	}()

	output, header, err = serveFunction(request)

	setHeader(header, responseWriter)
}

func setHeader(header map[string]string, responseWriter *fiber.Response) {
	accessControlExpose := "Access-Control-Expose-Headers"
	exposeHeader := string(responseWriter.Header.Peek(accessControlExpose))
	for key := range header {
		responseWriter.Header.Add(key, header[key])
		if exposeHeader == "" {
			exposeHeader = key
		} else {
			exposeHeader += ", " + key
		}
	}
	if exposeHeader != "" {
		responseWriter.Header.Set(accessControlExpose, exposeHeader)
	}
}

func finish(responseWriter *fiber.Response, err error, output response.APIResponse) {
	if err != nil {
		writeErrorResponse(responseWriter, err)
	} else {
		writeSuccessResponse(responseWriter, output)
	}
}

func writeErrorResponse(responseWriter *fiber.Response, err error) {
	responseWriter.SetStatusCode(500)
	responseWriter.SetBody([]byte(err.Error()))
}

func writeSuccessResponse(responseWriter *fiber.Response, output response.APIResponse) {
	output.Timestamp = time.Now().Format(constanta.DefaultDateFormat)
	bodyMessage := output.String()

	responseWriter.SetStatusCode(200)
	responseWriter.SetBody([]byte(bodyMessage))
}
