package sharecalculate

import ("os"
"bufio"
"fmt"
"strconv"
)

type filereader struct{
    filepointer *os.File
} 

// read the file having expenses and share details and store in filerow
func (fr filereader)read()[]string{
		
	scanner := bufio.NewScanner(fr.filepointer)
    scanner.Split(bufio.ScanLines) 
    var filerow []string
	
	if scanner.Scan(){ 
		scanner.Text()
		
	}
	for scanner.Scan() {
		filerow = append(filerow, scanner.Text())
	}
	return filerow
} 

func openfile()filereader{

    fi,err := os.OpenFile("expense.csv", os.O_RDWR,0755) 
    if err != nil{ 
        fmt.Println("Error in file open", err)
	}  
  return filereader{filepointer:fi}
}

//write the file per person share exchange to file 
func (fr filereader)write(o output){
	  
	  for _,s := range o.towrite{ 
		data := s.From +  "  gives  " + s.To + " $" + strconv.FormatFloat(s.Amtinv,'f',2,64)+"\n"   
		fr.filepointer.WriteString(data)
	  }
}

func (fr filereader)closefile(){
	fr.filepointer.Close()
}