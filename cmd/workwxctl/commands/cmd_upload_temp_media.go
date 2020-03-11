package commands

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/xen0n/go-workwx"
)

func cmdUploadTempMedia(c *cli.Context) error {
	cfg := mustGetConfig(c)
	mediaType := c.String(flagMediaType)

	filename := c.Args().Get(0)
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	media, err := workwx.NewMediaFromFile(f)
	if err != nil {
		return err
	}

	app := cfg.MakeWorkwxApp()

	var result *workwx.MediaUploadResult
	switch mediaType {
	case "image":
		result, err = app.UploadTempImageMedia(media)
	case "voice":
		result, err = app.UploadTempVoiceMedia(media)
	case "video":
		result, err = app.UploadTempVideoMedia(media)
	case "file":
		result, err = app.UploadTempFileMedia(media)
	default:
		return fmt.Errorf("unknown media type: %s", mediaType)
	}

	fmt.Printf("upload result = %+v\nerr = %+v\n", result, err)

	return err
}
