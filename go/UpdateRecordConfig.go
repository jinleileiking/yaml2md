package param


type T_req_UpdateRecordConfig{
    UniqueName string `form:"UniqueName"  valid:"true" `
    App string `form:"App"  valid:"true" `
    Pubdomain string `form:"Pubdomain"  valid:"true" `
    Mp4VodEnable string `form:"Mp4VodEnable" `
    MergeEnable string `form:"MergeEnable" `
    HlsBucket string `form:"HlsBucket" `
    Mp4Bucket string `form:"Mp4Bucket" `
    LiveFileNameM3U8 string `form:"LiveFileNameM3U8" `

}

type T_rsp_UpdateRecordConfig{
    Data struct {
        RetCode   int
        RetMsg    string
        
        
    }
}



