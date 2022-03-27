package crawler

import (
	"errors"
	"io"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

func getLinks(body io.Reader) (links []string) {
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			return links
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if "a" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}

				}
			}

		}
	}
}

func getTitle(body io.Reader) string {
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			return ""
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if "title" == token.Data {
				z.Next()
				return z.Token().Data
			}
		}
	}
}

func isExternal(s string, domain string) bool {
	if s == "" {
		return true
	}
	s = strings.Replace(s, "https://", "http://", -1)
	s = strings.ToLower(s)
	domain = strings.Replace(domain, "https://", "http://", -1)
	e := !(strings.Contains(s, domain) || s[0:1] == "/" || s[0:1] == "#")
	return e
}

func getDomain(u string) (string, error) {

	re := regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)`)

	submatchall := re.FindAllString(u, -1)
	for _, element := range submatchall {
		return strings.ToLower(element), nil
	}
	return "", errors.New("Broken entry point")
}

func containsUrl(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}
