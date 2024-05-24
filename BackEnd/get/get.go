// get data from api by getData(url, &data)
package get

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/tidwall/gjson"
)

var (
	mixinKeyEncTab = []int{
		46, 47, 18, 2, 53, 8, 23, 32, 15, 50, 10, 31, 58, 3, 45, 35, 27, 43, 5, 49,
		33, 9, 42, 19, 29, 28, 14, 39, 12, 38, 41, 13, 37, 48, 7, 16, 24, 55, 40,
		61, 26, 17, 0, 1, 60, 51, 30, 4, 22, 25, 54, 21, 56, 59, 6, 63, 57, 62, 11,
		36, 20, 34, 44, 52,
	}
	cache          sync.Map
	lastUpdateTime time.Time
)

func Get(urlStr string, data any) error {
	newUrlStr, err := signAndGenerateURL(urlStr)
	if err != nil {
		return fmt.Errorf("generate url: %s", err)
	}
	req, err := http.NewRequest("GET", newUrlStr, nil)
	if err != nil {
		return fmt.Errorf("new request: %s", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Referer", "https://www.bilibili.com/")

	cookie, err := os.ReadFile("./myCookie.txt")
	if err != nil {
		return fmt.Errorf("open cookie file failed: %s", err)
	}
	req.Header.Set("cookie", string(cookie))

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %s", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %s", err)
	}
	err = json.Unmarshal(body, data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal body: %s", err)
	}
	data.(struct{ Code int }).Code = int(gjson.GetBytes(body, "code").Int())
	return nil
}

func GetConcur(urls []string, datap [](*any)) error { // 并发获取数据
	var wg sync.WaitGroup
	var allErr error

	for i, url := range urls {
		nowdatap := datap[i]
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			err := Get(url, nowdatap)
			if err != nil {
				allErr = fmt.Errorf("%s\nError get from %s: %s", allErr, url, err)
			}
		}(url)
	}

	wg.Wait()
	return allErr
}

func signAndGenerateURL(urlStr string) (string, error) {
	urlObj, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}
	imgKey, subKey := getWbiKeysCached()
	query := urlObj.Query()
	params := map[string]string{}
	for k, v := range query {
		params[k] = v[0]
	}
	newParams := encWbi(params, imgKey, subKey)
	for k, v := range newParams {
		query.Set(k, v)
	}
	urlObj.RawQuery = query.Encode()
	newUrlStr := urlObj.String()
	return newUrlStr, nil
}

func encWbi(params map[string]string, imgKey, subKey string) map[string]string {
	mixinKey := getMixinKey(imgKey + subKey)
	currTime := strconv.FormatInt(time.Now().Unix(), 10)
	params["wts"] = currTime

	// Sort keys
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Remove unwanted characters
	for k, v := range params {
		v = sanitizeString(v)
		params[k] = v
	}

	// Build URL parameters
	query := url.Values{}
	for _, k := range keys {
		query.Set(k, params[k])
	}
	queryStr := query.Encode()

	// Calculate w_rid
	hash := md5.Sum([]byte(queryStr + mixinKey))
	params["w_rid"] = hex.EncodeToString(hash[:])
	return params
}

func getMixinKey(orig string) string {
	var str strings.Builder
	for _, v := range mixinKeyEncTab {
		if v < len(orig) {
			str.WriteByte(orig[v])
		}
	}
	return str.String()[:32]
}

func sanitizeString(s string) string {
	unwantedChars := []string{"!", "'", "(", ")", "*"}
	for _, char := range unwantedChars {
		s = strings.ReplaceAll(s, char, "")
	}
	return s
}

func updateCache() {
	if time.Since(lastUpdateTime).Minutes() < 10 {
		return
	}
	imgKey, subKey, err := getWbiKeys()
	if err != nil {
		fmt.Println("error getting wbi keys: ", err)
		return
	}
	cache.Store("imgKey", imgKey)
	cache.Store("subKey", subKey)
	lastUpdateTime = time.Now()
}

func getWbiKeysCached() (string, string) {
	updateCache()
	imgKeyI, _ := cache.Load("imgKey")
	subKeyI, _ := cache.Load("subKey")
	return imgKeyI.(string), subKeyI.(string)
}

func getWbiKeys() (string, string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.bilibili.com/x/web-interface/nav", nil)
	if err != nil {
		return "", "", fmt.Errorf("Error creating request: %s", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Referer", "https://www.bilibili.com/")
	resp, err := client.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("Error sending request: %s", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("Error reading response: %s", err)
	}
	json := string(body)
	imgURL := gjson.Get(json, "data.wbi_img.img_url").String()
	subURL := gjson.Get(json, "data.wbi_img.sub_url").String()
	imgKey := strings.Split(strings.Split(imgURL, "/")[len(strings.Split(imgURL, "/"))-1], ".")[0]
	subKey := strings.Split(strings.Split(subURL, "/")[len(strings.Split(subURL, "/"))-1], ".")[0]
	return imgKey, subKey, nil
}
