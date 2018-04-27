package file

import (
	de "github.com/DroiTaipei/droipkg"
)

// https://docs.google.com/spreadsheets/d/150AALDHHdKyKXzIk-PnrcNfTm9fEs6dO64vrcdcj0m0/edit#gid=1670635339
var (
	// Global
	OkCode          = 0
	OkCarrier       = de.CarrierDroiError("0000000 OK ")
	ErrUnknown      = de.CarrierDroiError("1120000 Unknown error. ")
	ErrRequest      = de.CarrierDroiError("1120001 Fail to accept request. Check path, query string, or body. ")
	ErrResponse     = de.CarrierDroiError("1120002 Fail to generate response. ")
	ErrRdbAPI       = de.CarrierDroiError("1120003 Fail to access RDB. ")
	ErrSlackWebhook = de.CarrierDroiError("1120004 Fail to send slack webhook. ")

	// FileUpload
	ErrFileUploadUnknown     = de.CarrierDroiError("1120100 Unknown error. ")
	ErrFileUploadConn        = de.CarrierDroiError("1120101 Fail to connect role. ")
	ErrFileUploadRequest     = de.CarrierDroiError("1120102 Request is rejected by role. ")
	ErrFileUploadResponse    = de.CarrierDroiError("1120103 Fail to get response from role. ")
	ErrFileUploadInvalidResp = de.CarrierDroiError("1120104 Invalid response. ")

	// File Checker
	ErrFileCheckerUnknown     = de.CarrierDroiError("1120200 Unknown error. ")
	ErrFileCheckerConn        = de.CarrierDroiError("1120201 Fail to connect role. ")
	ErrFileCheckerRequest     = de.CarrierDroiError("1120202 Request is rejected by role. ")
	ErrFileCheckerResponse    = de.CarrierDroiError("1120203 Fail to get response from role. ")
	ErrFileCheckerInvalidResp = de.CarrierDroiError("1120204 Invalid response. ")
	ErrFileCheckerScan        = de.CarrierDroiError("1120205 Fail to scan file. ")

	//CDN Publisher
	ErrCdnPublisherUnknown     = de.CarrierDroiError("1120300 Unknown error. ")
	ErrCdnPublisherConn        = de.CarrierDroiError("1120301 Fail to connect role. ")
	ErrCdnPublisherRequest     = de.CarrierDroiError("1120302 Request is rejected by role. ")
	ErrCdnPublisherResponse    = de.CarrierDroiError("1120303 Fail to get response from role. ")
	ErrCdnPublisherInvalidResp = de.CarrierDroiError("1120304 Invalid response. ")
	ErrCdnPublisherRdbAPI      = de.CarrierDroiError("1120305 Fail to access RDB API. ")
	ErrCdnPublisherGosun       = de.CarrierDroiError("1120306 Gosun returns error. ")
	ErrCdnPublisherGosunReq    = de.CarrierDroiError("1120307 Fail to request Gosun. ")
	ErrCdnPublisherGosunResp   = de.CarrierDroiError("1120308 Fail to handle response from Gosun. ")
	ErrCdnPublisherWangsu      = de.CarrierDroiError("1120309 Wangsu returns error. ")
	ErrCdnPublisherWangsuReq   = de.CarrierDroiError("1120310 Fail to request Wangsu. ")
	ErrCdnPublisherWangsuResp  = de.CarrierDroiError("1120311 Fail to handle response from Wangsu. ")
)
