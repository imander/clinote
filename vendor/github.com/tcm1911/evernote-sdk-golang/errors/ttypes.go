// Autogenerated by Thrift Compiler (1.0.0-dev)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package errors

import (
	"bytes"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

//Numeric codes indicating the type of error that occurred on the
//service.
//<dl>
//  <dt>UNKNOWN</dt>
//    <dd>No information available about the error</dd>
//  <dt>BAD_DATA_FORMAT</dt>
//    <dd>The format of the request data was incorrect</dd>
//  <dt>PERMISSION_DENIED</dt>
//    <dd>Not permitted to perform action</dd>
//  <dt>INTERNAL_ERROR</dt>
//    <dd>Unexpected problem with the service</dd>
//  <dt>DATA_REQUIRED</dt>
//    <dd>A required parameter/field was absent</dd>
//  <dt>LIMIT_REACHED</dt>
//    <dd>Operation denied due to data model limit</dd>
//  <dt>QUOTA_REACHED</dt>
//    <dd>Operation denied due to user storage limit</dd>
//  <dt>INVALID_AUTH</dt>
//    <dd>Username and/or password incorrect</dd>
//  <dt>AUTH_EXPIRED</dt>
//    <dd>Authentication token expired</dd>
//  <dt>DATA_CONFLICT</dt>
//    <dd>Change denied due to data model conflict</dd>
//  <dt>ENML_VALIDATION</dt>
//    <dd>Content of submitted note was malformed</dd>
//  <dt>SHARD_UNAVAILABLE</dt>
//    <dd>Service shard with account data is temporarily down</dd>
//  <dt>LEN_TOO_SHORT</dt>
//    <dd>Operation denied due to data model limit, where something such
//        as a string length was too short</dd>
//  <dt>LEN_TOO_LONG</dt>
//    <dd>Operation denied due to data model limit, where something such
//        as a string length was too long</dd>
//  <dt>TOO_FEW</dt>
//    <dd>Operation denied due to data model limit, where there were
//        too few of something.</dd>
//  <dt>TOO_MANY</dt>
//    <dd>Operation denied due to data model limit, where there were
//        too many of something.</dd>
//  <dt>UNSUPPORTED_OPERATION</dt>
//    <dd>Operation denied because it is currently unsupported.</dd>
//  <dt>TAKEN_DOWN</dt>
//    <dd>Operation denied because access to the corresponding object is
//        prohibited in response to a take-down notice.</dd>
//  <dt>RATE_LIMIT_REACHED</dt>
//    <dd>Operation denied because the calling application has reached
//        its hourly API call limit for this user.</dd>
//</dl>
type EDAMErrorCode int64

const (
	EDAMErrorCode_UNKNOWN               EDAMErrorCode = 1
	EDAMErrorCode_BAD_DATA_FORMAT       EDAMErrorCode = 2
	EDAMErrorCode_PERMISSION_DENIED     EDAMErrorCode = 3
	EDAMErrorCode_INTERNAL_ERROR        EDAMErrorCode = 4
	EDAMErrorCode_DATA_REQUIRED         EDAMErrorCode = 5
	EDAMErrorCode_LIMIT_REACHED         EDAMErrorCode = 6
	EDAMErrorCode_QUOTA_REACHED         EDAMErrorCode = 7
	EDAMErrorCode_INVALID_AUTH          EDAMErrorCode = 8
	EDAMErrorCode_AUTH_EXPIRED          EDAMErrorCode = 9
	EDAMErrorCode_DATA_CONFLICT         EDAMErrorCode = 10
	EDAMErrorCode_ENML_VALIDATION       EDAMErrorCode = 11
	EDAMErrorCode_SHARD_UNAVAILABLE     EDAMErrorCode = 12
	EDAMErrorCode_LEN_TOO_SHORT         EDAMErrorCode = 13
	EDAMErrorCode_LEN_TOO_LONG          EDAMErrorCode = 14
	EDAMErrorCode_TOO_FEW               EDAMErrorCode = 15
	EDAMErrorCode_TOO_MANY              EDAMErrorCode = 16
	EDAMErrorCode_UNSUPPORTED_OPERATION EDAMErrorCode = 17
	EDAMErrorCode_TAKEN_DOWN            EDAMErrorCode = 18
	EDAMErrorCode_RATE_LIMIT_REACHED    EDAMErrorCode = 19
)

func (p EDAMErrorCode) String() string {
	switch p {
	case EDAMErrorCode_UNKNOWN:
		return "UNKNOWN"
	case EDAMErrorCode_BAD_DATA_FORMAT:
		return "BAD_DATA_FORMAT"
	case EDAMErrorCode_PERMISSION_DENIED:
		return "PERMISSION_DENIED"
	case EDAMErrorCode_INTERNAL_ERROR:
		return "INTERNAL_ERROR"
	case EDAMErrorCode_DATA_REQUIRED:
		return "DATA_REQUIRED"
	case EDAMErrorCode_LIMIT_REACHED:
		return "LIMIT_REACHED"
	case EDAMErrorCode_QUOTA_REACHED:
		return "QUOTA_REACHED"
	case EDAMErrorCode_INVALID_AUTH:
		return "INVALID_AUTH"
	case EDAMErrorCode_AUTH_EXPIRED:
		return "AUTH_EXPIRED"
	case EDAMErrorCode_DATA_CONFLICT:
		return "DATA_CONFLICT"
	case EDAMErrorCode_ENML_VALIDATION:
		return "ENML_VALIDATION"
	case EDAMErrorCode_SHARD_UNAVAILABLE:
		return "SHARD_UNAVAILABLE"
	case EDAMErrorCode_LEN_TOO_SHORT:
		return "LEN_TOO_SHORT"
	case EDAMErrorCode_LEN_TOO_LONG:
		return "LEN_TOO_LONG"
	case EDAMErrorCode_TOO_FEW:
		return "TOO_FEW"
	case EDAMErrorCode_TOO_MANY:
		return "TOO_MANY"
	case EDAMErrorCode_UNSUPPORTED_OPERATION:
		return "UNSUPPORTED_OPERATION"
	case EDAMErrorCode_TAKEN_DOWN:
		return "TAKEN_DOWN"
	case EDAMErrorCode_RATE_LIMIT_REACHED:
		return "RATE_LIMIT_REACHED"
	}
	return "<UNSET>"
}

func EDAMErrorCodeFromString(s string) (EDAMErrorCode, error) {
	switch s {
	case "UNKNOWN":
		return EDAMErrorCode_UNKNOWN, nil
	case "BAD_DATA_FORMAT":
		return EDAMErrorCode_BAD_DATA_FORMAT, nil
	case "PERMISSION_DENIED":
		return EDAMErrorCode_PERMISSION_DENIED, nil
	case "INTERNAL_ERROR":
		return EDAMErrorCode_INTERNAL_ERROR, nil
	case "DATA_REQUIRED":
		return EDAMErrorCode_DATA_REQUIRED, nil
	case "LIMIT_REACHED":
		return EDAMErrorCode_LIMIT_REACHED, nil
	case "QUOTA_REACHED":
		return EDAMErrorCode_QUOTA_REACHED, nil
	case "INVALID_AUTH":
		return EDAMErrorCode_INVALID_AUTH, nil
	case "AUTH_EXPIRED":
		return EDAMErrorCode_AUTH_EXPIRED, nil
	case "DATA_CONFLICT":
		return EDAMErrorCode_DATA_CONFLICT, nil
	case "ENML_VALIDATION":
		return EDAMErrorCode_ENML_VALIDATION, nil
	case "SHARD_UNAVAILABLE":
		return EDAMErrorCode_SHARD_UNAVAILABLE, nil
	case "LEN_TOO_SHORT":
		return EDAMErrorCode_LEN_TOO_SHORT, nil
	case "LEN_TOO_LONG":
		return EDAMErrorCode_LEN_TOO_LONG, nil
	case "TOO_FEW":
		return EDAMErrorCode_TOO_FEW, nil
	case "TOO_MANY":
		return EDAMErrorCode_TOO_MANY, nil
	case "UNSUPPORTED_OPERATION":
		return EDAMErrorCode_UNSUPPORTED_OPERATION, nil
	case "TAKEN_DOWN":
		return EDAMErrorCode_TAKEN_DOWN, nil
	case "RATE_LIMIT_REACHED":
		return EDAMErrorCode_RATE_LIMIT_REACHED, nil
	}
	return EDAMErrorCode(0), fmt.Errorf("not a valid EDAMErrorCode string")
}

func EDAMErrorCodePtr(v EDAMErrorCode) *EDAMErrorCode { return &v }

func (p EDAMErrorCode) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *EDAMErrorCode) UnmarshalText(text []byte) error {
	q, err := EDAMErrorCodeFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

// This exception is thrown by EDAM procedures when a call fails as a result of
// a problem that a caller may be able to resolve.  For example, if the user
// attempts to add a note to their account which would exceed their storage
// quota, this type of exception may be thrown to indicate the source of the
// error so that they can choose an alternate action.
//
// This exception would not be used for internal system errors that do not
// reflect user actions, but rather reflect a problem within the service that
// the user cannot resolve.
//
// errorCode:  The numeric code indicating the type of error that occurred.
//   must be one of the values of EDAMErrorCode.
//
// parameter:  If the error applied to a particular input parameter, this will
//   indicate which parameter.
//
// Attributes:
//  - ErrorCode
//  - Parameter
type EDAMUserException struct {
	ErrorCode EDAMErrorCode `thrift:"errorCode,1,required" json:"errorCode"`
	Parameter *string       `thrift:"parameter,2" json:"parameter,omitempty"`
}

func NewEDAMUserException() *EDAMUserException {
	return &EDAMUserException{}
}

func (p *EDAMUserException) GetErrorCode() EDAMErrorCode {
	return p.ErrorCode
}

var EDAMUserException_Parameter_DEFAULT string

func (p *EDAMUserException) GetParameter() string {
	if !p.IsSetParameter() {
		return EDAMUserException_Parameter_DEFAULT
	}
	return *p.Parameter
}
func (p *EDAMUserException) IsSetParameter() bool {
	return p.Parameter != nil
}

func (p *EDAMUserException) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetErrorCode bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetErrorCode = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetErrorCode {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field ErrorCode is not set"))
	}
	return nil
}

