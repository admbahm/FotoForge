// Package cli defines the FotoForge command-line interface.
package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"path/filepath"

	"github.com/fotoforge/fotoforge/internal/config"
	"github.com/fotoforge/fotoforge/internal/db"
	"github.com/fotoforge/fotoforge/internal/logger"
	"github.com/fotoforge/fotoforge/internal/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// App owns process-scoped CLI dependencies.
type App struct {
	info        version.Info
	out, errOut io.Writer
	viper       *viper.Viper
	cfg         config.Config
	log         *slog.Logger
}

// New constructs the root command without touching the filesystem.
func New(info version.Info, out, errOut io.Writer) *cobra.Command {
	app := &App{info: info, out: out, errOut: errOut, viper: viper.New()}
	cmd := &cobra.Command{
		Use:          "fotoforge",
		Short:        "Safely audit and organize photo and video collections",
		Long:         "FotoForge audits, reconciles, organizes, and preserves media collections using deterministic, explainable, and reversible operations.",
		SilenceUsage: true, SilenceErrors: true, Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error { return cmd.Help() },
	}
	cmd.SetOut(out)
	cmd.SetErr(errOut)
	cmd.PersistentFlags().Bool("verbose", false, "enable verbose structured logging")
	cmd.PersistentFlags().String("config", "", "configuration file (default: .fotoforge.yaml or user config directory)")
	if err := app.viper.BindPFlag("verbose", cmd.PersistentFlags().Lookup("verbose")); err != nil {
		panic(err)
	}
	cmd.PersistentPreRunE = func(cmd *cobra.Command, _ []string) error { return app.initialize(cmd) }
	cmd.AddCommand(app.versionCommand())
	for _, spec := range commandSpecs {
		cmd.AddCommand(app.placeholderCommand(spec))
	}
	return cmd
}

func (a *App) initialize(cmd *cobra.Command) error {
	configFile, err := cmd.Flags().GetString("config")
	if err != nil {
		return fmt.Errorf("read config flag: %w", err)
	}
	if configFile != "" {
		a.viper.SetConfigFile(configFile)
	} else {
		a.viper.SetConfigName(".fotoforge")
		a.viper.SetConfigType("yaml")
		a.viper.AddConfigPath(".")
		if dir, err := config.UserConfigDir(); err == nil {
			a.viper.AddConfigPath(filepath.Join(dir, "fotoforge"))
		}
	}
	a.cfg, err = config.Load(a.viper)
	if err != nil {
		return err
	}
	a.log = logger.New(a.errOut, a.cfg.Verbose)
	a.log.DebugContext(cmd.Context(), "configuration loaded", "database", a.cfg.DatabasePath)
	return nil
}

func (a *App) versionCommand() *cobra.Command {
	var asJSON bool
	cmd := &cobra.Command{Use: "version", Short: "Print build and version information", Args: cobra.NoArgs}
	cmd.RunE = func(_ *cobra.Command, _ []string) error {
		if asJSON {
			encoder := json.NewEncoder(a.out)
			encoder.SetIndent("", "  ")
			return encoder.Encode(a.info)
		}
		_, err := fmt.Fprintln(a.out, a.info.String())
		return err
	}
	cmd.Flags().BoolVar(&asJSON, "json", false, "output build information as JSON")
	return cmd
}

type commandSpec struct{ name, short, detail string }

var commandSpecs = []commandSpec{
	{"audit", "Inventory a media collection without changing it", "Audit will safely inventory files and record observations in the catalog."},
	{"analyze", "Analyze cataloged media", "Analyze will derive deterministic findings from a completed audit."},
	{"report", "Generate an explainable collection report", "Report will present findings and provenance without modifying source media."},
	{"organize", "Plan reversible media organization", "Organize will create and apply explicit, reviewable organization plans."},
	{"quarantine", "Move selected media into reversible quarantine", "Quarantine will isolate explicitly selected files while retaining restoration metadata."},
	{"restore", "Restore media from quarantine", "Restore will return quarantined files to their recorded original locations."},
	{"purge", "Permanently remove explicitly approved quarantined data", "Purge will require an explicit reviewed plan and will never run automatically."},
	{"verify", "Verify collection and catalog integrity", "Verify will check content hashes, catalog records, and operation provenance."},
}

func (a *App) placeholderCommand(spec commandSpec) *cobra.Command {
	return &cobra.Command{Use: spec.name, Short: spec.short, Long: spec.detail + " This command is reserved for a future release.", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error { return a.runPlaceholder(cmd.Context(), spec) }}
}

func (a *App) runPlaceholder(ctx context.Context, spec commandSpec) error {
	catalog, err := db.Open(ctx, a.cfg.DatabasePath)
	if err != nil {
		return err
	}
	defer catalog.Close()
	a.log.InfoContext(ctx, "command is not implemented", "command", spec.name)
	_, err = fmt.Fprintf(a.out, "%s is not implemented yet; no media files were changed.\n", spec.name)
	return err
}
