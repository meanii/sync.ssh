# sync.ssh

> A powerful CLI tool built in golang that helps to sync your projects (files & directories) to your personal or private
> repository!

## Installation

Requirement: go 1.19

If you got golang in your machine, no you good to go, so we install it from the source

```shell
git clone https://github.com/meanii/sync.ssh.git && cd sync.ssh && bash installer.sh
```

so this command gonna clone the source and install it in your machine and inject all daemon services and start it!

## How to use it

so, after the installation there is few setups which you to do it!

- Create a private repository ( _you can use public repository but not recommanded_ )
- Get github token with CRUD permission
- Now, login sync.ssh with github by the token which you have created

```shell
sync.ssh auth --token=<TOKEN> --username=<GITHUB_USERNAME>
```

- So, the last step is you need to init the sync.ssh by passsing --repo=<REPO_NAME> flag

```shell
sync.ssh init --repo=<REPO_NAME> 
```

So, its done now!

I would recommanded you to check the sync.ssh's daemon service, if its running or not
if its running, good! you can now sync dirs/files

### How to sync a file and dir

its easier process, to sync any file you just need to run this below command

```shell
sync.ssh sync ~/Desktop/dotfile
```

Here, `dotfile` is a dir - it gonna sync everything inside the dir

In order to sync a file

```shell
sync.ssh sync ~/.gitconfig
```

Here, `.gitconfig` is a file here, it will only sync the file!

---

### Copyright & License

- Copyright (C)  2022 [meanii](https://github.com/meanii )
- Licensed under the terms of
  the [GNU General Public License v3.0](https://github.com/meanii/sync.ssh/blob/main/LICENSE)