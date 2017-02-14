package param


type T_req_GetRecordTask{
    RecID string `form:"RecID"  valid:"true" `

}

type T_rsp_GetRecordTask{
    Data struct {
        RetCode   int
        RetMsg    string
        UniqueName string `json:"omitempty"` 
        App string `json:"omitempty"` 
        Pubdomain string `json:"omitempty"` 
        Stream string `json:"omitempty"` 
        RecStatus string `json:"omitempty"` 
        RecType string `json:"omitempty"` 
        StartUnixTime string `json:"omitempty"` 
        EndUnixTime string `json:"omitempty"` 
        Ks3FullPathMp4 string `json:"omitempty"` 
        Ks3FileNameM3u8 string `json:"omitempty"` 
        
        
    }
}



