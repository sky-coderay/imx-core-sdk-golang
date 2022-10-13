/*
Immutable X API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0
Contact: support@immutable.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


type MetadataApi interface {

	/*
	AddMetadataSchemaToCollection Add metadata schema to collection

	Add metadata schema to collection

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param address Collection contract address
	@return ApiAddMetadataSchemaToCollectionRequest
	*/
	AddMetadataSchemaToCollection(ctx context.Context, address string) ApiAddMetadataSchemaToCollectionRequest

	// AddMetadataSchemaToCollectionExecute executes the request
	//  @return SuccessResponse
	AddMetadataSchemaToCollectionExecute(r ApiAddMetadataSchemaToCollectionRequest) (*SuccessResponse, *http.Response, error)

	/*
	GetMetadataSchema Get collection metadata schema

	Get collection metadata schema

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param address Collection contract address
	@return ApiGetMetadataSchemaRequest
	*/
	GetMetadataSchema(ctx context.Context, address string) ApiGetMetadataSchemaRequest

	// GetMetadataSchemaExecute executes the request
	//  @return []MetadataSchemaProperty
	GetMetadataSchemaExecute(r ApiGetMetadataSchemaRequest) ([]MetadataSchemaProperty, *http.Response, error)

	/*
	UpdateMetadataSchemaByName Update metadata schema by name

	Update metadata schema by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param address Collection contract address
	@param name Metadata schema name
	@return ApiUpdateMetadataSchemaByNameRequest
	*/
	UpdateMetadataSchemaByName(ctx context.Context, address string, name string) ApiUpdateMetadataSchemaByNameRequest

	// UpdateMetadataSchemaByNameExecute executes the request
	//  @return SuccessResponse
	UpdateMetadataSchemaByNameExecute(r ApiUpdateMetadataSchemaByNameRequest) (*SuccessResponse, *http.Response, error)
}

// MetadataApiService MetadataApi service
type MetadataApiService service

type ApiAddMetadataSchemaToCollectionRequest struct {
	ctx context.Context
	ApiService MetadataApi
	address string
	iMXSignature *string
	iMXTimestamp *string
	addMetadataSchemaToCollectionRequest *AddMetadataSchemaToCollectionRequest
}

// String created by signing wallet address and timestamp
func (r ApiAddMetadataSchemaToCollectionRequest) IMXSignature(iMXSignature string) ApiAddMetadataSchemaToCollectionRequest {
	r.iMXSignature = &iMXSignature
	return r
}

// Unix Epoc timestamp
func (r ApiAddMetadataSchemaToCollectionRequest) IMXTimestamp(iMXTimestamp string) ApiAddMetadataSchemaToCollectionRequest {
	r.iMXTimestamp = &iMXTimestamp
	return r
}

// add metadata schema to a collection
func (r ApiAddMetadataSchemaToCollectionRequest) AddMetadataSchemaToCollectionRequest(addMetadataSchemaToCollectionRequest AddMetadataSchemaToCollectionRequest) ApiAddMetadataSchemaToCollectionRequest {
	r.addMetadataSchemaToCollectionRequest = &addMetadataSchemaToCollectionRequest
	return r
}

func (r ApiAddMetadataSchemaToCollectionRequest) Execute() (*SuccessResponse, *http.Response, error) {
	return r.ApiService.AddMetadataSchemaToCollectionExecute(r)
}

/*
AddMetadataSchemaToCollection Add metadata schema to collection

Add metadata schema to collection

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param address Collection contract address
 @return ApiAddMetadataSchemaToCollectionRequest
*/
func (a *MetadataApiService) AddMetadataSchemaToCollection(ctx context.Context, address string) ApiAddMetadataSchemaToCollectionRequest {
	return ApiAddMetadataSchemaToCollectionRequest{
		ApiService: a,
		ctx: ctx,
		address: address,
	}
}

