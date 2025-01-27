package cmd

import (
	"context"
	"fmt"

	"github.com/gogf/gf-cli/v2/internal/consts"
	"github.com/gogf/gf-cli/v2/utility/mlog"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gbuild"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
)

var (
	Version = commandVersion{}
)

type commandVersion struct {
	g.Meta `name:"version" brief:"show version information of current binary"`
}

type commandVersionInput struct {
	g.Meta `name:"version"`
}
type commandVersionOutput struct{}

func (c commandVersion) Index(ctx context.Context, in commandVersionInput) (*commandVersionOutput, error) {
	info := gbuild.Info()
	if info["git"] == "" {
		info["git"] = "none"
	}
	mlog.Printf(`GoFrame CLI Tool %s, https://goframe.org`, consts.Version)
	gfVersion, err := c.getGFVersionOfCurrentProject()
	if err != nil {
		gfVersion = err.Error()
	} else {
		gfVersion = gfVersion + " in current go.mod"
	}
	mlog.Printf(`GoFrame Version: %s`, gfVersion)
	mlog.Printf(`CLI Installed At: %s`, gfile.SelfPath())
	if info["gf"] == "" {
		mlog.Print(`Current is a custom installed version, no installation information.`)
		return nil, nil
	}

	mlog.Print(gstr.Trim(fmt.Sprintf(`
CLI Built Detail:
  Go Version:  %s
  Git Commit:  %s
  Build Time:  %s
`, info["go"], info["git"], info["time"])))
	return nil, nil
}

// getGFVersionOfCurrentProject checks and returns the GoFrame version current project using.
func (c commandVersion) getGFVersionOfCurrentProject() (string, error) {
	goModPath := gfile.Join(gfile.Pwd(), "go.mod")
	if gfile.Exists(goModPath) {
		match, err := gregex.MatchString(`github.com/gogf/gf\s+([\w\d\.]+)`, gfile.GetContents(goModPath))
		if err != nil {
			return "", err
		}
		if len(match) > 1 {
			return match[1], nil
		}
		return "", gerror.New("cannot find goframe requirement in go.mod")
	} else {
		return "", gerror.New("cannot find go.mod")
	}
}
