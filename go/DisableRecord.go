package param


type T_req_DisableRecord{
    UniqueName string `form:"UniqueName"  valid:"true" `
    App string `form:"App"  valid:"true" `
    Pubdomain string `form:"Pubdomain"  valid:"true" `

}

type T_rsp_DisableRecord{
    Data struct {
        RetCode   int
        RetMsg    string
        
        
    }
}



