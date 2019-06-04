
::set ip=192.168.85.148
set ip=212.64.94.62
::set ip=172.17.0.11

cd ..
start gochart.exe
cd test1

start tcpclient.exe --ip=%ip%
start kcpclient.exe --ip=%ip%
start client.exe %ip%

