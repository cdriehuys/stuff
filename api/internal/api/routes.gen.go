//go:build go1.22

// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/oapi-codegen/runtime"
	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
)

// APIError defines model for APIError.
type APIError struct {
	// Fields An array of errors associated with specific fields.
	Fields *[]FieldError `json:"fields,omitempty"`

	// Message A high-level overview of the error condition.
	Message *string `json:"message,omitempty"`
}

// FieldError defines model for FieldError.
type FieldError struct {
	// Field The name of a field that failed validation.
	Field string `json:"field"`

	// Message A description of why the field is invalid.
	Message string `json:"message"`
}

// Model defines model for Model.
type Model struct {
	// CreatedAt The instant the vendor was added to the system.
	CreatedAt time.Time `json:"createdAt"`

	// Id A unique identifier for the model.
	Id int `json:"id"`

	// Model The unique vendor-provided identifier for the model.
	Model string `json:"model"`

	// Name A readable name for the vendor.
	Name string `json:"name"`

	// UpdatedAt The instant the vendor's information was last updated.
	UpdatedAt time.Time `json:"updatedAt"`

	// VendorID The ID of the vendor who produces the model.
	VendorID int `json:"vendorID"`
}

// ModelCollection defines model for ModelCollection.
type ModelCollection struct {
	Items []Model `json:"items"`
}

// NewModel defines model for NewModel.
type NewModel struct {
	// Model The unique vendor-provided identifier for the model.
	Model string `json:"model"`

	// Name A readable name for the vendor.
	Name *string `json:"name,omitempty"`
}

// NewVendor defines model for NewVendor.
type NewVendor struct {
	// Name A readable name for the vendor.
	Name string `json:"name" validate:"required,min=1,max=150"`
}

// Vendor defines model for Vendor.
type Vendor struct {
	// CreatedAt The instant the vendor was added to the system.
	CreatedAt time.Time `json:"created_at"`

	// Id A unique identifier for the vendor.
	Id int `json:"id"`

	// Name A readable name for the vendor.
	Name string `json:"name"`

	// UpdatedAt The instant the vendor's information was last updated.
	UpdatedAt time.Time `json:"updated_at"`
}

// VendorCollection defines model for VendorCollection.
type VendorCollection struct {
	Items []Vendor `json:"items"`
}

// PutModelsModelIDJSONRequestBody defines body for PutModelsModelID for application/json ContentType.
type PutModelsModelIDJSONRequestBody = NewModel

// PostVendorsJSONRequestBody defines body for PostVendors for application/json ContentType.
type PostVendorsJSONRequestBody = NewVendor

// PostVendorsVendorIDModelsJSONRequestBody defines body for PostVendorsVendorIDModels for application/json ContentType.
type PostVendorsVendorIDModelsJSONRequestBody = NewModel

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /models)
	GetModels(w http.ResponseWriter, r *http.Request)

	// (DELETE /models/{modelID})
	DeleteModelsModelID(w http.ResponseWriter, r *http.Request, modelID int)

	// (GET /models/{modelID})
	GetModelsModelID(w http.ResponseWriter, r *http.Request, modelID int)

	// (PUT /models/{modelID})
	PutModelsModelID(w http.ResponseWriter, r *http.Request, modelID int)

	// (GET /vendors)
	GetVendors(w http.ResponseWriter, r *http.Request)

	// (POST /vendors)
	PostVendors(w http.ResponseWriter, r *http.Request)

	// (DELETE /vendors/{vendorID})
	DeleteVendorsVendorID(w http.ResponseWriter, r *http.Request, vendorID int)

	// (GET /vendors/{vendorID})
	GetVendorsVendorID(w http.ResponseWriter, r *http.Request, vendorID int)
	// List vendor models
	// (GET /vendors/{vendorID}/models)
	GetVendorsVendorIDModels(w http.ResponseWriter, r *http.Request, vendorID int)
	// Add vendor model
	// (POST /vendors/{vendorID}/models)
	PostVendorsVendorIDModels(w http.ResponseWriter, r *http.Request, vendorID int)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetModels operation middleware
