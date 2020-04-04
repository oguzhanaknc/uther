package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"./imdb"
	"./messages"
	"./utils"
	"./wikipedia"

	tgbot "github.com/eternnoir/gotelebot"
	"github.com/eternnoir/gotelebot/types"
)

//TODO: Hatırlatıcı fonksyonunu yaz
type Todo struct {
	day  string
	time string
	job  string
	verb string
}
type hava struct {
	max         string
	min         string
	icon        string
	day         string
	description string
}

var days = [7]string{"pazartesi", "salı", "çarşamba", "perşembe", "cuma", "cumartesi", "pazar"}
var sadEmoji = [6]string{"😔", "😟", "😕", "😳", "🥺", "😩"}
var okayEmoji = [6]string{"🔥", "👌", "👉", "👍", "💪", "⏰"}

var bot = tgbot.InitTeleBot("1039012461:AAFEDMiz11PmWCpGpXWTo5QABo3yM3ciw_4")
var id int64 = 511092151
var chatID int64

//main : anafonksyon diğer iki fonksyonu paralel olarak çalıştırıyor
func main() {
	go bot.StartPolling(true, 60)
	go processNewMessage(bot)
	processNewInlineQuery(bot)

}

//processNewMEssage : bot üzerinden sürekli yeni mesajları takip eder
func processNewMessage(bot *tgbot.TeleBot) {
	newMsgChan := bot.Messages
	for {
		m := <-newMsgChan // Mesajı al yeni mesaj geldiğinde
		fmt.Printf("Get Message:%#v \n", m)
		if m.Text != "" { // Mesajın text mesajı olduğu  kontrol et

			go messageSender(int(m.Chat.Id), m.Text, m.Chat.FirstName)
		}
	}
}

// Gelen mesaja göre gerekli fonksyonları çalıştırır ve gerekli geri dönüşleri sağlar
func mainOrganizer(msg, name string) (string, string, string) {
	switch x := toArray(msg); x[0] {
	case "hatirlat", "hatırlat", "hatırla", "hatirla":
		t := newTodo(msg)
		if t.day != "null" && t.time != "null" {
			return okayEmoji[rand.Intn(len(okayEmoji)-1)] + " sana " + t.day + " " + t.time + "'da " + t.job + " " + t.verb + " hatırlatacağım.", "t-mp3", "done"
		}
	case "/start", "nasildi", "nasıldı":
		return name + messages.Messages("wellcome"), "t-mp3", "intro"
	case "/hatirlatici", "hatirlatici", "hatırlatıcı":
		return messages.Messages("reminder"), "text", "null"
	case "/havadurumu":
		return "Her gün sabah 09:00 da hava durumu bilgisini göndereceğim. İstediğin zaman ulaşmak için 'hava', 'hava durumu','hava nasıl' yada 'soğuk mu' yazarak ulaşabilirsin. ", "text", "null"
	case "hava", "hava durumu", "hava nasıl", "soğuk mu":
		msg, _ := weatherOrganizer(true)

		return msg, "text", "null"
	case "film":
		message, poster := movieOrganizer(imdb.Film(strings.Join(x[1:], ",")))
		return message, "null", poster
	case "/filmler":

		return "film boşluk film adı şeklinde yazdığın her filmin detaylarını iletebilirim. Örnek olarak: film Arrival", "text", "null"
	case "hi", "merhaba", "selam":
		return "Sanada merhaba 🤙", "text", "null"
	case "sağol", "teşekkürler", "süpersin":
		return "Her zaman 🤖", "text", "null"
	case "/wiki":
		return "aramak istediğin cümlenin sonuna 'nedir' veya 'kimdir' yazman yeterli", "text", "null"
	case "orada mısın", "ordamısın", "hey", "ordamisin":
		var file string = "at-your-service-sir"
		switch rand.Intn(4) {
		case 1:
			file = "ready"
		case 2:
			file = "at-your-service-sir"
		case 3:
			file = "ask0"
		}
		return "null", "mp3", file

	default:
		if strings.Contains(msg, "nedir") || strings.Contains(msg, "kimdir") {

			return search(msg), "text", "null"
		}
		return sadEmoji[rand.Intn(len(sadEmoji)-1)] + " ne demek istediğini anlamadım. unuttuysan nasıldı yazarak yardım alabilirsin. ", "t-mp3", "repeat"
	}
	return sadEmoji[rand.Intn(len(sadEmoji)-1)] + " ne demek istediğini anlamadım. unuttuysan nasıldı yazarak yardım alabilirsin. ", "t-mp3", "repeat"
}

//Yeni todo nesnesi oluşturacak fonksyon
func newTodo(msg string) *Todo {
	message := toArray(msg)
	t := new(Todo)
	t.day = witchDay(msg)
	t.job = strings.Join(message[3:len(message)-1], " ")
	t.time = timeSlover(message[2])
	t.verb = verbSlover(message[len(message)-1])
	return t
}

//hatırlatıcıda hangi günün yazıldığını arayan fonksyon
func witchDay(msg string) string {
	for i := 0; i < len(days)-1; i++ {
		if strings.Contains(msg, days[i]) {
			return days[i]
		}
	}
	return "null"
}

//string i array a çeviren fonksyon
func toArray(msg string) []string {
	return strings.Split(msg, " ")
}

