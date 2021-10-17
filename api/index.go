package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	tb "gopkg.in/tucnak/telebot.v2"
)

var syarat = map[string]string{"ktp": `*Syarat Pembuatan KTP*:
- Berusia 17 tahun
- Surat pengantar dari pihak Rukun Tetangga (RT) dan Rukun Warga (RW)
- Fotokopi Kartu Keluarga (KK)
- Surat keterangan pindah dari kota asal, jika Anda bukan asli warga setempat
- Surat keterangan pindah dari luar negeri, dan surat ini harus diterbitkan oleh Instansi Pelaksana bagi Warga Negara Indonesia (WNI) yang datang dari luar negeri karena pindah.
- Datang langsung ke kantor Keluruhan, di sini pula Anda akan diambil fotonya dan melakukan sidik jari.
`,
	"kk": `*Kartu Keluarga Baru Bagi Pasangan Baru*:

- Surat pengantar dari RT yang sudah distempel oleh RW
- Fotokopi buku nikah atau akta perkawinan
- Surat keterangan pindah (khusus untuk anggota keluarga pendatang)
 
*Penambahan Anggota Keluarga*

- Surat pengantar dari RT atau RW setempat
- Kartu keluarga (KK) yang lama
- Surat keterangan kelahiran anggota keluarga baru yang akan ditambahkan
 
*Penambahan Anggota Keluarga (Numpang KK)*

- Surat pengantar RT atau RW daerah setempat
- Kartu keluarga (KK) yang lama atau kartu keluarga yang akan ditumpangi
- Surat keterangan pindah datang (jika tidak satu daerah)
- Surat keterangan datang dari luar negeri (bagi WNI dari luar negeri)
- Paspor, izin tinggal tetap, dan surat keterangan catatan kepolisian (SKCK) atau surat tanda lapor diri (bagi WNA)
 
*Pengurangan Anggota Keluarga pada KK*

- Surat pengantar dari RT atau RW daerah setempat
- Kartu keluarga (KK) yang lama
- Surat keterangan kematian (bagi anggota keluarga yang meninggal dunia)
- Jika pengurangan terjadi karena ada anggota keluarga yang pindah, maka surat keterangan kematian diganti dengan surat keterangan pindah (bagi anggota keluarga yang pindah).
`}

func Handler(w http.ResponseWriter, r *http.Request) {
	b, err := tb.NewBot(tb.Settings{
		Token:       os.Getenv("TELEGRAM_TOKEN"),
		Synchronous: true,
	})

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(m *tb.Message) {
		msg := `Selamat datang demo dispenduk bot
		List bot command:
		/syarat: Syarat pembuatan dokumen. contoh "/syarat kk"
		
		---
		`
		b.Send(m.Sender, msg, &tb.SendOptions{
			ParseMode:             tb.ModeMarkdown,
			DisableWebPagePreview: true,
		})
	})
	b.Handle("/syarat", func(m *tb.Message) {
		response := fmt.Sprintf("Syarat pembuatan dokumen %s tidak ditemukan", m.Payload)
		if v, ok := syarat[strings.ToLower(m.Payload)]; ok {
			response = v
		}

		b.Send(m.Sender, response, &tb.SendOptions{
			ParseMode:             tb.ModeMarkdown,
			DisableWebPagePreview: true,
		})
	})
	b.Handle(tb.OnText, func(m *tb.Message) {
		_, _ = b.Send(m.Sender, "Maaf bos, ga ngerti!")
	})

	var u tb.Update

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
	}

	if err = json.Unmarshal(body, &u); err == nil {
		b.ProcessUpdate(u)
	}
}
