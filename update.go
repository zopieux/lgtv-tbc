package main

import "encoding/xml"

type SWUpdateRequest struct {
	XMLName       xml.Name `xml:"REQUEST"`
	Product       string   `xml:"PRODUCT_NM"`
	Model         string   `xml:"MODEL_NM"`
	Type          string   `xml:"SW_TYPE"`
	MajorVersion  string   `xml:"MAJOR_VER"`
	MinorVersion  string   `xml:"MINOR_VER"`
	Country       string   `xml:"COUNTRY"`
	CountryGroup  string   `xml:"COUNTRY_GROUP"`
	DeviceID      string   `xml:"DEVICE_ID"`
	AuthFlag      string   `xml:"AUTH_FLAG"`
	IgnoreDisable string   `xml:"IGNORE_DISABLE"`
	EcoInfo       string   `xml:"ECO_INFO"`
	ConfigKey     string   `xml:"CONFIG_KEY"`
	LanguageCode  string   `xml:"LANGUAGE_CODE"`
}

type SWUpdateResponse struct {
	XMLName            xml.Name `xml:"RESPONSE"`
	ResultCode         string   `xml:"RESULT_CD"`
	Message            string   `xml:"MSG"`
	RequestId          string   `xml:"REQ_ID"`
	ImageURL           string   `xml:"IMAGE_URL"`
	ImageSize          string   `xml:"IMAGE_SIZE"`
	ImageName          string   `xml:"IMAGE_NAME"`
	UpdateMajorVersion string   `xml:"UPDATE_MAJOR_VER"`
	UpdateMinorVersion string   `xml:"UPDATE_MINOR_VER"`
	ForceFlag          string   `xml:"FORCE_FLAG"`
	KE                 string   `xml:"KE"`
	TimeGTM            string   `xml:"GMT"`
	EcoInfo            string   `xml:"ECO_INFO"`
	CDNUrl             string   `xml:"CDN_URL"`
	Contents           string   `xml:"CONTENTS"`
}
