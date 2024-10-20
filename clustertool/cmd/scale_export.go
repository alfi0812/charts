package cmd

import (
    "github.com/spf13/cobra"
    "github.com/truecharts/public/clustertool/pkg/scale"
)

var scaleexport = &cobra.Command{
    Use:   "export",
    Short: "Export SCALE Apps to file",
    Run: func(cmd *cobra.Command, args []string) {
        scale.ExportApps()
    },
}

func init() {
    scaleCmd.AddCommand(scaleexport)
}
