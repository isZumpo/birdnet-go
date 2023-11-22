package realtime

import (
	"github.com/spf13/cobra"
	"github.com/tphakala/BirdNET-Go/internal/analysis"
	"github.com/tphakala/BirdNET-Go/internal/config"
)

// RealtimeCommand creates a new command for real-time audio analysis.
func Command(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "realtime",
		Short: "Analyze audio in realtime mode",
		Long:  "Start analyzing incoming audio data in real-time looking for bird calls.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return analysis.RealtimeAnalysis(ctx)
		},
	}

	setupFlags(cmd, ctx.Settings)

	return cmd
}

// setupRealtimeFlags configures flags specific to the realtime command.
func setupFlags(cmd *cobra.Command, settings *config.Settings) {
	cmd.Flags().StringVar(&settings.ClipPath, "clippath", "", "Path to save audio clips")
	cmd.Flags().StringVar(&settings.ClipType, "cliptype", "", "Audio clip type: wav, flac, mp3")
	cmd.Flags().StringVar(&settings.LogPath, "logpath", "", "Path to save log files")
	cmd.Flags().StringVar(&settings.LogFile, "logfile", "", "Filename for the log file")
	cmd.Flags().BoolVar(&settings.ProcessingTime, "processingtime", false, "Report processing time for each detection")

	config.BindFlags(cmd, settings)
}
