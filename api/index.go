package handler

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"os"
// 	"strings"

// 	tb "gopkg.in/tucnak/telebot.v2"
// )

// var syarat = map[string]string{"ktp": `*Syarat Pembuatan KTP*:
// - Berusia 17 tahun
// - Surat pengantar dari pihak Rukun Tetangga (RT) dan Rukun Warga (RW)
// - Fotokopi Kartu Keluarga (KK)
// - Surat keterangan pindah dari kota asal, jika Anda bukan asli warga setempat
// - Surat keterangan pindah dari luar negeri, dan surat ini harus diterbitkan oleh Instansi Pelaksana bagi Warga Negara Indonesia (WNI) yang datang dari luar negeri karena pindah.
// - Datang langsung ke kantor Keluruhan, di sini pula Anda akan diambil fotonya dan melakukan sidik jari.
// `,
// 	"kk": `*Kartu Keluarga Baru Bagi Pasangan Baru*:

// - Surat pengantar dari RT yang sudah distempel oleh RW
// - Fotokopi buku nikah atau akta perkawinan
// - Surat keterangan pindah (khusus untuk anggota keluarga pendatang)

// *Penambahan Anggota Keluarga*

// - Surat pengantar dari RT atau RW setempat
// - Kartu keluarga (KK) yang lama
// - Surat keterangan kelahiran anggota keluarga baru yang akan ditambahkan

// *Penambahan Anggota Keluarga (Numpang KK)*

// - Surat pengantar RT atau RW daerah setempat
// - Kartu keluarga (KK) yang lama atau kartu keluarga yang akan ditumpangi
// - Surat keterangan pindah datang (jika tidak satu daerah)
// - Surat keterangan datang dari luar negeri (bagi WNI dari luar negeri)
// - Paspor, izin tinggal tetap, dan surat keterangan catatan kepolisian (SKCK) atau surat tanda lapor diri (bagi WNA)

// *Pengurangan Anggota Keluarga pada KK*

// - Surat pengantar dari RT atau RW daerah setempat
// - Kartu keluarga (KK) yang lama
// - Surat keterangan kematian (bagi anggota keluarga yang meninggal dunia)
// - Jika pengurangan terjadi karena ada anggota keluarga yang pindah, maka surat keterangan kematian diganti dengan surat keterangan pindah (bagi anggota keluarga yang pindah).
// `}

// type WebhookBot struct {
// 	Token string
// 	Bot   *tb.Bot
// }

// func (wb *WebhookBot) Setup() {
// 	b, err := tb.NewBot(tb.Settings{
// 		Token:       wb.Token,
// 		Synchronous: true,
// 	})

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	// set instance bot
// 	wb.Bot = b

// 	wb.Bot.Handle("/start", wb.handlerStart)
// 	wb.Bot.Handle("/syarat", wb.handlerSyarat)
// 	wb.Bot.Handle(tb.OnText, func(m *tb.Message) {
// 		_, _ = b.Send(m.Sender, "Maaf bos, ga ngerti!")
// 	})
// }

// func (wb *WebhookBot) handlerStart(m *tb.Message) {
// 	msg := `Selamat datang demo dispenduk bot
// 	List bot command:
// 	/syarat <tipe>: Syarat pembuatan dokumen

// 	---
// 	`
// 	wb.Bot.Send(m.Sender, msg, &tb.SendOptions{
// 		ParseMode:             tb.ModeMarkdown,
// 		DisableWebPagePreview: true,
// 	})
// }

// func (wb *WebhookBot) handlerSyarat(m *tb.Message) {
// 	response := fmt.Sprintf("Syarat pembuatan dokumen %s tidak ditemukan", m.Payload)
// 	if v, ok := syarat[strings.ToLower(m.Payload)]; ok {
// 		response = v
// 	}

// 	wb.Bot.Send(m.Sender, response, &tb.SendOptions{
// 		ParseMode:             tb.ModeMarkdown,
// 		DisableWebPagePreview: true,
// 	})
// }

// func NewWebhookBot(token string) *WebhookBot {
// 	return &WebhookBot{Token: token}
// }

// func Handler(w http.ResponseWriter, r *http.Request) {
// 	b := NewWebhookBot(os.Getenv("TELEGRAM_TOKEN"))
// 	b.Setup()

// 	var u tb.Update

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	if err = json.Unmarshal(body, &u); err == nil {
// 		b.Bot.ProcessUpdate(u)
// 	}
// }

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"time"
)

type Hito struct {
	URI   string
	Title string
	fecha time.Time
}

type Response struct {
	Msg    string `json:"text"`
	ChatID int64  `json:"chat_id"`
	Method string `json:"method"`
}

var hitos = []Hito{
	Hito{
		URI:   "0.Repositorio",
		Title: "Datos básicos y repo",
		fecha: time.Date(2020, time.September, 29, 11, 30, 0, 0, time.UTC),
	},
	Hito{
		URI:   "1.Infraestructura",
		Title: "HUs y entidad principal",
		fecha: time.Date(2020, time.October, 6, 11, 30, 0, 0, time.UTC),
	},
	Hito{
		URI:   "2.Tests",
		Title: "Tests iniciales",
		fecha: time.Date(2020, time.October, 16, 11, 30, 0, 0, time.UTC),
	},
	Hito{
		URI:   "3.Contenedores",
		Title: "Contenedores",
		fecha: time.Date(2020, time.October, 26, 11, 30, 0, 0, time.UTC),
	},
	Hito{
		URI:   "4.CI",
		Title: "Integración continua",
		fecha: time.Date(2020, time.November, 6, 23, 59, 0, 0, time.UTC),
	},
	Hito{
		URI:   "5.Serverless",
		Title: "Trabajando con funciones serverless",
		fecha: time.Date(2020, time.November, 24, 11, 30, 0, 0, time.UTC),
	},
}

func Handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	var update tgbotapi.Update
	if err := json.Unmarshal(body, &update); err != nil {
		log.Fatal("Error en el update →", err)
	}
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	currentTime := time.Now()
	var next int
	var queda time.Duration
	for indice, hito := range hitos {
		if hito.fecha.After(currentTime) {
			next = indice
			queda = hito.fecha.Sub(currentTime)
		}
	}
	if update.Message.IsCommand() {
		text := ""
		if next == 0 {
			text = "Ninguna entrega próxima"
		} else {

			switch update.Message.Command() {
			case "kk":
				text = queda.String()
			case "kekeda":
				text = fmt.Sprintf("→ Próximo hito %s\n🔗 https://jj.github.io/IV/documentos/proyecto/%s\n📅 %s",
					hitos[next].Title,
					hitos[next].URI,
					hitos[next].fecha.String(),
				)
			default:
				text = "Usa /kk para lo que queda para el próximo hito, /kekeda para + detalle"
			}
		}
		data := Response{Msg: text,
			Method: "sendMessage",
			ChatID: update.Message.Chat.ID}

		msg, _ := json.Marshal(data)
		log.Printf("Response %s", string(msg))
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, string(msg))
	}
}
