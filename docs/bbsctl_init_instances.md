## bbsctl init instances

Initialize bigblueswarm instances file

### Synopsis

Create instances list file if it does not exists

```
bbsctl init instances [flags]
```

### Examples

```

bbsctl init instances --dest /path/to/file

bbsctl init instances
# generate the following file
#
# kind: InstanceList
# instances: {}

```

### Options

```
  -d, --dest string   File folder destination (default "$HOME/.bigblueswarm")
  -h, --help          help for instances
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.bigblueswarm/.bbsctl.yml) (default "$HOME/.bigblueswarm/.bbsctl.yml")
```

### SEE ALSO

* [bbsctl init](bbsctl_init.md)	 - Initialize a resource

