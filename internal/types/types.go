package types

type Student struct{
	Id int 
	Name string `validate:"required"`
	Gmail string `validate:"required"`
	Age int  `validate:"required"`
} 