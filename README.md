sorting hat
===========

`sortinghat` is a service for monitoring users coming into the system and adding them to
(local) unix group(s).

The particular motivating use case is creating and managing a local `rstudio-connect`
group to allow users on Posit Connect to RunAsCurrentUser without needing to
also set up a separate unix group for them, which can be a painful approval activity
in enterprise environments.
