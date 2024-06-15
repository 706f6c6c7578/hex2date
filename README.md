# hex2date
Convert hex, or decimal values, to UTC time stamps.

$ echo 67676767 | hex2date

Sunday, 22. December 2024 01:12:07 +0000 (UTC)

$ echo -n '68686868' | hex2date -tz all

$ date -d 'tomorrow 15:00' +%s | hex2date -d -tz Africa
