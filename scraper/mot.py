# ministry of energy
from scraper.scraper import WebScraper

class MOTScraper(WebScraper):
    def perform_scraping(self):
        pass

    def run(self):
        try:
            self.navigate_to_url()
            self.perform_scraping()
        except Exception as error:
            print(error)