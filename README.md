## Features

- Basic arithmetic operations (ADD, SUB, MUL, DIV)
- Variable loading and storing (LOAD, STORE)
- Control flow (JMP, JZ, JNZ)
- Stack-based execution
- Halt instruction

## Installation

```bash
git clone https://github.com/WillKirkmanM/bytecode
cd bytecode
go build -o bytecode cmd/interpreter/main.go
```

## Usage

### Run a bytecode file:
```bash
./bytecode examples/hello.bc
```

### Interactive mode:
```bash
./bytecode
```

## Instruction Set

| Opcode | Instruction | Description |
|--------|-------------|-------------|
| 0x00   | HALT        | Stop execution |
| 0x01   | LOAD_CONST  | Load constant to stack |
| 0x02   | LOAD_VAR    | Load variable to stack |
| 0x03   | STORE_VAR   | Store top of stack to variable |
| 0x04   | ADD         | Add two values from stack |
| 0x05   | SUB         | Subtract two values from stack |
| 0x06   | MUL         | Multiply two values from stack |
| 0x07   | DIV         | Divide two values from stack |
| 0x08   | JMP         | Jump to address |
| 0x09   | JZ          | Jump if zero |
| 0x0A   | JNZ         | Jump if not zero |
| 0x0B   | PRINT       | Print top of stack |

## Examples

See the `examples/` directory for sample programs demonstrating various features of the interpreter.