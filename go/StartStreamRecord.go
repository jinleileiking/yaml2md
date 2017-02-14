package param


type T_req_StartStreamRecord{
    UniqueName string `form:"UniqueName"  valid:"true" `
    App string `form:"App"  valid:"true" `
    Pubdomain string `form:"Pubdomain"  valid:"true" `
    Stream string `form:"Stream"  valid:"true" `
    Mp4VodEnable string `form:"Mp4VodEnable" `
    Ks3FileNameM3U8 string `form:"Ks3FileNameM3U8" `
    Ks3FullPathMP4 string `form:"Ks3FullPathMP4" `

}





