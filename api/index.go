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

type WebhookBot struct {
	Token string
	Bot   *tb.Bot
}

func (wb *WebhookBot) Setup() {
	b, err := tb.NewBot(tb.Settings{
		Token:       wb.Token,
		Synchronous: true,
	})

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
		return
	}
	// set instance bot
	wb.Bot = b

	wb.Bot.Handle("/start", wb.handlerStart)
	wb.Bot.Handle("/syarat", wb.handlerSyarat)
	wb.Bot.Handle(tb.OnText, func(m *tb.Message) {
		_, _ = b.Send(m.Sender, "Maaf bos, ga ngerti!")
	})
}

func (wb *WebhookBot) handlerStart(m *tb.Message) {
	msg := `Selamat datang demo dispenduk bot
	List bot command:
	/syarat <tipe>: Syarat pembuatan dokumen
	
	---
	`
	wb.Bot.Send(m.Sender, msg, &tb.SendOptions{
		ParseMode:             tb.ModeMarkdown,
		DisableWebPagePreview: true,
	})
}

func (wb *WebhookBot) handlerSyarat(m *tb.Message) {
	response := fmt.Sprintf("Syarat pembuatan dokumen %s tidak ditemukan", m.Payload)
	if v, ok := syarat[strings.ToLower(m.Payload)]; ok {
		response = v
	}

	wb.Bot.Send(m.Sender, response, &tb.SendOptions{
		ParseMode:             tb.ModeMarkdown,
		DisableWebPagePreview: true,
	})
}

func NewWebhookBot(token string) *WebhookBot {
	return &WebhookBot{Token: token}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	b := NewWebhookBot(os.Getenv("TELEGRAM_TOKEN"))
	b.Setup()

	var u tb.Update

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
	}

	if err = json.Unmarshal(body, &u); err == nil {
		b.Bot.ProcessUpdate(u)
	}
}
