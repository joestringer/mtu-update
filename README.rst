Usage
-----

```
$ ./mtu-update -h
Update the MTU inside network namespaces.

Usage:
  mtu-update [flags]

Flags:
  -h, --help                  help for mtu-update
  -m, --mtu int               Base MTU to configure on links (0 for autodetect) (default 1500)
  -p, --period int            Run once every x seconds (default 300)
  -t, --tunnel-overhead int   Expected tunnel overhead for overlay traffic (default 50)
  -v, --verbose               Print verbose debug log messages
```

Update the MTU across a k8s cluster:

```
$ kubectl create -f https://raw.githubusercontent.com/cilium/mtu-update/master/mtu-update.yaml
(Check the logs for the newly created pod to validate that the MTU was updated)
$ kubectl delete -f https://raw.githubusercontent.com/cilium/mtu-update/master/mtu-update.yaml
```

Contact
-------

If you have any questions feel free to contact us on `Slack <https://cilium.herokuapp.com/>`_.


License
-------

This project is licensed under the `Apache License, Version 2.0 <LICENSE>`_.
