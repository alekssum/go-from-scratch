package main

import (
	"encoding/xml"
	"fmt"
)

func main() {

	res := YMLCatalogTree{}
	xml.Unmarshal([]byte(xmlString), &res)

	fmt.Println(res)

}

// YMLCatalogTree is ...
type YMLCatalogTree struct {
	Shop shop `xml:"shop"`
}

type shop struct {
	Name       string     `xml:"name"`
	Company    string     `xml:"company"`
	URL        string     `xml:"url"`
	City       string     `xml:"city"`
	Sklads     sklads     `xml:"sklads"`
	Categories categories `xml:"categories"`
	Offers     offers     `xml:"offers"`
}

type sklads struct {
	Store []store `xml:"store"`
}

type store struct {
	ID       string   `xml:"id,attr"`
	Phone    string   `xml:"phone,attr"`
	Lat      string   `xml:"lat,attr"`
	Lon      string   `xml:"lon,attr"`
	Schedule string   `xml:"schedule,attr"`
	Address  string   `xml:"address,attr"`
	Workweek workweek `xml:"work-week"`
}

type workweek struct {
	Workday []workday `xml:"work-day"`
}

type workday struct {
	Weekday string `xml:"week-day,attr"`
	Start   string `xml:"start,attr"`
	End     string `xml:"end,attr"`
}

type categories struct {
	Category []category `xml:"category"`
}

type category struct {
	ID   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type offers struct {
	Offer []offer `xml:"offer"`
}

type offer struct {
	ID         string  `xml:"id,attr"`
	Available  bool    `xml:"available,attr"`
	GroupID    string  `xml:"group_id,attr"`
	Price      float64 `xml:"price"`
	CurrencyID string  `xml:"currencyId"`
	CategoryID string  `xml:"categoryId"`
	Store      struct {
		ID     string  `xml:"id,attr"`
		Amount float64 `xml:"amount,attr"`
	} `xml:"store"`
	Picture     string `xml:"picture"`
	Vendor      string `xml:"vendor"`
	VendorCode  string `xml:"vendorCode"`
	Name        string `xml:"name"`
	Description string `xml:"description"`
	Params      []struct {
		Name  string `xml:"name,attr"`
		Value string `xml:",chardata"`
	} `xml:"param"`
}

var xmlString string = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE yml_catalog SYSTEM "shops.dtd">
<yml_catalog date="2020-06-22 10:00">
	<shop>
		<name>
			Драгоценная орхидея
		</name>
		<company>
			Драгоценная орхидея
		</company>
		<url>
			Драгоценная орхидея
		</url>
		<city>
			Пермь
		</city>
		<currencies>
			<currency id="RUR" rate="1" plus="0" />
		</currencies>
		<sklads>
			<store id="01" phone="+79655741916" lat="58.013015" lon="56.233425" schedule="ежедневно 11.00-21.00" address="г.Пермь, Тополевый переулок, 5">
				<work-week>
					<work-day week-day="mo" start="11:00" end="21:00" />
					<work-day week-day="tu" start="11:00" end="21:00" />
					<work-day week-day="we" start="11:00" end="21:00" />
					<work-day week-day="th" start="11:00" end="21:00" />
					<work-day week-day="fr" start="11:00" end="21:00" />
					<work-day week-day="sa" start="11:00" end="21:00" />
					<work-day week-day="su" start="11:00" end="21:00" />
				</work-week>
			</store>
		</sklads>
		<categories>
			<category id="000000023">
				Кольца
			</category>
			<category id="000000037">
				Серьги
			</category>
			<category id="000000041">
				Браслеты
			</category>
			<category id="000000043">
				Колье
			</category>
			<category id="000000033">
				Подвески
			</category>
			<category id="000000046">
				Мужская коллекция
			</category>
			<category id="000000048">
				Броши
			</category>
			<category id="000000087">
				Сувениры
			</category>
			<category id="000000045">
				Пирсинг
			</category>
			<category id="000000042">
				Цепи
			</category>
		</categories>
		<offers>
			<offer id="UT008200055750639271" available="true" group_id="00012562-6_УТ0001372">
				<price>
					3950
				</price>
				<currencyId>
					RUR
				</currencyId>
				<categoryId>
					000000023
				</categoryId>
				<store id="01" amount="1.00" />
				<picture>
					https://permgold.ru/assets/upload/images/bone/УТ00820005.jpg
				</picture>
				<vendor>
					АПАРТ
				</vendor>
				<vendorCode>
					00012562-6
				</vendorCode>
				<name>
					Кольцо (925), 00012562-6
				</name>
				<description />
				<param name="Тип Изделия">
					Кольца
				</param>
				<param name="Вставка">
					Без вставки
				</param>
				<param name="Металл">
					Нет
				</param>
				<param name="Размер">
					18,0
				</param>
			</offer>
			<offer id="UT008200015750639268" available="true" group_id="00012563-6л_УТ0001372">
				<price>
					2950
				</price>
				<currencyId>
					RUR
				</currencyId>
				<categoryId>
					000000023
				</categoryId>
				<store id="01" amount="1.00" />
				<picture>
					https://permgold.ru/assets/upload/images/bone/УТ00820001.jpg
				</picture>
				<vendor>
					АПАРТ
				</vendor>
				<vendorCode>
					00012563-6л
				</vendorCode>
				<name>
					Кольцо (925 п), 00012563-6л
				</name>
				<description />
				<param name="Тип Изделия">
					Кольца
				</param>
				<param name="Вставка">
					Без вставки
				</param>
				<param name="Металл">
					Нет
				</param>
				<param name="Размер">
					17,5
				</param>
			</offer>	
		</offers>
	</shop>
</yml_catalog>`
