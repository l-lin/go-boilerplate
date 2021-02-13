package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

func NewCompletionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "completion [bash|zsh|fish|powershell]",
		Short: "Generate completion script",
		Long: `To load completions:

Bash:

$ source <(go-boilerplate completion bash)

# To load completions for each session, execute once:
Linux:
  $ go-boilerplate completion bash > /etc/bash_completion.d/go-boilerplate
MacOS:
  $ go-boilerplate completion bash > /usr/local/etc/bash_completion.d/go-boilerplate

Zsh:

# If shell completion is not already enabled in your environment you will need
# to enable it.  You can execute the following once:

$ echo "autoload -U compinit; compinit" >> ~/.zshrc

# To load completions for each session, execute once:
$ go-boilerplate completion zsh > "${fpath[1]}/_go-boilerplate"

# You will need to start a new shell for this setup to take effect.

Fish:

$ go-boilerplate completion fish | source

# To load completions for each session, execute once:
$ go-boilerplate completion fish > ~/.config/fish/completions/go-boilerplate.fish
`,
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		Args:                  cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "bash":
				cmd.Root().GenBashCompletion(os.Stdout)
			case "zsh":
				cmd.Root().GenZshCompletion(os.Stdout)
			case "fish":
				cmd.Root().GenFishCompletion(os.Stdout, true)
			case "powershell":
				cmd.Root().GenPowerShellCompletion(os.Stdout)
			}
		},
	}
}
