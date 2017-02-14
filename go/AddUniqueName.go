package param


type T_req_AddUniqueName{
    UniqueName string `form:"UniqueName"  valid:"true" `

}

type T_rsp_AddUniqueName{
    Data struct {
        RetCode   int
        RetMsg    string
        
        
    }
}



