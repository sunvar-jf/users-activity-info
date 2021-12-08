package fileparser

import (
	"fileparsemod/src/displayparser"
	"fileparsemod/src/helpers"
	"fileparsemod/src/helpers/timerhelp"
	"fileparsemod/src/model"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"time"
)

//Parse the file
func ParseFile(logpath string, excelpath string) {
	validationlog := Validatefilesinfolder(logpath)
	if validationlog == "No Error" {
		accessdata := readfile(logpath + helpers.Final_Log)
		userinfo := make(map[string][]model.UserInfo)
		//initialize the variable to 0 for looping below
		i := 0
		//split the lines
		splitdata := createlinesplit(accessdata)
		for _, v := range splitdata {
			//fmt.Println(v)
			//find indexof search string
			indexstr := strings.Index(v, helpers.Last_Login)
			if indexstr >= 0 {
				adduserinfo(v, userinfo)
			}
			//userinfo
			i++
		}

		yearmap, userlistlastlogin := ParseMonthlyUserInfo(userinfo)
		displayparser.GetExcelBook(excelpath, yearmap, userlistlastlogin)
	}
}

//this is for parsing data for excel
func ParseMonthlyUserInfo(userinfo map[string][]model.UserInfo) (map[int]map[string]int, map[string]string) {
	yearmap := buildyearmap()
	userlastlogin := map[string]string{}
	//fmt.Println(userlastlogin)
	for user, userinfolist := range userinfo {
		fmt.Println(user)
		logindates := []time.Time{}
		for _, userloginfo := range userinfolist {
			getmonthmap(userloginfo.LoginDate, yearmap)
			fmt.Println("time logged in :", userloginfo.LoginDate[0])
			logindates = append(logindates, userloginfo.LoginDate[0])
		}
		userlastlogin[user] = getlastloginmap(logindates).String()
	}
	return yearmap, userlastlogin
}

func getlastloginmap(timearr []time.Time) time.Time {
	timeslic := timesorted(timearr)[0]
	return timeslic
}

func timesorted(timearr []time.Time) timerhelp.TimeSlice {
	var dateSlice timerhelp.TimeSlice = timearr
	sort.Sort(sort.Reverse(dateSlice))
	return dateSlice
}

func getmonthmap(logindates []time.Time, yearmap map[int]map[string]int) map[int]map[string]int {
	for _, logindate := range logindates {
		year := logindate.Year()
		value1, isMapContainsKey := yearmap[year]
		if isMapContainsKey {
			month := logindate.Month().String()
			value := value1[month]
			value++
			//delete the key
			delete(value1, month)
			//set the new value
			value1[month] = value
			yearmap[year] = value1
		}
	}
	return yearmap
}

//build month map. default one
func buildmonthmap() map[string]int {
	// Initializing a map
	month := map[string]int{"January": 0, "February": 0, "March": 0, "April": 0, "May": 0, "June": 0, "July": 0, "August": 0, "September": 0,
		"October": 0, "November": 0, "December": 0}
	return month
}

//build month map. default one
func buildyearmap() map[int]map[string]int {
	var presentyear = time.Now().Year()
	yearmap := map[int]map[string]int{}
	yearmap[presentyear] = buildmonthmap()
	yearmap[presentyear-1] = buildmonthmap()
	return yearmap
}

//add user to the array of users if user doesn't exit. if exists then add the time to the UserInfo array.
func adduserinfo(loginput string, userinfo map[string][]model.UserInfo) map[string][]model.UserInfo {
	username, logintime := getusername(loginput)
	value := userinfo[username] //ok can be used to check true or false.
	//println("bool", ok)
	//initialize
	userlog := &model.UserInfo{}
	userlog.LoginDateTime = append(userlog.LoginDateTime, logintime) //setu logindatetime
	usertime, _ := time.Parse(time.RFC3339, logintime)
	userlog.LoginDate = append(userlog.LoginDate, usertime)
	//println(userlog.LoginDate[userlog.LoginDate.Len()-1].Date())
	delete(userinfo, username)      //delete the key from the object
	value = append(value, *userlog) //append the array with new entry
	userinfo[username] = value      //set the map key to new value
	return userinfo
}

//return //loginuser and logintime
func getusername(loginput string) (string, string) {
	//find indexof search string
	//2021-11-18T14:26:18.371Z|63ac819f617a9272|127.0.0.1|jfrt@01fjtwjafm4rtg0gjmsbse0f5n|PUT|/access/api/v1/users/last_login/testuser2|204|-1|66|1|JFrog Access Java Client/7.27.5 72705900  Artifactory/7.27.6 72706900
	indexstr := strings.Index(loginput, helpers.Last_Login)
	//get username from (index of last_login + length of last_login + 1(/)) till length of whole string
	usernamewithlog := loginput[indexstr+len(helpers.Last_Login)+1:]
	//get username from usernamewithlog
	username := strings.Split(usernamewithlog, "|")[0] //first split from above for | to get first index to get time.
	indextimestr := strings.Split(loginput, "|")       //split for | to get first index to get time.
	logintime := loginput[:len(indextimestr[0])]
	return username, logintime
}

//read the file
func readfile(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
		return err.Error()
	}
	//fmt.Println("Contents of file:", string(data))
	return string(data)
}

func Validatefilesinfolder(path string) string {
	validationlog := "No Error"
	files, err := ioutil.ReadDir(path)
	//check directy
	if err != nil {
		//fmt.Println("Directory reading error", err)
		validationlog = "Error reading directory"
		return validationlog //err.Error()
	}

	if len(files) == 0 {
		validationlog = "No files in directory"
		return validationlog
	}

	final_log_file := path + helpers.Final_Log
	//create final log file
	log, err := os.Create(final_log_file)
	fmt.Println(log.Name())
	if err != nil {
		validationlog = "Error creating file"
		return validationlog
	}

	//append all file content into single file
	for _, f := range files {
		if strings.Contains(f.Name(), ".log") && !strings.Contains(f.Name(), helpers.Final_Log) {
			data, _ := ioutil.ReadFile(path + f.Name())
			if strings.Contains(string(data), helpers.Access_Search) && strings.Contains(string(data), helpers.Last_Login) {
				f, _ := os.OpenFile(final_log_file, os.O_APPEND|os.O_WRONLY, 0600)
				defer f.Close()
				f.WriteString(string(data))
			}
		}
	}
	return validationlog
}

func createlinesplit(filedata string) []string {
	splitstr := "\n"
	tempArr := strings.Split(filedata, splitstr)
	return tempArr
}
