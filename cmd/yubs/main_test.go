package main

import "testing"

func Example_No_Argument() {
	goMain([]string{"./yubs"})
	// Output:
	// アクセストークンを入力してください。
}

func Example_Token() {
	goMain([]string{"./yubs", "--token"})
	// Output:
	// アクセストークンを入力してください。
}

func Example_Delete() {
	goMain([]string{"./yubs", "--delete"})
	// Output:
	// アクセストークンを入力してください。
}

func Example_Completion() {
	goMain([]string{"./yubs", "--generate-completions"})
	// Output:
	// GenerateCompletion
	// アクセストークンを入力してください。
}

func Example_Help() {
	goMain([]string{"./yubs", "--help"})
	// Output:
	// yubs [OPTIONS] [URLs...]
    // OPTIONS
	//     -t, --token <TOKEN>      アクセストークンを入力してください.
	//     -h, --help               ヘルプメッセージの表示.
	//     -v, --version            versionの表示.
	// ARGUMENT
	//     URL     コマンドラインで入力したURLを短縮URLにする。
}

func Test_Main(t *testing.T) {
	if status := goMain([]string{"./yubs", "-v"}); status != 0 {
		t.Error("Expected 0, got ", status)
	}
}
