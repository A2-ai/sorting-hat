#!/bin/bash

# Check if the script is running as root
if [[ $EUID -ne 0 ]]; then
    echo "This script must be run as root."
    exit 1
fi

# Define the starting UID
START_UID=4000

for ((i=15; i<17; i++)); do
    # Calculate the UID for the current user
    USER_UID=$((START_UID + i))

    # Create the user without a home directory and set the login shell to /bin/bash
    useradd -m -d /home/user$USER_UID -u $USER_UID user$USER_UID

    # Display information about the created user
    echo "User user$USER_UID created with UID $USER_UID and login shell /bin/bash"
done

# Display a message when the script is done
echo "User creation completed."