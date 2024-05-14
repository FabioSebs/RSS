from generator.generate_rss import RSSGenerator
from fastapi import FastAPI, HTTPException, Response
from typing import Dict
import xml.etree.ElementTree as ET
import os
from scraper.moe import MOEScraper

app = FastAPI()

filename = ""

@app.get("/")
async def root():
    return {"message": "Hello World"}

@app.get("/xml/moe", response_class=Response)
async def read_xml():
    scraper = MOEScraper(fname="moe", url="https://www.esdm.go.id/")
    # Logic for xml, scraping, and serving xmx
    
    # Define the path to the XML file
    file_path = f"moe.xml"
    
    # Check if the file exists
    if not os.path.exists(file_path):
        raise HTTPException(status_code=404, detail="File not found")
    
    # Read the content of the XML file
    with open(file_path, "r") as file:
        xml_content = file.read()
    
    # Return the content as an XML response
    return Response(content=xml_content, media_type="application/xml")

@app.get("/xml/mot", response_class=Response)
async def read_xml():
    scraper = MOEScraper(fname="mot", url="https://www.dephub.go.id")
    # Logic for xml, scraping, and serving xmx
    
    # Define the path to the XML file
    file_path = f"mot.xml"
    
    # Check if the file exists
    if not os.path.exists(file_path):
        raise HTTPException(status_code=404, detail="File not found")
    
    # Read the content of the XML file
    with open(file_path, "r") as file:
        xml_content = file.read()
    
    # Return the content as an XML response
    return Response(content=xml_content, media_type="application/xml")


# Run the application using `uvicorn`
if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="127.0.0.1", port=8000)