# sync.ssh

> A powerful CLI tool built in golang that helps to sync your projects (files & directories) to your personal or private
> repository!

## Installation

Requirement: go 1.19

If you have golang installed, you may run the command below to install sync. ssh

```shell
git clone https://github.com/meanii/sync.ssh.git && cd sync.ssh && bash installer.sh
```

Consequently, this command will clone the code, install it on your computer, inject all required daemon services, and
launch the programme!

## How to use it

Therefore, there are a few setups that need to be completed after installation.

- Create a private repository (public repositories are permitted but not advised).
- Obtain a github token with CRUD access.
- Now, log in to github using sync.ssh and the token you created.

```shell
sync.ssh auth --token=<TOKEN> --username=<GITHUB_USERNAME>
```

- The final step is to launch sync.ssh by using the —repo=<REPO_NAME> parameter.

```shell
sync.ssh init --repo=<REPO_NAME> 
```

So, it's done now!

I would recommend you to check the sync.ssh's daemon service, if it's running or not
if its running, good! you can now sync dirs/files

The daemon can be manually started via `sync.ssh deamon` will sync everything, so there!

### How to sync a file and dir

It's a simpler approach; all you have to do is run the command below to sync any file.

```shell
sync.ssh sync ~/Desktop/dotfile
```

Here, `dotfile` is a dir - it going to sync everything inside the dir

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