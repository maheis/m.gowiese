# Erstellt ein Build für das aktuelle Hostsystem
go build -o bin/m.gowiese

# Crossplatform Builds !

# Erstellt ein Build für Windows 64Bit
GOOS=windows GOOARCH=amd64 go build -o bin/m.gowiese.exe

# Erstellt ein Web Assembly Build
GOOS=js GOARCH=wasm go build -o bin/m.gowiese.wasm
