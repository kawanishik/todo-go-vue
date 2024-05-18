#!/bin/bash

DIR="$( cd "$( dirname "$( dirname "${BASH_SOURCE[0]}" )" )" &> /dev/null && pwd )"
mysql -h 127.0.0.1 -P 3306 -u root -p my_password -p < "$DIR/lib/database/init.sql"
