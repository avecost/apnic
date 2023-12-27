package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type APNICResponse struct {
	RdapConformance []string `json:"rdapConformance,omitempty"`
	Notices         []struct {
		Title       string   `json:"title,omitempty"`
		Description []string `json:"description,omitempty"`
		Links       []struct {
			Value string `json:"value,omitempty"`
			Rel   string `json:"rel,omitempty"`
			Href  string `json:"href,omitempty"`
			Type  string `json:"type,omitempty"`
		} `json:"links,omitempty"`
	} `json:"notices,omitempty"`
	Country string `json:"country,omitempty"`
	Events  []struct {
		EventAction string    `json:"eventAction,omitempty"`
		EventDate   time.Time `json:"eventDate,omitempty"`
	} `json:"events,omitempty"`
	Name    string `json:"name,omitempty"`
	Remarks []struct {
		Description []string `json:"description,omitempty"`
		Title       string   `json:"title,omitempty"`
	} `json:"remarks,omitempty"`
	Links []struct {
		Value string `json:"value,omitempty"`
		Rel   string `json:"rel,omitempty"`
		Href  string `json:"href,omitempty"`
		Type  string `json:"type,omitempty"`
	} `json:"links,omitempty"`
	Status          []string `json:"status,omitempty"`
	Type            string   `json:"type,omitempty"`
	EndAddress      string   `json:"endAddress,omitempty"`
	IPVersion       string   `json:"ipVersion,omitempty"`
	StartAddress    string   `json:"startAddress,omitempty"`
	ObjectClassName string   `json:"objectClassName,omitempty"`
	Handle          string   `json:"handle,omitempty"`
	Entities        []struct {
		Roles  []string `json:"roles,omitempty"`
		Events []struct {
			EventAction string    `json:"eventAction,omitempty"`
			EventDate   time.Time `json:"eventDate,omitempty"`
		} `json:"events,omitempty"`
		Links []struct {
			Value string `json:"value,omitempty"`
			Rel   string `json:"rel,omitempty"`
			Href  string `json:"href,omitempty"`
			Type  string `json:"type,omitempty"`
		} `json:"links,omitempty"`
		VcardArray      []any  `json:"vcardArray,omitempty"`
		ObjectClassName string `json:"objectClassName,omitempty"`
		Handle          string `json:"handle,omitempty"`
		Remarks         []struct {
			Description []string `json:"description,omitempty"`
			Title       string   `json:"title,omitempty"`
		} `json:"remarks,omitempty"`
	} `json:"entities,omitempty"`
	Cidr0Cidrs []struct {
		V4Prefix string `json:"v4prefix,omitempty"`
		Length   int    `json:"length,omitempty"`
	} `json:"cidr0_cidrs,omitempty"`
	Port43 string `json:"port43,omitempty"`
}

func IPwhois(ip string) (whois string, err error) {
	var respToJSON APNICResponse

	resp, err := http.Get("https://rdap.apnic.net/ip/" + ip)
	if err != nil {
		return "", err
	}

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(respData, &respToJSON)
	if err != nil {
		return "", err
	}

	ownerName := respToJSON.Name
	ownerDesc := ""
	telcoName := ""
	if len(respToJSON.Remarks) > 0 {
		ownerDesc = strings.Join(respToJSON.Remarks[0].Description, "")
		telcoName = fmt.Sprintf("%s - %s %s", respToJSON.Country, ownerDesc, ownerName)
	} else {
		telcoName = ownerName
	}

	return telcoName, nil
}

func main() {
	d, err := IPwhois("8.8.8.8")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(d)
}
