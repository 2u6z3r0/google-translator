package gtranslate

import (
	"net/http"
	"time"
)

var googleHostList = []string{"google.com"}

// come from http://m.news.xixik.com/content/25df0c00ac300bdc/
var googleHostAsia = []string{"google.com.hk", "google.mn", "google.co.kr", "google.co.jp", "google.com.vn", "google.la", "google.com.kh", "google.co.th", "google.com.my", "google.com.sg", "google.com.bn", "google.com.ph", "google.co.id", "google.kz", "google.kg", "google.com.tj", "google.co.uz", "google.tm", "google.com.af", "google.com.pk", "google.com.np", "google.co.in", "google.com.bd", "google.lk", "google.mv", "google.com.kw", "google.com.sa", "google.com.bh", "google.ae", "google.com.om", "google.jo", "google.co.il", "google.com.lb", "google.com.tr", "google.az", "google.am", "google.co.ls"}
var googleHostEurope = []string{"google.is", "google.dk", "google.no", "google.se", "google.fi", "google.ee", "google.lv", "google.lt", "google.ie", "google.co.uk", "google.gg", "google.je", "google.im", "google.fr", "google.nl", "google.be", "google.lu", "google.de", "google.at", "google.ch", "google.li", "google.pt", "google.es", "google.com.gi", "google.ad", "google.it", "google.com.mt", "google.sm", "google.gr", "google.ru", "google.com.by", "google.com.ua", "google.pl", "google.cz", "google.sk", "google.hu", "google.si", "google.hr", "google.ba", "google.me", "google.rs", "google.mk", "google.bg", "google.ro", "google.md"}
var googleHostAfrica = []string{"google.com.eg", "google.com.ly", "google.dz", "google.co.ma", "google.sn", "google.gm", "google.ml", "google.bf", "google.com.sl", "google.ci", "google.com.gh", "google.tg", "google.bj", "google.ne", "google.com.ng", "google.sh", "google.cm", "google.td", "google.cf", "google.ga", "google.cg", "google.cd", "google.it.ao", "google.com.et", "google.dj", "google.co.ke", "google.co.ug", "google.co.tz", "google.rw", "google.bi", "google.mw", "google.co.mz", "google.mg", "google.sc", "google.mu", "google.co.zm", "google.co.zw", "google.co.bw", "google.com.na", "google.co.za"}
var googleHostAtlantic = []string{"google.com.au", "google.com.nf", "google.co.nz", "google.com.sb", "google.com.fj", "google.fm", "google.ki", "google.nr", "google.tk", "google.ws", "google.as", "google.to", "google.nu", "google.co.ck", "google.com.do", "google.tt", "google.com.co", "google.com.ec", "google.co.ve", "google.gy", "google.com.pe", "google.com.bo", "google.com.py", "google.com.br", "google.com.uy", "google.com.ar", "google.cl"}
var googleHostAmerica = []string{"google.gl", "google.com.mx", "google.com.gt", "google.com.bz", "google.com.sv", "google.hn", "google.com.ni", "google.co.cr", "google.com.pa", "google.bs", "google.com.cu", "google.com.jm", "google.ht"}
var sw chan string

func init() {
	googleHostList = append(googleHostList, googleHostAsia...)
	googleHostList = append(googleHostList, googleHostEurope...)
	googleHostList = append(googleHostList, googleHostAfrica...)
	googleHostList = append(googleHostList, googleHostAtlantic...)
	googleHostList = append(googleHostList, googleHostAmerica...)
	sw = make(chan string, len(googleHostList))
	for _, v := range googleHostList {
		sw <- v
	}
}

// TranslationParams is a util struct to pass as parameter to indicate how to translate
type TranslationParams struct {
	From       string
	To         string
	Retry      int
	RetryDelay time.Duration
	GoogleHost string
	Client     *http.Client
}

// TranslateWithParams translate a text with simple params as string
func Translate(text string, params TranslationParams) (string, error) {
	var googleHost string
	if params.GoogleHost == "" {
		select {
		case googleHost = <-sw:
			defer func() { sw <- googleHost }()
		default:
			googleHost = "google.com"
		}
	} else {
		googleHost = params.GoogleHost
	}
	translated, err := translate(text, params.From, params.To, googleHost, true, params.Retry, params.RetryDelay, params.Client)
	if err != nil {
		return "", err
	}
	return translated, nil
}
