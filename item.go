package sharecalculate

import ("strings"
"strconv"
"fmt"
)

type item struct{
    name string
    split []details
    total float64
    totweightage int
} 

type items struct{
     itemarr []item
}

//itemise each person expense and contribution
func (itemsdata items)updateshare()persons{

	var persondata persons
	for _,i := range itemsdata.itemarr{ 
		for j:=0;j<len(i.split);j++{ 
			//persondata = i.split[j].personmap(i.total,i.totweightage,persondata) 
			name,amtgiven,amttaken :=  i.split[j].personmap(i.total,i.totweightage)
			persondata.updateperson(name,amtgiven,amttaken)
		}
	}
	return persondata
}

//itemise the file data and build item array
func getitems(filerow []string) items{ 
	data := item{}
	datas := items{}
	firstitem := true
	var err error 
	for _,dm := range filerow{
		 a := strings.Split(dm,",")
		 if a[0] != "" && !firstitem{ 
		   datas.itemarr  = append(datas.itemarr,data)
		   data = item{}
		   data.name = a[0]
		  } 
		  if firstitem{ 
			data.name = a[0]
			firstitem = false
		  }
           dt := details{} 
		   dt.personname = a[1] 
		   dt.amtgiven,err =  strconv.ParseFloat(a[2],64) 
		   dt.weightage,err = strconv.Atoi(a[3]) 
           if err != nil{
			   fmt.Println("error")
		   }
  		   data.total = data.total + dt.amtgiven
		   data.totweightage = data.totweightage + dt.weightage
		   data.split = append(data.split,dt) 
	}
	datas.itemarr  = append(datas.itemarr,data)
    return datas
}