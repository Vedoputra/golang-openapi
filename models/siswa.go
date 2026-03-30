package models

type Siswa struct {
	ID    int    `json:"id" example:"1"`
	Nama  string `json:"nama" example:"Budi"`
	Kelas string `json:"kelas" example:"XI RPL 1"`
	Umur  int    `json:"umur" example:"17"`
}

type CreateSiswaRequest struct {
	Nama  string `json:"nama" example:"Budi"`
	Kelas string `json:"kelas" example:"XI RPL 1"`
	Umur  int    `json:"umur" example:"17"`
}

type UpdateSiswaRequest struct {
	Nama  string `json:"nama" example:"Budi Update"`
	Kelas string `json:"kelas" example:"XI RPL 2"`
	Umur  int    `json:"umur" example:"18"`
}

type ErrorResponse struct {
	Message string `json:"message" example:"Data tidak ditemukan"`
}

// dummy in-memory
var DataSiswa = []Siswa{
	{ID: 1, Nama: "Budi", Kelas: "XI RPL 1", Umur: 17},
	{ID: 2, Nama: "Siti", Kelas: "XI RPL 2", Umur: 16},
}
var IDCounter = 3