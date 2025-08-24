# Nome do binário (sem extensão)
APP := projeto-go

# Número padrão de workers para o modo concorrente (pode sobrescrever: make run-conc WORKERS=12)
WORKERS ?= 8

# Ajustes por SO (extensão .exe no Windows, comando de delete e como executar o binário)
ifeq ($(OS),Windows_NT)
  EXT := .exe
  RM := del /Q
  RUN := .\\
else
  EXT :=
  RM := rm -f
  RUN := ./
endif

BIN := $(APP)$(EXT)

.PHONY: build run-seq run-conc run-build-seq run-build-conc clean fmt vet

## --------- Alvos principais ---------

# Compila o binário
build:
	go build -o $(BIN)

# Roda sem compilar (sequencial)
run-seq:
	go run . -mode=seq

# Roda sem compilar (concorrente) - use WORKERS para ajustar
run-conc:
	go run . -mode=conc -workers $(WORKERS)

# Compila e roda o binário (sequencial)
run-build-seq: build
	$(RUN)$(BIN) -mode=seq

# Compila e roda o binário (concorrente)
run-build-conc: build
	$(RUN)$(BIN) -mode=conc -workers $(WORKERS)

# Limpa o binário
clean:
	-$(RM) $(BIN)

## --------- Qualidade (opcionais) ---------

# Formata o código
fmt:
	go fmt ./...

# Vetoriza/análises estáticas simples
vet:
	go vet ./...
