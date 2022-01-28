# init new package
# PKG=github.com/andrdru/gosvc/test make new
.PHONY: new
new:
	go run template/init.go -pkg ${PKG}
