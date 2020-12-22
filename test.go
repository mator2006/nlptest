package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func processGETRequest(appkey string, token string, text string, audioSaveFile string, format string, sampleRate int) {
	/**
	 * 设置HTTPS GET请求：
	 * 1.使用HTTPS协议
	 * 2.语音识别服务域名：nls-gateway.cn-shanghai.aliyuncs.com
	 * 3.语音识别接口请求路径：/stream/v1/tts
	 * 4.设置必须请求参数：appkey、token、text、format、sample_rate
	 * 5.设置可选请求参数：voice、volume、speech_rate、pitch_rate
	 */
	var url string = "https://nls-gateway.cn-shanghai.aliyuncs.com/stream/v1/tts"
	url = url + "?appkey=" + appkey
	url = url + "&token=" + token
	url = url + "&text=" + text
	url = url + "&format=" + format
	url = url + "&sample_rate=" + strconv.Itoa(sampleRate)
	// voice 发音人，可选，默认是xiaoyun。
	url = url + "&voice=" + "pt_paufibcpftsmflbp_mator"
	// volume 音量，范围是0~100，可选，默认50。
	// url = url + "&volume=" + strconv.Itoa(50)
	// speech_rate 语速，范围是-500~500，可选，默认是0。
	// url = url + "&speech_rate=" + strconv.Itoa(0)
	// pitch_rate 语调，范围是-500~500，可选，默认是0。
	// url = url + "&pitch_rate=" + strconv.Itoa(0)
	fmt.Println(url)
	/**
	 * 发送HTTPS GET请求，处理服务端的响应。
	 */
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("The GET request failed!")
		panic(err)
	}
	defer response.Body.Close()
	contentType := response.Header.Get("Content-Type")
	body, _ := ioutil.ReadAll(response.Body)
	if "audio/mpeg" == contentType {
		file, _ := os.Create(audioSaveFile)
		defer file.Close()
		file.Write([]byte(body))
		fmt.Println("The GET request succeed!")
	} else {
		// ContentType 为 null 或者为 "application/json"
		statusCode := response.StatusCode
		fmt.Println("The HTTP statusCode: " + strconv.Itoa(statusCode))
		fmt.Println("The GET request failed: " + string(body))
	}
}
func processPOSTRequest(appkey string, token string, text string, audioSaveFile string, format string, sampleRate int) {
	/**
	 * 设置HTTPS POST请求：
	 * 1.使用HTTPS协议
	 * 2.语音合成服务域名：nls-gateway.cn-shanghai.aliyuncs.com
	 * 3.语音合成接口请求路径：/stream/v1/tts
	 * 4.设置必须请求参数：appkey、token、text、format、sample_rate
	 * 5.设置可选请求参数：voice、volume、speech_rate、pitch_rate
	 */
	var url string = "https://nls-gateway.cn-shanghai.aliyuncs.com/stream/v1/tts"
	bodyContent := make(map[string]interface{})
	bodyContent["appkey"] = appkey
	bodyContent["text"] = text
	bodyContent["token"] = token
	bodyContent["format"] = format
	bodyContent["sample_rate"] = sampleRate
	// voice 发音人，可选，默认是xiaoyun。
	// bodyContent["voice"] = "xiaoyun"
	// volume 音量，范围是0~100，可选，默认50。
	// bodyContent["volume"] = 50
	// speech_rate 语速，范围是-500~500，可选，默认是0。
	// bodyContent["speech_rate"] = 0
	// pitch_rate 语调，范围是-500~500，可选，默认是0。
	// bodyContent["pitch_rate"] = 0
	bodyJson, err := json.Marshal(bodyContent)
	if err != nil {
		panic(nil)
	}
	fmt.Println(string(bodyJson))
	/**
	 * 发送HTTPS POST请求，处理服务端的响应。
	 */
	response, err := http.Post(url, "application/json;charset=utf-8", bytes.NewBuffer([]byte(bodyJson)))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	contentType := response.Header.Get("Content-Type")
	body, _ := ioutil.ReadAll(response.Body)
	if "audio/mpeg" == contentType {
		file, _ := os.Create(audioSaveFile)
		defer file.Close()
		file.Write([]byte(body))
		fmt.Println("The POST request succeed!")
	} else {
		// ContentType 为 null 或者为 "application/json"
		statusCode := response.StatusCode
		fmt.Println("The HTTP statusCode: " + strconv.Itoa(statusCode))
		fmt.Println("The POST request failed: " + string(body))
	}
}
func main() {
	var appkey string = "paufiBCpFTSmFLbp"
	var token string = "c46ad278b07649bab565d9a9b46ffa9c"
	var text string = "你在工作中可能会遇到同时给你说不要重复发明轮子，其实这个说的就是第一：不要做重复的事情，第二：站在巨人的肩膀上。现在有了互联网以及开源的精神，我们的很多在产品中需要实现的功能组件都可以在互联网上找到，我们可以直接拿来用，这样我们就不用重复做东西了，这得益于开源奉献精神"
	var textUrlEncode = text
	textUrlEncode = url.QueryEscape(textUrlEncode)
	textUrlEncode = strings.Replace(textUrlEncode, "+", "%20", -1)
	textUrlEncode = strings.Replace(textUrlEncode, "*", "%2A", -1)
	textUrlEncode = strings.Replace(textUrlEncode, "%7E", "~", -1)
	fmt.Println(textUrlEncode)
	var audioSaveFile string = "syAudio.wav"
	var format string = "wav"
	var sampleRate int = 16000
	processGETRequest(appkey, token, textUrlEncode, audioSaveFile, format, sampleRate)
	// processPOSTRequest(appkey, token, text, audioSaveFile, format, sampleRate)
}