func (p *EDAMUserException) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		temp := EDAMErrorCode(v)
		p.ErrorCode = temp
	}
	return nil
}

func (p *EDAMUserException) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Parameter = &v
	}
	return nil
}

func (p *EDAMUserException) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("EDAMUserException"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *EDAMUserException) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("errorCode", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:errorCode: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.ErrorCode)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.errorCode (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:errorCode: ", p), err)
	}
	return err
}

func (p *EDAMUserException) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetParameter() {
		if err := oprot.WriteFieldBegin("parameter", thrift.STRING, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:parameter: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Parameter)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.parameter (2) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:parameter: ", p), err)
		}
	}
	return err
}

func (p *EDAMUserException) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EDAMUserException(%+v)", *p)
}

func (p *EDAMUserException) Error() string {
	return p.String()
}

// This exception is thrown by EDAM procedures when a call fails as a result of
// a problem in the service that could not be changed through caller action.
//
// errorCode:  The numeric code indicating the type of error that occurred.
//   must be one of the values of EDAMErrorCode.
//
// message:  This may contain additional information about the error
//
// rateLimitDuration:  Indicates the minimum number of seconds that an application should
//   expect subsequent API calls for this user to fail. The application should not retry
//   API requests for the user until at least this many seconds have passed. Present only
//   when errorCode is RATE_LIMIT_REACHED,
//
// Attributes:
//  - ErrorCode
//  - Message
//  - RateLimitDuration
type EDAMSystemException struct {
	ErrorCode         EDAMErrorCode `thrift:"errorCode,1,required" json:"errorCode"`
	Message           *string       `thrift:"message,2" json:"message,omitempty"`
	RateLimitDuration *int32        `thrift:"rateLimitDuration,3" json:"rateLimitDuration,omitempty"`
}

