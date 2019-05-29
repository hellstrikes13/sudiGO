mentor: sathish vj + github
http://golangtraining.in
http://learngobyexample.com

1)to make GO binary
#go build -o <filename.go>
will result executable filename
if -o and filename is omitted it shall take by default the underlying Directory name

2) to cross compile in GO
change the env var GOOS
#go env
export GOOS=linux // i m making binary for linux machine on my Mac
go build -o <filename.go>
#file filename.go // shall result in Linux type binary
Similary way can be done done on windows
#go tool dist list // will display the supported OS platforms by GO..


3)to invoke HELP on GO and open browser hit http://localhost:6000
#godoc -http=localhost:6060

