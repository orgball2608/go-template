package config

var (
	StatusActive   = "on"
	StatusInactive = "off"
)

const (
	USER    = "USER"
	TEACHER = "TEACHER"
)

const (
	TypeSendEmail = "sendEmail"
)

type Constants struct {
	LoginFailLimit                            int `env:"LOGIN_FAIL_LIMIT,default=5"`
	LoginFailDurationMinutes                  int `env:"LOGIN_FAIL_DURATION_MINUTES,default=20"`
	ResetPassWordDurationHours                int `env:"RESET_PASSWORD_DURATION_HOURS,default=24"`
	ResendRequestResetPasswordDurationSeconds int `env:"RESEND_RESET_PASSWORD_DURATION_SECONDS,default=60"`
	SystemTimeOutSecond                       int `env:"SYSTEM_TIMEOUT_SECOND,default=30"`
}
