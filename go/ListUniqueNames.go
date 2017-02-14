package param


type T_req_ListUniqueNames{

}

type T_rsp_ListUniqueNames{
    Data struct {
        RetCode   int
        RetMsg    string
        UniqueNames String Array `json:"omitempty"` 
        
        
    }
}



