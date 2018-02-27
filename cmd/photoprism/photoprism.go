package main

import (
	"github.com/photoprism/photoprism"
	"github.com/urfave/cli"
	"os"
	"fmt"
)

func main() {
	config := photoprism.NewConfig()

	app := cli.NewApp()
	app.Name = "PhotoPrism"
	app.Usage = "Sort, view and archive photos on your local computer"
	app.Version = "0.0.1"
	app.Copyright = "Michael Mayer <michael@liquidbytes.net>"
	app.Flags = globalCliFlags
	app.Commands = []cli.Command{
		{
			Name:  "config",
			Usage: "Displays global configuration values",
			Action: func(c *cli.Context) error {
				config.SetValuesFromFile(photoprism.getExpandedFilename(c.GlobalString("config-file")))

				config.SetValuesFromCliContext(c)

				fmt.Printf("<name>                <value>\n")
				fmt.Printf("config-file           %s\n", config.ConfigFile)
				fmt.Printf("darktable-cli         %s\n", config.DarktableCli)
				fmt.Printf("originals-path        %s\n", config.OriginalsPath)
				fmt.Printf("thumbnails-path       %s\n", config.ThumbnailsPath)
				fmt.Printf("import-path           %s\n", config.ImportPath)
				fmt.Printf("export-path           %s\n", config.ExportPath)

				return nil
			},
		},
		{
			Name:  "import",
			Usage: "Imports photo from a directory",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "import-directory, d",
					Usage: "Import directory",
					Value: "~/Pictures/Import",
				},
			},
			Action: func(c *cli.Context) error {
				config.SetValuesFromFile(photoprism.getExpandedFilename(c.GlobalString("config-file")))

				config.SetValuesFromCliContext(c)

				fmt.Printf("Importing photos from %s\n", config.ImportPath)

				importer := photoprism.NewImporter(config.OriginalsPath)

				importer.ImportPhotosFromDirectory(config.ImportPath)

				fmt.Println("Done.")

				return nil
			},
		},
	}

	app.Run(os.Args)
}

var globalCliFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "config-file, c",
		Usage: "Config filename",
		Value: "~/.photoprism",
	},
	cli.StringFlag{
		Name:  "darktable-cli",
		Usage: "Darktable CLI app",
		Value: "/Applications/darktable.app/Contents/MacOS/darktable-cli",
	},
	cli.StringFlag{
		Name:  "storage-path",
		Usage: "Storage path",
		Value: "~/Photos",
	},
}