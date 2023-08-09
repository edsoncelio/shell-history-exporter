# go-history-exporter
Small tool to export the zsh_history to a private Gist.


## Requirements
* A valid Github personal token with gist permissions exported with the name GH_TOKEN

## How to use
```
# build 
go build . 

# execute
export GH_TOKEN=xxxxxxxxxx
./go-history-exporter
``````

Expected output:
```
➜ ./go-history-exporter
zsh_history backup tool!
Reading zsh_history file...
Creating Gist..
Gist created successfully!!
```
