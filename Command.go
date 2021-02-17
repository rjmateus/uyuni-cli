package main

type Command struct{
	id string
	local bool
	localProcess func()
	remoteCommand string
}