GOOS="linux" GOARCH="amd64" go build -o "./bin/ccmu" 
GOOS="windows" GOARCH="amd64" go build -o "./bin/ccmu.exe"
GOOS="darwin" GOARCH="amd64" go build -o "./bin/ccmu_mac"

cd bin

printf 'ccmu %s\n' $(certutil -hashfile ccmu SHA256 |& head -2 |& tail -1) >> checksum_sha256.txt
printf 'ccmu.exe %s\n' $(certutil -hashfile ccmu.exe SHA256 |& head -2 |& tail -1) >> checksum_sha256.txt
printf 'ccmu_mac %s\n' $(certutil -hashfile ccmu_mac SHA256 |& head -2 |& tail -1) >> checksum_sha256.txt