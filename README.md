# netlinkd
 Linux daemon interacts with netlink and provides information to a process via a linux socket
 
## installation
 
netlinkd can be installed using yum or apt package managers

```
sudo apt install https://github.com/fsc-demo-wim/netlinkd/releases/download/v0.1.0/netlinkd_0.1.0_linux_amd64.deb
```

netlinkd package can be installed using the installation script which detects the operating system type and installs the relevant package:

```
sudo curl -sL https://raw.githubusercontent.com/fsc-demo-wim/netlinkd/master/get.sh | sudo bash

sudo curl -sL https://raw.githubusercontent.com/fsc-demo-wim/netlinkd/master/get.sh | sudo bash -s -- -v 0.1.0
``` 

