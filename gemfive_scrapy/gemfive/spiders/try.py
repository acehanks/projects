# -*- coding: utf-8 -*-
import scrapy
from scrapy.spiders import CrawlSpider, Rule
from scrapy.linkextractors import LinkExtractor


class TrySpider(CrawlSpider):
    name = "try"
    #allowed_domains = ["http://www.gemfive.com/my/mobile-phones"]
    start_urls = (
        'http://www.gemfive.com/my/mobile-phones',
    )
    download_delay= 5.0


    rules = [
        
        #Rule(LinkExtractor(allow=['http://www.gemfive.com/my/mobile-phones\?p=\w*']), callback='parse'),
        Rule(LinkExtractor(allow=['http://www.gemfive.com/my/mobile-phones\?p=\w*']), callback= 'parse_item', follow=True)
    ]


    def parse_item(self, response):

    	for info in response.css('div.product-info'):
    		yield{
    		'price': info.css('span.price span::text').extract_first(),
    		'brand': info.css('div.product-brand::text').extract_first(),
    		'phone': info.css('h2.product-name a::attr(title)').extract_first(),
    		}

'''
next_page = response.css('li.next a::attr(href)').extract_first()
if next_page is not None:
    next_page = response.urljoin(next_page)
    yield scrapy.Request(next_page, callback=self.parse)
'''

    