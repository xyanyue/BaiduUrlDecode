func BaiduRealURL(urls string) string {
	u, err := url.Parse(urls)
	if err != nil {
    panic("URL 格式错误")
	}

	m, _ := url.ParseQuery(u.RawQuery)

	if murl, ok := m["url"]; ok && u.Host == "www.baidu.com" {
		urls = "https://" + u.Host + u.Path + "?url=" + murl[0]
		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
		res, err := client.Head(urls)
		if err != nil {
			panic("BAIDU_REAL_URL_HEAD_ERR错误")
		}
		return res.Header.Get("Location")
	}
	return ""
}
