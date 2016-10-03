/*
 * testmessage360simpl2_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ) on 10/03/2016
 */
package conference_pkg


import(
	"testmessage360simpl2_lib/models_pkg"
	"github.com/apimatic/unirest-go"
	"testmessage360simpl2_lib"
	"testmessage360simpl2_lib/apihelper_pkg"
)
/*
 * Client structure as interface implementation
 */
type CONFERENCE_IMPL struct { }

/**
 * View Conference JSON
 * @param    string        conferencesid     parameter: Required
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) CreateViewConferenceJSON (
            conferencesid string) (string, error) {
        //the base uri for api requests
    _queryBuilder := testmessage360simpl2_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/conferences/viewconference.json"

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

        "conferencesid" : conferencesid,

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
 * View Conference XML
 * @param    string        conferencesid     parameter: Required
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) CreateViewConferenceXML (
            conferencesid string) (string, error) {
        //the base uri for api requests
    _queryBuilder := testmessage360simpl2_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/conferences/viewconference.xml"

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

        "conferencesid" : conferencesid,

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
 * View Participant JSON
 * @param    string        conferenceSid      parameter: Required
 * @param    string        participantSid     parameter: Required
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) CreateViewParticipantJSON (
            conferenceSid string,
            participantSid string) (string, error) {
        //the base uri for api requests
    _queryBuilder := testmessage360simpl2_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/conferences/viewparticipant.json"

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

        "ConferenceSid" : conferenceSid,
        "ParticipantSid" : participantSid,

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
 * View Participant XML
 * @param    string        conferenceSid      parameter: Required
 * @param    string        participantSid     parameter: Required
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) CreateViewParticipantXML (
            conferenceSid string,
            participantSid string) (string, error) {
        //the base uri for api requests
    _queryBuilder := testmessage360simpl2_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/conferences/viewparticipant.xml"

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

        "ConferenceSid" : conferenceSid,
        "ParticipantSid" : participantSid,

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
 * List Conference JSON
 * @param    *int64                                      page             parameter: Optional
 * @param    *int64                                      pageSize         parameter: Optional
 * @param    *string                                     friendlyName     parameter: Optional
 * @param    models_pkg.InterruptedCallStatusEnum        status           parameter: Optional
 * @param    *string                                     dateCreated      parameter: Optional
 * @param    *string                                     dateUpdated      parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) CreateListConferenceJSON (
            page *int64,
            pageSize *int64,
            friendlyName *string,
            status models_pkg.InterruptedCallStatusEnum,
            dateCreated *string,
            dateUpdated *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := testmessage360simpl2_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/conferences/listconference.json"

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
        "FriendlyName" : friendlyName,
        "Status" : models_pkg.InterruptedCallStatusEnumToValue(status),
        "DateCreated" : dateCreated,
        "DateUpdated" : dateUpdated,

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
 * List Conference JSON
 * @param    *int64                                      page             parameter: Optional
 * @param    *int64                                      pageSize         parameter: Optional
 * @param    *string                                     friendlyName     parameter: Optional
 * @param    models_pkg.InterruptedCallStatusEnum        status           parameter: Optional
 * @param    *string                                     dateCreated      parameter: Optional
 * @param    *string                                     dateUpdated      parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) CreateListConferenceXML (
            page *int64,
            pageSize *int64,
            friendlyName *string,
            status models_pkg.InterruptedCallStatusEnum,
            dateCreated *string,
            dateUpdated *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := testmessage360simpl2_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/conferences/listconference.xml"

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
        "FriendlyName" : friendlyName,
        "Status" : models_pkg.InterruptedCallStatusEnumToValue(status),
        "DateCreated" : dateCreated,
        "DateUpdated" : dateUpdated,

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
 * Add Participant in conference JSON 
 * @param    string        conferencesid         parameter: Required
 * @param    string        participantnumber     parameter: Required
 * @param    int64         tocountrycode         parameter: Required
 * @param    *bool         muted                 parameter: Optional
 * @param    *bool         deaf                  parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) AddParticipantJSON (
            conferencesid string,
            participantnumber string,
            tocountrycode int64,
            muted *bool,
            deaf *bool) (string, error) {
        //the base uri for api requests
    _queryBuilder := testmessage360simpl2_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/conferences/addParticipant.json"

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

        "conferencesid" : conferencesid,
        "participantnumber" : participantnumber,
        "tocountrycode" : tocountrycode,
        "muted" : muted,
        "deaf" : deaf,

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
 * Add Participant in conference XML 
 * @param    string        conferencesid         parameter: Required
 * @param    string        participantnumber     parameter: Required
 * @param    int64         tocountrycode         parameter: Required
 * @param    *bool         muted                 parameter: Optional
 * @param    *bool         deaf                  parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) AddParticipantXML (
            conferencesid string,
            participantnumber string,
            tocountrycode int64,
            muted *bool,
            deaf *bool) (string, error) {
        //the base uri for api requests
    _queryBuilder := testmessage360simpl2_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/conferences/addParticipant.xml"

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

        "conferencesid" : conferencesid,
        "participantnumber" : participantnumber,
        "tocountrycode" : tocountrycode,
        "muted" : muted,
        "deaf" : deaf,

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
 * List Participant JSON
 * @param    string        conferenceSid     parameter: Required
 * @param    *int64        page              parameter: Optional
 * @param    *int64        pagesize          parameter: Optional
 * @param    *bool         muted             parameter: Optional
 * @param    *bool         deaf              parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) CreateListParticipantJSON (
            conferenceSid string,
            page *int64,
            pagesize *int64,
            muted *bool,
            deaf *bool) (string, error) {
        //the base uri for api requests
    _queryBuilder := testmessage360simpl2_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/conferences/listparticipant.json"

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

        "ConferenceSid" : conferenceSid,
        "Page" : page,
        "Pagesize" : pagesize,
        "Muted" : muted,
        "Deaf" : deaf,

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
 * List Participant XML
 * @param    string        conferenceSid     parameter: Required
 * @param    *int64        page              parameter: Optional
 * @param    *int64        pagesize          parameter: Optional
 * @param    *bool         muted             parameter: Optional
 * @param    *bool         deaf              parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) CreateListParticipantXML (
            conferenceSid string,
            page *int64,
            pagesize *int64,
            muted *bool,
            deaf *bool) (string, error) {
        //the base uri for api requests
    _queryBuilder := testmessage360simpl2_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/conferences/listparticipant.xml"

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

        "ConferenceSid" : conferenceSid,
        "Page" : page,
        "Pagesize" : pagesize,
        "Muted" : muted,
        "Deaf" : deaf,

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

