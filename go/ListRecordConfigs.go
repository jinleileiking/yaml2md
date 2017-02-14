package param


type T_req_ListRecordConfigs{
    UniqueName string `form:"UniqueName"  valid:"true" `
    App string `form:"App"  valid:"true" `
    Pubdomain string `form:"Pubdomain"  valid:"true" `

}

type T_rsp_ListRecordConfigs{
    Data struct {
        RetCode   int
        RetMsg    string
        PubdomainConfigs struct {
        App string `json:"omitempty"` 
        Pubdomain string `json:"omitempty"` 
        Mp4VodEnable string `json:"omitempty"` 
        HlsVodEnable string `json:"omitempty"` 
        HlsBucket string `json:"omitempty"` 
        Mp4Bucket string `json:"omitempty"` 
        MergeEnable string `json:"omitempty"` 
        LiveFileNameM3U8 string `json:"omitempty"` 
        
        }
    }
}



