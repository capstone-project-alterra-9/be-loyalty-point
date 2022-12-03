package repository

// import (
// 	"capstone-project/entity"

// 	"github.com/labstack/echo/v4"
// )


// func (p *Pagination) GetOffset() int {  
//     return (p.GetPage() - 1) * p.GetLimit() 
// }   

// func (p *Pagination) GetLimit() int {   
//     if p.Limit == 0 {   
//         p.Limit = 10    
//     }   
//     return p.Limit  
// }   

// func (p *Pagination) GetPage() int {    
//     if p.Page == 0 {    
//         p.Page = 1  
//     }   
//     return p.Page   
// }   

// func (p *Pagination) GetSort() string { 
//     if p.Sort == "" {   
//         p.Sort = "Id desc"  
//     }   
//     return p.Sort   
// }   

// func Paginate(value interface{}, pagination *repository.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {  
//     var totalRows int64 
//     db.Model(value).Count(&totalRows)   

//     pagination.TotalRows = totalRows    
//     totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))    
//     pagination.TotalPages = totalPages  

//     return func(db *gorm.DB) *gorm.DB { 
//         return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())   
//     }   
// }   
