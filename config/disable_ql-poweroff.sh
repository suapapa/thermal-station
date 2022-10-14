# /usr/bin/disable_poweroff.sh

#!/bin/bash
echo -n -e '\x1b\x69\x55\x41\x00\x00' > $1