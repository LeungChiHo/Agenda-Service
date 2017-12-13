package cmd

import (
	"fmt"
	"testing"
)

func TestRegister(t *testing.T) {
	fmt.Println("=====> In TEST of UserRegister")
	UserRegisterCmd.Flags().Set("username", "Alice")
	UserRegisterCmd.Flags().Set("password", "123")
	UserRegisterCmd.Flags().Set("email", "Alice@163.com")
	UserRegisterCmd.Flags().Set("phone", "123")
	UserRegisterCmd.Run(UserRegisterCmd, nil)
}

func TestLogin(t *testing.T) {
	fmt.Println("=====> In TEST of UserLogin")
	UserLoginCmd.Run(UserLoginCmd, nil)
	UserLoginCmd.Flags().Set("username", "Alice")
	UserLoginCmd.Flags().Set("password", "123")
	UserLoginCmd.Run(UserLoginCmd, nil)
}

func TestShowAllUsers(t *testing.T) {
	fmt.Println("=====> In TEST of ListAllUser")
	usersCmd.Run(usersCmd, nil)
}

func TestCreateNewMeeting(t *testing.T) {
	fmt.Println("=====> In TEST of MeetingCreate")
	MeetingCreateCmd.Flags().Set("title", "testMeeting")
	MeetingCreateCmd.Flags().Set("members", "testUser0,testUser1")
	MeetingCreateCmd.Flags().Set("starttime", "2000/01/01/00:00")
	MeetingCreateCmd.Flags().Set("endtime", "2001/01/01/00:00")
	MeetingCreateCmd.Run(MeetingCreateCmd, nil)
}

func TestShowAllMeetings(t *testing.T) {
	fmt.Println("=====> In TEST of ListAllMeeting")
	meetingsCmd.Run(meetingsCmd, nil)
}

func TestUserDelete(t *testing.T) {
	fmt.Println("=====> In TEST of UserDelete")
	UserDeleteCmd.Run(UserDeleteCmd, nil)
	//usersCmd.Run(usersCmd, nil)
}
