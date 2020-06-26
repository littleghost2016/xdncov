package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/storage"
)

const (
	// BaseURL 域名URL
	BaseURL = "https://xxcapp.xidian.edu.cn"
	// LoginURL 登录URL
	LoginURL = "https://xxcapp.xidian.edu.cn/uc/wap/login/check"
	// SaveURL 提交结果URL
	SaveURL = "https://xxcapp.xidian.edu.cn/ncov/wap/open-report/save"
	// MyUserAgent 模拟手机UA
	MyUserAgent = "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.45 Mobile Safari/537.36 Edg/84.0.522.20"
)

// PostSaveForm 提交晨午检表单
func PostSaveForm(newClient *colly.Collector, config StudentConfig) {
	savePostForm := map[string]string{
		"province": config.Province,
		"city":     config.City,
		"district": config.District,
		"address":  config.Address,
		"ymtys":    strconv.Itoa(config.Ymtys),
		"tw":       strconv.Itoa(config.Tw),
		"sfzx":     strconv.Itoa(config.Sfzx),
		"sfcyglq":  strconv.Itoa(config.Sfcyglq),
		"sfyzz":    strconv.Itoa(config.Sfyzz),
	}
	err := newClient.Post(SaveURL, savePostForm)
	if err != nil {
		log.Fatal(2, err)
	}
}

// SignIn HTTP主要步骤
func SignIn(config StudentConfig) {
	firstPostClient := colly.NewCollector()
	firstPostClient.UserAgent = MyUserAgent

	firstPostSuccessFlag := false
	if config.Cookie != "" {
		firstPostClient.OnRequest(func(request *colly.Request) {
			firstPostClient.SetCookies(BaseURL, storage.UnstringifyCookies(config.Cookie))
		})

	}
	firstPostClient.OnResponse(func(r *colly.Response) {
		tempResponse := UnmarshalHTTPResponse(r.Body)
		if tempResponse.M != "" {
			if tempResponse.M == "操作成功" || tempResponse.M == "您已上报过" {
				firstPostSuccessFlag = true
				fmt.Println(config.ID, "第一次提交即成功")
			}
		}
	})
	PostSaveForm(firstPostClient, config)

	if !firstPostSuccessFlag {
		firstLoginClient := firstPostClient.Clone()
		loginFlag := Login(firstLoginClient, strconv.Itoa(config.ID), config.Password)
		if loginFlag {
			fmt.Println(config.ID, "登陆成功")
		}

		secondPostClient := firstLoginClient.Clone()
		secondPostClient.OnResponse(func(response *colly.Response) {
			newCookie := storage.StringifyCookies(secondPostClient.Cookies(response.Request.URL.String()))
			config.Cookie = newCookie
			UpdateConfig(config)
			fmt.Println(string(response.Body))
		})
		PostSaveForm(secondPostClient, config)
	}
}

// Login 当持久化未能通过时，模拟登录以获得cookie
func Login(newClient *colly.Collector, id string, password string) (loginFlag bool) {
	loginFlag = false

	newClient.OnResponse(func(r *colly.Response) {
		tempResponse := UnmarshalHTTPResponse(r.Body)
		if tempResponse.M != "" {
			if tempResponse.M == "操作成功" {
				loginFlag = true
			}
		}
	})

	err := newClient.Post(LoginURL, map[string]string{
		"username": id,
		"password": password,
	})
	if err != nil {
		log.Fatal(1, err)
	}

	return
}
