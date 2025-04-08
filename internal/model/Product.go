package model

/**
  Notice declaration here: ID    int     `json:"id"`
  ID is in caps because we want it to be public field. Now this filed would begenerated from json response
  which would be something like { "id" : 1} for example.
  These two fields would never match because its case sensitive. So to fix that we have added tags
  `json: id` This struct tag is telling GO on how to map between the json field and this struct field.
  Go supports many other tags like db:"..." or xml:"..."

  Go uses reflect.StructTag.Get("json") under the hood
**/
type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Rating      Rating  `json:"rating"`
}

type Rating struct {
	Rate  float64 `json:"rate"`
	Count int     `json:"count"`
}
