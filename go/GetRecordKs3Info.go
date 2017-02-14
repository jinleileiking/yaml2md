package param


type T_req_GetRecordKs3Info{
    UniqueName string `form:"UniqueName"  valid:"true" `
    App string `form:"App"  valid:"true" `
    Pubdomain string `form:"Pubdomain"  valid:"true" `
    Stream string `form:"Stream"  valid:"true" `
    UnixTime string `form:"UnixTime"  valid:"true" `
    Vdoid string `form:"Vdoid"  valid:"true" `

}

type T_rsp_GetRecordKs3Info{
    Data struct {
        RetCode   int
        RetMsg    string
        LiveFileNameM3U8 string `json:"omitempty"` 
        PhraseFileNameMp4 string `json:"omitempty"` 
        
        
    }
}