func (siw *ServerInterfaceWrapper) GetModels(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetModels(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// DeleteModelsModelID operation middleware
func (siw *ServerInterfaceWrapper) DeleteModelsModelID(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "modelID" -------------
	var modelID int

	err = runtime.BindStyledParameterWithOptions("simple", "modelID", r.PathValue("modelID"), &modelID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "modelID", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteModelsModelID(w, r, modelID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetModelsModelID operation middleware
func (siw *ServerInterfaceWrapper) GetModelsModelID(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "modelID" -------------
	var modelID int

	err = runtime.BindStyledParameterWithOptions("simple", "modelID", r.PathValue("modelID"), &modelID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "modelID", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetModelsModelID(w, r, modelID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// PutModelsModelID operation middleware
func (siw *ServerInterfaceWrapper) PutModelsModelID(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "modelID" -------------
	var modelID int

	err = runtime.BindStyledParameterWithOptions("simple", "modelID", r.PathValue("modelID"), &modelID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "modelID", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutModelsModelID(w, r, modelID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetVendors operation middleware
func (siw *ServerInterfaceWrapper) GetVendors(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetVendors(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// PostVendors operation middleware
func (siw *ServerInterfaceWrapper) PostVendors(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostVendors(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// DeleteVendorsVendorID operation middleware
func (siw *ServerInterfaceWrapper) DeleteVendorsVendorID(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "vendorID" -------------
	var vendorID int

	err = runtime.BindStyledParameterWithOptions("simple", "vendorID", r.PathValue("vendorID"), &vendorID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "vendorID", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteVendorsVendorID(w, r, vendorID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetVendorsVendorID operation middleware
func (siw *ServerInterfaceWrapper) GetVendorsVendorID(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "vendorID" -------------
	var vendorID int

	err = runtime.BindStyledParameterWithOptions("simple", "vendorID", r.PathValue("vendorID"), &vendorID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "vendorID", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetVendorsVendorID(w, r, vendorID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetVendorsVendorIDModels operation middleware
func (siw *ServerInterfaceWrapper) GetVendorsVendorIDModels(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "vendorID" -------------
	var vendorID int

	err = runtime.BindStyledParameterWithOptions("simple", "vendorID", r.PathValue("vendorID"), &vendorID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "vendorID", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetVendorsVendorIDModels(w, r, vendorID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// PostVendorsVendorIDModels operation middleware
func (siw *ServerInterfaceWrapper) PostVendorsVendorIDModels(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "vendorID" -------------
	var vendorID int

	err = runtime.BindStyledParameterWithOptions("simple", "vendorID", r.PathValue("vendorID"), &vendorID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "vendorID", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostVendorsVendorIDModels(w, r, vendorID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{})
}

// ServeMux is an abstraction of http.ServeMux.
type ServeMux interface {
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type StdHTTPServerOptions struct {
	BaseURL          string
	BaseRouter       ServeMux
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, m ServeMux) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseRouter: m,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, m ServeMux, baseURL string) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseURL:    baseURL,
		BaseRouter: m,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options StdHTTPServerOptions) http.Handler {
	m := options.BaseRouter

	if m == nil {
		m = http.NewServeMux()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	m.HandleFunc("GET "+options.BaseURL+"/models", wrapper.GetModels)
	m.HandleFunc("DELETE "+options.BaseURL+"/models/{modelID}", wrapper.DeleteModelsModelID)
	m.HandleFunc("GET "+options.BaseURL+"/models/{modelID}", wrapper.GetModelsModelID)
	m.HandleFunc("PUT "+options.BaseURL+"/models/{modelID}", wrapper.PutModelsModelID)
	m.HandleFunc("GET "+options.BaseURL+"/vendors", wrapper.GetVendors)
	m.HandleFunc("POST "+options.BaseURL+"/vendors", wrapper.PostVendors)
	m.HandleFunc("DELETE "+options.BaseURL+"/vendors/{vendorID}", wrapper.DeleteVendorsVendorID)
	m.HandleFunc("GET "+options.BaseURL+"/vendors/{vendorID}", wrapper.GetVendorsVendorID)
	m.HandleFunc("GET "+options.BaseURL+"/vendors/{vendorID}/models", wrapper.GetVendorsVendorIDModels)
	m.HandleFunc("POST "+options.BaseURL+"/vendors/{vendorID}/models", wrapper.PostVendorsVendorIDModels)

	return m
}

type GetModelsRequestObject struct {
}

type GetModelsResponseObject interface {
	VisitGetModelsResponse(w http.ResponseWriter) error
}

type GetModels200JSONResponse ModelCollection

func (response GetModels200JSONResponse) VisitGetModelsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type DeleteModelsModelIDRequestObject struct {
	ModelID int `json:"modelID"`
}

type DeleteModelsModelIDResponseObject interface {
	VisitDeleteModelsModelIDResponse(w http.ResponseWriter) error
}

type DeleteModelsModelID204Response struct {
}

func (response DeleteModelsModelID204Response) VisitDeleteModelsModelIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type DeleteModelsModelID404JSONResponse APIError

func (response DeleteModelsModelID404JSONResponse) VisitDeleteModelsModelIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type GetModelsModelIDRequestObject struct {
	ModelID int `json:"modelID"`
}

type GetModelsModelIDResponseObject interface {
	VisitGetModelsModelIDResponse(w http.ResponseWriter) error
}

type GetModelsModelID200JSONResponse Model

func (response GetModelsModelID200JSONResponse) VisitGetModelsModelIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetModelsModelID404JSONResponse APIError

func (response GetModelsModelID404JSONResponse) VisitGetModelsModelIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type PutModelsModelIDRequestObject struct {
	ModelID int `json:"modelID"`
	Body    *PutModelsModelIDJSONRequestBody
}

type PutModelsModelIDResponseObject interface {
	VisitPutModelsModelIDResponse(w http.ResponseWriter) error
}

type PutModelsModelID200JSONResponse Model

func (response PutModelsModelID200JSONResponse) VisitPutModelsModelIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PutModelsModelID400JSONResponse APIError

func (response PutModelsModelID400JSONResponse) VisitPutModelsModelIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PutModelsModelID404JSONResponse APIError

func (response PutModelsModelID404JSONResponse) VisitPutModelsModelIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type GetVendorsRequestObject struct {
}

type GetVendorsResponseObject interface {
	VisitGetVendorsResponse(w http.ResponseWriter) error
}

type GetVendors200JSONResponse VendorCollection

func (response GetVendors200JSONResponse) VisitGetVendorsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostVendorsRequestObject struct {
	Body *PostVendorsJSONRequestBody
}

type PostVendorsResponseObject interface {
	VisitPostVendorsResponse(w http.ResponseWriter) error
}

type PostVendors201JSONResponse Vendor

func (response PostVendors201JSONResponse) VisitPostVendorsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type PostVendors400JSONResponse APIError

func (response PostVendors400JSONResponse) VisitPostVendorsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type DeleteVendorsVendorIDRequestObject struct {
	VendorID int `json:"vendorID"`
}

type DeleteVendorsVendorIDResponseObject interface {
	VisitDeleteVendorsVendorIDResponse(w http.ResponseWriter) error
}

type DeleteVendorsVendorID204Response struct {
}

func (response DeleteVendorsVendorID204Response) VisitDeleteVendorsVendorIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type DeleteVendorsVendorID404JSONResponse APIError

func (response DeleteVendorsVendorID404JSONResponse) VisitDeleteVendorsVendorIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type GetVendorsVendorIDRequestObject struct {
	VendorID int `json:"vendorID"`
}

type GetVendorsVendorIDResponseObject interface {
	VisitGetVendorsVendorIDResponse(w http.ResponseWriter) error
}

type GetVendorsVendorID200JSONResponse Vendor

func (response GetVendorsVendorID200JSONResponse) VisitGetVendorsVendorIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetVendorsVendorID404JSONResponse APIError

func (response GetVendorsVendorID404JSONResponse) VisitGetVendorsVendorIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type GetVendorsVendorIDModelsRequestObject struct {
	VendorID int `json:"vendorID"`
}

type GetVendorsVendorIDModelsResponseObject interface {
	VisitGetVendorsVendorIDModelsResponse(w http.ResponseWriter) error
}

type GetVendorsVendorIDModels200JSONResponse ModelCollection

func (response GetVendorsVendorIDModels200JSONResponse) VisitGetVendorsVendorIDModelsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetVendorsVendorIDModels404JSONResponse APIError

func (response GetVendorsVendorIDModels404JSONResponse) VisitGetVendorsVendorIDModelsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type PostVendorsVendorIDModelsRequestObject struct {
	VendorID int `json:"vendorID"`
	Body     *PostVendorsVendorIDModelsJSONRequestBody
}

type PostVendorsVendorIDModelsResponseObject interface {
	VisitPostVendorsVendorIDModelsResponse(w http.ResponseWriter) error
}

type PostVendorsVendorIDModels201JSONResponse Model

func (response PostVendorsVendorIDModels201JSONResponse) VisitPostVendorsVendorIDModelsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type PostVendorsVendorIDModels400JSONResponse APIError

func (response PostVendorsVendorIDModels400JSONResponse) VisitPostVendorsVendorIDModelsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PostVendorsVendorIDModels404JSONResponse APIError

func (response PostVendorsVendorIDModels404JSONResponse) VisitPostVendorsVendorIDModelsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (GET /models)
	GetModels(ctx context.Context, request GetModelsRequestObject) (GetModelsResponseObject, error)

	// (DELETE /models/{modelID})
	DeleteModelsModelID(ctx context.Context, request DeleteModelsModelIDRequestObject) (DeleteModelsModelIDResponseObject, error)

	// (GET /models/{modelID})
	GetModelsModelID(ctx context.Context, request GetModelsModelIDRequestObject) (GetModelsModelIDResponseObject, error)

	// (PUT /models/{modelID})
	PutModelsModelID(ctx context.Context, request PutModelsModelIDRequestObject) (PutModelsModelIDResponseObject, error)

	// (GET /vendors)
	GetVendors(ctx context.Context, request GetVendorsRequestObject) (GetVendorsResponseObject, error)

	// (POST /vendors)
	PostVendors(ctx context.Context, request PostVendorsRequestObject) (PostVendorsResponseObject, error)

	// (DELETE /vendors/{vendorID})
	DeleteVendorsVendorID(ctx context.Context, request DeleteVendorsVendorIDRequestObject) (DeleteVendorsVendorIDResponseObject, error)

	// (GET /vendors/{vendorID})
	GetVendorsVendorID(ctx context.Context, request GetVendorsVendorIDRequestObject) (GetVendorsVendorIDResponseObject, error)
	// List vendor models
	// (GET /vendors/{vendorID}/models)
	GetVendorsVendorIDModels(ctx context.Context, request GetVendorsVendorIDModelsRequestObject) (GetVendorsVendorIDModelsResponseObject, error)
	// Add vendor model
	// (POST /vendors/{vendorID}/models)
	PostVendorsVendorIDModels(ctx context.Context, request PostVendorsVendorIDModelsRequestObject) (PostVendorsVendorIDModelsResponseObject, error)
}

type StrictHandlerFunc = strictnethttp.StrictHTTPHandlerFunc
type StrictMiddlewareFunc = strictnethttp.StrictHTTPMiddlewareFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// GetModels operation middleware
func (sh *strictHandler) GetModels(w http.ResponseWriter, r *http.Request) {
	var request GetModelsRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetModels(ctx, request.(GetModelsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetModels")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetModelsResponseObject); ok {
		if err := validResponse.VisitGetModelsResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// DeleteModelsModelID operation middleware
func (sh *strictHandler) DeleteModelsModelID(w http.ResponseWriter, r *http.Request, modelID int) {
	var request DeleteModelsModelIDRequestObject

	request.ModelID = modelID

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteModelsModelID(ctx, request.(DeleteModelsModelIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteModelsModelID")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(DeleteModelsModelIDResponseObject); ok {
		if err := validResponse.VisitDeleteModelsModelIDResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetModelsModelID operation middleware
func (sh *strictHandler) GetModelsModelID(w http.ResponseWriter, r *http.Request, modelID int) {
	var request GetModelsModelIDRequestObject

	request.ModelID = modelID

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetModelsModelID(ctx, request.(GetModelsModelIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetModelsModelID")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetModelsModelIDResponseObject); ok {
		if err := validResponse.VisitGetModelsModelIDResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PutModelsModelID operation middleware
func (sh *strictHandler) PutModelsModelID(w http.ResponseWriter, r *http.Request, modelID int) {
	var request PutModelsModelIDRequestObject

	request.ModelID = modelID

	var body PutModelsModelIDJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PutModelsModelID(ctx, request.(PutModelsModelIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PutModelsModelID")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PutModelsModelIDResponseObject); ok {
		if err := validResponse.VisitPutModelsModelIDResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetVendors operation middleware
func (sh *strictHandler) GetVendors(w http.ResponseWriter, r *http.Request) {
	var request GetVendorsRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetVendors(ctx, request.(GetVendorsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetVendors")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetVendorsResponseObject); ok {
		if err := validResponse.VisitGetVendorsResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostVendors operation middleware
func (sh *strictHandler) PostVendors(w http.ResponseWriter, r *http.Request) {
	var request PostVendorsRequestObject

	var body PostVendorsJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostVendors(ctx, request.(PostVendorsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostVendors")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostVendorsResponseObject); ok {
		if err := validResponse.VisitPostVendorsResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// DeleteVendorsVendorID operation middleware
func (sh *strictHandler) DeleteVendorsVendorID(w http.ResponseWriter, r *http.Request, vendorID int) {
	var request DeleteVendorsVendorIDRequestObject

	request.VendorID = vendorID

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteVendorsVendorID(ctx, request.(DeleteVendorsVendorIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteVendorsVendorID")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(DeleteVendorsVendorIDResponseObject); ok {
		if err := validResponse.VisitDeleteVendorsVendorIDResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetVendorsVendorID operation middleware
func (sh *strictHandler) GetVendorsVendorID(w http.ResponseWriter, r *http.Request, vendorID int) {
	var request GetVendorsVendorIDRequestObject

	request.VendorID = vendorID

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetVendorsVendorID(ctx, request.(GetVendorsVendorIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetVendorsVendorID")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetVendorsVendorIDResponseObject); ok {
		if err := validResponse.VisitGetVendorsVendorIDResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetVendorsVendorIDModels operation middleware
func (sh *strictHandler) GetVendorsVendorIDModels(w http.ResponseWriter, r *http.Request, vendorID int) {
	var request GetVendorsVendorIDModelsRequestObject

	request.VendorID = vendorID

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetVendorsVendorIDModels(ctx, request.(GetVendorsVendorIDModelsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetVendorsVendorIDModels")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetVendorsVendorIDModelsResponseObject); ok {
		if err := validResponse.VisitGetVendorsVendorIDModelsResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostVendorsVendorIDModels operation middleware
func (sh *strictHandler) PostVendorsVendorIDModels(w http.ResponseWriter, r *http.Request, vendorID int) {
	var request PostVendorsVendorIDModelsRequestObject

	request.VendorID = vendorID

	var body PostVendorsVendorIDModelsJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostVendorsVendorIDModels(ctx, request.(PostVendorsVendorIDModelsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostVendorsVendorIDModels")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostVendorsVendorIDModelsResponseObject); ok {
		if err := validResponse.VisitPostVendorsVendorIDModelsResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}
