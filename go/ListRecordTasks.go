package param


type T_req_ListRecordTasks{
    UniqueName string `form:"UniqueName"  valid:"true" `
    App string `form:"App"  valid:"true" `
    Pubdomain string `form:"Pubdomain"  valid:"true" `
    Stream string `form:"Stream"  valid:"true" `
    RecType string `form:"RecType" `
    OrderTime Int `form:"OrderTime" `
    Marker string `form:"Marker" `
    Limit string `form:"Limit" `
    StartUnixTime string `form:"StartUnixTime" `
    EndUnixTime string `form:"EndUnixTime" `

}

type T_rsp_ListRecordTasks{
    Data struct {
        RetCode   int
        RetMsg    string
        UniqueName string `json:"omitempty"` 
        App string `json:"omitempty"` 
        Pubdomain string `json:"omitempty"` 
        Stream string `json:"omitempty"` 
        OrderTime Int `json:"omitempty"` 
        RecType string `json:"omitempty"` 
        Total Int `json:"omitempty"` 
        Count Int `json:"omitempty"` 
        Recs struct {
        RecID string `json:"omitempty"` 
        RecStatus string `json:"omitempty"` 
        RecType string `json:"omitempty"` 
        StartUnixTime string `json:"omitempty"` 
        EndUnixTime string `json:"omitempty"` 
        Ks3FullPathMp4 string `json:"omitempty"` 
        Ks3FullPathM3u8 string `json:"omitempty"` 
        
        }
    }
}



