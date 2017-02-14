package param


type T_req_UpdateM3U8Merge{
    RecID string `form:"RecID"  valid:"true" `
    Status string `form:"Status"  valid:"true" `

}

type T_rsp_UpdateM3U8Merge{
    Data struct {
        RetCode   int
        RetMsg    string
        
        
    }
}



