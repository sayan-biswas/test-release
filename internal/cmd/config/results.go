package config

import (
	"github.com/sayan-biswas/kubectl-tekton/internal/helper"
	"github.com/sayan-biswas/kubectl-tekton/internal/results/config"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericiooptions"
	"k8s.io/cli-runtime/pkg/printers"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/scheme"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

// ResultsOptions holds the command-line options configure tekton results
type ResultsOptions struct {
	PrintFlags  *genericclioptions.PrintFlags
	PrinterFunc printers.ResourcePrinterFunc
	IOStreams   *genericiooptions.IOStreams
	Config      config.Config

	View   bool
	Reset  bool
	Host   string
	Token  string
	Client string
}

var (
	resultsLong = templates.LongDesc(i18n.T(`
		Configure tekton results client.

		You can use --output jsonpath={...} to extract specific values using a jsonpath expression.
	`))

	resultsExample = templates.Examples(`
		# Configure all parameters interactively
		kubectl tekton config results

		# Configure specific parameters interactively
		kubectl tekton config results host token
	`)
)

func Results(s *genericiooptions.IOStreams) *cobra.Command {
	o := &ResultsOptions{
		PrintFlags: genericclioptions.NewPrintFlags("").WithTypeSetter(scheme.Scheme).WithDefaultOutput("yaml"),
		IOStreams:  s,
	}

	c := &cobra.Command{
		Use:     "results",
		Short:   i18n.T("Configure tekton results"),
		Long:    resultsLong,
		Example: resultsExample,
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(o.Complete(cmd, args))
			cmdutil.CheckErr(o.Validate())
			cmdutil.CheckErr(o.Run(args))
		},
	}

	o.PrintFlags.AddFlags(c)
	c.Flags().BoolVarP(&o.View, "view", "", false, "View tekton results config")
	c.Flags().BoolVarP(&o.Reset, "reset", "", false, "Reset tekton results config")
	c.Flags().StringVarP(&o.Host, "host", "", "", "Configure tekton result host")
	c.Flags().StringVarP(&o.Token, "token", "", "", "Configure tekton result token")
	c.Flags().StringVarP(&o.Client, "client", "", "", "Configure tekton result client type")

	return c
}

// Complete completes the required command-line options
func (o *ResultsOptions) Complete(cmd *cobra.Command, args []string) error {
	//if len(args) != 0 {
	//	return cmdutil.UsageErrorf(cmd, "unexpected arguments: %v", args)
	//}

	printer, err := o.PrintFlags.ToPrinter()
	if err != nil {
		return err
	}
	o.PrinterFunc = printer.PrintObj

	return nil
}

// Validate makes sure that provided values for command-line options are valid
func (o *ResultsOptions) Validate() error {
	return nil
}

// Run performs the execution of 'config View' sub command
func (o *ResultsOptions) Run(args []string) (err error) {
	o.Config, err = config.NewConfig()
	if err != nil {
		return
	}

	if len(args) > 0 {
		return o.Config.Set(helper.ParseArgs(args))
	}

	if o.Reset {
		return o.Config.Reset()
	}

	if o.View {
		err = o.PrinterFunc(o.Config.RawConfig(), o.IOStreams.Out)
		if err != nil {
			return err
		}
	}
	return nil
}
