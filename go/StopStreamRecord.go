package param


type T_req_StopStreamRecord{
    UniqueName string `form:"UniqueName"  valid:"true" `
    App string `form:"App"  valid:"true" `
    Pubdomain string `form:"Pubdomain"  valid:"true" `
    Stream string `form:"Stream"  valid:"true" `
    RecID string `form:"RecID"  valid:"true" `

}

type T_rsp_StopStreamRecord{
    Data struct {
        RetCode   int
        RetMsg    string
        
        
    }
}



