package godot

import (
	"github.com/shadowapex/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

//func NewHTTPClientFromPointer(ptr gdnative.Pointer) HTTPClient {
func newHTTPClientFromPointer(ptr gdnative.Pointer) HTTPClient {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := HTTPClient{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Hyper-text transfer protocol client (sometimes called "User Agent"). Used to make HTTP requests to download web content, upload files and other data or to communicate with various services, among other use cases. Note that this client only needs to connect to a host once (see [method connect_to_host]) to send multiple requests. Because of this, methods that take URLs usually take just the part after the host instead of the full URL, as the client is already connected to a host. See [method request] for a full example and to get started. A [code]HTTPClient[/code] should be reused between multiple requests or to connect to different hosts instead of creating one client per request. Supports SSL and SSL server certificate verification. HTTP status codes in the 2xx range indicate success, 3xx redirection (i.e. "try again, but over here"), 4xx something was wrong with the request, and 5xx something went wrong on the server's side. For more information on HTTP, see https://developer.mozilla.org/en-US/docs/Web/HTTP (or read RFC 2616 to get it straight from the source: https://tools.ietf.org/html/rfc2616).
*/
type HTTPClient struct {
	Reference
	owner gdnative.Object
}

func (o *HTTPClient) BaseClass() string {
	return "HTTPClient"
}

/*
        Closes the current connection, allowing reuse of this [code]HTTPClient[/code].
	Args: [], Returns: void
*/
func (o *HTTPClient) Close() {
	//log.Println("Calling HTTPClient.Close()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "close")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Connect to a host. This needs to be done before any requests are sent. The host should not have http:// prepended but will strip the protocol identifier if provided. If no [code]port[/code] is specified (or [code]-1[/code] is used), it is automatically set to 80 for HTTP and 443 for HTTPS (if [code]use_ssl[/code] is enabled). [code]verify_host[/code] will check the SSL identity of the host if set to [code]true[/code].
	Args: [{ false host String} {-1 true port int} {False true use_ssl bool} {True true verify_host bool}], Returns: enum.Error
*/

/*
        Undocumented
	Args: [], Returns: StreamPeer
*/
func (o *HTTPClient) GetConnection() StreamPeerImplementer {
	//log.Println("Calling HTTPClient.GetConnection()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "get_connection")

	// Call the parent method.
	// StreamPeer
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newStreamPeerFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(StreamPeerImplementer)
	}

	return &ret
}

/*
        Returns the response's body length.
	Args: [], Returns: int
*/
func (o *HTTPClient) GetResponseBodyLength() gdnative.Int {
	//log.Println("Calling HTTPClient.GetResponseBodyLength()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "get_response_body_length")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the response's HTTP status code.
	Args: [], Returns: int
*/
func (o *HTTPClient) GetResponseCode() gdnative.Int {
	//log.Println("Calling HTTPClient.GetResponseCode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "get_response_code")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the response headers.
	Args: [], Returns: PoolStringArray
*/
func (o *HTTPClient) GetResponseHeaders() gdnative.PoolStringArray {
	//log.Println("Calling HTTPClient.GetResponseHeaders()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "get_response_headers")

	// Call the parent method.
	// PoolStringArray
	retPtr := gdnative.NewEmptyPoolStringArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolStringArrayFromPointer(retPtr)
	return ret
}

/*
        Returns all response headers as dictionary where the case-sensitivity of the keys and values is kept like the server delivers it. A value is a simple String, this string can have more than one value where "; " is used as separator. Structure: ("key":"value1; value2") Example: (content-length:12), (Content-Type:application/json; charset=UTF-8)
	Args: [], Returns: Dictionary
*/
func (o *HTTPClient) GetResponseHeadersAsDictionary() gdnative.Dictionary {
	//log.Println("Calling HTTPClient.GetResponseHeadersAsDictionary()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "get_response_headers_as_dictionary")

	// Call the parent method.
	// Dictionary
	retPtr := gdnative.NewEmptyDictionary()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewDictionaryFromPointer(retPtr)
	return ret
}

/*
        Returns a STATUS_* enum constant. Need to call [method poll] in order to get status updates.
	Args: [], Returns: enum.HTTPClient::Status
*/

/*
        If [code]true[/code] this [code]HTTPClient[/code] has a response available.
	Args: [], Returns: bool
*/
func (o *HTTPClient) HasResponse() gdnative.Bool {
	//log.Println("Calling HTTPClient.HasResponse()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "has_response")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *HTTPClient) IsBlockingModeEnabled() gdnative.Bool {
	//log.Println("Calling HTTPClient.IsBlockingModeEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "is_blocking_mode_enabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        If [code]true[/code] this [code]HTTPClient[/code] has a response that is chunked.
	Args: [], Returns: bool
*/
func (o *HTTPClient) IsResponseChunked() gdnative.Bool {
	//log.Println("Calling HTTPClient.IsResponseChunked()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "is_response_chunked")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        This needs to be called in order to have any request processed. Check results with [method get_status]
	Args: [], Returns: enum.Error
*/

/*
        Generates a GET/POST application/x-www-form-urlencoded style query string from a provided dictionary, e.g.: [codeblock] var fields = {"username": "user", "password": "pass"} String queryString = httpClient.query_string_from_dict(fields) returns:= "username=user&password=pass" [/codeblock]
	Args: [{ false fields Dictionary}], Returns: String
*/
func (o *HTTPClient) QueryStringFromDict(fields gdnative.Dictionary) gdnative.String {
	//log.Println("Calling HTTPClient.QueryStringFromDict()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromDictionary(fields)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "query_string_from_dict")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Reads one chunk from the response.
	Args: [], Returns: PoolByteArray
*/
func (o *HTTPClient) ReadResponseBodyChunk() gdnative.PoolByteArray {
	//log.Println("Calling HTTPClient.ReadResponseBodyChunk()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "read_response_body_chunk")

	// Call the parent method.
	// PoolByteArray
	retPtr := gdnative.NewEmptyPoolByteArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolByteArrayFromPointer(retPtr)
	return ret
}

/*
        Sends a request to the connected host. The URL parameter is just the part after the host, so for [code]http://somehost.com/index.php[/code], it is [code]index.php[/code]. Headers are HTTP request headers. For available HTTP methods, see [code]METHOD_*[/code]. To create a POST request with query strings to push to the server, do: [codeblock] var fields = {"username" : "user", "password" : "pass"} var queryString = httpClient.query_string_from_dict(fields) var headers = ["Content-Type: application/x-www-form-urlencoded", "Content-Length: " + str(queryString.length())] var result = httpClient.request(httpClient.METHOD_POST, "index.php", headers, queryString) [/codeblock]
	Args: [{ false method int} { false url String} { false headers PoolStringArray} { true body String}], Returns: enum.Error
*/

/*
        Sends a raw request to the connected host. The URL parameter is just the part after the host, so for [code]http://somehost.com/index.php[/code], it is [code]index.php[/code]. Headers are HTTP request headers. For available HTTP methods, see [code]METHOD_*[/code]. Sends the body data raw, as a byte array and does not encode it in any way.
	Args: [{ false method int} { false url String} { false headers PoolStringArray} { false body PoolByteArray}], Returns: enum.Error
*/

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *HTTPClient) SetBlockingMode(enabled gdnative.Bool) {
	//log.Println("Calling HTTPClient.SetBlockingMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "set_blocking_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false connection StreamPeer}], Returns: void
*/
func (o *HTTPClient) SetConnection(connection StreamPeer) {
	//log.Println("Calling HTTPClient.SetConnection()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(connection.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "set_connection")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the size of the buffer used and maximum bytes to read per iteration. see [method read_response_body_chunk]
	Args: [{ false bytes int}], Returns: void
*/
func (o *HTTPClient) SetReadChunkSize(bytes gdnative.Int) {
	//log.Println("Calling HTTPClient.SetReadChunkSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(bytes)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "set_read_chunk_size")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// HTTPClientImplementer is an interface that implements the methods
// of the HTTPClient class.
type HTTPClientImplementer interface {
	ReferenceImplementer
	Close()
	GetConnection() StreamPeerImplementer
	GetResponseBodyLength() gdnative.Int
	GetResponseCode() gdnative.Int
	GetResponseHeaders() gdnative.PoolStringArray
	GetResponseHeadersAsDictionary() gdnative.Dictionary
	HasResponse() gdnative.Bool
	IsBlockingModeEnabled() gdnative.Bool
	IsResponseChunked() gdnative.Bool
	QueryStringFromDict(fields gdnative.Dictionary) gdnative.String
	ReadResponseBodyChunk() gdnative.PoolByteArray
	SetBlockingMode(enabled gdnative.Bool)
	SetConnection(connection StreamPeer)
	SetReadChunkSize(bytes gdnative.Int)
}
