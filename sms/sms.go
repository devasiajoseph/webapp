package sms

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/devasiajoseph/webapp/core"
)

var otpURL = "https://api.msg91.com/api/v5/otp?template_id=%s&mobile=%s&authkey=%s&otp=%s"

func SendOTP(otp string, mobile string, templateID string) error {
	url := fmt.Sprintf(otpURL, templateID, mobile, core.Config.SmsAuthKey, otp)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request")
		return err
	}

	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Error making request")
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error reading response")
		fmt.Println(err)
		fmt.Println(string(body))
		return err
	}
	return err
}

/*func ActivateOTP(otp string, mobile string) error {
	url := fmt.Sprintf(otpURL, otpTemplateID, mobile, authKey, otp)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request")
		return err
	}

	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Error making request")
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error reading response")
		return err
	}

	fmt.Println(res)
	fmt.Println(string(body))

	return err
}*/
