> Avatar - Your Personal Working from Home Cheater, oh no, Assitant

# Avatar

COVID-19 is here, so is the working-from-home arrangement. It is hard, I know. 

Steam is on sale again with your favourite game; Netflix is advertising the season 2 of your favourite show; Amazon is pushing you new deals just at your fingertips. 

Don't go crazy! Here comes Avatar!

A tool that automates your daily working-from-home routines as much as possible. 

Join us!

It will eventually create a virtual "you" that exposes your working patterns every day.

A digital twin of you. 

Steam, Netflix and Amazon, no matter what they are, you won't miss a thing. 

I promise.

Avatar - Think Lasy Differently.

## Contents
- [Language](#language)
- [Requirements](#requirements)
- [Usage](#usage)
- [Future Plans](#plans)
- [License](#license)

## Language
Go, as I am learning Go recently. 

## Requirements
It uses some third-party libraries. Mainly the following:
```
gopkg.in/ini.v1: to read ini config file
github.com/go-vgo/robotgo: to automate GUI operations
```
So if you work on Windows, then the robotgo may give you some workload. It requires 64bit GCC where you can download from http://mingw-w64.org/doku.php/download

## Usage
Currently, Avatar automates two things: 

1. Email distribution to receivers;
1. Message distribution to a team on Microsoft Teams

Please go to avatar.ini first to change the settings for your SMTP;

Then, open up Microsoft Teams.

Go to emails.csv and teams.csv to manipulate a schedule when you want to send those emails/messages as well as how texts you want to send. 

You can invoke the program by

```
go run avatar.go
```

or 

```
go build avatar.go
```
and invoke the program such as avatar.exe (use powershell if you can, tested)

You need to follow the prompts to calibrate Avatar first so that it knows where to click on Teams to bring up the input filed for the team message. This can be automated in the future. 

After the calibration, you are good to go. 

Avatar will send out the emails as well as sending out the team messages according to the schedule you specified in the CSV files. 

## Plans

Well, this is not enough. 

I will tentatively incorporate other features in longer-term:

1. Get rid of calibration: automatically identify where to type messages and send;
1. Behavioural analysis: perform network activity analysis to identify users' behavioural patterns every day;
1. Behavioural imitation: with the pattern revealed, let Avatar manipulate network packages or manipulate relevant programs to send out network packages that build up similar behavioural patterns.

Yes, we are targeting network auditing! 

It wants to create a digital you!

Avatar is more than an automation tool. Well, hopefully.

## License
MIT