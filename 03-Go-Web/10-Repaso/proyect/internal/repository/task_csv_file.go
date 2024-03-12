package repository

type TaskCSV struct {
	ID    int     `csv:"id"`
	Name  string  `csv:"name"`
	Price float64 `csv:"price"`
}

type TaskCSVFile struct {
	FileName string
}