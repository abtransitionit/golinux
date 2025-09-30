package release

import (
	"context"

	"github.com/abtransitionit/gocore/logx"
)

func listValue(ctx context.Context, logger logx.Logger) (string, error) {
	logger.Info("To implement: Display user defined values about a relase")
	return "", nil
}

// var releaseValueShortDesc = "Display user defined values about a relase"

// // Parent command
// var valueCmd = &cobra.Command{
// 	Use:   "value ChartName Namespace",
// 	Short: releaseValueShortDesc,
// 	Example: `
// 	xxx kbe-cilicium  kube-system
// 	`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("\n🟦",releaseValueShortDesc)
// 		// error : exit : the number of args must be exactly 2
// 		if len(args) != 2 {
// 			fmt.Fprintln(os.Stderr, "❌ Error: you must specify a release name follow by the k8s namespace.")
// 			return
// 		}
// 		cli := fmt.Sprintf(`helm get values %s -n %s`, args[0], args[1])
// 		output,cerr,err := config.PlayQueryHelm(cli)
// 		if err != nil { fmt.Fprintln(os.Stderr, cerr)}
// 		fmt.Println(output)
// },
// }
