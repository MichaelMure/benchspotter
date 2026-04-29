package commands

import (
	"cmp"
	"context"
	"fmt"
	"io"

	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/spf13/cobra"

	"benchspotter/commands/execenv"
	"benchspotter/commands/inputs"
	"benchspotter/engine"
)

type showDiffOptions struct {
	session string
}

func newShowDiffCommand(env *execenv.Env) *cobra.Command {
	options := showDiffOptions{}

	cmd := &cobra.Command{
		Use:     "diff",
		Short:   "Show the (git) source diff when the benchmark was run",
		PreRunE: execenv.LoadRepo(env),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runShowDiffCommand(cmd.Context(), env, options)
		},
	}

	flags := cmd.Flags()

	flags.StringVar(&options.session, "session", "", "The session to show the diff for")

	return cmd
}

func runShowDiffCommand(ctx context.Context, env *execenv.Env, options showDiffOptions) error {
	var err error
	var selection *engine.SessionInfo

	if len(options.session) == 0 {
		const recallKey = "show_diff_session"
		preSelected := env.Repo.GetRecall(recallKey)
		selection, err = inputs.SelectSession(ctx, env, preSelected,
			func(info *engine.SessionInfo) bool {
				return info.HasGitDiff()
			})
		if err != nil {
			return err
		}

		err = env.Repo.SetRecall(recallKey, selection.Id)
		if err != nil {
			return err
		}
	} else {
		sessions, err := engine.LocateSessions(env.Repo.Storage())
		if err != nil {
			return err
		}
		for _, s := range sessions {
			if s.Id == options.session {
				selection = s
				break
			}
		}
		if selection == nil {
			return fmt.Errorf("session %q not found", options.session)
		}
	}
	if !selection.HasGitDiff() {
		return fmt.Errorf("this session doesn't have a git diff")
	}

	f, err := selection.OpenFile(engine.GitDiffFilename)
	if err != nil {
		return err
	}
	defer f.Close()

	// limit to 10MB
	diff, err := io.ReadAll(io.LimitReader(f, 10*1024*1024))
	if err != nil {
		return err
	}

	switch env.Format {
	case execenv.FormatText:
		lexer := cmp.Or(lexers.Get("diff"), lexers.Fallback)
		formatter := formatters.TTY16m
		style := cmp.Or(styles.Get("monokai"), styles.Fallback)
		it, err := lexer.Tokenise(nil, string(diff))
		if err != nil {
			return err
		}
		viewport, runFn := env.Viewport(ctx)
		if line := engine.FormatMachineLine(selection.Machine, selection.GoVersion); line != "" {
			fmt.Fprintln(viewport, env.Style.TonedDown(line))
			fmt.Fprintln(viewport)
		}
		if err := formatter.Format(viewport, style, it); err != nil {
			return err
		}
		return runFn()
	case execenv.FormatRaw:
		_, err = env.Out.Write(diff)
		return err
	case execenv.FormatJSON:
		type showDiffJSON struct {
			SessionID   string `json:"session_id"`
			SessionName string `json:"session_name"`
			Diff        string `json:"diff"`
		}
		return env.Out.PrintJSON(showDiffJSON{
			SessionID:   selection.Id,
			SessionName: selection.HumanName,
			Diff:        string(diff),
		})
	default:
		return fmt.Errorf("unsupported format %v for show diff (text, json, raw)", env.Format)
	}
}
