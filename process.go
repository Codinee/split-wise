package sharecalculate
// Controls the file process of calculating share for each person 
func Runfirst(){
	
	fp := openfile()
    filedata :=  fp.read()
	itemsdata := getitems(filedata)
	persondata := itemsdata.updateshare()
	writedata := persondata.calculateshare() 
	fp.write(writedata)
	fp.closefile()
 
}