package alpha

import (
	"github.com/kyma-project/cli/cmd/kyma/alpha/create"
	"github.com/kyma-project/cli/cmd/kyma/alpha/deploy"
	"github.com/kyma-project/cli/cmd/kyma/alpha/sign"
	"github.com/kyma-project/cli/internal/cli"
	"github.com/spf13/cobra"
)

// NewCmd creates a new Kyma CLI command
func NewCmd(o *cli.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "alpha",
		Short: "Experimental commands",
		Long: `Alpha commands are experimental, unreleased features that should only be used by the Kyma team. Use at your own risk.
`,
	}

	cmd.AddCommand(create.NewCmd(o))
	cmd.AddCommand(deploy.NewCmd(deploy.NewOptions(o)))
	cmd.AddCommand(sign.NewCmd(o))

	return cmd
}
