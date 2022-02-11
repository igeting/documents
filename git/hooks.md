# hooks

## local hooks
```
cd .git/hooks
vi post-commit

#!/usr/bin/sh
./task (task file at project root)
```

## remote hooks

### git server
```
adduser git
passwd git
su - git
mkdir -p /home/git/.ssh
vi /home/git/.ssh/authorized_keys
chmod 700 /home/git/.ssh
chmod 600 /home/git/.ssh/authorized_keys
```

### make hooks
```
cd /home/git
git init --bare test.git
cd /home/git/test.git/hooks
vi post-receive

#!/usr/bin/sh
./task (task file at project root)
```