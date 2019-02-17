package resource

import "github.com/gobuffalo/packr"

func SQLBox() packr.Box {
	return packr.NewBox("./sql")
}
