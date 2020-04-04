package messages

var wellcomeMessage string = ` Hoş geldin. Kişisel asistanın olarak yapılacaklarını aklımda tutabilirim. /hatirlatici
Film önerisinde bulunabilir veya bir film hakkında detayları iletebilirim. /filmler
Günlük olarak bulunduğun şehirdeki hava durumu bilgilerini iletebilirim.  /havadurumu
istediğin bilgiyi wikipedia üzerinden araştırıp sana getirebilirim. /wiki 
buraya tekrar erişmek için nasıldı veya nasildi yazabilirsin.`
var reminderMessage string = `hatırlat, hatirlat, hatırla,hatirla kelimelerinden birini kullanan ki amacını anlayabileyim. 
devamında hangi gün ve hangi vakit ne yapacağını  söyle. Örnek olarak hatırlat cuma akşamı Ahmeti arayacağım. Not: İstanbul Türkçesi ile 
yazmanı tavsiye ederim. Tüm komutlardı görmek için nasıldı veya nasildi yazabilirsin.`

//Messages : önceden tanımlanış uzun metinleri döndüren bir fonksyon
func Messages(message string) string {
	switch message {
	case "wellcome":
		return wellcomeMessage
	case "reminder":
		return reminderMessage
	}
	return "bu konuda henüz yardımcı olamam"
}
