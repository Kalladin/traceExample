package rdb

import (
	"time"
)

const (
	// StageFlag Sandbox Value in database
	StageFlagSand = 1
	// StageFlag Production Value in database
	StageFlagProd = 2
	StageSand     = "Sandbox"
	StageProd     = "Production"

	StatusFlagValid   = 1
	StatusFlagDeleted = 2
	StatusValid       = "Valid"
	StatusDeleted     = "Deleted"
	Unknown           = "Unknown"
)

// Application data structure
type Application struct {
	ID              string    `json:"ID" gorm:"primary_key"`
	Name            string    `json:"name"`
	PackageName     string    `json:"packageName"`
	DeveloperID     string    `json:"developerID"`
	URL             string    `json:"URL"`
	Description     string    `json:"description"`
	Icon            string    `json:"icon"`
	ClientKey       string    `json:"clientKey"`
	RestAPIKey      string    `json:"restApiKey"`
	CloudCodeKey    string    `json:"cloudCodeKey"`
	MasterKey       string    `json:"masterKey"`
	MasterKeyCount  int       `json:"masterKeyCount"`
	SecretKey       string    `json:"secretKey"`
	Preference      string    `json:"preference"`
	AuthPublicData  string    `json:"authPublicData"`
	AuthPrivateData string    `json:"authPrivateData"`
	CreatedTime     time.Time `gorm:"column:creation_time" json:"createdTime"`
	Status          int       `json:"status"`
	RunningVer      string    `json:"runningVer"`
	WebConfig       string    `json:"webConfig"`
	StageFlag       int       `json:"stageFlag"`
	QiniuAccount    string    `json:"qiniuAccount"`
	CRC32AppID      int       `gorm:"column:crc32_app_id"`
	ModRID          int       `gorm:"column:mod_r_id"`
}

// TableName return the table name in database for gorm using
func (s Application) TableName() string {
	return "baas.application"
}

// FieldMap return the map between json field name and db field name
func (s Application) FieldMap() map[string]string {
	return map[string]string{
		"name":            "name",
		"packageName":     "package_name",
		"developerID":     "developer_id",
		"URL":             "url",
		"description":     "description",
		"icon":            "icon",
		"clientKey":       "client_key",
		"restApiKey":      "rest_api_key",
		"cloudCodeKey":    "cloud_code_key",
		"masterKey":       "master_key",
		"masterKeyCount":  "master_key_count",
		"secretKey":       "secret_key",
		"preference":      "preference",
		"authPublicData":  "auth_public_data",
		"authPrivateData": "auth_private_data",
		"createdTime":     "creation_time",
		"status":          "status",
		"runningVer":      "running_ver",
		"webConfig":       "web_config",
		"stageFlag":       "stage_flag",
		"qiniuAccount":    "qiniu_account",
	}
}

// GetStage return the stage of app
func (s Application) GetStage() string {
	switch s.StageFlag {
	case StageFlagSand:
		return StageSand
	case StageFlagProd:
		return StageProd
	default:
		return Unknown
	}
}

// GetStatus return the status of app
func (s Application) GetStatus() string {
	switch s.Status {
	case StatusFlagValid:
		return StatusValid
	case StatusFlagDeleted:
		return StatusDeleted
	default:
		return Unknown
	}
}

// QueryAppPrefixPayload - payload sent to Get Sandbox/Production App with App Prefix
type QueryAppPrefixPayload struct {
	ID     LikeCriterion `json:"ID"`
	Status int           `json:"status"`
}

// LikeCriterion - For using fuzzy search
type LikeCriterion struct {
	Like string `json:"$like"`
}

// UploadFile data structure
type UploadFile struct {
	AppID            string    `gorm:"column:app_id" json:"appID"`
	FIDRaw           int64     `gorm:"column:fid_raw" json:"-"`
	FID              string    `gorm:"column:fid" json:"fileID"`
	UID              string    `gorm:"column:uid" json:"userID"`
	Path             string    `json:"path"`
	Type             string    `json:"type"`
	Size             int64     `json:"size"`
	ModifyTime       time.Time `json:"modifyTime"`
	CreatedTime      time.Time `json:"createdTime"`
	FileDescription  string    `json:"fileDescription"`
	MD5              string    `gorm:"column:md5"`
	CDN              string    `gorm:"column:cdn"`
	CDNMap           string    `gorm:"column:cdn_map"`
	DfsGroup         int       `json:"dfsGroup"`
	DfsPath          string    `json:"dfsPath"`
	Status           int64     `json:"status"`
	StatusUpdateTime int64     `json:"statusUpdateTime"`
	ObjectID         string    `gorm:"column:object_id" json:"objectID"`
	CRC32AppID       int       `gorm:"column:crc32_app_id" json:"-"`
	ModRID           int       `gorm:"column:mod_r_id" json:"-"`
}

// TableName return the table name in database for gorm using
func (s UploadFile) TableName() string {
	return "baas.upload_file_mod50"
}

// FieldMap return the map between json field name and db field name
func (s UploadFile) FieldMap() map[string]string {
	return map[string]string{
		"appID":            "app_id",
		"fileID":           "fid_raw",
		"userID":           "uid",
		"path":             "path",
		"type":             "type",
		"size":             "size",
		"modifyTime":       "modify_time",
		"createdTime":      "created_time",
		"fileDescription":  "file_description",
		"MD5":              "md5",
		"CDN":              "cdn",
		"CDNMap":           "cdn_map",
		"dfsGroup":         "dfs_group",
		"dfsPath":          "dfs_path",
		"status":           "status",
		"statusUpdateTime": "status_update_time",
		"objectID":         "object_id",
	}
}

// BulkCreate data structure
type BulkCreateFile struct {
	Files []UploadFile `json:"files"`
}

// BulkUpdate data structure
type BulkUpdateFile struct {
	Files []map[string]interface{} `json:"files"`
}
