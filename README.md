mentor: sathish vj + github
http://golangtraining.in
http://learngobyexample.com

### to make GO binary
will result executable filename
#go build -o <filename.go>  
if -o and filename is omitted it shall take by default the underlying Directory name

### to cross compile in GO
change the env var GOOS<br>
#go env<br>
 i m making binary for linux machine on my Mac<br>
#export GOOS=linux <br>
#go build -o <filename.go>
  shall result in Linux type binary<br>
#file filename.go<br>

#go tool dist list // will display the supported OS platforms by GO..<br>


### to invoke HELP on GO and open browser hit http://localhost:6000
#godoc -http=localhost:6060

