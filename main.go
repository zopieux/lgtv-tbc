// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	addr   = flag.String("addr", "", "Address to listen on")
	notify = flag.Bool("notify", false, "Whether to send a notification at startup")

	httpClient = http.Client{
		Timeout:   time.Second * 4,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}
)

const (
	// Chosen by fair dice roll.
	deviceSecret   = "ucYDZmdTK4ixk/hWxgoD1KFA0IgGR1FIbA2gsQ8nwuqFmqLyV5MhPtOIDaL1iBc95WXpb0xcOwfN1DVAVyL04L0zvQmkiVc9Tj0NDyj3EehwsVLq29zgwCXqIiQerLmtkiP/FKSgI/9EBE4h7wBpRA=="
	sessionID      = "87KVz5+Pl7QX1caiwJ9wirvV3qE="
	deviceUniqueId = "97bdf14c9d8f02e92e5a381e0c414bef17e9ae74c2c80ee3b7a7f1c4bf2d524e1a6e7f4175a8d1a1f2e7114ea39edc6764f5c6ef9cc2ff23f69c004758d87b9a"
)

var (
	premiumAppList = []App{
		{
			ID:             "netflix",
			Premium:        true,
			Name:           "Netflix",
			CategoryID:     "003",
			CategoryName:   "Entertainment",
			CurrencyCode:   "USD",
			Price:          0,
			DisplayPrice:   "$0.00",
			Evaluation:     8.01,
			AgeType:        "0",
			IconColor:      "#ffffff",
			IconURL:        "http://ngfts.lge.com/fts/gftsDownload.lge?biz_code=APP_STORE&func_code=APP_ICON&file_path=/appstore/app/icon/20150421/21790820472215287LARGE_APP_ICON_130x130_webos.png",
			StubAppVersion: "0.0.4",
		},
		{
			ID:             "tv.twitch.tv.starshot.lg",
			Premium:        true,
			Name:           "Twitch",
			CategoryID:     "003",
			CategoryName:   "Entertainment",
			CurrencyCode:   "USD",
			Price:          0,
			DisplayPrice:   "$0.00",
			Evaluation:     7.42,
			AgeType:        "16",
			IconColor:      "#9146ff",
			IconURL:        "http://ngfts.lge.com/fts/gftsDownload.lge?biz_code=APP_STORE&func_code=APP_ICON&file_path=/appstore/app/icon/20210111/35910902.png",
			StubAppVersion: "0.0.14",
		},
	}

	blockedAppList = []BlockedApp{}

	noticeIndex uint64 = 420000
)

func generateAuthentication() *SdpAuthentication {
	auth := &SdpAuthentication{
		InitServices: InitServices{
			Authentication: Authentication{
				DeviceSecret:   deviceSecret,
				SessionID:      sessionID,
				DeviceUniqueID: deviceUniqueId,
			},
			ActivationDate: ActivationDate{
				UpdatedActivationDate: "N",
				UpdatedMessage:        "0",
			},
			Services: Services{
				LoggingStatus: LoggingStatus{
					LogUnit: []LogUnit{
						{
							Name:                 "normal_log",
							Switch:               "on",
							RequestInterval:      "600",
							RetryIntervalOnError: "0",
							RetryIntervalOnFail:  "0",
							RetryCount:           "0",
						},
					},
					CheckInterval: "3600",
					Whitelist: Whitelist{
						Checksum: "01010101010101010101010101010101",
						UpdateYn: "Y",
						URL:      "http://sne.lge.com/fake-whitelist",
						Version:  41,
					},
				},
				Country: Country{
					Code:      "CH",
					Threecode: "CHE",
					RicCode:   "EIC",
				},
				PremiumAppList: AppList{
					AppCount: len(premiumAppList),
					AppList:  premiumAppList,
				},
				BlockedAppList: BlockedAppList{
					AppCount: len(blockedAppList),
					AppList:  blockedAppList,
				},
				EulaMappingList: EulaMappingList{
					MaxEulaMappingGroupID: "GR00000632",
					EulaInfo:              []EulaInfo{},
				},
				EulaList: EulaList{
					Eulas:             []Eulas{},
					EulaMergedFileURL: "http://ngfts.lge.com/fts/gftsDownload.lge?biz_code=MEMBERSHIP&func_code=TERMS&file_path=/terms/comp/201808140003_ch_comp.zip",
				},
				LauncherPromotion: LauncherPromotion{
					PromotionID:    "",
					PromotionCount: 0,
				},
				ServerDomain: ServerDomain{
					URL:     "http://ngfts.lge.com/domainurl/webOS4.0/v1.0/server_domain_list_v4.0.json",
					Version: "1.0",
				},
			},
			DeviceFeatureResult: DeviceFeatureResult{
				Fck:               154,
				RcmdCategoryCount: 0,
			},
		},
	}
	for _, service := range []string{
		"english_nlp", "wedge_service", "mycontent_service", "mycontent_youtube", "magic_tips", "magic_tips_youtube",
		"magic_tips_web", "membership_service", "alibaba_genie", "amazon_echo", "channel_plus", "efs_service",
		"freeview_play", "french_nlp", "google_assistant", "google_home", "magic_tips_genre", "magic_tips_person",
		"magic_tips_server", "mycontent_movie", "mycontent_tvshow", "naver_clova", "personalized_recommendation",
	} {
		auth.InitServices.Services.ItemList = append(auth.InitServices.Services.ItemList, ItemList{
			Name:   service,
			Switch: "off",
		})
	}
	if *notify {
		auth.InitServices.Services.NoticeList.LastUpdatedTime = strconv.FormatInt(time.Now().Unix(), 10)
		content := "Hi there! “LG TV: Take Back Control” works just fine.<br>Have a nice day!"
		auth.InitServices.Services.NoticeList.Notice = []Notice{
			{
				ID:              fmt.Sprintf("NT%09d", noticeIndex),
				AlertTypeFlag:   "N",
				ToastTypeFlag:   "Y",
				LgStoreTypeFlag: "N",
				ExecuteAppID:    "com.webos.app.notificationcenter",
				StartDate:       time.Now().Format("2006.01.02"),
				EndDate:         strconv.FormatInt(time.Now().Add(time.Hour*24).Unix(), 10),
				Language:        "en-US",
				Title:           content,
				Content:         content,
				TitleContent:    content,
			},
		}
		noticeIndex++
	}
	return auth
}

func generateInitServicesResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if bytes, err := json.Marshal(generateAuthentication()); err == nil {
		w.Write(bytes)
		log.Print("Sent back no-bloatware /initservices reply")
	} else {
		w.WriteHeader(500)
		log.Printf("Error marshalling: %v", err)
	}
}

func generateUpdateResponse(w http.ResponseWriter, r *http.Request) {
	timeLoc, err := time.LoadLocation("GMT")
	if err != nil {
		log.Printf("Could not load GTM TZ: %v", err)
		w.WriteHeader(500)
		return
	}
	updateResponse := SWUpdateResponse{
		ResultCode:         "900",
		Message:            "Success",
		RequestId:          "00000000080808080808",
		TimeGTM:            time.Now().In(timeLoc).Format("02 Jan 2006 15:04:05 GMT"),
		EcoInfo:            "01",
		ImageURL:           "",
		ImageSize:          "",
		ImageName:          "",
		UpdateMajorVersion: "",
		UpdateMinorVersion: "",
		ForceFlag:          "",
		KE:                 "",
		CDNUrl:             "",
		Contents:           "",
	}
	var xmlBytesResp []byte
	if xmlBytesResp, err = xml.Marshal(updateResponse); err != nil {
		log.Printf("Could not encode XML response: %v", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-type", "application/octet-stream;charset=UTF-8")
	w.Write([]byte(base64.StdEncoding.EncodeToString(xmlBytesResp)))
	log.Printf("Sent back 'already up-to-date' /CheckSWAutoUpdate.laf reply")
}

func handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	url := r.Header.Get("X-Forwarded-Proto") + "://" + r.Header.Get("X-Forwarded-Host") + r.URL.String()
	log.Printf("Request: %s %v | %v", r.Method, url, r.Header)

	// Prevent accidental DNS poisoning loops.
	if strings.HasPrefix(r.Host, "127.") || strings.HasPrefix(r.Host, "localhost") {
		w.WriteHeader(400)
		return
	}

	// Milliseconds.
	w.Header().Set("X-Server-Timer", strconv.FormatInt(time.Now().UnixNano()/1e6, 10))

	if r.Method == "POST" && r.URL.Path == "/rest/sdp/v9.0/initservices" {
		generateInitServicesResponse(w, r)
		return
	}

	if r.Method == "POST" && (
		r.URL.Path == "/CheckSWAutoUpdate.laf" ||
			r.URL.Path == "/CheckSWManualUpdate.laf") {
		generateUpdateResponse(w, r)
		return
	}

	request, err := http.NewRequestWithContext(context.Background(), r.Method, url, r.Body)
	if err != nil {
		log.Printf("Error creating upstream request: %v", err)
		w.WriteHeader(500)
		return
	}
	response, err := httpClient.Do(request)
	if err != nil {
		log.Printf("Error on upstream request: %v", err)
		w.WriteHeader(500)
		return
	}
	log.Printf("Upstream response: %d | %v", response.StatusCode, response.Header)
	w.WriteHeader(response.StatusCode)
	for key, value := range response.Header {
		if key == "X-Real-IP" || key == "Connection" || strings.HasPrefix(key, "X-Forward") {
			continue
		}
		for _, v := range value {
			w.Header().Add(key, v)
		}
	}
	if _, err := io.Copy(w, response.Body); err != nil {
		log.Printf("Error copying response bytes: %v", err)
	}
}

func main() {
	flag.Parse()
	log.Printf("Listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, http.HandlerFunc(handler)))
}
