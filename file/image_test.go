package file

import (
	"fmt"
	"testing"

	"github.com/devasiajoseph/webapp/core"
	"github.com/devasiajoseph/webapp/db/postgres"
)

func TestBlankImage(t *testing.T) {
	core.Start()
	postgres.InitDb()
	img := GetBlankImage()
	fmt.Println(img.Src)
}
