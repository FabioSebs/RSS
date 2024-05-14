# RSS Feed Repository

**Mission** - Create an RSS Feed for updating news on key ministries regarding these topics

- FE standards
- CO2 standard 
- Electric Vehicles 

---

**Ministries** ([List](https://www.dpr.go.id/en/index/link))

- [Ministry of Energy and Mineral Resources](https://www.esdm.go.id/)
- [Ministry of Transportation](https://www.dephub.go.id)


## Information

**RSS** - a technology used to subscribe to web content, such as **blogs**, **news** **sites**, **podcasts**, and other **online publications**. Instead of visiting each website individually to check for new updates, users can subscribe to RSS feeds, which are **XML files** containing summaries or full texts of articles, along with metadata like publication date and authorship

**XML** - It's a markup language similar to HTML, but **whereas HTML is designed to display data, XML is designed to store and transport data**. XML is widely used for storing and exchanging structured data between different systems, such as web services, databases, and applications


## Methodology

**RSS Reader** - for this project we will use Feedly! Users can easily subscribe to RSS feeds by entering the **URL of the website**. If this url is supported by feedly it automatically will generate the XML. Hoevever .....

**Problem** - the websites we need to get the XML from don't support RSS and will have to be generated periodicially and manually! In order to do so we will be ....

**Solution** - creating a scraper that automatically generates the xml for all of the sites by parsing the data! Once generated the XML will be hosted on an AWS EC2 server for consumption by feedly. This will allow for a url to be given 

