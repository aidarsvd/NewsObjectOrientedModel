package main

import "fmt"

type News struct {
	status string
	totalResults int
	articles [3]NewsObjects
}
type NewsObjects struct {
	source Source
	author string
	title string
	description string
	url string
	urlToImage string
	publishedAt string
}


type Source struct {
	id string
	name string
}


func(article *News) showSource(indexOfArticle int){
	name := article.articles[indexOfArticle].source.name
	fmt.Printf("Name of the resource: %v\n", name)
}

func (article *News) showAuthor(indexOfArticle int) {
	author := article.articles[indexOfArticle].author
	fmt.Printf("Author of the article: %v\n", author)
}

func (article News) showTitle(indexOfArticle int) {
	title := article.articles[indexOfArticle].title
	fmt.Printf("Title of the article: %v\n", title)
}

func (article News) showDescription(indexOfArticle int) {
	descr := article.articles[indexOfArticle].description
	fmt.Printf("Description of the article: %v\n", descr)
}

func (article News) showDate(indexOfArticle int) {
	date := article.articles[indexOfArticle].publishedAt
	fmt.Printf("Article published at: %v\n", date)
}

func (size News)getLen()int{
	return len(size.articles)
}

func showAll(size int, article News) {
	for i:=0; i<size; i++{
		article.showSource(i)
		article.showAuthor(i)
		article.showTitle(i)
		article.showDescription(i)
		article.showDate(i)
	}
}

func main() {
	var article News
	article.status = "ok"
	article.totalResults = 32
	article.articles[0] = NewsObjects{
		source: Source{
			id:   "rbc",
			name: "RBC",
		},
		author:      "nil",
		title:       "Ростуризм упрекнули в отсутствии единой методики подсчета туристов - РБК",
		description: "Ростуризм должен разъяснить региональным властям, как правильно считать туристический поток. На местах чиновники могут завышать показатели, что не позволяет предпринимателям достоверно рассчитать доходность проектов",
		url:         "https://www.rbc.ru/business/09/12/2020/5fcf6a639a7947630700608f",
		urlToImage:  "https://s0.rbk.ru/v6_top_pics/media/img/8/01/756074900470018.jpg",
		publishedAt: "2020-12-09T06:24:00Z"}
	article.articles[1] = NewsObjects{
		source: Source{
			id:   "Ria.ru",
			name: "Ria.ru",
		},
		author:      "nil",
		title:       "Гинцбург уточнил рекомендации о запрете на алкоголь во время вакцинации - РИА НОВОСТИ",
		description: "Речь о полном запрете на алкоголь во время вакцинации от коронавируса не идет, но рекомендуется воздерживаться от него в течение трех дней после каждой... РИА Новости, 09.12.2020",
		url:         "https://ria.ru/20201209/vaktsinatsiya-1588332309.html",
		urlToImage:  "https://cdn24.img.ria.ru/images/sharing/article/1588332309.jpg?15875579151607496721",
		publishedAt: "2020-12-09T05:47:33Z"}
	article.articles[2] = NewsObjects{
		source: Source{
			id:   "Meduza",
			name: "Meduza",
		},
		author:      "nil",
		title:       "ПСЖ и «Истанбул» не доиграли матч Лиги чемпионов. Футболисты обвинили арбитра в расизме - Meduza",
		description: "Матч Лиги чемпионов между французским ПСЖ и турецким клубом «Истанбул Башакшехир», проходивший вечером 8 декабря в Париже, прервали после того, как игроки команды гостей обвинили одного из арбитров в расизме.",
		url:         "https://meduza.io/news/2020/12/09/pszh-i-istanbul-ne-doigrali-match-ligi-chempionov-futbolisty-obvinili-arbitra-v-rasizme",
		urlToImage:  "https://meduza.io/imgly/share/1607492817/news/2020/12/09/pszh-i-istanbul-ne-doigrali-match-ligi-chempionov-futbolisty-obvinili-arbitra-v-rasizme",
		publishedAt: "2020-12-09T05:30:00Z"}

	article.showSource(2)
	//showAll(article.getLen(), article)
}