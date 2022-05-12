GO = go
CC = x86_64-w64-mingw32-gcc
CXX = x86_64-w64-mingw32-g++
HOST = x86_64-w64-mingw32
LD_FLAGS = -s -w -H=windowsgui -extldflags=-static
GOOS = windows
GOARCH = amd64

BASENAME = TwitchToGC
SYSO = $(BASENAME).syso
OUTFILE = $(BASENAME).exe

all: $(OUTFILE)

$(OUTFILE): $(SYSO)
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=1 CC=$(CC) CXX=$(CXX) HOST=$(HOST) go build -ldflags "$(LD_FLAGS)" -v -o $(OUTFILE)

$(SYSO):
	x86_64-w64-mingw32-windres $(BASENAME).rc -O coff -o $(SYSO)
