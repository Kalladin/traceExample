Elastic Search

# Raise ram used by jvm
vim ${ES_HOME}/config/jvm.options
# Change settings to
-Xms4g
-Xmx4g

# Raise max file descriptor
`sudo vim /etc/security/limits.conf`

# Add limits

```
*  -  nofile  100000
* soft memlock unlimited
* hard memlock unlimited
* -    nproc  10000
```

# Raise virtual memory
sudo vim /etc/sysctl.conf
# Add setting
`vm.max_map_count=262144`

# Disable Swap
sudo swapoff -a

#Comment out swap partition
/etc/fstab
