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

//func NewStreamPeerSSLFromPointer(ptr gdnative.Pointer) StreamPeerSSL {
func newStreamPeerSSLFromPointer(ptr gdnative.Pointer) StreamPeerSSL {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := StreamPeerSSL{}
	obj.SetBaseObject(owner)

	return obj
}

/*
SSL Stream peer. This object can be used to connect to SSL servers.
*/
type StreamPeerSSL struct {
	StreamPeer
	owner gdnative.Object
}

func (o *StreamPeerSSL) BaseClass() string {
	return "StreamPeerSSL"
}

/*

	Args: [{ false stream StreamPeer}], Returns: enum.Error
*/

/*
        Connect to a peer using an underlying [StreamPeer] "stream", when "validate_certs" is true, [code]StreamPeerSSL[/code] will validate that the certificate presented by the peer matches the "for_hostname".
	Args: [{ false stream StreamPeer} {False true validate_certs bool} { true for_hostname String}], Returns: enum.Error
*/

/*
        Disconnect from host.
	Args: [], Returns: void
*/
func (o *StreamPeerSSL) DisconnectFromStream() {
	//log.Println("Calling StreamPeerSSL.DisconnectFromStream()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StreamPeerSSL", "disconnect_from_stream")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Return the status of the connection, one of STATUS_* enum.
	Args: [], Returns: enum.StreamPeerSSL::Status
*/

// StreamPeerSSLImplementer is an interface that implements the methods
// of the StreamPeerSSL class.
type StreamPeerSSLImplementer interface {
	StreamPeerImplementer
	DisconnectFromStream()
}
