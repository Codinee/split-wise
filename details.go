package sharecalculate 

type details struct{ 
    personname string 
    amtgiven float64 
    weightage int 
} 
type Orow struct{
	From string `json:"fromperson"`
	To string `json:"toperson"`
	Amtinv float64 `json:"amtinvolved"` 
}

type output struct{
	towrite []Orow
}

func (dt details) personmap(totalamt float64, totalwt int) (string,float64,float64){ 

	 amttaken := totalamt*float64(dt.weightage)/float64(totalwt)
	 //p = updateperson(dt.personname,dt.amtgiven,amttaken,p)
	 return dt.personname,dt.amtgiven,amttaken 
}
