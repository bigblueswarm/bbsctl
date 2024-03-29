## bbsctl init config

Initialize bbsctl configuration

### Synopsis

Create bbsctl if not exists and initialize a basic configuration

```
bbsctl init config [flags]
```

### Examples

```

bbsctl init config --dest /path/to/files

bbsctl init config
# generates the following file
#
# bbs: ""
# apiKey: ""

bbsctl init config --bbs http://bbs.example.com
# generates the following file
# 
# bbs: http://bbs.example.com
# apiKey: ""

bbsctl init config --bbs http://bbs.example.com --key api_key
# generates the following file
#
# bbs: http://bbs.example.com
# apiKey: api_key

```

### Options

```
  -b, --bbs string    BigBlueSwarm url
  -d, --dest string   Configuration file folder destination (default "$HOME/.bigblueswarm")
  -h, --help          help for config
  -k, --key string    BigBlueSwarm admin api key
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.bigblueswarm/.bbsctl.yml) (default "$HOME/.bigblueswarm/.bbsctl.yml")
```

### SEE ALSO

* [bbsctl init](bbsctl_init.md)	 - Initialize a resource

