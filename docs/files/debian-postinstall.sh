#!/bin/sh

# SPDX-FileCopyrightText: 2021 The NGI Pointer Secure-Scuttlebutt Team of 2020/2021
#
# SPDX-License-Identifier: CC0-1.0

set -e

# create a user to run the server as
adduser --system --home /var/lib/go-ssb-room go-ssb-room
chown go-ssb-room /var/lib/go-ssb-room

if [ "$1" = "configure" ] || [ "$1" = "abort-upgrade" ] || [ "$1" = "abort-deconfigure" ] || [ "$1" = "abort-remove" ] ; then
	# This will only remove masks created by d-s-h on package removal.
	deb-systemd-helper unmask go-ssb-room.service >/dev/null || true

	# was-enabled defaults to true, so new installations run enable.
	if deb-systemd-helper --quiet was-enabled go-ssb-room.service; then
		# Enables the unit on first installation, creates new
		# symlinks on upgrades if the unit file has changed.
		deb-systemd-helper enable go-ssb-room.service >/dev/null || true
	else
		# Update the statefile to add new symlinks (if any), which need to be
		# cleaned up on purge. Also remove old symlinks.
		deb-systemd-helper update-state go-ssb-room.service >/dev/null || true
	fi
fi

if [ "$1" = "configure" ] || [ "$1" = "abort-upgrade" ] || [ "$1" = "abort-deconfigure" ] || [ "$1" = "abort-remove" ] ; then
	if [ -d /run/systemd/system ]; then
		systemctl --system daemon-reload >/dev/null || true
		if [ -n "$2" ]; then
			_dh_action=restart
		else
			_dh_action=start
		fi
		deb-systemd-invoke $_dh_action go-ssb-room.service >/dev/null || true
	fi
fi

# welcome message
cat <<EOF
> Welcome !

go-ssb-room has been installed as a systemd service.

It will store it's files (roomdb and cookie secrets) under /var/lib/go-ssb-room.
This is also where you would put custom translations.

For more configuration background see /usr/share/go-ssb-room/README.md
or visit the code repo at https://github.com/ssb-ngi-pointer/go-ssb-room/tree/master/docs

Like outlined in that document, we highly encourage using nginx with certbot for TLS termination.
We also supply an example config for this. You can find it under /usr/share/go-ssb-room/nginx-example.conf

> Important 

Before you start using room server via the systemd service, you need to at least change the https domain in the systemd service.

Edit /etc/systemd/system/go-ssb-room.service and then run this command to reflect the changes:

sudo systemctl daemon-reload

> Running the room server:

To start/stop go-ssb-room:

sudo systemctl start go-ssb-room
sudo systemctl stop go-ssb-room

To enable/disable go-ssb-room starting automatically on boot:

sudo systemctl enable go-ssb-room
sudo systemctl disable go-ssb-room

To reload go-ssb-room:

sudo systemctl restart go-ssb-room

To view go-ssb-room logs:

sudo journalctl -f -u go-ssb-room

EOF
