## bbsctl init tenant

Initialize a new bigblueswarm tenant configuration file

### Synopsis

Initialize a new bigblueswarm tenant configuration file if not exits

```
bbsctl init tenant [flags]
```

### Examples

```

bbsctl init tenant --host bbs.example.com
# generates the following file
#
# kind: Tenant
# spec:
#    host: bbs.example.com
# instances: []

bbsctl init tenant --host bbs.example.com --secret dummy_secret
# generates the following file
#
# kind: Tenant
# spec:
#    host: bbs.example.com
#		 secret: dummy_secret
# instances: []

bbsctl init tenant --host bbs.example.com --dest /path/to/file

bbsctl init tenant --host bbs.example.com --meeting_pool 10
# generates the following file
#
# kind: Tenant
# spec:
#    host: bbs.example.com
#    meeting_pool: 10
# instances: []

bbsctl init tenant --host bbs.example.com --meeting_pool 10 --user_pool 100
# generates the following file
#
# kind: Tenant
# spec:
#    host: bbs.example.com
#    meeting_pool: 10
#    user_pool: 100
# instances: []
	
```

### Options

```
  -d, --dest string        File folder destination (default "$HOME/.bigblueswarm")
  -h, --help               help for tenant
      --host string        Tenant hostname
      --meeting_pool int   Tenant meeting pool. This means the tenant can't create more meetings than the configured meeting pool. -1 is ignored. (default -1)
      --secret string      Tenant secret
      --user_pool int      Tenant user pool. This means the tenant can't have more users than the configured user pool. -1 is ignored. (default -1)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.bigblueswarm/.bbsctl.yml) (default "$HOME/.bigblueswarm/.bbsctl.yml")
```

### SEE ALSO

* [bbsctl init](bbsctl_init.md)	 - Initialize a resource

