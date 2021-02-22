package sharecalculate

type person struct{
	name string
	amtgiven float64
	amttaken float64
}

type persons struct{ 
	personsarr []person
} 

type sharedetail struct{
	name string
	sum float64
} 

type sharearr struct{ 
	sarr []sharedetail
} 

//main each persons expense/contribution
func (p *persons)updateperson(name string,amtgiven float64,amttaken float64){
     found := false 
     r := person{}
     for i,_ := range p.personsarr{ 
		if name == p.personsarr[i].name{ 
		 		
				 p.personsarr[i].amtgiven = p.personsarr[i].amtgiven + amtgiven 
				 p.personsarr[i].amttaken = p.personsarr[i].amttaken + amttaken 
			     found = true 
		 }
	 }   
     
	 if !found{ 
		 r.name = name
		 r.amtgiven= amtgiven
		 r.amttaken =amttaken
		 p.personsarr = append(p.personsarr,r)
	 } 
}

//find amt to be paid by each person to other 
func (p persons) calculateshare()output{ 
	  var a1, a2 sharearr 
	  s := sharedetail{}

	  for _,r := range p.personsarr{ 
		if r.amtgiven > r.amttaken{ 
			s.name = r.name 
			s.sum = r.amtgiven - r.amttaken
			a1 = a1.sorter(s)

		}else if r.amtgiven < r.amttaken{ 
			s.name = r.name 
			s.sum = r.amttaken - r.amtgiven
			a2 = a2.sorter(s)
		}
	  }
	  o:= a1.getpairing(a2)
	  return o 
}

func (a1 sharearr)getpairing(a2 sharearr) output{
  o := output{}
  out := a1.equalshare(&a2)  
  o.towrite = append(o.towrite,out.towrite...)
  out.towrite = nil
  for i:= 0;i<len(a1.sarr);i++{
	  for j:=0;j<len(a2.sarr);j++{ 
	   if a1.sarr[i].sum > a2.sarr[j].sum{
		 
    	 data := Orow{}
		 data.From = a2.sarr[j].name
		 data.To= a1.sarr[i].name
		 data.Amtinv = a2.sarr[j].sum
		 o.towrite = append(o.towrite,data)
		 a1.sarr[i].sum = a1.sarr[i].sum - a2.sarr[j].sum
		 copy(a2.sarr[j:],a2.sarr[j+1:]) 
		 a2.sarr = a2.sarr[:len(a2.sarr)-1] 
		 out = a1.equalshare(&a2)
		 o.towrite = append(o.towrite,out.towrite...)
		 out.towrite = nil
	  }
	 }
  }  
  
  for i:=0;i<len(a1.sarr);i++{ 	
	for j:=0;j<len(a2.sarr);j++{ 

	  if a1.sarr[i].sum  < a2.sarr[j].sum{
	
	     data := Orow{}
		 data.From = a2.sarr[j].name
		 data.To= a1.sarr[i].name
		 data.Amtinv = a1.sarr[i].sum
	     o.towrite = append(o.towrite,data)
	  	 a2.sarr[j].sum = a2.sarr[j].sum - a1.sarr[i].sum 
		 copy(a1.sarr[i:],a1.sarr[i+1:]) 
		 a1.sarr = a1.sarr[:len(a1.sarr)-1] 
		 out = a1.equalshare(&a2)
		 o.towrite = append(o.towrite,out.towrite...)
		 out.towrite  = nil
	  }  
   }
  } 

  if len(a1.sarr) > 0 && len(a2.sarr) > 0{
	out = a1.getpairing(a2)
  }

  o.towrite = append(o.towrite,out.towrite...)
  return o
}

func (a1 *sharearr)equalshare(a2 *sharearr ) output{ 

	o:= output{}
	for i := 0;i < len(a1.sarr);i++{ 
		for j:= 0;j < len(a2.sarr);j++{ 
		 
		 if a1.sarr[i].sum == a2.sarr[j].sum { 

		   data := Orow{}
		   data.From = a2.sarr[j].name
		   data.To= a1.sarr[i].name
		   data.Amtinv = a2.sarr[j].sum
		   o.towrite = append(o.towrite,data)
		   copy(a1.sarr[i:],a1.sarr[i+1:]) 
		   a1.sarr = a1.sarr[:len(a1.sarr)-1]
		   copy(a2.sarr[j:],a2.sarr[j+1:]) 
		   a2.sarr = a2.sarr[:len(a2.sarr)-1]
		 }
	   } 
	}
	return o
}

func (a sharearr)sorter(s sharedetail)sharearr{

  added := false 
  l := len(a.sarr)
  for i:=0;i<l;i++{ 
	  if a.sarr[i].sum >= s.sum{
		  a.sarr = append(a.sarr,s)
		  copy(a.sarr[i+1:],a.sarr[i:]) 
		  a.sarr[i] = s 
		  added = true
		  break
	  }
  }
  if !added{ 
	  a.sarr = append(a.sarr,s)
  }	
  return a 
}