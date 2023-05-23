package main
//import flag "github.com/spf13/pflag"
//これはいらない
type options struct {
	help bool
	version bool
}

func buildOptions(args []string) (*options, *flag.FlagSet) {
	opts := &options{}
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args[0])) }
	flags.BoolVarP(&opts.help, "help", "h", false, "へルプメッセージを表示する")
	flags.BoolVarP(&opts.version, "version", "v", false, "バージョンを表示する")
	return opts, flags
}

func perform(opts *options, args []string) *yubsError {
	fmt.Println("Hello World")
	return nil
}

func parseOptions(args []string) (*options, []string, *yubsError) {
	opts, flags := buildOptions(args)
	flags.Parse(args[1:])
	if opts.help {
		fmt.Println(helpMessage(args[0]))
		return nil, nil, &yubsError{statusCode: 0, message: ""}
	}
	if opts.token == "" {
		return nil, nil, &yubsError{statusCode: 3, message: "no token was given"}
	}
	return opts, flags.Args(), nil
}

func goMain(args []string) int {
	opts, args, err := parseOptions(args)
	if err != nil {
		if err.statusCode != 0 {
			fmt.Println(err.Error())
		}
		return err.statusCode
	}
	if err := perform(opts, args); err != nil {
		fmt.Println(err.Error())
		return err.statusCode
	}
	return 0
}