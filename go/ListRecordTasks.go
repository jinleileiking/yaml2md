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





