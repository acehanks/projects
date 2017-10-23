# -*- coding: utf-8 -*-
import scrapy


class TripSpider(scrapy.Spider):
    name = "trip"
    #allowed_domains = ["www.tripadvisor.com.my/"]
    start_urls = (
        'https://www.tripadvisor.com.my/Restaurants-g60763-c11-New_York_City_New_York.html/',
    )
    download_delay= 1.5

    def parse(self, response):
        page = response.css('h3.title a::attr(href)').extract()

        for url in page:
        	ab_url= response.urljoin(url)
        	#yield{'URL': ab_url}
        	yield scrapy.Request(ab_url, callback= self.parse_page)

        next_page= response.css('div.unified.pagination.js_pageLinks a::attr(href)')[1].extract()
        #next_page= response.xpath('//a[text()="Next"]/@href').extract()
        if next_page is not None:
            next_page = response.urljoin(next_page)
            yield scrapy.Request(next_page, callback=self.parse)	


    def parse_page(self, response):

    	
    	name= response.xpath('//h1/text()').extract_first(),
    	rating= response.css('div.rs.rating').xpath('//span/img/@alt')[4].extract(),
    	street= response.css('span.street-address::text').extract_first()
    	location= response.css('span.locality::text')[2].extract(),
    	city_side= response.css('div.name::text').extract_first()

    	yield{
    	'Resturant': name,
    	'Rating': rating,
    	'Location': city_side,
    	'locality': location,
    	'Street': street
    	}
    	
