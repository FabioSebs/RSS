# RSS Feed Repository

**Mission** - Create an RSS Feed for updating news on key ministries regarding these topics

*English*

>vehicle |
>car |
>mobile |
>transport |
>electric | 
>hybrid |
>battery |
>policy |
>self-driving |
>regulation |
>law |
>legislation | 
>standard |
>initiative |
>emission |
>pollution |
>carbon |
>co2 |
>greenhouse | 
>sustainability | 
>environment |
>climate | 
>renewable |
>energy |
>efficiency |
>motorcycle |
>bus |
>trucks | 

*Indonesia*

> kendaraan |
> mobil |
> seluler | 
> transportasi |
> listrik |
> hibrida |
> baterai |
> kebijakan |
> self-driving |
> regulasi | 
> hukum |
> legislasi | 
> standar | 
> inisiatif |
> emisi |
> polusi |
> karbon |
> co2 |
> rumah kaca |
> keberlanjutan |
> lingkungan |
> iklim |
> terbarukan |
> energi |
> efisiensi | 
> sepeda motor |
> bus |
> truk |

---

## Information

**RSS** - a technology used to subscribe to web content, such as *blogs*, *news* *sites*, *podcasts*, and other *online publications*. Instead of visiting each website individually to check for new updates, users can subscribe to RSS feeds, which are *XML files* containing summaries or full texts of articles, along with metadata like publication date and authorship

**XML** - It's a markup language similar to HTML, but *whereas HTML is designed to display data, XML is designed to store and transport data*. XML is widely used for storing and exchanging structured data between different systems, such as web services, databases, and applications

**OPML** - OPML is a way to bundle many RSS feeds in one file to easily share among others !

## Challenges

- most webistes don't natively support RSS anymore :(
- constant update / scraping of webistes need to be performed
- email alerts are not supported 
- need to translate indonesian to english and vice versa!


## Solutions

- using a rss reader like Feedly!
- creating a scraper that generates the xml for every news source!
- create a scheduler that updates the xml on a timely basis!
- creating an email service with all members list in an excel file!
- using google translate api to translate our news titles!

## Methodology

**Feedly** - this will be the RSS reader of choice unless we choose to have a self-hosted one. Simply upload the OPML file into Feedly and should see all of the feeds! 


**WebScraper** - The scraper can go into any website that has **wanted information** such as 
- news articles
- posts
- publications
- etc

> Once I am given a url I can code the neccessary program to get the information and generate it into XML

**Generator** - RSS 2.0 is a standard that has been around longer than decades. I am following their standards when generating XML. Golang is being used to generate the xml.

## Live Demo

I am using my server so far as the web server to expose the RSS to Feedly. Here is a video demonstration!

- One stop place for our web data!
- Collobaritively view the web data!
- Append more sites for more web data!

> Example of the RSS Endpoint: https://fabrzy.dev/icct/v1/icct/moe/xml

### LINKS TO NEWS SOURCES

- https://en.baochinhphu.vn/policies.htm (vietnam)
- https://en.antaranews.com/latest-news (jakarta)
- https://apnews.com/hub/thailand-government (thailand)

