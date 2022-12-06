package api

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
)

// mimeType represents various MIME type used API responses.
type mimeType string

const (
	// Means no response type.
	mimeNone mimeType = ""
	// Means response type is JSON.
	mimeJSON mimeType = "application/json"
	// Means response type is XML.
	mimeXML mimeType = "application/xml"

	XmlHeader = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
)

// WriteSuccessResponseJSON writes success headers and response if any,
// with content-type set to `application/json`.
func WriteSuccessResponseJSON(w http.ResponseWriter, response interface{}) {
	writeResponse(w, http.StatusOK, encodeResponseJSON(response), mimeJSON)
}

// WriteSuccessResponseXML writes success headers and response if any,
// with content-type set to `application/xml`.
func WriteSuccessResponseXML(w http.ResponseWriter, response interface{}) {
	writeResponse(w, http.StatusOK, encodeResponseXml(response), mimeXML)
}

// writeRedirectSeeOther writes Location header with http status 303
func writeRedirectSeeOther(w http.ResponseWriter, location string) {
	w.Header().Set(Location, location)
	writeResponse(w, http.StatusSeeOther, nil, mimeNone)
}

func writeSuccessResponseHeadersOnly(w http.ResponseWriter) {
	writeResponse(w, http.StatusOK, nil, mimeNone)
}

func writeResponse(w http.ResponseWriter, statusCode int, response []byte, mType mimeType) error {
	if statusCode == 0 {
		statusCode = 200
	}
	// Similar check to http.checkWriteHeaderCode
	if statusCode < 100 || statusCode > 999 {
		// todo add log
		//logger.Error(fmt.Sprintf("invalid WriteHeader code %v", statusCode))
		statusCode = http.StatusInternalServerError
	}
	setCommonHeaders(w)
	if mType != mimeNone {
		w.Header().Set(ContentType, string(mType))
	}
	w.Header().Set(ContentLength, strconv.Itoa(len(response)))
	w.WriteHeader(statusCode)
	if response != nil {
		if _, err := w.Write(response); err != nil {
			return err
		}
	}
	return nil
}

// Encodes the response headers into JSON format.
func encodeResponseJSON(v interface{}) []byte {
	var bytesBuffer bytes.Buffer
	e := json.NewEncoder(&bytesBuffer)
	_ = e.Encode(v)
	return bytesBuffer.Bytes()
}

type XmlCommonResult struct {
	//ResponseWithResponseInfo xml.Name `xml:"ResponseWithResponseInfo"`
	Item interface{}
	Info interface{}
}

// Encodes the response headers into XML format.
func encodeResponseXml(v interface{}) []byte {
	var bytesBuffer bytes.Buffer
	bytesBuffer.WriteString(XmlHeader)
	vv := XmlCommonResult{}
	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice:
		vv.Item = v
	default:
		vv.Info = v
	}
	buf, _ := xml.MarshalIndent(vv, "", "  ")
	bytesBuffer.Write(buf)
	fmt.Println(bytesBuffer.String())
	//fmt.Println(strings.Replace(strData, "RedPacketQueryRequest", "xml", -1))
	return bytesBuffer.Bytes()
}
