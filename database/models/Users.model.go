package models

type Users struct {
	Generic
	FailedSignInAttempts int
	Login                string
	RecoveryAnswerHash   string
	RecoveryQuestion     string
	Passwords            Passwords    `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Secrets              Secrets      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TokenSecrets         TokenSecrets `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
