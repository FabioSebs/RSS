# RSS Feed Repository

**Mission** - Create an RSS Feed for updating news on key ministries regarding these topics

- FE standards
- CO2 standard 
- Electric Vehicles 

---

**Ministries** 

- [Ministry of Energy and Mineral Resources](https://www.esdm.go.id/)
    - [Reports](https://www.esdm.go.id/id/publikasi/lain-lain)

- [Ministry of Transportation](https://www.dephub.go.id)
    - [Publications](https://www.dephub.go.id/post/kategori/publikasi-daftar-publikasi)
    - [Booklet](https://www.dephub.go.id/post/kategori/publikasi-booklet)
    - [News Letter](https://dephub.go.id/post/kategori/publikasi-newsletter)


## Information

**RSS** - a technology used to subscribe to web content, such as **blogs**, **news** **sites**, **podcasts**, and other **online publications**. Instead of visiting each website individually to check for new updates, users can subscribe to RSS feeds, which are **XML files** containing summaries or full texts of articles, along with metadata like publication date and authorship

**XML** - It's a markup language similar to HTML, but **whereas HTML is designed to display data, XML is designed to store and transport data**. XML is widely used for storing and exchanging structured data between different systems, such as web services, databases, and applications

**OPML** - OPML is a way to bundle many RSS feeds in one file to easily share among others 

## Challenges

**Problem** - the websites we need to get the XML from **don't support RSS** and will have to be generated periodicially! In order to do so we will be ....

**Solution** - creating a scraper that automatically generates the xml for all of the sites by parsing the data! Once generated the XML will be hosted on a web server for consumption by feedly. This will allow for a url to be given and read by our RSS Reader (Feedly)


## Methodology

**Feedly** - this will be the RSS reader of my choice unless we choose to have a self-hosted one. Simply upload the OPML file into Feedly and should see all of the feeds.


**WebScraper** - The scraper can go into any website that has **wanted information** such as 

- news articles
- posts
- publications
- etc

Once I am given a url I can code the neccessary program to get the information and generate it into XML

**Generator** - RSS 2.0 is a standard that has been around longer than decades. I am following their standards when generating XML. Golang is being used to generate the xml.

## Live Demo

I am using my server so far as the web server / url we need to feed to our RSS Reader (Feedly). I can add more to the feed and we can all be invited to the feed to view it collobaritvely if we get Feedly Enterprise. These are some of the benefits

- One stop place for our web data
- Collobaritively view the web data
- Append more sites for more web data

> Example of the RSS Endpoint: https://fabrzy.dev/icct/v1/icct/moe/xml

### LINKS TO NEWS SOURCES

- https://en.baochinhphu.vn/policies.htm (vietnam)
- https://en.antaranews.com/latest-news (jakarta)
- https://apnews.com/hub/thailand-government (thailand)

