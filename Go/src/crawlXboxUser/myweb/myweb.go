package myweb

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
	"strings"
	"sync"
	"io/ioutil" 
	"../myfile"
)

type CrawlJob struct{
	Url string
	File string
}

func Crawl( urls string, client *http.Client ) (*http.Response, error) {
	time.Sleep(1000 * time.Millisecond)
	
	req, err := http.NewRequest("GET", urls, nil)
	req.Header.Add("Authorization", "XBL3.0 x=981467893410297204;eyJlbmMiOiJBMTI4Q0JDK0hTMjU2IiwiYWxnIjoiUlNBLU9BRVAiLCJjdHkiOiJKV1QiLCJ6aXAiOiJERUYiLCJ4NXQiOiIxZlVBejExYmtpWklFaE5KSVZnSDFTdTVzX2cifQ.HSWGM91Ylj5IeiGOtawoHylELTJJo6IimzrJk6kfBXsD5mqEAo1E5nIzYiLXEY_lgwTXTa7Y1yxJxrC6uUgxnrJ8-iHzH42sK7OGf8Nh2zBJP114Z92GClZdxYYCNTbhZ3GzdQcdZaIbnQDILKXCl9U1BuxYgKbkpcfR7r5Xf6k.1sZ7aeXAZMSUit-r7xr7SQ.8YuShgmKDryKwm2rAzDAydn-yRncqz4J_lloEm0LxO7uAq8DOY4_8gdN5w2fAdbHkY4U3FTC82BN-RmJmjxHXTxprAAWcWS2MRYEeoBK0RWGbvVWsHgSSoDttFuk0O3vy1fVLriepihKLkB_kLeZADaKvBOwgqN__ruld-9kQ3O5-wApI89-uFT1xniyHn1ba1yvyj6Gl-geNeVWeKhe123uyclWrsANYLazMXivt54oY_nXGfd8ChPHZvgndwq-Z6_JQxBziYUf8CC0zbQu7S4KnriKmg2xPJEzn3-JvxJCpNs1EnJGJse4rqv6rkk0XK2PtgPNeHGUz2mmBcJJHG5LROMRA2XI2HzCLa5a05IJFfrjHl6erYbqVWH9-8s5shtzp9bJf1JbDzlEx5fjNBlsmbsH5YNA-h19mNHr-pQS89qwiL0grkN7Tv-7UHxTKIYS12DNVyDuwVygKQIj3WKHyhclpOeGiVRRZbbr1e-4gtjVEkyWVS0Au3z1GcAbSLjLPZpQ1EDMeImD2vaTk5Irx2N6dz-zpWuedVeHAybo5WaHt-kI9Hg_uW99Rt4totf28r_Wa9XvLrY3S6hKv9GLZB_OpO6dbIlXi08BitIT1FrO5tKUbdn1iBsnLdDY37fCvT5yj9ksmub4Jp3cH9_lV8PyC49USbiNg6lUvxOgqO13KJJtL1oM_JHYQG1hY_mo7jg5Pp_7iSxqrbYNmZfJ_5ElWgJsUqYQXiFvjjBk0SOehqhn6R_w_yrIQ_-ElVNzu7u5KiWUsuMwaO-rRxeV8hiIn_ywp-MIp-iA1Dvxs3SGIqMKF__meCk8fhfG03qequy_GqMtiTF2H5DLZT-p_YP-c66XY10SXc7vdC-VnhVVMG7qQJGz3my6zP-iRoeyUN5PIHskT05h2PT4gphaaBisHUPWPJwtWENtw-khy5lRrirykKnIt1N662d8T8HAkaajQm0kLKGlWeYQ0cegMNRtRbZl2Cl67wp0TXLm7gqIZW86ZwIlhOzHN8bv1adBs94mIPCeza-tEmltB2kGM93HJa9aim03enSURTxpWPGjY7JES-Uv2_KHRBRHHdH8JdLSuc6wusLjmGlA5ysYxbdsXoZSJvOYOt9BeBkKGraVjh-pa9yUYiKTGzM8TPhZFU6F4wkoWtWaquOlkQrXkMlfge5F8zmiz4dqW9bU93Um2Q-vp82mH3uRxXzoRbWHxOxG0deQwqVunaxJug.IQly0XDjezOrNQu8j1EZk6uN0fmpmhxskbV35BVq87I")
//									 XBL3.0 x=8652283878512036862 ;eyJlbmMiOiJBMTI4Q0JDK0hTMjU2IiwiYWxnIjoiUlNBLU9BRVAiLCJjdHkiOiJKV1QiLCJ6aXAiOiJERUYiLCJ4NXQiOiIxZlVBejExYmtpWklFaE5KSVZnSDFTdTVzX2cifQ.FUVQ941OdSL7AMQialNvJHQCAvoRDOy2DMi4A-gpnMU3gmuit86sgid2ZBMYslMszJwgyg_I5pnA5AEtY2LjY7vL02fsaOFW4Ocg90R3AoIoQ3HK8uHrdXM5FkhaUe3lUg1A5SjYx5zH57pnWCx8ZgEQPpXtOidjxqWc-MyR8ss.HWiXS2epjqJpEND-XCMNUQ.k4a6begQazcibe5cF_tZOicxDFJmeJYi5xanPeNf-qL_G-KNoUBsk8RQQwyQ7rc_FTc5TIDis2neulT8xp8-MrrDAsHs4ax7KyrAz2x-LAmYByCxkdb8zZ1pRWCQWeQsrnnNIFWhNruTRmDs_zAa3VusVIh4owZXbThyBsHPGWrodAzElLrjNRNsY5KKEXxxUyw-YA4uqU2xMP_bNZPetyppxd6DvB7zm4o4KWiJPiHS1cucL6Acp5kZNYpDKjO8-sEcqTg_oK_lYnSu3fFdq55xSaISc37qO2_m0oJsqhSZA6zn16-7MNgFuTHCyhzDn3T6jtalTuQzBvYyZnDYE51D2EQa4CWlf3rX7i0MteKG3-Y6Wm36fqH6n_MxKbQ-YrnUUUogUw4Wm16gwBVG6W3jGYq-QT55B1QGdsYwiig2tOugeu0SaBXReV-Sz2tkIzuBJy2-2-LOxVUPkcVO_jlpEhS4fMZRl9tWIMqRqL4MhJt9TyZda9CShyZyrq-eo8gP09r9u9GkKAWNyIbMrTDArQ4ZQ2cbxTEsXpj0vIKsbXpN9fnyOsV78Fk__W-uw5Y3k9rn-Nl3xuMdHPo0EA0iEl7e1YInqvQbx1MXHXVjbSsd0dR3o8bO9eZlgP5R4DoaPinL9X99OQ790g4iQuAiJ5MCUSQMH84iiI-zd4OpIIQ1btRjHPt0vPDKSllEPnOdFTpPB-NLkXzCdfNPq6rGi8wYLpXw9v-3MbJCjfpWl0RfSh4Q3-qFJRPIk_2_jnuGRyLSyMn4SKd8elWRWNMKAIWFTuAA8UqIzdVsUsieQthr5R6YUIVQ7x5tcSNPM9TfIrw2lXrtLd3xeLBMT3q7vN8LLIUfWasK7B45o54sOspF2XhF5_o1QOuVLC6aVs9M_Z7Yn8ICqyX4CTtlze1qCAC8Xh4okDOCX28AB6zShmTAGtdkhP6EImhwmXSAA550qc44oATgEn34E7oGyJ4IIuAObfgBBSFktg_wmww5jYFaQ-OrtweG4h-XlewJu6TQQF02VDI3fpRZbBxFY1-bJ_chG2zIr13RPI2eUg3ktHL1VkPQWlbohMnOyzh067-1rqAokOjWy5iuTrcGO5OS-cdNDJT1VufAuEXSPBH57CmIDESEQo13G5mEqYLEHMZJgHAv153U3_ayZPY_BxDcC2twiquY7LbBxYIK3lcwoqjTYmPoG5bZ2tXlVt9W2Xg17sSqgNMw_poglFCzcw.lCR4-mlD2tKFKES2o3Y_xGB3WWFm6rqLte4JkRBXmRE
//																   eyJlbmMiOiJBMTI4Q0JDK0hTMjU2IiwiYWxnIjoiUlNBLU9BRVAiLCJjdHkiOiJKV1QiLCJ6aXAiOiJERUYiLCJ4NXQiOiIxZlVBejExYmtpWklFaE5KSVZnSDFTdTVzX2cifQ.FUVQ941OdSL7AMQialNvJHQCAvoRDOy2DMi4A-gpnMU3gmuit86sgid2ZBMYslMszJwgyg_I5pnA5AEtY2LjY7vL02fsaOFW4Ocg90R3AoIoQ3HK8uHrdXM5FkhaUe3lUg1A5SjYx5zH57pnWCx8ZgEQPpXtOidjxqWc-MyR8ss.HWiXS2epjqJpEND-XCMNUQ.k4a6begQazcibe5cF_tZOicxDFJmeJYi5xanPeNf-qL_G-KNoUBsk8RQQwyQ7rc_FTc5TIDis2neulT8xp8-MrrDAsHs4ax7KyrAz2x-LAmYByCxkdb8zZ1pRWCQWeQsrnnNIFWhNruTRmDs_zAa3VusVIh4owZXbThyBsHPGWrodAzElLrjNRNsY5KKEXxxUyw-YA4uqU2xMP_bNZPetyppxd6DvB7zm4o4KWiJPiHS1cucL6Acp5kZNYpDKjO8-sEcqTg_oK_lYnSu3fFdq55xSaISc37qO2_m0oJsqhSZA6zn16-7MNgFuTHCyhzDn3T6jtalTuQzBvYyZnDYE51D2EQa4CWlf3rX7i0MteKG3-Y6Wm36fqH6n_MxKbQ-YrnUUUogUw4Wm16gwBVG6W3jGYq-QT55B1QGdsYwiig2tOugeu0SaBXReV-Sz2tkIzuBJy2-2-LOxVUPkcVO_jlpEhS4fMZRl9tWIMqRqL4MhJt9TyZda9CShyZyrq-eo8gP09r9u9GkKAWNyIbMrTDArQ4ZQ2cbxTEsXpj0vIKsbXpN9fnyOsV78Fk__W-uw5Y3k9rn-Nl3xuMdHPo0EA0iEl7e1YInqvQbx1MXHXVjbSsd0dR3o8bO9eZlgP5R4DoaPinL9X99OQ790g4iQuAiJ5MCUSQMH84iiI-zd4OpIIQ1btRjHPt0vPDKSllEPnOdFTpPB-NLkXzCdfNPq6rGi8wYLpXw9v-3MbJCjfpWl0RfSh4Q3-qFJRPIk_2_jnuGRyLSyMn4SKd8elWRWNMKAIWFTuAA8UqIzdVsUsieQthr5R6YUIVQ7x5tcSNPM9TfIrw2lXrtLd3xeLBMT3q7vN8LLIUfWasK7B45o54sOspF2XhF5_o1QOuVLC6aVs9M_Z7Yn8ICqyX4CTtlze1qCAC8Xh4okDOCX28AB6zShmTAGtdkhP6EImhwmXSAA550qc44oATgEn34E7oGyJ4IIuAObfgBBSFktg_wmww5jYFaQ-OrtweG4h-XlewJu6TQQF02VDI3fpRZbBxFY1-bJ_chG2zIr13RPI2eUg3ktHL1VkPQWlbohMnOyzh067-1rqAokOjWy5iuTrcGO5OS-cdNDJT1VufAuEXSPBH57CmIDESEQo13G5mEqYLEHMZJgHAv153U3_ayZPY_BxDcC2twiquY7LbBxYIK3lcwoqjTYmPoG5bZ2tXlVt9W2Xg17sSqgNMw_poglFCzcw.lCR4-mlD2tKFKES2o3Y_xGB3WWFm6rqLte4JkRBXmRE
	req.Header.Add("x-xbl-contract-version", "2")
	resp, err := client.Do(req)

	defer func() {
		if err:=recover(); err!=nil{
			fmt.Printf( "got err: %v", err )
		}
	}()

	return resp, err
}

