#!/bin/sh
if ! which dlv ; then
	PATH="${GOPATH}/bin:$PATH"
fi
if [ "$DEBUG_AS_ROOT" = "true" ]; then
	exec sudo dlv "$@"
else
	exec dlv "$@"
fi