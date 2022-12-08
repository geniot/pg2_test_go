all: clean build

clean:
	rm bin/pg2_test -f
	rm bin/pg2_test.gcw -f

build:
	go build -o bin/pg2_test geniot.com/geniot/pg2_test_go/cmd/pg2_test

mips:
	CC='/opt/gcw0-toolchain/usr/bin/mipsel-gcw0-linux-uclibc-gcc' \
	CGO_CFLAGS='-I/opt/gcw0-toolchain/usr/mipsel-gcw0-linux-uclibc/sysroot/usr/include -I/opt/gcw0-toolchain/usr/mipsel-gcw0-linux-uclibc/sysroot/usr/include/libpng16 -D_REENTRANT' \
	 CGO_ENABLED=1 \
	 CGO_LDFLAGS='-L/opt/gcw0-toolchain/usr/mipsel-gcw0-linux-uclibc/sysroot/usr/lib -lSDL2 -lpng16' \
	 GOARCH=mipsle \
	 GOMIPS=softfloat \
	 GOOS=linux \
	 PKG_CONFIG='/opt/gcw0-toolchain/usr/bin/pkg-config' \
	 go build -o bin/pg2_test.gcw geniot.com/geniot/pg2_test_go/cmd/pg2_test

squash:
	mksquashfs bin/pg2_test.gcw resources/media/pg2test.png resources/default.gcw0.desktop bin/pg2_test.opk -all-root -no-xattrs -noappend -no-exports

opk: clean mips squash

#on PG2 use opkrun


