package main

import (
	"fmt"
	"os"
	"path/filepath"
	flag "github.com/spf13/pflag"
	"github.com/SOICHIRO-NISHIO-github/yubs"
)

const VERSION = "0.1.16"

func versionString(args []string) string {
	prog := "yubs"
	if len(args) > 0 {
		prog = filepath.Base(args[0])
	}
	return fmt.Sprintf("%s version %s", prog, VERSION)
}

/*
ヘルプメッセージの構築
*/
func helpMessage(args []string) string {
	prog := "yubs"
	if len(args) > 0 {
		prog = filepath.Base(args[0])
	}
	return fmt.Sprintf(`%s [OPTIONS] [URLs...]
OPTIONS
    -t, --token <TOKEN>      アクセストークンを入力してください.
    -d, --delete             指定された短縮URLの削除.
    -h, --help               ヘルプメッセージの表示.
    -v, --version            versionの表示.
ARGUMENT
    URL     コマンドラインで入力したURLを短縮URLにする。`, prog)
}

type yubsError struct {
	statusCode int
	message    string
}

func (e yubsError) Error() string {
	return e.message
}

type flags struct {
	deleteFlag    bool
	listGroupFlag bool
	helpFlag      bool
	versionFlag   bool
}

type runOpts struct {
	token  string
	qrcode string
	config string
	group  string
}

/*
This struct holds the values of the options.
*/
type options struct {
	runOpt  *runOpts
	flagSet *flags
}

func newOptions() *options {
	return &options{runOpt: &runOpts{}, flagSet: &flags{}}
}

func (opts *options) mode(args []string) yubs.Mode {
	switch {
	case opts.flagSet.listGroupFlag:
		return yubs.ListGroup
	case len(args) == 0:
		return yubs.List
	case opts.flagSet.deleteFlag:
		return yubs.Delete
	case opts.runOpt.qrcode != "":
		return yubs.QRCode
	default:
		return yubs.Shorten
	}
}

/*
オプションを定義し、オプションへのポインタとフラグセットへのポインタを返します。
*/
func buildOptions(args []string) (*options, *flag.FlagSet) {
	opts := newOptions()
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args)) }
	flags.StringVarP(&opts.runOpt.token, "token", "t", "", "Bitly-apiのアクセストークンを入力してください.")
	flags.BoolVarP(&opts.flagSet.listGroupFlag, "list-group", "L", false, "list the groups. 隠しコマンド.")
	flags.BoolVarP(&opts.flagSet.helpFlag, "help", "h", false, "ヘルプメッセージの表示.")
	flags.BoolVarP(&opts.flagSet.versionFlag, "version", "v", false, "versionの表示.")
	return opts, flags
}

/*
parseOptions は、指定されたコマンド ライン引数からオプションを解析します。
*/
func parseOptions(args []string) (*options, []string, *yubsError) {
	opts, flags := buildOptions(args)
	flags.Parse(args[1:])
	if opts.flagSet.helpFlag {
		fmt.Println(helpMessage(args))
		return nil, nil, &yubsError{statusCode: 0, message: ""}
	}
	if opts.flagSet.versionFlag {
		fmt.Println(versionString(args))
		return nil, nil, &yubsError{statusCode: 0, message: ""}
	}
	if opts.runOpt.token == "" {
		return nil, nil, &yubsError{statusCode: 3, message: "no token was given"}
	}
	return opts, flags.Args(), nil
}

func shortenEach(bitly *yubs.Bitly, config *yubs.Config, url string) error {
	result, err := bitly.Shorten(config, url)
	fmt.Println("main_1")
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

func deleteEach(bitly *yubs.Bitly, config *yubs.Config, url string) error {
	fmt.Println("main_2")
	return bitly.Delete(config, url)
}

func listUrls(bitly *yubs.Bitly, config *yubs.Config) error {
	fmt.Println("main_3")
	urls, err := bitly.List(config)
	if err != nil {
		return err
	}
	for _, url := range urls {
		fmt.Println(url)
	}
	return nil
}

func listGroups(bitly *yubs.Bitly, config *yubs.Config) error {
	groups, err := bitly.Groups(config)
	fmt.Println("main_4")
	if err != nil {
		return err
	}
	for i, group := range groups {
		fmt.Printf("GUID[%d] %s\n", i, group.Guid)
	}
	return nil
}

func performImpl(args []string, executor func(url string) error) *yubsError {
	for _, url := range args {
		err := executor(url)
		if err != nil {
			return makeError(err, 3)
		}
	}
	return nil
}

func perform(opts *options, args []string) *yubsError {
	bitly := yubs.NewBitly(opts.runOpt.group)
	config := yubs.NewConfig(opts.runOpt.config, opts.mode(args))
	config.Token = opts.runOpt.token
	switch config.RunMode {
	case yubs.List:
		err := listUrls(bitly, config)
		return makeError(err, 1)
	case yubs.ListGroup:
		err := listGroups(bitly, config)
		return makeError(err, 2)
	case yubs.Delete:
		return performImpl(args, func(url string) error {
			return deleteEach(bitly, config, url)
		})
	case yubs.Shorten:
		return performImpl(args, func(url string) error {
			return shortenEach(bitly, config, url)
		})
	}
	return nil
}

func makeError(err error, status int) *yubsError {
	if err == nil {
		return nil
	}
	ue, ok := err.(*yubsError)
	if ok {
		return ue
	}
	return &yubsError{statusCode: status, message: err.Error()}
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

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}