//havadurumunu yöneten ve düzeleyen fonksyon
func weatherOrganizer(who bool) (string, string) {
	say := "Günaydın 🙌 "
	weather := new(hava)
	weather.day, weather.description, weather.icon, weather.max, weather.min = utils.Havadurumu()
	var icon string
	if who {
		say = ""
	}
	switch weather.description {
	case "kapalı":
		icon = "☁️"
	case "yağmurlu":
		icon = "🌧️"
	case "güneşli":
		icon = "🌞"
	case "parçalı bulutlu":
		icon = "⛅"
	case "açık":
		icon = "☀️"
	case "az bulutlu":
		icon = "⛅"
	case "hafif yağmur":
		icon = "🌧️"
	default:
	}

	max, _ := strconv.ParseFloat(weather.max, 64)
	max = math.Floor(max)
	var imax int = int(max)
	str := strconv.Itoa(imax)
	return say + "bu gün hava " + weather.description + " " + icon + " En fazla " + str + " en az " + weather.min + " derece.", str
}

//zamana göre işlem yapmayı sağlayan fonksyon
func timemainOrganizer() {

	for {
		today := time.Now()
		if string(today.Format("15:04:05")) == "21:12:40" {
			msg, file := weatherOrganizer(false)
			bot.SendVoice(int(id), "./morning.ogg", nil)
			bot.SendVoice(int(id), "./caged_temp_current_0.ogg", nil)
			bot.SendVoice(int(id), "./ ("+file+").ogg", nil)
			bot.SendVoice(int(id), "./caged_temp_c.ogg", nil)
			bot.SendMessage(int(id), msg, nil)
			go timemainOrganizer()
			break
		}
		time.Sleep(1 * time.Second)
	}
}

// aradığın filme göre düzenleme yapan fonksyon
func movieOrganizer(title, filmtype, year, votes, poster, ferr string) (string, string) {
	message := "🎬 Film Adı : " + title + "🎬 👨‍🎨Oyuncular : " + filmtype + "👩‍🎨 📆 Çekim yılı : " + year + "📆 Puanı : " + votes + " ✔"
	if ferr == "Movie not found!" {
		return "nil", "404"
	}

	return message, poster

}

//hatırlatıcı için yüklem mapping
func verbSlover(verb string) string {
	var verbs = make(map[string]string)
	verbs["yapacağım"] = "yapmanı"
	verbs["gideceğim"] = "gitmeni"
	verbs["alacağım"] = "almanı"
	verbs["yazacağım"] = "yazmanı"
	verbs["okuyacağım"] = "okumanı"
	verbs["arayacağım"] = "aramanı"
	verbs["izleyeceğim"] = "izlemeni"
	verbs["dinleyeceğim"] = "dinlemeni"
	return verbs[verb]
}

//hatırlatıcı için zaman mapping
func timeSlover(time string) string {
	var times = make(map[string]string)
	times["sabah"] = "09:00"
	times["akşam"] = "19:00"
	times["öğlen"] = "12:00"
	times["gece"] = "22:00"
	if times[time] == "" {
		return "null"
	}
	return times[time]
}

//geri dönütlere göre mesajları döndüren fonksyon
func messageSender(chatid int, msg, FirstName string) {
	if message, typeof, file := mainOrganizer(msg, FirstName); typeof == "text" {

		bot.SendMessage(chatid, message, nil)

	} else if typeof == "404" {
		bot.SendMessage(chatid, "Aradığın filmi 🎬 bulamadım "+sadEmoji[rand.Intn(len(sadEmoji)-1)]+" yada böyle bir film yok 🚫", nil)

	} else if typeof == "t-mp3" {
		bot.SendMessage(chatid, message, nil)
		bot.SendVoice(chatid, "./sounds/"+file+".ogg", nil)

	} else if typeof == "mp3" {
		bot.SendVoice(chatid, "./sounds/"+file+".ogg", nil)

	} else if typeof == "test" {
		wikipedia.GetWiki("ankara")
	} else {
		var a tgbot.SendPhotoOptional
		a.Caption = &message
		bot.SendPhoto(chatid, file, &a)
	}
}

//inline
func processNewInlineQuery(bot *tgbot.TeleBot) {
	newQuery := bot.InlineQuerys
	for {
		q := <-newQuery
		fmt.Printf("Get NewInlineQuery:%#v \n", q)
		if q.Query != "" { // Only return result when query string not empty.
			result1 := types.NewInlineQueryResultArticl()
			result1.Id = "1"
			result1.Title = "Example"
			result1.MessageText = "Hi" + q.Query
			_, err := bot.AnswerInlineQuery(q.Id, []interface{}{result1}, nil)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

//wikipedia araması için örnek kod
func search(search string) string {

	var without string
	if strings.Contains(search, "nedir") {
		without = strings.ReplaceAll(search, "nedir", "")
	} else {
		without = strings.ReplaceAll(search, "kimdir", "")
	}

	regular := strings.TrimSpace(without)
	regular = strings.ReplaceAll(regular, " ", "_")

	result, status := wikipedia.GetWiki(regular)
	if !status {
		result = strings.ReplaceAll(result, "&&&&", without) + sadEmoji[rand.Intn(len(sadEmoji)-1)]
	}
	return result
}
