package AvatarScreenOperation

import (
	"../AvatarConfig"
	"encoding/csv"
	"github.com/go-vgo/robotgo"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var teamsSchedule map[int]string

func AvatarScreenOperation(cfg *AvatarConfig.Screen_cfg) {

	if !cfg.Enable {
		//do nothing yet, as we know Screen Operation is our main feature so far.
	}
	//First, we need to know where to click on.
	//There is a method in Robotgo to invoke a window through PID search, unfortunately it did not work on my development machine, Windows 10
	//So, let us use a manual calibration method to invoke Teams.

	//1st where is Teams on taskbar?
	var Teams_pos [2]int
	//2nd where is Teams Icon on the left?
	var TeamsIcon_pos [2]int
	//3rd where is the team I want to send message to?
	var TeamName_pos [2]int
	//4th where is the input field for messages
	var Input_field [2]int

	pressed := 0

	log.Print("Okay, please hover your mouse over Teams on Taskbar and press s")
	for pressed < 4 {
		key_store := robotgo.AddEvent("s")
		x, y := robotgo.GetMousePos()

		if key_store {
			switch pressed {
			case 0:
				Teams_pos[0] = x
				Teams_pos[1] = y
				pressed += 1 //2
				log.Print(Teams_pos)
				log.Print("Okay, please press Teams on Taskbar, and then put your mouse over Teams Icon on the left and press s")
			case 1:
				TeamsIcon_pos[0] = x
				TeamsIcon_pos[1] = y
				pressed += 1 //3
				log.Print(TeamsIcon_pos)
				log.Print("Okay, please put your mouse over the team you want to talk to and press s")
			case 2:
				TeamName_pos[0] = x
				TeamName_pos[1] = y
				pressed += 1 //3
				log.Print(TeamName_pos)
				log.Print("Okay, please put your mouse over the input field and press s")
			case 3:
				Input_field[0] = x
				Input_field[1] = y
				pressed += 1 //3
				log.Print(Input_field)
				log.Print("Thank you, now your Avatar is now fully calibrated!")
			}
		}
	}

	//set up a map to store schedule information
	teamsSchedule = make(map[int]string)

	msgcsv, err := os.Open("teams.csv")

	if err != nil {
		log.Fatal("No CSV File Found for Teams")
	}

	msgs := csv.NewReader(msgcsv)

	//ignore header
	singlemsg, err := msgs.Read()

	for {
		singlemsg, err = msgs.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		val_hr, _ := strconv.Atoi(singlemsg[0])
		val_min, _ := strconv.Atoi(singlemsg[1])

		msgindex := val_hr*60 + val_min

		teamsSchedule[msgindex] = strings.Trim(singlemsg[2], "'")

		log.Print(msgindex, teamsSchedule[msgindex])
	}

	for {

		currentIndex := time.Now().Hour()*60 + time.Now().Minute()

		if teamsSchedule[currentIndex] == "" {
			continue
		}
		log.Print(teamsSchedule[currentIndex])
		//now, click the teams icon on taskbar
		//Teams_pos
		robotgo.MoveMouseSmooth(Teams_pos[0], Teams_pos[1])
		robotgo.MouseClick("left", true)

		//then, click Teams Icon
		//TeamsIcon_pos
		robotgo.MoveMouseSmooth(TeamsIcon_pos[0], TeamsIcon_pos[1])
		robotgo.MouseClick("left", true)

		//then, click the team
		//TeamName_pos
		robotgo.MoveMouseSmooth(TeamName_pos[0], TeamName_pos[1])
		robotgo.MouseClick("left", true)

		//then, click the input field
		//Input_field
		robotgo.MoveMouseSmooth(Input_field[0], Input_field[1])
		robotgo.MouseClick("left", true)

		//then, type the msg in
		robotgo.TypeStr(teamsSchedule[currentIndex])
		robotgo.KeyTap("enter")

		delete(teamsSchedule, currentIndex)
	}

}
