#!/bin/bash

# expects /opt/sortinghat/sortinghat to exist
# mkdir -p /opt/sortinghat
# mv sorting-hat /opt/sortinghat/sortinghat
cp sortinghat.service /etc/systemd/system/sortinghat.service
cp sortinghat.timer /etc/systemd/system/sortinghat.timer

systemctl daemon-reload                # Reload systemd to recognize the new timer
systemctl enable sortinghat.timer      # Enable the timer to start on boot
systemctl start sortinghat.timer       # Start the timer immediately
