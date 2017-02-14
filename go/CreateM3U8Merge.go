package param


type T_req_CreateM3U8Merge{
    RecID string `form:"RecID"  valid:"true" `

}

type T_rsp_CreateM3U8Merge{
    Data struct {
        RetCode   int
        RetMsg    string
        
        
    }
}



