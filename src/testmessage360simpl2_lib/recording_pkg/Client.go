/*
 * testmessage360simpl2_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ) on 10/03/2016
 */
package recording_pkg


import(
	"github.com/apimatic/unirest-go"
	"testmessage360simpl2_lib"
	"testmessage360simpl2_lib/apihelper_pkg"
)
/*
 * Client structure as interface implementation
 */
type RECORDING_IMPL struct { }

/**
 * View Recording JSON
 * @param    string        recordingSid     parameter: Required
 * @return	Returns the string response from the API call
 */
func (me *RECORDING_IMPL) CreateViewRecordingJSON (
            recordingSid string) (string, error) {
        //the base uri for api requests
    _queryBuilder := testmessage360simpl2_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/recording/viewrecording.json"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "RecordingSid" : recordingSid,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, testmessage360simpl2_lib.Config.AccountSid, testmessage360simpl2_lib.Config.AuthToken)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil
}

/**
 * View Recording XML
 * @param    string        recordingSid     parameter: Required
 * @return	Returns the string response from the API call
 */
func (me *RECORDING_IMPL) CreateViewRecordingXML (
            recordingSid string) (string, error) {
        //the base uri for api requests
    _queryBuilder := testmessage360simpl2_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/recording/viewrecording.xml"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "RecordingSid" : recordingSid,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, testmessage360simpl2_lib.Config.AccountSid, testmessage360simpl2_lib.Config.AuthToken)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil
}

/**
 * Delete Recording Record JSON
 * @param    string        recordingSid     parameter: Required
 * @return	Returns the string response from the API call
 */
func (me *RECORDING_IMPL) CreateDeleteRecordingJSON (
            recordingSid string) (string, error) {
        //the base uri for api requests
    _queryBuilder := testmessage360simpl2_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/recording/deleterecording.json"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "RecordingSid" : recordingSid,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, testmessage360simpl2_lib.Config.AccountSid, testmessage360simpl2_lib.Config.AuthToken)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil
}

/**
 * Delete Recording Record XML
 * @param    string        recordingSid     parameter: Required
 * @return	Returns the string response from the API call
 */
func (me *RECORDING_IMPL) CreateDeleteRecordingXML (
            recordingSid string) (string, error) {
        //the base uri for api requests
    _queryBuilder := testmessage360simpl2_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/recording/deleterecording.xml"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "RecordingSid" : recordingSid,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, testmessage360simpl2_lib.Config.AccountSid, testmessage360simpl2_lib.Config.AuthToken)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil
}

/**
 * List Recording XML
 * @param    *int64         page            parameter: Optional
 * @param    *int64         pageSize        parameter: Optional
 * @param    *string        dateCreated     parameter: Optional
 * @param    *string        callSid         parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *RECORDING_IMPL) CreateListRecordingXML (
            page *int64,
            pageSize *int64,
            dateCreated *string,
            callSid *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := testmessage360simpl2_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/recording/listrecording.xml"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "Page" : page,
        "PageSize" : pageSize,
        "DateCreated" : dateCreated,
        "CallSid" : callSid,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, testmessage360simpl2_lib.Config.AccountSid, testmessage360simpl2_lib.Config.AuthToken)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil
}

/**
 * List Recording  JSON
 * @param    *int64         page            parameter: Optional
 * @param    *int64         pageSize        parameter: Optional
 * @param    *string        dateCreated     parameter: Optional
 * @param    *string        callSid         parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *RECORDING_IMPL) CreateListRecordingJSON (
            page *int64,
            pageSize *int64,
            dateCreated *string,
            callSid *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := testmessage360simpl2_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/recording/listrecording.json"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "Page" : page,
        "PageSize" : pageSize,
        "DateCreated" : dateCreated,
        "CallSid" : callSid,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, testmessage360simpl2_lib.Config.AccountSid, testmessage360simpl2_lib.Config.AuthToken)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil
}

