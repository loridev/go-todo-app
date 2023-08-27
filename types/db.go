package types

import "gorm.io/gorm"

type ConnectionOpener func(dsn string) gorm.Dialector
