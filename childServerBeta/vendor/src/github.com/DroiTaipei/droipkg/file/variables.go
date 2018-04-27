package file

const (
	// file status
	StatusDfsCreated    = 0
	StatusDfsUploading  = 1
	StatusDfsCompleted  = 2
	StatusDfsDelivering = 3
	StatusDfsDelete     = 4
	StatusDfsConcatErr  = 5
	StatusDfsCorrupt    = 6
	StatusCdnAssignJob  = 7
	StatusCdnRetrieved  = 8
	StatusCdnCorrupt    = 9
	StatusCdnCompleted  = 10

	LogAppID     = "AppID"
	LogRequestID = "RequestID"
	LogFileID    = "FileID"
	LogVendor    = "Vendor"
	// Log fields: Majesty
	LogOptional    = "Op"
	LogRoleID      = "RoleID"
	LogEventID     = "EventID"
	LogEventMsg    = "EventMsg"
	LogCallbackMsg = "CallbackMsg"

	// Log fields: CDN Publisher
	LogCdnURL             = "CdnURL"
	LogSourceURL          = "SourceURL"
	LogScanURL            = "ScanURL"
	LogPublishAction      = "Act"
	LogItemID             = "ItemID"
	LogElapsedTime        = "ElapsedTime" // seconds
	LogWangsuPersistentId = "PersistentId"

	// CDN vendors
	VendorGosunName    = "gosun"
	VendorGosunDomain  = "newmarket1.oo523.com"
	VendorWangsuName   = "wangsu"
	VendorWangsuDomain = "droibaas.yy845.com"
	VendorQiniuName    = "qiniu"
	VendorQiniuDomain  = "droibaascdn.com"

	// Slack Integration
	SlackChannelGeneral = "file_service"
	SlackChannelAlert   = "file_service_alert"
	SlackChannelTest    = "file_service_test"
)

// See http://mkdocs/file-majesty
var (
	// Majesty events
	EventInit       uint8 = 1
	EventDeleteFile uint8 = 2
	EventUpdateFile uint8 = 3
	// File Upload and Abyss
	EventSaved uint8 = 16
	// CDN Publisher
	EventCdnSentRequest uint8 = 20
	EventCdnReady       uint8 = 21
	EventCdnFailed      uint8 = 22
	EventCdnDeleted     uint8 = 23
	EventCdnMarkDelete  uint8 = 24
	EventCdnUpdated     uint8 = 25
	// qiniu file operations
	EventFileScanPass      uint8 = 40
	EventFileScanFail      uint8 = 41
	EventFileScanSpecious  uint8 = 42
	EventFileScanError     uint8 = 43
	EventImageScanPass     uint8 = 44
	EventImageScanFail     uint8 = 45
	EventImageScanSpecious uint8 = 46
	EventImageScanError    uint8 = 47

	// Events - the mapping between eventID and name
	Events = map[uint8]string{
		// Global
		EventInit:       "EventInit",
		EventDeleteFile: "EventDeleteFile",
		EventUpdateFile: "EventUpdateFile",
		// File Upload and Abyss
		EventSaved: "EventSaved",
		// CDN Publisher
		EventCdnSentRequest: "EventCdnSentRequest",
		EventCdnReady:       "EventCdnReady",
		EventCdnFailed:      "EventCdnFailed",
		EventCdnDeleted:     "EventCdnDeleted",
		EventCdnMarkDelete:  "EventCdnMarkDelete",
		EventCdnUpdated:     "EventCdnUpdated",
		// qiniu file operations
		EventFileScanPass:      "EventFileScanPass",
		EventFileScanFail:      "EventFileScanFail",
		EventFileScanSpecious:  "EventFileScanSpecious",
		EventFileScanError:     "EventFileScanError",
		EventImageScanPass:     "EventImageScanPass",
		EventImageScanFail:     "EventImageScanFail",
		EventImageScanSpecious: "EventImageScanSpecious",
		EventImageScanError:    "EventImageScanError",
	}

	// roles
	RoleUnknown      uint8
	RoleMajesty      uint8 = 1
	RoleFileUpload   uint8 = 2
	RoleFileChecker  uint8 = 3
	RoleCdnPublisher uint8 = 4
	RoleCloudOps     uint8 = 5
	RoleAbyss        uint8 = 6

	// Roles - the mapping between roleID and name
	Roles = map[uint8]string{
		RoleUnknown:      "Unknown",
		RoleMajesty:      "Majesty",
		RoleFileUpload:   "FileUpload",
		RoleFileChecker:  "FileChecker",
		RoleCdnPublisher: "CdnPublisher",
		RoleCloudOps:     "CloudOps",
		RoleAbyss:        "Abyss",
	}

	OpCodePublish uint8
	OpCodeDelete  uint8 = 1
	OpCodeUpdate  uint8 = 2

	// OpCodes - the mapping between operationID and name
	OpCodes = map[uint8]string{
		OpCodePublish: "publish",
		OpCodeDelete:  "delete",
		OpCodeUpdate:  "update",
	}
)