func GetCookieToken( urls string, client *http.Client ) (*http.Response, error) {
	time.Sleep(1000 * time.Millisecond)
	
	req, err := http.NewRequest("POST", urls, nil)

	//req.SetBasicAuth("api", "key-3ax6xnjp29jd6fds4gc373sgvjxteol0")
	//生成Form 因为Form是个Values对象来着
	form := url.Values{
		"NAPExp" : {`Wed, 29-Jul-2015 17:18:07 GMT`},
		"NAP" : {`V%3D1.9%26E%3D1079%26C%3DW-RGts2-XiGqxo-CMFGJiao_0Y4MAu3RhNfRAM1WIuLc50QDE5p4FA%26W%3D1`},
		"ANON" : {`A%3D9F24D14E9D78435FA0578B04FFFFFFFF%26E%3D10d3%26W%3D1`},
		"ANONExp" : {`Fri, 06-Nov-2015 18:18:07 GMT`},
		"t" : {`EgCSAQMAAAAEgAAAC4AA668myyn5owsKIUAQAER/GVJQBuBfF9MOhMZtXbTMbf7J5++Jy61G/GfiQI3QIZ3dppPCxJ4+XbWI1X/4tEbSYZSn4iBeiIcNjXV84KC4z+jUNFGq3knhaqHICLMMxZVfS/z7NbjnXrSN7hCSyam3ouuHMBCOQGlpR0Y0vVEwSeYBAWcAAQEAAAMAy0gu0lzSNFVc0jRVv3YEAAoTIAAQEQA4NzUyMzY4MjFAcXEuY29tAFgAAB04NzUyMzY4MjElcXEuY29tQHBhc3Nwb3J0LmNvbQAAAAFDTgAGMTAwMDg0AAAAAAQJAgAAhd9tQBAGQQAEV2FuZwAHSHVhcWluZwAAAAAAAAAAAAAAAAAAjVPCvJY0PF4AAFzSNFVeeatVAAAAAAAAAAAAAAAAEAAyMjEuMjIwLjI1My4xNDIAAgEAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAObJzfQjoD7eAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA`}}
	//对form进行编码
	req.Body = ioutil.NopCloser(strings.NewReader(form.Encode()))
	
	req.Header.Add("Cookie", `optimizelySegments=%7B%222409710566%22%3A%22gc%22%2C%222444610318%22%3A%22false%22%2C%222452430334%22%3A%22none%22%2C%222453730225%22%3A%22search%22%7D; optimizelyEndUserId=oeu1429524900277r0.7018708982504904; optimizelyBuckets=%7B%7D; UtcOffsetMinutes=480; s_pers=%20s_vnum%3D1461060901327%2526vn%253D1%7C1461060901327%3B%20s_lastvisit%3D1429524901334%7C1524132901334%3B%20s_nr%3D1429524902073%7C1461060902073%3B%20s_invisit%3Dtrue%7C1429526702078%3B; s_sess=%20s_ria%3Dflash%252016%257Csilverlight%2520not%2520detected%3B%20s_cc%3Dtrue%3B%20s_sq%3Dmsxboxcomzhcn%252Cmsxboxcomglobal%253D%252526pid%25253Dwww%2525252F%252526pidt%25253D1%252526oid%25253Dhttps%2525253A%2525252F%2525252Faccount.xbox.com%2525252FAccount%2525252FSignin%2525253FreturnUrl%2525253Dhttp%252525253a%252525252f%252525252fwww.xbox.com%252525252fzh-CN%252525252f%252526ot%25253DA%3B; X1WY=3MKccOUvcxFCu3zNK5tD5fOODZMcswLPfkGU6o3U8jo0V172KJVTMFg; s_vi=[CS]v1|2A9A69278519447C-6000060840001E28[CE]; defCulture=zh-CN; MUID=10D3E56336FC66813FC0E2EC32FC67A0; graceIncr=0`)
	req.Header.Add("Referer", `https://login.live.com/ppsecure/post.srf?wa=wsignin1.0&rpsnv=12&ct=1429525071&rver=6.2.6289.0&wp=MBI_SSL&wreply=https:%2F%2Faccount.xbox.com:443%2Fpassport%2FsetCookies.ashx%3Frru%3Dhttps%253a%252f%252faccount.xbox.com%252fzh-CN%252fAccount%252fSignin%253freturnUrl%253dhttp%25253a%25252f%25252fwww.xbox.com%25252fzh-CN%25252f&lc=2052&id=292543&cbcxt=0&bk=1429525073&uaid=cd4fa222f2544c36839f6dbe6d7c40dc`)
	req.Header.Add("Origin", `https://login.live.com`)
	req.Header.Add("Host", `account.xbox.com`)
	req.Header.Add("Content-Type", `application/x-www-form-urlencoded`)
	req.Header.Add("Content-Length", `719`)
	req.Header.Add("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8`)
	req.Header.Add("Accept-Encoding", `gzip, deflate`)
	req.Header.Add("Accept-Language", `zh-CN,zh;q=0.8`)
	req.Header.Add("Connection", `keep-alive`)
	req.Header.Add("Cache-Control", `max-age=0`)
	req.Header.Add("User-Agent", `Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.94 Safari/537.36`)
//	req.Header.Add("Transfer-Encoding", "chunked")
	req.ContentLength = int64(strings.NewReader(form.Encode()).Len())
	resp, err := client.Do(req)
	fmt.Printf( "%+v\n", resp )
//	cookies := resp.Header.Get("Cache-Control")
	cookies := resp.Cookies()
	fmt.Printf( "%+v\n", cookies )

	defer func() {
		if err:=recover(); err!=nil{
			fmt.Printf( "got err: %v", err )
		}
	}()

	return resp, err
}


func CrawlAndSaveWorker( linkChannel chan CrawlJob, client *http.Client, rootPath string, wg *sync.WaitGroup){
	defer wg.Done()

	for job := range linkChannel{
		if myfile.FileExists( job.File ){
			continue
		}

		resp, err := Crawl( job.Url, client )
		if err != nil {
			fmt.Printf( "GET error %v\n", err )
		}
		if resp == nil{
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)  
		if err != nil {
		 	fmt.Println("Read error is: %v\n", err)
		}
		
		myfile.SaveFile( job.File, body )
		fmt.Printf( "Url:%s, Size:%d \n", job.Url, len(body) )
		resp.Body.Close()

	}
}