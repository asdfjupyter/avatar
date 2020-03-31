package AvatarEmail

import (
	"../AvatarConfig"
	"encoding/csv"
	"io"
	"log"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"time"
)

type Email struct {
	Receiver string
	Subject  string
	Body     string
}

var emailSchedule map[int]Email

func AvatarEmail(cfg *AvatarConfig.Email_cfg) {

	if !cfg.Enable {
		//do nothing yet, as we know Email is our main feature so far.
	}

	//set up a map to store schedule information
	emailSchedule = make(map[int]Email)

	emailcsv, err := os.Open("emails.csv")

	if err != nil {
		log.Fatal("No CSV File Found for Emails")
	}

	emails := csv.NewReader(emailcsv)

	//ignore header
	singleemail, err := emails.Read()

	for {
		//read all lines in and populate the email schedule.
		singleemail, err = emails.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		val_hr, _ := strconv.Atoi(singleemail[0])
		val_min, _ := strconv.Atoi(singleemail[1])

		emailIndex := val_hr*60 + val_min

		emailSchedule[emailIndex] = Email{singleemail[2], strings.Trim(singleemail[3], "'"), strings.Trim(singleemail[4], "'")}
	}

	auth := smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.SMTP)

	for {

		//Sending emails directly through SMTP is not recommended.
		//use Gmail as an example as I tested, you need to open up App Password first.
		//Your Gmail password won't work and I do not recommend you to allow insecure Apps.
		//So, please go to https://myaccount.google.com/
		//Click Security
		//Click Enable 2-Step Verification
		//Add App Passwords
		//Select "Mail", then copy and paste the password generated to the avatar.ini file.

		currentIndex := time.Now().Hour()*60 + time.Now().Minute()

		if emailSchedule[currentIndex].Receiver == "" {
			continue
		}

		to := []string{emailSchedule[currentIndex].Receiver}
		msg := []byte("To:" + emailSchedule[currentIndex].Receiver + "\r\n" +
			"Subject:" + emailSchedule[currentIndex].Subject + "\r\n" +
			"\r\n" +
			emailSchedule[currentIndex].Body + "\r\n")

		delete(emailSchedule, currentIndex)

		err := smtp.SendMail(cfg.SMTP+":"+cfg.PORT, auth, cfg.MasterEmail, to, msg)
		if err != nil {
			log.Fatal(err)
		}

	}

}
