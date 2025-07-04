# Erstellt ein Build für das aktuelle Hostsystem
go build -o build/m.gowiese

# ?! Laut miqt soll so gebaut werden, damit die Binary kleiner werden
# go build -o build/m.gowiese -ldflags '-s -w'

# Crossplatform Builds !

# Erstellt ein Build für Windows 64Bit
# GOOS=windows GOOARCH=amd64 go build -o build/m.gowiese.exe

# Erstellt ein Web Assembly Build
# GOOS=js GOARCH=wasm go build -o build/m.gowiese.wasm