func NewEDAMSystemException() *EDAMSystemException {
	return &EDAMSystemException{}
}

func (p *EDAMSystemException) GetErrorCode() EDAMErrorCode {
	return p.ErrorCode
}

var EDAMSystemException_Message_DEFAULT string

func (p *EDAMSystemException) GetMessage() string {
	if !p.IsSetMessage() {
		return EDAMSystemException_Message_DEFAULT
	}
	return *p.Message
}

var EDAMSystemException_RateLimitDuration_DEFAULT int32

func (p *EDAMSystemException) GetRateLimitDuration() int32 {
	if !p.IsSetRateLimitDuration() {
		return EDAMSystemException_RateLimitDuration_DEFAULT
	}
	return *p.RateLimitDuration
}
func (p *EDAMSystemException) IsSetMessage() bool {
	return p.Message != nil
}

func (p *EDAMSystemException) IsSetRateLimitDuration() bool {
	return p.RateLimitDuration != nil
}

func (p *EDAMSystemException) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetErrorCode bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetErrorCode = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetErrorCode {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field ErrorCode is not set"))
	}
	return nil
}

func (p *EDAMSystemException) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		temp := EDAMErrorCode(v)
		p.ErrorCode = temp
	}
	return nil
}

func (p *EDAMSystemException) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Message = &v
	}
	return nil
}

