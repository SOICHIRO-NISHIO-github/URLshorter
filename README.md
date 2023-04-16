# YUbS
入力したURLを短縮URLにする。
##
Bit.lyは入力したURLを短縮URLにできるWebapiのサービスであるがサイトにわざわざ移動し、URLを入力してなど、すこし面倒である。そこでCLIでBit.lyを利用することでよりスムーズにURLを短縮できるようにした。
--help　引数の説明
--url urlを入力する引数
##
出力形式　　
JSON
##
surl [GLOBAL_OPTS] <COMMAND>
GLOBAL_OPTS
  -t, --token <TOKEN>   specify the API token.
  -v, --verbose         verbose mode.
  -h, --help            print the help message and exit.
  -V, --version         print the version and exit.
COMMAND
  create   create new shorten url from the given url.
  remove   remove the given shortened urls.
  info     show the information of the given shortened urls.
  list     list the shortened urls and the corresponding urls.
  update   update corresponding url of the given shortened url.
  ##


