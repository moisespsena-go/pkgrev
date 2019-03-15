package pkgrev

import (
	"os"
	"time"

	"github.com/aghape/media/media_library"

	"github.com/hashicorp/go-version"
	"github.com/moisespsena-go/aorm"
)

type Project struct {
	aorm.AuditedSDModel
	Package string `sql:"index"`
	Name    string `sql:"index"`
	Oss     []Os
}

type Os struct {
	aorm.KeyStringSerial
	Name      string
	ProjectID string `sql:"index;not null"`
	Archs     []Arch
}

type Arch struct {
	aorm.KeyStringSerial
	Name     string
	OsID     string `sql:"index;not null"`
	Versions []Version
}

type Version struct {
	aorm.AuditedSDModel
	Version version.Version `sql:"index"`

	ArchID string `sql:"index;not null"`
	Files  []File
}

type File struct {
	// SHA256SUM
	ID   string `gorm:"size:32;primary_key"`
	Path string `sql:"index"`
	Mode os.FileMode
	Url  string
	BuildDate time.Time

	File media_library.File

	VersionID string  `sql:"index;not null"`
	Version   Version `sql:"not null" gorm:"preload:Version"`

	aorm.Audited
	aorm.SoftDeleteAudited
}