// Execute executes the request
//  @return SuccessResponse
func (a *MetadataApiService) AddMetadataSchemaToCollectionExecute(r ApiAddMetadataSchemaToCollectionRequest) (*SuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetadataApiService.AddMetadataSchemaToCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/collections/{address}/metadata-schema"
	localVarPath = strings.Replace(localVarPath, "{"+"address"+"}", url.PathEscape(parameterToString(r.address, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.iMXSignature == nil {
		return localVarReturnValue, nil, reportError("iMXSignature is required and must be specified")
	}
	if r.iMXTimestamp == nil {
		return localVarReturnValue, nil, reportError("iMXTimestamp is required and must be specified")
	}
	if r.addMetadataSchemaToCollectionRequest == nil {
		return localVarReturnValue, nil, reportError("addMetadataSchemaToCollectionRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["IMX-Signature"] = parameterToString(*r.iMXSignature, "")
	localVarHeaderParams["IMX-Timestamp"] = parameterToString(*r.iMXTimestamp, "")
	// body params
	localVarPostBody = r.addMetadataSchemaToCollectionRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetMetadataSchemaRequest struct {
	ctx context.Context
	ApiService MetadataApi
	address string
}

func (r ApiGetMetadataSchemaRequest) Execute() ([]MetadataSchemaProperty, *http.Response, error) {
	return r.ApiService.GetMetadataSchemaExecute(r)
}

/*
GetMetadataSchema Get collection metadata schema

Get collection metadata schema

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param address Collection contract address
 @return ApiGetMetadataSchemaRequest
*/
func (a *MetadataApiService) GetMetadataSchema(ctx context.Context, address string) ApiGetMetadataSchemaRequest {
	return ApiGetMetadataSchemaRequest{
		ApiService: a,
		ctx: ctx,
		address: address,
	}
}

// Execute executes the request
//  @return []MetadataSchemaProperty
func (a *MetadataApiService) GetMetadataSchemaExecute(r ApiGetMetadataSchemaRequest) ([]MetadataSchemaProperty, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []MetadataSchemaProperty
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetadataApiService.GetMetadataSchema")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/collections/{address}/metadata-schema"
	localVarPath = strings.Replace(localVarPath, "{"+"address"+"}", url.PathEscape(parameterToString(r.address, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateMetadataSchemaByNameRequest struct {
	ctx context.Context
	ApiService MetadataApi
	address string
	name string
	iMXSignature *string
	iMXTimestamp *string
	metadataSchemaRequest *MetadataSchemaRequest
}

// String created by signing wallet address and timestamp
func (r ApiUpdateMetadataSchemaByNameRequest) IMXSignature(iMXSignature string) ApiUpdateMetadataSchemaByNameRequest {
	r.iMXSignature = &iMXSignature
	return r
}

// Unix Epoc timestamp
func (r ApiUpdateMetadataSchemaByNameRequest) IMXTimestamp(iMXTimestamp string) ApiUpdateMetadataSchemaByNameRequest {
	r.iMXTimestamp = &iMXTimestamp
	return r
}

// update metadata schema
func (r ApiUpdateMetadataSchemaByNameRequest) MetadataSchemaRequest(metadataSchemaRequest MetadataSchemaRequest) ApiUpdateMetadataSchemaByNameRequest {
	r.metadataSchemaRequest = &metadataSchemaRequest
	return r
}

func (r ApiUpdateMetadataSchemaByNameRequest) Execute() (*SuccessResponse, *http.Response, error) {
	return r.ApiService.UpdateMetadataSchemaByNameExecute(r)
}

/*
UpdateMetadataSchemaByName Update metadata schema by name

Update metadata schema by name

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param address Collection contract address
 @param name Metadata schema name
 @return ApiUpdateMetadataSchemaByNameRequest
*/
func (a *MetadataApiService) UpdateMetadataSchemaByName(ctx context.Context, address string, name string) ApiUpdateMetadataSchemaByNameRequest {
	return ApiUpdateMetadataSchemaByNameRequest{
		ApiService: a,
		ctx: ctx,
		address: address,
		name: name,
	}
}

// Execute executes the request
//  @return SuccessResponse
func (a *MetadataApiService) UpdateMetadataSchemaByNameExecute(r ApiUpdateMetadataSchemaByNameRequest) (*SuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetadataApiService.UpdateMetadataSchemaByName")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/collections/{address}/metadata-schema/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"address"+"}", url.PathEscape(parameterToString(r.address, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterToString(r.name, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.iMXSignature == nil {
		return localVarReturnValue, nil, reportError("iMXSignature is required and must be specified")
	}
	if r.iMXTimestamp == nil {
		return localVarReturnValue, nil, reportError("iMXTimestamp is required and must be specified")
	}
	if r.metadataSchemaRequest == nil {
		return localVarReturnValue, nil, reportError("metadataSchemaRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["IMX-Signature"] = parameterToString(*r.iMXSignature, "")
	localVarHeaderParams["IMX-Timestamp"] = parameterToString(*r.iMXTimestamp, "")
	// body params
	localVarPostBody = r.metadataSchemaRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}