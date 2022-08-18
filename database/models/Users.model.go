package models

type Users struct {
	Generic
	FailedSignInAttempts int
	Login                string           `json:"login"`
	RecoveryAnswerHash   string           `json:"recoveryAnswerHash"`
	RecoveryQuestion     string           `json:"recoveryQuestion"`
	AuthCodes            AuthCodes        `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DeletedSecretIDs     DeletedSecretIDs `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Passwords            Passwords        `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Secrets              Secrets          `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TokenSecrets         TokenSecrets     `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
