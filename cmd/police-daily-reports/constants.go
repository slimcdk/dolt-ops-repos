package main

const (
	DoltRepo    string = "slimcdk /danish-daily-police-reports"
	APIRoot     string = "https://politi.dk/api/"
	APIEndpoint string = APIRoot + "news/getNewsResults"
)

var Districts []string = []string{
	"Bornholms-politi",
	"Fyns-Politi",
	"Koebenhavns-Vestegns-Politi",
	"Midt-og-Vestsjaellands-Politi",
	"Nordjyllands-Politi",
	"Nordsjaellands-Politi",
	"Sydsjaellands-og-Lolland-Falsters-Politi",
	"OEstjyllands-Politi",
}
