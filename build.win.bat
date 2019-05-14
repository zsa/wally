windres -i dist/win64/wally.rc -O coff -o wally.syso
go-bindata -prefix "ui/build" -o assets.go ui\build\index.dist.js
go build -tags dist -ldflags "-H windowsgui" -o dist\win64\wally.exe
go build -o dist\win64\wally-cli.exe cli\main.go
del wally.syso
upx dist\win64\wally.exe
upx dist\win64\wally-cli.exe
