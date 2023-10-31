# GitHub Copilot: 
# In einem Makefile ist das erste Ziel, das definiert wird, das Standardziel. 
# Das bedeutet, wenn Sie make ohne Argumente ausführen, wird dieses Ziel ausgeführt. 
# In Ihrem Fall ist analyze das erste Ziel und somit das Standardziel.

# Wenn Sie ein anderes Ziel als Standardziel festlegen möchten, können Sie ein Ziel 
# namens default am Anfang Ihres Makefiles definieren und es von dem gewünschten Ziel 
# abhängig machen. Hier ist ein Beispiel, wie Sie build als Standardziel festlegen können:
default: build

analyze:
	@go vet ./...
	@go run honnef.co/go/tools/cmd/staticcheck@latest --checks=all ./...

build: analyze test
	@mkdir -p ./build
	@go build -o ./build/tictactoe

coverage: test
	@mkdir -p ./coverage
	@go test -coverprofile=./coverage/coverage.out ./...
	@go tool cover -html=./coverage/coverage.out -o ./coverage/coverage.html
	@open ./coverage/coverage.html

test: analyze
	@go test -cover ./...

# GitHub Copilot: 
# `.PHONY` ist eine spezielle Zielmarkierung in Makefiles, die dazu dient, zu erklären, 
# dass das genannte Ziel nicht mit einer Datei oder einem Verzeichnis übereinstimmt. 

# In Ihrem Fall:
.PHONY: analyze build coverage test 
# sagt es `make`, dass `analyze`, `build`, `coverage` und `test` keine Dateien sind. 

# Dies ist nützlich, wenn Sie ein Ziel haben, das den gleichen Namen wie eine Datei oder ein Verzeichnis hat. 
# Ohne `.PHONY` würde `make` das Ziel nicht ausführen, wenn eine Datei oder ein Verzeichnis mit dem 
# gleichen Namen existiert und aktuell ist (d.h., es hat sich seit der letzten Ausführung von `make` nicht geändert).

# Mit `.PHONY` teilen Sie `make` mit, dass es das Ziel unabhängig vom Vorhandensein oder Status einer gleichnamigen 
# Datei oder eines Verzeichnisses ausführen soll.