func (p *EDAMSystemException) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.RateLimitDuration = &v
	}
	return nil
}

func (p *EDAMSystemException) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("EDAMSystemException"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *EDAMSystemException) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("errorCode", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:errorCode: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.ErrorCode)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.errorCode (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:errorCode: ", p), err)
	}
	return err
}

func (p *EDAMSystemException) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetMessage() {
		if err := oprot.WriteFieldBegin("message", thrift.STRING, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:message: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Message)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.message (2) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:message: ", p), err)
		}
	}
	return err
}

func (p *EDAMSystemException) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetRateLimitDuration() {
		if err := oprot.WriteFieldBegin("rateLimitDuration", thrift.I32, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:rateLimitDuration: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.RateLimitDuration)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.rateLimitDuration (3) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:rateLimitDuration: ", p), err)
		}
	}
	return err
}

func (p *EDAMSystemException) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EDAMSystemException(%+v)", *p)
}

func (p *EDAMSystemException) Error() string {
	return p.String()
}

// This exception is thrown by EDAM procedures when a caller asks to perform
// an operation on an object that does not exist.  This may be thrown based on an invalid
// primary identifier (e.g. a bad GUID), or when the caller refers to an object
// by another unique identifier (e.g. a User's email address).
//
// identifier:  A description of the object that was not found on the server.
//   For example, "Note.notebookGuid" when a caller attempts to create a note in a
//   notebook that does not exist in the user's account.
//
// key:  The value passed from the client in the identifier, which was not
//   found. For example, the GUID that was not found.
//
// Attributes:
//  - Identifier
//  - Key
type EDAMNotFoundException struct {
	Identifier *string `thrift:"identifier,1" json:"identifier,omitempty"`
	Key        *string `thrift:"key,2" json:"key,omitempty"`
}

func NewEDAMNotFoundException() *EDAMNotFoundException {
	return &EDAMNotFoundException{}
}

var EDAMNotFoundException_Identifier_DEFAULT string

func (p *EDAMNotFoundException) GetIdentifier() string {
	if !p.IsSetIdentifier() {
		return EDAMNotFoundException_Identifier_DEFAULT
	}
	return *p.Identifier
}

var EDAMNotFoundException_Key_DEFAULT string

func (p *EDAMNotFoundException) GetKey() string {
	if !p.IsSetKey() {
		return EDAMNotFoundException_Key_DEFAULT
	}
	return *p.Key
}
func (p *EDAMNotFoundException) IsSetIdentifier() bool {
	return p.Identifier != nil
}

func (p *EDAMNotFoundException) IsSetKey() bool {
	return p.Key != nil
}

func (p *EDAMNotFoundException) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *EDAMNotFoundException) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Identifier = &v
	}
	return nil
}

func (p *EDAMNotFoundException) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Key = &v
	}
	return nil
}

func (p *EDAMNotFoundException) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("EDAMNotFoundException"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *EDAMNotFoundException) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetIdentifier() {
		if err := oprot.WriteFieldBegin("identifier", thrift.STRING, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:identifier: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Identifier)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.identifier (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:identifier: ", p), err)
		}
	}
	return err
}

func (p *EDAMNotFoundException) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetKey() {
		if err := oprot.WriteFieldBegin("key", thrift.STRING, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:key: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Key)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.key (2) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:key: ", p), err)
		}
	}
	return err
}

func (p *EDAMNotFoundException) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EDAMNotFoundException(%+v)", *p)
}

func (p *EDAMNotFoundException) Error() string {
	return p.String()
}