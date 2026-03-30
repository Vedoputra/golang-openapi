package handlers

import (
	"golang-api-crud/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetAllSiswa godoc
//
//	@Summary		Ambil semua data siswa
//	@Description	List semua siswa (dummy in-memory)
//	@Tags			Siswa
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Siswa
//	@Failure		500	{object}	models.ErrorResponse
//	@Router			/siswa [get]
func GetAllSiswa(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(models.DataSiswa)
}

// GetSiswaByID godoc
//
//	@Summary		Ambil siswa berdasarkan ID
//	@Description	Detail siswa by ID
//	@Tags			Siswa
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"ID Siswa"
//	@Success		200	{object}	models.Siswa
//	@Failure		400	{object}	models.ErrorResponse
//	@Failure		404	{object}	models.ErrorResponse
//	@Router			/siswa/{id} [get]
func GetSiswaByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Message: "ID tidak valid"})
	}

	for _, s := range models.DataSiswa {
		if s.ID == id {
			return c.Status(fiber.StatusOK).JSON(s)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{Message: "Siswa tidak ditemukan"})
}

// CreateSiswa godoc
//
//	@Summary		Tambah siswa baru
//	@Description	Create siswa
//	@Tags			Siswa
//	@Accept			json
//	@Produce		json
//	@Param			siswa	body		models.CreateSiswaRequest	true	"Data siswa"
//	@Success		201		{object}	models.Siswa
//	@Failure		400		{object}	models.ErrorResponse
//	@Router			/siswa [post]
func CreateSiswa(c *fiber.Ctx) error {
	var req models.CreateSiswaRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Message: "Body tidak valid"})
	}
	if req.Nama == "" || req.Kelas == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Message: "nama & kelas wajib diisi"})
	}

	newSiswa := models.Siswa{
		ID:    models.IDCounter,
		Nama:  req.Nama,
		Kelas: req.Kelas,
		Umur:  req.Umur,
	}
	models.IDCounter++
	models.DataSiswa = append(models.DataSiswa, newSiswa)

	return c.Status(fiber.StatusCreated).JSON(newSiswa)
}

// UpdateSiswa godoc
//
//	@Summary		Update siswa
//	@Description	Update siswa by ID
//	@Tags			Siswa
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int						true	"ID Siswa"
//	@Param			siswa	body		models.UpdateSiswaRequest	true	"Data update siswa"
//	@Success		200		{object}	models.Siswa
//	@Failure		400		{object}	models.ErrorResponse
//	@Failure		404		{object}	models.ErrorResponse
//	@Router			/siswa/{id} [put]
func UpdateSiswa(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Message: "ID tidak valid"})
	}

	var req models.UpdateSiswaRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Message: "Body tidak valid"})
	}

	for i, s := range models.DataSiswa {
		if s.ID == id {
			if req.Nama != "" {
				models.DataSiswa[i].Nama = req.Nama
			}
			if req.Kelas != "" {
				models.DataSiswa[i].Kelas = req.Kelas
			}
			if req.Umur != 0 {
				models.DataSiswa[i].Umur = req.Umur
			}
			return c.Status(fiber.StatusOK).JSON(models.DataSiswa[i])
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{Message: "Siswa tidak ditemukan"})
}

// DeleteSiswa godoc
//
//	@Summary		Hapus siswa
//	@Description	Delete siswa by ID
//	@Tags			Siswa
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"ID Siswa"
//	@Success		200	{object}	map[string]string
//	@Failure		400	{object}	models.ErrorResponse
//	@Failure		404	{object}	models.ErrorResponse
//	@Router			/siswa/{id} [delete]
func DeleteSiswa(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Message: "ID tidak valid"})
	}

	for i, s := range models.DataSiswa {
		if s.ID == id {
			models.DataSiswa = append(models.DataSiswa[:i], models.DataSiswa[i+1:]...)
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Siswa berhasil dihapus"})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{Message: "Siswa tidak ditemukan"})
}