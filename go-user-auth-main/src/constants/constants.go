package constants

const (
	ServiceName        = "ACCOUNTS"
	ExpDevice   string = "app"
)

const (
	UserAgent       string = "User-Agent"
	ClientIP        string = "Client-Ip"
	Platform        string = "Platform"
	AppId           string = "App-Id"
	DeviceType      string = "Device-Type"
	UserAgentId     string = "User-Agent-Id"
	MobileIdentiyId string = "Mobile-Identity-Id"
)

const (
	AuthType = "Bearer"
)

const (
	ExperienceDeviceType = "EXPAPP"
)

const (
	GuestUserType = 1
	GuestRegLevel = "Guest"
)

const (
	AuthTimeStampLayout = "2006-01-02T15:04:05-07:00"
)

const (
	NoUserId = 0
)

const (
	MobileIdCaptureGuestUser    = "GUEST"
	MobileIdCaptureRegUser      = "RECOGNIZED"
	MobileIdCaptureEmailRegUser = "EMAILREG"
	EndpointMobileIDCapture     = "mobileidentity/capture"
)

type LoggerFields string

const (
	FuncName LoggerFields = "function"
)

const (
	LogFileKey   string = "log_file"
	Auth         string = "auth"
	Login        string = "login"
	Compliance   string = "compliance"
	EmailVerify  string = "email_verify"
	Maid         string = "maid"
	Password     string = "password"
	Registration string = "registration"
	Subscription string = "subscription"
	UserUpdate   string = "user_update"
	Health       string = "health"
	Activity     string = "activity"
	Convert      string = "convert"
	Member       string = "member"
	Segment      string = "segment"
)
