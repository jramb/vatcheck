package main

/*
	2016 by J Ramb
*/

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type VatQuery struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  string   `xml:"Header"`
	Body    struct {
		CheckVat struct {
			XMLName xml.Name `xml:"urn:ec.europa.eu:taxud:vies:services:checkVat:types checkVat"`
			Country string   `xml:"countryCode"`
			Vat     string   `xml:"vatNumber"`
		}
	} `xml:"Body"`
}

type VatResponse struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    struct {
		CheckVatResponse struct {
			XMLName     xml.Name `xml:"urn:ec.europa.eu:taxud:vies:services:checkVat:types checkVatResponse"`
			Country     string   `xml:"countryCode"`
			Vat         string   `xml:"vatNumber"`
			RequestDate string   `xml:"requestDate"`
			Valid       string   `xml:"valid"`
			Name        string   `xml:"name"`
			Address     string   `xml:"address"`
		}
	} `xml:"Body"`
}

/*
<soap:Envelope xmlns:soap=" http://schemas.xmlsoap.org/soap/envelope/">
      <soap:Body>
      <checkVatResponse xmlns="urn:ec.europa.eu:taxud:vies:services:chackVat:types">
        <countryCode>MS</countryCode>
        <vatNumber>TESTVATNUMBER</vatNumber>
        <requestDate>"YYYY-MM-DD+HH:MM"</requestDate>
        <valid>false</valid>
        <name>---</name>
        <address>---</address>
      </checkVatResponse>
   </soap:Body>
</soap:Envelope>
*/

/*
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	<soap:Body>
		<checkVatResponse xmlns="urn:ec.europa.eu:taxud:vies:services:checkVat:types">
			<countryCode>SE</countryCode>
			<vatNumber>516404316701</vatNumber>
			<requestDate>2016-02-18+01:00</requestDate>
			<valid>true</valid>
			<name>Skanska SK a.s. Filial Slovakia</name>
			<address>112 74 STOCKHOLM</address>
		</checkVatResponse>
	</soap:Body>
</soap:Envelope>
*/

func main() {
	argv := os.Args[1:]
	if len(argv) != 1 {
		fmt.Println("Need to give VAT number as an argument")
		fmt.Println("(2016 by J Ramb)")
		os.Exit(1)
		return
	}
	vatComb := argv[0]
	country, vat := vatComb[0:2], vatComb[2:]
	//fmt.Println("Checking ", country, vat)
	v := &VatQuery{}
	v.Body.CheckVat.Country = country //"SE"
	v.Body.CheckVat.Vat = vat         //"516404316701"

	buf, err := xml.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	body := bytes.NewBuffer(buf)
	_ = body
	//fmt.Println(body)
	r, err := http.Post("http://ec.europa.eu/taxation_customs/vies/services/checkVatService", "text/xml", body)
	response, _ := ioutil.ReadAll(r.Body)
	//fmt.Println(string(response))
	var result VatResponse
	xml.Unmarshal(response, &result)
	if result.Body.CheckVatResponse.Valid == "false" {
		fmt.Printf("Result: %s%s = %s\n",
			result.Body.CheckVatResponse.Country,
			result.Body.CheckVatResponse.Vat,
			result.Body.CheckVatResponse.Valid,
		)
	} else {
		fmt.Printf(`Result: %s%s = %s
Name:  %s
       %s
Date:  %s
`,
			result.Body.CheckVatResponse.Country,
			result.Body.CheckVatResponse.Vat,
			result.Body.CheckVatResponse.Valid,
			result.Body.CheckVatResponse.Name,
			result.Body.CheckVatResponse.Address,
			result.Body.CheckVatResponse.RequestDate,
		)
	}
	/*
		enc := xml.NewEncoder(os.Stdout)
		enc.Indent("", "  ")
		if err := enc.Encode(v); err != nil {
			fmt.Printf("error: %v\n", err)
		}
	*/
}
