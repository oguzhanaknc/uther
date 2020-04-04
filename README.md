# Uther

Uther asistan şeklinde tasarlanmış bir telegram botu 
## Kurulum

Use the package manager [git](https://github.com/oguzhanaknc/uther.git) to install foobar.

```bash
git clone https://github.com/oguzhanaknc/github-blog.git
```
```go
go get -u ./...
```
## Kullanım

```go
go run ./main.go
```

## Açıklamlar
Uther main.go dosyası temel işlemeri yaparken yan klasörlerde api'lar tarafından gelen verileri getiren dosyları içerir. \
Komutlar: 
- /start : botu başlat
- /filmler : filmler hakkında yardım
- /hatırlatıcı : hatırlatıcı hakkında yardım
- /wiki : wikipedia üzerinde arama yardımı
- hatılat + [gün] + [zaman] + [iş] + [yüklem] : hatırlatıcı oluştur
- film + [film adı] : film hakkında bilgi sunar
- [parametre] + nedir / [parametre] + kimdir : wikipedia üzerinden arama yapar 
- hey - test mesajı gönderir
   

## License
[MIT](https://choosealicense.com/licenses/mit/)
