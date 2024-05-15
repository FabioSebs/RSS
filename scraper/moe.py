# ministry of energy
from scraper.scraper import WebScraper
from selenium.webdriver.common.by import By
import time
from generator.generate_rss import RSSGenerator


class MOEScraper(WebScraper):
    def perform_scraping(self):
        reports = self.driver.find_elements(By.CSS_SELECTOR, "div.list-berita div.col-md-4 div.berita-item")
        rssGen = RSSGenerator()
        
        # Main RSS feed setup
        rssGen.feed.title("Ministry of Energy")
        rssGen.feed.link(href="https://www.esdm.go.id/id/publikasi/lain-lain", rel="self")
        
        # Get the logo from the first report if available
        if reports:
            logo_url = reports[0].find_element(By.CSS_SELECTOR, "img").get_attribute("src")
            rssGen.feed.logo(logo_url)
        
        rssGen.feed.subtitle("reports")
        rssGen.feed.description("ministry of energy")
        
        for report in reports:
            try:
                fe = rssGen.feed.add_entry()
                
                # Entry details
                fe.title(report.find_element(By.CSS_SELECTOR, "h4.title").text)
                
                link_element = report.find_element(By.CSS_SELECTOR, "a.btn-download")
                fe.link(href=link_element.get_attribute("href"), rel="self")
                
                # logo_element = report.find_element(By.CSS_SELECTOR, "img")
                # fe.logo(logo_element.get_attribute("src"))
                
                # fe.subtitle(report.find_element(By.CSS_SELECTOR, "h4.title").text)  # Assuming subtitle should be the title text
                fe.description("ministry of energy")
                
                # Add the entry to the feed
                rssGen.feed.add_entry(fe)
            except Exception as e:
                print(f"An error occurred: {e}")
                continue

        rssGen.feed.rss_file("moe.xml")

    def run(self):
        try:
            self.navigate_to_url()
            self.perform_scraping()
        except Exception as error:
            print(error)
