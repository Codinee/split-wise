package sharecalculate

type iooperator interface{ 
    read() []string
    write(string)
}