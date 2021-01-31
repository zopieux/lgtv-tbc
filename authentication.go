// SPDX-License-Identifier: GPL-3.0-or-later

package main

type SdpAuthentication struct {
	InitServices InitServices `json:"initServices"`
}
type Authentication struct {
	DeviceSecret   string `json:"deviceSecret"`
	SessionID      string `json:"sessionID"`
	DeviceUniqueID string `json:"deviceUniqueID"`
}
type ActivationDate struct {
	UpdatedActivationDate string `json:"updatedActivationDate"`
	UpdatedMessage        string `json:"updatedMessage"`
}
type ItemList struct {
	Name   string `json:"name"`
	Switch string `json:"switch"`
}
type LogUnit struct {
	Name                 string `json:"name"`
	Switch               string `json:"switch"`
	RequestInterval      string `json:"requestInterval"`
	RetryIntervalOnError string `json:"retryIntervalOnError"`
	RetryIntervalOnFail  string `json:"retryIntervalOnFail"`
	RetryCount           string `json:"retryCount"`
}
type Whitelist struct {
	Checksum string `json:"checksum"`
	UpdateYn string `json:"updateYn"`
	URL      string `json:"url"`
	Version  int    `json:"version"`
}
type LoggingStatus struct {
	LogUnit       []LogUnit `json:"logUnit"`
	CheckInterval string    `json:"checkInterval"`
	Whitelist     Whitelist `json:"whitelist"`
}
type Notice struct {
	ID              string `json:"id"`
	AlertTypeFlag   string `json:"alertTypeFlag"`
	ToastTypeFlag   string `json:"toastTypeFlag"`
	LgStoreTypeFlag string `json:"lgStoreTypeFlag"`
	ExecuteAppID    string `json:"executeAppId"`
	StartDate       string `json:"startDate"`
	EndDate         string `json:"endDate"`
	Language        string `json:"language"`
	Title           string `json:"title"`
	Content         string `json:"content"`
	TitleContent    string `json:"titleContent"`
}
type NoticeList struct {
	LastUpdatedTime string   `json:"lastUpdatedTime"`
	Notice          []Notice `json:"notice"`
}
type Country struct {
	Code      string `json:"code"`
	Threecode string `json:"threecode"`
	RicCode   string `json:"ricCode"`
}
type App struct {
	ID             string  `json:"id"`
	Premium        bool    `json:"bPremium"`
	Name           string  `json:"name"`
	CategoryID     string  `json:"categoryId"`
	CategoryName   string  `json:"categoryName"`
	CurrencyCode   string  `json:"currencyCode"`
	Price          float64 `json:"price"`
	DisplayPrice   string  `json:"displayPrice"`
	Evaluation     float64 `json:"evaluationAverage"`
	AgeType        string  `json:"ageType"`
	IconColor      string  `json:"iconColor"`
	IconURL        string  `json:"iconURL"`
	StubAppVersion string  `json:"stubAppVersion"`
}
type BlockedApp struct {
	ID string `json:"id"`
}
type AppList struct {
	AppCount int   `json:"appCount"`
	AppList  []App `json:"appList"`
}
type BlockedAppList struct {
	AppCount int          `json:"appCount"`
	AppList  []BlockedApp `json:"appList"`
}
type EulaInfo struct {
	AdditionalSelectAll []string `json:"additionalSelectAll,omitempty"`
	EulaGroupName       string   `json:"eulaGroupName"`
	GeneralSelectAll    []string `json:"generalSelectAll,omitempty"`
	Overview            []string `json:"overview,omitempty"`
	SettingKey          string   `json:"settingKey"`
	Mandatory           []string `json:"mandatory,omitempty"`
	Notice              []string `json:"notice,omitempty"`
}
type EulaMappingList struct {
	MaxEulaMappingGroupID string     `json:"maxEulaMappingGroupId"`
	EulaInfo              []EulaInfo `json:"eulaInfo"`
}
type Eulas struct {
	AgreementNoticeFlag    string `json:"agreementNoticeFlag"`
	EulaFileName           string `json:"eulaFileName"`
	EulaFileURL            string `json:"eulaFileUrl"`
	EulaManagementTypeCode string `json:"eulaManagementTypeCode"`
	EulaManagementTypeName string `json:"eulaManagementTypeName"`
	EulaMandatoryFlag      string `json:"eulaMandatoryFlag"`
	EulaVersionID          string `json:"eulaVersionId"`
	RequiredEulaID         string `json:"requiredEulaID"`
}
type EulaList struct {
	Eulas             []Eulas `json:"eulas"`
	EulaMergedFileURL string  `json:"eulaMergedFileUrl"`
}
type LauncherPromotion struct {
	PromotionID    string `json:"promotionId"`
	PromotionCount int    `json:"promotionCount"`
}
type ServerDomain struct {
	URL     string `json:"url"`
	Version string `json:"version"`
}
type Services struct {
	ItemList                 []ItemList        `json:"itemList"`
	NeedCheckHardwareFeature string            `json:"needCheckHardwareFeature"`
	LoggingStatus            LoggingStatus     `json:"loggingStatus"`
	NoticeList               NoticeList        `json:"noticeList"`
	Country                  Country           `json:"country"`
	PremiumAppList           AppList           `json:"premiumAppList"`
	BlockedAppList           BlockedAppList    `json:"blockedAppList"`
	EulaMappingList          EulaMappingList   `json:"eulaMappingList"`
	EulaList                 EulaList          `json:"eulaList"`
	LauncherPromotion        LauncherPromotion `json:"launcherPromotion"`
	ServerDomain             ServerDomain      `json:"serverDomain"`
}
type DeviceFeatureResult struct {
	Fck               int `json:"fck"`
	RcmdCategoryCount int `json:"rcmdCategoryCount"`
}
type InitServices struct {
	Authentication      Authentication      `json:"authentication"`
	ActivationDate      ActivationDate      `json:"activationDate"`
	Services            Services            `json:"services"`
	DeviceFeatureResult DeviceFeatureResult `json:"deviceFeatureResult"`
}
