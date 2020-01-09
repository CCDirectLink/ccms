GOOS="linux" GOARCH="amd64" go build -o "./bin/ccmu" 
GOOS="windows" GOARCH="amd64" go build -o "./bin/ccmu.exe"
GOOS="darwin" GOARCH="amd64" go build -o "./bin/ccmu_mac"

cd bin



export PATH=$PATH:"/C/Program Files/7-zip/"

sha256sum ccmu_mac > checksum_sha256_mac.txt
7z.exe a -ttar -so -an ccmu_mac checksum_sha256_mac.txt | 7z.exe a -si ccmu_mac.tgz


sha256sum ccmu  > checksum_sha256_linux.txt
7z.exe a -ttar -so -an ccmu checksum_sha256_linux.txt | 7z.exe a -si ccmu_linux.tgz




cd ..
