package AvatarConfig

type General_cfg struct {
	Name        string
	Description string
	Version     float64
	MyMaster    string
	Modules     []string
}

type Email_cfg struct {
	Enable      bool
	MasterEmail string
	Username    string
	Password    string
	SMTP        string
	PORT        string
	//Messages    string
}

type Screen_cfg struct {
	Enable bool
}
