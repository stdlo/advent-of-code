day=$1
cp -r day00 day$day
fd . -tf day$day -x sd 00 $day
