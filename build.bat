:: remove dir
rd /s/q release
:: make dir
md release
::go build -ldflags "-H windowsgui" -o chat.exe
go build -o chat.exe
MOVE chat.exe release\
COPY favicon.ico release\favicon.ico
XCOPY asset\*.* release\asset\  /s /e
XCOPY view\*.* release\view\  /s /e
@pause