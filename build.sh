#!/bin/sh

TMP_FOLDER=tmp
SOURCE=pg2test.go
TARGET=pg2test.gcw
OPK_NAME=pg2test.opk

rm -rf ${TMP_FOLDER}
mkdir ${TMP_FOLDER}

CC='/opt/gcw0-toolchain/usr/bin/mipsel-gcw0-linux-uclibc-gcc' \
 CGO_CFLAGS='-I/opt/gcw0-toolchain/usr/mipsel-gcw0-linux-uclibc/sysroot/usr/include -D_REENTRANT' \
 CGO_ENABLED=1 \
 CGO_LDFLAGS='-L/opt/gcw0-toolchain/usr/mipsel-gcw0-linux-uclibc/sysroot/usr/lib -lSDL2' \
 GOARCH=mipsle \
 GOMIPS=softfloat \
 GOOS=linux \
 PKG_CONFIG='/opt/gcw0-toolchain/usr/bin/pkg-config' \
 go build -o ${TMP_FOLDER}/${TARGET} ${SOURCE}

echo ${OPK_NAME}

# create default.gcw0.desktop
cat > ${TMP_FOLDER}/default.gcw0.desktop <<EOF
[Desktop Entry]
Name=PocketGo2 test
Comment=Test your PocketGo2
Exec=pg2test.gcw
Terminal=false
Type=Application
StartupNotify=true
Icon=pg2test
Categories=applications;
EOF

# create opk
FLIST="media"
FLIST="${FLIST} ${TMP_FOLDER}/pg2test.gcw"
FLIST="${FLIST} pg2test.png"
FLIST="${FLIST} ${TMP_FOLDER}/default.gcw0.desktop"

mksquashfs ${FLIST} ${TMP_FOLDER}/${OPK_NAME} -all-root -no-xattrs -noappend -no-exports