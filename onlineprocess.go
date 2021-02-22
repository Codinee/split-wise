package sharecalculate


import("net/http"
	"encoding/json"
)

type Oitem struct { 
	Iname string `json:"itemname"`
	Fname string `json:"friendname"`
	Paid float64 `json:"paidamt"`
	Weight int `json:"weightage"`
}

type itemslist struct{
	list []Oitem
}


func (i *itemslist) myhandler(w http.ResponseWriter, r *http.Request){ 
	   switch r.URL.Path{
		  case "/submit":
			    i := itemslist{}
				if r.Method == "POST"{ 
					json.NewDecoder(r.Body).Decode(&i.list)
					itemsdata := i.movedata() 
					persondata := itemsdata.updateshare()
					writedata := persondata.calculateshare() 
					json.NewEncoder(w).Encode(writedata.towrite)
				}
     	}

}

func (i *itemslist)movedata() items{ 
	previtem := "" 
	data := item{}
	datas := items{}
	for _,ilist := range i.list{
        if ilist.Iname != previtem && previtem != ""{ 
			datas.itemarr = append(datas.itemarr,data)
			data = item{}     		
		} 
		data.name = ilist.Iname
		data.total = data.total + ilist.Paid
		data.totweightage = data.totweightage + ilist.Weight
		d := details{}
		d.personname =  ilist.Fname 
	    d.weightage = ilist.Weight
		d.amtgiven = ilist.Paid
		data.split = append(data.split,d) 
		previtem = ilist.Iname
	}
	datas.itemarr = append(datas.itemarr,data)
    return datas
}

func RunAPI(){ 
  myitem := itemslist{}
  http.HandleFunc("/submit",myitem.myhandler)
  http.ListenAndServe(":8080",nil)
}
