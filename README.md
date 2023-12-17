<img src='https://github.com/A2-ai/sorting-hat/assets/3196313/2122e6dc-32f6-468c-89c4-30d9eea08a85' width=100 height=100> sorting hat
===========

  `sortinghat` is a service for monitoring users coming into the system and adding them to
(local) unix group(s).

The particular motivating use case is creating and managing a local `rstudio-connect`
group to allow users on Posit Connect to RunAsCurrentUser without needing to
also set up a separate unix group for them, which can be a painful approval activity
in enterprise environments.

## Installation and Activation

After installing, update either the systemd service file command to include a `--dir` flag
or update the `/etc/sortinghat/config.yml` file to include a `dir` key,
then startup the service via:

```
systemctl daemon-reload                # Reload systemd to recognize the new timer
systemctl enable sortinghat.timer      # Enable the timer to start on boot
systemctl start sortinghat.timer       # Start the timer immediately
```

## Usage

```shell
sortinghat scan --dir=/home
```

to get some additional information about what users were added

```shell
sortinghat scan --dir=/home --loglevel=debug
```


```bash
root@fb44d2942662:/workspaces/sorting-hat# go run main.go watch --dir=/home --loglevel=debug
time=2023-11-08T18:18:42.977Z level=DEBUG msg="found in rstudio-connect group" user=user25 n_grps=2
time=2023-11-08T18:18:42.977Z level=DEBUG msg="found in rstudio-connect group" user=user4009 n_grps=2
time=2023-11-08T18:18:42.977Z level=DEBUG msg="found in rstudio-connect group" user=user21 n_grps=2
time=2023-11-08T18:18:42.977Z level=DEBUG msg="found in rstudio-connect group" user=user27 n_grps=2
time=2023-11-08T18:18:42.977Z level=DEBUG msg="found in rstudio-connect group" user=user28 n_grps=2
time=2023-11-08T18:18:42.977Z level=DEBUG msg="found in rstudio-connect group" user=user4006 n_grps=2
time=2023-11-08T18:18:42.977Z level=DEBUG msg="found in rstudio-connect group" user=user4007 n_grps=2
time=2023-11-08T18:18:42.977Z level=DEBUG msg="found in rstudio-connect group" user=user23 n_grps=2
time=2023-11-08T18:18:42.977Z level=DEBUG msg="found in rstudio-connect group" user=user20 n_grps=2
time=2023-11-08T18:18:42.977Z level=DEBUG msg="found in rstudio-connect group" user=user29 n_grps=2
time=2023-11-08T18:18:42.977Z level=DEBUG msg="found in rstudio-connect group" user=user4008 n_grps=2
time=2023-11-08T18:18:42.977Z level=DEBUG msg="found in rstudio-connect group" user=user22 n_grps=2
time=2023-11-08T18:18:42.977Z level=DEBUG msg="found in rstudio-connect group" user=user1 n_grps=2
time=2023-11-08T18:18:42.977Z level=DEBUG msg="found in rstudio-connect group" user=user24 n_grps=2
time=2023-11-08T18:18:42.977Z level=DEBUG msg="found in rstudio-connect group" user=user26 n_grps=2
time=2023-11-08T18:18:42.977Z level=DEBUG msg="found in rstudio-connect group" user=user2 n_grps=2
time=2023-11-08T18:18:42.977Z level=DEBUG msg="found in rstudio-connect group" user=user3 n_grps=2
time=2023-11-08T18:18:42.977Z level=DEBUG msg="found in rstudio-connect group" user=user4 n_grps=2
time=2023-11-08T18:18:42.977Z level=DEBUG msg="found in rstudio-connect group" user=user4005 n_grps=2
time=2023-11-08T18:18:42.977Z level=INFO msg="ran user check on rstudio-connect group membership" users_added=0
time=2023-11-08T18:18:42.977Z level=DEBUG msg="users added" names=[]
root@fb44d2942662:/workspaces/sorting-hat# go run main.go watch --dir=/home
time=2023-11-08T18:27:57.872Z level=INFO msg="ran user check on rstudio-connect group membership" users_added=0 users_checked=19
```
