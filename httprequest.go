package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/storage"
)

var (
	// BaseURL 域名URL
	BaseURL string
	// LoginURL 登录URL
	LoginURL string
	// SaveURL 提交结果URL
	SaveURL string
	// MyUserAgent 模拟手机UA
	MyUserAgent string
)

func SetMainConfig(mainConfig MainConfig) {
	BaseURL = mainConfig.BaseURL
	LoginURL = mainConfig.LoginURL
	SaveURL = mainConfig.SaveURL
	MyUserAgent = mainConfig.MyUserAgent
	// fmt.Println(BaseURL, LoginURL, SaveURL, MyUserAgent)
}

// PostSaveForm 提交晨午检表单
func PostSaveForm(newClient *colly.Collector, config StudentConfig) {
	savePostForm := map[string]string{
		"province": config.Province,
		"city":     config.City,
		"area":     config.Area,
		"address":  config.Address,
		"szgjcs":   config.Szgjcs,
		"szcs":     config.Szcs,
		"szgj":     config.Szgj,
		"zgfxdq":   config.Zgfxdq,
		"mjry":     config.Mjry,
		"csmjry":   config.Csmjry,
		"tw":       config.Tw,
		"sfzx":     config.Sfzx,
		"sfcyglq":  config.Sfcyglq,
		"sfcxtz":   config.Sfcxtz,
		"sfjcbh":   config.Sfjcbh,
		"sfcxzysx": config.Sfcxzysx,
		"qksm":     config.Qksm,
		"sfyyjc":   config.Sfyyjc,
		"jcjgqr":   config.Jcjgqr,
		"remark":   config.Remark,
		"sfjcwhry": config.Sfjcwhry,
		"sfjchbry": config.Sfjchbry,
		"gllx":     config.Gllx,
		"glksrq":   config.Glksrq,
		"jcbhlx":   config.Jcbhlx,
		"jcbhrq":   config.Jcbhrq,
		"ismoved":  config.Ismoved,
		"bztcyy":   config.Bztcyy,
		"sftjhb":   config.Sftjhb,
		"sftjwh":   config.Sftjwh,
		"sfjcjwry": config.Sfjcjwry,
		"jcjg":     config.Jcjg,
	}
	err := newClient.Post(SaveURL, savePostForm)
	if err != nil {
		log.Fatal(2, err)
	}

}

// PostWX 向server酱服务端发送消息
func PostWX(rstText string, SCKEY string) {
	url := "https://sc.ftqq.com/" + SCKEY + ".send"
	method := "POST"
	payload := strings.NewReader("text=" + rstText + "&desp=119club np!")
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	//body, err := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))
}

// SignIn HTTP请求主要函数
func SignIn(config StudentConfig) {
	firstPostClient := colly.NewCollector()
	firstPostClient.UserAgent = MyUserAgent

	firstPostSuccessFlag := false
	if config.Cookie != "" {
		firstPostClient.OnRequest(func(request *colly.Request) {
			firstPostClient.SetCookies(BaseURL, storage.UnstringifyCookies(config.Cookie))
		})
		StandardLog(config.ID, "正在尝试使用cookie提交  ")
		firstPostClient.OnResponse(func(r *colly.Response) {
			message := ""
			tempResponse := UnmarshalHTTPResponse(r.Body)
			if tempResponse.M != "" {
				if tempResponse.M == "操作成功" {
					firstPostSuccessFlag = true
					message = "使用cookie提交成功"
					config.LastestUpdateTime = time.Now()
					UpdateConfig(config)
				} else if tempResponse.M == "您已上报过" {
					firstPostSuccessFlag = true
					message = "使用cookie时查询到本阶段已上报过"
				} else {
					firstPostSuccessFlag = true
					message = "使用cookie时返回的message为：" + tempResponse.M
				}
			}
			StandardLog(config.ID, message)
			// Server酱
			if config.SCKEY != "SCU89912...f4a70230" && config.SCKEY != "" {
				PostWX(message, config.SCKEY)
			}
		})
		PostSaveForm(firstPostClient, config)
	}

	if !firstPostSuccessFlag {
		firstLoginClient := firstPostClient.Clone()
		loginFlag := Login(firstLoginClient, strconv.Itoa(config.ID), config.Password)
		if loginFlag {
			StandardLog(config.ID, "登陆成功")
		} else {
			StandardLog(config.ID, "登陆失败")
			os.Exit(1)
		}

		secondPostClient := firstLoginClient.Clone()
		secondPostClient.OnResponse(func(response *colly.Response) {
			tempResponse := UnmarshalHTTPResponse(response.Body)
			message := ""
			if tempResponse.M != "" {
				if tempResponse.M == "操作成功" {
					message = "登陆后提交成功"
					newCookie := storage.StringifyCookies(secondPostClient.Cookies(response.Request.URL.String()))
					config.Cookie = newCookie
					config.LastestUpdateTime = time.Now()
					UpdateConfig(config)
				} else if tempResponse.M == "您已上报过" {
					message = "登陆后查询到本阶段已上报过"
				} else {
					message = "登陆后返回的message为：" + tempResponse.M
				}
			} else {
				message = "提交失败，或返回信息无法处理"
			}
			StandardLog(config.ID, message)
			// Server酱
			if config.SCKEY != "SCU89912...f4a70230" && config.SCKEY != "" {
				PostWX(message, config.SCKEY)
			}
		})
		PostSaveForm(secondPostClient, config)
	}

	StandardLog(0, "+++++")
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
